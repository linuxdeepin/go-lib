// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package pulse

import (
	"os"
	"testing"

	"github.com/linuxdeepin/go-lib/xdg/basedir"
)

// Please manually enable test CFLAGS in ./pulse.go

func TestQueryInfo(t *testing.T) {
	homeDir := basedir.GetUserHomeDir()
	_, err := os.Stat(homeDir)
	if os.IsNotExist(err) {
		t.Skip("home dir is not exist")
	}

	_ = GetContext()
	ctx := GetContextForced()
	if ctx == nil {
		t.Skip("Can't connect to pulseaudio.")
		return
	}
	_, err = ctx.GetServer()
	if err != nil {
		t.Fatal("Can't query server info", err)
	}
}
