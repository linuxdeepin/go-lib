package desktopappinfo

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"pkg.deepin.io/lib/keyfile"
	xdgdir "pkg.deepin.io/lib/xdg/basedir"
	"strings"
	"sync"
)

const (
	MainSection            = "Desktop Entry"
	KeyType                = "Type"
	KeyVersion             = "Version"
	KeyName                = "Name"
	KeyGenericName         = "GenericName"
	KeyNoDisplay           = "NoDisplay"
	KeyComment             = "Comment"
	KeyIcon                = "Icon"
	KeyHidden              = "Hidden"
	KeyOnlyShowIn          = "OnlyShowIn"
	KeyNotShowIn           = "NotShowIn"
	KeyTryExec             = "TryExec"
	KeyExec                = "Exec"
	KeyPath                = "Path"
	KeyTerminal            = "Terminal"
	KeyMimeType            = "MimeType"
	KeyCategories          = "Categories"
	KeyKeywords            = "Keywords"
	KeyStartupNotify       = "StartupNotify"
	KeyStartupWMClass      = "StartupWMClass"
	KeyURL                 = "URL"
	KeyActions             = "Actions"
	KeyDBusDBusActivatable = "DBusActivatable"

	TypeApplication = "Application"
	TypeLink        = "Link"
	TypeDirectory   = "Directory"

	envDesktopEnv = "XDG_CURRENT_DESKTOP"
	desktopExt    = ".desktop"
)

var xdgDataDirs []string
var xdgAppDirs []string

func init() {
	xdgDataDirs = xdgdir.GetSystemDataDirs()
	xdgDataDirs = append(xdgDataDirs, xdgdir.GetUserDataDir())

	xdgAppDirs = make([]string, len(xdgDataDirs))
	for i, dir := range xdgDataDirs {
		xdgAppDirs[i] = filepath.Join(dir, "applications")
	}
}

type DesktopAppInfo struct {
	*keyfile.KeyFile
	filename string
	id       string

	name string
	icon string
}

func NewDesktopAppInfo(id string) *DesktopAppInfo {
	if !strings.HasSuffix(id, desktopExt) {
		id += desktopExt
	}
	if filepath.IsAbs(id) {
		ai, _ := NewDesktopAppInfoFromFile(id)
		return ai
	}

	for _, appDir := range xdgAppDirs {
		file := filepath.Join(appDir, id)
		ai, err := newDesktopAppInfoFromFile(file)
		if err == nil {
			// length of desktopExt is 8
			ai.id = id[:len(id)-8]
			return ai
		}
	}
	return nil
}

var ErrInvalidFileExt = errors.New("the file extension is not " + desktopExt)
var ErrInvalidFirstSection = errors.New("first section is not " + MainSection)
var ErrInvalidType = errors.New("type is not " + TypeApplication)

func newDesktopAppInfoFromFile(filename string) (*DesktopAppInfo, error) {
	kfile := keyfile.NewKeyFile()
	err := kfile.LoadFromFile(filename)
	if err != nil {
		return nil, err
	}
	f, err := NewDesktopAppInfoFromKeyFile(kfile)
	if err != nil {
		return nil, err
	}
	f.filename = filename
	return f, nil
}

func NewDesktopAppInfoFromKeyFile(kfile *keyfile.KeyFile) (*DesktopAppInfo, error) {
	f := &DesktopAppInfo{
		KeyFile: kfile,
	}

	sections := f.GetSections()
	if !(len(sections) > 0 && sections[0] == MainSection) {
		return nil, ErrInvalidFirstSection
	}

	type0, _ := f.GetValue(MainSection, KeyType)
	if type0 != TypeApplication {
		return nil, ErrInvalidType
	}

	f.name, _ = f.GetLocaleString(MainSection, KeyName, "")
	icon, _ := f.GetString(MainSection, KeyIcon)
	/* Work around a common mistake in desktop files */
	if !filepath.IsAbs(icon) {
		ext := filepath.Ext(icon)
		switch ext {
		case ".png", ".xpm", ".svg":
			icon = strings.TrimSuffix(icon, ext)
		}
	}
	f.icon = icon
	return f, nil
}

func NewDesktopAppInfoFromFile(filename string) (*DesktopAppInfo, error) {
	if !strings.HasSuffix(filename, desktopExt) {
		return nil, ErrInvalidFileExt
	}
	filename, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	ai, err := newDesktopAppInfoFromFile(filename)
	if err != nil {
		return nil, err
	}
	ai.id = getId(filename)
	return ai, err
}

// filename must has suffix desktopExt
// example:
// /usr/share/applications/a.desktop -> a
// /usr/share/applications/kde4/a.desktop -> kde4/a
// /xxxx/dir/a.desktop -> /xxxx/dir/a
func getId(filename string) string {
	filename = strings.TrimSuffix(filename, desktopExt)
	i := strings.Index(filename, "/applications/")
	if i == -1 {
		return filename
	}

	dir := filename[:i]

	var installed bool
	for _, d := range xdgDataDirs {
		if d == dir {
			installed = true
			break
		}
	}

	if installed {
		// length of "/applications/" is 14
		return filename[i+14:]
	}
	return filename
}

func (ai *DesktopAppInfo) GetId() string {
	return ai.id
}

func (ai *DesktopAppInfo) IsInstalled() bool {
	i := strings.Index(ai.filename, "/applications/")
	if i == -1 {
		return false
	}

	dir := ai.filename[:i]
	for _, d := range xdgDataDirs {
		if d == dir {
			return true
		}
	}
	return false
}

func (ai *DesktopAppInfo) GetFileName() string {
	return ai.filename
}

func (ai *DesktopAppInfo) GetIsHiden() bool {
	hidden, _ := ai.GetBool(MainSection, KeyHidden)
	return hidden
}

func (ai *DesktopAppInfo) GetNoDisplay() bool {
	noDisplay, _ := ai.GetBool(MainSection, KeyNoDisplay)
	return noDisplay
}

var currentDesktops []string
var currentDesktopsOnce sync.Once

func getCurrentDesktop() []string {
	currentDesktopsOnce.Do(func() {
		desktopEnv := os.Getenv(envDesktopEnv)
		currentDesktops = strings.Split(desktopEnv, ":")
	})
	return currentDesktops
}

func (ai *DesktopAppInfo) GetShowIn(desktopEnvs []string) bool {
	if len(desktopEnvs) == 0 {
		desktopEnvs = getCurrentDesktop()
	}

	onlyShowIn, _ := ai.GetStringList(MainSection, KeyOnlyShowIn)
	notShowIn, _ := ai.GetStringList(MainSection, KeyNotShowIn)

	for _, env := range desktopEnvs {
		for _, de := range onlyShowIn {
			if env == de {
				return true
			}
		}

		for _, de := range notShowIn {
			if env == de {
				return false
			}
		}
	}

	return len(onlyShowIn) == 0
}

func (ai *DesktopAppInfo) ShouldShow() bool {
	if ai.GetNoDisplay() {
		return false
	}
	return ai.GetShowIn(nil)
}

func (ai *DesktopAppInfo) GetName() string {
	return ai.name
}

func (ai *DesktopAppInfo) GetGenericName() string {
	gname, _ := ai.GetLocaleString(MainSection, KeyGenericName, "")
	return gname
}

func (ai *DesktopAppInfo) GetDisplayName() string {
	return ai.GetName()
}

func (ai *DesktopAppInfo) GetCategories() []string {
	categories, _ := ai.GetStringList(MainSection, KeyCategories)
	return categories
}

func (ai *DesktopAppInfo) GetKeywords() []string {
	keywords, _ := ai.GetLocaleStringList(MainSection, KeyKeywords, "")
	return keywords
}

func (ai *DesktopAppInfo) GetStartupWMClass() string {
	wmclass, _ := ai.GetString(MainSection, KeyStartupWMClass)
	return wmclass
}

func (ai *DesktopAppInfo) GetStartupNotify() bool {
	sn, _ := ai.GetBool(MainSection, KeyStartupNotify)
	return sn
}

func (ai *DesktopAppInfo) GetIcon() string {
	return ai.icon
}

func (ai *DesktopAppInfo) GetCommandline() string {
	exec, _ := ai.GetString(MainSection, KeyExec)
	return exec
}

func (ai *DesktopAppInfo) GetTryExec() string {
	tryExec, _ := ai.GetString(MainSection, KeyTryExec)
	return tryExec
}

func (ai *DesktopAppInfo) GetTerminal() bool {
	useTerminal, _ := ai.GetBool(MainSection, KeyTerminal)
	return useTerminal
}

func (ai *DesktopAppInfo) Launch(timestamp uint32, files []string) error {
	return _launch(ai, ai.GetCommandline(), timestamp, files)
}

const launchScript = `export GIO_LAUNCHED_DESKTOP_FILE_PID=$$;exec $@`

func _launch(ai *DesktopAppInfo, cmdline string, timestamp uint32, files []string) error {
	if cmdline == "" {
		return errors.New("command line is empty")
	}

	exeargs, err := splitExec(cmdline)
	if err != nil {
		return err
	}

	exeargs, err = ai.expandFieldCode(exeargs, files)
	if err != nil {
		return err
	}

	useTerminal := ai.GetTerminal()
	if useTerminal {
		termExec, termExecArg := getDefaultTerminal()
		exeargs = append([]string{termExec, termExecArg}, exeargs...)
	}

	shellArgs := append([]string{"/dev/stdin"}, exeargs...)
	cmd := exec.Command("/bin/sh", shellArgs...)
	cmd.Stdin = strings.NewReader(launchScript)
	cmd.Env = append(os.Environ(), "GIO_LAUNCHED_DESKTOP_FILE="+ai.GetFileName())

	startupNotify := ai.GetStartupNotify()
	if startupNotify && len(exeargs) > 0 {
		snId := getStartupNotifyId(exeargs[0], timestamp)
		cmd.Env = append(cmd.Env, "DESKTOP_STARTUP_ID="+snId)
	}

	err = cmd.Start()
	go cmd.Wait()
	return err
}

// [Desktop Action new-window]
// or [Full_Screenshot Shortcut Group]
func isDesktopAction(name string) bool {
	return strings.HasPrefix(name, "Desktop Action") ||
		strings.HasSuffix(name, "Shortcut Group")
}

func (ai *DesktopAppInfo) GetActions() []DesktopAction {
	var actions []DesktopAction
	for _, section := range ai.GetSections() {
		if isDesktopAction(section) {
			name, _ := ai.GetLocaleString(section, KeyName, "")
			exec, _ := ai.GetString(section, KeyExec)
			action := DesktopAction{
				Name:   name,
				Exec:   exec,
				parent: ai,
			}
			actions = append(actions, action)
		}
	}
	return actions
}

type DesktopAction struct {
	parent *DesktopAppInfo
	Name   string
	Exec   string
}

func (action *DesktopAction) Launch(timestamp uint32, files []string) error {
	ai := action.parent
	return _launch(ai, action.Exec, timestamp, files)
}

func stringSliceContains(slice []string, str string) bool {
	for _, v := range slice {
		if str == v {
			return true
		}
	}
	return false
}

var launchCount uint32

func getStartupNotifyId(exec string, timestamp uint32) string {
	pid := os.Getpid()
	prog := filepath.Base(os.Args[0])
	hostname, _ := os.Hostname()
	execBase := filepath.Base(exec)
	launchCount++
	return fmt.Sprintf("%s-%d-%s-%s-%d_TIME%d", prog, pid, hostname, execBase, launchCount, timestamp)
}
