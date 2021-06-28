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
	require.Nil(t, err)

	name = "root2"
	passwd, err = GetPasswdByName(name)
	require.Nil(t, passwd)
	assert.Equal(t, err, &UserNotFoundError{Name: name})
}

func TestGetPasswdByUid(t *testing.T) {
	uid := uint32(0)
	passwd, err := GetPasswdByUid(uid)
	require.Nil(t, err)
	assert.Equal(t, passwd.Name, "root")
}

func TestGetPasswdEntry(t *testing.T) {
	passwds := GetPasswdEntry()
	assert.NotEqual(t, len(passwds), 0)
	assert.Equal(t, passwds[0].Name, "root")
}
