// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package lunar

func getNewMoonJD(jd0 float64) float64 {
	jd := NewtonIteration(
		func(x float64) float64 {
			return ModPi(GetEarthEclipticLongitudeForSun(x) - GetMoonEclipticLongitudeEC(x))
		}, jd0)
	return jd
}
