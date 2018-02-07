/*
 * Copyright (C) 2017 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package notify

import (
	"image"
	"image/draw"
)

// This is a raw data image format which describes the width, height, rowstride,
// has alpha, bits per sample, channels and image data respectively.
type Image struct {
	Width         int32
	Height        int32
	RowStride     int32
	HasAlpha      bool
	BitsPerSample int32
	Channels      int32
	Pix           []byte
}

func (n *Notification) SetImage(img image.Image) {
	n.SetHint(HintImageData, NewImage(img))
}

func toRGBA(src image.Image) *image.RGBA {
	b := src.Bounds()
	m := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(m, m.Bounds(), src, b.Min, draw.Src)
	return m
}

func NewImage(src image.Image) *Image {
	return newImageFromRGBA(toRGBA(src))
}

func newImageFromRGBA(img *image.RGBA) *Image {
	opaque := img.Opaque()
	b := img.Bounds()
	w := b.Dx()
	h := b.Dy()

	nchan := 4
	if opaque {
		nchan = 3
	}
	stride := w * nchan
	pix := make([]byte, stride*h)

	// fill data to pix
	// img.Pix idx
	i := 0
	// pix idx
	p := 0

	if opaque {
		for y := 0; y < h; y++ {
			// copy one row
			for x := 0; x < w; x++ {
				pix[p] = img.Pix[i]     // R
				pix[p+1] = img.Pix[i+1] // G
				pix[p+2] = img.Pix[i+2] // B
				// skip A

				p += nchan
				i += (nchan + 1)
			}
		}
	} else {
		for y := 0; y < h; y++ {
			// copy one row
			copy(pix[p:], img.Pix[i:i+stride])
			p += stride
			i += img.Stride
		}
	}

	return &Image{
		Width:         int32(w),
		Height:        int32(h),
		RowStride:     int32(stride),
		HasAlpha:      !opaque,
		BitsPerSample: 8,
		Channels:      int32(nchan),
		Pix:           pix,
	}
}
