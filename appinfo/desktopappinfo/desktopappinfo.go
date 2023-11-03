// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package desktopappinfo

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/godbus/dbus/v5"
	gio "github.com/linuxdeepin/go-gir/gio-2.0"
	"github.com/linuxdeepin/go-lib/appinfo"
	"github.com/linuxdeepin/go-lib/keyfile"
	"github.com/linuxdeepin/go-lib/shell"
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

	envDesktopEnv         = "XDG_CURRENT_DESKTOP"
	desktopExt            = ".desktop"
	gsSchemaStartdde      = "com.deepin.dde.startdde"
	enableInvoker         = "ENABLE_TURBO_INVOKER"
	turboInvokerFailedMsg = "Failed to invoke: Booster:"
	turboInvokerErrMsg    = "deepin-turbo-invoker: error"
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

func (ai *DesktopAppInfo) Launch(files []string, launchContext *appinfo.AppLaunchContext) error {
	return launch(ai, ai.GetCommandline(), files, launchContext)
}

func (ai *DesktopAppInfo) StartCommand(files []string, launchContext *appinfo.AppLaunchContext) (*exec.Cmd, error) {
	return startCommand(ai, ai.GetCommandline(), files, launchContext, false)
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

var _startddeGs *gio.Settings
var _startddeGsMu sync.Mutex

func getStartddeGs() *gio.Settings {
	_startddeGsMu.Lock()
	if _startddeGs == nil {
		_startddeGs = gio.NewSettings(gsSchemaStartdde)
	}
	_startddeGsMu.Unlock()
	return _startddeGs
}

// isAction 表示是否是启动一个 desktop action。
func shouldUseTurboInvoker(ai *DesktopAppInfo, isAction bool, turboInvokerPath string, ctx *appinfo.AppLaunchContext) bool {
	// NOTE: 启用 turbo invoker 的条件：
	// 必须 isAction 为 false，即必须不是 desktop action 快捷动作，而是 desktop 文件主要的那个。
	// 必须有 deepin-turbo-invoker 程序，它在 deepin-turbo 包。
	// 必须 desktop 文件中有 X-Deepin-TurboType 字段的值不为空。
	// 如果开了应用代理添加了 prefix 或 suffix 则不用 turbo。
	// 环境变量 ENABLE_TURBO_INVOKER  == 1
	// 或者环境变量 ENABLE_TURBO_INVOKER  == 空且 gsettings com.deepin.dde.startdde turbo-invoker-enabled 设置为 true。
	// 因此当环境变量 ENABLE_TURBO_INVOKER == 0 时，可以禁用此功能。
	if isAction || turboInvokerPath == "" {
		return false
	}

	turboType, _ := ai.GetString(MainSection, "X-Deepin-TurboType")
	if turboType == "" {
		return false
	}

	// 应用代理是通过添加 prefix 和 suffix 来实现的
	for _, prefix := range ctx.GetCmdPrefixes() {
		if strings.Contains(prefix, "proxychains") {
			return false
		}
	}
	for _, suffix := range ctx.GetCmdSuffixes() {
		if strings.Contains(suffix, "--proxy-server") {
			return false
		}
	}

	envEnabled := os.Getenv(enableInvoker)
	if envEnabled == "1" {
		return true
	}

	gsEnabled := getStartddeGs().GetBoolean("turbo-invoker-enabled")
	return envEnabled == "" && gsEnabled
}

type failDetector struct {
	buf  bytes.Buffer
	done bool // 是否完成检测
	ch   chan struct{}
	mu   sync.Mutex
}

func newFailDetector(ch chan struct{}) *failDetector {
	return &failDetector{
		ch: ch,
	}
}

func (w *failDetector) markDone() {
	w.mu.Lock()

	w.done = true
	w.buf = bytes.Buffer{}

	w.mu.Unlock()
}

func (w *failDetector) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.done {
		return len(p), nil
	}

	n, err = w.buf.Write(p)
	data := w.buf.Bytes()
	if bytes.Contains(data, []byte(turboInvokerFailedMsg)) || bytes.Contains(data, []byte(turboInvokerErrMsg)) {
		w.done = true
		w.buf = bytes.Buffer{}
		// 告知 turbo invoker 启动失败
		w.ch <- struct{}{}
	} else {
		// 把最后一个 \n 前面的数据清除
		idx := bytes.LastIndexByte(data, '\n')
		if idx != -1 {
			newData := make([]byte, len(data)-idx-1)
			copy(newData, data[idx+1:])
			w.buf.Reset()
			w.buf.Write(newData)
		}
	}
	// 防止 buf 数据量太大
	if w.buf.Len() > 50000 {
		w.done = true
		w.buf = bytes.Buffer{} // 释放 w.buf 占用的内存
	}
	return n, err
}

func startCommand(ai *DesktopAppInfo, cmdline string, files []string, launchContext *appinfo.AppLaunchContext, isAction bool) (*exec.Cmd, error) {
	turboInvokerPath, _ := exec.LookPath("deepin-turbo-invoker")

	if shouldUseTurboInvoker(ai, isAction, turboInvokerPath, launchContext) {
		args := []string{"--no-wait", "--desktop-file"}
		args = append(args, ai.GetFileName())

		if ai.IsDesktopOverrideExecSet() {
			args = append(args, "--desktop-override-exec="+ai.GetDesktopOverrideExec())
		}

		for _, file := range files {
			args = append(args, file)
		}
		cmd := exec.Command(turboInvokerPath, args...)
		failCh := make(chan struct{}, 2)
		fdOut := newFailDetector(failCh)
		fdErr := newFailDetector(failCh)
		cmd.Stdout = fdOut
		cmd.Stderr = fdErr

		err := cmd.Start()
		if err == nil {
			select {
			case <-time.After(2 * time.Second):
				// turbo invoker 启动成功
				fdOut.markDone()
				fdErr.markDone()
				return cmd, nil
			case <-failCh:
				// turbo invoker 启动失败
				go func() {
					// NOTE: 注意回收 cmd 进程资源
					_ = cmd.Wait()
				}()
			}
		}
	}

	if cmdline == "" {
		return nil, errors.New("command line is empty")
	}

	// get working dir
	workingDir := ai.GetPath()
	if workingDir == "" {
		// fallback to user home dir
		workingDir = basedir.GetUserHomeDir()

		// fallback to fs root /
		if workingDir == "" {
			workingDir = "/"
		}
	}

	exeargs, err := splitExec(cmdline)
	if err != nil {
		return nil, err
	}

	exeargs, err = ai.expandFieldCode(exeargs, files)
	if err != nil {
		return nil, err
	}

	useTerminal := ai.GetTerminal()
	if useTerminal {
		termExec, termExecArg := getDefaultTerminal()
		exeargs = append([]string{termExec, termExecArg}, exeargs...)
	}
	env := os.Environ()
	if launchContext != nil {
		cmdPrefixes := launchContext.GetCmdPrefixes()
		cmdSuffixes := launchContext.GetCmdSuffixes()
		exeargs = append(cmdPrefixes, exeargs...)
		exeargs = append(exeargs, cmdSuffixes...)
		if launchContext.GetEnv() != nil {
			env = launchContext.GetEnv()
		}
	}

	var launchScriptBuf bytes.Buffer
	launchScriptBuf.WriteString(`export GIO_LAUNCHED_DESKTOP_FILE_PID=$$;exec`)
	for _, arg := range exeargs {
		launchScriptBuf.WriteByte(' ')
		launchScriptBuf.WriteString(shell.Encode(arg))
	}

	cmd := exec.Command("/bin/sh", "-c", launchScriptBuf.String())
	cmd.Env = append(env, "GIO_LAUNCHED_DESKTOP_FILE="+ai.GetFileName())
	cmd.Dir = workingDir

	var snId string
	startupNotify := ai.GetStartupNotify()
	if startupNotify && launchContext != nil &&
		launchContext.GetTimestamp() != 0 {
		snId, _ = launchContext.GetStartupNotifyId(ai, files)
		cmd.Env = append(cmd.Env, "DESKTOP_STARTUP_ID="+snId)
	}

	err = cmd.Start()
	return cmd, err
}

func launch(ai *DesktopAppInfo, cmdline string, files []string, launchContext *appinfo.AppLaunchContext) error {
	cmd, err := startCommand(ai, cmdline, files, launchContext, false)
	go func() {
		_ = cmd.Wait()
	}()
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
				Section: section,
				Name:    name,
				Exec:    exec,
				parent:  ai,
			}
			actions = append(actions, action)
		}
	}
	return actions
}

type DesktopAction struct {
	parent *DesktopAppInfo

	Section string
	Name    string
	Exec    string
}

func (action *DesktopAction) Launch(files []string, launchContext *appinfo.AppLaunchContext) error {
	ai := action.parent
	return launch(ai, action.Exec, files, launchContext)
}

func (action *DesktopAction) StartCommand(files []string, launchContext *appinfo.AppLaunchContext) (*exec.Cmd, error) {
	ai := action.parent
	return startCommand(ai, action.Exec, files, launchContext, true)
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
