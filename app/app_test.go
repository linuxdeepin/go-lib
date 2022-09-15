// +build ignore

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package app

import (
	"os"
)

func ExampleApp() {
	app := New("test app", "this is a test app", "version 0.0.0")
	subcmd := app.ParseCommandLine(os.Args[1:])
	if app.StartProfile() != nil {
		// error handle
		return
	}
	switch subcmd {
	// deal with subcmd
	}
}
