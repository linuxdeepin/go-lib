// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package gsettings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_StartMonitor(t *testing.T) {
	require.Nil(t, StartMonitor())
}
