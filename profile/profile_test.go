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
