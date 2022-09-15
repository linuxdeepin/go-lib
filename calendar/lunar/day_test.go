// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package lunar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay(t *testing.T) {
	cc := New(2012)

	day := cc.SolarDayToLunarDay(1, 1)

	assert.Equal(t, day.MonthName(), "腊月")
	assert.Equal(t, day.DayName(), "初八")
	assert.Equal(t, day.GanZhiYear(), "辛卯")
	assert.Equal(t, day.GanZhiMonth(), "庚子")
	assert.Equal(t, day.GanZhiDay(), "辛酉")
	assert.Equal(t, day.YearZodiac(), "兔")
	assert.Equal(t, day.Festival(), "腊八节")
	assert.Equal(t, day.SolarTermName(), "")

	day = cc.SolarDayToLunarDay(2, 4)

	assert.Equal(t, day.MonthName(), "正月")
	assert.Equal(t, day.DayName(), "十三")
	assert.Equal(t, day.GanZhiYear(), "壬辰")
	assert.Equal(t, day.GanZhiMonth(), "壬寅")
	assert.Equal(t, day.GanZhiDay(), "乙未")
	assert.Equal(t, day.YearZodiac(), "龙")
	assert.Equal(t, day.Festival(), "")
	assert.Equal(t, day.SolarTermName(), "立春")

	day = cc.SolarDayToLunarDay(5, 20)

	assert.Equal(t, day.MonthName(), "四月")
	assert.Equal(t, day.DayName(), "三十")
	assert.Equal(t, day.GanZhiYear(), "壬辰")
	assert.Equal(t, day.GanZhiMonth(), "乙巳")
	assert.Equal(t, day.GanZhiDay(), "辛巳")
	assert.Equal(t, day.YearZodiac(), "龙")
	assert.Equal(t, day.Festival(), "")
	assert.Equal(t, day.SolarTermName(), "小满")

	day = cc.SolarDayToLunarDay(6, 10)

	assert.Equal(t, day.MonthName(), "闰四月")
	assert.Equal(t, day.DayName(), "廿一")
	assert.Equal(t, day.GanZhiYear(), "壬辰")
	assert.Equal(t, day.GanZhiMonth(), "丙午")
	assert.Equal(t, day.GanZhiDay(), "壬寅")
	assert.Equal(t, day.YearZodiac(), "龙")
	assert.Equal(t, day.Festival(), "")
	assert.Equal(t, day.SolarTermName(), "")

	cc = New(2016)
	day = cc.SolarDayToLunarDay(1, 1)
	assert.Equal(t, day.MonthName(), "冬月") // 十一月
	assert.Equal(t, day.DayName(), "廿二")
	assert.Equal(t, day.GanZhiYear(), "乙未")
	assert.Equal(t, day.GanZhiMonth(), "戊子")
	assert.Equal(t, day.GanZhiDay(), "壬午")
	assert.Equal(t, day.YearZodiac(), "羊")
	assert.Equal(t, day.Festival(), "")
	assert.Equal(t, day.SolarTermName(), "")

	day = cc.SolarDayToLunarDay(2, 26)

	assert.Equal(t, day.MonthName(), "正月")
	assert.Equal(t, day.DayName(), "十九")
	assert.Equal(t, day.GanZhiYear(), "丙申")
	assert.Equal(t, day.GanZhiMonth(), "庚寅")
	assert.Equal(t, day.GanZhiDay(), "戊寅")
	assert.Equal(t, day.YearZodiac(), "猴")
	assert.Equal(t, day.Festival(), "")
	assert.Equal(t, day.SolarTermName(), "")

	day = cc.SolarDayToLunarDay(3, 5)

	assert.Equal(t, day.MonthName(), "正月")
	assert.Equal(t, day.DayName(), "廿七")
	assert.Equal(t, day.GanZhiYear(), "丙申")
	assert.Equal(t, day.GanZhiMonth(), "辛卯")
	assert.Equal(t, day.GanZhiDay(), "丙戌")
	assert.Equal(t, day.YearZodiac(), "猴")
	assert.Equal(t, day.Festival(), "")
	assert.Equal(t, day.SolarTermName(), "惊蛰")

	day = cc.SolarDayToLunarDay(6, 9)

	assert.Equal(t, day.MonthName(), "五月")
	assert.Equal(t, day.DayName(), "初五")
	assert.Equal(t, day.GanZhiYear(), "丙申")
	assert.Equal(t, day.GanZhiMonth(), "甲午")
	assert.Equal(t, day.GanZhiDay(), "壬戌")
	assert.Equal(t, day.YearZodiac(), "猴")
	assert.Equal(t, day.Festival(), "端午节")
	assert.Equal(t, day.SolarTermName(), "")

	day = cc.SolarDayToLunarDay(12, 31)

	assert.Equal(t, day.MonthName(), "腊月")
	assert.Equal(t, day.DayName(), "初三")
	assert.Equal(t, day.GanZhiYear(), "丙申")
	assert.Equal(t, day.GanZhiMonth(), "庚子")
	assert.Equal(t, day.GanZhiDay(), "丁亥")
	assert.Equal(t, day.YearZodiac(), "猴")
	assert.Equal(t, day.Festival(), "")
	assert.Equal(t, day.SolarTermName(), "")
}
