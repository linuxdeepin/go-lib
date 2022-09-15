// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

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
