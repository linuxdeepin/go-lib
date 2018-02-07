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

package basedir

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestGetUserHomeDir(t *testing.T) {
	Convey("GetUserHomeDir", t, func() {
		os.Setenv("HOME", "/home/test")
		dir := GetUserHomeDir()
		So(dir, ShouldEqual, "/home/test")
	})
}

func TestGetUserDataDir(t *testing.T) {
	Convey("GetUserDataDir", t, func() {
		os.Setenv("HOME", "/home/test")
		os.Setenv("XDG_DATA_HOME", "")
		dir := GetUserDataDir()
		So(dir, ShouldEqual, "/home/test/.local/share")

		os.Setenv("XDG_DATA_HOME", "a invalid path")
		dir = GetUserDataDir()
		So(dir, ShouldEqual, "/home/test/.local/share")

		os.Setenv("XDG_DATA_HOME", "/home/test/xdg")
		dir = GetUserDataDir()
		So(dir, ShouldEqual, "/home/test/xdg")
	})
}

func TestFilterNotAbs(t *testing.T) {
	Convey("filterNotAbs", t, func() {
		result := filterNotAbs([]string{
			"a/is/invald", "b/is invalid", "c is invald", "/d/is/ok", "/e/is/ok/"})
		So(result, ShouldResemble, []string{"/d/is/ok", "/e/is/ok"})
	})
}

func TestGetSystemDataDirs(t *testing.T) {
	Convey("GetSystemDataDirs", t, func() {
		os.Setenv("XDG_DATA_DIRS", "/a:/b:/c")
		dirs := GetSystemDataDirs()
		So(dirs, ShouldResemble, []string{"/a", "/b", "/c"})

		os.Setenv("XDG_DATA_DIRS", "/a:/b/:/c/")
		dirs = GetSystemDataDirs()
		So(dirs, ShouldResemble, []string{"/a", "/b", "/c"})

		os.Setenv("XDG_DATA_DIRS", "/a:/b/:c is invald")
		dirs = GetSystemDataDirs()
		So(dirs, ShouldResemble, []string{"/a", "/b"})

		os.Setenv("XDG_DATA_DIRS", "a/is/invald:b/is invalid :c is invald")
		dirs = GetSystemDataDirs()
		So(dirs, ShouldResemble, []string{"/usr/local/share", "/usr/share"})

		os.Setenv("XDG_DATA_DIRS", "")
		dirs = GetSystemDataDirs()
		So(dirs, ShouldResemble, []string{"/usr/local/share", "/usr/share"})
	})
}

func TestGetUserConfigDir(t *testing.T) {
	Convey("GetUserConfigDir", t, func() {
		os.Setenv("XDG_CONFIG_HOME", "")
		os.Setenv("HOME", "/home/test")
		dir := GetUserConfigDir()
		So(dir, ShouldEqual, "/home/test/.config")

		os.Setenv("XDG_CONFIG_HOME", "/home/test/myconfig")
		dir = GetUserConfigDir()
		So(dir, ShouldEqual, "/home/test/myconfig")
	})
}

func TestGetSystemConfigDirs(t *testing.T) {
	Convey("GetSystemDirs", t, func() {
		os.Setenv("XDG_CONFIG_DIRS", "/a:/b:/c")
		dirs := GetSystemConfigDirs()
		So(dirs, ShouldResemble, []string{"/a", "/b", "/c"})

		os.Setenv("XDG_CONFIG_DIRS", "")
		dirs = GetSystemConfigDirs()
		So(dirs, ShouldResemble, []string{"/etc/xdg"})
	})
}

func TestGetUserCacheDir(t *testing.T) {
	Convey("GetUserCacheDir", t, func() {
		os.Setenv("XDG_CACHE_HOME", "/cache/user/a")
		dir := GetUserCacheDir()
		So(dir, ShouldEqual, "/cache/user/a")

		os.Setenv("XDG_CACHE_HOME", "")
		os.Setenv("HOME", "/home/test")
		dir = GetUserCacheDir()
		So(dir, ShouldEqual, "/home/test/.cache")
	})
}

func TestGetUserRuntimeDir(t *testing.T) {
	Convey("GetUserRuntimeDir", t, func() {
		os.Setenv("XDG_RUNTIME_DIR", "/runtime/user/test")
		dir, err := GetUserRuntimeDir(true)
		So(err, ShouldBeNil)
		So(dir, ShouldEqual, "/runtime/user/test")

		os.Setenv("XDG_RUNTIME_DIR", "")
		dir, err = GetUserRuntimeDir(true)
		So(err, ShouldNotBeNil)
		So(dir, ShouldEqual, "")

		os.Setenv("XDG_RUNTIME_DIR", "a invalid path")
		dir, err = GetUserRuntimeDir(true)
		So(err, ShouldNotBeNil)
		So(dir, ShouldEqual, "")

		os.Setenv("XDG_RUNTIME_DIR", "")
		dir, err = GetUserRuntimeDir(false)
		So(err, ShouldBeNil)
		So(dir, ShouldEqual, fmt.Sprintf("/tmp/goxdg-runtime-dir-fallback-%d", os.Getuid()))
	})
}
