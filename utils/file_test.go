/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package utils

import (
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCopyDir(t *testing.T) {
	src := "testdata/copy-src"
	dest := "testdata/copy-dest"

	pwd, _ := os.Getwd()
	defer func() {
		_ = os.Chdir(pwd)
	}()

	err := os.RemoveAll(dest)
	if err != nil {
		return
	}

	err = CopyDir(src, dest)
	if err != nil {
		t.Error(err)
		return
	}

	sf, _ := os.Open(src)
	defer sf.Close()
	df, _ := os.Open(dest)
	defer df.Close()

	sNames, _ := sf.Readdirnames(-1)
	dNames, _ := df.Readdirnames(-1)
	sort.Strings(sNames)
	sort.Strings(dNames)
	for i := 0; i < len(sNames); i++ {
		assert.Equal(t, sNames[i], dNames[i])
	}
}

func TestCreateFile(t *testing.T) {
	err := CreateFile("")
	assert.NotNil(t, err)

	file := "testdata/create-testfile"
	err = CreateFile(file)
	require.Nil(t, err)
	os.Remove(file)

	file = "testdata/xxx/create-testfile"
	err = CreateFile(file)
	assert.NotNil(t, err)
}

func TestSymlinkFile(t *testing.T) {
	var datas = []struct {
		src     string
		dest    string
		success bool
	}{
		{
			src:     "testdata/testfile",
			dest:    "testdata/test_symlink",
			success: true,
		},
		{
			src:     "testdata/testfile",
			dest:    "testdata/test1",
			success: false,
		},
		{
			src:     "testdata/testfile_xxx",
			dest:    "testdata/test_symlink",
			success: false,
		},
	}

	for _, data := range datas {
		if data.success {
			assert.Equal(t, SymlinkFile(data.src, data.dest), nil)
			os.Remove(data.dest)
		} else {
			assert.NotEqual(t, SymlinkFile(data.src, data.dest), nil)
		}
	}
}

func TestGetFiles(t *testing.T) {
	var datas = []struct {
		dir    string
		length int
		ret    bool
	}{
		{
			dir:    "testdata/test-get_files",
			length: 2,
			ret:    true,
		},
		{
			dir:    "testdata/xxx",
			length: 0,
			ret:    false,
		},
		{
			dir:    "testdata/testfile",
			length: 0,
			ret:    false,
		},
	}

	for _, data := range datas {
		files, err := GetFilesInDir(data.dir)
		assert.Equal(t, len(files), data.length)
		if data.ret {
			assert.Equal(t, err, nil)
		} else {
			assert.NotEqual(t, err, nil)
		}
	}
}
