// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package mime

import (
	"fmt"
	"path"
	dutils "github.com/linuxdeepin/go-lib/utils"
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
