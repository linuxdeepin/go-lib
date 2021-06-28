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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	f := NewKeyFile()
	err := f.LoadFromData([]byte(desktopFileContent1))
	assert.NotNil(t, err)
	assert.IsType(t, err, EntryNotInSectionError{})
	t.Log(err)

	f = NewKeyFile()
	err = f.LoadFromData([]byte(desktopFileContent2))
	assert.NotNil(t, err)
	assert.IsType(t, err, ParseError{})
	t.Log(err)

	f = NewKeyFile()
	err = f.LoadFromData([]byte(desktopFileContent3))
	assert.NotNil(t, err)
	t.Log(err)

	f = NewKeyFile()
	keyReg := regexp.MustCompile(`^[A-Za-z0-9\-]+$`)
	f.SetKeyRegexp(keyReg)
	err = f.LoadFromData([]byte(desktopFileContent4))
	assert.NotNil(t, err)
	t.Log(err)

	ret := f.SetValue("Desktop Entry", "Abc+", "123")
	assert.False(t, ret)

	ret = f.SetValue("Desktop Entry", "Abc", "123")
	assert.True(t, ret)

	f = NewKeyFile()
	err = f.LoadFromData([]byte(desktopFileContent0))
	require.Nil(t, err)

	assert.Equal(t, f.GetSections(), []string{"Desktop Entry"})

	v, err := f.GetValue("Desktop Entry", "Type")
	require.Nil(t, err)
	assert.Equal(t, v, "Application")

	_, err = f.GetValue("Desktop Entry", "x")
	assert.NotNil(t, err)

	_, err = f.GetValue("X", "X")
	assert.NotNil(t, err)

	assert.Equal(t, f.GetKeys("Desktop Entry"), []string{
		"Encoding", "Type", "X-Created-By", "Categories",
		"Icon", "Exec", "Name",
	})

	assert.Equal(t, f.GetSectionComments("Desktop Entry"), "#!/usr/bin/env xdg-open\n")

	assert.Equal(t, f.GetKeyComments("Desktop Entry", "Icon"), "# icon comments")
	assert.Equal(t, f.GetKeyComments("Desktop Entry", "Categories"), "")
}

func TestLoadFromFile(t *testing.T) {
	f := NewKeyFile()
	err := f.LoadFromFile("testdata/deepin-screenshot.desktop")
	require.Nil(t, err)

	localeName, err := f.GetLocaleString("Desktop Entry", "Name", "zh_CN")
	assert.Equal(t, localeName, "深度截图")
	require.Nil(t, err)

	files, err := filepath.Glob("/usr/share/applications/*.desktop")
	require.Nil(t, err)
	for _, file := range files {
		f := NewKeyFile()
		err := f.LoadFromFile(file)
		require.Nil(t, err)
	}
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
	f := NewKeyFile()
	keyFileContent := keyFileContent0 + string([]byte{0xff, 0xfe, 0xfd})
	err := f.LoadFromData([]byte(keyFileContent))
	require.Nil(t, err)

	list, err := f.GetStringList("Test", "strlist0")
	require.Nil(t, err)
	assert.Equal(t, list, []string{"a", "b", "c", "d"})

	list, err = f.GetStringList("Test", "strlist1")
	require.Nil(t, err)
	assert.Equal(t, list, []string{"a", "b", "c"})

	list, err = f.GetStringList("Test", "strlist2")
	require.Nil(t, err)
	assert.Equal(t, list, []string{""})

	list, err = f.GetStringList("Test", "strlist3")
	require.Nil(t, err)
	require.Nil(t, list)

	list, err = f.GetStringList("Test", "strlist4")
	require.Nil(t, err)
	assert.Equal(t, list, []string{"abc"})

	t.Log(err)

	list, err = f.GetStringList("Test", "strlist5")
	require.Nil(t, err)
	assert.Equal(t, list, []string{"abc\\befg"})
	t.Log(err)

	list, err = f.GetStringList("Test", "strlist6")
	require.Nil(t, err)
	assert.Equal(t, list, []string{"a;bc", "def"})

	_, err = f.GetStringList("Test", "strlist7")
	assert.NotNil(t, err)
	assert.IsType(t, err, ValueInvalidUTF8Error{})
	t.Log(err)
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
	f := NewKeyFile()
	keyFileContent := keyFileContent1 + string([]byte{0xff, 0xfe, 0xfd})
	err := f.LoadFromData([]byte(keyFileContent))
	require.Nil(t, err)

	str, err := f.GetString("Test", "str0")
	require.Nil(t, err)
	assert.Equal(t, str, "abcdef")

	str, err = f.GetString("Test", "str1")
	require.Nil(t, err)
	assert.Equal(t, str, "line\n<-newline\r<-break\t<-table <-space")

	str, err = f.GetString("Test", "str2")
	require.Nil(t, err)
	assert.Equal(t, str, "abcdef")

	str, err = f.GetString("Test", "str3")
	require.Nil(t, err)
	assert.Equal(t, str, "abc\\bdef")

	str, err = f.GetString("Test", "str4")
	require.Nil(t, err)
	assert.Equal(t, str, "abc\\;def")

	_, err = f.GetString("Test", "str5")
	assert.IsType(t, err, ValueInvalidUTF8Error{})
}

func TestSetString(t *testing.T) {
	f := NewKeyFile()
	const s0 = "space newline\ncarriage-return\rtab\tbackslash\\"
	f.SetString("Test", "str0", s0)
	str0, err := f.GetString("Test", "str0")
	require.Nil(t, err)
	assert.Equal(t, str0, s0)
	val0, err := f.GetValue("Test", "str0")
	require.Nil(t, err)
	assert.Equal(t, val0, `space newline\ncarriage-return\rtab\tbackslash\\`)
}

func TestSetStringList(t *testing.T) {
	f := NewKeyFile()
	strlist := []string{"space ", "newline\n", "carriage\rreturn", "tab\t", "backslash\\", "List;Separator;"}
	f.SetStringList("Test", "strlist", strlist)
	strlist1, err := f.GetStringList("Test", "strlist")
	require.Nil(t, err)
	assert.Equal(t, strlist1, strlist)

	strlistValue, err := f.GetValue("Test", "strlist")
	require.Nil(t, err)
	assert.Equal(t, strlistValue, `space\s;newline\n;carriage\rreturn;tab\t;backslash\\;List\;Separator\;;`)
}

func TestSetBoolList(t *testing.T) {
	f := NewKeyFile()
	blist := []bool{true, true, false, false, true, false}
	f.SetBoolList("Test", "blist", blist)
	blist1, err := f.GetBoolList("Test", "blist")
	require.Nil(t, err)
	assert.Equal(t, blist1, blist)

	blistStr, err := f.GetValue("Test", "blist")
	require.Nil(t, err)
	assert.Equal(t, blistStr, "true;true;false;false;true;false;")
}

func TestSetIntList(t *testing.T) {
	f := NewKeyFile()
	ints := []int{-345, -1, 0, 1, 3, 5, 7, 9, 11989}
	f.SetIntList("Test", "ints", ints)
	ints1, err := f.GetIntList("Test", "ints")
	require.Nil(t, err)
	assert.Equal(t, ints1, ints)

	intsStr, err := f.GetValue("Test", "ints")
	require.Nil(t, err)
	assert.Equal(t, intsStr, "-345;-1;0;1;3;5;7;9;11989;")

}

const keyFileContent2 = `[Test]
KeyA=aaaa
KeyB=1234567890
KeyC=true

[Main]
KeyD=keyfile

`

func TestSaveToWriter(t *testing.T) {
	f := NewKeyFile()
	f.SetValue("Test", "KeyA", "aaaa")
	f.SetValue("Test", "KeyB", "1234567890")
	f.SetValue("Test", "KeyC", "true")

	f.SetValue("Main", "KeyD", "keyfile")

	var buf bytes.Buffer
	err := f.SaveToWriter(&buf)
	require.Nil(t, err)
	assert.Equal(t, buf.String(), keyFileContent2)
}
