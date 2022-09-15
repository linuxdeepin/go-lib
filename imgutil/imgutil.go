// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package imgutil

import (
	"bufio"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"os"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	"github.com/linuxdeepin/go-lib/gdkpixbuf"
	"github.com/linuxdeepin/go-lib/strv"
)

const (
	FormatGIF  = "gif"
	FormatJPEG = "jpeg"
	FormatPNG  = "png"
	FormatBMP  = "bmp"
	FormatTIFF = "tiff"
)

var supportedFormats = strv.Strv([]string{FormatGIF, FormatJPEG, FormatPNG,
	FormatBMP, FormatTIFF})

func GetSupportedFormats() strv.Strv {
	return supportedFormats
}

func Load(filename string) (image.Image, error) {
	img, err := loadViaGdkPixbuf(filename)
	if err == nil {
		return img, nil
	}
	return loadCommon(filename)
}

func loadCommon(filename string) (image.Image, error) {
	fh, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	br := bufio.NewReader(fh)
	img, _, err := image.Decode(br)
	return img, err
}

func loadViaGdkPixbuf(filename string) (image.Image, error) {
	pb, err := gdkpixbuf.NewPixbufFromFile(filename)
	if err != nil {
		return nil, err
	}
	img, err := gdkpixbuf.ToImage(pb)
	gdkpixbuf.FreePixbuf(pb)
	return img, err
}

func SavePng(img image.Image, filename string, enc *png.Encoder) error {
	fh, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fh.Close()
	bw := bufio.NewWriter(fh)

	if enc == nil {
		enc = &png.Encoder{}
	}

	err = enc.Encode(bw, img)
	if err != nil {
		return err
	}
	err = bw.Flush()
	return err
}

func IsSupported(filename string) bool {
	format, err := SniffFormat(filename)
	if err != nil {
		return false
	}

	if supportedFormats.Contains(format) {
		return true
	}
	return false
}

func CanDecodeConfig(filename string) bool {
	fh, err := os.Open(filename)
	if err != nil {
		return false
	}

	defer fh.Close()
	reader := bufio.NewReader(fh)
	_, _, err = image.DecodeConfig(reader)
	return err == nil
}
