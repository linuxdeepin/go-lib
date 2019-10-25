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
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDay(t *testing.T) {
	Convey("year 2012", t, func() {
		cc := New(2012)

		Convey("01-01", func() {
			day := cc.SolarDayToLunarDay(1, 1)

			So(day.MonthName(), ShouldEqual, "腊月")
			So(day.DayName(), ShouldEqual, "初八")
			So(day.GanZhiYear(), ShouldEqual, "辛卯")
			So(day.GanZhiMonth(), ShouldEqual, "庚子")
			So(day.GanZhiDay(), ShouldEqual, "辛酉")
			So(day.YearZodiac(), ShouldEqual, "兔")
			So(day.Festival(), ShouldEqual, "腊八节")
			So(day.SolarTermName(), ShouldEqual, "")
		})

		Convey("02-04", func() {
			day := cc.SolarDayToLunarDay(2, 4)

			So(day.MonthName(), ShouldEqual, "正月")
			So(day.DayName(), ShouldEqual, "十三")
			So(day.GanZhiYear(), ShouldEqual, "壬辰")
			So(day.GanZhiMonth(), ShouldEqual, "壬寅")
			So(day.GanZhiDay(), ShouldEqual, "乙未")
			So(day.YearZodiac(), ShouldEqual, "龙")
			So(day.Festival(), ShouldEqual, "")
			So(day.SolarTermName(), ShouldEqual, "立春")
		})

		Convey("05-20", func() {
			day := cc.SolarDayToLunarDay(5, 20)

			So(day.MonthName(), ShouldEqual, "四月")
			So(day.DayName(), ShouldEqual, "三十")
			So(day.GanZhiYear(), ShouldEqual, "壬辰")
			So(day.GanZhiMonth(), ShouldEqual, "乙巳")
			So(day.GanZhiDay(), ShouldEqual, "辛巳")
			So(day.YearZodiac(), ShouldEqual, "龙")
			So(day.Festival(), ShouldEqual, "")
			So(day.SolarTermName(), ShouldEqual, "小满")
		})

		Convey("06-10", func() {
			day := cc.SolarDayToLunarDay(6, 10)

			So(day.MonthName(), ShouldEqual, "闰四月")
			So(day.DayName(), ShouldEqual, "廿一")
			So(day.GanZhiYear(), ShouldEqual, "壬辰")
			So(day.GanZhiMonth(), ShouldEqual, "丙午")
			So(day.GanZhiDay(), ShouldEqual, "壬寅")
			So(day.YearZodiac(), ShouldEqual, "龙")
			So(day.Festival(), ShouldEqual, "")
			So(day.SolarTermName(), ShouldEqual, "")
		})
	})

	Convey("year 2016", t, func() {
		cc := New(2016)
		Convey("01-01", func() {
			day := cc.SolarDayToLunarDay(1, 1)
			So(day.MonthName(), ShouldEqual, "冬月") // 十一月
			So(day.DayName(), ShouldEqual, "廿二")
			So(day.GanZhiYear(), ShouldEqual, "乙未")
			So(day.GanZhiMonth(), ShouldEqual, "戊子")
			So(day.GanZhiDay(), ShouldEqual, "壬午")
			So(day.YearZodiac(), ShouldEqual, "羊")
			So(day.Festival(), ShouldEqual, "")
			So(day.SolarTermName(), ShouldEqual, "")
		})

		Convey("02-26", func() {
			day := cc.SolarDayToLunarDay(2, 26)

			So(day.MonthName(), ShouldEqual, "正月")
			So(day.DayName(), ShouldEqual, "十九")
			So(day.GanZhiYear(), ShouldEqual, "丙申")
			So(day.GanZhiMonth(), ShouldEqual, "庚寅")
			So(day.GanZhiDay(), ShouldEqual, "戊寅")
			So(day.YearZodiac(), ShouldEqual, "猴")
			So(day.Festival(), ShouldEqual, "")
			So(day.SolarTermName(), ShouldEqual, "")
		})

		Convey("03-05", func() {
			day := cc.SolarDayToLunarDay(3, 5)

			So(day.MonthName(), ShouldEqual, "正月")
			So(day.DayName(), ShouldEqual, "廿七")
			So(day.GanZhiYear(), ShouldEqual, "丙申")
			So(day.GanZhiMonth(), ShouldEqual, "辛卯")
			So(day.GanZhiDay(), ShouldEqual, "丙戌")
			So(day.YearZodiac(), ShouldEqual, "猴")
			So(day.Festival(), ShouldEqual, "")
			So(day.SolarTermName(), ShouldEqual, "惊蛰")
		})

		Convey("06-09", func() {
			day := cc.SolarDayToLunarDay(6, 9)

			So(day.MonthName(), ShouldEqual, "五月")
			So(day.DayName(), ShouldEqual, "初五")
			So(day.GanZhiYear(), ShouldEqual, "丙申")
			So(day.GanZhiMonth(), ShouldEqual, "甲午")
			So(day.GanZhiDay(), ShouldEqual, "壬戌")
			So(day.YearZodiac(), ShouldEqual, "猴")
			So(day.Festival(), ShouldEqual, "端午节")
			So(day.SolarTermName(), ShouldEqual, "")
		})

		Convey("12-31", func() {
			day := cc.SolarDayToLunarDay(12, 31)

			So(day.MonthName(), ShouldEqual, "腊月")
			So(day.DayName(), ShouldEqual, "初三")
			So(day.GanZhiYear(), ShouldEqual, "丙申")
			So(day.GanZhiMonth(), ShouldEqual, "庚子")
			So(day.GanZhiDay(), ShouldEqual, "丁亥")
			So(day.YearZodiac(), ShouldEqual, "猴")
			So(day.Festival(), ShouldEqual, "")
			So(day.SolarTermName(), ShouldEqual, "")
		})
	})
}
