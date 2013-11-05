package dlib

import "dlib/gobject-2.0"
import "unsafe"

/*
#include <stdlib.h>
#include <stdint.h>
#include <gio/gio.h>
#cgo pkg-config:gio-2.0

unsigned int _array_length(void* _array) { void** array = _array; unsigned int i = 0; while(array && array[i] != 0) i++; return i;}
*/
import "C"

func _GoStringToGString(x string) *C.gchar {
	if x == "\x00" {
		return nil
	}
	return (*_Ctype_gchar)(C.CString(x))
}
func _GoBoolToCBool(x bool) C.gboolean {
	if x {
		return 1
	}
	return 0
}

func _GStringToGoString(str *_Ctype_gchar) string {
	return C.GoString((*_Ctype_char)(unsafe.Pointer(str)))
}

type SettingsLike interface {
	gobject.ObjectLike
	InheritedFromGSettings() *C.GSettings
}

type Settings struct {
	gobject.Object
}

func ToSettings(objlike gobject.ObjectLike) *Settings {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Settings)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Settings)(obj)
	}
	panic("cannot cast to Settings")
}

func (this0 *Settings) InheritedFromGSettings() *C.GSettings {
	if this0 == nil {
		return nil
	}
	return (*C.GSettings)(this0.C)
}

func (this0 *Settings) GetStaticType() gobject.Type {
	return gobject.Type(C.g_settings_get_type())
}

func SettingsGetType() gobject.Type {
	return (*Settings)(nil).GetStaticType()
}
func NewSettings(schema_id0 string) *Settings {
	schema_id1 := _GoStringToGString(schema_id0)
	defer C.free(unsafe.Pointer(schema_id1))
	ret1 := C.g_settings_new(schema_id1)
	var ret2 *Settings
	ret2 = (*Settings)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func NewSettingsWithPath(schema_id0 string, path0 string) *Settings {
	schema_id1 := _GoStringToGString(schema_id0)
	defer C.free(unsafe.Pointer(schema_id1))
	path1 := _GoStringToGString(path0)
	defer C.free(unsafe.Pointer(path1))
	ret1 := C.g_settings_new_with_path(schema_id1, path1)
	var ret2 *Settings
	ret2 = (*Settings)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func SettingsListRelocatableSchemas() []string {
	ret1 := C.g_settings_list_relocatable_schemas()
	var ret2 []string
	ret2 = make([]string, uint(C._array_length(unsafe.Pointer(ret1))))
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
	}
	return ret2
}
func SettingsListSchemas() []string {
	ret1 := C.g_settings_list_schemas()
	var ret2 []string
	ret2 = make([]string, uint(C._array_length(unsafe.Pointer(ret1))))
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
	}
	return ret2
}
func SettingsSync() {
	C.g_settings_sync()
}
func (this0 *Settings) Apply() {
	var this1 *C.GSettings
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	C.g_settings_apply(this1)
}
func (this0 *Settings) Delay() {
	var this1 *C.GSettings
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	C.g_settings_delay(this1)
}
func (this0 *Settings) GetBoolean(key0 string) bool {
	var this1 *C.GSettings
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 := _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_settings_get_boolean(this1, key1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Settings) GetChild(name0 string) *Settings {
	var this1 *C.GSettings
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	name1 := _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	ret1 := C.g_settings_get_child(this1, name1)
	var ret2 *Settings
	ret2 = (*Settings)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Settings) GetDouble(key0 string) float64 {
	var this1 *C.GSettings
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 := _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_settings_get_double(this1, key1)
	var ret2 float64
	ret2 = float64(ret1)
	return ret2
}
func (this0 *Settings) GetEnum(key0 string) int {
	var this1 *C.GSettings
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 := _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_settings_get_enum(this1, key1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Settings) GetFlags(key0 string) int {
	var this1 *C.GSettings
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 := _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_settings_get_flags(this1, key1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Settings) GetHasUnapplied() bool {
	var this1 *C.GSettings
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	ret1 := C.g_settings_get_has_unapplied(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Settings) GetInt(key0 string) int {
	var this1 *C.GSettings
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 := _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_settings_get_int(this1, key1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}

func (this0 *Settings) GetString(key0 string) string {
	var this1 *C.GSettings
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 := _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_settings_get_string(this1, key1)

	var ret2 string

	ret2 = _GStringToGoString(ret1)
	C.free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Settings) GetStrv(key0 string) []string {
	var this1 *C.GSettings
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 := _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_settings_get_strv(this1, key1)
	var ret2 []string
	ret2 = make([]string, uint(C._array_length(unsafe.Pointer(ret1))))
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return ret2
}
func (this0 *Settings) GetUint(key0 string) int {
	var this1 *C.GSettings
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 := _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_settings_get_uint(this1, key1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Settings) IsWritable(name0 string) bool {
	var this1 *C.GSettings
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	name1 := _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	ret1 := C.g_settings_is_writable(this1, name1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Settings) ListChildren() []string {
	var this1 *C.GSettings
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	ret1 := C.g_settings_list_children(this1)
	var ret2 []string
	ret2 = make([]string, uint(C._array_length(unsafe.Pointer(ret1))))
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return ret2
}
func (this0 *Settings) ListKeys() []string {
	var this1 *C.GSettings
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	ret1 := C.g_settings_list_keys(this1)
	var ret2 []string
	ret2 = make([]string, uint(C._array_length(unsafe.Pointer(ret1))))
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return ret2
}
func (this0 *Settings) Reset(key0 string) {
	var this1 *C.GSettings
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 := _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	C.g_settings_reset(this1, key1)
}
func (this0 *Settings) Revert() {
	var this1 *C.GSettings
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	C.g_settings_revert(this1)
}
func (this0 *Settings) SetBoolean(key0 string, value0 bool) bool {
	var this1 *C.GSettings
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 := _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 := _GoBoolToCBool(value0)
	ret1 := C.g_settings_set_boolean(this1, key1, value1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Settings) SetDouble(key0 string, value0 float64) bool {
	var this1 *C.GSettings
	var value1 C.gdouble
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 := _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = C.gdouble(value0)
	ret1 := C.g_settings_set_double(this1, key1, value1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Settings) SetEnum(key0 string, value0 int) bool {
	var this1 *C.GSettings
	var value1 C.gint
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 := _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = C.gint(value0)
	ret1 := C.g_settings_set_enum(this1, key1, value1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Settings) SetFlags(key0 string, value0 int) bool {
	var this1 *C.GSettings
	var value1 C.guint
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 := _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = C.guint(value0)
	ret1 := C.g_settings_set_flags(this1, key1, value1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Settings) SetInt(key0 string, value0 int) bool {
	var this1 *C.GSettings
	var value1 C.gint
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 := _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = C.gint(value0)
	ret1 := C.g_settings_set_int(this1, key1, value1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Settings) SetString(key0 string, value0 string) bool {
	var this1 *C.GSettings
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 := _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 := _GoStringToGString(value0)
	defer C.free(unsafe.Pointer(value1))
	ret1 := C.g_settings_set_string(this1, key1, value1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Settings) SetStrv(key0 string, value0 []string) bool {
	var this1 *C.GSettings
	var value1 **C.gchar
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 := _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = (**C.gchar)(C.malloc(C.size_t(int(unsafe.Sizeof(*value1)) * (len(value0) + 1))))
	defer C.free(unsafe.Pointer(value1))
	for i, e := range value0 {
		(*(*[999999]*C.gchar)(unsafe.Pointer(value1)))[i] = _GoStringToGString(e)
		defer C.free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(value1)))[i]))
	}
	(*(*[999999]*C.char)(unsafe.Pointer(value1)))[len(value0)] = nil
	ret1 := C.g_settings_set_strv(this1, key1, value1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Settings) SetUint(key0 string, value0 int) bool {
	var this1 *C.GSettings
	var value1 C.guint
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 := _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = C.guint(value0)
	ret1 := C.g_settings_set_uint(this1, key1, value1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
