/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package lunar

func getNewMoonJD(jd0 float64) float64 {
	jd := NewtonIteration(
		func(x float64) float64 {
			return ModPi(GetEarthEclipticLongitudeForSun(x) - GetMoonEclipticLongitudeEC(x))
		}, jd0)
	return jd
}
