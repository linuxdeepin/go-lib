// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package graphic

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	originImg               = "testdata/origin_1920x1080.jpg"
	originImgWidth          = 1920
	originImgHeight         = 1080
	originImgDominantColorH = 219.6
	originImgDominantColorS = 0.5625
	originImgDominantColorV = 0.3137

	originImgPngSmall = "testdata/origin_small_200x200.png"
	originImgPngIcon1 = "testdata/origin_icon_1_48x48.png"
	originImgPngIcon2 = "testdata/origin_icon_2_48x48.png"
	originIconWidth   = 48
	originIconHeight  = 48

	originImgNotImage = "testdata/origin_not_image"
)

// data uri for originImgPngIcon2
const testDataUri = `data:image/png;base64,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`

func delta(x, y float64) float64 {
	if x >= y {
		return x - y
	}
	return y - x
}

func TestLoadImage(t *testing.T) {
	_, err := LoadImage(originImg)
	if err != nil {
		t.Error(err)
	}
}

func TestCompositeImage(t *testing.T) {
	resultFile := "testdata/test_compositeimage.png"
	err := CompositeImage(originImgPngSmall, originImgPngIcon1, resultFile, 0, 0, FormatPng)
	if err != nil {
		t.Error(err)
	}
	err = CompositeImage(resultFile, originImgPngIcon2, resultFile, 24, 24, FormatPng)
	if err != nil {
		t.Error(err)
	}
}

func TestCompositeImageSet(t *testing.T) {
	resultFile := "testdata/test_compositeimageset.png"
	compInfos := []CompositeInfo{
		{originImgPngIcon1, 0, 0, 10},
		{originImgPngIcon2, 24, 24, 0},
	}
	err := CompositeImageSet(originImgPngSmall, compInfos, resultFile, FormatPng)
	if err != nil {
		t.Error(err)
	}
}

func TestCompositeImageUri(t *testing.T) {
	resultFile := "testdata/test_compositeimageuri.png"
	srcImageUri, _ := ConvertImageToDataUri(originImgPngSmall)
	compImageUri1, _ := ConvertImageToDataUri(originImgPngIcon1)
	compImageUri2, _ := ConvertImageToDataUri(originImgPngIcon2)
	resultDataUri, _ := CompositeImageUri(srcImageUri, compImageUri1, 0, 0, FormatPng)
	resultDataUri, _ = CompositeImageUri(resultDataUri, compImageUri2, 24, 24, FormatPng)
	err := ConvertDataUriToImage(resultDataUri, resultFile, FormatPng)
	if err != nil {
		t.Error(err)
	}
}

func TestClipImage(t *testing.T) {
	resultFile := "testdata/test_clipimage_100x200.png"
	err := ClipImage(originImg, resultFile, 0, 0, 100, 200, FormatPng)
	if err != nil {
		t.Error(err)
	}
	w, h, err := GetImageSize(resultFile)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int(w), 100)
	assert.Equal(t, int(h), 200)

	resultFile = "testdata/test_clipimage_160x160.png"
	err = ClipImage(originImg, resultFile, 40, 40, 160, 160, FormatPng)
	if err != nil {
		t.Error(err)
	}
	w, h, err = GetImageSize(resultFile)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int(w), 160)
	assert.Equal(t, int(h), 160)
}

func TestConvertImage(t *testing.T) {
	resultFile := "testdata/test_convertimage.png"
	err := ConvertImage(originImg, resultFile, FormatPng)
	if err != nil {
		t.Error(err)
	}
}

func TestConvertImageToDataUri(t *testing.T) {
	dataUri, err := ConvertImageToDataUri(originImgPngIcon2)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, dataUri, testDataUri)
}

func TestConvertDataUriToImage(t *testing.T) {
	resultFile := "testdata/test_convertdatauri.png"
	err := ConvertDataUriToImage(testDataUri, resultFile, FormatPng)
	if err != nil {
		t.Error(err)
	}
}

func TestLoadImageFromDataUri(t *testing.T) {
	img, err := LoadImageFromDataUri(testDataUri)
	if err != nil {
		t.Error(err)
	}
	w, h := GetSize(img)
	assert.Equal(t, w, originIconWidth)
	assert.Equal(t, h, originIconHeight)
}

func TestFillImage(t *testing.T) {
	resultFile := "testdata/test_flllimage_tile_200x200.png"
	err := FillImage(originImg, resultFile, 200, 200, FillTile, FormatPng)
	if err != nil {
		t.Error(err)
	}

	resultFile = "testdata/test_flllimage_tile_1600x1000.png"
	err = FillImage(originImg, resultFile, 1600, 1000, FillTile, FormatPng)
	if err != nil {
		t.Error(err)
	}

	resultFile = "testdata/test_flllimage_center_400x400.png"
	err = FillImage(originImg, resultFile, 400, 400, FillCenter, FormatPng)
	if err != nil {
		t.Error(err)
	}

	resultFile = "testdata/test_flllimage_center_1600x1000.png"
	err = FillImage(originImg, resultFile, 1600, 1000, FillCenter, FormatPng)
	if err != nil {
		t.Error(err)
	}
}

func TestFillImageCache(t *testing.T) {
	_, _, err := FillImageCache(originImg, 1024, 768, FillTile, FormatPng)
	if err != nil {
		t.Skip("Fill image cache failed:" + err.Error())
		return
	}
	_, useCache, err := FillImageCache(originImg, 1024, 768, FillTile, FormatPng)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, useCache, true)
}

func TestFlipImageHorizontal(t *testing.T) {
	resultFile := "testdata/test_flipimagehorizontal.png"
	err := FlipImageHorizontal(originImg, resultFile, FormatPng)
	if err != nil {
		t.Error(err)
	}
}

func TestFlipImageVertical(t *testing.T) {
	resultFile := "testdata/test_flipimagevertical.png"
	err := FlipImageVertical(originImg, resultFile, FormatPng)
	if err != nil {
		t.Error(err)
	}
}

func TestGetDominantColor(t *testing.T) {
	h, s, v, err := GetDominantColorOfImage(originImg)
	if err != nil {
		t.Error(err)
	}
	if delta(h, originImgDominantColorH) > 1 ||
		delta(s, originImgDominantColorS) > 0.1 ||
		delta(v, originImgDominantColorV) > 0.1 {
		t.Error("dh, ds, dv, h, s, v = ", delta(h, originImgDominantColorH), delta(s, originImgDominantColorS), delta(v, originImgDominantColorV), h, s, v)
	}
}

// Test that a subset of RGB space can be converted to HSV and back to within
// 1/256 tolerance.
func TestHsv(t *testing.T) {
	for r := 0; r < 255; r += 7 {
		for g := 0; g < 255; g += 5 {
			for b := 0; b < 255; b += 3 {
				r0, g0, b0 := uint8(r), uint8(g), uint8(b)
				h, s, v := Rgb2Hsv(r0, g0, b0)
				r1, g1, b1 := Hsv2Rgb(h, s, v)
				if delta(float64(r0), float64(r1)) > 1 || delta(float64(g0), float64(g1)) > 1 || delta(float64(b0), float64(b1)) > 1 {
					t.Fatalf("r0, g0, b0 = %d, %d, %d   r1, g1, b1 = %d, %d, %d", r0, g0, b0, r1, g1, b1)
				}
			}
		}
	}
}

func TestGetImageSize(t *testing.T) {
	w, h, err := GetImageSize(originImg)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int(w), originImgWidth)
	assert.Equal(t, int(h), originImgHeight)
}

func TestGetImageFormat(t *testing.T) {
	format, err := GetImageFormat(originImg)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, format, FormatJpeg)
}

func TestIsSupportedImage(t *testing.T) {
	assert.Equal(t, IsSupportedImage(originImg), true)
	assert.Equal(t, IsSupportedImage(originImgNotImage), false)
	assert.Equal(t, IsSupportedImage("<file not exists>"), false)
}

func TestScaleImage(t *testing.T) {
	resultFile := "testdata/test_scaleimage_500x600.png"
	err := ScaleImage(originImg, resultFile, 500, 600, FormatPng)
	if err != nil {
		t.Error(err)
	}
	w, h, err := GetImageSize(resultFile)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int(w), 500)
	assert.Equal(t, int(h), 600)
}

func TestScaleImagePrefer(t *testing.T) {
	resultFile := "testdata/test_scaleimageprefer_500x600.png"
	err := ScaleImagePrefer(originImg, resultFile, 500, 600, FormatPng)
	if err != nil {
		t.Error(err)
	}
	w, h, err := GetImageSize(resultFile)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int(w), 500)
	assert.Equal(t, int(h), 600)
}

func TestScaleImageCache(t *testing.T) {
	_, _, err := ScaleImageCache(originImg, 200, 200, FormatPng)
	if err != nil {
		t.Skip("Scale image cache failed:" + err.Error())
		return
	}
	_, useCache, err := ScaleImageCache(originImg, 200, 200, FormatPng)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, useCache, true)
}

func TestThumbnailImage(t *testing.T) {
	resultFile := "testdata/test_thumbnail.png"
	maxWidth, maxHeight := 200, 200
	err := ThumbnailImage(originImg, resultFile, maxWidth, maxHeight, FormatPng)
	if err != nil {
		t.Error(err)
	}
	w, h, _ := GetImageSize(resultFile)
	assert.Equal(t, int(w) <= maxWidth, true)
	assert.Equal(t, int(h) <= maxHeight, true)
}

func TestThumbnailImageCache(t *testing.T) {
	_, _, err := ThumbnailImageCache(originImg, 200, 200, FormatPng)
	if err != nil {
		t.Skip("Thumb image cache failed:" + err.Error())
		return
	}
	_, useCache, err := ThumbnailImageCache(originImg, 200, 200, FormatPng)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, useCache, true)
}

func TestRotateImageLeft(t *testing.T) {
	resultFile := "testdata/test_rotateimageleft.png"
	err := RotateImageLeft(originImg, resultFile, FormatPng)
	if err != nil {
		t.Error(err)
	}
}

func TestRotateImageRight(t *testing.T) {
	resultFile := "testdata/test_rotateimageright.png"
	err := RotateImageRight(originImg, resultFile, FormatPng)
	if err != nil {
		t.Error(err)
	}
}

func TestGetPreferScaleClipRect(t *testing.T) {
	data := []struct {
		refWidth, refHeight, imgWidth, imgHeight int
		x, y, w, h                               int
	}{
		{1024, 768, 512, 100, 189, 0, 133, 100},
		{1024, 768, 100, 384, 0, 154, 100, 75},

		{1024, 768, 512, 384, 0, 0, 512, 384},
		{1024, 768, 1024, 768, 0, 0, 1024, 768},
		{1440, 900, 1440, 900, 0, 0, 1440, 900},

		{1024, 768, 1920, 1080, 240, 0, 1440, 1080},
	}
	for _, d := range data {
		x, y, w, h, err := GetPreferScaleClipRect(d.refWidth, d.refHeight, d.imgWidth, d.imgHeight)
		assert.Equal(t, err, nil)
		assert.Equal(t, x, d.x)
		assert.Equal(t, y, d.y)
		assert.Equal(t, w, d.w)
		assert.Equal(t, h, d.h)
	}

	// check same clip rectangle size with original width and height
	for i := 1; i < 10000; i++ {
		x, y, w, h, _ := GetPreferScaleClipRect(i, 768, i, 768)
		assert.Equal(t, x, 0)
		assert.Equal(t, y, 0)
		assert.Equal(t, w, i)
		assert.Equal(t, h, 768)
	}

	// check exceptions
	var err error
	_, _, _, _, err = GetPreferScaleClipRect(0, 0, 100, 100)
	assert.Error(t, err)
	_, _, _, _, err = GetPreferScaleClipRect(1024, 768, 512, 0)
	assert.Error(t, err)
	_, _, _, _, err = GetPreferScaleClipRect(1024, 768, 0, 384)
	assert.Error(t, err)
	_, _, _, _, err = GetPreferScaleClipRect(1024, 768, 0, 0)
	assert.Error(t, err)
}

func TestNewImageWithColor(t *testing.T) {
	resultFile := "testdata/test_newimagewithcolor.png"
	err := NewImageWithColor(resultFile, 10, 10, uint8(12), uint8(200), uint8(12), uint8(220), FormatPng)
	if err != nil {
		t.Error(err)
	}
}

func TestGetIcons(t *testing.T) {
	var datas = []struct {
		dir    string
		images map[string]string
		ret    bool
	}{
		{
			dir: "testdata/test-get_images",
			images: map[string]string{
				"testdata/test-get_images/1.png": "testdata/test-get_images/1.png",
				"testdata/test-get_images/2.png": "testdata/test-get_images/2.png",
				"testdata/test-get_images/3.png": "testdata/test-get_images/3.png",
			},
			ret: true,
		},
		{
			dir:    "testdata/test-get_images-noimage",
			images: nil,
			ret:    true,
		},
		{
			dir:    "testdata/origin_icon.txt",
			images: nil,
			ret:    false,
		},
	}

	for _, data := range datas {
		icons, err := GetImagesInDir(data.dir)
		if data.ret {
			assert.Equal(t, err, nil)
			assert.Equal(t, len(icons), len(data.images))
			for _, v := range icons {
				assert.Equal(t, v, data.images[v])
			}
		} else {
			assert.Error(t, err)
			assert.Equal(t, len(icons), 0)
		}
	}
}

func TestSniffImageFormat(t *testing.T) {
	var datas = []struct {
		file       string
		formatName string
	}{
		{
			file:       "testdata/origin_1920x1080.jpg",
			formatName: "jpeg",
		},
		{
			file:       "testdata/origin_icon_1_48x48.bmp",
			formatName: "bmp",
		},
		{
			file:       "testdata/origin_icon_1_48x48.gif",
			formatName: "gif",
		},
		{
			file:       "testdata/origin_icon_1_48x48.png",
			formatName: "png",
		},
		{
			file:       "testdata/origin_not_image",
			formatName: "",
		},
		{
			file:       "testdata/sniff_format.tiff",
			formatName: "tiff",
		},
	}

	for _, data := range datas {
		formatName, err := SniffImageFormat(data.file)
		require.NoError(t, err)
		assert.NoError(t, err)
		assert.Equal(t, formatName, data.formatName)
	}
}
