// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package utils

import (
	"fmt"
	"io"
	"os"
	"path"
)

func CopyFile(src, dest string) (err error) {
	if dest == src {
		return nil
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

func CopyDir(src, dest string) error {
	sInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	return iterCopyDir(src, dest, sInfo.Mode())
}

func MoveDir(src, dest string) error {
	if !IsDir(src) {
		return fmt.Errorf("%q not a dir", src)
	}

	err := CopyDir(src, dest)
	if err != nil {
		return err
	}

	return os.RemoveAll(src)
}

func MoveFile(src, dest string) error {
	if IsDir(src) {
		return fmt.Errorf("%q not a file", src)
	}

	err := CopyFile(src, dest)
	if err != nil {
		return err
	}

	return os.RemoveAll(src)
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

func SymlinkFile(src, dest string) error {
	if !IsFileExist(src) {
		return fmt.Errorf("The src file '%s' not exist.", src)
	}

	if IsFileExist(dest) {
		if !IsSymlink(dest) {
			return fmt.Errorf("The dest file '%s' already exists.", dest)
		}

		err := os.Remove(dest)
		if err != nil {
			return err
		}
	}

	return os.Symlink(src, dest)
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

func CreateFile(filename string) error {
	fp, err := os.Create(filename)
	if err != nil {
		return err
	}
	return fp.Close()
}

func GetFilesInDir(dir string) ([]string, error) {
	if !IsDir(dir) {
		return nil, fmt.Errorf("The '%s' is not a dir", dir)
	}

	fp, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	names, err := fp.Readdirnames(0)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, name := range names {
		file := path.Join(dir, name)
		files = append(files, file)
	}

	return files, nil
}

// 同步写入文件
func SyncWriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

func iterCopyDir(src, dest string, mode os.FileMode) error {
	sr, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sr.Close()

	finfos, err := sr.Readdir(0)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, mode)
	if err != nil {
		return err
	}

	for _, fi := range finfos {
		sTmp := path.Join(src, fi.Name())
		dTmp := path.Join(dest, fi.Name())

		var err error
		if fi.IsDir() {
			err = iterCopyDir(sTmp, dTmp, fi.Mode())
		} else {
			err = CopyFile(sTmp, dTmp)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

