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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetGroupByName(t *testing.T) {
	name := "root"
	group, err := GetGroupByName(name)
	require.NoError(t, err)
	assert.Equal(t, group.Gid, uint32(0))

	name = "root2"
	group, err = GetGroupByName(name)
	require.Nil(t, group)
	assert.Equal(t, err, &GroupNotFoundError{Name: name})
}

func TestGetGroupByGid(t *testing.T) {
	uid := uint32(0)
	group, err := GetGroupByGid(uid)
	require.NoError(t, err)
	assert.Equal(t, group.Name, "root")
}

func TestGetGroupEntry(t *testing.T) {
	groups := GetGroupEntry()
	assert.NotEqual(t, len(groups), 0)
}
