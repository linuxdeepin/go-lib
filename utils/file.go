/**
 * Copyright (c) 2011 ~ 2013 Deepin, Inc.
 *               2011 ~ 2013 jouyouyun
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
	"io/ioutil"
	"os"
)

func CopyFile(src, dest string) bool {
	if ok := IsFileExist(src); !ok && len(dest) < 1 {
		return false
	}

	contents, err := ioutil.ReadFile(src)
	if err != nil {
		return false
	}

	f, err1 := os.Create(dest + "~")
	if err1 != nil {
		return false
	}
	defer f.Close()

	if _, err := f.Write(contents); err != nil {
		return false
	}
	f.Sync()

	if err := os.Rename(dest+"~", dest); err != nil {
		return false
	}

	return true
}

func IsFileExist(path string) bool {
	// if is uri path, ensure it decoded
	path = DecodeURI(path)
	if len(path) < 1 {
		return false
	}
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// TODO refactor code, exist or exists
func EnsureDirExists(dir string) error {
	return os.MkdirAll(dir, 0755)
}

func EnsureDirExistsWithPerm(dir string, perm os.FileMode) error {
	// TODO if dir exists with wrong perm, fix it
	return os.MkdirAll(dir, perm)
}
