// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

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
	_, _ = io.Copy(h, f)
	return fmt.Sprintf("%x", h.Sum(nil)), true
}
