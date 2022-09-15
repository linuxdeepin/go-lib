// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package graphic

import (
	"image"
	"os"
	dutils "github.com/linuxdeepin/go-lib/utils"
)

// GetImageSize return image's width and height.
func GetImageSize(imgfile string) (w, h int, err error) {
	f, err := os.Open(imgfile)
	if err != nil {
		return
	}
	defer f.Close()
	config, _, err := image.DecodeConfig(f)
	return config.Width, config.Height, err
}

func GetSize(img image.Image) (w, h int) {
	w = img.Bounds().Dx()
	h = img.Bounds().Dy()
	return
}

// GetImageFormat return image format, such as "png", "jpeg".
func GetImageFormat(imgfile string) (format Format, err error) {
	f, err := os.Open(imgfile)
	if err != nil {
		return
	}
	defer f.Close()
	_, name, err := image.DecodeConfig(f)
	format = Format(name)
	return
}

// IsSupportedImage check if image file is supported.
func IsSupportedImage(imgfile string) bool {
	f, err := os.Open(imgfile)
	if err != nil {
		return false
	}
	defer f.Close()
	_, _, err = image.DecodeConfig(f)
	return err == nil
}

func GetImagesInDir(dir string) ([]string, error) {
	files, err := dutils.GetFilesInDir(dir)
	if err != nil {
		return nil, err
	}

	var images []string
	for _, file := range files {
		if !IsSupportedImage(file) {
			continue
		}

		images = append(images, file)
	}

	return images, nil
}
