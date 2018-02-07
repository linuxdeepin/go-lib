/*
 * Copyright (C) 2016 ~ 2018 Deepin Technology Co., Ltd.
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

package procfs

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"strings"
	"testing"
)

const packagePath = "/pkg.deepin.io/lib/procfs/"

func TestGetFile(t *testing.T) {
	Convey("getFile", t, func() {
		p := Process(1)
		So(p.getFile("cwd"), ShouldEqual, "/proc/1/cwd")
	})
}

func TestExist(t *testing.T) {
	Convey("Exist", t, func() {
		p := Process(os.Getpid())
		So(p.Exist(), ShouldBeTrue)
	})
}

func TestCmdline(t *testing.T) {
	Convey("Cmdline", t, func() {
		p := Process(os.Getpid())
		cmdline, err := p.Cmdline()
		So(err, ShouldBeNil)
		t.Log("cmdline:", cmdline)
		So(len(cmdline) > 0, ShouldBeTrue)
		So(strings.Contains(cmdline[0], packagePath), ShouldBeTrue)
	})
}

func TestCwd(t *testing.T) {
	Convey("Cwd", t, func() {
		p := Process(os.Getpid())
		cwd, err := p.Cwd()
		So(err, ShouldBeNil)
		t.Log("cwd:", cwd)

		osWd, err1 := os.Getwd()
		So(err1, ShouldBeNil)
		So(cwd, ShouldEqual, osWd)
	})
}

func TestExe(t *testing.T) {
	Convey("Exe", t, func() {
		p := Process(os.Getpid())
		exe, err := p.Exe()
		So(err, ShouldBeNil)
		t.Log("exe:", exe)
		So(strings.Contains(exe, packagePath), ShouldBeTrue)
	})
}

func TestEnvVars(t *testing.T) {
	vars := EnvVars{
		"PWD=/a/b/c",
	}
	Convey("EnvVars.Lookup", t, func() {
		pwd, ok := vars.Lookup("PWD")
		So(pwd, ShouldEqual, "/a/b/c")
		So(ok, ShouldBeTrue)

		abc, ok := vars.Lookup("abc")
		So(abc, ShouldEqual, "")
		So(ok, ShouldBeFalse)
	})

	Convey("EnvVars.Get", t, func() {
		pwd := vars.Get("PWD")
		So(pwd, ShouldEqual, "/a/b/c")

		abc := vars.Get("abc")
		So(abc, ShouldEqual, "")
	})
}

func TestEnvion(t *testing.T) {
	Convey("Envion", t, func() {
		p := Process(os.Getpid())
		environ, err := p.Environ()
		So(err, ShouldBeNil)
		So(len(environ) > 0, ShouldBeTrue)
		for _, aVar := range environ {
			t.Log(string(aVar))
		}

		path, ok := environ.Lookup("PATH")
		So(ok, ShouldBeTrue)
		So(path != "", ShouldBeTrue)

		home, ok := environ.Lookup("HOME")
		So(ok, ShouldBeTrue)
		So(home != "", ShouldBeTrue)

		xxx, ok := environ.Lookup("XXXXXXXXXXXXXXX")
		So(ok, ShouldBeFalse)
		So(xxx, ShouldEqual, "")
	})
}

func TestStatus(t *testing.T) {
	Convey("Status", t, func() {
		p := Process(os.Getpid())
		status, err := p.Status()
		So(err, ShouldBeNil)
		So(status, ShouldNotBeEmpty)

		// test lookup
		val, err := status.lookup("XXX")
		So(val, ShouldBeBlank)
		So(err, ShouldResemble, StatusFieldNotFoundErr{"XXX"})
		So(err.Error(), ShouldEqual, "field XXX is not found in proc status file")

		// test Uids
		uids, err := status.Uids()
		So(err, ShouldBeNil)
		t.Log("uids:", uids)
		So(uids[0], ShouldEqual, uint(os.Getuid()))
		So(uids[1], ShouldEqual, uint(os.Geteuid()))

		// test PPid
		ppid, err := status.PPid()
		So(err, ShouldBeNil)
		t.Log("ppid:", ppid)
		So(ppid, ShouldBeGreaterThan, 0)
	})
}
