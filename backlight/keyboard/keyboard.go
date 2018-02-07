/*
 * Copyright (C) 2017 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

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
