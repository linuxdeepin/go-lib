package appinfo

type AppInfo interface {
	GetId() string
	GetName() string
	GetIcon() string
	GetExecutable() string
	GetFileName() string
	GetCommandline() string
	Launch(files []string, launchContext *AppLaunchContext) error
	GetStartupWMClass() string
}
