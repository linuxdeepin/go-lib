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
	"image/color"
)

// NewImageWithColor create a new image file with target size and rgba.
func NewImageWithColor(dstfile string, width, height int, r, g, b, a uint8, f Format) (err error) {
	dstimg, err := NewWithColor(width, height, r, g, b, a)
	if err != nil {
		return
	}
	err = SaveImage(dstfile, dstimg, f)
	dstimg.Pix = nil
	return
}

// NewWithColor create a new image object with target size and rgba.
func NewWithColor(width, height int, r, g, b, a uint8) (dstimg *image.RGBA, err error) {
	dstimg = image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	color := color.RGBA{R: r, G: g, B: b, A: a}
	w, h := GetSize(dstimg)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			dstimg.SetRGBA(i, j, color)
		}
	}
	return
}
