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
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

// GetDominantColorOfImage return the dominant hsv color of a image.
func GetDominantColorOfImage(imgfile string) (h, s, v float64) {
	var defH, defS, defV float64 = 200, 0.5, 0.8 // default hsv

	// open the image file
	fr, err := os.Open(imgfile)
	if err != nil {
		log.Printf(err.Error()) // TODO
		return defH, defS, defV
	}
	defer fr.Close()

	img, _, err := image.Decode(fr)
	if err != nil {
		log.Printf(err.Error()) // TODO
		return defH, defS, defV
	}

	// loop all points in image
	var sumR, sumG, sumB, count uint64
	mx := img.Bounds().Max.X
	my := img.Bounds().Max.Y
	count = uint64(mx * my)
	if count == 0 {
		return defH, defS, defV
	}
	if mx == 0 && my == 0 {
		return defH, defS, defV
	}
	for x := 1; x <= mx; x++ {
		for y := 1; y <= my; y++ {
			c := img.At(x, y)
			rr, gg, bb, _ := c.RGBA()
			r, g, b := rr>>8, gg>>8, bb>>8
			sumR += uint64(r)
			sumG += uint64(g)
			sumB += uint64(b)
		}
	}

	h, s, v = RGB2HSV(uint8(sumR/count), uint8(sumG/count), uint8(sumB/count))
	log.Printf("h=%f, s=%f, v=%f", h, s, v) // TODO
	return
}
