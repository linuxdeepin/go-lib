// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package calendar

import "strings"

type Day struct {
	Year, Month, Day int
}

type festival struct {
	name      string
	startYear int
}

var solarFestivals = map[int][]festival{
	101: {
		{
			name:      "元旦",
			startYear: 0,
		},
	},
	214: {
		{
			name:      "情人节",
			startYear: 0,
		},
	},
	305: {
		{
			name:      "学雷锋纪念日",
			startYear: 1963,
		},
	},
	308: {
		{
			name:      "妇女节",
			startYear: 0,
		},
	},
	312: {
		{
			name:      "植树节",
			startYear: 0,
		},
	},
	401: {
		{
			name:      "愚人节",
			startYear: 0,
		},
	},
	415: {
		{
			name:      "全民国家安全教育日",
			startYear: 0,
		},
	},
	501: {
		{
			name:      "劳动节",
			startYear: 0,
		},
	},
	504: {
		{
			name:      "青年节",
			startYear: 0,
		},
	},
	601: {
		{
			name:      "儿童节",
			startYear: 0,
		},
	},
	701: {
		{
			name:      "建党节",
			startYear: 0,
		},
		{
			name:      "香港回归纪念日",
			startYear: 1997,
		},
	},
	801: {
		{
			name:      "建军节",
			startYear: 0,
		},
	},
	903: {
		{
			name:      "抗日战争胜利纪念日",
			startYear: 1945,
		},
	},
	910: {
		{
			name:      "教师节",
			startYear: 0,
		},
	},
	1001: {
		{
			name:      "国庆节",
			startYear: 0,
		},
	},
	1213: {
		{
			name:      "南京大屠杀死难者国家公祭日",
			startYear: 1937,
		},
	},
	1220: {
		{
			name:      "澳门回归纪念",
			startYear: 1999,
		},
	},
	1224: {
		{
			name:      "平安夜",
			startYear: 0,
		},
	},
	1225: {
		{
			name:      "圣诞节",
			startYear: 0,
		},
	},
	1226: {
		{
			name:      "毛泽东诞辰纪念",
			startYear: 0,
		},
	},
}

func (d *Day) Festival() string {
	year := d.Year
	month := d.Month
	day := d.Day
	name := ""
	var festivals []string
	if (month == 5) || (month == 6) {
		name = festivalForFatherAndMother(year, month, day)
		if name != "" {
			festivals = append(festivals, name)
		}
	}
	key := month*100 + day
	if solarFestival, ok := solarFestivals[key]; ok {
		for _, festival := range solarFestival {
			if festival.startYear <= year {
				festivals = append(festivals, festival.name)
			}
		}
	}
	return strings.Join(festivals, ",")
}

func festivalForFatherAndMother(year, month, day int) string {
	var disparityMotherDay, disparityFatherDay, fatherDay, i, motherDay int
	var leapYear int
	for i = 1900; i <= year; i++ {
		if (i%400 == 0) || ((i%100 != 0) && (i%4 == 0)) {
			leapYear = leapYear + 1
		}
	}
	if month == 5 {
		disparityMotherDay = (((year-1899)*365 + leapYear) - (31 + 30 + 31 + 31 + 30 + 31 + 30 + 31)) % 7
		motherDay = 14 - disparityMotherDay
		if day == motherDay {
			return "母亲节"
		} else {
			return ""
		}
	}
	if month == 6 {
		disparityFatherDay = (((year-1899)*365 + leapYear) - (30 + 31 + 31 + 30 + 31 + 30 + 31)) % 7
		fatherDay = 21 - disparityFatherDay
		if day == fatherDay {
			return "父亲节"
		} else {
			return ""
		}
	}

	return ""

}
