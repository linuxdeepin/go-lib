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
	Convey("year 2012", t, func(c C) {
		cc := New(2012)

		c.Convey("01-01", func(c C) {
			day := cc.SolarDayToLunarDay(1, 1)

			c.So(day.MonthName(), ShouldEqual, "腊月")
			c.So(day.DayName(), ShouldEqual, "初八")
			c.So(day.GanZhiYear(), ShouldEqual, "辛卯")
			c.So(day.GanZhiMonth(), ShouldEqual, "庚子")
			c.So(day.GanZhiDay(), ShouldEqual, "辛酉")
			c.So(day.YearZodiac(), ShouldEqual, "兔")
			c.So(day.Festival(), ShouldEqual, "腊八节")
			c.So(day.SolarTermName(), ShouldEqual, "")
		})

		c.Convey("02-04", func(c C) {
			day := cc.SolarDayToLunarDay(2, 4)

			c.So(day.MonthName(), ShouldEqual, "正月")
			c.So(day.DayName(), ShouldEqual, "十三")
			c.So(day.GanZhiYear(), ShouldEqual, "壬辰")
			c.So(day.GanZhiMonth(), ShouldEqual, "壬寅")
			c.So(day.GanZhiDay(), ShouldEqual, "乙未")
			c.So(day.YearZodiac(), ShouldEqual, "龙")
			c.So(day.Festival(), ShouldEqual, "")
			c.So(day.SolarTermName(), ShouldEqual, "立春")
		})

		c.Convey("05-20", func(c C) {
			day := cc.SolarDayToLunarDay(5, 20)

			c.So(day.MonthName(), ShouldEqual, "四月")
			c.So(day.DayName(), ShouldEqual, "三十")
			c.So(day.GanZhiYear(), ShouldEqual, "壬辰")
			c.So(day.GanZhiMonth(), ShouldEqual, "乙巳")
			c.So(day.GanZhiDay(), ShouldEqual, "辛巳")
			c.So(day.YearZodiac(), ShouldEqual, "龙")
			c.So(day.Festival(), ShouldEqual, "")
			c.So(day.SolarTermName(), ShouldEqual, "小满")
		})

		c.Convey("06-10", func(c C) {
			day := cc.SolarDayToLunarDay(6, 10)

			c.So(day.MonthName(), ShouldEqual, "闰四月")
			c.So(day.DayName(), ShouldEqual, "廿一")
			c.So(day.GanZhiYear(), ShouldEqual, "壬辰")
			c.So(day.GanZhiMonth(), ShouldEqual, "丙午")
			c.So(day.GanZhiDay(), ShouldEqual, "壬寅")
			c.So(day.YearZodiac(), ShouldEqual, "龙")
			c.So(day.Festival(), ShouldEqual, "")
			c.So(day.SolarTermName(), ShouldEqual, "")
		})
	})

	Convey("year 2016", t, func(c C) {
		cc := New(2016)
		c.Convey("01-01", func(c C) {
			day := cc.SolarDayToLunarDay(1, 1)
			c.So(day.MonthName(), ShouldEqual, "冬月") // 十一月
			c.So(day.DayName(), ShouldEqual, "廿二")
			c.So(day.GanZhiYear(), ShouldEqual, "乙未")
			c.So(day.GanZhiMonth(), ShouldEqual, "戊子")
			c.So(day.GanZhiDay(), ShouldEqual, "壬午")
			c.So(day.YearZodiac(), ShouldEqual, "羊")
			c.So(day.Festival(), ShouldEqual, "")
			c.So(day.SolarTermName(), ShouldEqual, "")
		})

		c.Convey("02-26", func(c C) {
			day := cc.SolarDayToLunarDay(2, 26)

			c.So(day.MonthName(), ShouldEqual, "正月")
			c.So(day.DayName(), ShouldEqual, "十九")
			c.So(day.GanZhiYear(), ShouldEqual, "丙申")
			c.So(day.GanZhiMonth(), ShouldEqual, "庚寅")
			c.So(day.GanZhiDay(), ShouldEqual, "戊寅")
			c.So(day.YearZodiac(), ShouldEqual, "猴")
			c.So(day.Festival(), ShouldEqual, "")
			c.So(day.SolarTermName(), ShouldEqual, "")
		})

		c.Convey("03-05", func(c C) {
			day := cc.SolarDayToLunarDay(3, 5)

			c.So(day.MonthName(), ShouldEqual, "正月")
			c.So(day.DayName(), ShouldEqual, "廿七")
			c.So(day.GanZhiYear(), ShouldEqual, "丙申")
			c.So(day.GanZhiMonth(), ShouldEqual, "辛卯")
			c.So(day.GanZhiDay(), ShouldEqual, "丙戌")
			c.So(day.YearZodiac(), ShouldEqual, "猴")
			c.So(day.Festival(), ShouldEqual, "")
			c.So(day.SolarTermName(), ShouldEqual, "惊蛰")
		})

		c.Convey("06-09", func(c C) {
			day := cc.SolarDayToLunarDay(6, 9)

			c.So(day.MonthName(), ShouldEqual, "五月")
			c.So(day.DayName(), ShouldEqual, "初五")
			c.So(day.GanZhiYear(), ShouldEqual, "丙申")
			c.So(day.GanZhiMonth(), ShouldEqual, "甲午")
			c.So(day.GanZhiDay(), ShouldEqual, "壬戌")
			c.So(day.YearZodiac(), ShouldEqual, "猴")
			c.So(day.Festival(), ShouldEqual, "端午节")
			c.So(day.SolarTermName(), ShouldEqual, "")
		})

		c.Convey("12-31", func(c C) {
			day := cc.SolarDayToLunarDay(12, 31)

			c.So(day.MonthName(), ShouldEqual, "腊月")
			c.So(day.DayName(), ShouldEqual, "初三")
			c.So(day.GanZhiYear(), ShouldEqual, "丙申")
			c.So(day.GanZhiMonth(), ShouldEqual, "庚子")
			c.So(day.GanZhiDay(), ShouldEqual, "丁亥")
			c.So(day.YearZodiac(), ShouldEqual, "猴")
			c.So(day.Festival(), ShouldEqual, "")
			c.So(day.SolarTermName(), ShouldEqual, "")
		})
	})
}
