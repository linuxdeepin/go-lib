// Theme checker
package checker

import (
	"fmt"
	"path"
	"pkg.deepin.io/lib/gio-2.0"
	dutils "pkg.deepin.io/lib/utils"
)

const (
	mimeTypeXCursor = "image/x-xcursor"
)

// uri: ex "file:///usr/share/themes/Deepin/index.theme"
func IsGtkTheme(uri string) (bool, error) {
	var conditions = []string{
		"gtk-2.0",
		"gtk-3.0",
		"metacity-1",
	}
	dir := path.Dir(dutils.DecodeURI(uri))
	return isFilesInDir(conditions, dir)
}

// uri: ex "file:///usr/share/icons/Deepin/index.theme"
func IsIconTheme(uri string) (bool, error) {
	file := dutils.DecodeURI(uri)
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

// uri: ex "file:///usr/share/icons/Deepin/index.theme"
func IsCursorTheme(uri string) (bool, error) {
	file := dutils.DecodeURI(uri)
	parent := path.Dir(file)
	tmp := path.Join(parent, "cursors", "left_ptr")
	mime, err := queryFileMime(tmp)
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

func queryFileMime(file string) (string, error) {
	if !dutils.IsFileExist(file) {
		return "", fmt.Errorf("Not found the file '%s'", file)
	}

	var gf = gio.FileNewForPath(file)
	defer gf.Unref()

	info, err := gf.QueryInfo(
		gio.FileAttributeStandardContentType,
		gio.FileQueryInfoFlagsNone, nil)
	if err != nil {
		return "", err
	}
	defer info.Unref()

	return info.GetAttributeString(
		gio.FileAttributeStandardContentType), nil
}
