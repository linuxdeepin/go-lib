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
