/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

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
