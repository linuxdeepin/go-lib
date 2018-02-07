/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
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

package shadow

import (
	"syscall"
	"testing"

	C "gopkg.in/check.v1"
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
