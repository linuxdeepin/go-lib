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

package gdkpixbuf

import (
	"fmt"
	. "launchpad.net/gocheck"
	"os"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) {
	TestingT(t)
}

type gdkpixbufTester struct{}

var _ = Suite(&gdkpixbufTester{})

const (
	originImg       = "testdata/origin_1920x1080.jpg"
	originImgWidth  = 1920
	originImgHeight = 1080

	originImgIconBmp = "testdata/origin_icon_48x48.bmp"
	originImgIconGif = "testdata/origin_icon_48x48.gif"
	originImgIconTxt = "testdata/origin_icon_48x48.txt"

	originImgJpgClear               = "testdata/origin_1920x1080_clear.jpg"
	originImgJpgClearDominantColorH = 205
	originImgJpgClearDominantColorS = 0.69
	originImgJpgClearDominantColorV = 0.42

	originImgJpgMix         = "testdata/origin_1920x1080_mix.jpg"
	originImgPngSmall       = "testdata/origin_small_200x200.png"
	originImgPngSmallWidth  = 200
	originImgPngSmallHeight = 200
	originImgIconPng1       = "testdata/origin_icon_1_48x48.png"
	originImgIconPng2       = "testdata/origin_icon_2_48x48.png"
	originImgIconWidth      = 48
	originImgIconHeight     = 48
)

func (*gdkpixbufTester) TestGetImageSize(c *C) {
	w, h, _ := GetImageSize(originImg)
	c.Check(w, Equals, originImgWidth)
	c.Check(h, Equals, originImgHeight)
}

func (*gdkpixbufTester) TestGetImageFormat(c *C) {
	var f Format
	f, _ = GetImageFormat(originImgIconPng1)
	c.Check(f, Equals, FormatPng)
	f, _ = GetImageFormat(originImgIconBmp)
	c.Check(f, Equals, FormatBmp)
	// TODO
	// f, _ = GetImageFormat(originImgIconGif)
	// c.Check(f, Equals, FormatGif)
	_, err := GetImageFormat(originImgIconTxt)
	c.Check(err, NotNil)
}

func (*gdkpixbufTester) TestIsSupportedImage(c *C) {
	c.Check(IsSupportedImage(originImgIconPng1), Equals, true)
	c.Check(IsSupportedImage(originImgIconBmp), Equals, true)
	c.Check(IsSupportedImage(originImgIconGif), Equals, true)
	c.Check(IsSupportedImage(originImgIconTxt), Equals, false)
}

func (*gdkpixbufTester) TestGetDominantColor(c *C) {
	h, s, v, err := GetDominantColorOfImage(originImgJpgClear)
	if err != nil {
		c.Error(err)
	}
	if delta(h, originImgJpgClearDominantColorH) > 1 ||
		delta(s, originImgJpgClearDominantColorS) > 0.1 ||
		delta(v, originImgJpgClearDominantColorV) > 0.1 {
		c.Error("h, s, v = ", h, s, v)
	}
}
func delta(x, y float64) float64 {
	if x >= y {
		return x - y
	}
	return y - x
}

func (*gdkpixbufTester) TestBlurImage(c *C) {
	resultFile := "testdata/test_blurimage.png"
	err := BlurImage(originImg, resultFile, 50, 1, FormatPng)
	if err != nil {
		c.Error(err)
	}
}

func (*gdkpixbufTester) TestBlurImageCache(c *C) {
	resultFile, useCache, err := BlurImageCache(originImg, 50, 1, FormatPng)
	if err != nil {
		c.Error(err)
	}
	resultFile, useCache, err = BlurImageCache(originImg, 50, 1, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(useCache, Equals, true)
	os.Remove(resultFile)
	resultFile, useCache, err = BlurImageCache(originImg, 50, 1, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(useCache, Equals, false)
	fmt.Println("TestBlurImageCache:", useCache, resultFile)
}

func (*gdkpixbufTester) BenchmarkBlurImage(c *C) {
	for i := 0; i < c.N; i++ {
		resultFile := fmt.Sprintf("testdata/test_blurimage_%d.png", i)
		err := BlurImage(originImg, resultFile, 50, 1, FormatPng)
		if err != nil {
			c.Error(err)
		}
	}
}

func (*gdkpixbufTester) TestClipImage(c *C) {
	resultFile := "testdata/test_clipimage_100x200.png"
	err := ClipImage(originImg, resultFile, 0, 0, 100, 200, FormatPng)
	if err != nil {
		c.Error(err)
	}
	w, h, err := GetImageSize(resultFile)
	if err != nil {
		c.Error(err)
	}
	c.Check(int(w), Equals, 100)
	c.Check(int(h), Equals, 200)

	resultFile = "testdata/test_clipimage_160x160.png"
	err = ClipImage(originImg, resultFile, 40, 40, 160, 160, FormatPng)
	if err != nil {
		c.Error(err)
	}
	w, h, err = GetImageSize(resultFile)
	if err != nil {
		c.Error(err)
	}
	c.Check(int(w), Equals, 160)
	c.Check(int(h), Equals, 160)
}

func (*gdkpixbufTester) TestConvertImage(c *C) {
	var f Format
	resultFilePng := "testdata/test_convertimage.png"
	ConvertImage(originImgPngSmall, resultFilePng, FormatPng)
	f, _ = GetImageFormat(resultFilePng)
	c.Check(f, Equals, FormatPng)

	resultFileJpg := "testdata/test_convertimage.jpg"
	ConvertImage(originImgPngSmall, resultFileJpg, FormatJpeg)
	f, _ = GetImageFormat(resultFileJpg)
	c.Check(f, Equals, FormatJpeg)

	resultFileBmp := "testdata/test_convertimage.bmp"
	ConvertImage(originImgPngSmall, resultFileBmp, FormatBmp)
	f, _ = GetImageFormat(resultFileBmp)
	c.Check(f, Equals, FormatBmp)

	resultFileIco := "testdata/test_convertimage.ico"
	ConvertImage(originImgPngSmall, resultFileIco, FormatIco)
	f, _ = GetImageFormat(resultFileIco)
	c.Check(f, Equals, FormatIco)

	resultFileTiff := "testdata/test_convertimage.tiff"
	ConvertImage(originImgPngSmall, resultFileTiff, FormatTiff)
	f, _ = GetImageFormat(resultFileTiff)
	c.Check(f, Equals, FormatTiff)

	// TODO
	// resultFileGif := "testdata/test_convertimage.gif"
	// ConvertImage(originImgPngSmall, resultFileGif, FormatGif)
	// f, _ = GetImageFormat(resultFileGif)
	// c.Check(f, Equals, FormatGif)

	// resultFileXpm := "testdata/test_convertimage.xpm"
	// ConvertImage(originImgPngSmall, resultFileXpm, FormatXpm)
	// f, _ = GetImageFormat(resultFileXpm)
	// c.Check(f, Equals, FormatXpm)
}

func (*gdkpixbufTester) TestFlipImageHorizontal(c *C) {
	resultFileHorizontal := "testdata/test_flipimage_horizontal.png"
	err := FlipImageHorizontal(originImg, resultFileHorizontal, FormatPng)
	if err != nil {
		c.Error(err)
	}
}

func (*gdkpixbufTester) TestFlipImageVertical(c *C) {
	resultFileVertical := "testdata/test_flipimage_vertical.png"
	err := FlipImageVertical(originImg, resultFileVertical, FormatPng)
	if err != nil {
		c.Error(err)
	}
}

func (*gdkpixbufTester) TestScaleImage(c *C) {
	resultFile := "testdata/test_scaleimage_500x600.png"
	err := ScaleImage(originImg, resultFile, 500, 600, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		c.Error(err)
	}
	w, h, err := GetImageSize(resultFile)
	if err != nil {
		c.Error(err)
	}
	c.Check(int(w), Equals, 500)
	c.Check(int(h), Equals, 600)
}

func (*gdkpixbufTester) TestScaleImagePrefer(c *C) {
	resultFile := "testdata/test_scaleimageprefer_500x600.png"
	err := ScaleImagePrefer(originImg, resultFile, 500, 600, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		c.Error(err)
	}
	w, h, err := GetImageSize(resultFile)
	if err != nil {
		c.Error(err)
	}
	c.Check(int(w), Equals, 500)
	c.Check(int(h), Equals, 600)
}

func (*gdkpixbufTester) TestThumbnailImage(c *C) {
	resultFile := "testdata/test_thumbnail.png"
	maxWidth, maxHeight := 200, 200
	err := ThumbnailImage(originImg, resultFile, maxWidth, maxHeight, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		c.Error(err)
	}
	w, h, _ := GetImageSize(resultFile)
	c.Check(int(w) <= maxWidth, Equals, true)
	c.Check(int(h) <= maxHeight, Equals, true)
}

func (*gdkpixbufTester) TestScaleImageCache(c *C) {
	resultFile, useCache, err := ScaleImageCache(originImg, 500, 600, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		c.Error(err)
	}
	resultFile, useCache, err = ScaleImageCache(originImg, 500, 600, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(useCache, Equals, true)
	os.Remove(resultFile)
	resultFile, useCache, err = ScaleImageCache(originImg, 500, 600, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(useCache, Equals, false)
	fmt.Println("TestScaleImageCache:", useCache, resultFile)
}

func (*gdkpixbufTester) TestScaleImagePreferCache(c *C) {
	resultFile, useCache, err := ScaleImagePreferCache(originImg, 500, 600, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		c.Error(err)
	}
	resultFile, useCache, err = ScaleImagePreferCache(originImg, 500, 600, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(useCache, Equals, true)
	os.Remove(resultFile)
	resultFile, useCache, err = ScaleImagePreferCache(originImg, 500, 600, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(useCache, Equals, false)
	fmt.Println("TestScaleImagePreferCache:", useCache, resultFile)
}

func (*gdkpixbufTester) TestRotateImageLeft(c *C) {
	resultFile := "testdata/test_rotateimageleft.png"
	err := RotateImageLeft(originImg, resultFile, FormatPng)
	if err != nil {
		c.Error(err)
	}
}

func (*gdkpixbufTester) TestRotateImageRight(c *C) {
	resultFile := "testdata/test_rotateimageright.png"
	err := RotateImageRight(originImg, resultFile, FormatPng)
	if err != nil {
		c.Error(err)
	}
}

func (*gdkpixbufTester) TestRotateImageUpsizedown(c *C) {
	resultFile := "testdata/test_rotateimageupsidedown.png"
	err := RotateImageUpsizedown(originImg, resultFile, FormatPng)
	if err != nil {
		c.Error(err)
	}
}

func (*gdkpixbufTester) ManualTestScreenshotImage(c *C) {
	InitGdk()
	resultFile := "testdata/test_screenshot.png"
	err := ScreenshotImage(resultFile, FormatPng)
	if err != nil {
		c.Error(err)
	}
}
