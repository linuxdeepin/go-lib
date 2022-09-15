// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package util

// IsLeapYear 公历闰年判断
func IsLeapYear(year int) bool {
	return (year&3) == 0 && year%100 != 0 || year%400 == 0
}

var monthDays = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// GetSolarMonthDays 获取公历月份的天数
func GetSolarMonthDays(year, month int) int {
	if month == 2 && IsLeapYear(year) {
		return 29
	} else {
		return monthDays[month-1]
	}
}

// GetYearDaysCount 获取某年的天数
func GetYearDaysCount(year int) int {
	if IsLeapYear(year) {
		return 366
	}
	return 365
}

// ToJulianDate 计算Gregorian日期的儒略日数，以TT当天中午12点为准(结果是整数)。
// 算法摘自 http://en.wikipedia.org/wiki/Julian_day
func ToJulianDate(year, month, day int) int {
	var a int = (14 - month) / 12
	var y int = year + 4800 - a
	var m int = month + 12*a - 3
	return day + (153*m+2)/5 + 365*y + y/4 - y/100 + y/400 - 32045
}

// ToJulianDateHMS 计算Gregorian时间的儒略日数
// 算法摘自 http://en.wikipedia.org/wiki/Julian_day
func ToJulianDateHMS(year, month, day, hour, minute int, second float64) float64 {
	jdn := ToJulianDate(year, month, day)
	return float64(jdn) + (float64(hour)-12)/24.0 + float64(minute)/1440.0 + second/86400.0
}

// Gregorian历TT2000年1月1日中午12点的儒略日
const J2000 = 2451545.0

// GetJulianThousandYears 计算儒略千年数
func GetJulianThousandYears(jd float64) float64 {
	//1000年的日数
	const DaysOf1000Years = 365250.0
	return (jd - J2000) / DaysOf1000Years
}

// GetJulianCentury 计算儒略世纪数
func GetJulianCentury(jd float64) float64 {
	// 100年的日数
	const DaysOfCentury = 36525.0
	return (jd - J2000) / DaysOfCentury
}

// GetWeekday 计算Gregorian日历的星期几
// 算法摘自 http://en.wikipedia.org/wiki/Zeller%27s_congruence
// 返回星期几的数字表示，1-6表示星期一到星期六，0表示星期日
func GetWeekday(y, m, d int) int {
	if m <= 2 {
		y -= 1
		m += 12
	}
	c := int(y / 100)
	y = y % 100
	w := (d + 13*(m+1)/5 + y + (y / 4) + (c / 4) - 2*c - 1) % 7
	if w < 0 {
		w += 7
	}
	return w
}

// GetDeltaT 计算地球时和UTC的时差，算法摘自
// http://eclipse.gsfc.nasa.gov/SEhelp/deltatpoly2004.html NASA网站
// ∆T = TT - UT 此算法在-1999年到3000年有效
func GetDeltaT(year, month int) float64 {
	y := float64(year) + (float64(month)-0.5)/12
	switch {
	case year < -500:
		u := (float64(year) - 1820) / 100
		return -20 + 32*u*u

	case year < 500:
		u := y / 100
		u2 := u * u
		u3 := u2 * u
		u4 := u3 * u
		u5 := u4 * u
		u6 := u5 * u
		return 10583.6 - 1014.41*u + 33.78311*u2 - 5.952053*u3 - 0.1798452*u4 + 0.022174192*u5 + 0.0090316521*u6

	case year < 1600:
		u := (y - 1000) / 100
		u2 := u * u
		u3 := u2 * u
		u4 := u3 * u
		u5 := u4 * u
		u6 := u5 * u
		return 1574.2 - 556.01*u + 71.23472*u2 + 0.319781*u3 - 0.8503463*u4 - 0.005050998*u5 + 0.0083572073*u6

	case year < 1700:
		t := y - 1600
		t2 := t * t
		t3 := t2 * t
		return 120 - 0.9808*t - 0.01532*t2 + t3/7129

	case year < 1800:
		t := y - 1700
		t2 := t * t
		t3 := t2 * t
		t4 := t3 * t
		return 8.83 + 0.1603*t - 0.0059285*t2 + 0.00013336*t3 - t4/1174000

	case year < 1860:
		t := y - 1800
		t2 := t * t
		t3 := t2 * t
		t4 := t3 * t
		t5 := t4 * t
		t6 := t5 * t
		t7 := t6 * t
		return 13.72 - 0.332447*t + 0.0068612*t2 + 0.0041116*t3 - 0.00037436*t4 + 0.0000121272*t5 - 0.0000001699*t6 + 0.000000000875*t7

	case year < 1900:
		t := y - 1860
		t2 := t * t
		t3 := t2 * t
		t4 := t3 * t
		t5 := t4 * t
		return 7.62 + 0.5737*t - 0.251754*t2 + 0.01680668*t3 - 0.0004473624*t4 + t5/233174

	case year < 1920:
		t := y - 1900
		t2 := t * t
		t3 := t2 * t
		t4 := t3 * t
		return -2.79 + 1.494119*t - 0.0598939*t2 + 0.0061966*t3 - 0.000197*t4

	case year < 1941:
		t := y - 1920
		t2 := t * t
		t3 := t2 * t
		return 21.20 + 0.84493*t - 0.076100*t2 + 0.0020936*t3

	case year < 1961:
		t := y - 1950
		t2 := t * t
		t3 := t2 * t
		return 29.07 + 0.407*t - t2/233 + t3/2547

	case year < 1986:
		t := y - 1975
		t2 := t * t
		t3 := t2 * t
		return 45.45 + 1.067*t - t2/260 - t3/718

	case year < 2005:
		t := y - 2000
		t2 := t * t
		t3 := t2 * t
		t4 := t3 * t
		t5 := t4 * t
		return 63.86 + 0.3345*t - 0.060374*t2 + 0.0017275*t3 + 0.000651814*t4 + 0.00002373599*t5

	case year < 2050:
		t := y - 2000
		t2 := t * t
		return 62.92 + 0.32217*t + 0.005589*t2

	case year < 2150:
		u := (y - 1820) / 100
		u2 := u * u
		return -20 + 32*u2 - 0.5628*(2150-y)

	default:
		u := (y - 1820) / 100
		u2 := u * u
		return -20 + 32*u2
	}
}

// JDUTC2BeijingTime 儒略日 UTC 时间转换到北京时间
func JDUTC2BeijingTime(utcJD float64) float64 {
	return utcJD + 8.0/24.0
}

// JDBeijingTime2UTC 儒略日 北京时间到 UTC 时间
func JDBeijingTime2UTC(bjtJD float64) float64 {
	return bjtJD - 8.0/24.0
}
