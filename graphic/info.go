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
	"os"
	dutils "pkg.deepin.io/lib/utils"
)

// GetImageSize return image's width and height.
func GetImageSize(imgfile string) (w, h int, err error) {
	f, err := os.Open(imgfile)
	if err != nil {
		return
	}
	defer f.Close()
	config, _, err := image.DecodeConfig(f)
	return config.Width, config.Height, err
}

func GetSize(img image.Image) (w, h int) {
	w = img.Bounds().Dx()
	h = img.Bounds().Dy()
	return
}

// GetImageFormat return image format, such as "png", "jpeg".
func GetImageFormat(imgfile string) (format Format, err error) {
	f, err := os.Open(imgfile)
	if err != nil {
		return
	}
	defer f.Close()
	_, name, err := image.DecodeConfig(f)
	format = Format(name)
	return
}

// IsSupportedImage check if image file is supported.
func IsSupportedImage(imgfile string) bool {
	f, err := os.Open(imgfile)
	if err != nil {
		return false
	}
	defer f.Close()
	_, _, err = image.DecodeConfig(f)
	if err != nil {
		return false
	}
	return true
}

func GetImagesInDir(dir string) ([]string, error) {
	files, err := dutils.GetFilesInDir(dir)
	if err != nil {
		return nil, err
	}

	var images []string
	for _, file := range files {
		if !IsSupportedImage(file) {
			continue
		}

		images = append(images, file)
	}

	return images, nil
}
