// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

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
