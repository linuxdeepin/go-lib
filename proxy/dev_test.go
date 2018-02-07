//+build dev

/*
 * Copyright (C) 2016 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package proxy

import (
	"gir/glib-2.0"
	C "gopkg.in/check.v1"
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
