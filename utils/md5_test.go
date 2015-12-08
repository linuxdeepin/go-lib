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
)

func (*testWrapper) TestMD5Sum(c *C.C) {
	testStr := "hello world"
	if ret, ok := SumStrMd5(testStr); !ok {
		c.Errorf("SumStrMd5 '%s' Faild", testStr)
		return
	} else {
		c.Check(ret, C.Equals, "5eb63bbbe01eeed093cb22bb8f5acdc3")
	}

	testFile := "testdata/testfile"
	if ret, ok := SumFileMd5(testFile); !ok {
		c.Errorf("SumFileMd5 '%s' Failed", testFile)
		return
	} else {
		c.Check(ret, C.Equals, "0a75266cc21da8c88a940b00d4d535b7")
	}
}
