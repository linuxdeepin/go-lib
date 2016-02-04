/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package utils

import (
	"fmt"
	"os"
	"path"
)

var DefaultCachePrefix = os.Getenv("HOME") + "/.cache/deepin"

func GenerateCacheFilePath(keyword string) (dstfile string) {
	cachePathFormat := DefaultCachePrefix + "/%s"
	md5, _ := SumStrMd5(keyword)
	dstfile = fmt.Sprintf(cachePathFormat, md5)
	EnsureDirExist(path.Dir(dstfile))
	return
}

func GenerateCacheFilePathWithPrefix(prefix, keyword string) (dstfile string) {
	graphicCacheFormat := DefaultCachePrefix + "/%s/%s"
	md5, _ := SumStrMd5(keyword)
	dstfile = fmt.Sprintf(graphicCacheFormat, prefix, md5)
	EnsureDirExist(path.Dir(dstfile))
	return
}
