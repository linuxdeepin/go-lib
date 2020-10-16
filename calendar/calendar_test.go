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

package calendar

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_SolarToLunar(t *testing.T) {
	Convey("2016-02-19", t, func(c C) {
		dayInfo, _ := SolarToLunar(2016, 2, 19)
		c.So(dayInfo.LunarMonthName, ShouldEqual, "正月")
		c.So(dayInfo.LunarDayName, ShouldEqual, "十二")
		c.So(dayInfo.Term, ShouldEqual, "雨水")
		c.So(dayInfo.Zodiac, ShouldEqual, "猴")
		c.So(dayInfo.GanZhiYear, ShouldEqual, "丙申")
		c.So(dayInfo.GanZhiMonth, ShouldEqual, "庚寅")
		c.So(dayInfo.GanZhiDay, ShouldEqual, "辛未")
	})

	Convey("2012-01-23", t, func(c C) {
		dayInfo, _ := SolarToLunar(2012, 1, 23)
		c.So(dayInfo.LunarMonthName, ShouldEqual, "正月")
		c.So(dayInfo.LunarDayName, ShouldEqual, "初一")
		c.So(dayInfo.Term, ShouldEqual, "")
		c.So(dayInfo.LunarFestival, ShouldEqual, "春节")
		c.So(dayInfo.Zodiac, ShouldEqual, "龙")
		c.So(dayInfo.GanZhiYear, ShouldEqual, "壬辰")
		c.So(dayInfo.GanZhiMonth, ShouldEqual, "辛丑")
		c.So(dayInfo.GanZhiDay, ShouldEqual, "癸未")

	})

	Convey("2014-06-01", t, func(c C) {
		dayInfo, _ := SolarToLunar(2014, 6, 1)
		c.So(dayInfo.LunarMonthName, ShouldEqual, "五月")
		c.So(dayInfo.LunarDayName, ShouldEqual, "初四")
		c.So(dayInfo.Term, ShouldEqual, "")
		c.So(dayInfo.LunarFestival, ShouldEqual, "")
		c.So(dayInfo.SolarFestival, ShouldEqual, "儿童节")
		c.So(dayInfo.Zodiac, ShouldEqual, "马")
		c.So(dayInfo.GanZhiYear, ShouldEqual, "甲午")
		c.So(dayInfo.GanZhiMonth, ShouldEqual, "己巳")
		c.So(dayInfo.GanZhiDay, ShouldEqual, "癸卯")
	})

	Convey("2057-10-27", t, func(c C) {
		dayInfo, _ := SolarToLunar(2057, 10, 27)
		c.So(dayInfo.LunarMonthName, ShouldEqual, "九月")
		c.So(dayInfo.LunarDayName, ShouldEqual, "三十")
		c.So(dayInfo.Term, ShouldEqual, "")
		c.So(dayInfo.LunarFestival, ShouldEqual, "")
		c.So(dayInfo.Zodiac, ShouldEqual, "牛")
		c.So(dayInfo.GanZhiYear, ShouldEqual, "丁丑")
		c.So(dayInfo.GanZhiMonth, ShouldEqual, "庚戌")
		c.So(dayInfo.GanZhiDay, ShouldEqual, "丁巳")
	})
}
