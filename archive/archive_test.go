// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package archive

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
		src      string
		dest     string
		errIsNil bool
	}{
		{
			src:      "testdata/tar-compress-datas",
			dest:     "testdata/tmp-compress.tar",
			errIsNil: true,
		},
		{
			src:      "testdata/tar-compress-datas",
			dest:     "testdata/tmp-compress.tar.gz",
			errIsNil: true,
		},
		{
			src:      "testdata/tar-compress-datas",
			dest:     "testdata/tmp-compress.tar.bz2",
			errIsNil: true,
		},
		{
			src:      "testdata/tar-compress-datas",
			dest:     "testdata/tmp-compress.zip",
			errIsNil: true,
		},
		{
			src:      "testdata/xxxxx",
			dest:     "testdata/xxxxx",
			errIsNil: false,
		},
	}

	for _, info := range infos {
		err := CompressDir(info.src, info.dest)
		if !info.errIsNil {
			assert.NotEqual(s.T(), nil, err)
		} else {
			assert.Equal(s.T(), nil, err)
		}
	}
}

func (s *UnitTestSuite) Test_CompresssFiles() {
	var infos = []struct {
		files    []string
		dest     string
		errIsNil bool
	}{
		{
			files:    []string{"testdata/tar-compress-datas"},
			dest:     "testdata/tmp-compress.tar",
			errIsNil: true,
		},
		{
			files:    []string{"testdata/tar-compress-datas"},
			dest:     "testdata/tmp-compress.tar.gz",
			errIsNil: true,
		},
		{
			files:    []string{"testdata/tar-compress-datas"},
			dest:     "testdata/tmp-compress.tar.bz2",
			errIsNil: true,
		},
		{
			files:    []string{"testdata/tar-compress-datas"},
			dest:     "testdata/tmp-compress.zip",
			errIsNil: true,
		},
		{
			files:    []string{"testdata/xxxxx"},
			dest:     "testdata/xxxxx",
			errIsNil: false,
		},
	}
	for _, info := range infos {
		err := CompressFiles(info.files, info.dest)
		if !info.errIsNil {
			assert.NotEqual(s.T(), nil, err)
		} else {
			assert.Equal(s.T(), nil, err)
		}
	}
}

func (s *UnitTestSuite) Test_ExtracteFile() {
	var infos = []struct {
		src      string
		dest     string
		fileNum  int
		errIsNil bool
	}{
		{
			src:      "testdata/tmp-compress.tar",
			dest:     "testdata/tmp-extracte",
			fileNum:  1,
			errIsNil: true,
		},
		{
			src:      "testdata/tmp-compress.tar.gz",
			dest:     "testdata/tmp-extracte",
			fileNum:  1,
			errIsNil: true,
		},
		{
			src:      "testdata/tmp-compress.tar.bz2",
			dest:     "testdata/tmp-extracte",
			fileNum:  0,
			errIsNil: true,
		},
		{
			src:      "testdata/tmp-compress.zip",
			dest:     "testdata/tmp-extracte",
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
		files, err := Extracte(info.src, info.dest)

		if !info.errIsNil {
			assert.NotEqual(s.T(), nil, err)
			continue
		}

		assert.Equal(s.T(), nil, err)
		assert.Equal(s.T(), info.fileNum, len(files))

		// 此处需要移除解压的目录, 否则在执行其它用例时会因为文件数不对而导致测试用例执行失败
		os.Remove(info.dest)
	}
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UnitTestSuite))
}
