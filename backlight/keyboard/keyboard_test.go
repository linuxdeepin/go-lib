/*
 * Copyright (C) 2017 ~ 2018 Deepin Technology Co., Ltd.
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
