// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package gdkpixbuf

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	originImg               = "testdata/origin_1920x1080.jpg"
	originImgWidth          = 1920
	originImgHeight         = 1080
	originImgDominantColorH = 219.6
	originImgDominantColorS = 0.5625
	originImgDominantColorV = 0.3137

	originImgIconPng = "testdata/origin_icon_1_48x48.png"
	originImgIconBmp = "testdata/origin_icon_1_48x48.bmp"
	originImgIconGif = "testdata/origin_icon_1_48x48.gif"
	originImgIconTxt = "testdata/origin_icon_1_48x48.txt"

	originImgPngSmall = "testdata/origin_small_200x200.png"

	originImgNotImage = "testdata/origin_not_image"
)

func TestGetImageSize(t *testing.T) {
	w, h, _ := GetImageSize(originImg)
	assert.Equal(t, w, originImgWidth)
	assert.Equal(t, h, originImgHeight)
}

func TestGetImageFormat(t *testing.T) {
	var f Format
	f, _ = GetImageFormat(originImgIconPng)
	assert.Equal(t, f, FormatPng)
	f, _ = GetImageFormat(originImgIconBmp)
	assert.Equal(t, f, FormatBmp)
	_, err := GetImageFormat(originImgIconTxt)
	assert.Error(t, err)
}

func TestIsSupportedImage(t *testing.T) {
	assert.Equal(t, IsSupportedImage(originImgIconPng), true)
	assert.Equal(t, IsSupportedImage(originImgIconBmp), true)
	assert.Equal(t, IsSupportedImage(originImgIconGif), true)
	assert.Equal(t, IsSupportedImage(originImgIconTxt), false)
	assert.Equal(t, IsSupportedImage(originImgNotImage), false)
	assert.Equal(t, IsSupportedImage("<file not exists>"), false)
}

func TestGetDominantColor(t *testing.T) {
	h, s, v, err := GetDominantColorOfImage(originImg)
	if err != nil {
		assert.Error(t, err)
	}
	if delta(h, originImgDominantColorH) > 1 ||
		delta(s, originImgDominantColorS) > 0.1 ||
		delta(v, originImgDominantColorV) > 0.1 {
		t.Log(t, "h, s, v = ", h, s, v)
	}
}
func delta(x, y float64) float64 {
	if x >= y {
		return x - y
	}
	return y - x
}

func TestBlurImage(t *testing.T) {
	resultFile := "testdata/test_blurimage.png"
	err := BlurImage(originImg, resultFile, 50, 1, FormatPng)
	if err != nil {
		assert.Error(t, err)
	}
	//assert.Equal(t, sumFileMd5(resultFile), "1b6781963a66148aed343325d08cfec0")
}

func TestBlurImageCache(t *testing.T) {
	err := InitGdk()
	if err != nil {
		t.Skip(err.Error())
		return
	}
	_, _, err = BlurImageCache(originImg, 50, 1, FormatPng)
	if err != nil {
		assert.Error(t, err)
	}
	resultFile, useCache, err := BlurImageCache(originImg, 50, 1, FormatPng)
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, useCache, true)
	_ = os.Remove(resultFile)
	_, useCache, err = BlurImageCache(originImg, 50, 1, FormatPng)
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, useCache, false)
}

func TestClipImage(t *testing.T) {
	resultFile := "testdata/test_clipimage_100x200.png"
	err := ClipImage(originImg, resultFile, 0, 0, 100, 200, FormatPng)
	if err != nil {
		assert.Error(t, err)
	}
	w, h, err := GetImageSize(resultFile)
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, int(w), 100)
	assert.Equal(t, int(h), 200)
	// assert.Equal(t, sumFileMd5(resultFile), "0b31f921d5e9478e83eb98eda6b4252d")

	resultFile = "testdata/test_clipimage_160x160.png"
	err = ClipImage(originImg, resultFile, 40, 40, 160, 160, FormatPng)
	if err != nil {
		assert.Error(t, err)
	}
	w, h, err = GetImageSize(resultFile)
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, int(w), 160)
	assert.Equal(t, int(h), 160)
	// assert.Equal(t, sumFileMd5(resultFile), "ee9530add4ccf8cf77763fdb8e1d7394")
}

func TestConvertImage(t *testing.T) {
	var f Format
	resultFilePng := "testdata/test_convertimage.png"
	_ = ConvertImage(originImgPngSmall, resultFilePng, FormatPng)
	f, _ = GetImageFormat(resultFilePng)
	assert.Equal(t, f, FormatPng)
	// assert.Equal(t, sumFileMd5(resultFilePng), "597e22ed7a9633950908d50e9309be21")

	resultFileJpg := "testdata/test_convertimage.jpg"
	_ = ConvertImage(originImgPngSmall, resultFileJpg, FormatJpeg)
	f, _ = GetImageFormat(resultFileJpg)
	assert.Equal(t, f, FormatJpeg)
	// assert.Equal(t, sumFileMd5(resultFileJpg), "00c7c47613c39e0321cf4df6ab87fd45")

	resultFileBmp := "testdata/test_convertimage.bmp"
	_ = ConvertImage(originImgPngSmall, resultFileBmp, FormatBmp)
	f, _ = GetImageFormat(resultFileBmp)
	assert.Equal(t, f, FormatBmp)
	// assert.Equal(t, sumFileMd5(resultFileBmp), "aef8c2806dfa927f996b18785e58650c")

	resultFileIco := "testdata/test_convertimage.ico"
	_ = ConvertImage(originImgPngSmall, resultFileIco, FormatIco)
	f, _ = GetImageFormat(resultFileIco)
	assert.Equal(t, f, FormatIco)
	// assert.Equal(t, sumFileMd5(resultFileIco), "530442dfd38b9118e944d30d82a9cd37")

	resultFileTiff := "testdata/test_convertimage.tiff"
	_ = ConvertImage(originImgPngSmall, resultFileTiff, FormatTiff)
	f, _ = GetImageFormat(resultFileTiff)
	assert.Equal(t, f, FormatTiff)
	// FIXME: this hash won't be the same on different arch.
	// assert.Equal(t, sumFileMd5(resultFileTiff), "2d28a01653464896e02c14de58e7487c")
}

func TestFlipImageHorizontal(t *testing.T) {
	resultFileHorizontal := "testdata/test_flipimage_horizontal.png"
	err := FlipImageHorizontal(originImg, resultFileHorizontal, FormatPng)
	if err != nil {
		assert.Error(t, err)
	}
	//assert.Equal(t, sumFileMd5(resultFileHorizontal), "c6cb10065fea5865c2151b012ebe426c")
}

func TestFlipImageVertical(t *testing.T) {
	resultFileVertical := "testdata/test_flipimage_vertical.png"
	err := FlipImageVertical(originImg, resultFileVertical, FormatPng)
	if err != nil {
		assert.Error(t, err)
	}
	// assert.Equal(t, sumFileMd5(resultFileVertical), "9683192bfd772e024e8f9ece58aa0035")
}

func TestScaleImage(t *testing.T) {
	resultFile := "testdata/test_scaleimage_500x600.png"
	err := ScaleImage(originImg, resultFile, 500, 600, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		assert.Error(t, err)
	}
	w, h, err := GetImageSize(resultFile)
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, int(w), 500)
	assert.Equal(t, int(h), 600)
	// assert.Equal(t, sumFileMd5(resultFile), "14b8e27e743b6eb66e1c98745dbbd5cf")
}

func TestScaleImagePrefer(t *testing.T) {
	resultFile := "testdata/test_scaleimageprefer_500x600.png"
	err := ScaleImagePrefer(originImg, resultFile, 500, 600, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		assert.Error(t, err)
	}
	w, h, err := GetImageSize(resultFile)
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, int(w), 500)
	assert.Equal(t, int(h), 600)
	// assert.Equal(t, sumFileMd5(resultFile), "066950b89d6296e7860ae811d41f47e1")
}

func TestThumbnailImage(t *testing.T) {
	resultFile := "testdata/test_thumbnail.png"
	maxWidth, maxHeight := 200, 200
	err := ThumbnailImage(originImg, resultFile, maxWidth, maxHeight, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		assert.Error(t, err)
	}
	w, h, _ := GetImageSize(resultFile)
	assert.Equal(t, int(w) <= maxWidth, true)
	assert.Equal(t, int(h) <= maxHeight, true)
	// assert.Equal(t, sumFileMd5(resultFile), "ee26f7cafacecb95ae6fa4e4333a0ba1")
}

func TestScaleImageCache(t *testing.T) {
	err := InitGdk()
	if err != nil {
		t.Skip(err.Error())
		return
	}
	_, _, err = ScaleImageCache(originImg, 500, 600, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		assert.Error(t, err)
	}
	resultFile, useCache, err := ScaleImageCache(originImg, 500, 600, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, useCache, true)
	_ = os.Remove(resultFile)
	_, useCache, err = ScaleImageCache(originImg, 500, 600, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, useCache, false)
}

func TestScaleImagePreferCache(t *testing.T) {
	err := InitGdk()
	if err != nil {
		t.Skip(err.Error())
		return
	}
	_, _, err = ScaleImagePreferCache(originImg, 500, 600, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		assert.Error(t, err)
	}
	resultFile, useCache, err := ScaleImagePreferCache(originImg, 500, 600, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, useCache, true)
	_ = os.Remove(resultFile)
	_, useCache, err = ScaleImagePreferCache(originImg, 500, 600, GDK_INTERP_HYPER, FormatPng)
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, useCache, false)
}

func TestRotateImageLeft(t *testing.T) {
	resultFile := "testdata/test_rotateimageleft.png"
	err := RotateImageLeft(originImg, resultFile, FormatPng)
	if err != nil {
		assert.Error(t, err)
	}
	// assert.Equal(t, sumFileMd5(resultFile), "e07d7c97138985843ffab23e26d4bc8d")
}

func TestRotateImageRight(t *testing.T) {
	resultFile := "testdata/test_rotateimageright.png"
	err := RotateImageRight(originImg, resultFile, FormatPng)
	if err != nil {
		assert.Error(t, err)
	}
	// assert.Equal(t, sumFileMd5(resultFile), "04032dac2648bdd03e07b2b8d38c5848")
}

func TestRotateImageUpsizedown(t *testing.T) {
	resultFile := "testdata/test_rotateimageupsidedown.png"
	err := RotateImageUpsizedown(originImg, resultFile, FormatPng)
	if err != nil {
		assert.Error(t, err)
	}
	// assert.Equal(t, sumFileMd5(resultFile), "6fb1ca20243db7769ba9bb9521e95012")
}

func ManualTestScreenshotImage(t *testing.T) {
	err := InitGdk()
	if err != nil {
		t.Skip(err.Error())
		return
	}
	resultFile := "testdata/test_screenshot.png"
	err = ScreenshotImage(resultFile, FormatPng)
	if err != nil {
		assert.Error(t, err)
	}
}
