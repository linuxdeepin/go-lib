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
	"image/color"
)

// NewImageWithColor create a new image file with target size and rgba.
func NewImageWithColor(dstfile string, width, height int, r, g, b, a uint8, f Format) (err error) {
	dstimg, err := NewWithColor(width, height, r, g, b, a)
	if err != nil {
		return
	}
	err = SaveImage(dstfile, dstimg, f)
	dstimg.Pix = nil
	return
}

// NewWithColor create a new image object with target size and rgba.
func NewWithColor(width, height int, r, g, b, a uint8) (dstimg *image.RGBA, err error) {
	dstimg = image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	color := color.RGBA{R: r, G: g, B: b, A: a}
	w, h := GetSize(dstimg)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			dstimg.SetRGBA(i, j, color)
		}
	}
	return
}
