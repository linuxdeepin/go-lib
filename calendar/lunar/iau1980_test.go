// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package lunar

import (
	"testing"
)

func Test_CalcEarthLongitudeNutation(t *testing.T) {
	n := CalcEarthLongitudeNutation(1.2345)
	t.Log(n)
}

func Test_CalcEarthObliquityNutation(t *testing.T) {
	n := CalcEarthObliquityNutation(1.2345)
	t.Log(n)
}
