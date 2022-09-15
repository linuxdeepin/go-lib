// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package gdkpixbuf

import (
	"github.com/linuxdeepin/go-lib/graphic"
	"github.com/linuxdeepin/go-lib/utils"
)

// function links to lib/graphic

func generateCacheFilePath(keyword string) string {
	return utils.GenerateCacheFilePathWithPrefix("graphic", keyword)
}

// Rgb2Hsv convert color format from RGB(r, g, b=[0..255]) to HSV(h=[0..360), s,v=[0..1]).
func Rgb2Hsv(r, g, b uint8) (h, s, v float64) {
	return graphic.Rgb2Hsv(r, g, b)
}

// Hsv2Rgb convert color format from HSV(h=[0..360), s,v=[0..1]) to RGB(r, g, b=[0..255]).
func Hsv2Rgb(h, s, v float64) (r, g, b uint8) {
	return graphic.Hsv2Rgb(h, s, v)
}

// GetPreferScaleClipRect get the maximum rectangle in center of
// image which with the same scale to reference width/heigh.
func GetPreferScaleClipRect(refWidth, refHeight, imgWidth, imgHeight int) (x, y, w, h int, err error) {
	x, y, w, h, err = graphic.GetPreferScaleClipRect(refWidth, refHeight, imgWidth, imgHeight)
	return
}
