// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package graphic

import (
	"fmt"
	"image"
	"image/draw"
)

// TODO: is FillXXX() still necessary?

// FillStyle define the type to fill image.
type FillStyle string

const (
	FillTile   FillStyle = "tile"   // 平铺
	FillCenter FillStyle = "center" // 居中
)

// FillImage generate a new image file in target width and height through
// source image file, there are many fill sytles to choice from.
func FillImage(srcfile, dstfile string, width, height int, style FillStyle, f Format) (err error) {
	srcimg, err := LoadImage(srcfile)
	if err != nil {
		return
	}
	dstimg, err := Fill(srcimg, width, height, style)
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
	dstfile = generateCacheFilePath(fmt.Sprintf("FillImageCache%s%d%d%s%s", srcfile, width, height, style, f))
	if isFileExists(dstfile) {
		// return cache file
		useCache = true
		return
	}
	err = FillImage(srcfile, dstfile, width, height, style, f)
	return
}

// FillImage generate a new image in target width and height through
// source image, there are many fill sytles to choice from.
func Fill(srcimg image.Image, width, height int, style FillStyle) (dstimg *image.RGBA, err error) {
	switch style {
	case FillTile:
		dstimg = doFillImageInTileStyle(srcimg, width, height)
	case FillCenter:
		dstimg = doFillImageInCenterStyle(srcimg, width, height)
	default:
		err = fmt.Errorf("unknown fill style %v", style)
		return
	}
	return
}

func doFillImageInTileStyle(srcimg image.Image, width, height int) (dstimg *image.RGBA) {
	dstimg = image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	iw, ih := GetSize(srcimg)

	endx := width - 1
	endy := height - 1
	for x := 0; x <= endx; x += iw {
		for y := 0; y <= endy; y += ih {
			draw.Draw(dstimg, image.Rect(x, y, x+iw, y+ih), srcimg, image.Point{0, 0}, draw.Src)
		}
	}

	return
}

func doFillImageInCenterStyle(srcimg image.Image, width, height int) (dstimg *image.RGBA) {
	dstimg = image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	iw, ih := GetSize(srcimg)

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
