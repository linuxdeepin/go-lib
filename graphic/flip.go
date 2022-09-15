// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package graphic

import (
	"image"
)

// FlipImageHorizontal flip image in horizontal direction, and save as
// target format.
func FlipImageHorizontal(srcfile, dstfile string, f Format) (err error) {
	srcimg, err := LoadImage(srcfile)
	if err != nil {
		return err
	}
	dstimg := doFlipImageHorizontal(srcimg)
	err = SaveImage(dstfile, dstimg, f)
	dstimg.Pix = nil
	return
}

// FlipImageVertical  flip image  in vertical  direction, and  save as
// target format.
func FlipImageVertical(srcfile, dstfile string, f Format) (err error) {
	srcimg, err := LoadImage(srcfile)
	if err != nil {
		return err
	}
	dstimg := doFlipImageVertical(srcimg)
	err = SaveImage(dstfile, dstimg, f)
	dstimg.Pix = nil
	return
}

func doFlipImageHorizontal(srcimg image.Image) (dstimg *image.RGBA) {
	w, h := GetSize(srcimg)
	dstimg = image.NewRGBA(image.Rect(0, 0, w, h))

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			dstimg.Set(w-x-1, y, srcimg.At(x, y))
		}
	}

	return
}

func doFlipImageVertical(srcimg image.Image) (dstimg *image.RGBA) {
	w, h := GetSize(srcimg)
	dstimg = image.NewRGBA(image.Rect(0, 0, w, h))

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			dstimg.Set(x, h-y-1, srcimg.At(x, y))
		}
	}

	return
}
