/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

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
