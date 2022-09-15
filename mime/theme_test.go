// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package mime

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsGtkTheme(t *testing.T) {
	ok, err := isGtkTheme("testdata/Deepin/index.theme")
	require.NoError(t, err)
	assert.Equal(t, ok, true)
}

func TestIsIconTheme(t *testing.T) {
	ok, err := isIconTheme("testdata/Deepin/index.theme")
	require.NoError(t, err)
	assert.Equal(t, ok, true)
}

func TestIsCursorTheme(t *testing.T) {
	ok, err := isCursorTheme("testdata/Deepin/index.theme")
	require.NoError(t, err)
	assert.Equal(t, ok, true)
}
