/**
 * Copyright (c) 2011 ~ 2015 Deepin, Inc.
 *               2013 ~ 2015 jouyouyun
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

package archive

import (
	C "launchpad.net/gocheck"
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
			dest:     "/tmp/tar-compress.tar",
			errIsNil: true,
		},
		{
			files:    []string{"testdata/xxxxx"},
			dest:     "/tmp/xxxxx",
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
			src:      "testdata/tar-extracte-data.tar",
			dest:     "/tmp/tar-extracte",
			fileNum:  2,
			errIsNil: true,
		},
		{
			src:      "testdata/xxxxx",
			dest:     "/tmp/xxxxx",
			errIsNil: false,
		},
	}

	for _, info := range infos {
		files, err := tarExtracteFile(info.src, info.dest)
		if !info.errIsNil {
			c.Check(err, C.Not(C.Equals), nil)
			continue
		}

		c.Check(err, C.Equals, nil)
		c.Check(info.fileNum, C.Equals, len(files))
		os.RemoveAll(info.dest)
	}
}
