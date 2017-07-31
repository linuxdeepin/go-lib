/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package utils

import (
	C "gopkg.in/check.v1"
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
