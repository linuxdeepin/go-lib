/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package gzip

import (
	C "gopkg.in/check.v1"
	"os"
	"testing"
)

type testWrapper struct{}

func init() {
	C.Suite(&testWrapper{})
}

func Test(t *testing.T) {
	C.TestingT(t)
}

func (*testWrapper) TestTarCompresssFiles(c *C.C) {
	var infos = []struct {
		files    []string
		dest     string
		errIsNil bool
	}{
		{
			files:    []string{"testdata/tar-compress-datas"},
			dest:     "testdata/tmp-compress.tar.gz",
			errIsNil: true,
		},
		{
			files:    []string{"testdata/xxxxx"},
			dest:     "testdata/xxxxx",
			errIsNil: false,
		},
	}
	for _, info := range infos {
		err := tarCompressFiles(info.files, info.dest)
		if !info.errIsNil {
			c.Check(err, C.Not(C.Equals), nil)
		} else {
			c.Check(err, C.Equals, nil)
		}
		os.Remove(info.dest)
	}
}

func (*testWrapper) TestTarExtracteFile(c *C.C) {
	var infos = []struct {
		src      string
		dest     string
		fileNum  int
		errIsNil bool
	}{
		{
			src:      "testdata/tar-extracte-data.tar.gz",
			dest:     "testdata/tmp-extracte",
			fileNum:  2,
			errIsNil: true,
		},
		{
			src:      "testdata/xxxxx",
			dest:     "testdata/xxxxx",
			errIsNil: false,
		},
	}

	for _, info := range infos {
		files, err := tarExtracte(info.src, info.dest)
		if !info.errIsNil {
			c.Check(err, C.Not(C.Equals), nil)
			continue
		}

		c.Check(err, C.Equals, nil)
		c.Check(info.fileNum, C.Equals, len(files))
		os.RemoveAll(info.dest)
	}
}
