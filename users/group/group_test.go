// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

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
