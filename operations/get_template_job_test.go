package operations_test

import (
	. "pkg.linuxdeepin.com/lib/operations"
	. "github.com/smartystreets/goconvey/convey"
	"net/url"
	"testing"
)

func TestGetTemplateJob(t *testing.T) {
	Convey("Get template from directory which consists of directories", t, func() {
		uri, err := url.Parse("./testdata")
		So(err, ShouldBeNil)
		op := NewGetTemplateJob(uri)
		So(len(op.Execute()), ShouldEqual, 0)
	})
}
