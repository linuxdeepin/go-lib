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
	"image/draw"
	"os"
	"pkg.deepin.io/lib/utils"
)

func generateCacheFilePath(keyword string) (dstfile string) {
	return utils.GenerateCacheFilePathWithPrefix("graphic", keyword)
}

func openFileOrCreate(file string) (*os.File, error) {
	return os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0644)
}

func isFileExists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	} else {
		return false
	}
}

func ensureDirExists(dir string) {
	os.MkdirAll(dir, 0755)
}

// convert image.Image to *image.RGBA
func convertToRGBA(img image.Image) (rgba *image.RGBA) {
	b := img.Bounds()
	r := image.Rect(0, 0, b.Dx(), b.Dy())
	rgba = image.NewRGBA(r)
	draw.Draw(rgba, rgba.Bounds(), img, b.Min, draw.Src)
	return
}
