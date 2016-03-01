//+build dev

/**
 * Copyright (C) 2016 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package proxy

import (
	"gir/glib-2.0"
	C "launchpad.net/gocheck"
	"pkg.deepin.io/lib/log"
	"testing"
)

func Test(t *testing.T) { C.TestingT(t) }

type testWrapper struct{}

func init() {
	C.Suite(&testWrapper{})
}

func (*testWrapper) TestMain(c *C.C) {
	logger.SetLogLevel(log.LevelDebug)
	SetupProxy()
	logger.Info("start loop...")
	glib.StartLoop()
}
