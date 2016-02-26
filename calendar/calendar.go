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
