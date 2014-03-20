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

// ClipImage clip any recognized format image and save to target format image.
func ClipImage(srcfile, dstfile string, x0, y0, x1, y1 int32, f Format) (err error) {
	srcimg, err := LoadImage(srcfile)
	if err != nil {
		return
	}
	dstimg := doClipImage(srcimg, int(x0), int(y0), int(x1), int(y1))
	return SaveImage(dstfile, dstimg, f)
}

// FIXME return draw.Image or *image.RGBA
func doClipImage(srcimg image.Image, x0, y0, x1, y1 int) (dstimg draw.Image) {
	dstimg = image.NewRGBA(image.Rect(0, 0, x1-x0, y1-y0))
	draw.Draw(dstimg, dstimg.Bounds(), srcimg, image.Point{x0, y0}, draw.Src)
	return
}
