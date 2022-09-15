// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package util

import (
	"testing"
)

func Test_GetDateFromJulianDay(t *testing.T) {
	y, m, d := GetDateFromJulianDay(2457438.0)
	t.Log(y, m, d)
	if y == 2016 && m == 2 && d == 19 {
		t.Log("ok")
	} else {
		t.Error("fail")
	}

	y, m, d = GetDateFromJulianDay(2248528.0)
	t.Log(y, m, d)
	if y == 1444 && m == 2 && d == 19 {
		t.Log("ok")
	} else {
		t.Error("fail")
	}
}

func Test_GetTimeFromJulianDay(t *testing.T) {
	h, m, s := GetTimeFromJulianDay(2457438.09546)
	t.Log(h, m, s)
	if h == 14 && m == 17 && s == 27 {
		t.Log("ok")
	} else {
		t.Error("fail")
	}

	h, m, s = GetTimeFromJulianDay(2457438.09851)
	t.Log(h, m, s)
	if h == 14 && m == 21 && s == 51 {
		t.Log("ok")
	} else {
		t.Error("fail")
	}
}

func Test_GetDateTimeFromJulianDay(t *testing.T) {
	dt := GetDateTimeFromJulianDay(2457438.10454)
	t.Log(dt)
	if dt.String() == "2016-02-19 14:29:22 +0000 UTC" {
		t.Log("ok")
	} else {
		t.Error("fail")
	}
}
