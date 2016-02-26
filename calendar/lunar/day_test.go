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
	"testing"
)

func TestDay(t *testing.T) {
	Convey("year 2012", t, func() {
		cc := New(2012)

		Convey("1-1", func() {
			day := cc.SolarDayToLunarDay(1, 1)

			So(day.MonthName(), ShouldEqual, "腊月")
			So(day.DayName(), ShouldEqual, "初八")
			So(day.GanZhiMonth(), ShouldEqual, "庚子")
			So(day.Festival(), ShouldEqual, "腊八节")
			So(day.SolarTermName(), ShouldEqual, "")
		})

		Convey("2-4", func() {
			day := cc.SolarDayToLunarDay(2, 4)

			So(day.MonthName(), ShouldEqual, "正月")
			So(day.DayName(), ShouldEqual, "十三")
			So(day.GanZhiMonth(), ShouldEqual, "壬寅")
			So(day.Festival(), ShouldEqual, "")
			So(day.SolarTermName(), ShouldEqual, "立春")
		})

		Convey("5-20", func() {
			day := cc.SolarDayToLunarDay(5, 20)

			So(day.MonthName(), ShouldEqual, "四月")
			So(day.DayName(), ShouldEqual, "三十")
			So(day.GanZhiMonth(), ShouldEqual, "乙巳")
			So(day.Festival(), ShouldEqual, "")
			So(day.SolarTermName(), ShouldEqual, "小满")
		})

		Convey("6-10", func() {
			day := cc.SolarDayToLunarDay(6, 10)

			So(day.MonthName(), ShouldEqual, "闰四月")
			So(day.DayName(), ShouldEqual, "廿一")
			So(day.GanZhiMonth(), ShouldEqual, "丙午")
			So(day.Festival(), ShouldEqual, "")
			So(day.SolarTermName(), ShouldEqual, "")
		})
	})

	Convey("year 2016", t, func() {
		cc := New(2016)

		Convey("2-26", func() {
			day := cc.SolarDayToLunarDay(2, 26)

			So(day.MonthName(), ShouldEqual, "正月")
			So(day.DayName(), ShouldEqual, "十九")
			So(day.GanZhiMonth(), ShouldEqual, "庚寅")
			So(day.Festival(), ShouldEqual, "")
			So(day.SolarTermName(), ShouldEqual, "")
		})

		Convey("3-5", func() {
			day := cc.SolarDayToLunarDay(3, 5)

			So(day.MonthName(), ShouldEqual, "正月")
			So(day.DayName(), ShouldEqual, "廿七")
			So(day.GanZhiMonth(), ShouldEqual, "辛卯")
			So(day.Festival(), ShouldEqual, "")
			So(day.SolarTermName(), ShouldEqual, "惊蛰")
		})

		Convey("6-9", func() {
			day := cc.SolarDayToLunarDay(6, 9)

			So(day.MonthName(), ShouldEqual, "五月")
			So(day.DayName(), ShouldEqual, "初五")
			So(day.GanZhiMonth(), ShouldEqual, "甲午")
			So(day.Festival(), ShouldEqual, "端午节")
			So(day.SolarTermName(), ShouldEqual, "")
		})

	})
}
