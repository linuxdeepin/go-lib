package operations

import (
	"fmt"
	"math"
	"pkg.deepin.io/lib/gio-2.0"
	"pkg.deepin.io/lib/timer"
	"strings"
)

const (
	_NsecPerMicrosec                      = 1000
	_SecondsNeededForReliableTransferRate = 15
	_MaximumDisplayedFileNameLength       = 50
)

func getTargetFileWithCustomName(src *gio.File, destDir *gio.File, destFsType string, sameFs bool, customName string) *gio.File {
	copyname := ""
	var dest *gio.File
	if customName != "" {
		copyname = customName
		copyname, _ = makeFileNameValidForDestFs(copyname, destFsType)
		dest, _ = destDir.GetChildForDisplayName(copyname)
	}

	if dest == nil && !sameFs {
		info, _ := src.QueryInfo(strings.Join(
			[]string{
				gio.FileAttributeStandardCopyName,
				gio.FileAttributeTrashOrigPath,
			}, ","),
			gio.FileQueryInfoFlagsNone,
			nil)
		if info != nil {
			copyname := ""
			if src.HasUriScheme("trash") {
				copyname = info.GetAttributeString(gio.FileAttributeTrashOrigPath)
			}

			if copyname == "" {
				info.GetAttributeString(gio.FileAttributeStandardCopyName)
			}

			if copyname != "" {
				copyname, _ = makeFileNameValidForDestFs(copyname, destFsType)
				dest, _ = destDir.GetChildForDisplayName(copyname)
			}

			info.Unref()
		}
	}

	if dest == nil {
		basename := src.GetBasename()
		basename, _ = makeFileNameValidForDestFs(basename, destFsType)
		dest = destDir.GetChild(basename)
	}

	return dest
}

func getTargetFile(src *gio.File, destDir *gio.File, destFsType string, sameFs bool) *gio.File {
	return getTargetFileWithCustomName(src, destDir, destFsType, sameFs, "")
}

// CopyMoveJob copy or move the files or directories
type CopyMoveJob struct {
	*CommonJob
	isMove            bool
	files             []*gio.File
	destination       *gio.File
	desktopLocation   *gio.File
	fakeDisplaySource *gio.File
	debutingFiles     map[string]bool
	targetName        string
	flags             gio.FileCopyFlags

	lastReportTime uint64
}

const (
	_CopyMoveSignalMoving            string = "moving"
	_CopyMoveSignalCopying           string = "copying"
	_CopyMoveSignalCopyingMovingDone string = "copying-moving-done"
	_CopyMoveSignalCreatingDir       string = "creating-dir"
)

// TODO: emitting moving/copying signal every 200ms,
// not emitting every file or directory (this would be too slow according to kio).
func (job *CopyMoveJob) emitMoving(srcURL string) {
	job.Emit(_CopyMoveSignalMoving, srcURL)
}

func (job *CopyMoveJob) emitCopying(srcURL string) {
	job.Emit(_CopyMoveSignalCopying, srcURL)
}

func (job *CopyMoveJob) emitCopyingMovingDone(srcURL string, destURL string) {
	job.Emit(_CopyMoveSignalCopyingMovingDone, srcURL, destURL)
}

func (job *CopyMoveJob) emitCreatingDir(srcURL string) {
	job.Emit(_CopyMoveSignalCreatingDir, srcURL)
}

func (job *CopyMoveJob) ListenMoving(fn func(string)) (func(), error) {
	return job.ListenSignal(_CopyMoveSignalMoving, fn)
}

func (job *CopyMoveJob) ListenCopying(fn func(string)) (func(), error) {
	return job.ListenSignal(_CopyMoveSignalCopying, fn)
}

func (job *CopyMoveJob) ListenCopyingMovingDone(fn func(string, string)) (func(), error) {
	return job.ListenSignal(_CopyMoveSignalCopyingMovingDone, fn)
}

func (job *CopyMoveJob) ListenCreatingDir(fn func(string)) (func(), error) {
	return job.ListenSignal(_CopyMoveSignalCreatingDir, fn)
}

func (job *CopyMoveJob) finalize() {
	for _, file := range job.files {
		file.Unref()
	}
	if job.destination != nil {
		job.destination.Unref()
	}
	if job.desktopLocation != nil {
		job.desktopLocation.Unref()
	}
	if job.fakeDisplaySource != nil {
		job.fakeDisplaySource.Unref()
	}
	job.CommonJob.finalize()
}

func (job *CopyMoveJob) getDest(
	src *gio.File,
	destDir *gio.File,
	sameFs bool,
	destFsType *string,
	uniqueName bool,
	uniqueNameNr *int) *gio.File {
	if uniqueName {
		dest := getUniqueTargetFile(src, destDir, sameFs, *destFsType, *uniqueNameNr)
		*uniqueNameNr++
		return dest
	} else if job.targetName != "" {
		return getTargetFileWithCustomName(src, destDir, *destFsType, sameFs, job.targetName)
	}

	return getTargetFile(src, destDir, *destFsType, sameFs)
}

func (job *CopyMoveJob) removeTargetRecursively(src *gio.File, toplevelDest *gio.File, file *gio.File) bool {
	var err error
	enumerator, err := file.EnumerateChildren(
		gio.FileAttributeStandardName,
		gio.FileQueryInfoFlagsNofollowSymlinks,
		job.cancellable)

	stop := false
	if enumerator != nil {
		err = nil
		var info *gio.FileInfo
		for !job.isAborted() {
			info, err = enumerator.NextFile(job.cancellable)
			if info == nil {
				break
			}
			child := file.GetChild(info.GetName())
			if !job.removeTargetRecursively(src, toplevelDest, child) {
				stop = true
				break
			}
			child.Unref()
			info.Unref()
		}

		enumerator.Close(job.cancellable)
		enumerator.Unref()
		// nautilus_file_changes_queue_file_removed(file)
		return true
	}

	gerr := err.(gio.GError)
	errCode := gio.IOErrorEnum(gerr.Code)
	if errCode != gio.IOErrorEnumNotDirectory || errCode != gio.IOErrorEnumCancelled {
		if !job.skipAllError {
			primaryText := Tr("Error while copying \"%B\"")                                    // src
			secondaryText := Tr("Could not remove files from the already existing folder %F.") //file
			detailText := err.Error()

			response := job.uiDelegate.AskSkip(primaryText, secondaryText, detailText, UIFlagsMulti)
			switch response.Code() {
			case ResponseCancel:
				job.Abort()
			case ResponseSkip:
				if response.ApplyToAll() {
					job.skipAllError = true
				}
			}
		}

		stop = true
	}

	if stop {
		return false
	}

	err = nil
	var ok bool

	// sorry to the files that is deleted, the undo operation won't work.
	ok, err = file.Delete(job.cancellable)
	if ok {
		// nautilus_file_changes_queue_file_removed(file)
		return true
	}

	gerr = err.(gio.GError)
	errCode = gio.IOErrorEnum(gerr.Code)
	if job.skipAllError || errCode == gio.IOErrorEnumCancelled {
		return false
	}

	primaryText := Tr("Error while copying “%B”.")                        // src
	secondaryText := Tr("Could not remove the already existing file %F.") // file
	detailText := err.Error()

	response := job.uiDelegate.AskSkip(primaryText, secondaryText, detailText, UIFlagsMulti)
	switch response.Code() {
	case ResponseCancel:
		job.Abort()
	case ResponseSkip:
		if response.ApplyToAll() {
			job.skipAllError = true
		}
	}

	return false
}

type _CreateDestState int

const (
	_CreateDestStateRetry _CreateDestState = iota
	_CreateDestStateFailed
	_CreateDestStateSuccess
)

func (job *CopyMoveJob) createDestDir(src *gio.File, destW *GFileWrapper, sameFs bool, destFsType *string) _CreateDestState {
	handledInvalidName := *destFsType != ""
retry:
	job.emitCreatingDir(destW.GetUri())
	ok, err := destW.MakeDirectory(job.cancellable)
	if ok {
		// TODO: undo
		// nautilus_file_changes_queue_file_added(info.dest)
		return _CreateDestStateSuccess
	}
	gerr := err.(gio.GError)
	errCode := gio.IOErrorEnum(gerr.Code)
	if errCode == gio.IOErrorEnumCancelled {
		return _CreateDestStateFailed
	} else if errCode == gio.IOErrorEnumInvalidFilename && !handledInvalidName {
		handledInvalidName = true
		destDir := destW.GetParent()

		if destDir != nil {
			*destFsType = queryFsType(destDir, job.cancellable)
			newDest := getTargetFile(src, destDir, *destFsType, sameFs)
			destDir.Unref()

			if !destW.Equal(newDest) {
				destW.Reset(newDest)
				return _CreateDestStateSuccess
			}
			newDest.Unref()
		}
	}

	primaryText := Tr("")
	secondaryText := ""
	detailText := ""

	if errCode == gio.IOErrorEnumPermissionDenied {
		secondaryText = Tr("The folder \"%B\" cannot be copied because you do not have permission to create it in the destination") // src
	} else {
		secondaryText = Tr("There was an error creating the folder \"%B\"") // src
		detailText = err.Error()
	}

	response := job.uiDelegate.AskSkip(primaryText, secondaryText, detailText, UIFlagsRetry)
	switch response.Code() {
	case ResponseCancel:
		job.Abort()
	case ResponseRetry:
		goto retry
	}

	return _CreateDestStateFailed
}

func (job *CopyMoveJob) copyMoveDirectory(
	src *gio.File,
	destW *GFileWrapper,
	sameFs bool,
	createDest bool,
	parentDestFsType *string,
	skippedFile *bool,
	readonlySourceFs bool) bool {

	if createDest {
		switch job.createDestDir(src, destW, sameFs, parentDestFsType) {
		case _CreateDestStateRetry:
			return false
		case _CreateDestStateFailed:
			*skippedFile = true
			return false
		}

		if job.debutingFiles != nil {
			job.debutingFiles[destW.GetUri()] = true
		}
	}

	localSkippedFile := false
	destFsType := ""

	skipError := job.shouldSkipDir(src)

retry:
	enumerator, enumerateErr := src.EnumerateChildren(gio.FileAttributeStandardName,
		gio.FileQueryInfoFlagsNofollowSymlinks,
		job.cancellable,
	)

	if enumerator != nil {
		var err error
		var childInfo *gio.FileInfo
		for !job.isAborted() {
			if skipError {
				childInfo, err = enumerator.NextFile(job.cancellable)
			} else {
				childInfo, _ = enumerator.NextFile(job.cancellable)
			}
			if childInfo == nil {
				break
			}

			srcFile := src.GetChild(childInfo.GetName())
			job.copyMoveFile(srcFile, destW.File, sameFs, false, &destFsType, false, &localSkippedFile, readonlySourceFs)

			srcFile.Unref()
			childInfo.Unref()
		}

		enumerator.Close(job.cancellable)
		enumerator.Unref()

		if err != nil {
			gerr := err.(gio.GError)
			errCode := gio.IOErrorEnum(gerr.Code)
			if errCode != gio.IOErrorEnumCancelled {
				primaryText := ""
				if job.isMove {
					primaryText = Tr("Error while moving.")
				} else {
					primaryText = Tr("Error while copying.")
				}

				secondaryText := ""
				detailText := ""

				if errCode == gio.IOErrorEnumPermissionDenied {
					secondaryText = Tr("Files in the folder \"%B\" cannot be copied because you do not have permissions to see them.") // src
				} else {
					secondaryText = Tr("There was an error getting information about the files in the folder \"%B\".") // src
					detailText = err.Error()
				}

				response := job.uiDelegate.AskSkip(primaryText, secondaryText, detailText, UIFlagsNone)
				switch response.Code() {
				case ResponseCancel:
					job.Abort()
				case ResponseSkip:
					localSkippedFile = true
				}
			}
		}

		job.setProcessedAmount(job.processedAmount[AmountUnitDirectories]+1, AmountUnitDirectories)
		job.reportCopyProgress()
		if job.debutingFiles != nil {
			job.debutingFiles[destW.GetUri()] = createDest
		}
	} else {
		gerr := enumerateErr.(gio.GError)
		errCode := gio.IOErrorEnum(gerr.Code)
		primaryText := ""
		if errCode != gio.IOErrorEnumCancelled {
			if job.isMove {
				primaryText = Tr("Error while moving")
			} else {
				primaryText = Tr("Error while copying")
			}

			secondaryText := ""
			detailText := ""

			if errCode == gio.IOErrorEnumPermissionDenied {
				secondaryText = Tr("The folder “%B” cannot be copied because you do not have permissions to read it.") //, src);
			} else {
				secondaryText = Tr("There was an error reading the folder “%B”.") //, src);
				detailText = gerr.Message
			}

			response := job.uiDelegate.AskSkip(primaryText, secondaryText, detailText, UIFlagsRetry)
			switch response.Code() {
			case ResponseCancel:
				job.Abort()
			case ResponseSkip:
				localSkippedFile = true
			case ResponseRetry:
				goto retry
			}
		}
	}

	if createDest {
		flags := gio.FileCopyFlagsNofollowSymlinks
		if readonlySourceFs {
			flags |= gio.FileCopyFlagsTargetDefaultPerms
		}

		src.CopyAttributes(destW.File, flags, job.cancellable)
	}

	if !job.isAborted() && job.isMove && !localSkippedFile {
		ok, err := src.Delete(job.cancellable)
		if !ok && !job.skipAllError {
			primaryText := Tr("Error while moving \"%B\"") //src
			secondaryText := Tr("Could not remove the source directory")
			detailText := err.Error()

			response := job.uiDelegate.AskSkip(primaryText, secondaryText, detailText, UIFlagsMulti)
			switch response.Code() {
			case ResponseSkip:
				localSkippedFile = true
				if response.ApplyToAll() {
					job.skipAllError = true
				}
			case ResponseCancel:
				job.Abort()
			}
		}
	}

	if localSkippedFile {
		*skippedFile = true
	}

	return true
}

func (job *CopyMoveJob) needRetry(
	gerr gio.GError,
	src *gio.File,
	destDir *gio.File,
	destW *GFileWrapper,
	overwrite *bool,
	uniqueName bool,
	uniqueNameNr *int,
	handledInvalidName *bool,
	skippedFile *bool,
	destFsType *string,
	sameFs *bool,
	readonlySourceFs bool) bool {

	errCode := gio.IOErrorEnum(gerr.Code)
	if !*handledInvalidName && errCode == gio.IOErrorEnumInvalidFilename {
		*handledInvalidName = true
		*destFsType = queryFsType(destDir, job.cancellable)

		var newDest *gio.File
		if uniqueName {
			newDest = getUniqueTargetFile(src, destDir, *sameFs, *destFsType, *uniqueNameNr)
		} else {
			newDest = getTargetFile(src, destDir, *destFsType, *sameFs)
		}

		if !destW.Equal(newDest) {
			destW.Reset(newDest)
			return true
		}
		newDest.Unref()
	}

	// conflict
	if !*overwrite && errCode == gio.IOErrorEnumExists {
		fmt.Println("file existing, get unique name?", uniqueName)
		if uniqueName {
			destW.Reset(getUniqueTargetFile(src, destDir, *sameFs, *destFsType, *uniqueNameNr))
			(*uniqueNameNr)++
			return true
		}

		isMerge := false
		if isDir(destW.File) && isDir(src) {
			isMerge = true
		}

		if (isMerge && job.mergeAll) || (!isMerge && job.replaceAll) {
			*overwrite = true
			return true
		}

		fmt.Println("skip all conflict?", job.skipAllConflict)
		if job.skipAllConflict {
			return false
		}

		// TODO:
		response := job.uiDelegate.ConflictDialog()
		fmt.Println(response)
		switch response.Code() {
		case ResponseCancel:
			job.Abort()
		case ResponseSkip:
			if response.ApplyToAll() {
				job.skipAllConflict = true
			}
		case ResponseOverwrite:
			if response.ApplyToAll() {
				if isMerge {
					job.mergeAll = true
				} else {
					job.replaceAll = true
				}
			}
			*overwrite = true
			return true
		case ResponseAutoRename:
			newDest := getUniqueTargetFile(src, destDir, *sameFs, *destFsType, *uniqueNameNr)
			(*uniqueNameNr)++
			destW.Reset(newDest)
			return true
		}
	} else if *overwrite && errCode == gio.IOErrorEnumIsDirectory {
		if job.removeTargetRecursively(src, destW.File, destW.File) {
			return true
		}
	} else if errCode == gio.IOErrorEnumWouldRecurse || errCode == gio.IOErrorEnumWouldMerge {
		isMerge := errCode == gio.IOErrorEnumWouldMerge
		wouldRecurse := errCode == gio.IOErrorEnumWouldRecurse

		/* Copying a dir onto file, first remove the file */
		if *overwrite && wouldRecurse {
			/* Copying a dir onto file, first remove the file */
			ok, err := destW.Delete(job.cancellable)
			gerr := err.(gio.GError)
			errCode := gio.IOErrorEnum(gerr.Code)
			if !ok && errCode != gio.IOErrorEnumNotFound {
				if job.skipAllError {
					return false
				}

				primaryText := ""
				if job.isMove {
					primaryText = Tr("Error while moving \"%B\"") // src
				} else {
					primaryText = Tr("Error while copying \"%B\"") // src
				}
				secondaryText := Tr("Could not remove the already existing file with the same name in %F.") //, dest_dir
				detailText := err.Error()

				response := job.uiDelegate.AskSkip(primaryText, secondaryText, detailText, UIFlagsMulti)
				switch response.Code() {
				case ResponseCancel:
					job.Abort()
				case ResponseSkip:
					if response.ApplyToAll() {
						job.skipAllError = true
					}
				}
				return false
			}

			// nautilus_file_changes_queue_file_removed(dest)
		}

		if isMerge {
			/* On merge we now write in the target directory, which may not
			   be in the same directory as the source, even if the parent is
			   (if the merged directory is a mountpoint). This could cause
			   problems as we then don't transcode filenames.
			   We just set same_fs to FALSE which is safe but a bit slower. */
			*sameFs = false
		}

		if !job.copyMoveDirectory(src, destW, *sameFs, wouldRecurse, destFsType, skippedFile, readonlySourceFs) {
			*handledInvalidName = true
			return true
		}
	} else if errCode != gio.IOErrorEnumCancelled {
		if job.skipAllError {
			return false
		}

		primaryText := fmt.Sprintf(Tr("Error while copying %s"), src)
		secondaryText := fmt.Sprintf(Tr("There was an error copying the file into %s"), destDir)
		detailText := gerr.Error()

		// TODO:
		// response = run_cancel_or_skip_warning (job,
		//                                        primaryText,
		//                                        secondaryText,
		//                                        detailText,
		//                                        source_info->num_files,
		//                                        source_info->num_files - transfer_info->num_files);
		response := job.uiDelegate.AskSkip(primaryText, secondaryText, detailText, UIFlagsMulti)
		switch response.Code() {
		case ResponseCancel:
			job.Abort()
		case ResponseSkip:
			if response.ApplyToAll() {
				job.skipAllError = true
			}
		}
	}

	return false
}

func (job *CopyMoveJob) reportCopyProgress() {
	now := timer.GetMonotonicTime().MicroSeconds()
	if job.lastReportTime != 0 &&
		math.Abs(float64(int64(job.lastReportTime)-now)) < 100*_NsecPerMicrosec {
		return
	}
	job.lastReportTime = uint64(now)
	job.emitProcessedAmount(job.processedAmount[AmountUnitBytes], AmountUnitBytes)
	job.emitProcessedAmount(job.processedAmount[AmountUnitFiles], AmountUnitFiles)
	job.emitProcessedAmount(job.processedAmount[AmountUnitDirectories], AmountUnitDirectories)
}

func newCopyFileProgressCallback(job *CopyMoveJob) gio.FileProgressCallback {
	var lastSize int64
	return func(currentNumBytes int64, totalNumBytes int64) {
		newSize := currentNumBytes - lastSize
		if newSize > 0 {
			job.setProcessedAmount(job.processedAmount[AmountUnitBytes]+newSize, AmountUnitBytes)
			lastSize = newSize
			job.reportCopyProgress()
		}
	}
}

// debutingFiles is not nil for toplevel items
func (job *CopyMoveJob) copyMoveFile(
	src *gio.File,
	destDir *gio.File,
	sameFs bool,
	uniqueName bool,
	destFsType *string,
	overwrite bool,
	skippedFile *bool,
	readonlySourceFs bool) {
	if job.shouldSkipFile(src) {
		fmt.Println("file should skip", src.GetUri())
		*skippedFile = true
		return
	}

	/* Don't allow recursive move/copy into itself.
	 * (We would get a file system error if we proceeded but it is nicer to
	 * detect and report it at this level) */
	if DirIsParentOf(src, destDir) {
		if job.skipAllError {
			*skippedFile = true
			return
		}

		primaryText := ""
		if job.isMove {
			primaryText = Tr("you cannot move a file into itself.")
		} else {
			primaryText = Tr("you cannot copy a file into itself.")
		}

		secondaryText := Tr("The destination folder is inside the source folder.")
		// TODO:
		response := job.uiDelegate.AskSkip(primaryText, secondaryText, "", UIFlagsMulti) /*, sourceInfo.NumFiles, sourceInfo.NumFiles-transferInfo.NumFiles*/
		switch response.Code() {
		case ResponseCancel:
			job.Abort()
		case ResponseSkip:
			if response.ApplyToAll() {
				job.skipAllError = true
			}
		}

		*skippedFile = true
		return
	}

	uniqueNameNr := 1
	dest := job.getDest(src, destDir, sameFs, destFsType, uniqueName, &uniqueNameNr)
	destW := NewGFileWrapper(dest)
	defer destW.Unref()

	/* Don't allow copying over the source or one of the parents of the source.
	 */
	if DirIsParentOf(dest, src) {
		if job.skipAllError {
			*skippedFile = true
			return
		}

		primaryText := ""
		if job.isMove {
			primaryText = Tr("You cannot move a file over itself.")
		} else {
			primaryText = Tr("You cannot copy a file over itself.")
		}
		secondaryText := Tr("The source file would be overwritten by the destination.")

		// TODO:
		response := job.uiDelegate.AskSkip(primaryText, secondaryText, "", UIFlagsMulti) /*source_info.num_files, source_info.num_files-transfer_info.num_files*/
		fmt.Println(primaryText, secondaryText, response)
		switch response.Code() {
		case ResponseCancel:
			job.Abort()
		case ResponseSkip:
			if response.ApplyToAll() {
				job.skipAllError = true
			}
		}

		*skippedFile = true
		return
	}

	handledInvalidName := *destFsType != ""
retry:
	flags := gio.FileCopyFlagsNofollowSymlinks
	if job.flags&gio.FileCopyFlagsNofollowSymlinks == 0 {
		flags = flags & ^gio.FileCopyFlagsNofollowSymlinks
	}
	if overwrite {
		flags |= gio.FileCopyFlagsOverwrite
	}
	if readonlySourceFs {
		flags |= gio.FileCopyFlagsTargetDefaultPerms
	}

	fmt.Println("job flags overwrite?", flags&gio.FileCopyFlagsOverwrite != 0)
	var err error
	var ok bool
	progressCb := newCopyFileProgressCallback(job)
	srcURL := src.GetUri()
	if job.isMove {
		job.emitMoving(srcURL)
		ok, err = src.Move(destW.File, flags, job.cancellable, progressCb)
	} else {
		job.emitCopying(src.GetUri())
		ok, err = src.Copy(destW.File, flags, job.cancellable, progressCb)
	}

	if ok {
		job.setProcessedAmount(job.processedAmount[AmountUnitFiles]+1, AmountUnitFiles)
		job.emitCopyingMovingDone(srcURL, destW.GetUri())
		job.reportCopyProgress()

		job.debutingFiles[srcURL] = true

		*skippedFile = false
		return
	}

	// this strange state will occur when the dest is nil.
	if err == nil {
		*skippedFile = true
		return
	}

	gerr := err.(gio.GError)
	if job.needRetry(gerr, src, destDir, destW, &overwrite, uniqueName, &uniqueNameNr, &handledInvalidName, skippedFile, destFsType, &sameFs, readonlySourceFs) {
		goto retry
	}

	*skippedFile = false
}

func (job *CopyMoveJob) copyFiles(destFsID string) {
	destFsType := ""

	job.reportCopyProgress()
	readonlyFs := isReadonlyFileSystem(job.files[0].GetParent())
	uniqueName := job.destination == nil
	for _, src := range job.files {
		if job.isAborted() {
			break
		}

		sameFs := false
		if destFsID != "" {
			sameFs = HasFsID(src, destFsID)
		}

		var dest *gio.File
		if job.destination != nil {
			dest = job.destination.Dup() // there is no Ref function, so Dup is used.
		} else {
			dest = src.GetParent()
		}

		if dest != nil {
			skippedFile := false
			job.copyMoveFile(src, dest, sameFs, uniqueName, &destFsType, false, &skippedFile, readonlyFs)
			dest.Unref()
		}
	}
}

func (job *CopyMoveJob) copyJob() {
	job.op = OpKindCopy
	job.scanSources(job.files)

	if job.isAborted() {
		fmt.Println("aborted copy job")
		return
	}

	var dest *gio.File
	if job.destination != nil {
		dest = job.destination.Dup() //Ref
	} else {
		dest = job.files[0].GetParent()
	}

	destFsID := job.verifyDestination(dest, uint64(job.totalAmount[AmountUnitBytes]))
	dest.Unref()

	if job.isAborted() {
		return
	}

	job.progressReporterTimer.Start()
	job.copyFiles(destFsID)
}

type _MoveFileCopyFallback struct {
	file      *gio.File
	overwrite bool
}

func newMoveCopyFileFallback(src *gio.File, overwrite bool) _MoveFileCopyFallback {
	var fallback _MoveFileCopyFallback
	fallback.file = src
	fallback.overwrite = overwrite
	return fallback
}

func getFilesFromFallbacks(fallbacks []_MoveFileCopyFallback) []*gio.File {
	files := []*gio.File{}
	for _, fallback := range fallbacks {
		files = append(files, fallback.file)
	}
	return files
}

func (job *CopyMoveJob) reportMoveProgress(total int, left int) {
	// TODO:
	// nautilus_progress_info_take_status (job->progress,
	// 				    f (_("Preparing to Move to “%B”"),
	// 				       move_job->destination));
	//
	// nautilus_progress_info_take_details (job->progress,
	// 				     f (ngettext ("Preparing to move %'d file",
	// 						  "Preparing to move %'d files",
	// 						  left), left));
	//
	// nautilus_progress_info_pulse_progress (job->progress);
}

func (job *CopyMoveJob) moveFilePrepare(
	src *gio.File,
	destDir *gio.File,
	sameFs bool,
	destFsType *string,
	fallbacks *[]_MoveFileCopyFallback,
	left int) {
	overwrite := false
	handledInvalidName := *destFsType != ""
	dest := getTargetFile(src, destDir, *destFsType, sameFs)

	/* Don't allow recursive move/copy into itself.
	 * (We would get a file system error if we proceeded but it is nicer to
	 * detect and report it at this level) */
	if DirIsParentOf(src, destDir) {
		if job.skipAllError {
			dest.Unref()
			return
		}

		primaryText := ""
		if job.isMove {
			primaryText = Tr("You cannot move a folder into itself.")
		} else {
			primaryText = Tr("You cannot copy a folder into itself.")
		}
		secondaryText := Tr("The destination folder is inside the source folder.")

		// response = run_warning (job,
		// 			primaryText,
		// 			secondaryText,
		// 			NULL,
		// 			files_left > 1,
		// 			CANCEL, SKIP_ALL, SKIP,
		// 			NULL);
		response := job.uiDelegate.AskSkip(primaryText, secondaryText, "", UIFlagsMulti)
		switch response.Code() {
		case ResponseCancel:
			job.Abort()
		case ResponseSkip:
			if response.ApplyToAll() {
				job.skipAllError = true
			}
		}

		dest.Unref()
		return
	}

	uniqueNameNr := 0
retry:
	flags := gio.FileCopyFlagsNofollowSymlinks | gio.FileCopyFlagsNoFallbackForMove
	if overwrite {
		flags |= gio.FileCopyFlagsOverwrite
	}

	srcURL := src.GetUri()
	job.emitMoving(srcURL)
	ok, err := src.Move(dest, flags, job.cancellable, nil)
	if ok {
		job.debutingFiles[dest.GetUri()] = true
		job.setProcessedAmount(job.processedAmount[AmountUnitFiles]+1, AmountUnitFiles)
		job.emitCopyingMovingDone(srcURL, dest.GetUri())
		return
	}

	gerr := err.(gio.GError)
	errCode := gio.IOErrorEnum(gerr.Code)

	if errCode == gio.IOErrorEnumInvalidFilename && !handledInvalidName {
		handledInvalidName = true
		*destFsType = queryFsType(destDir, job.cancellable)
		newDest := getTargetFile(src, destDir, *destFsType, sameFs)
		if !dest.Equal(newDest) {
			dest.Unref()
			dest = newDest
			goto retry
		} else {
			newDest.Unref()
		}
	} else if !overwrite && errCode == gio.IOErrorEnumExists {
		isMerge := false
		if isDir(dest) && isDir(src) {
			isMerge = true
		}

		if (isMerge && job.mergeAll) || (!isMerge && job.replaceAll) {
			overwrite = true
			goto retry
		}

		if job.skipAllConflict {
			dest.Unref()
			return
		}

		response := job.uiDelegate.ConflictDialog()
		switch response.Code() {
		case ResponseCancel:
			job.Abort()
		case ResponseSkip:
			if response.ApplyToAll() {
				job.skipAllConflict = true
			}
		case ResponseOverwrite:
			if response.ApplyToAll() {
				if isMerge {
					job.mergeAll = true
				} else {
					job.replaceAll = true
				}
			}
			overwrite = true
			goto retry
		case ResponseAutoRename:
			dest.Unref()
			dest = (getUniqueTargetFile(src, destDir, sameFs, *destFsType, uniqueNameNr))
			(uniqueNameNr)++
			goto retry
		}
	} else if errCode == gio.IOErrorEnumWouldRecurse ||
		errCode == gio.IOErrorEnumWouldMerge ||
		errCode == gio.IOErrorEnumNotSupported ||
		(overwrite && errCode == gio.IOErrorEnumIsDirectory) {
		fallback := newMoveCopyFileFallback(src, overwrite)
		*fallbacks = append(*fallbacks, fallback)
	} else if errCode != gio.IOErrorEnumCancelled {
		if job.skipAllError {
			dest.Unref()
			return
		}

		primaryText := Tr(`Error while moving "%B".`)                      // src
		secondaryText := Tr(`There was an error moving the file into %F.`) // destDir
		detailText := err.Error()

		response := job.uiDelegate.AskSkip(primaryText, secondaryText, detailText, UIFlagsMulti)
		switch response.Code() {
		case ResponseCancel:
			job.Abort()
		case ResponseSkip:
			if response.ApplyToAll() {
				job.skipAllError = true
			}
		}
	}
}

// moves all files that we can do without copy + delete
func (job *CopyMoveJob) moveFilesPrepare(destFsID string, destFsType *string) (fallbacks []_MoveFileCopyFallback) {
	total := len(job.files)
	left := total

	job.reportMoveProgress(total, left)

	for _, src := range job.files {
		if job.isAborted() {
			break
		}

		sameFs := false

		if destFsID != "" {
			sameFs = HasFsID(src, destFsID)
		}

		job.moveFilePrepare(src, job.destination, sameFs, destFsType, &fallbacks, left)
		left--
		job.reportMoveProgress(total, left)
	}

	return
}

func (job *CopyMoveJob) moveFiles(
	fallbacks []_MoveFileCopyFallback,
	destFsID string,
	destFsType *string,
) {
	job.reportCopyProgress()

	for _, fallback := range fallbacks {
		src := fallback.file
		sameFs := false
		if destFsID != "" {
			sameFs = HasFsID(src, destFsID)
		}

		skippedFile := false
		job.copyMoveFile(src, job.destination, sameFs, false, destFsType, fallback.overwrite, &skippedFile, false)
	}
}

func (job *CopyMoveJob) moveJob() {
	destFsType := ""

	job.op = OpKindMove
	destFsID := job.verifyDestination(job.destination, uint64(job.totalAmount[AmountUnitBytes]))

	if job.isAborted() {
		fmt.Println("aborted move job")
		return
	}

	fallbacks := job.moveFilesPrepare(destFsID, &destFsType)

	if job.isAborted() {
		return
	}

	fallbackFiles := getFilesFromFallbacks(fallbacks)
	job.op = OpKindMove
	job.scanSources(fallbackFiles)

	if job.isAborted() {
		return
	}

	job.verifyDestination(job.destination, uint64(job.totalAmount[AmountUnitBytes]))

	if job.isAborted() {
		return
	}

	job.moveFiles(fallbacks, destFsID, &destFsType)
}

// Execute a move job or a copy job.
func (job *CopyMoveJob) Execute() {
	defer job.finalize()
	defer job.emitDone()

	jobName := "copy"
	if job.isMove {
		jobName = "move"
	}
	fmt.Printf("execute %s job\n", jobName)
	if job.isMove {
		job.moveJob()
	} else {
		job.copyJob()
	}
}

func newCopyMoveJob(srcs []*gio.File, destDir *gio.File, targetName string, flags gio.FileCopyFlags, uiDelegate IUIDelegate) *CopyMoveJob {
	job := &CopyMoveJob{
		CommonJob:     newCommon(uiDelegate),
		files:         srcs,
		destination:   destDir,
		isMove:        false,
		targetName:    targetName,
		debutingFiles: map[string]bool{},
		flags:         flags,
	}

	return job
}

func newCopyMoveJobFromURL(srcURLs []string, destDirURL string, targetName string, flags gio.FileCopyFlags, uiDelegate IUIDelegate) *CopyMoveJob {
	srcs := locationListFromUriList(srcURLs)
	destDir := gio.FileNewForCommandlineArg(destDirURL)

	return newCopyMoveJob(srcs, destDir, targetName, flags, uiDelegate)
}

// NewCopyJob creates a copy job.
func NewCopyJob(srcURLs []string, destDirURL string, targetName string, flags gio.FileCopyFlags, uiDelegate IUIDelegate) *CopyMoveJob {
	return newCopyMoveJobFromURL(srcURLs, destDirURL, targetName, flags, uiDelegate)
}

func markAsMoveJob(job *CopyMoveJob) *CopyMoveJob {
	isMove := true
	targetIsMapping := false
	haveNonmappingSource := false

	if job.destination.HasUriScheme("burn") {
		targetIsMapping = true
	}

	for _, src := range job.files {
		if !src.HasUriScheme("burn") {
			haveNonmappingSource = true
			break
		}
	}

	if targetIsMapping && haveNonmappingSource {
		/* never move to "burn:///", but fall back to copy.
		 * This is a workaround, because otherwise the source files would be removed.
		 */
		isMove = false
	}

	job.isMove = isMove
	return job
}

// NewMoveJob creates a move job.
func NewMoveJob(srcURLs []string, destDirURL string, targetName string, flags gio.FileCopyFlags, uiDelegate IUIDelegate) *CopyMoveJob {
	return markAsMoveJob(newCopyMoveJobFromURL(srcURLs, destDirURL, targetName, flags, uiDelegate))
}
