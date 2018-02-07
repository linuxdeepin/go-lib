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
