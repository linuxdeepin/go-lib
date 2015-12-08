package mime

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestQueryURI(t *testing.T) {
	Convey("Query uri mime", t, func() {
		var infos = []struct {
			uri  string
			mime string
		}{
			{
				uri:  "testdata/data.txt",
				mime: "text/plain",
			},
			{
				uri:  "testdata/Deepin/index.theme",
				mime: MimeTypeGtk,
			},
		}

		for _, info := range infos {
			m, err := Query(info.uri)
			So(m, ShouldEqual, info.mime)
			So(err, ShouldBeNil)
		}
	})
}
