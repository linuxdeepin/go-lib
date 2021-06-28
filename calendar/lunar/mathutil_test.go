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
