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
	_image "image"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"os"
)

// Converts from any recognized format to PNG.
func ConvertToPNG(src, dest string) (err error) {
	sf, err := os.Open(src)
	if err != nil {
		return
	}
	defer sf.Close()
	df, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer df.Close()

	img, _, err := _image.Decode(sf)
	if err != nil {
		return
	}
	return png.Encode(df, img)
}

// Clip any recognized format image and save to PNG.
func ClipPNG(src, dest string, x0, y0, x1, y1 int32) (err error) {
	sf, err := os.Open(src)
	if err != nil {
		return
	}
	defer sf.Close()

	df, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer df.Close()

	imgSrc, _, err := _image.Decode(sf)
	if err != nil {
		return
	}

	imgDest := _image.NewRGBA(_image.Rect(int(x0), int(y0), int(x1), int(y1)))
	draw.Draw(imgDest, imgDest.Bounds(), imgSrc, _image.Point{0, 0}, draw.Src)
	return png.Encode(df, imgDest)
}
