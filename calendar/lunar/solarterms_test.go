// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package lunar

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/linuxdeepin/go-lib/calendar/util"
)

func TestSolarTerms(t *testing.T) {

	assert.Equal(t, GetSolarTermName(DongZhi), "冬至")

	var stNameList []string
	for i := -1; i < 25; i++ {
		name := GetSolarTermName(i)
		stNameList = append(stNameList, name)
	}
	stNameListStr := strings.Join(stNameList, ",")
	const testStNameListStr = ",春分,清明,谷雨,立夏,小满,芒种,夏至,小暑,大暑,立秋,处暑,白露,秋分,寒露,霜降,立冬,小雪,大雪,冬至,小寒,大寒,立春,雨水,惊蛰,"
	assert.Equal(t, stNameListStr, testStNameListStr)

	dongZhiJD := GetSolarTermJD(2016, DongZhi)
	dt := util.GetDateTimeFromJulianDay(dongZhiJD)
	assert.Equal(t, dt.String(), "2016-12-21 10:44:08 +0000 UTC")
}
