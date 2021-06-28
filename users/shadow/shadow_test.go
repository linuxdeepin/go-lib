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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetShadowByName(t *testing.T) {
	name := "root"
	shadow, err := GetShadowByName(name)

	if err != nil {
		// Permission denied, current user has no access to shadow file
		require.Nil(t, shadow)
		assert.Equal(t, err, syscall.EACCES)
	} else {
		assert.Equal(t, shadow.Name, "root")
	}
}

func TestGetShadowEntry(t *testing.T) {
	shadows := GetShadowEntry()
	if len(shadows) > 0 {
		assert.Equal(t, shadows[0].Name, "root")
	}
}
