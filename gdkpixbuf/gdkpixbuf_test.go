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
	originClearImage        = "testdata/origin_1920x1080_clear.jpg"
	originMixImage          = "testdata/origin_1920x1080_mix.jpg"
	originWidth             = 1920
	originHeight            = 1080
	originTestImageSmall    = "testdata/origin_small_200x200.png"
	originSmallWidth        = 200
	originSmallHeight       = 200
	originTestImageIcon1    = "testdata/origin_icon_1_48x48.png"
	originTestImageIcon2    = "testdata/origin_icon_2_48x48.png"
	originIconWidth         = 48
	originIconHeight        = 48
	originImgDominantColorH = 205
	originImgDominantColorS = 0.69
	originImgDominantColorV = 0.42
)

// TODO
// info

func (*gdkpixbufTester) TestBlurImage(c *C) {
	resultFile := "testdata/test_blurimage.png"
	err := BlurImage(originMixImage, resultFile, 50, 1, FormatPng)
	if err != nil {
		c.Error(err)
	}
}

func (*gdkpixbufTester) TestBlurImageCache(c *C) {
	resultFile, useCache, err := BlurImageCache(originMixImage, 50, 1, FormatPng)
	if err != nil {
		c.Error(err)
	}
	resultFile, useCache, err = BlurImageCache(originMixImage, 50, 1, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(useCache, Equals, true)
	os.Remove(resultFile)
	resultFile, useCache, err = BlurImageCache(originMixImage, 50, 1, FormatPng)
	if err != nil {
		c.Error(err)
	}
	c.Check(useCache, Equals, false)
	fmt.Println("TestBlurImageCache:", useCache, resultFile)
}

func (*gdkpixbufTester) BenchmarkBlurImage(c *C) {
	for i := 0; i < c.N; i++ {
		resultFile := fmt.Sprintf("testdata/test_blurimage_%d.png", i)
		err := BlurImage(originMixImage, resultFile, 50, 1, FormatPng)
		if err != nil {
			c.Error(err)
		}
	}
}
