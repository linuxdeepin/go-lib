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

package util

import (
	"testing"
	// "time"
)

func Test_IsLeapYear(t *testing.T) {
	if IsLeapYear(1996) && IsLeapYear(2000) && !IsLeapYear(2001) {
		t.Log("ok")
	} else {
		t.Error("fail")
	}
}

func Test_GetSolarMonthDays(t *testing.T) {
	if 31 == GetSolarMonthDays(2001, 1) {
		t.Log("ok")
	} else {
		t.Error("fail")
	}
	if 29 == GetSolarMonthDays(2016, 2) {
		t.Log("ok")
	} else {
		t.Error("fail")
	}
	if 28 == GetSolarMonthDays(2015, 2) {
		t.Log("ok")
	} else {
		t.Error("fail")
	}
}

func Test_ToJulianDate(t *testing.T) {
	if 2443230 == ToJulianDate(1977, 3, 27) &&
		2453522 == ToJulianDate(2005, 5, 31) {
		t.Log("ok")
	} else {
		t.Error("fail")
	}
}

func Test_ToJulianDateHMS(t *testing.T) {
	t.Log(ToJulianDateHMS(1977, 3, 27, 6, 6, 6))
	if 2.4432297542361114e+06 == ToJulianDateHMS(1977, 3, 27, 6, 6, 6) {
		t.Log("ok")
	} else {
		t.Error("fail")
	}
}

func Test_GetJulianThousandYears(t *testing.T) {
	t.Log(GetJulianThousandYears(2.2324937542361114e+06))
	if -0.5997296256369298 == GetJulianThousandYears(2.2324937542361114e+06) {
		t.Log("ok")
	} else {
		t.Error("fail")
	}
}

func Test_GetJulianCentury(t *testing.T) {
	t.Log(GetJulianCentury(2.2324937542361114e+06))
	if -5.997296256369297 == GetJulianCentury(2.2324937542361114e+06) {
		t.Log("ok")
	} else {
		t.Error("fail")
	}
}

func Test_GetWeekday(t *testing.T) {
	if 4 == GetWeekday(2015, 1, 1) {
		t.Log("ok")
	} else {
		t.Error("fail")
	}
}
