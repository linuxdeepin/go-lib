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
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func SumStrMd5(str string) (string, bool) {
	return fmt.Sprintf("%x", md5.Sum([]byte(str))), true
}

func SumFileMd5(filename string) (string, bool) {
	f, err := os.Open(filename)
	if err != nil {
		return "", false
	}
	defer f.Close()
	h := md5.New()
	io.Copy(h, f)
	return fmt.Sprintf("%x", h.Sum(nil)), true
}
