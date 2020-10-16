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

package lunar

import (
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

func TestMathUtil(t *testing.T) {
	Convey("convert", t, func(c C) {

		c.Convey("ToRadians", func(c C) {
			c.So(ToRadians(90), ShouldEqual, math.Pi/2)
		})
		c.Convey("SecondsToRadians", func(c C) {
			c.So(SecondsToRadians(648000), ShouldEqual, math.Pi)
		})
		c.Convey("SecondsToDegrees", func(c C) {
			c.So(SecondsToDegrees(648000), ShouldEqual, 180)
		})
		c.Convey("DmsToDegrees", func(c C) {
			c.So(DmsToDegrees(40, 11, 15), ShouldEqual, 40.1875)
		})
		c.Convey("DmsToSeconds", func(c C) {
			c.So(DmsToSeconds(40, 11, 15), ShouldEqual, 144675)
		})
		c.Convey("DmsToRadians", func(c C) {
			c.So(DmsToRadians(40, 11, 15), ShouldAlmostEqual, 0.7014041931452)
		})

	})

	Convey("modpi", t, func(c C) {
		c.Convey("Mod2Pi", func(c C) {
			rad := Mod2Pi(3 * math.Pi)
			c.So(rad, ShouldEqual, math.Pi)
			rad = Mod2Pi(-math.Pi)
			c.So(rad, ShouldEqual, math.Pi)
		})
		c.Convey("ModPi", func(c C) {
			rad := ModPi(2 * math.Pi)
			c.So(rad, ShouldEqual, 0)
			rad = ModPi(-2 * math.Pi)
			c.So(rad, ShouldEqual, 0)
		})
	})

	Convey("NewtonIteration", t, func(c C) {
		var n float64 = 2
		x := NewtonIteration(
			func(x float64) float64 {
				return x*x - n
			},
			1.4)
		c.So(x, ShouldAlmostEqual, math.Sqrt(2), 1e-7)
	})
}
