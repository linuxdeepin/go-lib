// gtk_init should be invokded first.
package operations

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -std=c99
// #include <stdlib.h>
// char* get_icon_for_app(char* file_path, int size);
// char* get_icon_for_file(char* icons, int size);
import "C"
import "unsafe"
import (
	"pkg.linuxdeepin.com/lib/gio-2.0"
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

func GetIconForApp(icon *gio.Icon, size int) string {
	return getIcon(icon, size, func(icon *C.char, size C.int) *C.char {
		return C.get_icon_for_app(icon, size)
	})
}

func GetIconForFile(icon *gio.Icon, size int) string {
	return getIcon(icon, size, func(icon *C.char, size C.int) *C.char {
		return C.get_icon_for_file(icon, size)
	})
}
