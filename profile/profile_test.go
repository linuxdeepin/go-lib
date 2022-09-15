// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package profile

func ExampleConfig_Start() {
	p := &Config{CPUProfile: "test"}
	if err := p.Start(); err != nil {
		// error handle
		return
	}
}

func ExampleConfig_Stop() {
	p := &Config{CPUProfile: "cpu.prof", NoShutdownHook: true}
	if err := p.Start(); err != nil {
		// error handle
		return
	}
	p.Stop()
}
