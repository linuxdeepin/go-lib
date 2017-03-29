package desktopappinfo

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"path/filepath"
	"pkg.deepin.io/lib/keyfile"
	"testing"
)

func TestNewDesktopAppInfoFromKeyFile(t *testing.T) {
	Convey("NewDesktopAppInfoFromKeyFile failed, empty keyfile", t, func() {
		kfile := keyfile.NewKeyFile()
		ai, err := NewDesktopAppInfoFromKeyFile(kfile)
		So(err, ShouldNotBeNil)
		So(err, ShouldEqual, ErrInvalidFirstSection)
		So(ai, ShouldBeNil)
		t.Log(err)
	})

	Convey("NewDesktopAppInfoFromKeyFile failed, invalid type", t, func() {
		kfile := keyfile.NewKeyFile()
		kfile.SetValue(MainSection, KeyType, "xxxx")
		kfile.SetValue(MainSection, KeyName, "Deepin Screenshot")
		ai, err := NewDesktopAppInfoFromKeyFile(kfile)
		So(err, ShouldNotBeNil)
		So(err, ShouldEqual, ErrInvalidType)
		So(ai, ShouldBeNil)
		t.Log(err)

	})

	Convey("NewDesktopAppInfoFromKeyFile sucess", t, func() {
		kfile := keyfile.NewKeyFile()
		kfile.SetValue(MainSection, KeyType, TypeApplication)
		kfile.SetValue(MainSection, KeyName, "Deepin Screenshot")
		kfile.SetValue(MainSection, KeyExec, "deepin-screenshot --icon")
		kfile.SetValue(MainSection, KeyIcon, "deepin-screenshot.png")
		ai, err := NewDesktopAppInfoFromKeyFile(kfile)
		So(err, ShouldBeNil)
		So(ai, ShouldNotBeNil)

		So(ai.GetCommandline(), ShouldEqual, "deepin-screenshot --icon")
		So(ai.GetName(), ShouldEqual, "Deepin Screenshot")
		So(ai.GetIcon(), ShouldEqual, "deepin-screenshot")
	})
}

func TestNewDesktopAppInfo(t *testing.T) {
	Convey("NewDesktopAppInfo", t, func() {
		testDataDir, err := filepath.Abs("testdata")
		So(err, ShouldBeNil)
		xdgDataDirs = []string{testDataDir}
		xdgAppDirs = []string{filepath.Join(testDataDir, "applications")}

		Convey("NewDesktopAppInfoFromFile", func() {
			ai, err := NewDesktopAppInfoFromFile("testdata/applications/deepin-screenshot.desktop")
			So(err, ShouldBeNil)
			So(ai, ShouldNotBeNil)

			So(err, ShouldBeNil)
			So(ai.GetIcon(), ShouldEqual, "deepin-screenshot")
			So(ai.GetId(), ShouldEqual, "deepin-screenshot")
		})

		Convey("NewDesktopAppInfo By Id", func() {
			ai := NewDesktopAppInfo("deepin-screenshot")
			So(ai, ShouldNotBeNil)
			So(ai.GetId(), ShouldEqual, "deepin-screenshot")

			ai = NewDesktopAppInfo("deepin-screenshot.desktop")
			So(ai, ShouldNotBeNil)

			ai = NewDesktopAppInfo("not-exist")
			So(ai, ShouldBeNil)
		})

	})
}

func init() {
	os.Setenv(envDesktopEnv, "Deepin:GNOME")
}

func TestGetCurrentDestkops(t *testing.T) {
	Convey("getCurrentDesktop", t, func() {
		So(getCurrentDesktop(), ShouldResemble, []string{"Deepin", "GNOME"})
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
	Convey("GetShowIn destkop env: Deepin", t, func() {
		currentDesktops = []string{"Deepin"}

		Convey("GetShowIn OnlyShowIn=Deepin", func() {
			kfile := keyfile.NewKeyFile()
			err := kfile.LoadFromData([]byte(desktopFileContent0))
			So(err, ShouldBeNil)
			ai, err := NewDesktopAppInfoFromKeyFile(kfile)
			So(err, ShouldBeNil)
			So(ai, ShouldNotBeNil)
			So(ai.GetShowIn(nil), ShouldBeTrue)
		})

		Convey("GetShowIn OnlyShowIn=GNOME", func() {
			kfile := keyfile.NewKeyFile()
			err := kfile.LoadFromData([]byte(desktopFileContent1))
			So(err, ShouldBeNil)
			ai, err := NewDesktopAppInfoFromKeyFile(kfile)
			So(err, ShouldBeNil)
			So(ai, ShouldNotBeNil)
			So(ai.GetShowIn(nil), ShouldBeFalse)
		})

		Convey("GetShowIn OnlyShowIn undefined and NotShowIn undefined", func() {
			kfile := keyfile.NewKeyFile()
			err := kfile.LoadFromData([]byte(desktopFileContent2))
			So(err, ShouldBeNil)
			ai, err := NewDesktopAppInfoFromKeyFile(kfile)
			So(err, ShouldBeNil)
			So(ai, ShouldNotBeNil)
			So(ai.GetShowIn(nil), ShouldBeTrue)
		})

		Convey("GetShowIn OnlyShowIn=GNOME;Deepin", func() {
			kfile := keyfile.NewKeyFile()
			err := kfile.LoadFromData([]byte(desktopFileContent3))
			So(err, ShouldBeNil)
			ai, err := NewDesktopAppInfoFromKeyFile(kfile)
			So(err, ShouldBeNil)
			So(ai, ShouldNotBeNil)
			So(ai.GetShowIn(nil), ShouldBeTrue)
		})

		Convey("GetShowIn OnlyShowIn=Xfce;KDE", func() {
			kfile := keyfile.NewKeyFile()
			err := kfile.LoadFromData([]byte(desktopFileContent4))
			So(err, ShouldBeNil)
			ai, err := NewDesktopAppInfoFromKeyFile(kfile)
			So(err, ShouldBeNil)
			So(ai, ShouldNotBeNil)
			So(ai.GetShowIn(nil), ShouldBeFalse)
		})

		Convey("GetShowIn NotShowIn=Deepin", func() {
			kfile := keyfile.NewKeyFile()
			err := kfile.LoadFromData([]byte(desktopFileContent5))
			So(err, ShouldBeNil)
			ai, err := NewDesktopAppInfoFromKeyFile(kfile)
			So(err, ShouldBeNil)
			So(ai, ShouldNotBeNil)
			So(ai.GetShowIn(nil), ShouldBeFalse)
		})

		Convey("GetShowIn NotShowIn=KDE;GNOME;", func() {
			kfile := keyfile.NewKeyFile()
			err := kfile.LoadFromData([]byte(desktopFileContent6))
			So(err, ShouldBeNil)
			ai, err := NewDesktopAppInfoFromKeyFile(kfile)
			So(err, ShouldBeNil)
			So(ai, ShouldNotBeNil)
			So(ai.GetShowIn(nil), ShouldBeTrue)
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
	Convey("GetExecutable Exec undefined", t, func() {
		kfile := keyfile.NewKeyFile()
		err := kfile.LoadFromData([]byte(desktopFileContent6))
		So(err, ShouldBeNil)
		ai, err := NewDesktopAppInfoFromKeyFile(kfile)
		So(err, ShouldBeNil)
		So(ai, ShouldNotBeNil)
		So(ai.GetExecutable(), ShouldEqual, "")
		So(ai.IsExecutableOk(), ShouldBeFalse)
	})

	Convey("GetExecutable Exec=sh", t, func() {
		kfile := keyfile.NewKeyFile()
		err := kfile.LoadFromData([]byte(desktopFileContent7))
		So(err, ShouldBeNil)
		ai, err := NewDesktopAppInfoFromKeyFile(kfile)
		So(err, ShouldBeNil)
		So(ai, ShouldNotBeNil)
		So(ai.GetExecutable(), ShouldEqual, "sh")
		So(ai.IsExecutableOk(), ShouldBeTrue)
	})

	Convey("GetExecutable Exec=sh $", t, func() {
		kfile := keyfile.NewKeyFile()
		err := kfile.LoadFromData([]byte(desktopFileContent8))
		So(err, ShouldBeNil)
		ai, err := NewDesktopAppInfoFromKeyFile(kfile)
		So(err, ShouldBeNil)
		So(ai, ShouldNotBeNil)
		So(ai.GetExecutable(), ShouldEqual, "")
		So(ai.IsExecutableOk(), ShouldBeFalse)
	})

	Convey("GetExecutable Exec=/bin/sh", t, func() {
		kfile := keyfile.NewKeyFile()
		err := kfile.LoadFromData([]byte(desktopFileContent9))
		So(err, ShouldBeNil)
		ai, err := NewDesktopAppInfoFromKeyFile(kfile)
		So(err, ShouldBeNil)
		So(ai, ShouldNotBeNil)
		So(ai.GetExecutable(), ShouldEqual, "/bin/sh")
		So(ai.IsExecutableOk(), ShouldBeTrue)
	})

	Convey("GetExecutable Exec=/bin/sh/notexist", t, func() {
		kfile := keyfile.NewKeyFile()
		err := kfile.LoadFromData([]byte(desktopFileContent10))
		So(err, ShouldBeNil)
		ai, err := NewDesktopAppInfoFromKeyFile(kfile)
		So(err, ShouldBeNil)
		So(ai, ShouldNotBeNil)
		So(ai.GetExecutable(), ShouldEqual, "/bin/sh/notexist")
		So(ai.IsExecutableOk(), ShouldBeFalse)
	})

	Convey("GetExecutable Exec=notexist", t, func() {
		kfile := keyfile.NewKeyFile()
		err := kfile.LoadFromData([]byte(desktopFileContent11))
		So(err, ShouldBeNil)
		ai, err := NewDesktopAppInfoFromKeyFile(kfile)
		So(err, ShouldBeNil)
		So(ai, ShouldNotBeNil)
		So(ai.GetExecutable(), ShouldEqual, "notexist")
		So(ai.IsExecutableOk(), ShouldBeFalse)
	})
}

func Test_splitExec(t *testing.T) {
	Convey("splitExec", t, func() {
		parts, err := splitExec(`abc def ghi`)
		So(err, ShouldBeNil)
		So(parts, ShouldResemble, []string{"abc", "def", "ghi"})

		parts, err = splitExec(`"abc" def "ghi"`)
		So(err, ShouldBeNil)
		So(parts, ShouldResemble, []string{"abc", "def", "ghi"})

		parts, err = splitExec(`"abc 123" def "ghi"`)
		So(err, ShouldBeNil)
		So(parts, ShouldResemble, []string{"abc 123", "def", "ghi"})

		parts, err = splitExec(`abc def "" ghi`)
		So(err, ShouldBeNil)
		So(parts, ShouldResemble, []string{"abc", "def", "", "ghi"})

		parts, err = splitExec(`"abc's" def`)
		So(err, ShouldBeNil)
		So(parts, ShouldResemble, []string{"abc's", "def"})

		parts, err = splitExec(`"abc\\s" def`)
		So(err, ShouldBeNil)
		So(parts, ShouldResemble, []string{"abc\\s", "def"})

		parts, err = splitExec(`abc   def ghi`)
		So(err, ShouldBeNil)
		So(parts, ShouldResemble, []string{"abc", "def", "ghi"})

		parts, err = splitExec(`"abc"`)
		So(err, ShouldBeNil)
		So(parts, ShouldResemble, []string{"abc"})

		_, err = splitExec(`"abcdef`)
		So(err, ShouldEqual, ErrQuotingNotClosed)

		_, err = splitExec(`"abcdef\"`)
		So(err, ShouldEqual, ErrQuotingNotClosed)

		_, err = splitExec(`"$abcdef"`)
		So(err, ShouldResemble, ErrCharNotEscaped{'$'})

		_, err = splitExec("\"`abcdef\" def")
		So(err, ShouldResemble, ErrCharNotEscaped{'`'})

		_, err = splitExec(`"abc\def"`)
		So(err, ShouldResemble, ErrInvalidEscapeSequence{'d'})

		_, err = splitExec(`#echo hello world`)
		So(err, ShouldResemble, ErrReservedCharNotQuoted{'#'})

		_, err = splitExec(`(1)`)
		So(err, ShouldResemble, ErrReservedCharNotQuoted{'('})

		_, err = splitExec(`"abc"def`)
		So(err, ShouldEqual, ErrNoSpaceAfterQuoting)

	})
}

func Test_expandFieldCode(t *testing.T) {
	Convey("expandFieldCode", t, func() {
		icon := "test_icon"
		desktopFile := "test.desktop"
		files := []string{"/dir1/dir2/a", "/dir1/dir2/b"}
		translatedName := "translatedName"
		cmdline := []string{"start", "%f", "%i", "%c", "%k", "end"}
		parts, err := expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
		So(err, ShouldBeNil)
		So(parts, ShouldResemble, []string{"start", files[0], "--icon", icon, translatedName, desktopFile, "end"})

		cmdline = []string{"start", "%G", "end"}
		parts, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
		So(err, ShouldEqual, ErrBadFieldCode)

		cmdline = []string{"start", "%d", "%d", "%v", "end"}
		parts, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
		So(err, ShouldBeNil)
		So(parts, ShouldResemble, []string{"start", "end"})

		cmdline = []string{"start", "%F", "end"}
		parts, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
		So(err, ShouldBeNil)
		So(parts, ShouldResemble, []string{"start", "/dir1/dir2/a", "/dir1/dir2/b", "end"})

		cmdline = []string{"start", "%u", "end"}
		parts, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
		So(err, ShouldBeNil)
		So(parts, ShouldResemble, []string{"start", "file:///dir1/dir2/a", "end"})

		cmdline = []string{"start", "%U", "end"}
		parts, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
		So(err, ShouldBeNil)
		So(parts, ShouldResemble, []string{"start", "file:///dir1/dir2/a", "file:///dir1/dir2/b", "end"})

		cmdline = []string{"start", "%%", "%abc", "end"}
		parts, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
		So(err, ShouldBeNil)
		So(parts, ShouldResemble, []string{"start", "%", "%abc", "end"})

		files = []string{"file:///home/tp/2017%E5%B9%B403%E6%9C%88-%E6%B7%B1%E5%BA%A6%E9%9B%86%E7%BB%93-%E7%94%B5%E5%AD%90%E7%89%88.pdf"}
		cmdline = []string{"/opt/Foxitreader/FoxitReader.sh", "%F"}
		parts, err = expandFieldCode(cmdline, files, translatedName, icon, desktopFile)
		So(err, ShouldBeNil)
		So(parts, ShouldResemble, []string{"/opt/Foxitreader/FoxitReader.sh", "/home/tp/2017年03月-深度集结-电子版.pdf"})
	})
}
