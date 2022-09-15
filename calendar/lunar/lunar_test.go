// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package lunar

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	cc := New(2012)
	const dateTimeFormatStr = "2006-01-02 15:04:05"

	var dtList1 []string
	for _, month := range cc.Months {
		dtList1 = append(dtList1, month.ShuoTime.Format(dateTimeFormatStr))
	}
	dtListStr := strings.Join(dtList1, ",")
	const testDtListStr1 = "2011-11-25 14:09:41,2011-12-25 02:06:27,2012-01-23 15:39:24,2012-02-22 06:34:40,2012-03-22 22:37:08,2012-04-21 15:18:22,2012-05-21 07:46:59,2012-06-19 23:02:06,2012-07-19 12:24:02,2012-08-17 23:54:27,2012-09-16 10:10:36,2012-10-15 20:02:30,2012-11-14 06:08:05,2012-12-13 16:41:37"
	// t.Logf("%q\n", dtListStr)
	assert.Equal(t, dtListStr, testDtListStr1)

	var dtList2 []string
	for _, stTime := range cc.SolarTermTimes {
		dtList2 = append(dtList2, stTime.Format(dateTimeFormatStr))
	}
	const testDtListStr2 = "2011-12-22 13:30:01,2012-01-06 06:43:54,2012-01-21 00:09:48,2012-02-04 18:22:22,2012-02-19 14:17:35,2012-03-05 12:21:01,2012-03-20 13:14:23,2012-04-04 17:05:34,2012-04-20 00:12:02,2012-05-05 10:19:39,2012-05-20 23:15:30,2012-06-05 14:25:52,2012-06-21 07:08:46,2012-07-07 00:40:42,2012-07-22 18:00:50,2012-08-07 10:30:31,2012-08-23 01:06:48,2012-09-07 13:28:59,2012-09-22 22:48:56,2012-10-08 05:11:41,2012-10-23 08:13:32,2012-11-07 08:25:56,2012-11-22 05:50:07,2012-12-07 01:18:55,2012-12-21 19:11:35"
	dtListStr = strings.Join(dtList2, ",")
	// t.Logf("%q\n", dtListStr)
	assert.Equal(t, dtListStr, testDtListStr2)
}

func Test_GetYearZodiac(t *testing.T) {
	assert.Equal(t, GetYearZodiac(2012), "龙")
	assert.Equal(t, GetYearZodiac(2014), "马")
	assert.Equal(t, GetYearZodiac(2015), "羊")
	assert.Equal(t, GetYearZodiac(2016), "猴")
	assert.Equal(t, GetYearZodiac(2017), "鸡")
}

func Test_cyclical(t *testing.T) {
	var ganZhiList []string
	for i := 0; i < 80; i++ {
		ganZhiList = append(ganZhiList, cyclical(i))
	}
	ganzhiListStr := strings.Join(ganZhiList, ",")
	// t.Logf("%q\n", ganzhiListStr )
	const testGanzhiListStr = "甲子,乙丑,丙寅,丁卯,戊辰,己巳,庚午,辛未,壬申,癸酉,甲戌,乙亥,丙子,丁丑,戊寅,己卯,庚辰,辛巳,壬午,癸未,甲申,乙酉,丙戌,丁亥,戊子,己丑,庚寅,辛卯,壬辰,癸巳,甲午,乙未,丙申,丁酉,戊戌,己亥,庚子,辛丑,壬寅,癸卯,甲辰,乙巳,丙午,丁未,戊申,己酉,庚戌,辛亥,壬子,癸丑,甲寅,乙卯,丙辰,丁巳,戊午,己未,庚申,辛酉,壬戌,癸亥,甲子,乙丑,丙寅,丁卯,戊辰,己巳,庚午,辛未,壬申,癸酉,甲戌,乙亥,丙子,丁丑,戊寅,己卯,庚辰,辛巳,壬午,癸未"
	assert.Equal(t, ganzhiListStr, testGanzhiListStr)
}

func Test_GetYearGanZhi(t *testing.T) {
	assert.Equal(t, GetYearGanZhi(2012), "壬辰")
	assert.Equal(t, GetYearGanZhi(2016), "丙申")
	assert.Equal(t, GetYearGanZhi(2020), "庚子")
}
