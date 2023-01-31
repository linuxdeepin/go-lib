// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package gsettings

import (
	"testing"
)

func Test_StartMonitor(t *testing.T) {
	// 依赖环境的测试
	err := StartMonitor()
	if err != nil {
		t.Skip("failed:" + err.Error())
		return
	}
}
