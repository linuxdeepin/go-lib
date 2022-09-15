// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package lunar

import (
	"math"

	"github.com/linuxdeepin/go-lib/calendar/util"
)

// 参数 T 儒略世纪数
// 返回 弧度
func GetMoonEclipticParameter(T float64) (Lp, D, M, Mp, F, E float64) {
	T2 := T * T
	T3 := T2 * T
	T4 := T3 * T

	/*月球平黄经*/
	Lp = Mod2Pi(ToRadians(218.3164591 + 481267.88134236*T - 0.0013268*T2 + T3/538841.0 - T4/65194000.0))

	/*月日距角*/
	D = Mod2Pi(ToRadians(297.8502042 + 445267.1115168*T - 0.0016300*T2 + T3/545868.0 - T4/113065000.0))

	/*太阳平近点角*/
	M = Mod2Pi(ToRadians(357.5291092 + 35999.0502909*T - 0.0001536*T2 + T3/24490000.0))

	/*月亮平近点角*/
	Mp = Mod2Pi(ToRadians(134.9634114 + 477198.8676313*T + 0.0089970*T2 + T3/69699.0 - T4/14712000.0))

	/*月球经度参数(到升交点的平角距离)*/
	F = Mod2Pi(ToRadians(93.2720993 + 483202.0175273*T - 0.0034029*T2 - T3/3526000.0 + T4/863310000.0))

	/* 反映地球轨道偏心率变化的辅助参量 */
	E = 1 - 0.002516*T - 0.0000074*T2
	return
}

/*计算月球地心黄经周期项的和*/
func CalcMoonECLongitudePeriodic(D, M, Mp, F, E float64) float64 {
	var EI float64
	for _, l := range MoonLongitude {
		theta := l.D*D + l.M*M + l.Mp*Mp + l.F*F
		EI += l.EiA * math.Sin(theta) * math.Pow(E, math.Abs(l.M))
	}
	// fmt.Printf("EI = %f\n", EI)
	return EI
}

/*计算金星摄动,木星摄动以及地球扁率摄动对月球地心黄经的影响, T 是儒略世纪数，Lp和F单位是弧度*/
// A1 = 119.75 + 131.849 * T                                             （4.13式）
// A2 = 53.09 + 479264.290 * T                                           （4.14式）
// A3 = 313.45 + 481266.484 * T                                          （4.15式）
func CalcMoonLongitudePerturbation(T, Lp, F float64) float64 {
	A1 := Mod2Pi(ToRadians(119.75 + 131.849*T))
	A2 := Mod2Pi(ToRadians(53.09 + 479264.290*T))

	return 3958.0*math.Sin(A1) + 1962.0*math.Sin(Lp-F) + 318.0*math.Sin(A2)
}

/*计算月球地心黄经*/
// jd 儒略日
// 返回 弧度
func GetMoonEclipticLongitudeEC(jd float64) float64 {
	T := util.GetJulianCentury(jd)
	Lp, D, M, Mp, F, E := GetMoonEclipticParameter(T)
	// Lp 计算是正确的
	// fmt.Printf("Lp = %f\n", Lp)

	/*计算月球地心黄经周期项*/
	EI := CalcMoonECLongitudePeriodic(D, M, Mp, F, E)

	/*修正金星,木星以及地球扁率摄动*/
	EI += CalcMoonLongitudePerturbation(T, Lp, F)

	longitude := Lp + ToRadians(EI/1000000.0)

	/*计算天体章动干扰*/
	longitude += CalcEarthLongitudeNutation(T)
	return longitude
}
