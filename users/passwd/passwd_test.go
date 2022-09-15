// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package passwd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPasswdByName(t *testing.T) {
	name := "root"
	passwd, err := GetPasswdByName(name)
	assert.Equal(t, passwd.Home, "/root")
	assert.Equal(t, passwd.Uid, uint32(0))
	require.NoError(t, err)

	name = "root2"
	passwd, err = GetPasswdByName(name)
	require.Nil(t, passwd)
	assert.Equal(t, err, &UserNotFoundError{Name: name})
}

func TestGetPasswdByUid(t *testing.T) {
	uid := uint32(0)
	passwd, err := GetPasswdByUid(uid)
	require.NoError(t, err)
	assert.Equal(t, passwd.Name, "root")
}

func TestGetPasswdEntry(t *testing.T) {
	passwds := GetPasswdEntry()
	assert.NotEqual(t, len(passwds), 0)
	assert.Equal(t, passwds[0].Name, "root")
}
