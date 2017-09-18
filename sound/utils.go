/*
 * Copyright (C) 2017 ~ 2017 Deepin Technology Co., Ltd.
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

package sound

import (
	"fmt"
	"os"
	"path"
	"pkg.deepin.io/lib/utils"
)

var (
	errInvalidEvent = fmt.Errorf("invalid theme event")
)

func findThemeFile(theme, event string) (string, error) {
	var home = os.Getenv("HOME")
	var themeDirs = []string{
		path.Join(home, ".local/share/sounds"),
		"/usr/local/share/sounds",
		"/usr/share/sounds",
	}

	for _, dir := range themeDirs {
		// TODO: fix non ogg/oga event
		// TODO: handle sound theme 'index.theme'
		file := path.Join(dir, theme, "stereo", event+".ogg")
		if utils.IsFileExist(file) {
			return file, nil
		}
		file = path.Join(dir, theme, "stereo", event+".oga")
		if utils.IsFileExist(file) {
			return file, nil
		}
	}
	return "", errInvalidEvent
}
