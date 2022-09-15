// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package graphic

import (
	"image"
)

// RotateImageLeft rotate image to left side.
func RotateImageLeft(srcfile, dstfile string, f Format) (err error) {
	srcimg, err := LoadImage(srcfile)
	if err != nil {
		return err
	}
	dstimg := doRotateImageLeft(srcimg)
	err = SaveImage(dstfile, dstimg, f)
	dstimg.Pix = nil
	return
}

// RotateImageLeft rotate image to right side.
func RotateImageRight(srcfile, dstfile string, f Format) (err error) {
	srcimg, err := LoadImage(srcfile)
	if err != nil {
		return err
	}
	dstimg := doRotateImageRight(srcimg)
	err = SaveImage(dstfile, dstimg, f)
	dstimg.Pix = nil
	return
}

func doRotateImageLeft(srcimg image.Image) (dstimg *image.RGBA) {
	w, h := GetSize(srcimg)
	dstimg = image.NewRGBA(image.Rect(0, 0, h, w))

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			dstimg.Set(y, w-x-1, srcimg.At(x, y))
		}
	}

	return
}

func doRotateImageRight(srcimg image.Image) (dstimg *image.RGBA) {
	w, h := GetSize(srcimg)
	dstimg = image.NewRGBA(image.Rect(0, 0, h, w))

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			dstimg.Set(h-y-1, x, srcimg.At(x, y))
		}
	}

	return
}
