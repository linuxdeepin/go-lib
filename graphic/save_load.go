// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package graphic

import (
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

// LoadImage load image file and return image.Image object.
func LoadImage(imgfile string) (img image.Image, err error) {
	f, err := os.Open(imgfile)
	if err != nil {
		return
	}
	defer f.Close()
	img, _, err = image.Decode(f)
	return
}

// SaveImage save image.Image object to target file.
func SaveImage(dstfile string, m image.Image, f Format) (err error) {
	df, err := openFileOrCreate(dstfile)
	if err != nil {
		return
	}
	defer df.Close()
	return doSaveImage(df, m, f)
}

func doSaveImage(w io.Writer, m image.Image, f Format) (err error) {
	switch f {
	case FormatPng:
		err = png.Encode(w, m)
	case FormatJpeg:
		err = jpeg.Encode(w, m, nil)
	case FormatBmp:
		err = bmp.Encode(w, m)
	case FormatTiff:
		err = tiff.Encode(w, m, nil)
	default:
		err = png.Encode(w, m)
	}
	return
}
