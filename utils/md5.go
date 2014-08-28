/**
 * Copyright (c) 2011 ~ 2014 Deepin, Inc.
 *               2013 ~ 2014 jouyouyun
 *
 * Author:      jouyouyun <jouyouwen717@gmail.com>
 * Maintainer:  jouyouyun <jouyouwen717@gmail.com>
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

package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
)

// TODO refactor code, use doSumStrMd5 instead of SumStrMd5
func doSumStrMd5(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func md5ByteToStr(bytes [16]byte) string {
	str := ""

	for _, b := range bytes {
		s := strconv.FormatInt(int64(b), 16)
		if len(s) == 1 {
			str += "0" + s
		} else {
			str += s
		}
	}

	return str
}

func SumStrMd5(str string) (string, bool) {
	if len(str) < 1 {
		return "", false
	}

	md5Byte := md5.Sum([]byte(str))
	md5Str := md5ByteToStr(md5Byte)
	if len(md5Str) < 32 {
		return "", false
	}

	return md5Str, true
}

func SumFileMd5(filename string) (string, bool) {
	if !IsFileExist(filename) {
		return "", false
	}

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", false
	}

	md5Byte := md5.Sum(contents)
	md5Str := md5ByteToStr(md5Byte)
	if len(md5Str) < 32 {
		return "", false
	}

	return md5Str, true
}

//SysMd5Sum will call sh to exec md5sum to get the md5 of file
func SysMd5Sum(filename string) (string, bool) {
	if !IsFileExist(filename) {
		return "", false
	}

	cmdLine := "md5sum -b " + filename
	cmd := exec.Command("/bin/sh", "-c", cmdLine)
	out, err := cmd.Output()
	if nil != err {
		return "", false
	}

	md5Str := strings.Split(string(out), " ")[0]
	if len(md5Str) < 32 {
		return "", false
	}

	return md5Str, true
}
