package mime

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsGtkTheme(t *testing.T) {
	Convey("Deepin is gtk theme", t, func() {
		ok, err := isGtkTheme("testdata/Deepin/index.theme")
		So(ok, ShouldEqual, true)
		So(err, ShouldBeNil)
	})
}

func TestIsIconTheme(t *testing.T) {
	Convey("Deepin is icon theme", t, func() {
		ok, err := isIconTheme("testdata/Deepin/index.theme")
		So(ok, ShouldEqual, true)
		So(err, ShouldBeNil)
	})
}

func TestIsCursorTheme(t *testing.T) {
	Convey("Deepin is cursor theme", t, func() {
		ok, err := isCursorTheme("testdata/Deepin/index.theme")
		So(ok, ShouldEqual, true)
		So(err, ShouldBeNil)
	})
}
