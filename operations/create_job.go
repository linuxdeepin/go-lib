package operations

import (
	"fmt"
	"path/filepath"
	"pkg.deepin.io/lib/gio-2.0"
	"strings"
	"unicode/utf8"
)

// TODO: add flags?? like CreateFlagAutoRename

type _TargetInfo struct {
	filename               string
	newFilename            string
	filenameIsUtf8         bool
	count                  int
	destFsType             string
	handledInvalidFilename bool
	maxLength              int
}

// CreateJob is used to create file, directory or symbol link.
// The uri of created file will be returned.
type CreateJob struct {
	*CommonJob
	destDir     *gio.File
	filename    string
	makeDir     bool
	makeLink    bool
	src         *gio.File
	srcData     []byte
	length      int
	createdFile *gio.File
	// position GtkPoint
	// hasPosition bool
	// done_callback
	// done_callback_data
}

func (job *CreateJob) finalize() {
	if job.createdFile != nil {
		job.createdFile.Unref()
	}

	if job.src != nil {
		job.src.Unref()
	}

	job.destDir.Unref()
	job.CommonJob.finalize()
}

func (job *CreateJob) getFilename(destFsType string) (string, bool) {
	filename := job.filename
	filenameIsUtf8 := false
	if filename != "" {
		filenameIsUtf8 = utf8.ValidString(filename)
	} else {
		if job.makeDir {
			// TODO
			filename = Tr("Untitled Folder") // TODO: doc
			filenameIsUtf8 = true
		} else {
			if job.src != nil {
				basename := job.src.GetBasename()
				filename = fmt.Sprintf(Tr("Untitled %s"), basename) // TODO: doc
			}

			if filename == "" {
				filename = Tr("Untitled Document") // TODO: doc
				filenameIsUtf8 = true
			}
		}
	}

	// TODO: destFsType is empty which makes makeFileNameValidForDestFs usefuless,
	// why call this function on nautilus???
	filename, _ = makeFileNameValidForDestFs(filename, destFsType)

	return filename, filenameIsUtf8
}

func (job *CreateJob) fillInitContent(out *gio.FileOutputStream) (bool, error) {
	// cannot use cast to convert []byte to []int directly.
	data := []uint8{}
	for _, i := range job.srcData {
		data = append(data, uint8(i))
	}

	_, res, err := out.WriteAll(data, job.cancellable)
	if res {
		res, err = out.Close(job.cancellable)
		// TODO:
		// if res && job.undoInfo != nil {
		// 	// nautilus_file_undo_info_create_set_data(NAUTILUS_FILE_UNDO_INFO_CREATE(common.undoInfo), dest, job.data, job.length)
		// }
	}

	return res, err
}

func (job *CreateJob) createTarget(dest *gio.File) (bool, error) {
	var err error
	res := false

	if job.makeDir {
		res, err = dest.MakeDirectory(job.cancellable)

		// TODO:
		// if res && job.CommonJob.undoInfo != nil {
		// nautilus_file_undo_info_create_set_data(NAUTILUS_FILE_UNDO_INFO_CREATE(job.CommonJob.undoInfo), dest, NULL, 0)
		// }
		return res, err
	}

	// create from template
	if job.src != nil {
		// TODO: use own copy job to make sure the name get the same format.
		res, err = job.src.Copy(dest, gio.FileCopyFlagsNone, job.cancellable, nil)
		// TODO:
		// if res && job.undoInfo != nil {
		// uri := job.src.GetUri()
		// nautilus_file_undo_info_create_set_data(NAUTILUS_FILE_UNDO_INFO_CREATE(job.CommonJob.undoInfo), dest, uri, 0)
		// }
		return res, err
	}

	out, err := dest.Create(gio.FileCreateFlagsNone, job.cancellable)
	if err != nil {
		return false, err
	}
	defer out.Unref()

	return job.fillInitContent(out)
}

func (job *CreateJob) getNewName(targetInfo *_TargetInfo) string {
	filenameBase := FilenameStripExtension(targetInfo.filename)
	offset := len(filenameBase)
	suffix := targetInfo.filename[offset:]
	return fmt.Sprintf("%s %d%s", filenameBase, targetInfo.count, suffix)
}

func (job *CreateJob) getDest(targetInfo *_TargetInfo) *gio.File {
	filename := targetInfo.newFilename
	if filename == "" {
		filename = targetInfo.filename
	}

	var dest *gio.File
	if targetInfo.filenameIsUtf8 {
		dest, _ = job.destDir.GetChildForDisplayName(filename)
	}
	if dest == nil {
		dest = job.destDir.GetChild(filename)
	}
	return dest
}

func (job *CreateJob) needRetry(targetInfo *_TargetInfo, err gio.GError) bool {
	retry := false

	filename := targetInfo.filename
	maxLength := targetInfo.maxLength
	destFsType := targetInfo.destFsType

	// TODO: compact code
	errCode := gio.IOErrorEnum(err.Code)
	if errCode == gio.IOErrorEnumInvalidFilename &&
		!targetInfo.handledInvalidFilename {
		targetInfo.handledInvalidFilename = true

		newFilename := ""
		if targetInfo.count == 1 {
			newFilename = filename
		} else {
			filename2 := job.getNewName(targetInfo)
			if maxLength > 0 && len(filename2) > maxLength {
				newFilename = ShortenUtf8String(filename2, len(filename2)-maxLength)
			}
			if newFilename == "" {
				newFilename = filename2
			}
		}

		newFilename, ok := makeFileNameValidForDestFs(newFilename, destFsType)
		if ok {
			targetInfo.newFilename = newFilename
			retry = true
		}
	} else if errCode == gio.IOErrorEnumExists {
		targetInfo.count++
		filename2 := job.getNewName(targetInfo)

		if maxLength > 0 && len(filename2) > maxLength {
			newFilename := ShortenUtf8String(filename2, len(filename2)-maxLength)
			if newFilename != "" {
				filename2 = newFilename
			}
		}

		filename2, _ = makeFileNameValidForDestFs(filename2, destFsType)
		targetInfo.newFilename = filename2
		retry = true
	} else if errCode != gio.IOErrorEnumCancelled {
		primaryText := ""
		// TODO: doc
		if job.makeDir {
			primaryText = Tr("Error while creating directory %B") //, dest)
		} else {
			primaryText = Tr("Error while creating file %B") //, dest)
		}
		secondaryText := Tr("There was error creating the directory in %F.") //, job.destDir)
		detailText := err.Error()

		response := job.uiDelegate.AskSkip(primaryText, secondaryText, detailText, UIFlagsNone)

		switch response.Code() {
		case ResponseCancel:
			job.Abort()
		case ResponseSkip: // skip, do nothing
		}
	}

	return retry
}

// Execute create job
func (job *CreateJob) Execute() error {
	defer job.finalize()
	if job.makeLink {
		job.linkJob()
		return nil
	}

	// the minimum size of a file or directory is 4k(one block) usually.
	requiredSize := uint64(4096)
	if job.src != nil {
		// this is a little trick for making defer work.
		err := func() error {
			info, err := job.src.QueryInfo(strings.Join(
				[]string{
					gio.FileAttributeStandardType,
					gio.FileAttributeStandardSize,
				}, ","),
				gio.FileQueryInfoFlagsNone, job.cancellable)
			if err != nil {
				return err
			}
			defer info.Unref()

			if info.GetFileType() != gio.FileTypeRegular {
				info.Unref()
				return fmt.Errorf("%v is not a file", job.src.GetUri())
			}

			size := info.GetSize()
			if size > 0 {
				requiredSize = uint64(size)
			}
			return nil
		}()
		if err != nil {
			return err
		}
	}

	// TODO:
	// nautilus_progress_info_start(job.common.progress)
	// undo info??

	targetInfo := &_TargetInfo{
		handledInvalidFilename: false,
		destFsType:             queryFsType(job.destDir, job.cancellable),
		maxLength:              getMaxNameLength(job.destDir),
		count:                  1,
	}

	job.verifyDestination(job.destDir, requiredSize) //math.MaxUint64)

	if job.isAborted() {
		return fmt.Errorf("job is aborted")
	}

	targetInfo.filename, targetInfo.filenameIsUtf8 = job.getFilename(targetInfo.destFsType)

	dest := job.getDest(targetInfo)

retry:
	ok, err := job.createTarget(dest)

	if ok {
		job.createdFile = dest // Ref??
		// TODO:
		// nautilus_file_changes_queue_file_added(dest)
		// if job.hasPosition {
		// 	nautilus_file_changes_queue_file_position_set(dest, job.position, job.CommonJob.screenNum)
		// } else {
		// 	nautilus_file_changes_queue_schedule_position_remove(dest)
		// }
		job.setResult(dest.GetUri())
		return nil
	}

	if job.needRetry(targetInfo, err.(gio.GError)) {
		dest.Unref()
		dest = job.getDest(targetInfo)
		goto retry
	}

	job.setResult(dest.GetUri())
	dest.Unref()
	return nil
}

func newCreateJob(destDir *gio.File, src *gio.File, makeDir bool, filename string, initContent []byte, uiDelegate IUIDelegate) *CreateJob {
	job := &CreateJob{
		CommonJob: newCommon(uiDelegate),
		src:       src,
		destDir:   destDir,
		makeDir:   makeDir,
		filename:  filename,
		srcData:   initContent,
	}
	return job
}

// NewCreateFileJob is used to create a new file.
// @param destURL: parent dir which contains the new directory.
func NewCreateFileJob(destDirURL string, filename string, initContent []byte, uiDelegate IUIDelegate) *CreateJob {
	destDir := gio.FileNewForCommandlineArg(destDirURL)
	return newCreateJob(destDir, nil, false, filename, initContent, uiDelegate)
}

// NewCreateFileFromTemplateJob creates new file from template.
// @param destURL: parent dir which contains the new directory.
func NewCreateFileFromTemplateJob(destDirURL string, templateURL string, uiDelegate IUIDelegate) *CreateJob {
	src := gio.FileNewForCommandlineArg(templateURL)
	destDir := gio.FileNewForCommandlineArg(destDirURL)
	job := newCreateJob(destDir, src, false, "", []byte{}, uiDelegate)
	return job
}

// NewCreateDirectoryJob creates directory
// @param destURL: parent dir which contains the new directory.
func NewCreateDirectoryJob(destDirURL string, dirname string, uiDelegate IUIDelegate) *CreateJob {
	destDir := gio.FileNewForCommandlineArg(destDirURL)
	return newCreateJob(destDir, nil, true, dirname, []byte{}, uiDelegate)
}

func getLinkName(name string, count int, maxLength int) string {
	if count < 0 {
		count = 0
	}

	format := ""
	useCount := false
	if count <= 2 {
		switch count {
		default:
			fallthrough
		case 0:
			format = "%s"
		case 1:
			format = Tr("Link to %s") // TODO: doc
		case 2:
			format = Tr("Another link to %s") // TODO: doc
		}
	} else {
		switch count % 10 {
		case 1:
			format = Tr("%'dst link to %s") // TODO: doc
		case 2:
			format = Tr("%'dnd link to %s") // TODO: doc
		case 3:
			format = Tr("%'drd link to %s") // TODO: doc
		default:
			format = Tr("%'dth link to %s") // TODO: doc
		}
		useCount = true
	}

	linkName := ""
	if useCount {
		linkName = fmt.Sprintf(format, count, name)
	} else {
		linkName = fmt.Sprintf(format, name)
	}

	unshortenedLength := len(linkName)
	if maxLength > 0 && unshortenedLength > maxLength {
		newName := ShortenUtf8String(name, unshortenedLength-maxLength)
		if newName != "" {
			if useCount {
				linkName = fmt.Sprintf(format, count, newName)
			} else {
				linkName = fmt.Sprintf(format, newName)
			}
		}
	}

	return linkName
}

func getTargetFileForLink(src *gio.File, destDir *gio.File, destFsType string, count int) *gio.File {
	maxLength := getMaxNameLength(destDir)
	var dest *gio.File
	srcInfo, _ := src.QueryInfo(gio.FileAttributeStandardEditName, gio.FileQueryInfoFlagsNone, nil)
	if srcInfo != nil {
		editName := srcInfo.GetAttributeString(gio.FileAttributeStandardEditName)
		if editName != "" {
			newName := getLinkName(editName, count, maxLength)
			makeFileNameValidForDestFs(newName, destFsType)
			dest, _ = destDir.GetChildForDisplayName(newName)
		}
		srcInfo.Unref()
	}

	if dest == nil {
		basename := src.GetBasename()
		makeFileNameValidForDestFs(basename, destFsType)

		if utf8.ValidString(basename) {
			newName := getLinkName(basename, count, maxLength)
			makeFileNameValidForDestFs(newName, destFsType)
			dest, _ = destDir.GetChildForDisplayName(newName)
		}

		if dest == nil {
			newName := ""
			if count == 1 {
				newName = fmt.Sprintf("%s.link", basename) // TODO: doc
			} else {
				newName = fmt.Sprintf("%s.link%d", basename, count) // TODO: doc
			}
			makeFileNameValidForDestFs(newName, destFsType)
			dest = destDir.GetChild(newName)
		}
	}

	return dest
}

func getAbsPathForSymlink(file *gio.File, destination *gio.File) string {
	if file.IsNative() || destination.IsNative() {
		return file.GetPath()
	}

	root := file.Dup()
	for {
		parent := root.GetParent()
		if parent == nil {
			break
		}

		root.Unref()
		root = parent
	}

	relative := root.GetRelativePath(file)
	root.Unref()
	return filepath.Join("/", relative)
}

func (job *CreateJob) linkFile(src *gio.File, destFsType *string) {
	destDir := job.destDir
	srcDir := src.GetParent()
	count := 0
	handledInvalidName := false
	if srcDir.Equal(destDir) {
		count = 1
	}
	srcDir.Unref()
	dest := getTargetFileForLink(src, destDir, *destFsType, count)

retry:
	var err error
	var ok bool
	// notLocal := false // TODO: used for error message.
	path := getAbsPathForSymlink(src, dest)
	if path == "" {
		// notLocal = true
	} else if ok, err = dest.MakeSymbolicLink(path, job.cancellable); ok {
		// TODO: undo
		// nautilus_file_changes_queue_file_added (dest);
		dest.Unref()
		return
	}

	if err == nil {
		dest.Unref()
		return
	}

	gerr := err.(gio.GError)
	errCode := gio.IOErrorEnum(gerr.Code)
	if errCode == gio.IOErrorEnumInvalidFilename && handledInvalidName {
		handledInvalidName = true
		*destFsType = queryFsType(destDir, job.cancellable)
		newDest := getTargetFileForLink(src, destDir, *destFsType, count)
		if newDest.Equal(dest) {
			newDest.Unref()
		} else {
			dest.Unref()
			dest = newDest
			goto retry
		}
	}

	switch errCode {
	case gio.IOErrorEnumExists:
		dest.Unref()
		dest = getTargetFileForLink(src, destDir, *destFsType, count)
		count++
		goto retry
	case gio.IOErrorEnumCancelled: // do nothing
	default:
		// TODO: send error message.
		// AskSkip is for multi files.
		// primaryText := Tr("Error while creating link to %B.") //, src);
		// secondaryText := ""
		// detailText := ""
		//
		// if notLocal {
		// 	secondaryText = Tr("Symbolic links only supported for local files")
		// } else if errCode == gio.IOErrorEnumNotSupported {
		// 	secondaryText = Tr("The target doesn't support symbolic links.")
		// } else {
		// 	secondaryText = Tr("There was an error creating the symlink in %F.") //, dest_dir;
		// 	detailText = err.Error()
		// }
		//
		// uiFlags := UIFlagsNone
		// response := job.UI().AskSkip(primaryText, secondaryText, detailText, uiFlags)
		// switch response.Code() {
		// case ResponseCancel:
		// 	job.Abort()
		// case ResponseSkip:
		// }
	}

	dest.Unref()
}

func (job *CreateJob) linkJob() {
	destFsType := ""
	job.verifyDestination(job.destDir, 0) // link use 0 storage.

	if job.isAborted() {
		return
	}

	if job.isAborted() {
		return
	}

	job.linkFile(job.src, &destFsType)
}

func newCreateLinkJob(src *gio.File, destDir *gio.File, uiDelegate IUIDelegate) *CreateJob {
	job := newCreateJob(destDir, src, false, "", []byte{}, uiDelegate)
	job.makeLink = true
	return job
}

// NewCreateLinkJob creates symbol link.
func NewCreateLinkJob(srcURL string, destDirURL string, uiDelegate IUIDelegate) *CreateJob {
	src := gio.FileNewForCommandlineArg(srcURL)
	destDir := gio.FileNewForCommandlineArg(destDirURL)
	return newCreateLinkJob(src, destDir, uiDelegate)
}
