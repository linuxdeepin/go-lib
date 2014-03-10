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

// #cgo pkg-config: glib-2.0 gdk-pixbuf-2.0
// #cgo LDFLAGS: -lm
// #include <stdlib.h>
// #include "blur_pict.h"
import "C"
import "unsafe"
import "fmt"

// BlurImage generate blur effect to an image.
// TODO Format always is PNG
func BlurImage(srcfile, dstfile string, sigma, numsteps float64, f Format) (err error) {
	ok := generateBlurPict(srcfile, dstfile, sigma, numsteps)
	if !ok {
		err = fmt.Errorf("generate blur pict failed")
	}
	return
}

func generateBlurPict(srcfile, dstfile string, sigma, numsteps float64) bool {
	src := C.CString(srcfile)
	defer C.free(unsafe.Pointer(src))
	dest := C.CString(dstfile)
	defer C.free(unsafe.Pointer(dest))

	ok := C.generate_blur_pict(src, dest, C.double(sigma), C.double(numsteps))
	if ok == 0 {
		return false
	}
	return true
}
