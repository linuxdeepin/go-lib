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

package group

import (
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

func (*testWrapper) TestGetGroupByName(c *C.C) {
	name := "root"
	group, err := GetGroupByName(name)
	c.Check(err, C.IsNil)
	c.Check(group.Gid, C.Equals, uint32(0))

	name = "root2"
	group, err = GetGroupByName(name)
	c.Check(group, C.IsNil)
	c.Check(err, C.DeepEquals, &GroupNotFoundError{Name: name})
}

func (*testWrapper) TestGetGroupByGid(c *C.C) {
	uid := uint32(0)
	group, err := GetGroupByGid(uid)
	c.Check(err, C.IsNil)
	c.Check(group.Name, C.Equals, "root")
}

func (*testWrapper) TestGetGroupEntry(c *C.C) {
	groups := GetGroupEntry()
	c.Check(len(groups), C.Not(C.Equals), 0)
	c.Check(groups[0].Name, C.Equals, "root")
}
