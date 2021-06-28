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
