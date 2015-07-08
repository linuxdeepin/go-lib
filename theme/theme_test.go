package theme

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestMergeThemeList(t *testing.T) {
	Convey("Merge theme list", t, func() {
		src := []string{"Deepin", "Adwaita", "Zukitwo"}
		target := []string{"Deepin", "Evolve"}
		ret := []string{"Deepin", "Adwaita", "Zukitwo", "Evolve"}

		So(mergeThemeList(src, target), ShouldResemble, ret)
	})
}

func TestSetQt4Theme(t *testing.T) {
	Convey("Set qt4 theme", t, func() {
		config := "/tmp/Trolltech.conf"
		So(setQt4Theme(config), ShouldEqual, true)
		os.Remove(config)
	})
}
