package gio

// #include <gdk/gdk.h>
// #cgo pkg-config: gdk-3.0
import "C"
import "unsafe"
import "pkg.deepin.io/lib/gobject-2.0"

func GetGdkAppLaunchContext() *AppLaunchContext {
	ret1 := C.gdk_display_get_app_launch_context(C.gdk_display_get_default())
	var ret2 *AppLaunchContext
	ret2 = (*AppLaunchContext)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
