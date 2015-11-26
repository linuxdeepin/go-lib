package gio

// #include <gdk/gdk.h>
// #include <stdint.h>
// void free(void*);
// #cgo pkg-config: gdk-3.0
import "C"
import "unsafe"
import "pkg.deepin.io/lib/gobject-2.0"

type GdkAppLaunchContext struct {
	AppLaunchContext
}

func (this0 *GdkAppLaunchContext) GetStaticType() gobject.Type {
	return gobject.Type(C.gdk_app_launch_context_get_type())
}

func (this0 *GdkAppLaunchContext) SetTimestamp(t uint32) *GdkAppLaunchContext {
	var this1 *C.GdkAppLaunchContext
	if this0 != nil {
		this1 = (*C.GdkAppLaunchContext)(this0.InheritedFromGAppLaunchContext())
	}
	variable0 := C.guint32(t)
	C.gdk_app_launch_context_set_timestamp(this1, variable0)
	return this0
}

func (this0 *GdkAppLaunchContext) SetIconName(iconName string) *GdkAppLaunchContext {
	var this1 *C.GdkAppLaunchContext
	var value0 *C.char
	if this0 != nil {
		this1 = (*C.GdkAppLaunchContext)(this0.InheritedFromGAppLaunchContext())
	}
	value0 = _GoStringToGString(iconName)
	defer C.free(unsafe.Pointer(value0))
	C.gdk_app_launch_context_set_icon_name(this1, value0)
	return this0
}

func GetGdkAppLaunchContext() *GdkAppLaunchContext {
	ret1 := C.gdk_display_get_app_launch_context(C.gdk_display_get_default())
	var ret2 *GdkAppLaunchContext
	ret2 = (*GdkAppLaunchContext)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
