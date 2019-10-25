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
	"pkg.deepin.io/lib/calendar/lunar"
)

/*

参考资料：
算法系列之十七：日历生成算法-中国公历（格里历）（上）
https://blog.csdn.net/orbit/article/details/7749723

算法系列之十七：日历生成算法-中国公历（格里历）（下）
https://blog.csdn.net/orbit/article/details/7825004

算法系列之十八：用天文方法计算二十四节气（上）
https://blog.csdn.net/orbit/article/details/7910220

算法系列之十八：用天文方法计算二十四节气（下）
https://blog.csdn.net/orbit/article/details/7944248

算法系列之十九：用天文方法计算日月合朔（新月）
https://blog.csdn.net/orbit/article/details/8223751


算法系列之二十：计算中国农历（一）
https://blog.csdn.net/orbit/article/details/9210413

算法系列之二十：计算中国农历（二）
https://blog.csdn.net/orbit/article/details/9337377

参考项目：github.com/oyyq99999/ChineseLunarCalendar

*/

type LunarDayInfo struct {
	GanZhiYear     string // 农历年的干支
	GanZhiMonth    string // 农历月的干支
	GanZhiDay      string // 农历日的干支
	LunarMonthName string // 农历月名
	LunarDayName   string // 农历日名
	LunarLeapMonth int32  // 未使用
	Zodiac         string // 农历年的生肖
	Term           string // 农历节气
	SolarFestival  string // 公历节日
	LunarFestival  string // 农历节日
	Worktime       int32  // 未使用
}

func SolarToLunar(year, month, day int) (LunarDayInfo, bool) {
	solarDay := Day{
		Year:  year,
		Month: month,
		Day:   day,
	}
	cc := lunar.New(year)
	lunarDay := cc.SolarDayToLunarDay(month, day)

	dayInfo := LunarDayInfo{
		GanZhiYear:     lunarDay.GanZhiYear(),
		GanZhiMonth:    lunarDay.GanZhiMonth(),
		GanZhiDay:      lunarDay.GanZhiDay(),
		LunarMonthName: lunarDay.MonthName(),
		LunarDayName:   lunarDay.DayName(),
		Term:           lunarDay.SolarTermName(),
		SolarFestival:  solarDay.Festival(),
		LunarFestival:  lunarDay.Festival(),
		Zodiac:         lunarDay.YearZodiac(),
	}
	return dayInfo, true
}
