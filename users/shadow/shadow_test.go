/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package shadow

import (
	"syscall"
	"testing"

	C "launchpad.net/gocheck"
)

type testWrapper struct{}

func init() {
	C.Suite(&testWrapper{})
}

func Test(t *testing.T) {
	C.TestingT(t)
}

func (*testWrapper) TestGetShadowByName(c *C.C) {
	name := "root"
	shadow, err := GetShadowByName(name)

	if err != nil {
		// Permission denied, current user has no access to shadow file
		c.Check(shadow, C.IsNil)
		c.Check(err, C.Equals, syscall.EACCES)
	} else {
		c.Check(shadow.Name, C.Equals, "root")
	}
}

func (*testWrapper) TestGetShadowEntry(c *C.C) {
	shadows := GetShadowEntry()
	if len(shadows) > 0 {
		c.Check(shadows[0].Name, C.Equals, "root")
	}
}
