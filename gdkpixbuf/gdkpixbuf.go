// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package gdkpixbuf

// #cgo pkg-config: gdk-3.0 gdk-pixbuf-xlib-2.0 gdk-x11-3.0 x11
// #cgo LDFLAGS: -lm
// #include <stdlib.h>
// #include "gdk_pixbuf_utils.h"
import "C"

import (
	"errors"
	"fmt"
	"image"
	"unsafe"

	x "github.com/linuxdeepin/go-x11-client"
	"github.com/linuxdeepin/go-lib/utils"
)

// Format defines the type of image format.
type Format string

// Registered image format.
const (
	FormatPng  Format = "png"
	FormatJpeg Format = "jpeg"
	FormatBmp  Format = "bmp"
	FormatIco  Format = "ico"
	FormatTiff Format = "tiff"
	// TODO
	// FormatGif  Format = "gif"
	// FormatXpm  Format = "xpm"
)

// InitGdkXlib initialize gdk and xlib, should not be used with InitGdk().
func InitGdkXlib() (err error) {
	ret := C.init_gdk_xlib()
	if ret == 0 {
		err = fmt.Errorf("initialize gdk xlib failed %v", ret)
	}
	return
}

// InitGdk initialize gdk.
func InitGdk() (err error) {
	ret := C.init_gdk()
	if ret == 0 {
		err = fmt.Errorf("initialize gdk failed %v", ret)
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
	defaultError := fmt.Errorf("save image to xpixmap failed, %s", destFile)
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

func GetPixels(pixbuf *C.GdkPixbuf) []byte {
	var cLength C.guint
	ptr := C.gdk_pixbuf_get_pixels_with_length(pixbuf, &cLength)
	return C.GoBytes(unsafe.Pointer(ptr), C.int(cLength))
}

func ToImage(pixbuf *C.GdkPixbuf) (image.Image, error) {
	width, height, err := GetSize(pixbuf)
	if err != nil {
		return nil, err
	}
	colorspace := C.gdk_pixbuf_get_colorspace(pixbuf)
	if colorspace != C.GDK_COLORSPACE_RGB {
		return nil, errors.New("colorspace is not RGB")
	}
	nChannels := C.gdk_pixbuf_get_n_channels(pixbuf)
	if nChannels != 4 {
		return nil, errors.New("n channels != 4")
	}
	bitsPerSample := C.gdk_pixbuf_get_bits_per_sample(pixbuf)
	if bitsPerSample != 8 {
		return nil, errors.New("bits per sample != 8")
	}
	stride := C.gdk_pixbuf_get_rowstride(pixbuf)

	pixels := GetPixels(pixbuf)

	img := &image.RGBA{
		Pix:    pixels,
		Stride: int(stride),
		Rect:   image.Rect(0, 0, width, height),
	}
	return img, nil
}

// Info

// GetImageSize return image file's width and height.
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

// GetSize return pixbuf's width and height.
func GetSize(pixbuf *C.GdkPixbuf) (w, h int, err error) {
	defaultError := fmt.Errorf("get GdkPixbuf size failed, %v", pixbuf)
	if pixbuf == nil {
		err = defaultError
		return
	}
	w = int(C.gdk_pixbuf_get_width(pixbuf))
	h = int(C.gdk_pixbuf_get_height(pixbuf))
	return
}

// GetImageFormat return image file's format, such as "png", "jpeg".
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
	return err == nil
}

// Clip

func ClipImage(srcFile, destFile string, srcX, srcY, width, height int, f Format) (err error) {
	srcPixbuf, err := NewPixbufFromFile(srcFile)
	defer FreePixbuf(srcPixbuf)
	if err != nil {
		return
	}
	destPixbuf, err := CopyAreaSimple(srcPixbuf, srcX, srcY, width, height)
	defer FreePixbuf(destPixbuf)
	if err != nil {
		return
	}
	err = Save(destPixbuf, destFile, f)
	return
}

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

// Dominant color

// GetDominantColorOfImage return the dominant hsv color of an image.
func GetDominantColorOfImage(imgFile string) (h, s, v float64, err error) {
	pixbuf, err := NewPixbufFromFile(imgFile)
	defer FreePixbuf(pixbuf)
	if err != nil {
		return
	}
	return GetDominantColor(pixbuf)
}

func GetDominantColor(pixbuf *C.GdkPixbuf) (h, s, v float64, err error) {
	defaultError := fmt.Errorf("get dominant color of pixbuf failed, %v", pixbuf)
	var r, g, b float64
	ret := C.get_dominant_color(pixbuf, (*C.double)(&r), (*C.double)(&g), (*C.double)(&b))
	if ret == 0 {
		err = defaultError
		return
	}
	h, s, v = Rgb2Hsv(uint8(r), uint8(g), uint8(b))
	return
}

// Convert

func ConvertImage(srcFile, destFile string, f Format) (err error) {
	srcPixbuf, err := NewPixbufFromFile(srcFile)
	defer FreePixbuf(srcPixbuf)
	if err != nil {
		return
	}
	err = Save(srcPixbuf, destFile, f)
	return
}

// Flip

func FlipImageHorizontal(srcFile, destFile string, f Format) (err error) {
	return doFlipImage(srcFile, destFile, true, f)
}

func FlipImageVertical(srcFile, destFile string, f Format) (err error) {
	return doFlipImage(srcFile, destFile, false, f)
}

func doFlipImage(srcFile, destFile string, horizontal bool, f Format) (err error) {
	srcPixbuf, err := NewPixbufFromFile(srcFile)
	defer FreePixbuf(srcPixbuf)
	if err != nil {
		return
	}
	destPixbuf, err := Flip(srcPixbuf, horizontal)
	defer FreePixbuf(destPixbuf)
	if err != nil {
		return
	}
	err = Save(destPixbuf, destFile, f)
	return
}

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

// Scale and Thumbnail

// ScaleImage returns a new image file with the given width and
// height created by resizing the given image.
func ScaleImage(srcFile, destFile string, newWidth, newHeght int, interpType GdkInterpType, f Format) (err error) {
	srcPixbuf, err := NewPixbufFromFile(srcFile)
	defer FreePixbuf(srcPixbuf)
	if err != nil {
		return
	}
	destPixbuf, err := ScaleSimple(srcPixbuf, newWidth, newHeght, interpType)
	defer FreePixbuf(destPixbuf)
	if err != nil {
		return
	}
	err = Save(destPixbuf, destFile, f)
	return
}

// ScaleImagePrefer resize image file to new width and heigh, and
// maintain the original proportions unchanged.
func ScaleImagePrefer(srcFile, destFile string, newWidth, newHeght int, interpType GdkInterpType, f Format) (err error) {
	srcPixbuf, err := NewPixbufFromFile(srcFile)
	defer FreePixbuf(srcPixbuf)
	if err != nil {
		return
	}
	destPixbuf, err := ScalePrefer(srcPixbuf, newWidth, newHeght, interpType)
	defer FreePixbuf(destPixbuf)
	if err != nil {
		return
	}
	err = Save(destPixbuf, destFile, f)
	return
}

// ThumbnailImage resize target image file with limited maximum width and height.
func ThumbnailImage(srcFile, destFile string, maxWidth, maxHeight int, interpType GdkInterpType, f Format) (err error) {
	srcPixbuf, err := NewPixbufFromFile(srcFile)
	defer FreePixbuf(srcPixbuf)
	if err != nil {
		return
	}
	destPixbuf, err := Thumbnail(srcPixbuf, maxWidth, maxHeight, interpType)
	defer FreePixbuf(destPixbuf)
	if err != nil {
		return
	}
	err = Save(destPixbuf, destFile, f)
	return
}

func ScaleImageCache(srcFile string, newWidth, newHeght int, interpType GdkInterpType, f Format) (destFile string, useCache bool, err error) {
	destFile = generateCacheFilePath(fmt.Sprintf("ScaleImageCache%s%d%d%d%s", srcFile, newWidth, newHeght, interpType, f))
	if utils.IsFileExist(destFile) {
		// return cache file
		useCache = true
		return
	}
	err = ScaleImage(srcFile, destFile, newWidth, newHeght, interpType, f)
	return
}

func ScaleImagePreferCache(srcFile string, newWidth, newHeght int, interpType GdkInterpType, f Format) (destFile string, useCache bool, err error) {
	destFile = generateCacheFilePath(fmt.Sprintf("ScaleImageCache%s%d%d%d%s", srcFile, newWidth, newHeght, interpType, f))
	if utils.IsFileExist(destFile) {
		// return cache file
		useCache = true
		return
	}
	err = ScaleImagePrefer(srcFile, destFile, newWidth, newHeght, interpType, f)
	return
}

func ScaleSimple(srcPixbuf *C.GdkPixbuf, newWidth, newHeght int, interpType GdkInterpType) (destPixbuf *C.GdkPixbuf, err error) {
	defaultError := fmt.Errorf("scale pixbuf failed, %v, %v, %v, %v", srcPixbuf, newWidth, newHeght, interpType)
	destPixbuf = C.gdk_pixbuf_scale_simple(srcPixbuf, C.int(newWidth), C.int(newHeght), C.GdkInterpType(interpType))
	if destPixbuf == nil {
		err = defaultError
		return
	}
	return
}

// ScalePrefer resize pixbuf to new width and heigh, and maintain the
// original proportions unchanged.
func ScalePrefer(srcPixbuf *C.GdkPixbuf, newWidth, newHeight int, interpType GdkInterpType) (destPixbuf *C.GdkPixbuf, err error) {
	iw, ih, err := GetSize(srcPixbuf)
	if err != nil {
		return
	}
	x, y, w, h, err := GetPreferScaleClipRect(newWidth, newHeight, iw, ih)
	if err != nil {
		return
	}
	clipPixbuf, err := CopyAreaSimple(srcPixbuf, x, y, w, h)
	defer FreePixbuf(clipPixbuf)
	if err != nil {
		return
	}
	destPixbuf, err = ScaleSimple(clipPixbuf, newWidth, newHeight, interpType)
	return
}

// Thumbnail resize pixbuf with limited maximum width and height.
func Thumbnail(srcPixbuf *C.GdkPixbuf, maxWidth, maxHeight int, interpType GdkInterpType) (destPixbuf *C.GdkPixbuf, err error) {
	// get new width and heigh
	var newWidth, newHeight int
	w, h, err := GetSize(srcPixbuf)
	if err != nil {
		return
	}
	scale := float32(w) / float32(h)
	newWidth = maxWidth
	newHeight = int(float32(newWidth) / scale)
	if newHeight > maxHeight {
		newHeight = maxHeight
		newWidth = int(float32(newHeight) * scale)
	}
	return ScaleSimple(srcPixbuf, newWidth, newHeight, interpType)
}

// Rotate

func RotateImageLeft(srcFile, destFile string, f Format) (err error) {
	err = doRotateImage(srcFile, destFile, GDK_PIXBUF_ROTATE_COUNTERCLOCKWISE, f)
	return
}

func RotateImageRight(srcFile, destFile string, f Format) (err error) {
	err = doRotateImage(srcFile, destFile, GDK_PIXBUF_ROTATE_CLOCKWISE, f)
	return
}

func RotateImageUpsizedown(srcFile, destFile string, f Format) (err error) {
	err = doRotateImage(srcFile, destFile, GDK_PIXBUF_ROTATE_UPSIDEDOWN, f)
	return
}

func doRotateImage(srcFile, destFile string, angle GdkPixbufRotation, f Format) (err error) {
	srcPixbuf, err := NewPixbufFromFile(srcFile)
	defer FreePixbuf(srcPixbuf)
	if err != nil {
		return
	}
	destPixbuf, err := RotateSimple(srcPixbuf, angle)
	defer FreePixbuf(destPixbuf)
	if err != nil {
		return
	}
	err = Save(destPixbuf, destFile, f)
	return
}

func RotateSimple(srcPixbuf *C.GdkPixbuf, angle GdkPixbufRotation) (destPixbuf *C.GdkPixbuf, err error) {
	defaultError := fmt.Errorf("rotate pixbuf failed, %v, %v", srcPixbuf, angle)
	destPixbuf = C.gdk_pixbuf_rotate_simple(srcPixbuf, C.GdkPixbufRotation(angle))
	if destPixbuf == nil {
		err = defaultError
		return
	}
	return
}

// XLib

// ConvertImageToXpixmap convert image file to x pixmap.
func ConvertImageToXpixmap(imgFile string) (xpixmap x.Pixmap, err error) {
	pixbuf, err := NewPixbufFromFile(imgFile)
	defer FreePixbuf(pixbuf)
	if err != nil {
		return
	}
	xpixmap, err = ConvertPixbufToXpixmap(pixbuf)
	return
}

func ConvertPixbufToXpixmap(pixbuf *C.GdkPixbuf) (xpixmap x.Pixmap, err error) {
	defaultError := fmt.Errorf("convert pixbuf to xpixmap failed, %v", pixbuf)
	xpixmap = x.Pixmap(C.convert_pixbuf_to_xpixmap(pixbuf))
	if xpixmap == 0 {
		err = defaultError
		return
	}
	return
}

func ConvertXpixmapToPixbuf(xpixmap x.Pixmap, width, height int) (pixbuf *C.GdkPixbuf, err error) {
	defaultError := fmt.Errorf("convert xpixmap to pixbuf failed, %v", xpixmap)
	pixbuf = C.convert_xpixmap_to_pixbuf(C.Pixmap(xpixmap), C.int(width), C.int(height))
	if pixbuf == nil {
		err = defaultError
		return
	}
	return
}

func ScreenshotImage(file string, f Format) (err error) {
	pixbuf, err := Screenshot()
	defer FreePixbuf(pixbuf)
	if err != nil {
		return
	}
	err = Save(pixbuf, file, f)
	return
}
func Screenshot() (pixbuf *C.GdkPixbuf, err error) {
	defaultError := fmt.Errorf("take a screenshot failed")
	pixbuf = C.screenshot()
	if pixbuf == nil {
		err = defaultError
		return
	}
	return
}
