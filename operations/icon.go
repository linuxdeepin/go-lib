// gtk_init should be invokded first.
package operations

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -std=c99
// #include <stdlib.h>
// char* icon_name_to_path_with_check_xpm(const char* name, int size);
// char* get_icon_for_app(char* file_path, int size);
// char* get_icon_for_file(char* icons, int size);
import "C"
import "unsafe"
import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"pkg.deepin.io/lib/gio-2.0"
)

func getIcon(icon *gio.Icon, size int, fn func(*C.char, C.int) *C.char) string {
	iconStr := icon.ToString()
	if iconStr == "" {
		return ""
	}

	cIconStr := C.CString(iconStr)
	defer C.free(unsafe.Pointer(cIconStr))

	cIcon := fn(cIconStr, C.int(size))
	defer C.free(unsafe.Pointer(cIcon))

	return C.GoString(cIcon)
}

func GetThemeIconForApp(icon *gio.Icon, size int) string {
	return getIcon(icon, size, func(icon *C.char, size C.int) *C.char {
		return C.get_icon_for_app(icon, size)
	})
}

func GetThemeIconForFile(icon *gio.Icon, size int) string {
	return getIcon(icon, size, func(icon *C.char, size C.int) *C.char {
		return C.get_icon_for_file(icon, size)
	})
}

func GetThemeIconFromIconName(iconName string, size int) string {
	if iconName == "" {
		return ""
	}

	cIconName := C.CString(iconName)
	defer C.free(unsafe.Pointer(cIconName))

	cIcon := C.icon_name_to_path_with_check_xpm(cIconName, C.int(size))
	defer C.free(unsafe.Pointer(cIcon))

	return C.GoString(cIcon)
}

const (
	_UserExecutable os.FileMode = 0500
)

func isUserExecutable(perm os.FileMode) bool {
	return perm&_UserExecutable != 0
}

func GetThemeIcon(file string, size int) string {
	icon := ""
	if filepath.Ext(file) == ".desktop" {
		u, _ := url.Parse(file)
		stat, err := os.Stat(u.Path)
		if err != nil {
			fmt.Println(err)
		} else {
			if isUserExecutable(stat.Mode().Perm()) {
				app := gio.NewDesktopAppInfoFromFilename(u.Path)
				if app != nil {
					defer app.Unref()
					gicon := app.GetIcon()
					if gicon != nil {
						icon = GetThemeIconForApp(gicon, size)
					}
				}
			}
		}
	}

	if icon == "" {
		file := gio.FileNewForCommandlineArg(file)
		if file == nil {
			return GetThemeIconFromIconName(icon, size)
		}
		defer file.Unref()

		info, _ := file.QueryInfo(gio.FileAttributeStandardIcon, gio.FileQueryInfoFlagsNone, nil)
		if info == nil {
			return icon
		}
		defer info.Unref()

		gicon := info.GetIcon()
		if gicon != nil {
			icon = GetThemeIconForFile(gicon, size)
		}
	}

	return icon
}
