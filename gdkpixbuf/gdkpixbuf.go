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

// #cgo pkg-config: gdk-3.0 gdk-pixbuf-xlib-2.0 gdk-x11-3.0 x11
// #cgo LDFLAGS: -lm
// #include <stdlib.h>
// #include "gdk_pixbuf_utils.h"
import "C"
import "unsafe"

import (
	"fmt"
	"github.com/BurntSushi/xgb/xproto"
)

// Format defines the type of image format.
type Format string

// Registered image format.
const (
	FormatPng  Format = "png"
	FormatJpeg Format = "jpeg"
	FormatBmp  Format = "bmp"
	FormatIco  Format = "ico"
	FormatGif  Format = "gif"
	FormatTiff Format = "tiff"
	FormatXpm  Format = "xpm"
)

func InitGdkXlib() (err error) {
	ret := C.init_gdk_xlib()
	if ret == 0 {
		err = fmt.Errorf("initialize gdk xlib failed %v", ret)
	}
	return
}

func FreePixbuf(pixbuf *C.GdkPixbuf) {
	if pixbuf != nil {
		C.g_object_unref(C.gpointer(pixbuf))
	}
}

// Save and Load

func Save(pixbuf *C.GdkPixbuf, destFile string, f Format) (err error) {
	defaultError := fmt.Errorf("render image to xpixmap failed, %s", destFile)
	cDestFile := C.CString(destFile)
	defer C.free(unsafe.Pointer(cDestFile))
	cFormat := C.CString(string(f))
	defer C.free(unsafe.Pointer(cFormat))
	ret := C.save(pixbuf, cDestFile, cFormat)
	if ret == 0 {
		err = defaultError
		return
	}
	return
}

func NewPixbufFromFile(imgFile string) (pixbuf *C.GdkPixbuf, err error) {
	defaultError := fmt.Errorf("render image to xpixmap failed, %s", imgFile)
	cImgFile := C.CString(imgFile)
	defer C.free(unsafe.Pointer(cImgFile))

	// new gdk pixbuf from file
	pixbuf = C.new_pixbuf_from_file(cImgFile)
	if pixbuf == nil {
		err = defaultError
		return
	}
	return
}

// Info

// GetImageSize return image's width and height.
func GetImageSize(imgFile string) (w, h int, err error) {
	defaultError := fmt.Errorf("get image size failed, %s", imgFile)
	cImgFile := C.CString(imgFile)
	defer C.free(unsafe.Pointer(cImgFile))

	w = 0
	h = 0
	ret := C.get_image_size(cImgFile, (*C.int)(unsafe.Pointer(&w)), (*C.int)(unsafe.Pointer(&h)))
	if ret == 0 {
		err = defaultError
		return
	}
	return
}

// GetImageSize return image's width and height.
func GetPixbufSize(pixbuf *C.GdkPixbuf) (w, h int, err error) {
	defaultError := fmt.Errorf("get GdkPixbuf size failed, %v", pixbuf)
	if pixbuf == nil {
		err = defaultError
		return
	}
	w = int(C.gdk_pixbuf_get_width(pixbuf))
	h = int(C.gdk_pixbuf_get_height(pixbuf))
	return
}

// GetImageFormat return image format, such as "png", "jpeg".
func GetImageFormat(imgFile string) (f Format, err error) {
	defaultError := fmt.Errorf("get image format failed, %s", imgFile)
	cImgFile := C.CString(imgFile)
	defer C.free(unsafe.Pointer(cImgFile))

	f = Format(C.GoString(C.get_image_format(cImgFile)))
	if len(f) == 0 {
		err = defaultError
		return
	}
	return
}

// IsSupportedImage check if image file is supported.
func IsSupportedImage(imgFile string) bool {
	_, err := GetImageFormat(imgFile)
	if err != nil {
		return false
	}
	return true
}

// Clip

func CopyArea(srcPixbuf, destPixbuf *C.GdkPixbuf, srcX, srcY, width, height, destX, destY int) (err error) {
	defaultError := fmt.Errorf("copy pixbuf area failed, %v, %v", srcPixbuf, destPixbuf)
	if srcPixbuf == nil || destPixbuf == nil {
		err = defaultError
		return
	}
	C.gdk_pixbuf_copy_area(srcPixbuf, C.int(srcX), C.int(srcY), C.int(width), C.int(height), destPixbuf, C.int(destX), C.int(destY))
	return
}

func CopyAreaSimple(srcPixbuf *C.GdkPixbuf, srcX, srcY, width, height int) (destPixbuf *C.GdkPixbuf, err error) {
	defaultError := fmt.Errorf("copy pixbuf area simple failed, %v", srcPixbuf)
	if srcPixbuf == nil {
		err = defaultError
		return
	}
	destPixbuf = C.copy_area_simple(srcPixbuf, C.int(srcX), C.int(srcY), C.int(width), C.int(height))
	if destPixbuf == nil {
		err = defaultError
		return
	}
	return
}

// Convert

func ConvertPixbufToXpixmap(pixbuf *C.GdkPixbuf) (xpixmap xproto.Pixmap, err error) {
	defaultError := fmt.Errorf("convert pixbuf to xpixmap failed, %v", pixbuf)
	xpixmap = xproto.Pixmap(C.convert_pixbuf_to_xpixmap(pixbuf))
	if xpixmap == 0 {
		err = defaultError
		return
	}
	return
}

func ConvertXpixmapToPixbuf(xpixmap xproto.Pixmap, width, heigth int) (pixbuf *C.GdkPixbuf, err error) {
	defaultError := fmt.Errorf("convert xpixmap to pixbuf failed, %v", xpixmap)
	pixbuf = C.convert_xpixmap_to_pixbuf(C.Pixmap(xpixmap), C.int(width), C.int(heigth))
	if pixbuf == nil {
		err = defaultError
		return
	}
	return
}

func ConvertImgToXpixmap(imgFile string) (xpixmap xproto.Pixmap, err error) {
	// new gdk pixbuf from file
	pixbuf, err := NewPixbufFromFile(imgFile)
	defer FreePixbuf(pixbuf)
	if err != nil {
		return
	}
	// convert pixbuf to xpixmap
	xpixmap, err = ConvertPixbufToXpixmap(pixbuf)
	if err != nil {
		return
	}
	return
}

// Flip

func Flip(srcPixbuf *C.GdkPixbuf, horizontal bool) (destPixbuf *C.GdkPixbuf, err error) {
	defaultError := fmt.Errorf("flip pixbuf failed, %v, %v", srcPixbuf, horizontal)
	if horizontal {
		destPixbuf = C.gdk_pixbuf_flip(srcPixbuf, C.TRUE)
	} else {
		destPixbuf = C.gdk_pixbuf_flip(srcPixbuf, C.FALSE)
	}
	if destPixbuf == nil {
		err = defaultError
		return
	}
	return
}

// Resize

func ScaleSimple(srcPixbuf *C.GdkPixbuf, destWidth, destHeght int, interpType GdkInterpType) (destPixbuf *C.GdkPixbuf, err error) {
	defaultError := fmt.Errorf("scale pixbuf failed, %v, %v, %v, %v", srcPixbuf, destWidth, destHeght, interpType)
	destPixbuf = C.gdk_pixbuf_scale_simple(srcPixbuf, C.int(destWidth), C.int(destHeght), C.GdkInterpType(interpType))
	if destPixbuf == nil {
		err = defaultError
		return
	}
	return
}

// Rotate

func RotateSimple(srcPixbuf *C.GdkPixbuf, angle GdkPixbufRotation) (destPixbuf *C.GdkPixbuf, err error) {
	defaultError := fmt.Errorf("rotate pixbuf failed, %v, %v", srcPixbuf, angle)
	destPixbuf = C.gdk_pixbuf_rotate_simple(srcPixbuf, C.GdkPixbufRotation(angle))
	if destPixbuf == nil {
		err = defaultError
		return
	}
	return
}
