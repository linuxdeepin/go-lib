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

import (
	"fmt"
	"image"
	"image/draw"
)

// FillStyle define the type to fill image.
type FillStyle string

const (
	FillTile                  FillStyle = "tile"                    // 平铺
	FillCenter                          = "center"                  // 居中
	FillScale                           = "scale"                   // 拉伸
	FillProportionCenterScale           = "proportion-center-scale" // 等比居中拉伸
)

// FillImage generate a new image in target width and height through
// source image, there are many fill sytles to choice from.
func FillImage(srcfile, dstfile string, width, height int, style FillStyle, f Format) (err error) {
	srcimg, err := LoadImage(srcfile)
	if err != nil {
		return
	}
	dstimg, err := ImplFillImage(srcimg, width, height, style)
	if err != nil {
		return
	}
	err = SaveImage(dstfile, dstimg, f)
	dstimg.Pix = nil
	return
}

// FillImageCache generate a new image in target width and height through
// source image, and save it to cache directory, if already exists,
// just return it.
func FillImageCache(srcfile string, width, height int, style FillStyle, f Format) (dstfile string, useCache bool, err error) {
	dstfile = GenerateCacheFilePath(fmt.Sprintf("FillImageCache%s%d%d%s%s", srcfile, width, height, style, f))
	if isFileExists(dstfile) {
		// return cache file
		useCache = true
		return
	}
	err = FillImage(srcfile, dstfile, width, height, style, f)
	return
}

func ImplFillImage(srcimg image.Image, width, height int, style FillStyle) (dstimg *image.RGBA, err error) {
	switch style {
	case FillTile:
		dstimg = doFillImageInTileStyle(srcimg, width, height, style)
	case FillCenter:
		dstimg = doFillImageInCenterStyle(srcimg, width, height, style)
	case FillScale:
		dstimg = doFillImageInScaleStyle(srcimg, width, height, style)
	case FillProportionCenterScale:
		dstimg, err = doFillImageInProportionCenterScaleStyle(srcimg, width, height, style)
	default:
		// default to use FilleTile style
		dstimg = doFillImageInTileStyle(srcimg, width, height, style)
	}
	return
}

func doFillImageInTileStyle(srcimg image.Image, width, height int, style FillStyle) (dstimg *image.RGBA) {
	dstimg = image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	iw, ih := doGetImageSize(srcimg)

	endx := width - 1
	endy := height - 1
	for x := 0; x <= endx; x += iw {
		for y := 0; y <= endy; y += ih {
			draw.Draw(dstimg, image.Rect(x, y, x+iw, y+ih), srcimg, image.Point{0, 0}, draw.Src)
		}
	}

	return
}

func doFillImageInCenterStyle(srcimg image.Image, width, height int, style FillStyle) (dstimg *image.RGBA) {
	dstimg = image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	iw, ih := doGetImageSize(srcimg)

	var srcX, srcY, dstX, dstY, clipWidth, clipHeight int
	if width > iw {
		srcX = 0
		clipWidth = iw
		dstX = width/2 - iw/2
	} else {
		dstX = 0
		clipWidth = width
		srcX = iw/2 - width/2
	}
	if height > ih {
		srcY = 0
		clipHeight = ih
		dstY = height/2 - ih/2
	} else {
		dstY = 0
		clipHeight = height
		srcY = ih/2 - height/2
	}

	draw.Draw(dstimg, image.Rect(dstX, dstY, dstX+clipWidth, dstY+clipHeight), srcimg, image.Point{srcX, srcY}, draw.Src)
	return
}

func doFillImageInScaleStyle(srcimg image.Image, width, height int, style FillStyle) (dstimg *image.RGBA) {
	dstimg = doResizeNearestNeighbor(srcimg, width, height)
	return
}

func doFillImageInProportionCenterScaleStyle(srcimg image.Image, width, height int, style FillStyle) (dstimg *image.RGBA, err error) {
	iw, ih := doGetImageSize(srcimg)
	x, y, w, h, err := GetProportionCenterScaleRect(width, height, iw, ih)
	if err != nil {
		return
	}
	dstimg = ImplClipImage(srcimg, x, y, w, h)
	dstimg = doResizeNearestNeighbor(dstimg, width, height)
	return
}

// GetProportionCenterScaleRect get rectangle in image which with the same
// scale to reference width/heigh, and the rectangle will placed in
// center.
func GetProportionCenterScaleRect(refWidth, refHeight, imgWidth, imgHeight int) (x, y, w, h int, err error) {
	if refWidth*refHeight == 0 || imgWidth*imgHeight == 0 {
		err = fmt.Errorf("argument is invalid: ", refWidth, refHeight, imgWidth, imgHeight)
		return
	}
	scale := float32(refWidth) / float32(refHeight)
	w = imgWidth
	h = int(float32(w) / scale)
	if h < imgHeight {
		offsetY := (imgHeight - h) / 2
		x = 0
		y = 0 + offsetY
	} else {
		h = imgHeight
		w = int(float32(h) * scale)
		offsetX := (imgWidth - w) / 2
		x = 0 + offsetX
		y = 0
	}
	return
}
