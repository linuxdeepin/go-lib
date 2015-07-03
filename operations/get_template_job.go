package operations

type GetTemplateJob struct {
	templateDirURI string
}

func NewGetTemplateJob(templateDirURI string) *GetTemplateJob {
	return &GetTemplateJob{
		templateDirURI: templateDirURI,
	}
}

func shouldShow(property ListProperty) bool {
	return property.CanRead && !property.IsHidden && !property.IsBackup && property.MIME != "inode/directory"
}

func (job *GetTemplateJob) Execute() []string {
	files := []string{}
	listJob := NewListDirJob(job.templateDirURI, ListJobFlagNone)
	listJob.ListenProperty(func(property ListProperty) {
		if shouldShow(property) {
			files = append(files, property.URI)
		}
	})
	listJob.Execute()

	return files
}
