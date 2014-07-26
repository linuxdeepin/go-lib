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

package gdkpixbuf

// #cgo pkg-config: glib-2.0 gdk-pixbuf-2.0
// #cgo LDFLAGS: -lm
// #include <stdlib.h>
// #include "blur.h"
import "C"

import (
	"fmt"
	"pkg.linuxdeepin.com/lib/utils"
)

func Blur(pixbuf *C.GdkPixbuf, sigma, numsteps float64) (err error) {
	defaultError := fmt.Errorf("blur gdkpixbuf failed, pixbuf=%v, sigma=%v, numsteps=%v", pixbuf, sigma, numsteps)
	ret := C.blur(pixbuf, C.double(sigma), C.double(numsteps))
	if ret == 0 {
		err = defaultError
		return
	}
	return
}

// BlurImage generate blur effect to an image.
func BlurImage(srcFile, dstFile string, sigma, numsteps float64, f Format) (err error) {
	srcPixbuf, err := NewPixbufFromFile(srcFile)
	if err != nil {
		return
	}
	err = Blur(srcPixbuf, sigma, numsteps)
	if err != nil {
		return
	}
	err = Save(srcPixbuf, dstFile, f)
	return
}

// BlurImageCache generate and save the blurred image to cache
// directory, if target file already exists, just return it.
func BlurImageCache(srcFile string, sigma, numsteps float64, f Format) (dstFile string, useCache bool, err error) {
	dstFile = generateCacheFilePath(fmt.Sprintf("BlurImageCache%s%f%f%s", srcFile, sigma, numsteps, f))
	if utils.IsFileExist(dstFile) {
		// return cache file
		useCache = true
		return
	}
	err = BlurImage(srcFile, dstFile, sigma, numsteps, f)
	return
}
