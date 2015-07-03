package operations_test

import (
	. "github.com/smartystreets/goconvey/convey"
	. "pkg.linuxdeepin.com/lib/operations"
	"testing"
)

func TestGetTemplateJob(t *testing.T) {
	Convey("Get template from directory which consists of directories", t, func() {
		uri := "./testdata"
		op := NewGetTemplateJob(uri)
		So(len(op.Execute()), ShouldEqual, 0)
	})
}
