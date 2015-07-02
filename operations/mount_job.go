package operations

type IMountUI interface {
	Username() string
	Domain() string
	Password() string
	IsAnonymous() bool
	RememberPasswordFlags() int

	// TODO: ???
	AskPassword()
	AskQuestion()
	ShowProcess()
	ShowUnmountProcess()
}

type MountJob struct {
	*CommonJob
}

func newMountJob() *MountJob {
	job := &MountJob{
		CommonJob: newCommon(nil), // TODO
	}
	return job
}

func NewMountJob() *MountJob {
	return newMountJob()
}
