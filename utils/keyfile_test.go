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
)

func (*testWrapper) TestKeyFile(c *C.C) {
	file := "testdata/test_keyfile.ini"
	group := "Test"
	key := "id"
	value := "1234"

	c.Check(WriteKeyToKeyFile(file, group, key, value), C.Equals, true)
	ifc, ok := ReadKeyFromKeyFile(file, group, key, "")
	c.Check(ok, C.Equals, true)

	c.Check(value, C.Equals, ifc.(string))

	os.Remove(file)
}
