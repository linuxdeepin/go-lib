package basedir

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestGetUserDataDir(t *testing.T) {
	Convey("GetUserDataDir", t, func() {
		os.Setenv("XDG_DATA_HOME", "")
		os.Setenv("HOME", "/home/test")
		dir := GetUserDataDir()
		So(dir, ShouldEqual, "/home/test/.local/share")

		os.Setenv("XDG_DATA_HOME", "/home/test/xdg")
		dir = GetUserDataDir()
		So(dir, ShouldEqual, "/home/test/xdg")
	})
}

func TestGetSystemDataDirs(t *testing.T) {
	Convey("GetSystemDirs", t, func() {
		os.Setenv("XDG_DATA_DIRS", "/a:/b:/c")
		dirs := GetSystemDataDirs()
		So(dirs[0], ShouldEqual, "/a")
		So(dirs[1], ShouldEqual, "/b")
		So(dirs[2], ShouldEqual, "/c")
		So(len(dirs), ShouldEqual, 3)

		os.Setenv("XDG_DATA_DIRS", "")
		dirs = GetSystemDataDirs()
		So(dirs[0], ShouldEqual, "/usr/local/share")
		So(dirs[1], ShouldEqual, "/usr/share")
		So(len(dirs), ShouldEqual, 2)

	})
}
