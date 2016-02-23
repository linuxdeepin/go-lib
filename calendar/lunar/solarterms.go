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
	"math"
	"pkg.deepin.io/lib/calendar/util"
)

var SolarTermNames = []string{
	"春分",
	"清明",
	"谷雨",
	"立夏",
	"小满",
	"芒种",
	"夏至",
	"小暑",
	"大暑",
	"立秋",
	"处暑",
	"白露",
	"秋分",
	"寒露",
	"霜降",
	"立冬",
	"小雪",
	"大雪",
	"冬至",
	"小寒",
	"大寒",
	"立春",
	"雨水",
	"惊蛰",
}

const (
	ChunFen int = iota
	QingMing
	GuYu
	LiXia
	XiaoMan
	MangZhong
	XiaZhi
	XiaoShu
	DaShu
	LiQiu
	ChuShu
	BaiLu
	QiuFen
	HanLu
	ShuangJiang
	LiDong
	XiaoXue
	DaXue
	DongZhi
	XiaoHan
	DaHan
	LiChun
	YuShui
	JingZhe
)

// GetSolarTermName 获取二十四节气名
func GetSolarTermName(order int) string {
	if 0 <= order && order <= 23 {
		return SolarTermNames[order]
	}
	return ""
}

// GetSolarTermJD 使用牛顿迭代法计算24节气的时间
// f(x) = Vsop87dEarthUtil.getEarthEclipticLongitudeForSun(x) - angle = 0
// year 年
// order 节气序号
// 返回 节气的儒略日力学时间 TD
func GetSolarTermJD(year, order int) float64 {
	const RADIANS_PER_TERM = math.Pi / 12.0
	angle := float64(order) * RADIANS_PER_TERM
	month := ((order+1)/2+2)%12 + 1
	// 春分 order 0
	// 3 月 20 号
	var day int = 6
	if order%2 == 0 {
		day = 20
	}

	jd0 := util.ToJulianDateHMS(year, month, day, 12, 0, 0.0)
	jd := NewtonIteration(func(x float64) float64 {
		return ModPi(GetEarthEclipticLongitudeForSun(x) - angle)
	}, jd0)

	return jd
}
