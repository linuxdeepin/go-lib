package operations_test

import (
	. "pkg.deepin.io/lib/operations"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestShortenUtf8String(t *testing.T) {
	Convey("shorten english", t, func() {
		a := "abc"
		b := ShortenUtf8String(a, 0)
		So(b, ShouldEqual, "abc")

		c := ShortenUtf8String(a, 1)
		So(c, ShouldEqual, "ab")
	})

	Convey("shorten chiness", t, func() {
		a := "ab我卡"

		b := ShortenUtf8String(a, 0)
		So(b, ShouldEqual, "ab我卡")

		c := ShortenUtf8String(a, 1)
		So(c, ShouldEqual, "ab我")

		d := ShortenUtf8String(a, 2)
		So(d, ShouldEqual, "ab")

		e := ShortenUtf8String(a, 3)
		So(e, ShouldEqual, "a")
	})

	Convey("the reduce num is out of range", t, func() {
		a := "aaa"
		b := ShortenUtf8String(a, len(a)+1)
		So(b, ShouldEqual, "")
	})

	Convey("the reduce num is a negative", t, func() {
		a := "xxx"
		b := ShortenUtf8String(a, -1)
		So(b, ShouldEqual, a)
	})
}
