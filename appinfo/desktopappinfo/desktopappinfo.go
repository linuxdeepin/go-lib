// SPDX-FileCopyrightText: 2018 - 2026 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package desktopappinfo

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/godbus/dbus/v5"
	"github.com/linuxdeepin/go-lib/keyfile"
	"github.com/linuxdeepin/go-lib/xdg/basedir"
)

const (
	MainSection        = "Desktop Entry"
	KeyType            = "Type"
	KeyVersion         = "Version"
	KeyName            = "Name"
	KeyGenericName     = "GenericName"
	KeyNoDisplay       = "NoDisplay"
	KeyComment         = "Comment"
	KeyIcon            = "Icon"
	KeyHidden          = "Hidden"
	KeyOnlyShowIn      = "OnlyShowIn"
	KeyNotShowIn       = "NotShowIn"
	KeyTryExec         = "TryExec"
	KeyExec            = "Exec"
	KeyPath            = "Path"
	KeyTerminal        = "Terminal"
	KeyMimeType        = "MimeType"
	KeyCategories      = "Categories"
	KeyKeywords        = "Keywords"
	KeyStartupNotify   = "StartupNotify"
	KeyStartupWMClass  = "StartupWMClass"
	KeyURL             = "URL"
	KeyActions         = "Actions"
	KeyDBusActivatable = "DBusActivatable"

	TypeApplication = "Application"
	TypeLink        = "Link"
	TypeDirectory   = "Directory"

	envDesktopEnv = "XDG_CURRENT_DESKTOP"
	desktopExt    = ".desktop"
)

// NOTE: these consts is copied from systemd-go
// https://github.com/coreos/go-systemd/blob/d843340ab4bd3815fda02e648f9b09ae2dc722a7/dbus/dbus.go#L30-L35
const (
	alpha    = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
	num      = `0123456789`
	alphaNum = alpha + num
)

var xdgDataDirs []string
var xdgAppDirs []string

func init() {
	xdgDataDirs = make([]string, 0, 3)
	xdgDataDirs = append(xdgDataDirs, basedir.GetUserDataDir())
	sysDataDirs := basedir.GetSystemDataDirs()
	xdgDataDirs = append(xdgDataDirs, sysDataDirs...)

	xdgAppDirs = make([]string, len(xdgDataDirs))
	for i, dir := range xdgDataDirs {
		xdgAppDirs[i] = filepath.Join(dir, "applications")
	}
}

func SetDataDirs(dirs []string) {
	xdgDataDirs = dirs
	xdgAppDirs = make([]string, len(xdgDataDirs))
	for i, dir := range xdgDataDirs {
		xdgAppDirs[i] = filepath.Join(dir, "applications")
	}
}

type DesktopAppInfo struct {
	*keyfile.KeyFile
	filename     string
	id           string
	name         string
	icon         string
	overrideExec string
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

// deprecated
func (ai *DesktopAppInfo) GetIsHiden() bool {
	hidden, _ := ai.GetBool(MainSection, KeyHidden)
	return hidden
}

func (ai *DesktopAppInfo) GetIsHidden() bool {
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
	if ai.GetIsHidden() {
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

func (ai *DesktopAppInfo) GetComment() string {
	comment, _ := ai.GetLocaleString(MainSection, KeyComment, "")
	return comment
}

func (ai *DesktopAppInfo) GetDisplayName() string {
	return ai.GetName()
}

func (ai *DesktopAppInfo) GetMimeTypes() []string {
	mimeTypes, _ := ai.GetStringList(MainSection, KeyMimeType)
	return mimeTypes
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

func (ai *DesktopAppInfo) GetPath() string {
	wd, _ := ai.GetString(MainSection, KeyPath)
	return wd
}

// TryExec is Path to an executable file on disk used to determine if the program is actually installed
func (ai *DesktopAppInfo) GetTryExec() string {
	tryExec, _ := ai.GetString(MainSection, KeyTryExec)
	return tryExec
}

// DBusActivatable is a boolean value specifying if D-Bus activation is supported for this application
func (ai *DesktopAppInfo) GetDBusActivatable() bool {
	b, _ := ai.GetBool(MainSection, KeyDBusActivatable)
	return b
}

func (ai *DesktopAppInfo) GetTerminal() bool {
	useTerminal, _ := ai.GetBool(MainSection, KeyTerminal)
	return useTerminal
}

func (ai *DesktopAppInfo) GetExecutable() string {
	cmdline := ai.GetCommandline()
	if cmdline == "" {
		return ""
	}
	parts, _ := splitExec(cmdline)
	if len(parts) == 0 {
		return ""
	}
	return parts[0]
}

func (ai *DesktopAppInfo) IsExecutableOk() bool {
	tryExec := ai.GetTryExec()
	if tryExec != "" {
		_, err := exec.LookPath(tryExec)
		return err == nil
	}

	exe := ai.GetExecutable()
	if exe == "" {
		return false
	}
	_, err := exec.LookPath(exe)
	return err == nil
}

func (ai *DesktopAppInfo) IsDesktopOverrideExecSet() bool {
	return ai.overrideExec != ""
}

func (ai *DesktopAppInfo) GetDesktopOverrideExec() string {
	return ai.overrideExec
}

func (ai *DesktopAppInfo) SetDesktopOverrideExec(exec string) {
	ai.overrideExec = exec
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
				Section: section,
				Name:    name,
				Exec:    exec,
			}
			actions = append(actions, action)
		}
	}
	return actions
}

type DesktopAction struct {
	Section string
	Name    string
	Exec    string
}

// PathBusEscape sanitizes a constituent string of a dbus ObjectPath using the
// rules that systemd uses for serializing special characters.
// NOTE: this function is copied from systemd-go
// https://github.com/coreos/go-systemd/blob/d843340ab4bd3815fda02e648f9b09ae2dc722a7/dbus/dbus.go#L47
func pathBusEscape(path string) string {
	// Special case the empty string
	if len(path) == 0 {
		return "_"
	}
	n := []byte{}
	for i := 0; i < len(path); i++ {
		c := path[i]
		if needsEscape(i, c) {
			e := fmt.Sprintf("_%x", c)
			n = append(n, []byte(e)...)
		} else {
			n = append(n, c)
		}
	}
	return string(n)
}

// needsEscape checks whether a byte in a potential dbus ObjectPath needs to be escaped
// NOTE: this function is copied from systemd-go
// https://github.com/coreos/go-systemd/blob/d843340ab4bd3815fda02e648f9b09ae2dc722a7/dbus/dbus.go#L38
func needsEscape(i int, b byte) bool {
	// Escape everything that is not a-z-A-Z-0-9
	// Also escape 0-9 if it's the first character
	return strings.IndexByte(alphaNum, b) == -1 ||
		(i == 0 && strings.IndexByte(num, b) != -1)
}

func GetDBusObjectFromAppDesktop(desktop string, service string, path string) (dbus.ObjectPath, error) {
	sessionBus, err := dbus.SessionBus()
	if err != nil {
		return "", err
	}

	escapeId := pathBusEscape(strings.TrimSuffix(desktop, ".desktop"))
	return sessionBus.Object(
		service,
		dbus.ObjectPath(path+"/"+escapeId),
	).Path(), nil
}
