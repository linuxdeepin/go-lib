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
	"pkg.deepin.io/lib/calendar/util"
	"testing"
)

func Test_getNewMoonJD(t *testing.T) {
	jd0 := float64(util.ToJulianDate(2012, 1, 10))
	t.Logf("jd0: %f", jd0)
	jd := getNewMoonJD(jd0)
	dt := util.GetDateTimeFromJulianDay(jd)
	t.Log(dt)
}
