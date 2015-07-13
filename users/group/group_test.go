// Copyright (c) 2015 Deepin Ltd. All rights reserved.
// Use of this source is govered by General Public License that can be found
// in the LICENSE file.
package group

import (
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
