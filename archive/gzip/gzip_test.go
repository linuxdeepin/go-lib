// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package gzip

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UnitTestSuite struct {
	suite.Suite
	testDataPath string
}

func (s *UnitTestSuite) SetupSuite() {
	s.testDataPath = "./testdata/tar-compress-datas"
	data := []byte("UOS Deepin")
	err := os.MkdirAll(s.testDataPath, 0777)
	require.NoError(s.T(), err)

	tmpfile, err := ioutil.TempFile(s.testDataPath, "data.dat")
	require.NoError(s.T(), err)
	defer tmpfile.Close()

	err = ioutil.WriteFile(tmpfile.Name(), data, 0777)
	require.NoError(s.T(), err)
}

func (s *UnitTestSuite) TearDownSuite() {
	s.testDataPath = "./testdata"
	err := os.RemoveAll(s.testDataPath)
	require.NoError(s.T(), err)
}

func (s *UnitTestSuite) Test_CompressDir() {
	var infos = []struct {
		files    string
		dest     string
		fileType int32
		errIsNil bool
	}{
		{
			files:    "testdata/tar-compress-datas",
			dest:     "testdata/tmp-compress.tar.gz",
			fileType: ArchiveTypeTar,
			errIsNil: true,
		},
		{
			files:    "testdata/tar-compress-datas",
			dest:     "testdata/tmp-compress.zip",
			fileType: ArchiveTypeZip,
			errIsNil: true,
		},
		{
			files:    "testdata/tar-compress-datas",
			dest:     "testdata/tmp-compress.rar",
			fileType: ArchiveTypeRar,
			errIsNil: true,
		},
		{
			files:    "testdata/xxxxx",
			dest:     "testdata/xxxxx",
			errIsNil: false,
		},
	}
	for _, info := range infos {
		err := CompressDir(info.files, info.dest, info.fileType)
		if !info.errIsNil {
			assert.NotEqual(s.T(), err, nil)
		} else {
			assert.Equal(s.T(), err, nil)
		}
	}
}

func (s *UnitTestSuite) Test_CompresssFiles() {
	var infos = []struct {
		files    []string
		dest     string
		fileType int32
		errIsNil bool
	}{
		{
			files:    []string{"testdata/tar-compress-datas"},
			dest:     "testdata/tmp-compress.tar.gz",
			fileType: ArchiveTypeTar,
			errIsNil: true,
		},
		{
			files:    []string{"testdata/tar-compress-datas"},
			dest:     "testdata/tmp-compress.zip",
			fileType: ArchiveTypeZip,
			errIsNil: true,
		},
		{
			files:    []string{"testdata/tar-compress-datas"},
			dest:     "testdata/tmp-compress.rar",
			fileType: ArchiveTypeRar,
			errIsNil: true,
		},
		{
			files:    []string{"testdata/xxxxx"},
			dest:     "testdata/xxxxx",
			errIsNil: false,
		},
	}
	for _, info := range infos {
		err := CompressFiles(info.files, info.dest, info.fileType)
		if !info.errIsNil {
			assert.NotEqual(s.T(), err, nil)
		} else {
			assert.Equal(s.T(), err, nil)
		}
	}
}

func (s *UnitTestSuite) Test_TarExtracteFile() {
	var infos = []struct {
		src      string
		dest     string
		fileType int32
		fileNum  int
		errIsNil bool
	}{
		{
			src:      "testdata/tmp-compress.tar.gz",
			dest:     "testdata/tmp-extracte",
			fileType: ArchiveTypeTar,
			fileNum:  1,
			errIsNil: true,
		},
		{
			src:      "testdata/tmp-compress.zip",
			dest:     "testdata/tmp-extracte",
			fileType: ArchiveTypeZip,
			fileNum:  0,
			errIsNil: true,
		},
		{
			src:      "testdata/tmp-compress.rar",
			dest:     "testdata/tmp-extracte",
			fileType: ArchiveTypeRar,
			fileNum:  0,
			errIsNil: true,
		},
		{
			src:      "testdata/xxxxx",
			dest:     "testdata/xxxxx",
			errIsNil: false,
		},
	}

	for _, info := range infos {
		files, err := Extracte(info.src, info.dest, info.fileType)
		if !info.errIsNil {
			assert.NotEqual(s.T(), err, nil)
			continue
		}

		assert.Equal(s.T(), err, nil)
		assert.Equal(s.T(), info.fileNum, len(files))

		// 此处需要移除解压的目录, 否则在执行其它用例时会因为文件数不对而导致测试用例执行失败
		os.Remove(info.dest)
	}
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UnitTestSuite))
}
