/*
 * Copyright (C) 2017 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package desktopappinfo

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	gio "pkg.deepin.io/gir/gio-2.0"
	"pkg.deepin.io/lib/appinfo"
	"pkg.deepin.io/lib/keyfile"
	"pkg.deepin.io/lib/shell"
	"pkg.deepin.io/lib/xdg/basedir"
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
		select {
		case w.ch <- struct{}{}:
		// 告知 turbo invoker 启动失败
		default:
			// 表示 ch 已经没有接收者了
		}
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
		args := []string{"--desktop-file"}
		args = append(args, ai.GetFileName())
		for _, file := range files {
			args = append(args, file)
		}
		cmd := exec.Command(turboInvokerPath, args...)
		failCh := make(chan struct{})
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

	if launchContext != nil {
		cmdPrefixes := launchContext.GetCmdPrefixes()
		cmdSuffixes := launchContext.GetCmdSuffixes()
		exeargs = append(cmdPrefixes, exeargs...)
		exeargs = append(exeargs, cmdSuffixes...)
	}

	var launchScriptBuf bytes.Buffer
	launchScriptBuf.WriteString(`export GIO_LAUNCHED_DESKTOP_FILE_PID=$$;exec`)
	for _, arg := range exeargs {
		launchScriptBuf.WriteByte(' ')
		launchScriptBuf.WriteString(shell.Encode(arg))
	}

	cmd := exec.Command("/bin/sh", "-c", launchScriptBuf.String())
	cmd.Env = append(os.Environ(), "GIO_LAUNCHED_DESKTOP_FILE="+ai.GetFileName())
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
