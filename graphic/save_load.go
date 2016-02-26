/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package graphic

import (
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
	default:
		err = png.Encode(w, m)
	}
	return
}
