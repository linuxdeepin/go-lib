// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package appinfo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AppLaunchContext(t *testing.T) {
	// 仅用于测试, 没有实际意义
	testTimeStamp := (uint32)(20210622)
	testCmdFixes := []string{"uos", "deepin"}

	appLaunchTest := &AppLaunchContext{
		timestamp:   testTimeStamp,
		cmdPrefixes: testCmdFixes,
		cmdSuffixes: testCmdFixes,
	}

	appLaunchTest.SetTimestamp(testTimeStamp)
	assert.Equal(t, appLaunchTest.timestamp, testTimeStamp)

	appLaunchTest.SetCmdPrefixes(testCmdFixes)
	assert.Equal(t, appLaunchTest.cmdPrefixes[0], "uos")
	assert.NotEqual(t, appLaunchTest.cmdPrefixes[1], "uos")

	appLaunchTest.SetCmdSuffixes(testCmdFixes)
	assert.Equal(t, appLaunchTest.cmdSuffixes[0], "uos")
	assert.NotEqual(t, appLaunchTest.cmdSuffixes[1], "uos")

	assert.Equal(t, appLaunchTest.GetTimestamp(), testTimeStamp)
	assert.Equal(t, appLaunchTest.GetCmdPrefixes()[0], "uos")
	assert.Equal(t, appLaunchTest.GetCmdSuffixes()[0], "uos")
}
