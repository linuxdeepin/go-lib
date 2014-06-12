/**
 * Copyright (c) 2013 ~ 2014 Deepin, Inc.
 *               2013 ~ 2014 Xu FaSheng
 *
 * Author:      Xu FaSheng <fasheng.xu@gmail.com>
 * Maintainer:  Xu FaSheng <fasheng.xu@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package graphic

import (
	"math"
)

// Rgb2Hsv convert color format from RGB(r, g, b=[0..255]) to HSV(h=[0..360), s,v=[0..1]).
func Rgb2Hsv(r, g, b uint8) (h, s, v float64) {
	fr := float64(r) / 255
	fg := float64(g) / 255
	fb := float64(b) / 255
	max := math.Max(math.Max(fr, fg), fb)
	min := math.Min(math.Min(fr, fg), fb)
	d := max - min

	if max == 0 {
		s = 0
	} else {
		s = d / max
	}

	v = max

	if max == min {
		h = 0
	} else {
		switch max {
		case fr:
			if fg >= fb {
				h = 60 * (fg - fb) / d
			} else {
				h = 60*(fg-fb)/d + 360
			}
		case fg:
			h = 60*(fb-fr)/d + 120
		case fb:
			h = 60*(fr-fg)/d + 240
		}
	}
	return
}

// Hsv2Rgb convert color format from HSV(h=[0..360), s,v=[0..1]) to RGB(r, g, b=[0..255]).
func Hsv2Rgb(h, s, v float64) (r, g, b uint8) {
	var fr, fg, fb float64
	hi := int(math.Floor(h/60)) % 6
	f := h/60 - float64(hi)
	p := v * (1 - s)
	q := v * (1 - f*s)
	t := v * (1 - (1-f)*s)
	switch hi {
	case 0:
		fr, fg, fb = v, t, p
	case 1:
		fr, fg, fb = q, v, p
	case 2:
		fr, fg, fb = p, v, t
	case 3:
		fr, fg, fb = p, q, v
	case 4:
		fr, fg, fb = t, p, v
	case 5:
		fr, fg, fb = v, p, q
	}
	r = uint8((fr * 255) + 0.5)
	g = uint8((fg * 255) + 0.5)
	b = uint8((fb * 255) + 0.5)
	return
}
