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

// ResizeImage returns a new image file with the given width and
// height created by resizing the given image.
func ResizeImage(srcfile, dstfile string, newWidth, newHeight int32, f Format) (err error) {
	srcimg, err := LoadImage(srcfile)
	if err != nil {
		return
	}
	dstimg := doResizeNearestNeighbor(srcimg, int(newWidth), int(newHeight))
	return SaveImage(dstfile, dstimg, f)
}

// TODO doResizeNearestNeighbor returns a new RGBA image with the given width and
// height created by resizing the given image using the nearest neighbor
// algorithm.
func doResizeNearestNeighbor(img image.Image, newWidth, newHeight int) (newimg draw.Image) {
	w := img.Bounds().Max.X
	h := img.Bounds().Max.Y
	newimg = image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	xr := (w<<16)/newWidth + 1
	yr := (h<<16)/newHeight + 1

	for yo := 0; yo < newHeight; yo++ {
		y2 := (yo * yr) >> 16
		for xo := 0; xo < newWidth; xo++ {
			x2 := (xo * xr) >> 16
			newimg.Set(xo, yo, img.At(x2, y2))
			//Much faster, but requires some image type.
			//newimg.Pix[offset] = img.Pix[y2*w+x2]
			//offset++
		}
	}
	return newimg
}
