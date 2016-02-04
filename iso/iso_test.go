/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package iso

import (
	C "launchpad.net/gocheck"
	"testing"
)

type testWrapper struct{}

func Test(t *testing.T) { C.TestingT(t) }

func init() {
	C.Suite(&testWrapper{})
}
