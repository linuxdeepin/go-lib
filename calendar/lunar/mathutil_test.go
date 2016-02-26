/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package lunar

import (
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

func TestMathUtil(t *testing.T) {
	Convey("convert", t, func() {

		Convey("ToRadians", func() {
			So(ToRadians(90), ShouldEqual, math.Pi/2)
		})
		Convey("SecondsToRadians", func() {
			So(SecondsToRadians(648000), ShouldEqual, math.Pi)
		})
		Convey("SecondsToDegrees", func() {
			So(SecondsToDegrees(648000), ShouldEqual, 180)
		})
		Convey("DmsToDegrees", func() {
			So(DmsToDegrees(40, 11, 15), ShouldEqual, 40.1875)
		})
		Convey("DmsToSeconds", func() {
			So(DmsToSeconds(40, 11, 15), ShouldEqual, 144675)
		})
		Convey("DmsToRadians", func() {
			So(DmsToRadians(40, 11, 15), ShouldAlmostEqual, 0.7014041931452)
		})

	})

	Convey("modpi", t, func() {
		Convey("Mod2Pi", func() {
			rad := Mod2Pi(3 * math.Pi)
			So(rad, ShouldEqual, math.Pi)
			rad = Mod2Pi(-math.Pi)
			So(rad, ShouldEqual, math.Pi)
		})
		Convey("ModPi", func() {
			rad := ModPi(2 * math.Pi)
			So(rad, ShouldEqual, 0)
			rad = ModPi(-2 * math.Pi)
			So(rad, ShouldEqual, 0)
		})
	})

	Convey("NewtonIteration", t, func() {
		var n float64 = 2
		x := NewtonIteration(
			func(x float64) float64 {
				return x*x - n
			},
			1.4)
		So(x, ShouldAlmostEqual, math.Sqrt(2), 1e-7)
	})
}
