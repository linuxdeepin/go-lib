/*
 * Copyright (C) 2016 ~ 2017 Deepin Technology Co., Ltd.
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

package userdir

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"path/filepath"
	"testing"
)

func TestParseValue(t *testing.T) {
	homeDir := "/home/test"
	Convey("parseValue", t, func() {
		value, err := parseValue([]byte(`"$HOME/Desktop/"`), homeDir)
		So(err, ShouldBeNil)
		So(value, ShouldEqual, "/home/test/Desktop")

		value, err = parseValue([]byte(`"/home/test/DesktopA"`), homeDir)
		So(err, ShouldBeNil)
		So(value, ShouldEqual, "/home/test/DesktopA")

		value, err = parseValue([]byte(`"$HOME/"`), homeDir)
		So(err, ShouldBeNil)
		So(value, ShouldEqual, "/home/test")

		value, err = parseValue([]byte(`"/"`), homeDir)
		So(err, ShouldBeNil)
		So(value, ShouldEqual, "/")

		value, err = parseValue([]byte(""), homeDir)
		So(err, ShouldNotBeNil)
		So(value, ShouldEqual, "")

		value, err = parseValue([]byte("$HOME"), homeDir)
		So(err, ShouldNotBeNil)
		So(value, ShouldEqual, "")

		value, err = parseValue([]byte(`"not abs"`), homeDir)
		So(err, ShouldNotBeNil)
		So(value, ShouldEqual, "")

	})
}

func TestParseUserDirsConfig(t *testing.T) {
	Convey("parseUserDirsConfig", t, func() {
		os.Setenv("HOME", "/home/test")
		cfg, err := parseUserDirsConfig("./testdata/user-dirs.dirs")
		So(err, ShouldBeNil)
		So(cfg, ShouldResemble, map[string]string{"XDG_DESKTOP_DIR": "/home/test/桌面", "XDG_DOCUMENTS_DIR": "/home/test/文档", "XDG_DOWNLOAD_DIR": "/home/test/下载", "XDG_MUSIC_DIR": "/home/test/音乐", "XDG_PICTURES_DIR": "/home/test/图片", "XDG_PUBLICSHARE_DIR": "/home/test/.Public", "XDG_TEMPLATES_DIR": "/home/test/.Templates", "XDG_VIDEOS_DIR": "/home/test/视频"})
	})
}

func TestGet(t *testing.T) {
	Convey("Get", t, func() {
		os.Setenv("HOME", "/home/test")
		testDataDir, err := filepath.Abs("./testdata")
		So(err, ShouldBeNil)

		os.Setenv("XDG_CONFIG_HOME", testDataDir)

		So(Get(Desktop), ShouldEqual, "/home/test/桌面")
		So(Get(Download), ShouldEqual, "/home/test/下载")
		So(Get(Templates), ShouldEqual, "/home/test/.Templates")
		So(Get(PublicShare), ShouldEqual, "/home/test/.Public")
		So(Get(Documents), ShouldEqual, "/home/test/文档")
		So(Get(Music), ShouldEqual, "/home/test/音乐")
		So(Get(Pictures), ShouldEqual, "/home/test/图片")
		So(Get(Videos), ShouldEqual, "/home/test/视频")
		So(Get("XXXX"), ShouldEqual, "/home/test")
	})
}

func TestReloadCache(t *testing.T) {
	Convey("ReloadCache", t, func() {
		os.Setenv("HOME", "/home/test")
		testDataDir, err := filepath.Abs("./testdata")
		So(err, ShouldBeNil)

		os.Setenv("XDG_CONFIG_HOME", testDataDir)
		So(Get(Desktop), ShouldEqual, "/home/test/桌面")

		testDataDir2, err := filepath.Abs("./testdata2")
		So(err, ShouldBeNil)
		os.Setenv("XDG_CONFIG_HOME", testDataDir2)
		err = ReloadCache()
		So(err, ShouldBeNil)
		So(Get(Desktop), ShouldEqual, "/home/test/MyDesktop")
	})
}
