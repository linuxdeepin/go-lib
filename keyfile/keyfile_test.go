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

package keyfile

import (
	"bytes"
	"path/filepath"
	"regexp"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
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

const desktopFileContent4 = `
[Desktop Entry]
abc/def=nokey
`

func TestLoadFromData(t *testing.T) {

	Convey("Test LoadFromData EntryNotInSectionError", t, func(c C) {
		f := NewKeyFile()
		err := f.LoadFromData([]byte(desktopFileContent1))
		c.So(err, ShouldNotBeNil)
		c.So(err, ShouldHaveSameTypeAs, EntryNotInSectionError{})
		t.Log(err)
	})

	Convey("Test LoadFromData ParseError", t, func(c C) {
		f := NewKeyFile()
		err := f.LoadFromData([]byte(desktopFileContent2))
		c.So(err, ShouldNotBeNil)
		c.So(err, ShouldHaveSameTypeAs, ParseError{})
		t.Log(err)
	})

	Convey("Test LoadFromData key is empty", t, func(c C) {
		f := NewKeyFile()
		err := f.LoadFromData([]byte(desktopFileContent3))
		c.So(err, ShouldNotBeNil)
		t.Log(err)
	})

	Convey("Test LoadFromData InvalidKeyError", t, func(c C) {
		f := NewKeyFile()
		keyReg := regexp.MustCompile(`^[A-Za-z0-9\-]+$`)
		f.SetKeyRegexp(keyReg)
		err := f.LoadFromData([]byte(desktopFileContent4))
		c.So(err, ShouldNotBeNil)
		t.Log(err)

		ret := f.SetValue("Desktop Entry", "Abc+", "123")
		c.So(ret, ShouldBeFalse)

		ret = f.SetValue("Desktop Entry", "Abc", "123")
		c.So(ret, ShouldBeTrue)
	})

	Convey("Test LoadFromData", t, func(c C) {
		f := NewKeyFile()
		err := f.LoadFromData([]byte(desktopFileContent0))
		c.So(err, ShouldBeNil)

		c.Convey("Test GetSections", func(c C) {
			c.So(f.GetSections(), ShouldResemble, []string{"Desktop Entry"})
		})

		c.Convey("Get value that does exist", func(c C) {
			v, err := f.GetValue("Desktop Entry", "Type")
			c.So(err, ShouldBeNil)
			c.So(v, ShouldEqual, "Application")
		})

		c.Convey("Get value the does not exist", func(c C) {
			_, err := f.GetValue("Desktop Entry", "x")
			c.So(err, ShouldNotBeNil)
		})

		c.Convey("Get value that section not exist", func(c C) {
			_, err := f.GetValue("X", "X")
			c.So(err, ShouldNotBeNil)
		})

		c.Convey("Test GetKeys", func(c C) {
			c.So(f.GetKeys("Desktop Entry"), ShouldResemble, []string{
				"Encoding", "Type", "X-Created-By", "Categories",
				"Icon", "Exec", "Name",
			})
		})

		c.Convey("Get section comments", func(c C) {
			c.So(f.GetSectionComments("Desktop Entry"), ShouldEqual, "#!/usr/bin/env xdg-open\n")
		})

		c.Convey("Get key comments", func(c C) {
			c.So(f.GetKeyComments("Desktop Entry", "Icon"), ShouldEqual, "# icon comments")
			c.So(f.GetKeyComments("Desktop Entry", "Categories"), ShouldEqual, "")
		})
	})
}

func TestLoadFromFile(t *testing.T) {
	Convey("LoadFromFile ok", t, func(c C) {
		f := NewKeyFile()
		err := f.LoadFromFile("testdata/deepin-screenshot.desktop")
		c.So(err, ShouldBeNil)

		localeName, err := f.GetLocaleString("Desktop Entry", "Name", "zh_CN")
		c.So(localeName, ShouldEqual, "深度截图")
		c.So(err, ShouldBeNil)
	})

	Convey("Load file in dir", t, func(c C) {
		files, err := filepath.Glob("/usr/share/applications/*.desktop")
		c.So(err, ShouldBeNil)
		for _, file := range files {
			f := NewKeyFile()
			err := f.LoadFromFile(file)
			c.So(err, ShouldBeNil)
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
	Convey("GetStringList", t, func(c C) {
		f := NewKeyFile()
		keyFileContent := keyFileContent0 + string([]byte{0xff, 0xfe, 0xfd})
		err := f.LoadFromData([]byte(keyFileContent))
		c.So(err, ShouldBeNil)

		list, err := f.GetStringList("Test", "strlist0")
		c.So(err, ShouldBeNil)
		c.So(list, ShouldResemble, []string{"a", "b", "c", "d"})

		list, err = f.GetStringList("Test", "strlist1")
		c.So(err, ShouldBeNil)
		c.So(list, ShouldResemble, []string{"a", "b", "c"})

		list, err = f.GetStringList("Test", "strlist2")
		c.So(err, ShouldBeNil)
		c.So(list, ShouldResemble, []string{""})

		list, err = f.GetStringList("Test", "strlist3")
		c.So(err, ShouldBeNil)
		c.So(list, ShouldBeNil)

		list, err = f.GetStringList("Test", "strlist4")
		c.So(err, ShouldBeNil)
		c.So(list, ShouldResemble, []string{"abc"})

		t.Log(err)

		list, err = f.GetStringList("Test", "strlist5")
		c.So(err, ShouldBeNil)
		c.So(list, ShouldResemble, []string{"abc\\befg"})
		t.Log(err)

		list, err = f.GetStringList("Test", "strlist6")
		c.So(err, ShouldBeNil)
		c.So(list, ShouldResemble, []string{"a;bc", "def"})

		_, err = f.GetStringList("Test", "strlist7")
		c.So(err, ShouldNotBeNil)
		c.So(err, ShouldHaveSameTypeAs, ValueInvalidUTF8Error{})
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
	Convey("GetString", t, func(c C) {
		f := NewKeyFile()
		keyFileContent := keyFileContent1 + string([]byte{0xff, 0xfe, 0xfd})
		err := f.LoadFromData([]byte(keyFileContent))
		c.So(err, ShouldBeNil)

		str, err := f.GetString("Test", "str0")
		c.So(err, ShouldBeNil)
		c.So(str, ShouldEqual, "abcdef")

		str, err = f.GetString("Test", "str1")
		c.So(err, ShouldBeNil)
		c.So(str, ShouldEqual, "line\n<-newline\r<-break\t<-table <-space")

		str, err = f.GetString("Test", "str2")
		c.So(err, ShouldBeNil)
		c.So(str, ShouldEqual, "abcdef")

		str, err = f.GetString("Test", "str3")
		c.So(err, ShouldBeNil)
		c.So(str, ShouldEqual, "abc\\bdef")

		str, err = f.GetString("Test", "str4")
		c.So(err, ShouldBeNil)
		c.So(str, ShouldEqual, "abc\\;def")

		_, err = f.GetString("Test", "str5")
		c.So(err, ShouldHaveSameTypeAs, ValueInvalidUTF8Error{})
	})
}

func TestSetString(t *testing.T) {
	Convey("SetString", t, func(c C) {
		f := NewKeyFile()
		const s0 = "space newline\ncarriage-return\rtab\tbackslash\\"
		f.SetString("Test", "str0", s0)
		str0, err := f.GetString("Test", "str0")
		c.So(err, ShouldBeNil)
		c.So(str0, ShouldEqual, s0)
		val0, err := f.GetValue("Test", "str0")
		c.So(err, ShouldBeNil)
		c.So(val0, ShouldEqual, `space newline\ncarriage-return\rtab\tbackslash\\`)
	})
}

func TestSetStringList(t *testing.T) {
	Convey("SetStringList", t, func(c C) {
		f := NewKeyFile()
		strlist := []string{"space ", "newline\n", "carriage\rreturn", "tab\t", "backslash\\", "List;Separator;"}
		f.SetStringList("Test", "strlist", strlist)
		strlist1, err := f.GetStringList("Test", "strlist")
		c.So(err, ShouldBeNil)
		c.So(strlist1, ShouldResemble, strlist)

		strlistValue, err := f.GetValue("Test", "strlist")
		c.So(err, ShouldBeNil)
		c.So(strlistValue, ShouldEqual, `space\s;newline\n;carriage\rreturn;tab\t;backslash\\;List\;Separator\;;`)
	})
}

func TestSetBoolList(t *testing.T) {
	Convey("SetBoolList", t, func(c C) {
		f := NewKeyFile()
		blist := []bool{true, true, false, false, true, false}
		f.SetBoolList("Test", "blist", blist)
		blist1, err := f.GetBoolList("Test", "blist")
		c.So(err, ShouldBeNil)
		c.So(blist1, ShouldResemble, blist)

		blistStr, err := f.GetValue("Test", "blist")
		c.So(err, ShouldBeNil)
		c.So(blistStr, ShouldEqual, "true;true;false;false;true;false;")
	})
}

func TestSetIntList(t *testing.T) {
	Convey("SetIntList", t, func(c C) {
		f := NewKeyFile()
		ints := []int{-345, -1, 0, 1, 3, 5, 7, 9, 11989}
		f.SetIntList("Test", "ints", ints)
		ints1, err := f.GetIntList("Test", "ints")
		c.So(err, ShouldBeNil)
		c.So(ints1, ShouldResemble, ints)

		intsStr, err := f.GetValue("Test", "ints")
		c.So(err, ShouldBeNil)
		c.So(intsStr, ShouldResemble, "-345;-1;0;1;3;5;7;9;11989;")

	})
}

const keyFileContent2 = `[Test]
KeyA=aaaa
KeyB=1234567890
KeyC=true

[Main]
KeyD=keyfile

`

func TestSaveToWriter(t *testing.T) {
	Convey("SaveToWriter", t, func(c C) {
		f := NewKeyFile()
		f.SetValue("Test", "KeyA", "aaaa")
		f.SetValue("Test", "KeyB", "1234567890")
		f.SetValue("Test", "KeyC", "true")

		f.SetValue("Main", "KeyD", "keyfile")

		var buf bytes.Buffer
		err := f.SaveToWriter(&buf)
		c.So(err, ShouldBeNil)
		c.So(buf.String(), ShouldEqual, keyFileContent2)
	})
}
