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

import gd "github.com/bolknote/go-gd"
import "image"
import "os"
import "strings"
import "fmt"

// FastBlurImage generate blur effect to an image through stack blur algorithm.
func FastBlurImage(srcfile, dstfile string, radius int32, f Format) (err error) {
	img, err := loadGDImage(srcfile)
	if err != nil {
		return
	}
	img.StackBlur(10, true)
	saveGDImage(img, dstfile, f)
	return
}

func loadGDImage(imgfile string) (img *gd.Image, err error) {
	file, err := os.Open(imgfile)
	if err != nil {
		return
	}
	defer file.Close()
	_, format, err := image.Decode(file)
	if err != nil {
		return
	}
	if strings.Contains(format, string(PNG)) {
		img = gd.CreateFromPng(imgfile)
	} else if strings.Contains(format, string(JPEG)) {
		img = gd.CreateFromJpeg(imgfile)
	} else {
		err = fmt.Errorf("unregisted image format")
	}
	return
}

func saveGDImage(img *gd.Image, dstfile string, f Format) {
	switch f {
	default:
		img.Png(dstfile)
	case PNG:
		img.Png(dstfile)
	case JPEG:
		img.Jpeg(dstfile, 9)
	}
}
