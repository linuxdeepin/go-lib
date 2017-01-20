package keyboard

import (
	"path/filepath"
	. "pkg.deepin.io/lib/backlight/common"
	"strings"
)

const controllersDir = "/sys/class/leds"

func List() ([]*Controller, error) {
	return list(controllersDir)
}

func list(dir string) ([]*Controller, error) {
	paths, err := ListControllerPaths(dir)
	if err != nil {
		return nil, err
	}
	controllers := make([]*Controller, 0, 1)
	for _, path := range paths {
		base := filepath.Base(path)
		if !strings.Contains(base, "kbd_backlight") {
			continue
		}
		c, err := NewController(path)
		if err != nil {
			continue
		}
		controllers = append(controllers, c)
	}
	return controllers, nil
}
