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

package graph

import (
	_image "image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func GetDominantColorOfImage(imagePath string) (h, s, v float64) {
	var def_h, def_s, def_v float64 = 200, 0.5, 0.8 // default hsv

	// open the image file
	fr, err := os.Open(imagePath)
	if err != nil {
		log.Printf(err.Error()) // TODO
		return def_h, def_s, def_v
	}
	defer fr.Close()

	img, _, err := _image.Decode(fr)
	if err != nil {
		log.Printf(err.Error()) // TODO
		return def_h, def_s, def_v
	}

	// loop all points in image
	var sum_r, sum_g, sum_b, count uint64
	mx := img.Bounds().Max.X
	my := img.Bounds().Max.Y
	count = uint64(mx * my)
	if count == 0 {
		return def_h, def_s, def_v
	}
	if mx == 0 && my == 0 {
		return def_h, def_s, def_v
	}
	for x := 1; x <= mx; x++ {
		for y := 1; y <= my; y++ {
			c := img.At(x, y)
			rr, gg, bb, _ := c.RGBA()
			r, g, b := rr>>8, gg>>8, bb>>8
			sum_r += uint64(r)
			sum_g += uint64(g)
			sum_b += uint64(b)
		}
	}

	h, s, v = RGB2HSV(uint8(sum_r/count), uint8(sum_g/count), uint8(sum_b/count))
	log.Printf("h=%f, s=%f, v=%f", h, s, v) // TODO
	return
}
