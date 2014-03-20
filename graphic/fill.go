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
	"image"
	"image/draw"
)

// FillStyle define the type to fill image.
type FillStyle string

const (
	FillTile         FillStyle = "tile"         // 平铺
	FillCenter                 = "center"       // 居中
	FillStretch                = "stretch"      // 拉伸
	FillScaleStretch           = "scalestretch" // 等比拉伸
)

// FillImage generate a new image in target width and height through
// source image, there are many fill sytles to choice from.
func FillImage(srcfile, dstfile string, width, height int32, style FillStyle, f Format) (err error) {
	srcimg, err := LoadImage(srcfile)
	if err != nil {
		return
	}
	dstimg := ImplFillImage(srcimg, int(width), int(height), style)
	return SaveImage(dstfile, dstimg, f)
}

// FIXME return draw.Image or *image.RGBA
func ImplFillImage(srcimg image.Image, width, height int, style FillStyle) (dstimg draw.Image) {
	switch style {
	case FillTile:
		dstimg = doFillImageInTileStyle(srcimg, width, height, style)
	case FillCenter:
		dstimg = doFillImageInCenterStyle(srcimg, width, height, style)
	case FillStretch:
		dstimg = doFillImageInStretchStyle(srcimg, width, height, style)
	case FillScaleStretch:
		dstimg = doFillImageInScaleStretchStyle(srcimg, width, height, style)
	default:
		// default to use FilleTile style
		dstimg = doFillImageInTileStyle(srcimg, width, height, style)
	}
	return
}

func doFillImageInTileStyle(srcimg image.Image, width, height int, style FillStyle) (dstimg draw.Image) {
	dstimg = image.NewRGBA(image.Rect(0, 0, width, height))
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

func doFillImageInCenterStyle(srcimg image.Image, width, height int, style FillStyle) (dstimg draw.Image) {
	dstimg = image.NewRGBA(image.Rect(0, 0, width, height))
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

func doFillImageInStretchStyle(srcimg image.Image, width, height int, style FillStyle) (dstimg draw.Image) {
	dstimg = doResizeNearestNeighbor(srcimg, width, height)
	return
}

func doFillImageInScaleStretchStyle(srcimg image.Image, width, height int, style FillStyle) (dstimg draw.Image) {
	iw, ih := doGetImageSize(srcimg)
	x0, y0, x1, y1 := GetScaleRectInImage(width, height, iw, ih)
	dstimg = ImplClipImage(srcimg, x0, y0, x1, y1)
	dstimg = doResizeNearestNeighbor(dstimg, width, height)
	return
}

// GetScaleRectInImage get rectangle in image which with the same
// scale to reference width/heigh, and the rectangle will placed in
// center.
func GetScaleRectInImage(refWidth, refHeight, imgWidth, imgHeight int) (x0, y0, x1, y1 int) {
	scale := float32(refWidth) / float32(refHeight)
	w := imgWidth
	h := int(float32(w) / scale)
	if h < imgHeight {
		offsetY := (imgHeight - h) / 2
		x0 = 0
		y0 = 0 + offsetY
		x1 = x0 + w
		y1 = y0 + h
	} else {
		h = imgHeight
		w = int(float32(h) * scale)
		offsetX := (imgWidth - w) / 2
		x0 = 0 + offsetX
		y0 = 0
		x1 = x0 + w
		y1 = y0 + h
	}
	return
}
