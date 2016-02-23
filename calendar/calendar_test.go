/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package calendar

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_SolarToLunar(t *testing.T) {
	Convey("2016-02-19", t, func() {
		dayInfo, _ := SolarToLunar(2016, 2, 19)
		So(dayInfo.LunarMonthName, ShouldEqual, "正月")
		So(dayInfo.LunarDayName, ShouldEqual, "十二")
		So(dayInfo.Term, ShouldEqual, "雨水")
		So(dayInfo.Zodiac, ShouldEqual, "猴")
		So(dayInfo.GanZhiYear, ShouldEqual, "丙申")
		So(dayInfo.GanZhiMonth, ShouldEqual, "庚寅")
		So(dayInfo.GanZhiDay, ShouldEqual, "辛未")
	})

	Convey("2012-01-23", t, func() {
		dayInfo, _ := SolarToLunar(2012, 1, 23)
		So(dayInfo.LunarMonthName, ShouldEqual, "正月")
		So(dayInfo.LunarDayName, ShouldEqual, "初一")
		So(dayInfo.Term, ShouldEqual, "")
		So(dayInfo.LunarFestival, ShouldEqual, "春节")
		So(dayInfo.Zodiac, ShouldEqual, "龙")
		So(dayInfo.GanZhiYear, ShouldEqual, "壬辰")
		So(dayInfo.GanZhiMonth, ShouldEqual, "辛丑")
		So(dayInfo.GanZhiDay, ShouldEqual, "癸未")

	})

	Convey("2014-06-01", t, func() {
		dayInfo, _ := SolarToLunar(2014, 6, 1)
		So(dayInfo.LunarMonthName, ShouldEqual, "五月")
		So(dayInfo.LunarDayName, ShouldEqual, "初四")
		So(dayInfo.Term, ShouldEqual, "")
		So(dayInfo.LunarFestival, ShouldEqual, "")
		So(dayInfo.SolarFestival, ShouldEqual, "国际儿童节")
		So(dayInfo.Zodiac, ShouldEqual, "马")
		So(dayInfo.GanZhiYear, ShouldEqual, "甲午")
		So(dayInfo.GanZhiMonth, ShouldEqual, "己巳")
		So(dayInfo.GanZhiDay, ShouldEqual, "癸卯")
	})
}
