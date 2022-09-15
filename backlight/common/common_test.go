// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Controller(t *testing.T) {
	c, err := NewController("../display/testdata/acpi_video0")
	require.NoError(t, err)
	brightness, err := c.GetBrightness()
	require.NoError(t, err)
	assert.Equal(t, brightness, 1)
	list, err := ListControllerPaths("../display/testdata")
	require.NoError(t, err)
	assert.Equal(t, list, []string{"../display/testdata/acpi_video0", "../display/testdata/intel_backlight"})
}
