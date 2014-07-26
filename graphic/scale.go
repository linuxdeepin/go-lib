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
)

// ScaleImage returns a new image file with the given width and
// height created by resizing the given image.
func ScaleImage(srcfile, dstfile string, newWidth, newHeight int, f Format) (err error) {
	srcimg, err := LoadImage(srcfile)
	if err != nil {
		return
	}
	dstimg := Scale(srcimg, newWidth, newHeight)
	err = SaveImage(dstfile, dstimg, f)
	dstimg.Pix = nil
	return
}

// ScaleImagePrefer resize image file to new width and heigh, and
// maintain the original proportions unchanged.
func ScaleImagePrefer(srcfile, dstfile string, newWidth, newHeight int, f Format) (err error) {
	srcimg, err := LoadImage(srcfile)
	if err != nil {
		return
	}
	dstimg, err := ScalePrefer(srcimg, newWidth, newHeight)
	if err != nil {
		return
	}
	err = SaveImage(dstfile, dstimg, f)
	dstimg.Pix = nil
	return
}

// ScaleImageCache resize any recognized format image file and save to cache
// directory, if already exists, just return it.
func ScaleImageCache(srcfile string, newWidth, newHeight int, f Format) (dstfile string, useCache bool, err error) {
	dstfile = generateCacheFilePath(fmt.Sprintf("ScaleImageCache%s%d%d%s", srcfile, newWidth, newHeight, f))
	if isFileExists(dstfile) {
		useCache = true
		return
	}
	err = ScaleImage(srcfile, dstfile, newWidth, newHeight, f)
	return
}

// ThumbnailImage resize target image file with limited maximum width and height.
func ThumbnailImage(srcfile, dstfile string, maxWidth, maxHeight int, f Format) (err error) {
	srcimg, err := LoadImage(srcfile)
	if err != nil {
		return
	}
	dstimg := Thumbnail(srcimg, maxWidth, maxHeight)
	err = SaveImage(dstfile, dstimg, f)
	dstimg.Pix = nil
	return
}

// ThumbnailImageCache resize target image file with limited maximum width
// and height, and save to cache directory, if already exists, just
// return it.
func ThumbnailImageCache(srcfile string, maxWidth, maxHeight int, f Format) (dstfile string, useCache bool, err error) {
	dstfile = generateCacheFilePath(fmt.Sprintf("ThumbnailImageCache%s%d%d%s", srcfile, maxWidth, maxHeight, f))
	if isFileExists(dstfile) {
		useCache = true
		return
	}
	err = ThumbnailImage(srcfile, dstfile, maxWidth, maxHeight, f)
	return
}

// Scale resize image object to new width and height.
func Scale(srcimg image.Image, newWidth, newHeight int) (dstimg *image.RGBA) {
	dstimg = doScaleNearestNeighbor(srcimg, newWidth, newHeight)
	return
}

// Thumbnail resize image object with limited maximum width and height.
func Thumbnail(srcimg image.Image, maxWidth, maxHeight int) (dstimg *image.RGBA) {
	// get new width and heigh
	var newWidth, newHeight int
	w, h := GetSize(srcimg)
	scale := float32(w) / float32(h)
	newWidth = maxWidth
	newHeight = int(float32(newWidth) / scale)
	if newHeight > maxHeight {
		newHeight = maxHeight
		newWidth = int(float32(newHeight) * scale)
	}
	return Scale(srcimg, newWidth, newHeight)
}

// ScalePrefer resize image object to new width and heigh, and
// maintain the original proportions unchanged.
func ScalePrefer(srcimg image.Image, newWidth, newHeight int) (dstimg *image.RGBA, err error) {
	iw, ih := GetSize(srcimg)
	x, y, w, h, err := GetPreferScaleClipRect(newWidth, newHeight, iw, ih)
	if err != nil {
		return
	}
	dstimg = Clip(srcimg, x, y, w, h)
	dstimg = Scale(dstimg, newWidth, newHeight)
	return
}

// TODO doScaleNearestNeighbor returns a new RGBA image with the given width and
// height created by resizing the given image using the nearest neighbor
// algorithm.
func doScaleNearestNeighbor(img image.Image, newWidth, newHeight int) (newimg *image.RGBA) {
	w := img.Bounds().Max.X
	h := img.Bounds().Max.Y
	newimg = image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	xr := (w<<16)/newWidth + 1
	yr := (h<<16)/newHeight + 1

	for yo := 0; yo < newHeight; yo++ {
		y2 := (yo * yr) >> 16
		for xo := 0; xo < newWidth; xo++ {
			x2 := (xo * xr) >> 16
			newimg.Set(xo, yo, img.At(x2, y2))
			//Much faster, but requires some image type.
			//newimg.Pix[offset] = img.Pix[y2*w+x2]
			//offset++
		}
	}
	return newimg
}

// GetPreferScaleClipRect get the maximum rectangle in center of
// image which with the same scale to reference width/heigh.
func GetPreferScaleClipRect(refWidth, refHeight, imgWidth, imgHeight int) (x, y, w, h int, err error) {
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
