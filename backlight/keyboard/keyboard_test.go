// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package keyboard

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	controllers, err := List()
	t.Log(err)
	if len(controllers) == 0 {
		t.Log("not found")
		return
	}
	for _, c := range controllers {
		t.Logf("%+v\n", c)
		br, _ := c.GetBrightness()
		t.Log("brightness", br)
	}
}

func Test_list(t *testing.T) {
	controllers, err := list("./testdata")
	require.NoError(t, err)
	assert.Len(t, controllers, 1)

	controller := controllers[0]
	assert.Equal(t, controller.Name, "xxx__kbd_backlight")
	assert.Equal(t, controller.MaxBrightness, 3)

	br, err := controller.GetBrightness()
	require.NoError(t, err)
	assert.Equal(t, br, 1)
}
