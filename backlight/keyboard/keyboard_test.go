package keyboard

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestList(t *testing.T) {
	controllers, err := List()
	t.Log(err)
	if len(controllers) == 0 {
		t.Log("not found")
		return
	}
	for _, c := range controllers {
		t.Logf("%+v\n", c)
		br, _ := c.GetBrightness()
		t.Log("brightness", br)
	}
}

func Test_list(t *testing.T) {
	Convey("Test list", t, func() {
		controllers, err := list("./testdata")
		So(err, ShouldBeNil)
		So(controllers, ShouldHaveLength, 1)

		Convey("Test Controller", func() {
			c := controllers[0]
			So(c.Name, ShouldEqual, "xxx::kbd_backlight")
			So(c.MaxBrightness, ShouldEqual, 3)

			br, err := c.GetBrightness()
			So(err, ShouldBeNil)
			So(br, ShouldEqual, 1)

		})
	})

}
