// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package lunar

import (
	"github.com/linuxdeepin/go-lib/calendar/util"
	"testing"
)

//行星日心黄经（L）、日心黄纬（B）和到太阳的距离（R)
// L0~5 B0~4
// func GetEarthL0 (t float64) float64 {
// func GetEarthL1 (t float64) float64 {
// func GetEarthL2 (t float64) float64 {
// func GetEarthL3 (t float64) float64 {
// func GetEarthL4 (t float64) float64 {
// func GetEarthL5 (t float64) float64 {
// func GetEarthB0 (t float64) float64 {
// func GetEarthB1 (t float64) float64 {
// func GetEarthB2 (t float64) float64 {
// func GetEarthB3 (t float64) float64 {
// func GetEarthB4 (t float64) float64 {

func Test_GetEarthXX(te *testing.T) {
	const t = 1.2345
	L0 := GetEarthL0(t)
	L1 := GetEarthL1(t)
	L2 := GetEarthL2(t)
	L3 := GetEarthL3(t)
	L4 := GetEarthL4(t)
	L5 := GetEarthL5(t)
	te.Log(L0)
	te.Log(L1)
	te.Log(L2)
	te.Log(L3)
	te.Log(L4)
	te.Log(L5)

	B0 := GetEarthB0(t)
	B1 := GetEarthB1(t)
	B2 := GetEarthB2(t)
	B3 := GetEarthB3(t)
	B4 := GetEarthB4(t)
	te.Log(B0)
	te.Log(B1)
	te.Log(B2)
	te.Log(B3)
	te.Log(B4)

	R0 := GetEarthR0(t)
	R1 := GetEarthR1(t)
	R2 := GetEarthR2(t)
	R3 := GetEarthR3(t)
	R4 := GetEarthR4(t)
	R5 := GetEarthR5(t)
	te.Log(R0)
	te.Log(R1)
	te.Log(R2)
	te.Log(R3)
	te.Log(R4)
	te.Log("R5 = ", R5)

}

func Test_GetSunEclipticLongitudeForEarth(t *testing.T) {
	jd := util.ToJulianDateHMS(2016, 2, 16, 6, 6, 6.0)
	L := GetSunEclipticLongitudeForEarth(jd)
	t.Log(L)
}

func Test_GetSunEclipticLatitudeForEarth(t *testing.T) {
	jd := util.ToJulianDateHMS(2016, 2, 16, 6, 6, 6.0)
	B := GetSunEclipticLatitudeForEarth(jd)
	t.Log(B)
}

func Test_GetSunRadiusForEarth(t *testing.T) {
	jd := util.ToJulianDateHMS(2016, 2, 16, 6, 6, 6.0)
	R := GetSunRadiusForEarth(jd)
	t.Log(R)
}
