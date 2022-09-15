// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package lunar

// Day 保存农历日信息
type Day struct {
	Year       int    // 公历年
	Month      int    // 公历月
	Day        int    // 公历日
	LunarYear  int    // 农历年
	LunarDay   int    // 农历日
	LunarMonth *Month // 农历月
	MonthZhi   int    // 农历日所在的月的地支
	SolarTerm  int    // 0~23 二十四节气 ，-1 非节气
}

// 十二月名
var monthNames = []string{"正", "二", "三", "四", "五", "六", "七", "八", "九", "十", "冬", "腊"}

// MonthName 获取当天的农历月名称
func (d *Day) MonthName() string {
	monthName := monthNames[d.LunarMonth.Name-1]
	if d.LunarMonth.IsLeap {
		return "闰" + monthName + "月"
	}
	return monthName + "月"
}

// 农历日名
// 月份分为大月和小月，大月三十天，小月二十九天
var dayNames = []string{
	"初一", "初二", "初三", "初四", "初五", "初六", "初七", "初八", "初九", "初十",
	"十一", "十二", "十三", "十四", "十五", "十六", "十七", "十八", "十九", "二十",
	"廿一", "廿二", "廿三", "廿四", "廿五", "廿六", "廿七", "廿八", "廿九", "三十",
}

// DayName 获取当天的农历日名
func (d *Day) DayName() string {
	return dayNames[d.LunarDay-1]
}

//农历节日
var lunarFestival = map[int]string{
	101:  "春节",
	115:  "元宵节",
	505:  "端午节",
	707:  "七夕节",
	815:  "中秋节",
	909:  "重阳节",
	1208: "腊八节",
}

// Festival 获取当天的农历节日名
// 没有则返回空字符串
func (d *Day) Festival() string {
	key := d.LunarMonth.Name*100 + d.LunarDay
	if name, ok := lunarFestival[key]; ok {
		return name
	}
	// 农历腊月（十二月）的最后个一天
	if d.LunarMonth.Name == 12 && d.LunarDay == d.LunarMonth.Days {
		return "除夕"
	}
	if d.SolarTerm == QingMing {
		return "清明节"
	}
	return ""
}

// SolarTermName 获取当天的二十四节气名
// 没有则返回空字符串
func (d *Day) SolarTermName() string {
	return GetSolarTermName(d.SolarTerm)
}

// GanZhiMonth 获取当天的月干支
func (d *Day) GanZhiMonth() string {
	return cyclical((d.Year-1900)*12 + d.MonthZhi + 12)
}

// GanZhiYear 获取当天的年干支
func (d *Day) GanZhiYear() string {
	return GetYearGanZhi(d.LunarYear)
}

// GanZhiYear 获取当天的日干支
func (d *Day) GanZhiDay() string {
	return GetDayGanZhi(d.Year, d.Month, d.Day)
}

// YearZodiac 获取当天的生肖
func (d *Day) YearZodiac() string {
	return GetYearZodiac(d.LunarYear)
}
