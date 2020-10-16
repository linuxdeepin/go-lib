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
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetFile(t *testing.T) {
	Convey("getFile", t, func(c C) {
		p := Process(1)
		c.So(p.getFile("cwd"), ShouldEqual, "/proc/1/cwd")
	})
}

func TestExist(t *testing.T) {
	Convey("Exist", t, func(c C) {
		p := Process(os.Getpid())
		c.So(p.Exist(), ShouldBeTrue)
	})
}

func TestCmdline(t *testing.T) {
	Convey("Cmdline", t, func(c C) {
		p := Process(os.Getpid())
		cmdline, err := p.Cmdline()
		c.So(err, ShouldBeNil)
		t.Log("cmdline:", cmdline)
		c.So(len(cmdline) > 0, ShouldBeTrue)
	})
}

func TestCwd(t *testing.T) {
	Convey("Cwd", t, func(c C) {
		p := Process(os.Getpid())
		cwd, err := p.Cwd()
		c.So(err, ShouldBeNil)
		t.Log("cwd:", cwd)

		osWd, err1 := os.Getwd()
		c.So(err1, ShouldBeNil)
		c.So(cwd, ShouldEqual, osWd)
	})
}

func TestExe(t *testing.T) {
	Convey("Exe", t, func(c C) {
		p := Process(os.Getpid())
		exe, err := p.Exe()
		c.So(err, ShouldBeNil)
		t.Log("exe:", exe)
		c.So(len(exe) > 0, ShouldBeTrue)
	})
}

func TestEnvVars(t *testing.T) {
	vars := EnvVars{
		"PWD=/a/b/c",
	}
	Convey("EnvVars.Lookup", t, func(c C) {
		pwd, ok := vars.Lookup("PWD")
		c.So(pwd, ShouldEqual, "/a/b/c")
		c.So(ok, ShouldBeTrue)

		abc, ok := vars.Lookup("abc")
		c.So(abc, ShouldEqual, "")
		c.So(ok, ShouldBeFalse)
	})

	Convey("EnvVars.Get", t, func(c C) {
		pwd := vars.Get("PWD")
		c.So(pwd, ShouldEqual, "/a/b/c")

		abc := vars.Get("abc")
		c.So(abc, ShouldEqual, "")
	})
}

func TestEnvion(t *testing.T) {
	Convey("Envion", t, func(c C) {
		p := Process(os.Getpid())
		environ, err := p.Environ()
		c.So(err, ShouldBeNil)
		c.So(len(environ) > 0, ShouldBeTrue)
		for _, aVar := range environ {
			t.Log(string(aVar))
		}

		path, ok := environ.Lookup("PATH")
		c.So(ok, ShouldBeTrue)
		c.So(path != "", ShouldBeTrue)

		home, ok := environ.Lookup("HOME")
		c.So(ok, ShouldBeTrue)
		c.So(home != "", ShouldBeTrue)

		xxx, ok := environ.Lookup("XXXXXXXXXXXXXXX")
		c.So(ok, ShouldBeFalse)
		c.So(xxx, ShouldEqual, "")
	})
}

func TestStatus(t *testing.T) {
	Convey("Status", t, func(c C) {
		p := Process(os.Getpid())
		status, err := p.Status()
		c.So(err, ShouldBeNil)
		c.So(status, ShouldNotBeEmpty)

		// test lookup
		val, err := status.lookup("XXX")
		c.So(val, ShouldBeBlank)
		c.So(err, ShouldResemble, StatusFieldNotFoundErr{"XXX"})
		c.So(err.Error(), ShouldEqual, "field XXX is not found in proc status file")

		// test Uids
		uids, err := status.Uids()
		c.So(err, ShouldBeNil)
		t.Log("uids:", uids)
		c.So(uids[0], ShouldEqual, uint(os.Getuid()))
		c.So(uids[1], ShouldEqual, uint(os.Geteuid()))

		// test PPid
		ppid, err := status.PPid()
		c.So(err, ShouldBeNil)
		t.Log("ppid:", ppid)
		c.So(ppid, ShouldBeGreaterThan, 0)
	})
}
