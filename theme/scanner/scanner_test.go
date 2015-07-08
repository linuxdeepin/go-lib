package scanner

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestListGtkTheme(t *testing.T) {
	Convey("List gtk theme", t, func() {
		list, err := ListGtkTheme("testdata")
		So(list, ShouldResemble, []string{
			"testdata/Deepin",
			"testdata/Theme1"})
		So(err, ShouldBeNil)
	})
}

func TestListIconTheme(t *testing.T) {
	Convey("List icon theme", t, func() {
		list, err := ListIconTheme("testdata")
		So(list, ShouldResemble, []string{
			"testdata/Deepin",
			"testdata/Theme1"})
		So(err, ShouldBeNil)
	})
}

func TestListCursorTheme(t *testing.T) {
	Convey("List cursor theme", t, func() {
		list, err := ListCursorTheme("testdata")
		So(list, ShouldResemble, []string{
			"testdata/Deepin",
			"testdata/Theme1"})
		So(err, ShouldBeNil)
	})
}
