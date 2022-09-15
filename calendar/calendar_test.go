// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package calendar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SolarToLunar(t *testing.T) {
	dayInfo, _ := SolarToLunar(2016, 2, 19)
	assert.Equal(t, dayInfo.LunarMonthName, "正月")
	assert.Equal(t, dayInfo.LunarDayName, "十二")
	assert.Equal(t, dayInfo.Term, "雨水")
	assert.Equal(t, dayInfo.Zodiac, "猴")
	assert.Equal(t, dayInfo.GanZhiYear, "丙申")
	assert.Equal(t, dayInfo.GanZhiMonth, "庚寅")
	assert.Equal(t, dayInfo.GanZhiDay, "辛未")

	dayInfo, _ = SolarToLunar(2012, 1, 23)
	assert.Equal(t, dayInfo.LunarMonthName, "正月")
	assert.Equal(t, dayInfo.LunarDayName, "初一")
	assert.Equal(t, dayInfo.Term, "")
	assert.Equal(t, dayInfo.LunarFestival, "春节")
	assert.Equal(t, dayInfo.Zodiac, "龙")
	assert.Equal(t, dayInfo.GanZhiYear, "壬辰")
	assert.Equal(t, dayInfo.GanZhiMonth, "辛丑")
	assert.Equal(t, dayInfo.GanZhiDay, "癸未")

	dayInfo, _ = SolarToLunar(2014, 6, 1)
	assert.Equal(t, dayInfo.LunarMonthName, "五月")
	assert.Equal(t, dayInfo.LunarDayName, "初四")
	assert.Equal(t, dayInfo.Term, "")
	assert.Equal(t, dayInfo.LunarFestival, "")
	assert.Equal(t, dayInfo.SolarFestival, "儿童节")
	assert.Equal(t, dayInfo.Zodiac, "马")
	assert.Equal(t, dayInfo.GanZhiYear, "甲午")
	assert.Equal(t, dayInfo.GanZhiMonth, "己巳")
	assert.Equal(t, dayInfo.GanZhiDay, "癸卯")

	dayInfo, _ = SolarToLunar(2057, 10, 27)
	assert.Equal(t, dayInfo.LunarMonthName, "九月")
	assert.Equal(t, dayInfo.LunarDayName, "三十")
	assert.Equal(t, dayInfo.Term, "")
	assert.Equal(t, dayInfo.LunarFestival, "")
	assert.Equal(t, dayInfo.Zodiac, "牛")
	assert.Equal(t, dayInfo.GanZhiYear, "丁丑")
	assert.Equal(t, dayInfo.GanZhiMonth, "庚戌")
	assert.Equal(t, dayInfo.GanZhiDay, "丁巳")
}
