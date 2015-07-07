package operations

import (
	"pkg.deepin.io/lib/gio-2.0"
	"strings"
)

// StatJob list some useful properties for properties window.
type StatJob struct {
	*CommonJob
	file *gio.File
}

// StatProperty is the useful properties for properties window.
// the size field of directory is not the size of the whole directory,
// using list job with includeHidden and recursive flags to get the real
// size of a directory.
type StatProperty struct {
	URI        string
	Name       string
	MIME       string
	LinkTarget string
	Mode       uint32
	Size       int64
	MTime      uint64
	ATime      uint64
	Owner      string
	OwnerReal  string
	Group      string
}

func newStatJob(file *gio.File) *StatJob {
	job := &StatJob{
		CommonJob: newCommon(nil),
		file:      file,
	}
	return job
}

func (job *StatJob) finalize() {
	if job.file != nil {
		job.file.Unref()
	}
}

// Execute StatJob.
func (job *StatJob) Execute() {
	defer job.finalize()
	defer job.emitDone()
	info, err := job.file.QueryInfo(strings.Join(
		[]string{
			gio.FileAttributeStandardDisplayName,
			gio.FileAttributeUnixMode,
			gio.FileAttributeStandardSize,
			gio.FileAttributeTimeModified,
			gio.FileAttributeTimeAccess,
			gio.FileAttributeStandardContentType,
			gio.FileAttributeOwnerUser,
			gio.FileAttributeOwnerUserReal,
			gio.FileAttributeOwnerGroup,
		}, ","),
		gio.FileQueryInfoFlagsNofollowSymlinks,
		nil)
	if err != nil {
		job.setError(err)
		return
	}
	defer info.Unref()

	contentType := info.GetContentType()
	linkTarget := info.GetSymlinkTarget()
	size := info.GetSize()
	mTime := info.GetAttributeUint64(gio.FileAttributeTimeModified)
	aTime := info.GetAttributeUint64(gio.FileAttributeTimeAccess)
	displayName := info.GetDisplayName()
	mode := uint32(info.GetAttributeUint32(gio.FileAttributeUnixMode))

	stat := StatProperty{
		Name:       displayName,
		URI:        job.file.GetUri(),
		MIME:       contentType,
		LinkTarget: linkTarget,
		Mode:       mode,
		Size:       size,
		MTime:      mTime,
		ATime:      aTime,
		Group:      info.GetAttributeString(gio.FileAttributeOwnerGroup),
		Owner:      info.GetAttributeString(gio.FileAttributeOwnerUser),
		OwnerReal:  info.GetAttributeString(gio.FileAttributeOwnerUserReal),
	}
	job.setResult(stat)
}

// NewStatJob creates a new StatJob.
func NewStatJob(uri string) *StatJob {
	dest := gio.FileNewForCommandlineArg(uri)
	return newStatJob(dest)
}
