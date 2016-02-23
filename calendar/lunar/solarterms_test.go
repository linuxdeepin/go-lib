/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package lunar

import (
	. "github.com/smartystreets/goconvey/convey"
	"pkg.deepin.io/lib/calendar/util"
	"strings"
	"testing"
)

func TestSolarTerms(t *testing.T) {
	Convey("GetSolarTermName", t, func() {

		So(GetSolarTermName(DongZhi), ShouldEqual, "冬至")

		var stNameList []string
		for i := -1; i < 25; i++ {
			name := GetSolarTermName(i)
			stNameList = append(stNameList, name)
		}
		stNameListStr := strings.Join(stNameList, ",")
		const testStNameListStr = ",春分,清明,谷雨,立夏,小满,芒种,夏至,小暑,大暑,立秋,处暑,白露,秋分,寒露,霜降,立冬,小雪,大雪,冬至,小寒,大寒,立春,雨水,惊蛰,"
		So(stNameListStr, ShouldEqual, testStNameListStr)
	})

	Convey("GetSolarTermJD", t, func() {
		dongZhiJD := GetSolarTermJD(2016, DongZhi)
		dt := util.GetDateTimeFromJulianDay(dongZhiJD)
		So(dt.String(), ShouldEqual, "2016-12-21 10:44:08 +0000 UTC")
	})
}
