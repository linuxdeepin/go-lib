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
	"fmt"
	"io"
	"os"
)

func CopyFile(src, dest string) (err error) {
	if dest == src {
		return fmt.Errorf("source and destination are same file")
	}

	sf, err := os.Open(src)
	if err != nil {
		return
	}
	defer sf.Close()

	df, err := os.Create(dest)
	if err != nil {
		return
	}
	defer df.Close()

	_, err = io.Copy(df, sf)
	return
}

func IsFileExist(path string) bool {
	// if is uri path, ensure it decoded
	path = DecodeURI(path)
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func IsDir(path string) bool {
	// if is uri path, ensure it decoded
	path = DecodeURI(path)
	f, err := os.Stat(path)
	if err != nil {
		return false
	}
	return f.IsDir()
}

func IsSymlink(path string) bool {
	// if is uri path, ensure it decoded
	path = DecodeURI(path)
	f, err := os.Lstat(path)
	if err != nil {
		return false
	}
	if f.Mode()&os.ModeSymlink == os.ModeSymlink {
		return true
	}
	return false
}

func EnsureDirExist(path string) error {
	return os.MkdirAll(path, 0755)
}

func EnsureDirExistWithPerm(path string, perm os.FileMode) error {
	// TODO if path exists with wrong perm, fix it
	return os.MkdirAll(path, perm)
}
