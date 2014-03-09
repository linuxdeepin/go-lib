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
	"image/draw"
)

func FlipImageHorizontal(srcfile, dstfile string, f Format) (err error) {
	srcimg, err := loadImage(srcfile)
	if err != nil {
		return err
	}
	dstimg := doFlipImageHorizontal(srcimg)
	return saveImage(dstfile, dstimg, f)
}

func FlipImageVertical(srcfile, dstfile string, f Format) (err error) {
	srcimg, err := loadImage(srcfile)
	if err != nil {
		return err
	}
	dstimg := doFlipImageVertical(srcimg)
	return saveImage(dstfile, dstimg, f)
}

// FIXME return draw.Image or *image.RGBA
func doFlipImageHorizontal(srcimg image.Image) (dstimg draw.Image) {
	w, h, _ := doGetImageSize(srcimg)
	dstimg = image.NewRGBA(image.Rect(0, 0, w, h))

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			dstimg.Set(w-x-1, y, srcimg.At(x, y))
		}
	}

	return
}

// FIXME return draw.Image or *image.RGBA
func doFlipImageVertical(srcimg image.Image) (dstimg draw.Image) {
	w, h, _ := doGetImageSize(srcimg)
	dstimg = image.NewRGBA(image.Rect(0, 0, w, h))

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			dstimg.Set(x, h-y-1, srcimg.At(x, y))
		}
	}

	return
}
