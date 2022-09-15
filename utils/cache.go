// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

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
	_ = EnsureDirExist(path.Dir(dstfile))
	return
}

func GenerateCacheFilePathWithPrefix(prefix, keyword string) (dstfile string) {
	graphicCacheFormat := DefaultCachePrefix + "/%s/%s"
	md5, _ := SumStrMd5(keyword)
	dstfile = fmt.Sprintf(graphicCacheFormat, prefix, md5)
	_ = EnsureDirExist(path.Dir(dstfile))
	return
}
