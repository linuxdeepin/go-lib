/*
 * Copyright (C) 2014 ~ 2017 Deepin Technology Co., Ltd.
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

package dbus

import (
	"os"
	"testing"
)

const testString = `This is a test!
This text should be read from the file that is created by this test.`

type unixFDTest struct{}

func (t unixFDTest) Test(fd UnixFD) (string, error) {
	var b [4096]byte
	file := os.NewFile(uintptr(fd), "testfile")
	defer file.Close()
	n, err := file.Read(b[:])
	if err != nil {
		return "", &dbusError{"com.github.guelfey.test.Error", nil}
	}
	return string(b[:n]), nil
}

func TestUnixFDs(t *testing.T) {
	conn, err := SessionBus()
	if err != nil {
		t.Fatal(err)
	}
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer w.Close()
	if _, err := w.Write([]byte(testString)); err != nil {
		t.Fatal(err)
	}
	name := conn.Names()[0]
	test := unixFDTest{}
	conn.Export(test, "/com/github/guelfey/test", "com.github.guelfey.test")
	var s string
	obj := conn.Object(name, "/com/github/guelfey/test")
	err = obj.Call("com.github.guelfey.test.Test", 0, UnixFD(r.Fd())).Store(&s)
	if err != nil {
		t.Fatal(err)
	}
	if s != testString {
		t.Fatal("got", s, "wanted", testString)
	}
}
