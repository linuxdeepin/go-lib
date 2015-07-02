package operations

import (
	"net/url"
	"pkg.linuxdeepin.com/lib/gio-2.0"
	"sort"
)

// AppInfo holds the Name and Id of a application.
type AppInfo struct {
	Name string
	ID   string
}

type byName []AppInfo

func (self byName) Less(i, j int) bool {
	return self[i].Name < self[j].Name
}

func (self byName) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self byName) Len() int {
	return len(self)
}

func newAppInfo(name string, id string) AppInfo {
	return AppInfo{
		Name: name,
		ID:   id,
	}
}

// LaunchAppInfo holds default launch application, recommended applications and all the applications.
type LaunchAppInfo struct {
	DefaultApp      AppInfo
	RecommendedApps []AppInfo
	OtherApps       []AppInfo
}

// LaunchAppJob is used to list all the possible apps for launch a MIME type.
type LaunchAppJob struct {
	*CommonJob
	file *gio.File
}

func (job *LaunchAppJob) finalize() {
	if job.file != nil {
		job.file.Unref()
	}
}

func appendApp(apps []*gio.AppInfo, appInfos *[]AppInfo, defaultAppID string) {
	for _, app := range apps {
		id := app.GetId()
		name := app.GetName()
		app.Unref()
		if id == defaultAppID {
			continue
		}
		*appInfos = append(*appInfos, newAppInfo(name, id))
	}
}

// Execute the LaunchAppJob.
func (job *LaunchAppJob) Execute() {
	defer job.finalize()
	defer job.emitDone()

	launchAppInfo := &LaunchAppInfo{
		RecommendedApps: []AppInfo{},
		OtherApps:       []AppInfo{},
	}
	info, err := job.file.QueryInfo(gio.FileAttributeStandardContentType, gio.FileQueryInfoFlagsNone, nil)
	if err != nil {
		job.setError(err)
		return
	}

	mimeType := info.GetContentType()
	defaultApp := gio.AppInfoGetDefaultForType(mimeType, false)
	defaultAppName := defaultApp.GetName()
	defaultAppID := defaultApp.GetId()
	launchAppInfo.DefaultApp = newAppInfo(defaultAppName, defaultAppID)

	apps := gio.AppInfoGetRecommendedForType(mimeType)
	appendApp(apps, &(launchAppInfo.RecommendedApps), defaultAppID)

	apps = gio.AppInfoGetFallbackForType(mimeType)
	appendApp(apps, &launchAppInfo.RecommendedApps, defaultAppID)

	apps = gio.AppInfoGetAll()
	appendApp(apps, &launchAppInfo.OtherApps, defaultAppID)

	sort.Sort(byName(launchAppInfo.RecommendedApps))
	sort.Sort(byName(launchAppInfo.OtherApps))

	job.setResult(launchAppInfo)
}

func newLaunchAppJob(file *gio.File) *LaunchAppJob {
	return &LaunchAppJob{
		CommonJob: newCommon(nil),
		file:      file,
	}
}

// NewLaunchAppJob creates a new LaunchAppJob.
func NewLaunchAppJob(uri *url.URL) *LaunchAppJob {
	file := uriToGFile(uri)
	return newLaunchAppJob(file)
}

// SetLaunchAppJob sets the default launch applications for a MIME type.
type SetLaunchAppJob struct {
	*CommonJob
	app      *gio.AppInfo
	mimeType string
}

func (job *SetLaunchAppJob) finalize() {
	if job.app != nil {
		job.app.Unref()
	}
}

// Execute the SetLaunchAppJob.
func (job *SetLaunchAppJob) Execute() {
	defer job.finalize()
	defer job.emitDone()

	_, err := job.app.SetAsDefaultForType(job.mimeType)
	job.setError(err)
	return
}

// NewSetLaunchAppJob creates a new SetLaunchAppJob.
func NewSetLaunchAppJob(id string, mimeType string) *SetLaunchAppJob {
	desktopApp := gio.NewDesktopAppInfo(id)
	app := gio.ToAppInfo(desktopApp)

	return &SetLaunchAppJob{
		CommonJob: newCommon(nil),
		app:       app,
		mimeType:  mimeType,
	}
}
