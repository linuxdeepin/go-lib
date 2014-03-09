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
	libgraphic "code.google.com/p/graphics-go/graphics"
	"image"
	"math"
)

func RotateImageLeft(srcfile, dstfile string, f Format) (err error) {
	srcimg, err := loadImage(srcfile)
	if err != nil {
		return err
	}
	w, h, err := doGetImageSize(srcimg)
	if err != nil {
		return err
	}
	dstimg := image.NewRGBA(image.Rect(0, 0, int(h), int(w)))
	err = libgraphic.Rotate(dstimg, srcimg, &libgraphic.RotateOptions{math.Pi / -2})
	if err != nil {
		return err
	}
	return saveImage(dstfile, dstimg, f)
}

func RotateImageRight(srcfile, dstfile string, f Format) (err error) {
	srcimg, err := loadImage(srcfile)
	if err != nil {
		return err
	}
	w, h, err := doGetImageSize(srcimg)
	if err != nil {
		return err
	}
	dstimg := image.NewRGBA(image.Rect(0, 0, int(h), int(w)))
	err = libgraphic.Rotate(dstimg, srcimg, &libgraphic.RotateOptions{math.Pi / 2})
	if err != nil {
		return err
	}
	return saveImage(dstfile, dstimg, f)
}
