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
	C "launchpad.net/gocheck"
	"os"
	"sort"
)

func (*testWrapper) TestCopyDir(c *C.C) {
	src := "testdata/copy-src"
	dest := "testdata/copy-dest"

	pwd, _ := os.Getwd()
	defer os.Chdir(pwd)

	err := os.RemoveAll(dest)
	if err != nil {
		return
	}

	err = CopyDir(src, dest)
	if err != nil {
		c.Error(err)
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
		c.Check(sNames[i], C.Equals, dNames[i])
	}
}
