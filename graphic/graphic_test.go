package graphic

import (
	C "launchpad.net/gocheck"
	"pkg.deepin.io/lib/utils"
	"testing"
)

const (
	originImg               = "testdata/origin_1920x1080.jpg"
	originImgWidth          = 1920
	originImgHeight         = 1080
	originImgDominantColorH = 198.6
	originImgDominantColorS = 0.40
	originImgDominantColorV = 0.43

	originImgPngSmall       = "testdata/origin_small_200x200.png"
	originImgPngSmallWidth  = 200
	originImgPngSmallHeight = 200
	originImgPngIcon1       = "testdata/origin_icon_1_48x48.png"
	originImgPngIcon2       = "testdata/origin_icon_2_48x48.png"
	originIconWidth         = 48
	originIconHeight        = 48

	originImgNotImage = "testdata/origin_not_image"
)

// data uri for originImgPngIcon2
const testDataUri = `data:image/png;base64,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`

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { C.TestingT(t) }

type testWrapper struct{}

var _ = C.Suite(&testWrapper{})

func delta(x, y float64) float64 {
	if x >= y {
		return x - y
	}
	return y - x
}

func sumFileMd5(f string) (md5 string) {
	md5, _ = utils.SumFileMd5(f)
	return
}

func (g *testWrapper) TestLoadImage(c *C.C) {
	_, err := LoadImage(originImg)
	if err != nil {
		c.Error(err)
	}
}

func (g *testWrapper) TestCompositeImage(c *C.C) {
	resultFile := "testdata/test_compositeimage.png"
	err := CompositeImage(originImgPngSmall, originImgPngIcon1, resultFile, 0, 0, FormatPng)
	if err != nil {
		c.Error(err)
	}
	err = CompositeImage(resultFile, originImgPngIcon2, resultFile, 24, 24, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(sumFileMd5(resultFile), C.Equals, "4209e6810ca64189048b25b3aac4b9da")
}

func (g *testWrapper) TestCompositeImageSet(c *C.C) {
	resultFile := "testdata/test_compositeimageset.png"
	compInfos := []CompositeInfo{
		{originImgPngIcon1, 0, 0, 10},
		{originImgPngIcon2, 24, 24, 0},
	}
	err := CompositeImageSet(originImgPngSmall, compInfos, resultFile, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(sumFileMd5(resultFile), C.Equals, "56e1fd0644351235098bed846f17bf12")
}

func (g *testWrapper) TestCompositeImageUri(c *C.C) {
	resultFile := "testdata/test_compositeimageuri.png"
	srcImageUri, _ := ConvertImageToDataUri(originImgPngSmall)
	compImageUri1, _ := ConvertImageToDataUri(originImgPngIcon1)
	compImageUri2, _ := ConvertImageToDataUri(originImgPngIcon2)
	resultDataUri, _ := CompositeImageUri(srcImageUri, compImageUri1, 0, 0, FormatPng)
	resultDataUri, _ = CompositeImageUri(resultDataUri, compImageUri2, 24, 24, FormatPng)
	err := ConvertDataUriToImage(resultDataUri, resultFile, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(sumFileMd5(resultFile), C.Equals, "4209e6810ca64189048b25b3aac4b9da")
}

func (g *testWrapper) TestClipImage(c *C.C) {
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
	c.Check(sumFileMd5(resultFile), C.Equals, "fa2b58bdc0aa09ae837e23828eda0445")

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
	c.Check(sumFileMd5(resultFile), C.Equals, "bce7057369bdc742faf0f2ef58bd048b")
}

func (g *testWrapper) TestConvertImage(c *C.C) {
	resultFile := "testdata/test_convertimage.png"
	err := ConvertImage(originImg, resultFile, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(sumFileMd5(resultFile), C.Equals, "d389abd6706e918a8061d33cd53f8a27")
}

func (g *testWrapper) TestConvertImageToDataUri(c *C.C) {
	dataUri, err := ConvertImageToDataUri(originImgPngIcon2)
	if err != nil {
		c.Error(err)
	}
	c.Check(dataUri, C.Equals, testDataUri)
}

func (g *testWrapper) TestConvertDataUriToImage(c *C.C) {
	resultFile := "testdata/test_convertdatauri.png"
	err := ConvertDataUriToImage(testDataUri, resultFile, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(sumFileMd5(resultFile), C.Equals, "adaef9427c10ba58e4e94aaee27c5486")
}

func (g *testWrapper) TestLoadImageFromDataUri(c *C.C) {
	img, err := LoadImageFromDataUri(testDataUri)
	if err != nil {
		c.Error(err)
	}
	w, h := GetSize(img)
	c.Check(w, C.Equals, originIconWidth)
	c.Check(h, C.Equals, originIconHeight)
}

func (g *testWrapper) TestFillImage(c *C.C) {
	resultFile := "testdata/test_flllimage_tile_200x200.png"
	err := FillImage(originImg, resultFile, 200, 200, FillTile, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(sumFileMd5(resultFile), C.Equals, "3f12c9173e7d41bc697ca0a2ae56d7e0")

	resultFile = "testdata/test_flllimage_tile_1600x1000.png"
	err = FillImage(originImg, resultFile, 1600, 1000, FillTile, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(sumFileMd5(resultFile), C.Equals, "bc2ce31bc1bfa42737557a76510b8d68")

	resultFile = "testdata/test_flllimage_center_400x400.png"
	err = FillImage(originImg, resultFile, 400, 400, FillCenter, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(sumFileMd5(resultFile), C.Equals, "3dda8c7af16b6580e448f32594e23ddd")

	resultFile = "testdata/test_flllimage_center_1600x1000.png"
	err = FillImage(originImg, resultFile, 1600, 1000, FillCenter, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(sumFileMd5(resultFile), C.Equals, "e0f0d511b2988aa24b9b2f5d29ef12b0")
}

func (g *testWrapper) TestFillImageCache(c *C.C) {
	FillImageCache(originImg, 1024, 768, FillTile, FormatPng)
	_, useCache, err := FillImageCache(originImg, 1024, 768, FillTile, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(useCache, C.Equals, true)
}

func (g *testWrapper) TestFlipImageHorizontal(c *C.C) {
	resultFile := "testdata/test_flipimagehorizontal.png"
	err := FlipImageHorizontal(originImg, resultFile, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(sumFileMd5(resultFile), C.Equals, "5447f8349de7c114f8a7f2b11aef481e")
}

func (g *testWrapper) TestFlipImageVertical(c *C.C) {
	resultFile := "testdata/test_flipimagevertical.png"
	err := FlipImageVertical(originImg, resultFile, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(sumFileMd5(resultFile), C.Equals, "914262cef701016c67f3260df994f0ae")
}

func (g *testWrapper) TestGetDominantColor(c *C.C) {
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

// Test that a subset of RGB space can be converted to HSV and back to within
// 1/256 tolerance.
func (g *testWrapper) TestHsv(c *C.C) {
	for r := 0; r < 255; r += 7 {
		for g := 0; g < 255; g += 5 {
			for b := 0; b < 255; b += 3 {
				r0, g0, b0 := uint8(r), uint8(g), uint8(b)
				h, s, v := Rgb2Hsv(r0, g0, b0)
				r1, g1, b1 := Hsv2Rgb(h, s, v)
				if delta(float64(r0), float64(r1)) > 1 || delta(float64(g0), float64(g1)) > 1 || delta(float64(b0), float64(b1)) > 1 {
					c.Fatalf("r0, g0, b0 = %d, %d, %d   r1, g1, b1 = %d, %d, %d", r0, g0, b0, r1, g1, b1)
				}
			}
		}
	}
}

func (g *testWrapper) TestGetImageSize(c *C.C) {
	w, h, err := GetImageSize(originImg)
	if err != nil {
		c.Error(err)
	}
	c.Check(int(w), C.Equals, originImgWidth)
	c.Check(int(h), C.Equals, originImgHeight)
}

func (g *testWrapper) TestGetImageFormat(c *C.C) {
	format, err := GetImageFormat(originImg)
	if err != nil {
		c.Error(err)
	}
	c.Check(format, C.Equals, FormatJpeg)
}

func (g *testWrapper) TestIsSupportedImage(c *C.C) {
	c.Check(IsSupportedImage(originImg), C.Equals, true)
	c.Check(IsSupportedImage(originImgNotImage), C.Equals, false)
	c.Check(IsSupportedImage("<file not exists>"), C.Equals, false)
}

func (g *testWrapper) TestScaleImage(c *C.C) {
	resultFile := "testdata/test_scaleimage_500x600.png"
	err := ScaleImage(originImg, resultFile, 500, 600, FormatPng)
	if err != nil {
		c.Error(err)
	}
	w, h, err := GetImageSize(resultFile)
	if err != nil {
		c.Error(err)
	}
	c.Check(int(w), C.Equals, 500)
	c.Check(int(h), C.Equals, 600)
	c.Check(sumFileMd5(resultFile), C.Equals, "66a43cf723d520e107cfa0284a605ebf")
}

func (g *testWrapper) TestScaleImagePrefer(c *C.C) {
	resultFile := "testdata/test_scaleimageprefer_500x600.png"
	err := ScaleImagePrefer(originImg, resultFile, 500, 600, FormatPng)
	if err != nil {
		c.Error(err)
	}
	w, h, err := GetImageSize(resultFile)
	if err != nil {
		c.Error(err)
	}
	c.Check(int(w), C.Equals, 500)
	c.Check(int(h), C.Equals, 600)
	c.Check(sumFileMd5(resultFile), C.Equals, "d523c11dfb54b43f92f89d7525d38455")
}

func (g *testWrapper) TestScaleImageCache(c *C.C) {
	ScaleImageCache(originImg, 200, 200, FormatPng)
	_, useCache, err := ScaleImageCache(originImg, 200, 200, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(useCache, C.Equals, true)
}

func (g *testWrapper) TestThumbnailImage(c *C.C) {
	resultFile := "testdata/test_thumbnail.png"
	maxWidth, maxHeight := 200, 200
	err := ThumbnailImage(originImg, resultFile, maxWidth, maxHeight, FormatPng)
	if err != nil {
		c.Error(err)
	}
	w, h, _ := GetImageSize(resultFile)
	c.Check(int(w) <= maxWidth, C.Equals, true)
	c.Check(int(h) <= maxHeight, C.Equals, true)
	c.Check(sumFileMd5(resultFile), C.Equals, "8852517df5cc9c1fbf12f110084574a9")
}

func (g *testWrapper) TestThumbnailImageCache(c *C.C) {
	ThumbnailImageCache(originImg, 200, 200, FormatPng)
	_, useCache, err := ThumbnailImageCache(originImg, 200, 200, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(useCache, C.Equals, true)
}

func (g *testWrapper) TestRotateImageLeft(c *C.C) {
	resultFile := "testdata/test_rotateimageleft.png"
	err := RotateImageLeft(originImg, resultFile, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(sumFileMd5(resultFile), C.Equals, "e28427752e8b1fa7a298909f04c3d1a8")
}

func (g *testWrapper) TestRotateImageRight(c *C.C) {
	resultFile := "testdata/test_rotateimageright.png"
	err := RotateImageRight(originImg, resultFile, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(sumFileMd5(resultFile), C.Equals, "e9142938dda0537c7e1bf9eb0a0345f6")
}

func (*testWrapper) TestGetPreferScaleClipRect(c *C.C) {
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
		c.Check(err, C.Equals, nil)
		c.Check(x, C.Equals, d.x)
		c.Check(y, C.Equals, d.y)
		c.Check(w, C.Equals, d.w)
		c.Check(h, C.Equals, d.h)
	}

	// check same clip rectangle size with original width and height
	for i := 1; i < 10000; i++ {
		x, y, w, h, _ := GetPreferScaleClipRect(i, 768, i, 768)
		c.Check(x, C.Equals, 0)
		c.Check(y, C.Equals, 0)
		c.Check(w, C.Equals, i)
		c.Check(h, C.Equals, 768)
	}

	// check exceptions
	var err error
	_, _, _, _, err = GetPreferScaleClipRect(0, 0, 100, 100)
	c.Check(err, C.NotNil)
	_, _, _, _, err = GetPreferScaleClipRect(1024, 768, 512, 0)
	c.Check(err, C.NotNil)
	_, _, _, _, err = GetPreferScaleClipRect(1024, 768, 0, 384)
	c.Check(err, C.NotNil)
	_, _, _, _, err = GetPreferScaleClipRect(1024, 768, 0, 0)
	c.Check(err, C.NotNil)
}

func (g *testWrapper) TestNewImageWithColor(c *C.C) {
	resultFile := "testdata/test_newimagewithcolor.png"
	err := NewImageWithColor(resultFile, 10, 10, uint8(12), uint8(200), uint8(12), uint8(220), FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(sumFileMd5(resultFile), C.Equals, "21ba3463aa8daabf86685ff92a7bc164")
}

func (*testWrapper) TestGetIcons(c *C.C) {
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
			c.Check(err, C.Equals, nil)
			c.Check(len(icons), C.Equals, len(data.images))
			for _, v := range icons {
				c.Check(v, C.Equals, data.images[v])
			}
		} else {
			c.Check(err, C.Not(C.Equals), nil)
			c.Check(len(icons), C.Equals, 0)
		}
	}
}
