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
	"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"pkg.deepin.io/lib/keyfile"
)

func TestNewDesktopAppInfoFromKeyFile(t *testing.T) {
	Convey("NewDesktopAppInfoFromKeyFile failed, empty keyfile", t, func(c C) {
		kfile := keyfile.NewKeyFile()
		ai, err := NewDesktopAppInfoFromKeyFile(kfile)
		c.So(err, ShouldNotBeNil)
		c.So(err, ShouldEqual, ErrInvalidFirstSection)
		c.So(ai, ShouldBeNil)
		t.Log(err)
	})

	Convey("NewDesktopAppInfoFromKeyFile failed, invalid type", t, func(c C) {
		kfile := keyfile.NewKeyFile()
		kfile.SetValue(MainSection, KeyType, "xxxx")
		kfile.SetValue(MainSection, KeyName, "Deepin Screenshot")
		ai, err := NewDesktopAppInfoFromKeyFile(kfile)
		c.So(err, ShouldNotBeNil)
		c.So(err, ShouldEqual, ErrInvalidType)
		c.So(ai, ShouldBeNil)
		t.Log(err)

	})

	Convey("NewDesktopAppInfoFromKeyFile sucess", t, func(c C) {
		kfile := keyfile.NewKeyFile()
		kfile.SetValue(MainSection, KeyType, TypeApplication)
		kfile.SetValue(MainSection, KeyName, "Deepin Screenshot")
		kfile.SetValue(MainSection, KeyExec, "deepin-screenshot --icon")
		kfile.SetValue(MainSection, KeyIcon, "deepin-screenshot.png")
		ai, err := NewDesktopAppInfoFromKeyFile(kfile)
		c.So(err, ShouldBeNil)
		c.So(ai, ShouldNotBeNil)

		c.So(ai.GetCommandline(), ShouldEqual, "deepin-screenshot --icon")
		c.So(ai.GetName(), ShouldEqual, "Deepin Screenshot")
		c.So(ai.GetIcon(), ShouldEqual, "deepin-screenshot")
	})
}

func TestNewDesktopAppInfo(t *testing.T) {
	Convey("NewDesktopAppInfo", t, func(c C) {
		testDataDir, err := filepath.Abs("testdata")
		c.So(err, ShouldBeNil)
		xdgDataDirs = []string{testDataDir}
		xdgAppDirs = []string{filepath.Join(testDataDir, "applications")}

		c.Convey("NewDesktopAppInfoFromFile", func(c C) {
			ai, err := NewDesktopAppInfoFromFile("testdata/applications/deepin-screenshot.desktop")
			c.So(err, ShouldBeNil)
			c.So(ai, ShouldNotBeNil)

			c.So(err, ShouldBeNil)
			c.So(ai.GetIcon(), ShouldEqual, "deepin-screenshot")
			c.So(ai.GetId(), ShouldEqual, "deepin-screenshot")
		})

		c.Convey("NewDesktopAppInfo By Id", func(c C) {
			ai := NewDesktopAppInfo("deepin-screenshot")
			c.So(ai, ShouldNotBeNil)
			c.So(ai.GetId(), ShouldEqual, "deepin-screenshot")

			ai = NewDesktopAppInfo("deepin-screenshot.desktop")
			c.So(ai, ShouldNotBeNil)

			ai = NewDesktopAppInfo("not-exist")
			c.So(ai, ShouldBeNil)
		})

	})
}

func TestGetActions(t *testing.T) {
	Convey("DesktopAppInfo.GetActions", t, func(c C) {
		ai, err := NewDesktopAppInfoFromFile("testdata/applications/deepin-screenshot.desktop")
		c.So(err, ShouldBeNil)
		actions := ai.GetActions()
		t.Log("actions:", actions)
		c.So(actions, ShouldHaveLength, 2)
		c.So(actions[0].Exec, ShouldEqual, "deepin-screenshot -f")
		c.So(actions[1].Exec, ShouldEqual, "deepin-screenshot -d 5")
	})
}

func init() {
	os.Setenv(envDesktopEnv, "Deepin:GNOME")
}

func TestGetCurrentDestkops(t *testing.T) {
	Convey("getCurrentDesktop", t, func(c C) {
		c.So(getCurrentDesktop(), ShouldResemble, []string{"Deepin", "GNOME"})
	})
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
	Convey("GetShowIn destkop env: Deepin", t, func(c C) {
		currentDesktops = []string{"Deepin"}

		c.Convey("GetShowIn OnlyShowIn=Deepin", func(c C) {
			kfile := keyfile.NewKeyFile()
			err := kfile.LoadFromData([]byte(desktopFileContent0))
			c.So(err, ShouldBeNil)
			ai, err := NewDesktopAppInfoFromKeyFile(kfile)
			c.So(err, ShouldBeNil)
			c.So(ai, ShouldNotBeNil)
			c.So(ai.GetShowIn(nil), ShouldBeTrue)
		})

		c.Convey("GetShowIn OnlyShowIn=GNOME", func(c C) {
			kfile := keyfile.NewKeyFile()
			err := kfile.LoadFromData([]byte(desktopFileContent1))
			c.So(err, ShouldBeNil)
			ai, err := NewDesktopAppInfoFromKeyFile(kfile)
			c.So(err, ShouldBeNil)
			c.So(ai, ShouldNotBeNil)
			c.So(ai.GetShowIn(nil), ShouldBeFalse)
		})

		c.Convey("GetShowIn OnlyShowIn undefined and NotShowIn undefined", func(c C) {
			kfile := keyfile.NewKeyFile()
			err := kfile.LoadFromData([]byte(desktopFileContent2))
			c.So(err, ShouldBeNil)
			ai, err := NewDesktopAppInfoFromKeyFile(kfile)
			c.So(err, ShouldBeNil)
			c.So(ai, ShouldNotBeNil)
			c.So(ai.GetShowIn(nil), ShouldBeTrue)
		})

		c.Convey("GetShowIn OnlyShowIn=GNOME;Deepin", func(c C) {
			kfile := keyfile.NewKeyFile()
			err := kfile.LoadFromData([]byte(desktopFileContent3))
			c.So(err, ShouldBeNil)
			ai, err := NewDesktopAppInfoFromKeyFile(kfile)
			c.So(err, ShouldBeNil)
			c.So(ai, ShouldNotBeNil)
			c.So(ai.GetShowIn(nil), ShouldBeTrue)
		})

		c.Convey("GetShowIn OnlyShowIn=Xfce;KDE", func(c C) {
			kfile := keyfile.NewKeyFile()
			err := kfile.LoadFromData([]byte(desktopFileContent4))
			c.So(err, ShouldBeNil)
			ai, err := NewDesktopAppInfoFromKeyFile(kfile)
			c.So(err, ShouldBeNil)
			c.So(ai, ShouldNotBeNil)
			c.So(ai.GetShowIn(nil), ShouldBeFalse)
		})

		c.Convey("GetShowIn NotShowIn=Deepin", func(c C) {
			kfile := keyfile.NewKeyFile()
			err := kfile.LoadFromData([]byte(desktopFileContent5))
			c.So(err, ShouldBeNil)
			ai, err := NewDesktopAppInfoFromKeyFile(kfile)
			c.So(err, ShouldBeNil)
			c.So(ai, ShouldNotBeNil)
			c.So(ai.GetShowIn(nil), ShouldBeFalse)
		})

		c.Convey("GetShowIn NotShowIn=KDE;GNOME;", func(c C) {
			kfile := keyfile.NewKeyFile()
			err := kfile.LoadFromData([]byte(desktopFileContent6))
			c.So(err, ShouldBeNil)
			ai, err := NewDesktopAppInfoFromKeyFile(kfile)
			c.So(err, ShouldBeNil)
			c.So(ai, ShouldNotBeNil)
			c.So(ai.GetShowIn(nil), ShouldBeTrue)
		})
	})
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
	Convey("GetExecutable Exec undefined", t, func(c C) {
		kfile := keyfile.NewKeyFile()
		err := kfile.LoadFromData([]byte(desktopFileContent6))
		c.So(err, ShouldBeNil)
		ai, err := NewDesktopAppInfoFromKeyFile(kfile)
		c.So(err, ShouldBeNil)
		c.So(ai, ShouldNotBeNil)
		c.So(ai.GetExecutable(), ShouldEqual, "")
		c.So(ai.IsExecutableOk(), ShouldBeFalse)
	})

	Convey("GetExecutable Exec=sh", t, func(c C) {
		kfile := keyfile.NewKeyFile()
		err := kfile.LoadFromData([]byte(desktopFileContent7))
		c.So(err, ShouldBeNil)
		ai, err := NewDesktopAppInfoFromKeyFile(kfile)
		c.So(err, ShouldBeNil)
		c.So(ai, ShouldNotBeNil)
		c.So(ai.GetExecutable(), ShouldEqual, "sh")
		c.So(ai.IsExecutableOk(), ShouldBeTrue)
	})

	Convey("GetExecutable Exec=sh $", t, func(c C) {
		kfile := keyfile.NewKeyFile()
		err := kfile.LoadFromData([]byte(desktopFileContent8))
		c.So(err, ShouldBeNil)
		ai, err := NewDesktopAppInfoFromKeyFile(kfile)
		c.So(err, ShouldBeNil)
		c.So(ai, ShouldNotBeNil)
		c.So(ai.GetExecutable(), ShouldEqual, "")
		c.So(ai.IsExecutableOk(), ShouldBeFalse)
	})

	Convey("GetExecutable Exec=/bin/sh", t, func(c C) {
		kfile := keyfile.NewKeyFile()
		err := kfile.LoadFromData([]byte(desktopFileContent9))
		c.So(err, ShouldBeNil)
		ai, err := NewDesktopAppInfoFromKeyFile(kfile)
		c.So(err, ShouldBeNil)
		c.So(ai, ShouldNotBeNil)
		c.So(ai.GetExecutable(), ShouldEqual, "/bin/sh")
		c.So(ai.IsExecutableOk(), ShouldBeTrue)
	})

	Convey("GetExecutable Exec=/bin/sh/notexist", t, func(c C) {
		kfile := keyfile.NewKeyFile()
		err := kfile.LoadFromData([]byte(desktopFileContent10))
		c.So(err, ShouldBeNil)
		ai, err := NewDesktopAppInfoFromKeyFile(kfile)
		c.So(err, ShouldBeNil)
		c.So(ai, ShouldNotBeNil)
		c.So(ai.GetExecutable(), ShouldEqual, "/bin/sh/notexist")
		c.So(ai.IsExecutableOk(), ShouldBeFalse)
	})

	Convey("GetExecutable Exec=notexist", t, func(c C) {
		kfile := keyfile.NewKeyFile()
		err := kfile.LoadFromData([]byte(desktopFileContent11))
		c.So(err, ShouldBeNil)
		ai, err := NewDesktopAppInfoFromKeyFile(kfile)
		c.So(err, ShouldBeNil)
		c.So(ai, ShouldNotBeNil)
		c.So(ai.GetExecutable(), ShouldEqual, "notexist")
		c.So(ai.IsExecutableOk(), ShouldBeFalse)
	})
}

func Test_splitExec(t *testing.T) {
	Convey("splitExec", t, func(c C) {
		parts, err := splitExec(`abc def ghi`)
		c.So(err, ShouldBeNil)
		c.So(parts, ShouldResemble, []string{"abc", "def", "ghi"})

		parts, err = splitExec(`"abc" def "ghi"`)
		c.So(err, ShouldBeNil)
		c.So(parts, ShouldResemble, []string{"abc", "def", "ghi"})

		parts, err = splitExec(`"abc 123" def "ghi"`)
		c.So(err, ShouldBeNil)
		c.So(parts, ShouldResemble, []string{"abc 123", "def", "ghi"})

		parts, err = splitExec(`abc def "" ghi`)
		c.So(err, ShouldBeNil)
		c.So(parts, ShouldResemble, []string{"abc", "def", "", "ghi"})

		parts, err = splitExec(`"abc's" def`)
		c.So(err, ShouldBeNil)
		c.So(parts, ShouldResemble, []string{"abc's", "def"})

		parts, err = splitExec(`"abc\\s" def`)
		c.So(err, ShouldBeNil)
		c.So(parts, ShouldResemble, []string{"abc\\s", "def"})

		parts, err = splitExec(`abc   def ghi`)
		c.So(err, ShouldBeNil)
		c.So(parts, ShouldResemble, []string{"abc", "def", "ghi"})

		parts, err = splitExec(`"abc"`)
		c.So(err, ShouldBeNil)
		c.So(parts, ShouldResemble, []string{"abc"})

		parts, err = splitExec(`"$abcdef"`)
		c.So(err, ShouldBeNil)
		c.So(parts, ShouldResemble, []string{"$abcdef"})

		parts, err = splitExec("\"`abcdef\" def")
		c.So(err, ShouldBeNil)
		c.So(parts, ShouldResemble, []string{"`abcdef", "def"})

		parts, err = splitExec(`sh -c 'if ! [ -e "/usr/bin/ibus-daemon" ] && [ "x$XDG_SESSION_TYPE" = "xwayland" ] ; then exec env IM_CONFIG_CHECK_ENV=1 deepin-terminal true; fi'`)
		c.So(err, ShouldBeNil)
		c.So(parts, ShouldResemble, []string{"sh", "-c", `if ! [ -e "/usr/bin/ibus-daemon" ] && [ "x$XDG_SESSION_TYPE" = "xwayland" ] ; then exec env IM_CONFIG_CHECK_ENV=1 deepin-terminal true; fi`})

		_, err = splitExec(`"abcdef`)
		c.So(err, ShouldEqual, ErrQuotingNotClosed)

		_, err = splitExec(`"abcdef\"`)
		c.So(err, ShouldEqual, ErrQuotingNotClosed)

		_, err = splitExec(`"abc\def"`)
		c.So(err, ShouldResemble, ErrInvalidEscapeSequence{'d'})

		_, err = splitExec(`#echo hello world`)
		c.So(err, ShouldResemble, ErrReservedCharNotQuoted{'#'})

		_, err = splitExec(`(1)`)
		c.So(err, ShouldResemble, ErrReservedCharNotQuoted{'('})

		_, err = splitExec(`"abc"def`)
		c.So(err, ShouldEqual, ErrNoSpaceAfterQuoting)

		parts, err = splitExec(`env WINEPREFIX="/home/tp/.wine" wine-stable C:\\windows\\command\\start.exe /Unix /home/tp/.wine/dosdevices/c:/users/tp/Start\ Menu/Programs/sc1.08/sc1.08.lnk`)
		c.So(err, ShouldBeNil)
		c.So(parts, ShouldResemble, []string{"env", "WINEPREFIX=/home/tp/.wine",
			"wine-stable", "C:\\\\windows\\\\command\\\\start.exe",
			"/Unix", "/home/tp/.wine/dosdevices/c:/users/tp/Start Menu/Programs/sc1.08/sc1.08.lnk"})

		parts, err = splitExec(`echo hello\ world`)
		c.So(err, ShouldBeNil)
		c.So(parts, ShouldResemble, []string{"echo", "hello world"})

		parts, err = splitExec(`echo hello\world`)
		c.So(err, ShouldBeNil)
		c.So(parts, ShouldResemble, []string{"echo", "helloworld"})

		_, err = splitExec(`echo hello\`)
		c.So(err, ShouldEqual, ErrEscapeCharAtEnd)
	})
}

func Test_expandFieldCode(t *testing.T) {
	Convey("expandFieldCode", t, func(c C) {
		icon := "test_icon"
		desktopFile := "test.desktop"
		files := []string{"/dir1/dir2/a", "/dir1/dir2/b"}
		translatedName := "translatedName"
		cmdline := []string{"start", "%f", "%i", "%c", "%k", "end"}
		parts, err := expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
		c.So(err, ShouldBeNil)
		c.So(parts, ShouldResemble, []string{"start", files[0], "--icon", icon, translatedName, desktopFile, "end"})

		cmdline = []string{"start", "%G", "end"}
		_, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
		c.So(err, ShouldEqual, ErrBadFieldCode)

		cmdline = []string{"start", "%d", "%d", "%v", "end"}
		parts, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
		c.So(err, ShouldBeNil)
		c.So(parts, ShouldResemble, []string{"start", "end"})

		cmdline = []string{"start", "%F", "end"}
		parts, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
		c.So(err, ShouldBeNil)
		c.So(parts, ShouldResemble, []string{"start", "/dir1/dir2/a", "/dir1/dir2/b", "end"})

		cmdline = []string{"start", "%u", "end"}
		parts, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
		c.So(err, ShouldBeNil)
		c.So(parts, ShouldResemble, []string{"start", "/dir1/dir2/a", "end"})

		cmdline = []string{"start", "%U", "end"}
		parts, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
		c.So(err, ShouldBeNil)
		c.So(parts, ShouldResemble, []string{"start", "/dir1/dir2/a", "/dir1/dir2/b", "end"})

		cmdline = []string{"start", "%%", "%abc", "end"}
		_, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
		c.So(err, ShouldEqual, ErrBadFieldCode)

		cmdline = []string{"start", "1%%", "end"}
		parts, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
		c.So(err, ShouldBeNil)
		c.So(parts, ShouldResemble, []string{"start", "1%", "end"})

		files = []string{"file:///home/tp/2017%E5%B9%B403%E6%9C%88-%E6%B7%B1%E5%BA%A6%E9%9B%86%E7%BB%93-%E7%94%B5%E5%AD%90%E7%89%88.pdf"}
		cmdline = []string{"/opt/Foxitreader/FoxitReader.sh", "%F"}
		parts, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
		c.So(err, ShouldBeNil)
		c.So(parts, ShouldResemble, []string{"/opt/Foxitreader/FoxitReader.sh", "/home/tp/2017年03月-深度集结-电子版.pdf"})

		files = []string{"/a/b/log"}
		cmdline = []string{"start", "file=%u", "end"}
		parts, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
		c.So(err, ShouldBeNil)
		c.So(parts, ShouldResemble, []string{"start", "file=/a/b/log", "end"})

		cmdline = []string{"start", "file=%u+++", "end"}
		parts, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
		c.So(err, ShouldBeNil)
		c.So(parts, ShouldResemble, []string{"start", "file=/a/b/log+++", "end"})

		cmdline = []string{"start", "icon:%i", "end"}
		parts, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
		c.So(err, ShouldBeNil)
		c.So(parts, ShouldResemble, []string{"start", "icon:--icon", "test_icon", "end"})
	})
}
