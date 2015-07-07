package operations

import (
	"pkg.deepin.io/lib/gio-2.0"
	"strings"
)

// DeleteJob delete or trash files/directories.
type DeleteJob struct {
	*CommonJob
	files         []*gio.File
	trash         bool
	shouldConfirm bool
	userCancel    bool
}

const (
	_DeleteJobSignalDeleting = "deleting"
	_DeleteJobSignalTrashing = "trashing"
)

func (job *DeleteJob) emitDeleting(deletedFileURL string) {
	job.Emit(_DeleteJobSignalDeleting, deletedFileURL)
}

func (job *DeleteJob) emitTrashing(trashingFile string) {
	job.Emit(_DeleteJobSignalTrashing, trashingFile)
}

func (job *DeleteJob) ListenDeleting(fn func(string)) (func(), error) {
	return job.ListenSignal(_DeleteJobSignalDeleting, fn)
}

func (job *DeleteJob) ListenTrashing(fn func(string)) (func(), error) {
	return job.ListenSignal(_DeleteJobSignalTrashing, fn)
}

// TODO: function is huge.
func (job *DeleteJob) deleteDirectory(
	dir *gio.File,
	skippedFile *bool,
	topLevel bool,
) {
	localSkippedFile := false

	skipError := job.shouldSkipDir(dir)

retry:
	enumerator, enumerateErr := dir.EnumerateChildren(strings.Join(
		[]string{
			gio.FileAttributeStandardName,
			gio.FileAttributeStandardSize,
		}, ","),
		gio.FileQueryInfoFlagsNofollowSymlinks,
		job.cancellable,
	)

	if enumerator != nil {
		var err error
		for !job.isAborted() {
			fileInfo, nextErr := enumerator.NextFile(job.cancellable)
			err = nextErr
			if fileInfo == nil {
				break
			}
			child := dir.GetChild(fileInfo.GetName())
			job.deleteFile(child, &localSkippedFile, false)
			fileInfo.Unref()
			child.Unref()
		}
		enumerator.Close(job.cancellable)
		enumerator.Unref()

		if !skipError && err != nil {
			errCode := gio.IOErrorEnum(err.(gio.GError).Code)
			if errCode != gio.IOErrorEnumCancelled {
				primaryText := Tr("error while deleting")
				secondaryText := ""
				detailText := ""
				if errCode == gio.IOErrorEnumPermissionDenied {
					secondaryText = Tr("permission denied")
				} else {
					secondaryText = Tr("error")
					detailText = err.Error()
				}

				response := job.uiDelegate.AskSkip(primaryText, secondaryText, detailText, UIFlagsMulti)
				switch response.Code() {
				case ResponseCancel:
					job.Abort()
				case ResponseSkip:
					if response.ApplyToAll() {
						localSkippedFile = true
					}
				default: // not reached
				}
			}
		}
	} else {
		errCode := gio.IOErrorEnum(enumerateErr.(gio.GError).Code)
		if errCode != gio.IOErrorEnumCancelled {
			primaryText := Tr("error while deleting")
			secondaryText := ""
			detailText := ""
			if errCode == gio.IOErrorEnumPermissionDenied {
				secondaryText = Tr("permission denied")
			} else {
				secondaryText = Tr("error")
				detailText = enumerateErr.Error()
			}

			response := job.uiDelegate.AskSkip(primaryText, secondaryText, detailText, UIFlagsMulti|UIFlagsRetry)
			switch response.Code() {
			case ResponseCancel:
				job.Abort()
			case ResponseSkip: // skip all, do nothing
				if response.ApplyToAll() {
					localSkippedFile = true
				}
			case ResponseRetry:
				goto retry // avoid recusive
			default: // not reached
			}
		}
	}

	if !job.isAborted() && !localSkippedFile {
		deletedDirURL := dir.GetUri()
		job.emitDeleting(deletedDirURL)
		ok, err := dir.Delete(job.cancellable)
		if ok {
			job.setProcessedAmount(job.processedAmount[AmountUnitDirectories]+1, AmountUnitDirectories)
			job.setProcessedAmount(job.processedAmount[AmountUnitSumOfFilesAndDirs]+1, AmountUnitSumOfFilesAndDirs)
			return
		}

		if !job.skipAllError {
			primaryText := Tr("error while deleting")
			secondaryText := Tr("could not remove the folder")
			detailText := err.Error()
			response := job.uiDelegate.AskSkip(primaryText, secondaryText, detailText, UIFlagsMulti)
			switch response.Code() {
			case ResponseCancel:
				job.Abort()
			case ResponseSkip:
				if response.ApplyToAll() {
					job.skipAllError = true
				}
				localSkippedFile = true
			}
		}
	}

	if localSkippedFile {
		*skippedFile = true
	}
}

func (job *DeleteJob) deleteFile(file *gio.File, skippedFile *bool, topLevel bool) {
	if job.shouldSkipFile(file) {
		*skippedFile = true
		return
	}
	deletedFileURL := file.GetUri()
	job.emitDeleting(deletedFileURL)
	ok, err := file.Delete(job.cancellable)
	if ok {
		job.setProcessedAmount(job.processedAmount[AmountUnitFiles]+1, AmountUnitFiles)
		job.setProcessedAmount(job.processedAmount[AmountUnitSumOfFilesAndDirs]+1, AmountUnitSumOfFilesAndDirs)
		return
	}

	errCode := gio.IOErrorEnum(err.(gio.GError).Code)
	if errCode == gio.IOErrorEnumNotEmpty {
		job.deleteDirectory(file, skippedFile, topLevel)
		return
	}

	*skippedFile = true

	if !job.skipAllError && errCode != gio.IOErrorEnumCancelled {
		primaryText := Tr("")
		secondaryText := Tr("")
		detailText := Tr("")
		response := job.uiDelegate.AskSkip(primaryText, secondaryText, detailText, UIFlagsMulti)
		switch response.Code() {
		case ResponseCancel:
			job.Abort()
		case ResponseSkip:
			// skip do nothing
			if response.ApplyToAll() {
				job.skipAllError = true
			}
		default: // not reached
		}
	}
}

func (job *DeleteJob) deleteFiles(files []*gio.File) int {
	nFilesSkipped := 0
	if job.isAborted() {
		return nFilesSkipped
	}

	job.op = OpKindDelete

	// all files that cannot be trashed sometimes will be deleted.
	if !job.trash {
		job.scanSources(files)
		job.processedAmount[AmountUnitSumOfFilesAndDirs] = job.processedAmount[AmountUnitFiles] + job.processedAmount[AmountUnitDirectories]
	}

	if job.isAborted() {
		return nFilesSkipped
	}

	for _, file := range files {
		skippedFile := false
		job.deleteFile(file, &skippedFile, true)

		if skippedFile {
			nFilesSkipped++
			job.skippedAmount[AmountUnitSumOfFilesAndDirs]++
			job.setProcessedAmount(job.processedAmount[AmountUnitSumOfFilesAndDirs]+1, AmountUnitSumOfFilesAndDirs)
		}
	}

	return nFilesSkipped
}

type _TrashState int

const (
	_TrashStateTrashed _TrashState = iota
	_TrashStateRetry
	_TrashStateSkip
	_TrashStateDelete
	_TrashStateCancel
)

func (job *DeleteJob) trashFile(file *gio.File) _TrashState {
	state := _TrashStateTrashed
	ok, err := file.Trash(job.cancellable)
	if ok {
		return state
	}

	skip := false
	if job.skipAllError {
		state = _TrashStateSkip
		skip = true
	}

	if job.deleteAll {
		state = _TrashStateDelete
		skip = true
	}

	if skip {
		return state
	}

	primaryText := Tr("trash failed, delete???")
	secondaryText := ""
	detailText := ""

	errCode := gio.IOErrorEnum(err.(gio.GError).Code)
	if errCode == gio.IOErrorEnumNotSupported {
		detailText = err.Error()
	} else if !file.IsNative() {
		secondaryText = Tr("")
	}

	response := job.uiDelegate.AskDelete(primaryText, secondaryText, detailText, UIFlagsMulti|UIFlagsRetry)

	switch response.Code() {
	case ResponseCancel:
		state = _TrashStateCancel
		job.userCancel = true
		job.Abort()
	case ResponseSkip:
		state = _TrashStateSkip
		if response.ApplyToAll() {
			job.skipAllError = true
		}
	case ResponseDelete:
		state = _TrashStateDelete
		if response.ApplyToAll() {
			job.deleteAll = true
		}
	case ResponseRetry:
		state = _TrashStateRetry
	default:
		state = _TrashStateSkip
	}

	return state
}

func (job *DeleteJob) trashFiles(files []*gio.File) int {
	nFilesSkipped := 0

	if job.isAborted() {
		return nFilesSkipped
	}

	totalFiles := len(files)
	job.totalAmount[AmountUnitSumOfFilesAndDirs] = int64(totalFiles)
	toDelete := []*gio.File{}

	for _, file := range files {
	retry:
		job.emitTrashing(file.GetUri())
		state := job.trashFile(file)
		if state != 0 {
			totalFiles--
			switch state {
			case _TrashStateRetry:
				goto retry
			case _TrashStateSkip:
				job.skippedAmount[AmountUnitSumOfFilesAndDirs]++
				job.setProcessedAmount(job.processedAmount[AmountUnitSumOfFilesAndDirs]+1, AmountUnitSumOfFilesAndDirs)
			case _TrashStateDelete:
				toDelete = append(toDelete, file)
			case _TrashStateCancel:
				// do nothing
			}
			continue
		}

		// files and directories has no difference on trashing job.
		job.setProcessedAmount(job.processedAmount[AmountUnitSumOfFilesAndDirs]+1, AmountUnitSumOfFilesAndDirs)
		continue
	}

	if len(toDelete) != 0 {
		job.deleteFiles(toDelete)
	}

	return nFilesSkipped
}

func (job *DeleteJob) canDeleteFilesWithoutConfirm(files []*gio.File) bool {
	for _, file := range files {
		if !job.canDeleteWithoutConfirm(file) {
			return false
		}
	}

	return true
}

func (job *DeleteJob) canDeleteWithoutConfirm(file *gio.File) bool {
	return file.HasUriScheme("burn") ||
		file.HasUriScheme("recent") ||
		file.HasUriScheme("x-nautilus-desktop")
}

func (job *DeleteJob) confirmDeleteFromTrash(files []*gio.File) bool {
	if !job.shouldConfirm {
		return true
	}

	fileNum := uint64(len(files))
	if fileNum == 0 {
		panic("file should be greater than 0")
	}

	primaryText := NTr("permanently delete from trash??", "", fileNum)
	secondaryText := Tr("")
	detailText := Tr("")
	response := job.uiDelegate.AskDeleteConfirmation(
		primaryText,
		secondaryText,
		detailText,
	)
	return response
}

func (job *DeleteJob) confirmDeleteDirectly(files []*gio.File) bool {
	if !job.shouldConfirm {
		return true
	}

	if job.canDeleteFilesWithoutConfirm(files) {
		return true
	}

	fileNum := uint64(len(files))
	// TODO: doc
	primaryText := NTr("delete???", "", fileNum)
	secondaryText := Tr("")
	detailText := Tr("")
	response := job.uiDelegate.AskDeleteConfirmation(
		primaryText,
		secondaryText,
		detailText,
	)
	return response
}

func (job *DeleteJob) init(files []*gio.File, uiDelegate IUIDelegate) {
	job.CommonJob = newCommon(uiDelegate)
	job.files = files
	job.userCancel = false
	job.progressUnit = AmountUnitSumOfFilesAndDirs

	// TODO:
	// if job.trash {
	// 	// inhibit power manager for trashing
	// } else {
	// 	// inhibit power manager for deleting
	// }
}

func (job *DeleteJob) finalize() {
	for _, file := range job.files {
		file.Unref()
	}
	job.CommonJob.finalize()
}

// Execute the delete job to delete files/directories.
func (job *DeleteJob) Execute() {
	defer job.finalize()
	defer job.emitDone()

	toTrashFiles := []*gio.File{}
	toDeleteFiles := []*gio.File{}

	confirmed := true
	skippedFileNum := 0
	mustConfirmDeleteInTrash := false
	mustConfirmDelete := false

	for _, file := range job.files {
		if job.trash && file.HasUriScheme("trash") {
			mustConfirmDeleteInTrash = true
			toDeleteFiles = append(toDeleteFiles, file)
		} else if job.canDeleteWithoutConfirm(file) {
			toDeleteFiles = append(toDeleteFiles, file)
		} else {
			if job.trash {
				toTrashFiles = append(toTrashFiles, file)
			} else {
				mustConfirmDelete = true
				toDeleteFiles = append(toDeleteFiles, file)
			}
		}
	}

	if len(toDeleteFiles) != 0 {
		confirmed = true
		if mustConfirmDeleteInTrash {
			confirmed = job.confirmDeleteFromTrash(toDeleteFiles)
		} else if mustConfirmDelete {
			confirmed = job.confirmDeleteDirectly(toDeleteFiles)
		}
		if confirmed {
			job.deleteFiles(toDeleteFiles)
		} else {
			job.userCancel = true
		}
	}

	if len(toTrashFiles) != 0 {
		skippedFileNum = job.trashFiles(toTrashFiles)
	}

	if skippedFileNum == len(job.files) {
		job.userCancel = true
	}

}

// create DeleteJob for delete files or trash files, using internally.
func newDeleteOrTrashJob(files []*gio.File, trash bool, shouldConfirm bool, uiDelegate IUIDelegate) *DeleteJob {
	job := &DeleteJob{
		trash:         trash,
		shouldConfirm: shouldConfirm,
	}
	job.init(files, uiDelegate)
	return job
}

// convenient function for creating delete or trash job.
func delOrTrash(urls []string, trash bool, shouldConfirm bool, uiDelegate IUIDelegate) *DeleteJob {
	files := make([]*gio.File, len(urls))
	for i, fileURL := range urls {
		files[i] = gio.FileNewForCommandlineArg(fileURL)
	}

	return newDeleteOrTrashJob(files, trash, shouldConfirm, uiDelegate)
}

// NewDeleteJob creates a new delete job to delete files or directories.
func NewDeleteJob(urls []string, shouldConfirm bool, uiDelegate IUIDelegate) *DeleteJob {
	return delOrTrash(urls, false, shouldConfirm, uiDelegate)
}

// NewTrashJob creates a new trash job to trash files or directories.
func NewTrashJob(urls []string, shouldConfirm bool, uiDelegate IUIDelegate) *DeleteJob {
	return delOrTrash(urls, true, shouldConfirm, uiDelegate)
}
