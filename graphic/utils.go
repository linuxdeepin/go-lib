// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package graphic

import (
	"image"
	"image/draw"
	"os"
	"github.com/linuxdeepin/go-lib/utils"
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

// convert image.Image to *image.RGBA
func convertToRGBA(img image.Image) (rgba *image.RGBA) {
	b := img.Bounds()
	r := image.Rect(0, 0, b.Dx(), b.Dy())
	rgba = image.NewRGBA(r)
	draw.Draw(rgba, rgba.Bounds(), img, b.Min, draw.Src)
	return
}
