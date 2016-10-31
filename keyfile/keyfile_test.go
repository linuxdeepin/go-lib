package keyfile

import (
	. "github.com/smartystreets/goconvey/convey"
	"path/filepath"
	"testing"
)

const desktopFileContent0 = `#!/usr/bin/env xdg-open

[Desktop Entry]
Encoding=UTF-8
Type=Application
X-Created-By=cxoffice-407693d8-5e19-11e6-845f-2b300706ad83
Categories=chat;
# icon comments
Icon=apps.com.qq.im.light
Exec="/opt/cxoffice/support/apps.com.qq.im.light/desktopdata/cxmenu/StartMenu.C^5E3A_users_crossover_Start^2BMenu/Programs/QQ轻聊版.lnk" %u
Name=QQ轻聊版`

const desktopFileContent1 = `
Encoding=UTF-8
Type=Application
`

const desktopFileContent2 = `
[Desktop Entry]
Encoding=UTF-8
Type=Application
X-Created-By=cxoffice-407693d8-5e19-11e6-845f-2b300706ad83
Categories: chat;
`

const desktopFileContent3 = `
[Desktop Entry]
=nokey
`

func TestLoadFromData(t *testing.T) {

	Convey("Test LoadFromData EntryNotInSectionError", t, func() {
		f := NewKeyFile()
		err := f.LoadFromData([]byte(desktopFileContent1))
		So(err, ShouldNotBeNil)
		So(err, ShouldHaveSameTypeAs, EntryNotInSectionError{})
		t.Log(err)
	})

	Convey("Test LoadFromData ParseError", t, func() {
		f := NewKeyFile()
		err := f.LoadFromData([]byte(desktopFileContent2))
		So(err, ShouldNotBeNil)
		So(err, ShouldHaveSameTypeAs, ParseError{})
		t.Log(err)
	})

	Convey("Test LoadFromData InvalidKeyError", t, func() {
		f := NewKeyFile()
		err := f.LoadFromData([]byte(desktopFileContent3))
		So(err, ShouldNotBeNil)
		So(err, ShouldHaveSameTypeAs, InvalidKeyError{})
		t.Log(err)
	})

	Convey("Test LoadFromData", t, func() {
		f := NewKeyFile()
		err := f.LoadFromData([]byte(desktopFileContent0))
		So(err, ShouldBeNil)

		Convey("Test GetSections", func() {
			So(f.GetSections(), ShouldResemble, []string{"Desktop Entry"})
		})

		Convey("Get value that does exist", func() {
			v, err := f.GetValue("Desktop Entry", "Type")
			So(err, ShouldBeNil)
			So(v, ShouldEqual, "Application")
		})

		Convey("Get value the does not exist", func() {
			_, err := f.GetValue("Desktop Entry", "x")
			So(err, ShouldNotBeNil)
		})

		Convey("Get value that section not exist", func() {
			_, err := f.GetValue("X", "X")
			So(err, ShouldNotBeNil)
		})

		Convey("Test GetKeys", func() {
			So(f.GetKeys("Desktop Entry"), ShouldResemble, []string{
				"Encoding", "Type", "X-Created-By", "Categories",
				"Icon", "Exec", "Name",
			})
		})

		Convey("Get section comments", func() {
			So(f.GetSectionComments("Desktop Entry"), ShouldEqual, "#!/usr/bin/env xdg-open\n")
		})

		Convey("Get key comments", func() {
			So(f.GetKeyComments("Desktop Entry", "Icon"), ShouldEqual, "# icon comments")
			So(f.GetKeyComments("Desktop Entry", "Categories"), ShouldEqual, "")
		})
	})
}

func TestLoadFromFile(t *testing.T) {
	Convey("LoadFromFile ok", t, func() {
		f := NewKeyFile()
		err := f.LoadFromFile("testdata/deepin-screenshot.desktop")
		So(err, ShouldBeNil)

		localeName, err := f.GetLocaleString("Desktop Entry", "Name", "zh_CN")
		So(localeName, ShouldEqual, "深度截图")
		So(err, ShouldBeNil)
	})

	Convey("Load file in dir", t, func() {
		files, err := filepath.Glob("/usr/share/applications/*.desktop")
		So(err, ShouldBeNil)
		for _, file := range files {
			f := NewKeyFile()
			err := f.LoadFromFile(file)
			So(err, ShouldBeNil)
		}
	})
}

const keyFileContent0 = `
[Test]
strlist0=a;b;c;d;
strlist1=a;b;c
strlist2=;
strlist3=
strlist4=abc\
strlist5=abc\befg
strlist6=a\;bc;def
strlist7=?`

func TestGetStringList(t *testing.T) {
	Convey("GetStringList", t, func() {
		f := NewKeyFile()
		keyFileContent := keyFileContent0 + string([]byte{0xff, 0xfe, 0xfd})
		err := f.LoadFromData([]byte(keyFileContent))
		So(err, ShouldBeNil)

		list, err := f.GetStringList("Test", "strlist0")
		So(err, ShouldBeNil)
		So(list, ShouldResemble, []string{"a", "b", "c", "d"})

		list, err = f.GetStringList("Test", "strlist1")
		So(err, ShouldBeNil)
		So(list, ShouldResemble, []string{"a", "b", "c"})

		list, err = f.GetStringList("Test", "strlist2")
		So(err, ShouldBeNil)
		So(list, ShouldResemble, []string{""})

		list, err = f.GetStringList("Test", "strlist3")
		So(err, ShouldBeNil)
		So(list, ShouldBeNil)

		list, err = f.GetStringList("Test", "strlist4")
		So(err, ShouldNotBeNil)
		So(err, ShouldHaveSameTypeAs, ParseValueAsStringError{})
		t.Log(err)

		list, err = f.GetStringList("Test", "strlist5")
		So(err, ShouldNotBeNil)
		So(err, ShouldHaveSameTypeAs, ParseValueAsStringError{})
		t.Log(err)

		list, err = f.GetStringList("Test", "strlist6")
		So(err, ShouldBeNil)
		So(list, ShouldResemble, []string{"a;bc", "def"})

		list, err = f.GetStringList("Test", "strlist7")
		So(err, ShouldNotBeNil)
		So(err, ShouldHaveSameTypeAs, ValueInvalidUTF8Error{})
		t.Log(err)
	})
}

const keyFileContent1 = `
[Test]
str0 = abcdef 
str1=line\n<-newline\r<-break\t<-table\s<-space
str2=abcdef\
str3=abc\bdef
str4=abc\;def
str5=?`

func TestGetString(t *testing.T) {
	Convey("GetString", t, func() {
		f := NewKeyFile()
		keyFileContent := keyFileContent1 + string([]byte{0xff, 0xfe, 0xfd})
		err := f.LoadFromData([]byte(keyFileContent))
		So(err, ShouldBeNil)

		str, err := f.GetString("Test", "str0")
		So(err, ShouldBeNil)
		So(str, ShouldEqual, "abcdef")

		str, err = f.GetString("Test", "str1")
		So(err, ShouldBeNil)
		So(str, ShouldEqual, "line\n<-newline\r<-break\t<-table <-space")

		str, err = f.GetString("Test", "str2")
		So(err, ShouldHaveSameTypeAs, ParseValueAsStringError{})

		str, err = f.GetString("Test", "str3")
		So(err, ShouldHaveSameTypeAs, ParseValueAsStringError{})

		str, err = f.GetString("Test", "str4")
		So(err, ShouldHaveSameTypeAs, ParseValueAsStringError{})

		str, err = f.GetString("Test", "str5")
		So(err, ShouldHaveSameTypeAs, ValueInvalidUTF8Error{})
	})
}
