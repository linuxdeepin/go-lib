/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package mime

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsGtkTheme(t *testing.T) {
	Convey("Deepin is gtk theme", t, func(c C) {
		ok, err := isGtkTheme("testdata/Deepin/index.theme")
		c.So(ok, ShouldEqual, true)
		c.So(err, ShouldBeNil)
	})
}

func TestIsIconTheme(t *testing.T) {
	Convey("Deepin is icon theme", t, func(c C) {
		ok, err := isIconTheme("testdata/Deepin/index.theme")
		c.So(ok, ShouldEqual, true)
		c.So(err, ShouldBeNil)
	})
}

func TestIsCursorTheme(t *testing.T) {
	Convey("Deepin is cursor theme", t, func(c C) {
		ok, err := isCursorTheme("testdata/Deepin/index.theme")
		c.So(ok, ShouldEqual, true)
		c.So(err, ShouldBeNil)
	})
}
