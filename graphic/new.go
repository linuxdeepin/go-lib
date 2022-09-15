// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

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
