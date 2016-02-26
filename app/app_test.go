// +build ignore

/**
 * Copyright (C) 2015 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

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
