// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package graphic

import (
	"fmt"
	"image"
)

// GetDominantColorOfImage return the dominant hsv color of an image.
func GetDominantColorOfImage(imgfile string) (h, s, v float64, err error) {
	img, err := LoadImage(imgfile)
	if err != nil {
		return
	}
	return doGetDominantColorOfImage(img)
}

func doGetDominantColorOfImage(img image.Image) (h, s, v float64, err error) {
	// loop all points in image
	var sumR, sumG, sumB, count uint64
	mx := img.Bounds().Max.X
	my := img.Bounds().Max.Y
	count = uint64(mx * my)
	if count == 0 {
		err = fmt.Errorf("image is empty")
		return
	}
	for x := 0; x < mx; x++ {
		for y := 0; y < my; y++ {
			c := img.At(x, y)
			rr, gg, bb, _ := c.RGBA()
			r, g, b := rr>>8, gg>>8, bb>>8
			sumR += uint64(r)
			sumG += uint64(g)
			sumB += uint64(b)
		}
	}

	h, s, v = Rgb2Hsv(uint8(sumR/count), uint8(sumG/count), uint8(sumB/count))
	return
}
