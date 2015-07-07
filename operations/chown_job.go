package operations

// #include <stdlib.h>
// #include <sys/types.h>
// #include <grp.h>
import "C"
import "unsafe"
import (
	"errors"
	"os"
	"os/user"
	"pkg.deepin.io/lib/gio-2.0"
	"strconv"
)

// ChownJob changes the owner or group of a file or directory.
type ChownJob struct {
	*CommonJob
	file  *gio.File
	owner string
	group string
}

func (job *ChownJob) finalize() {
	defer job.CommonJob.finalize()
	if job.file != nil {
		job.file.Unref()
	}
}

// Execute ChownJob.
func (job *ChownJob) Execute() {
	defer job.finalize()
	// -1 means not change.
	uid := -1
	gid := -1
	if job.owner != "" {
		newUser, err := user.Lookup(job.owner)
		if err != nil {
			job.setError(err)
			return
		}
		uid, _ = strconv.Atoi(newUser.Uid)
	}

	if job.group != "" {
		cGroupName := C.CString(job.group)
		defer C.free(unsafe.Pointer(cGroupName))
		group := C.getgrnam(cGroupName)
		if group == nil {
			job.setError(errors.New("no such a group"))
			return
		}
		gid = int(group.gr_gid)
	}

	job.setError(os.Chown(job.file.GetPath(), uid, gid))
}

func newChownJob(file *gio.File, newOwner string, newGroup string) *ChownJob {
	return &ChownJob{
		CommonJob: newCommon(nil),
		file:      file,
		owner:     newOwner,
		group:     newGroup,
	}
}

// NewChownJob creates a new ChownJob.
func NewChownJob(uri string, newOwner string, newGroup string) *ChownJob {
	file := gio.FileNewForCommandlineArg(uri)
	return newChownJob(file, newOwner, newGroup)
}
