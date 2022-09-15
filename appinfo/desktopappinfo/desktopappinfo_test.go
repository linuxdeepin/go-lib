// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package desktopappinfo

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/linuxdeepin/go-lib/keyfile"
)

func TestNewDesktopAppInfoFromKeyFile(t *testing.T) {
	kfile := keyfile.NewKeyFile()
	ai, err := NewDesktopAppInfoFromKeyFile(kfile)
	assert.Error(t, err)
	assert.Equal(t, err, ErrInvalidFirstSection)
	require.Nil(t, ai)
	t.Log(err)

	kfile = keyfile.NewKeyFile()
	kfile.SetValue(MainSection, KeyType, "xxxx")
	kfile.SetValue(MainSection, KeyName, "Deepin Screenshot")
	ai, err = NewDesktopAppInfoFromKeyFile(kfile)
	assert.Error(t, err)
	assert.Equal(t, err, ErrInvalidType)
	require.Nil(t, ai)
	t.Log(err)

	kfile = keyfile.NewKeyFile()
	kfile.SetValue(MainSection, KeyType, TypeApplication)
	kfile.SetValue(MainSection, KeyName, "Deepin Screenshot")
	kfile.SetValue(MainSection, KeyExec, "deepin-screenshot --icon")
	kfile.SetValue(MainSection, KeyIcon, "deepin-screenshot.png")
	ai, err = NewDesktopAppInfoFromKeyFile(kfile)
	require.NoError(t, err)
	assert.NotNil(t, ai)

	assert.Equal(t, ai.GetCommandline(), "deepin-screenshot --icon")
	assert.Equal(t, ai.GetName(), "Deepin Screenshot")
	assert.Equal(t, ai.GetIcon(), "deepin-screenshot")
}

func TestNewDesktopAppInfo(t *testing.T) {
	testDataDir, err := filepath.Abs("testdata")
	require.NoError(t, err)
	xdgDataDirs = []string{testDataDir}
	xdgAppDirs = []string{filepath.Join(testDataDir, "applications")}

	var testDesktopTypes []string
	testFileName := "testdata/applications/deepin-screenshot.desktop"
	ai, err := NewDesktopAppInfoFromFile(testFileName)
	require.NoError(t, err)
	assert.NotNil(t, ai)

	require.NoError(t, err)
	assert.Equal(t, ai.GetIcon(), "deepin-screenshot")
	assert.Equal(t, ai.GetId(), "deepin-screenshot")
	assert.Equal(t, ai.GetIsHiden(), false)
	assert.Equal(t, ai.GetIsHidden(), false)
	assert.Equal(t, ai.GetNoDisplay(), false)
	assert.Equal(t, ai.ShouldShow(), true)
	assert.Equal(t, ai.GetGenericName(), "Screen capturing application")
	assert.Equal(t, ai.GetComment(), "Screen capturing application")
	assert.Equal(t, ai.GetDisplayName(), "Deepin Screenshot")
	assert.Equal(t, ai.GetMimeTypes(), testDesktopTypes)
	assert.Equal(t, ai.GetCategories(), []string{"Graphics", "GTK"})
	assert.Equal(t, ai.GetKeywords(), testDesktopTypes)
	assert.Equal(t, ai.GetStartupWMClass(), "")
	assert.Equal(t, ai.GetStartupNotify(), false)
	assert.Equal(t, ai.GetPath(), "")
	assert.Equal(t, ai.GetDBusActivatable(), false)
	assert.Equal(t, ai.GetTerminal(), false)
	assert.Equal(t, ai.IsDesktopOverrideExecSet(), false)
	assert.Equal(t, ai.GetDesktopOverrideExec(), "")

	assert.Equal(t, getId(testFileName), "testdata/applications/deepin-screenshot")
	assert.Equal(t, getCurrentDesktop(), []string{"Deepin", "GNOME"})

	ai = NewDesktopAppInfo("deepin-screenshot")
	assert.NotNil(t, ai)
	assert.Equal(t, ai.GetId(), "deepin-screenshot")

	ai = NewDesktopAppInfo("deepin-screenshot.desktop")
	assert.NotNil(t, ai)

	ai = NewDesktopAppInfo("not-exist")
	require.Nil(t, ai)

}

func TestGetActions(t *testing.T) {
	ai, err := NewDesktopAppInfoFromFile("testdata/applications/deepin-screenshot.desktop")
	require.NoError(t, err)
	actions := ai.GetActions()
	t.Log("actions:", actions)
	assert.Len(t, actions, 2)
	assert.Equal(t, actions[0].Exec, "deepin-screenshot -f")
	assert.Equal(t, actions[1].Exec, "deepin-screenshot -d 5")
}

func init() {
	os.Setenv(envDesktopEnv, "Deepin:GNOME")
}

func TestGetCurrentDestkops(t *testing.T) {
	assert.Equal(t, getCurrentDesktop(), []string{"Deepin", "GNOME"})
}

const desktopFileContent0 = `
[Desktop Entry]
Type=Application
Name=xxx
OnlyShowIn=Deepin
`

const desktopFileContent1 = `
[Desktop Entry]
Type=Application
Name=xxx
OnlyShowIn=GNOME
`

const desktopFileContent2 = `
[Desktop Entry]
Type=Application
Name=xxx
`

const desktopFileContent3 = `
[Desktop Entry]
Type=Application
Name=xxx
OnlyShowIn=GNOME;Deepin
`

const desktopFileContent4 = `
[Desktop Entry]
Type=Application
Name=xxx
OnlyShowIn=Xfce;KDE
`

const desktopFileContent5 = `
[Desktop Entry]
Type=Application
Name=xxx
NotShowIn=Deepin
`

const desktopFileContent6 = `
[Desktop Entry]
Type=Application
Name=xxx
NotShowIn=KDE;GNOME;
`

func TestGetShowIn(t *testing.T) {
	currentDesktops = []string{"Deepin"}

	kfile := keyfile.NewKeyFile()
	err := kfile.LoadFromData([]byte(desktopFileContent0))
	require.NoError(t, err)
	ai, err := NewDesktopAppInfoFromKeyFile(kfile)
	require.NoError(t, err)
	assert.NotNil(t, ai)
	assert.True(t, ai.GetShowIn(nil))

	kfile = keyfile.NewKeyFile()
	err = kfile.LoadFromData([]byte(desktopFileContent1))
	require.NoError(t, err)
	ai, err = NewDesktopAppInfoFromKeyFile(kfile)
	require.NoError(t, err)
	assert.NotNil(t, ai)
	assert.False(t, ai.GetShowIn(nil))

	kfile = keyfile.NewKeyFile()
	err = kfile.LoadFromData([]byte(desktopFileContent2))
	require.NoError(t, err)
	ai, err = NewDesktopAppInfoFromKeyFile(kfile)
	require.NoError(t, err)
	assert.NotNil(t, ai)
	assert.True(t, ai.GetShowIn(nil))

	kfile = keyfile.NewKeyFile()
	err = kfile.LoadFromData([]byte(desktopFileContent3))
	require.NoError(t, err)
	ai, err = NewDesktopAppInfoFromKeyFile(kfile)
	require.NoError(t, err)
	assert.NotNil(t, ai)
	assert.True(t, ai.GetShowIn(nil))

	kfile = keyfile.NewKeyFile()
	err = kfile.LoadFromData([]byte(desktopFileContent4))
	require.NoError(t, err)
	ai, err = NewDesktopAppInfoFromKeyFile(kfile)
	require.NoError(t, err)
	assert.NotNil(t, ai)
	assert.False(t, ai.GetShowIn(nil))

	kfile = keyfile.NewKeyFile()
	err = kfile.LoadFromData([]byte(desktopFileContent5))
	require.NoError(t, err)
	ai, err = NewDesktopAppInfoFromKeyFile(kfile)
	require.NoError(t, err)
	assert.NotNil(t, ai)
	assert.False(t, ai.GetShowIn(nil))

	kfile = keyfile.NewKeyFile()
	err = kfile.LoadFromData([]byte(desktopFileContent6))
	require.NoError(t, err)
	ai, err = NewDesktopAppInfoFromKeyFile(kfile)
	require.NoError(t, err)
	assert.NotNil(t, ai)
	assert.True(t, ai.GetShowIn(nil))
}

const desktopFileContent7 = `
[Desktop Entry]
Type=Application
Name=shell
Exec=sh
`

const desktopFileContent8 = `
[Desktop Entry]
Type=Application
Name=shell
Exec=sh $
`

const desktopFileContent9 = `
[Desktop Entry]
Type=Application
Name=shell
Exec=/bin/sh
`

const desktopFileContent10 = `
[Desktop Entry]
Type=Application
Name=shell
Exec=/bin/sh/notexist
`

const desktopFileContent11 = `
[Desktop Entry]
Type=Application
Name=shell
Exec=notexist
`

func TestGetExecutable(t *testing.T) {
	kfile := keyfile.NewKeyFile()
	err := kfile.LoadFromData([]byte(desktopFileContent6))
	require.NoError(t, err)
	ai, err := NewDesktopAppInfoFromKeyFile(kfile)
	require.NoError(t, err)
	assert.NotNil(t, ai)
	assert.Equal(t, ai.GetExecutable(), "")
	assert.False(t, ai.IsExecutableOk())

	kfile = keyfile.NewKeyFile()
	err = kfile.LoadFromData([]byte(desktopFileContent7))
	require.NoError(t, err)
	ai, err = NewDesktopAppInfoFromKeyFile(kfile)
	require.NoError(t, err)
	assert.NotNil(t, ai)
	assert.Equal(t, ai.GetExecutable(), "sh")
	assert.True(t, ai.IsExecutableOk())

	kfile = keyfile.NewKeyFile()
	err = kfile.LoadFromData([]byte(desktopFileContent8))
	require.NoError(t, err)
	ai, err = NewDesktopAppInfoFromKeyFile(kfile)
	require.NoError(t, err)
	assert.NotNil(t, ai)
	assert.Equal(t, ai.GetExecutable(), "")
	assert.False(t, ai.IsExecutableOk())

	kfile = keyfile.NewKeyFile()
	err = kfile.LoadFromData([]byte(desktopFileContent9))
	require.NoError(t, err)
	ai, err = NewDesktopAppInfoFromKeyFile(kfile)
	require.NoError(t, err)
	assert.NotNil(t, ai)
	assert.Equal(t, ai.GetExecutable(), "/bin/sh")
	assert.True(t, ai.IsExecutableOk())

	kfile = keyfile.NewKeyFile()
	err = kfile.LoadFromData([]byte(desktopFileContent10))
	require.NoError(t, err)
	ai, err = NewDesktopAppInfoFromKeyFile(kfile)
	require.NoError(t, err)
	assert.NotNil(t, ai)
	assert.Equal(t, ai.GetExecutable(), "/bin/sh/notexist")
	assert.False(t, ai.IsExecutableOk())

	kfile = keyfile.NewKeyFile()
	err = kfile.LoadFromData([]byte(desktopFileContent11))
	require.NoError(t, err)
	ai, err = NewDesktopAppInfoFromKeyFile(kfile)
	require.NoError(t, err)
	assert.NotNil(t, ai)
	assert.Equal(t, ai.GetExecutable(), "notexist")
	assert.False(t, ai.IsExecutableOk())
}

func Test_splitExec(t *testing.T) {
	parts, err := splitExec(`abc def ghi`)
	require.NoError(t, err)
	assert.Equal(t, parts, []string{"abc", "def", "ghi"})

	parts, err = splitExec(`"abc" def "ghi"`)
	require.NoError(t, err)
	assert.Equal(t, parts, []string{"abc", "def", "ghi"})

	parts, err = splitExec(`"abc 123" def "ghi"`)
	require.NoError(t, err)
	assert.Equal(t, parts, []string{"abc 123", "def", "ghi"})

	parts, err = splitExec(`abc def "" ghi`)
	require.NoError(t, err)
	assert.Equal(t, parts, []string{"abc", "def", "", "ghi"})

	parts, err = splitExec(`"abc's" def`)
	require.NoError(t, err)
	assert.Equal(t, parts, []string{"abc's", "def"})

	parts, err = splitExec(`"abc\\s" def`)
	require.NoError(t, err)
	assert.Equal(t, parts, []string{"abc\\s", "def"})

	parts, err = splitExec(`abc   def ghi`)
	require.NoError(t, err)
	assert.Equal(t, parts, []string{"abc", "def", "ghi"})

	parts, err = splitExec(`"abc"`)
	require.NoError(t, err)
	assert.Equal(t, parts, []string{"abc"})

	parts, err = splitExec(`"$abcdef"`)
	require.NoError(t, err)
	assert.Equal(t, parts, []string{"$abcdef"})

	parts, err = splitExec("\"`abcdef\" def")
	require.NoError(t, err)
	assert.Equal(t, parts, []string{"`abcdef", "def"})

	parts, err = splitExec(`sh -c 'if ! [ -e "/usr/bin/ibus-daemon" ] && [ "x$XDG_SESSION_TYPE" = "xwayland" ] ; then exec env IM_CONFIG_CHECK_ENV=1 deepin-terminal true; fi'`)
	require.NoError(t, err)
	assert.Equal(t, parts, []string{"sh", "-c", `if ! [ -e "/usr/bin/ibus-daemon" ] && [ "x$XDG_SESSION_TYPE" = "xwayland" ] ; then exec env IM_CONFIG_CHECK_ENV=1 deepin-terminal true; fi`})

	_, err = splitExec(`"abcdef`)
	assert.Equal(t, err, ErrQuotingNotClosed)

	_, err = splitExec(`"abcdef\"`)
	assert.Equal(t, err, ErrQuotingNotClosed)

	_, err = splitExec(`"abc\def"`)
	assert.Equal(t, err, ErrInvalidEscapeSequence{'d'})

	_, err = splitExec(`#echo hello world`)
	assert.Equal(t, err, ErrReservedCharNotQuoted{'#'})

	_, err = splitExec(`(1)`)
	assert.Equal(t, err, ErrReservedCharNotQuoted{'('})

	_, err = splitExec(`"abc"def`)
	assert.Equal(t, err, ErrNoSpaceAfterQuoting)

	parts, err = splitExec(`env WINEPREFIX="/home/tp/.wine" wine-stable C:\\windows\\command\\start.exe /Unix /home/tp/.wine/dosdevices/c:/users/tp/Start\ Menu/Programs/sc1.08/sc1.08.lnk`)
	require.NoError(t, err)
	assert.Equal(t, parts, []string{"env", "WINEPREFIX=/home/tp/.wine",
		"wine-stable", "C:\\\\windows\\\\command\\\\start.exe",
		"/Unix", "/home/tp/.wine/dosdevices/c:/users/tp/Start Menu/Programs/sc1.08/sc1.08.lnk"})

	parts, err = splitExec(`echo hello\ world`)
	require.NoError(t, err)
	assert.Equal(t, parts, []string{"echo", "hello world"})

	parts, err = splitExec(`echo hello\world`)
	require.NoError(t, err)
	assert.Equal(t, parts, []string{"echo", "helloworld"})

	_, err = splitExec(`echo hello\`)
	assert.Equal(t, err, ErrEscapeCharAtEnd)
}

func Test_expandFieldCode(t *testing.T) {
	icon := "test_icon"
	desktopFile := "test.desktop"
	files := []string{"/dir1/dir2/a", "/dir1/dir2/b"}
	translatedName := "translatedName"
	cmdline := []string{"start", "%f", "%i", "%c", "%k", "end"}
	parts, err := expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
	require.NoError(t, err)
	assert.Equal(t, parts, []string{"start", files[0], "--icon", icon, translatedName, desktopFile, "end"})

	cmdline = []string{"start", "%G", "end"}
	_, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
	assert.Equal(t, err, ErrBadFieldCode)

	cmdline = []string{"start", "%d", "%d", "%v", "end"}
	parts, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
	require.NoError(t, err)
	assert.Equal(t, parts, []string{"start", "end"})

	cmdline = []string{"start", "%F", "end"}
	parts, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
	require.NoError(t, err)
	assert.Equal(t, parts, []string{"start", "/dir1/dir2/a", "/dir1/dir2/b", "end"})

	cmdline = []string{"start", "%u", "end"}
	parts, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
	require.NoError(t, err)
	assert.Equal(t, parts, []string{"start", "/dir1/dir2/a", "end"})

	cmdline = []string{"start", "%U", "end"}
	parts, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
	require.NoError(t, err)
	assert.Equal(t, parts, []string{"start", "/dir1/dir2/a", "/dir1/dir2/b", "end"})

	cmdline = []string{"start", "%%", "%abc", "end"}
	_, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
	assert.Equal(t, err, ErrBadFieldCode)

	cmdline = []string{"start", "1%%", "end"}
	parts, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
	require.NoError(t, err)
	assert.Equal(t, parts, []string{"start", "1%", "end"})

	files = []string{"file:///home/tp/2017%E5%B9%B403%E6%9C%88-%E6%B7%B1%E5%BA%A6%E9%9B%86%E7%BB%93-%E7%94%B5%E5%AD%90%E7%89%88.pdf"}
	cmdline = []string{"/opt/Foxitreader/FoxitReader.sh", "%F"}
	parts, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
	require.NoError(t, err)

	files = []string{"/a/b/log"}
	cmdline = []string{"start", "file=%u", "end"}
	parts, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
	require.NoError(t, err)
	assert.Equal(t, parts, []string{"start", "file=/a/b/log", "end"})

	cmdline = []string{"start", "file=%u+++", "end"}
	parts, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
	require.NoError(t, err)
	assert.Equal(t, parts, []string{"start", "file=/a/b/log+++", "end"})

	cmdline = []string{"start", "icon:%i", "end"}
	parts, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
	require.NoError(t, err)
	assert.Equal(t, parts, []string{"start", "icon:--icon", "test_icon", "end"})
}
