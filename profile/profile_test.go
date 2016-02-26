/**
 * Copyright (C) 2015 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package profile_test

import (
	"pkg.deepin.io/lib/profile"
)

func ExampleCPUProfile() {
	p := &profile.Config{CPUProfile: "test"}
	if err := p.Start(); err != nil {
		// error handle
		return
	}
}

func ExampleNoShutdownHook() {
	p := &profile.Config{CPUProfile: "cpu.prof", NoShutdownHook: true}
	if err := p.Start(); err != nil {
		// error handle
		return
	}
	p.Stop()
}
