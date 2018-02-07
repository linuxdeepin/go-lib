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

type LunarDayInfo struct {
	GanZhiYear     string
	GanZhiMonth    string
	GanZhiDay      string
	LunarMonthName string
	LunarDayName   string
	LunarLeapMonth int32
	Zodiac         string
	Term           string
	SolarFestival  string
	LunarFestival  string
	Worktime       int32
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
		GanZhiYear:     lunar.GetYearGanZhi(year),
		GanZhiMonth:    lunarDay.GanZhiMonth(),
		GanZhiDay:      lunar.GetDayGanZhi(year, month, day),
		LunarMonthName: lunarDay.MonthName(),
		LunarDayName:   lunarDay.DayName(),
		Term:           lunarDay.SolarTermName(),
		SolarFestival:  solarDay.Festival(),
		LunarFestival:  lunarDay.Festival(),
		Zodiac:         lunar.GetYearZodiac(year),
	}
	return dayInfo, true
}
