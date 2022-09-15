// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package imgutil

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	for _, name := range []string{"deepin-music.bmp", "deepin-music.gif",
		"deepin-music.jpg", "deepin-music.png", "deepin-music.tiff"} {
		filename := filepath.Join("testdata", name)
		img, err := Load(filename)
		if assert.Nil(t, err) {
			assert.Equal(t, image.Rect(0, 0, 48, 48), img.Bounds())
		}
	}
}

func TestIsSupported(t *testing.T) {
	for _, name := range []string{"deepin-music.bmp", "deepin-music.gif",
		"deepin-music.jpg", "deepin-music.png", "deepin-music.tiff"} {

		filename := filepath.Join("testdata", name)
		assert.True(t, IsSupported(filename), "should support "+name)
	}
}

func TestCanDecodeConfig(t *testing.T) {
	for _, name := range []string{"deepin-music.bmp", "deepin-music.gif",
		"deepin-music.jpg", "deepin-music.png", "deepin-music.tiff"} {

		filename := filepath.Join("testdata", name)
		assert.True(t, CanDecodeConfig(filename), "should decode config ok "+name)
	}
}

func TestSniffFormat(t *testing.T) {
	for _, ext := range []string{"bmp", "gif", "png", "tiff"} {
		filename := "testdata/deepin-music." + ext
		format, err := SniffFormat(filename)
		if assert.Nil(t, err) {
			assert.Equal(t, ext, format)
		}
	}

	format, err := SniffFormat("testdata/deepin-music.jpg")
	if assert.Nil(t, err) {
		assert.Equal(t, "jpeg", format)
	}
}

func TestSavePng(t *testing.T) {
	img, err := Load("testdata/deepin-music.png")
	if assert.Nil(t, err) {
		err = SavePng(img, "testdata/out1.png", nil)
		assert.Nil(t, err)
		os.Remove("testdata/out1.png")

		err = SavePng(img, "testdata/out2.png", &png.Encoder{
			CompressionLevel: png.NoCompression,
		})
		assert.Nil(t, err)
		os.Remove("testdata/out2.png")
	}
}
