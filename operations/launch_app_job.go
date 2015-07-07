package operations

import (
	"fmt"
	"pkg.deepin.io/lib/gio-2.0"
	"sort"
)

type byName []*gio.AppInfo

func (self byName) Less(i, j int) bool {
	return self[i].GetName() < self[j].GetName()
}

func (self byName) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self byName) Len() int {
	return len(self)
}

// GetLaunchAppJob is the base struct for get launch relatived job.
type GetLaunchAppJob struct {
	*CommonJob
	uri  string
	file *gio.File
}

func NewGetLaunchAppJob(uri string) *GetLaunchAppJob {
	return &GetLaunchAppJob{
		CommonJob: newCommon(nil),
		uri:       uri,
		file:      gio.FileNewForCommandlineArg(uri),
	}
}

func (job *GetLaunchAppJob) getContentType() string {
	info, err := job.file.QueryInfo(gio.FileAttributeStandardContentType, gio.FileQueryInfoFlagsNone, nil)
	if err != nil {
		job.setError(err)
		return ""
	}
	defer info.Unref()

	return info.GetContentType()
}

func (job *GetLaunchAppJob) finalize() {
	if job.file != nil {
		job.file.Unref()
	}
}

// GetDefaultLaunchAppJob will get the default launch app.
type GetDefaultLaunchAppJob struct {
	*GetLaunchAppJob
	mustSupportURI bool
}

func NewGetDefaultLaunchAppJob(uri string, mustSupportURI bool) *GetDefaultLaunchAppJob {
	return &GetDefaultLaunchAppJob{
		GetLaunchAppJob: NewGetLaunchAppJob(uri),
		mustSupportURI:  mustSupportURI,
	}
}

func (job *GetDefaultLaunchAppJob) Execute() {
	defer job.finalize()
	defer job.emitDone()

	if job.file == nil {
		job.setError(fmt.Errorf("No such a file: %q", job.uri))
		return
	}

	mimeType := job.getContentType()
	defaultApp := gio.AppInfoGetDefaultForType(mimeType, job.mustSupportURI)
	if defaultApp == nil {
		job.setError(fmt.Errorf("the default app of %q is not existed", job.uri))
		return
	}

	job.setResult(defaultApp)
}

// GetRecommendedLaunchAppsJob will get all recommended launch apps.
type GetRecommendedLaunchAppsJob struct {
	*GetLaunchAppJob
}

func NewGetRecommendedLaunchAppsJob(uri string) *GetRecommendedLaunchAppsJob {
	return &GetRecommendedLaunchAppsJob{
		GetLaunchAppJob: NewGetLaunchAppJob(uri),
	}
}

func (job *GetRecommendedLaunchAppsJob) Execute() {
	defer job.finalize()
	defer job.emitDone()

	if job.file == nil {
		job.setError(fmt.Errorf("No such a file: %q", job.uri))
		return
	}

	mimeType := job.getContentType()
	apps := gio.AppInfoGetRecommendedForType(mimeType)
	apps = append(apps, gio.AppInfoGetFallbackForType(mimeType)...)

	sort.Sort(byName(apps))
	job.setResult(apps)
}

// GetAllLaunchAppsJob will get all apps.
type GetAllLaunchAppsJob struct {
	*CommonJob
}

func NewGetAllLaunchAppsJob() *GetAllLaunchAppsJob {
	return &GetAllLaunchAppsJob{
		CommonJob: newCommon(nil),
	}
}

func (job *GetAllLaunchAppsJob) Execute() {
	defer job.finalize()
	defer job.emitDone()

	apps := gio.AppInfoGetAll()
	sort.Sort(byName(apps))

	job.setResult(apps)
}

// SetDefaultLaunchAppJob sets the default launch applications for a MIME type.
type SetDefaultLaunchAppJob struct {
	*CommonJob
	app      *gio.AppInfo
	mimeType string
}

func (job *SetDefaultLaunchAppJob) finalize() {
	if job.app != nil {
		job.app.Unref()
	}
}

// Execute the SetLaunchAppJob.
func (job *SetDefaultLaunchAppJob) Execute() {
	defer job.finalize()
	defer job.emitDone()

	_, err := job.app.SetAsDefaultForType(job.mimeType)
	job.setError(err)
}

// NewSetLaunchAppJob creates a new SetLaunchAppJob.
func NewSetDefaultLaunchAppJob(id string, mimeType string) *SetDefaultLaunchAppJob {
	desktopApp := gio.NewDesktopAppInfo(id)
	app := gio.ToAppInfo(desktopApp)

	return &SetDefaultLaunchAppJob{
		CommonJob: newCommon(nil),
		app:       app,
		mimeType:  mimeType,
	}
}
