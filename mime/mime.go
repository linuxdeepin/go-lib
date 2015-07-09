package mime

import (
	"fmt"
	"pkg.deepin.io/lib/gio-2.0"
	"pkg.deepin.io/lib/theme/checker"
	dutils "pkg.deepin.io/lib/utils"
)

const (
	MimeTypeGtk    = "application/x-gtk-theme"
	MimeTypeIcon   = "application/x-icon-theme"
	MimeTypeCursor = "application/x-cursor-theme"

	mimeTypeTheme = "application/x-theme"
)

func QueryURI(uri string) (string, error) {
	file := dutils.DecodeURI(uri)
	mime, err := doQueryFile(file)
	if err != nil {
		return "", err
	}

	if mime != mimeTypeTheme {
		return mime, nil
	}

	return queryThemeMime(file)
}

func queryThemeMime(file string) (string, error) {
	gtk, _ := checker.IsGtkTheme(file)
	if gtk {
		return MimeTypeGtk, nil
	}

	icon, _ := checker.IsIconTheme(file)
	if icon {
		return MimeTypeIcon, nil
	}

	cursor, _ := checker.IsCursorTheme(file)
	if cursor {
		return MimeTypeCursor, nil
	}

	return "", fmt.Errorf("The mime of '%s' not supported", file)
}

func doQueryFile(file string) (string, error) {
	if !dutils.IsFileExist(file) {
		return "", fmt.Errorf("Not found the file '%s'", file)
	}

	gf := gio.FileNewForPath(file)
	defer gf.Unref()

	info, err := gf.QueryInfo(gio.FileAttributeStandardContentType,
		gio.FileQueryInfoFlagsNone, nil)
	if err != nil {
		return "", err
	}
	defer info.Unref()

	return info.GetAttributeString(
		gio.FileAttributeStandardContentType), nil
}
