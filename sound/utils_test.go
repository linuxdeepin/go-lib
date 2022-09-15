// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package sound

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findThemeFile(t *testing.T) {
	file, err := findThemeFile("freedesktop", "service-login")
	if err != nil {
		assert.Equal(t, err, fmt.Errorf("invalid theme event"))
	} else {
		assert.Equal(t, file, "/usr/share/sounds/freedesktop/stereo/service-login.oga")
	}
}
