// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package lunar

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMathUtil(t *testing.T) {

	assert.Equal(t, ToRadians(90), math.Pi/2)
	assert.Equal(t, SecondsToRadians(648000), math.Pi)
	assert.Equal(t, SecondsToDegrees(648000), float64(180))
	assert.Equal(t, DmsToDegrees(40, 11, 15), 40.1875)
	assert.Equal(t, DmsToSeconds(40, 11, 15), float64(144675))
	assert.Equal(t, DmsToRadians(40, 11, 15), float64(0.7014041931452212))

	rad := Mod2Pi(3 * math.Pi)
	assert.Equal(t, rad, math.Pi)
	rad = Mod2Pi(-math.Pi)
	assert.Equal(t, rad, math.Pi)
	rad = ModPi(2 * math.Pi)
	assert.Equal(t, rad, float64(0))
	rad = ModPi(-2 * math.Pi)
	assert.Equal(t, rad, float64(0))
}
