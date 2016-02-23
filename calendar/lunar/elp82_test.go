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
	"testing"
)

func Test_GetMoonEclipticParameter(t *testing.T) {
	Lp, D, M, Mp, F, E := GetMoonEclipticParameter(1.2345)
	t.Log(Lp)
	t.Log(D)
	t.Log(M)
	t.Log(Mp)
	t.Log(F)
	t.Log(E)
}

func Test_CalcMoonECLongitudePeriodic(t *testing.T) {
	Lp, D, M, Mp, F, E := GetMoonEclipticParameter(1.2345)
	EI := CalcMoonECLongitudePeriodic(D, M, Mp, F, E)
	t.Log(Lp)
	t.Log(EI)
}

func Test_GetMoonEclipticLongitudeEC(t *testing.T) {
	rad := GetMoonEclipticLongitudeEC(2448724.5)
	t.Log(rad)
	// 目标 2.3242072713306796
	deg := ToDegrees(rad)
	t.Log(deg)
}
