/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package gdkpixbuf

import (
	"pkg.deepin.io/lib/graphic"
	"pkg.deepin.io/lib/utils"
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
