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
)

// ToRadians 角度转换为弧度
func ToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// ToDegrees 弧度转换为角度
func ToDegrees(rad float64) float64 {
	return rad * 180 / math.Pi
}

// SecondsToRadians 把角秒换算成弧度
func SecondsToRadians(seconds float64) float64 {
	return ToRadians(SecondsToDegrees(seconds))
}

// Mod2Pi 把角度限制在[0, 2π]之间
func Mod2Pi(r float64) float64 {
	for r < 0 {
		r += math.Pi * 2
	}
	for r > 2*math.Pi {
		r -= math.Pi * 2
	}
	return r
}

// ModPi 把角度限制在[-π, π]之间
func ModPi(r float64) float64 {
	for r < -math.Pi {
		r += math.Pi * 2
	}
	for r > math.Pi {
		r -= math.Pi * 2
	}
	return r
}

// SecondsToDegrees 把角秒换算成角度
func SecondsToDegrees(seconds float64) float64 {
	return seconds / 3600
}

// DmsToDegrees 把度分秒表示的角度换算成度
func DmsToDegrees(degrees int, mintues int, seconds float64) float64 {
	return float64(degrees) + float64(mintues)/60 + seconds/3600
}

// DmsToSeconds 把度分秒表示的角度换算成角秒(arcsecond)
func DmsToSeconds(d int, m int, s float64) float64 {
	return float64(d)*3600 + float64(m)*60 + s
}

// DmsToRadians 把度分秒表示的角度换算成弧度(rad)
func DmsToRadians(d int, m int, s float64) float64 {
	return ToRadians(DmsToDegrees(d, m, s))
}

// NewtonIteration 牛顿迭代法求解方程的根
func NewtonIteration(f func(float64) float64, x0 float64) float64 {
	const Epsilon = 1e-7
	const Delta = 5e-6
	var x float64

	for {
		x = x0
		fx := f(x)
		// 导数
		fpx := (f(x+Delta) - f(x-Delta)) / Delta / 2
		x0 = x - fx/fpx
		if math.Abs(x0-x) <= Epsilon {
			break
		}
	}
	return x
}
