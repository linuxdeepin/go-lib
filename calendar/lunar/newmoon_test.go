// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package lunar

import (
	"github.com/linuxdeepin/go-lib/calendar/util"
	"testing"
)

func Test_getNewMoonJD(t *testing.T) {
	jd0 := float64(util.ToJulianDate(2012, 1, 10))
	t.Logf("jd0: %f", jd0)
	jd := getNewMoonJD(jd0)
	dt := util.GetDateTimeFromJulianDay(jd)
	t.Log(dt)
}
