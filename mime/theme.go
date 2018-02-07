/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
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

package mime

import (
	"fmt"
	"path"
	dutils "pkg.deepin.io/lib/utils"
)

const (
	mimeTypeXCursor = "image/x-xcursor"
)

func queryThemeMime(file string) (string, error) {
	gtk, _ := isGtkTheme(file)
	if gtk {
		return MimeTypeGtk, nil
	}

	icon, _ := isIconTheme(file)
	if icon {
		return MimeTypeIcon, nil
	}

	cursor, _ := isCursorTheme(file)
	if cursor {
		return MimeTypeCursor, nil
	}

	return "", fmt.Errorf("The mime of '%s' not supported", file)
}

// file: ex "/usr/share/themes/Deepin/index.theme"
func isGtkTheme(file string) (bool, error) {
	var conditions = []string{
		"gtk-2.0",
		"gtk-3.0",
		"metacity-1",
	}
	dir := path.Dir(file)
	return isFilesInDir(conditions, dir)
}

// file: ex "/usr/share/icons/Deepin/index.theme"
func isIconTheme(file string) (bool, error) {
	kfile, err := dutils.NewKeyFileFromFile(file)
	if err != nil {
		return false, err
	}
	defer kfile.Free()

	_, err = kfile.GetString("Icon Theme", "Directories")
	if err != nil {
		return false, err
	}

	return true, nil
}

// file: ex "/usr/share/icons/Deepin/index.theme"
func isCursorTheme(file string) (bool, error) {
	parent := path.Dir(file)
	tmp := path.Join(parent, "cursors", "left_ptr")
	mime, err := doQueryFile(tmp)
	if err != nil {
		return false, err
	}

	if mime != mimeTypeXCursor {
		return false, fmt.Errorf("Invalid xcursor file '%s'", tmp)
	}
	return true, nil
}

func isFilesInDir(files []string, dir string) (bool, error) {
	if !dutils.IsDir(dir) {
		return false, fmt.Errorf("'%s' not a dir", dir)
	}

	for _, file := range files {
		tmp := path.Join(dir, file)
		if !dutils.IsFileExist(tmp) {
			return false, fmt.Errorf("Not found the file: '%s'", tmp)
		}
	}
	return true, nil
}
