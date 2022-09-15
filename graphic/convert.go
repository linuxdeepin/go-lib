// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package graphic

import (
	"fmt"
)

// ConvertImage converts from any recognized format to target format image.
func ConvertImage(srcfile, dstfile string, f Format) (err error) {
	srcimg, err := LoadImage(srcfile)
	if err != nil {
		return
	}
	return SaveImage(dstfile, srcimg, f)
}

// ConvertImageCache converts from any recognized format to cache
// directory, if already exists, just return it.
func ConvertImageCache(srcfile string, f Format) (dstfile string, useCache bool, err error) {
	dstfile = generateCacheFilePath(fmt.Sprintf("ConvertImageCache%s%s", srcfile, f))
	if isFileExists(dstfile) {
		useCache = true
		return
	}
	err = ConvertImage(srcfile, dstfile, f)
	return
}
