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
	"os"
)

// ClipImage clip any recognized format image and save to target format image.
func ClipImage(srcfile, dstfile string, x0, y0, x1, y1 int32, f Format) (err error) {
	sf, err := os.Open(srcfile)
	if err != nil {
		return
	}
	defer sf.Close()

	df, err := openFileOrCreate(dstfile)
	if err != nil {
		return
	}
	defer df.Close()

	srcimg, _, err := image.Decode(sf)
	if err != nil {
		return
	}

	dstimg := doClipImage(srcimg, x0, y0, x1, y1)
	return encodeImage(df, dstimg, f)
}

func doClipImage(srcimg image.Image, x0, y0, x1, y1 int32) (dstimg image.Image) {
	dstimg = image.NewRGBA(image.Rect(int(x0), int(y0), int(x1), int(y1)))
	draw.Draw(dstimg, dstimg.Bounds(), srcimg, image.Point{0, 0}, draw.Src)
	return
}
