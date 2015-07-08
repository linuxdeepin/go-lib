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
	C "launchpad.net/gocheck"
	"os"
	"pkg.deepin.io/lib/utils"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { C.TestingT(t) }

type testWrapper struct{}

var _ = C.Suite(&testWrapper{})

const (
	originImg               = "testdata/origin_1920x1080.jpg"
	originImgWidth          = 1920
	originImgHeight         = 1080
	originImgDominantColorH = 198.6
	originImgDominantColorS = 0.40
	originImgDominantColorV = 0.43

	originImgIconBmp = "testdata/origin_icon_48x48.bmp"
	originImgIconGif = "testdata/origin_icon_48x48.gif"
	originImgIconTxt = "testdata/origin_icon_48x48.txt"

	originImgPngSmall       = "testdata/origin_small_200x200.png"
	originImgPngSmallWidth  = 200
	originImgPngSmallHeight = 200
	originImgIconPng1       = "testdata/origin_icon_1_48x48.png"
	originImgIconPng2       = "testdata/origin_icon_2_48x48.png"
	originImgIconWidth      = 48
	originImgIconHeight     = 48

	originImgNotImage = "testdata/origin_not_image"
)

func sumFileMd5(f string) (md5 string) {
	md5, _ = utils.SumFileMd5(f)
	return
}

func (*testWrapper) TestGetImageSize(c *C.C) {
	w, h, _ := GetImageSize(originImg)
	c.Check(w, C.Equals, originImgWidth)
	c.Check(h, C.Equals, originImgHeight)
}

func (*testWrapper) TestGetImageFormat(c *C.C) {
	var f Format
	f, _ = GetImageFormat(originImgIconPng1)
	c.Check(f, C.Equals, FormatPng)
	f, _ = GetImageFormat(originImgIconBmp)
	c.Check(f, C.Equals, FormatBmp)
	_, err := GetImageFormat(originImgIconTxt)
	c.Check(err, C.NotNil)
}

func (*testWrapper) TestIsSupportedImage(c *C.C) {
	c.Check(IsSupportedImage(originImgIconPng1), C.Equals, true)
	c.Check(IsSupportedImage(originImgIconBmp), C.Equals, true)
	c.Check(IsSupportedImage(originImgIconGif), C.Equals, true)
	c.Check(IsSupportedImage(originImgIconTxt), C.Equals, false)
	c.Check(IsSupportedImage(originImgNotImage), C.Equals, false)
	c.Check(IsSupportedImage("<file not exists>"), C.Equals, false)
}

func (*testWrapper) TestGetDominantColor(c *C.C) {
	h, s, v, err := GetDominantColorOfImage(originImg)
	if err != nil {
		c.Error(err)
	}
	if delta(h, originImgDominantColorH) > 1 ||
		delta(s, originImgDominantColorS) > 0.1 ||
		delta(v, originImgDominantColorV) > 0.1 {
		c.Error("h, s, v = ", h, s, v)
	}
}
func delta(x, y float64) float64 {
	if x >= y {
		return x - y
	}
	return y - x
}

func (*testWrapper) TestBlurImage(c *C.C) {
	resultFile := "testdata/test_blurimage.png"
	err := BlurImage(originImg, resultFile, 50, 1, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(sumFileMd5(resultFile), C.Equals, "1b6781963a66148aed343325d08cfec0")
}

func (*testWrapper) TestBlurImageCache(c *C.C) {
	resultFile, useCache, err := BlurImageCache(originImg, 50, 1, FormatPng)
	if err != nil {
		c.Error(err)
	}
	resultFile, useCache, err = BlurImageCache(originImg, 50, 1, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(useCache, C.Equals, true)
	os.Remove(resultFile)
	resultFile, useCache, err = BlurImageCache(originImg, 50, 1, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(useCache, C.Equals, false)
}

func (*testWrapper) BenchmarkBlurImage(c *C.C) {
	for i := 0; i < c.N; i++ {
		resultFile := fmt.Sprintf("testdata/test_blurimage_%d.png", i)
		err := BlurImage(originImg, resultFile, 50, 1, FormatPng)
		if err != nil {
			c.Error(err)
		}
	}
}

func (*testWrapper) TestClipImage(c *C.C) {
	resultFile := "testdata/test_clipimage_100x200.png"
	err := ClipImage(originImg, resultFile, 0, 0, 100, 200, FormatPng)
	if err != nil {
		c.Error(err)
	}
	w, h, err := GetImageSize(resultFile)
	if err != nil {
		c.Error(err)
	}
	c.Check(int(w), C.Equals, 100)
	c.Check(int(h), C.Equals, 200)
	c.Check(sumFileMd5(resultFile), C.Equals, "0b31f921d5e9478e83eb98eda6b4252d")

	resultFile = "testdata/test_clipimage_160x160.png"
	err = ClipImage(originImg, resultFile, 40, 40, 160, 160, FormatPng)
	if err != nil {
		c.Error(err)
	}
	w, h, err = GetImageSize(resultFile)
	if err != nil {
		c.Error(err)
	}
	c.Check(int(w), C.Equals, 160)
	c.Check(int(h), C.Equals, 160)
	c.Check(sumFileMd5(resultFile), C.Equals, "ee9530add4ccf8cf77763fdb8e1d7394")
}

func (*testWrapper) TestConvertImage(c *C.C) {
	var f Format
	resultFilePng := "testdata/test_convertimage.png"
	ConvertImage(originImgPngSmall, resultFilePng, FormatPng)
	f, _ = GetImageFormat(resultFilePng)
	c.Check(f, C.Equals, FormatPng)
	c.Check(sumFileMd5(resultFilePng), C.Equals, "597e22ed7a9633950908d50e9309be21")

	resultFileJpg := "testdata/test_convertimage.jpg"
	ConvertImage(originImgPngSmall, resultFileJpg, FormatJpeg)
	f, _ = GetImageFormat(resultFileJpg)
	c.Check(f, C.Equals, FormatJpeg)
	c.Check(sumFileMd5(resultFileJpg), C.Equals, "00c7c47613c39e0321cf4df6ab87fd45")

	resultFileBmp := "testdata/test_convertimage.bmp"
	ConvertImage(originImgPngSmall, resultFileBmp, FormatBmp)
	f, _ = GetImageFormat(resultFileBmp)
	c.Check(f, C.Equals, FormatBmp)
	c.Check(sumFileMd5(resultFileBmp), C.Equals, "aef8c2806dfa927f996b18785e58650c")

	resultFileIco := "testdata/test_convertimage.ico"
	ConvertImage(originImgPngSmall, resultFileIco, FormatIco)
	f, _ = GetImageFormat(resultFileIco)
	c.Check(f, C.Equals, FormatIco)
	c.Check(sumFileMd5(resultFileIco), C.Equals, "530442dfd38b9118e944d30d82a9cd37")

	resultFileTiff := "testdata/test_convertimage.tiff"
	ConvertImage(originImgPngSmall, resultFileTiff, FormatTiff)
	f, _ = GetImageFormat(resultFileTiff)
	c.Check(f, C.Equals, FormatTiff)
	// FIXME: this hash won't be the same on different arch.
	// c.Check(sumFileMd5(resultFileTiff), C.Equals, "2d28a01653464896e02c14de58e7487c")
}

func (*testWrapper) TestFlipImageHorizontal(c *C.C) {
	resultFileHorizontal := "testdata/test_flipimage_horizontal.png"
	err := FlipImageHorizontal(originImg, resultFileHorizontal, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(sumFileMd5(resultFileHorizontal), C.Equals, "c6cb10065fea5865c2151b012ebe426c")
}

func (*testWrapper) TestFlipImageVertical(c *C.C) {
	resultFileVertical := "testdata/test_flipimage_vertical.png"
	err := FlipImageVertical(originImg, resultFileVertical, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(sumFileMd5(resultFileVertical), C.Equals, "9683192bfd772e024e8f9ece58aa0035")
}

func (*testWrapper) TestScaleImage(c *C.C) {
	resultFile := "testdata/test_scaleimage_500x600.png"
	err := ScaleImage(originImg, resultFile, 500, 600, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		c.Error(err)
	}
	w, h, err := GetImageSize(resultFile)
	if err != nil {
		c.Error(err)
	}
	c.Check(int(w), C.Equals, 500)
	c.Check(int(h), C.Equals, 600)
	c.Check(sumFileMd5(resultFile), C.Equals, "14b8e27e743b6eb66e1c98745dbbd5cf")
}

func (*testWrapper) TestScaleImagePrefer(c *C.C) {
	resultFile := "testdata/test_scaleimageprefer_500x600.png"
	err := ScaleImagePrefer(originImg, resultFile, 500, 600, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		c.Error(err)
	}
	w, h, err := GetImageSize(resultFile)
	if err != nil {
		c.Error(err)
	}
	c.Check(int(w), C.Equals, 500)
	c.Check(int(h), C.Equals, 600)
	c.Check(sumFileMd5(resultFile), C.Equals, "066950b89d6296e7860ae811d41f47e1")
}

func (*testWrapper) TestThumbnailImage(c *C.C) {
	resultFile := "testdata/test_thumbnail.png"
	maxWidth, maxHeight := 200, 200
	err := ThumbnailImage(originImg, resultFile, maxWidth, maxHeight, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		c.Error(err)
	}
	w, h, _ := GetImageSize(resultFile)
	c.Check(int(w) <= maxWidth, C.Equals, true)
	c.Check(int(h) <= maxHeight, C.Equals, true)
	c.Check(sumFileMd5(resultFile), C.Equals, "ee26f7cafacecb95ae6fa4e4333a0ba1")
}

func (*testWrapper) TestScaleImageCache(c *C.C) {
	resultFile, useCache, err := ScaleImageCache(originImg, 500, 600, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		c.Error(err)
	}
	resultFile, useCache, err = ScaleImageCache(originImg, 500, 600, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(useCache, C.Equals, true)
	os.Remove(resultFile)
	resultFile, useCache, err = ScaleImageCache(originImg, 500, 600, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(useCache, C.Equals, false)
}

func (*testWrapper) TestScaleImagePreferCache(c *C.C) {
	resultFile, useCache, err := ScaleImagePreferCache(originImg, 500, 600, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		c.Error(err)
	}
	resultFile, useCache, err = ScaleImagePreferCache(originImg, 500, 600, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(useCache, C.Equals, true)
	os.Remove(resultFile)
	resultFile, useCache, err = ScaleImagePreferCache(originImg, 500, 600, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(useCache, C.Equals, false)
}

func (*testWrapper) TestRotateImageLeft(c *C.C) {
	resultFile := "testdata/test_rotateimageleft.png"
	err := RotateImageLeft(originImg, resultFile, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(sumFileMd5(resultFile), C.Equals, "e07d7c97138985843ffab23e26d4bc8d")
}

func (*testWrapper) TestRotateImageRight(c *C.C) {
	resultFile := "testdata/test_rotateimageright.png"
	err := RotateImageRight(originImg, resultFile, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(sumFileMd5(resultFile), C.Equals, "04032dac2648bdd03e07b2b8d38c5848")
}

func (*testWrapper) TestRotateImageUpsizedown(c *C.C) {
	resultFile := "testdata/test_rotateimageupsidedown.png"
	err := RotateImageUpsizedown(originImg, resultFile, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(sumFileMd5(resultFile), C.Equals, "6fb1ca20243db7769ba9bb9521e95012")
}

func (*testWrapper) ManualTestScreenshotImage(c *C.C) {
	InitGdk()
	resultFile := "testdata/test_screenshot.png"
	err := ScreenshotImage(resultFile, FormatPng)
	if err != nil {
		c.Error(err)
	}
}
