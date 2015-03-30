package gobject

/*
#include "gobject.gen.h"
#include <string.h>

extern void g_free(void*);

#include "gobject.h"

extern uint32_t g_quark_from_string(const char*);
extern void g_object_set_qdata(GObject*, uint32_t, void*);

extern void g_type_init();

extern GParamSpec *g_param_spec_ref_sink(GParamSpec*);
extern void g_param_spec_unref(GParamSpec*);

typedef int32_t (*_GSourceFunc)(void*);
extern uint32_t g_timeout_add(uint32_t, _GSourceFunc, void*);
extern int32_t g_source_remove(uint32_t);

extern int32_t fqueue_dispatcher(void*);
static uint32_t _g_timeout_add_fqueue(uint32_t time) {
	return g_timeout_add(time, fqueue_dispatcher, 0);
}
#cgo pkg-config: gobject-2.0
*/
import "C"
import "unsafe"
import "runtime"
import "reflect"
import "sync"

const alot = 999999

type _GSList struct {
	data unsafe.Pointer
	next *_GSList
}

type _GList struct {
	data unsafe.Pointer
	next *_GList
	prev *_GList
}

type _GError struct {
	domain uint32
	code int32
	message *C.char
}
func (e _GError) ToGError() GError {
	return GError{e.domain, e.code, C.GoString(e.message)}
}

type GError struct {
	Domain uint32
	Code int32
	Message string
}
func (e GError) Error() string {
	return e.Message
}

func _GoStringToGString(x string) *C.char {
	if x == "\x00" {
		return nil
	}
	return C.CString(x)
}

func _GoBoolToCBool(x bool) C.int {
	if x { return 1 }
	return 0
}

func _CInterfaceToGoInterface(iface [2]unsafe.Pointer) interface{} {
	return *(*interface{})(unsafe.Pointer(&iface))
}

func _GoInterfaceToCInterface(iface interface{}) *unsafe.Pointer {
	return (*unsafe.Pointer)(unsafe.Pointer(&iface))
}


//export _GObject_go_callback_cleanup
func _GObject_go_callback_cleanup(gofunc unsafe.Pointer) {
	Holder.Release(gofunc)
}


// blacklisted: BaseFinalizeFunc (callback)
// blacklisted: BaseInitFunc (callback)
type BindingLike interface {
	ObjectLike
	InheritedFromGBinding() *C.GBinding
}

type Binding struct {
	Object
	
}

func ToBinding(objlike ObjectLike) *Binding {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Binding)(nil).GetStaticType()
	obj := ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Binding)(obj)
	}
	panic("cannot cast to Binding")
}

func (this0 *Binding) InheritedFromGBinding() *C.GBinding {
	if this0 == nil {
		return nil
	}
	return (*C.GBinding)(this0.C)
}

func (this0 *Binding) GetStaticType() Type {
	return Type(C.g_binding_get_type())
}

func BindingGetType() Type {
	return (*Binding)(nil).GetStaticType()
}
func (this0 *Binding) GetFlags() BindingFlags {
	var this1 *C.GBinding
	if this0 != nil {
		this1 = (*C.GBinding)(this0.InheritedFromGBinding())
	}
	ret1 := C.g_binding_get_flags(this1)
	var ret2 BindingFlags
	ret2 = BindingFlags(ret1)
	return ret2
}
func (this0 *Binding) GetSource() *Object {
	var this1 *C.GBinding
	if this0 != nil {
		this1 = (*C.GBinding)(this0.InheritedFromGBinding())
	}
	ret1 := C.g_binding_get_source(this1)
	var ret2 *Object
	ret2 = (*Object)(ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Binding) GetSourceProperty() string {
	var this1 *C.GBinding
	if this0 != nil {
		this1 = (*C.GBinding)(this0.InheritedFromGBinding())
	}
	ret1 := C.g_binding_get_source_property(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *Binding) GetTarget() *Object {
	var this1 *C.GBinding
	if this0 != nil {
		this1 = (*C.GBinding)(this0.InheritedFromGBinding())
	}
	ret1 := C.g_binding_get_target(this1)
	var ret2 *Object
	ret2 = (*Object)(ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Binding) GetTargetProperty() string {
	var this1 *C.GBinding
	if this0 != nil {
		this1 = (*C.GBinding)(this0.InheritedFromGBinding())
	}
	ret1 := C.g_binding_get_target_property(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *Binding) Unbind() {
	var this1 *C.GBinding
	if this0 != nil {
		this1 = (*C.GBinding)(this0.InheritedFromGBinding())
	}
	C.g_binding_unbind(this1)
}
type BindingFlags C.uint32_t
const (
	BindingFlagsDefault BindingFlags = 0
	BindingFlagsBidirectional BindingFlags = 1
	BindingFlagsSyncCreate BindingFlags = 2
	BindingFlagsInvertBoolean BindingFlags = 4
)
// blacklisted: BindingTransformFunc (callback)
// blacklisted: BoxedFreeFunc (callback)
type CClosure struct {
	Closure Closure
	Callback unsafe.Pointer
}
func CClosureMarshalBoolean_BoxedBoxed(closure0 *Closure, return_value0 *Value, n_param_values0 uint32, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_BOOLEAN__BOXED_BOXED(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalBoolean_Flags(closure0 *Closure, return_value0 *Value, n_param_values0 uint32, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_BOOLEAN__FLAGS(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalString_ObjectPointer(closure0 *Closure, return_value0 *Value, n_param_values0 uint32, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_STRING__OBJECT_POINTER(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Boolean(closure0 *Closure, return_value0 *Value, n_param_values0 uint32, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__BOOLEAN(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Boxed(closure0 *Closure, return_value0 *Value, n_param_values0 uint32, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__BOXED(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Char(closure0 *Closure, return_value0 *Value, n_param_values0 uint32, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__CHAR(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Double(closure0 *Closure, return_value0 *Value, n_param_values0 uint32, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__DOUBLE(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Enum(closure0 *Closure, return_value0 *Value, n_param_values0 uint32, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__ENUM(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Flags(closure0 *Closure, return_value0 *Value, n_param_values0 uint32, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__FLAGS(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Float(closure0 *Closure, return_value0 *Value, n_param_values0 uint32, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__FLOAT(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Int(closure0 *Closure, return_value0 *Value, n_param_values0 uint32, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__INT(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Long(closure0 *Closure, return_value0 *Value, n_param_values0 uint32, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__LONG(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Object(closure0 *Closure, return_value0 *Value, n_param_values0 uint32, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__OBJECT(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Param(closure0 *Closure, return_value0 *Value, n_param_values0 uint32, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__PARAM(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Pointer(closure0 *Closure, return_value0 *Value, n_param_values0 uint32, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__POINTER(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_String(closure0 *Closure, return_value0 *Value, n_param_values0 uint32, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__STRING(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Uchar(closure0 *Closure, return_value0 *Value, n_param_values0 uint32, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__UCHAR(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Uint(closure0 *Closure, return_value0 *Value, n_param_values0 uint32, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__UINT(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_UintPointer(closure0 *Closure, return_value0 *Value, n_param_values0 uint32, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__UINT_POINTER(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Ulong(closure0 *Closure, return_value0 *Value, n_param_values0 uint32, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__ULONG(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Variant(closure0 *Closure, return_value0 *Value, n_param_values0 uint32, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__VARIANT(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Void(closure0 *Closure, return_value0 *Value, n_param_values0 uint32, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__VOID(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalGeneric(closure0 *Closure, return_gvalue0 *Value, n_param_values0 uint32, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_gvalue1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_gvalue1 = (*C.GValue)(unsafe.Pointer(return_gvalue0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_generic(closure1, return_gvalue1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
// blacklisted: Callback (callback)
// blacklisted: ClassFinalizeFunc (callback)
// blacklisted: ClassInitFunc (callback)
type Closure struct {
	RefCount uint32
	MetaMarshalNouse uint32
	NGuards uint32
	NFnotifiers uint32
	NInotifiers uint32
	InInotify uint32
	Floating uint32
	DerivativeFlag uint32
	InMarshal uint32
	IsInvalid uint32
	Marshal unsafe.Pointer
	Data unsafe.Pointer
	Notifiers *ClosureNotifyData
}
func NewClosureObject(sizeof_closure0 uint32, object0 ObjectLike) *Closure {
	var sizeof_closure1 C.uint32_t
	var object1 *C.GObject
	sizeof_closure1 = C.uint32_t(sizeof_closure0)
	if object0 != nil {
		object1 = (*C.GObject)(object0.InheritedFromGObject())
	}
	ret1 := C.g_closure_new_object(sizeof_closure1, object1)
	var ret2 *Closure
	ret2 = (*Closure)(unsafe.Pointer(ret1))
	return ret2
}
func NewClosureSimple(sizeof_closure0 uint32, data0 unsafe.Pointer) *Closure {
	var sizeof_closure1 C.uint32_t
	var data1 unsafe.Pointer
	sizeof_closure1 = C.uint32_t(sizeof_closure0)
	data1 = unsafe.Pointer(data0)
	ret1 := C.g_closure_new_simple(sizeof_closure1, data1)
	var ret2 *Closure
	ret2 = (*Closure)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Closure) Invalidate() {
	var this1 *C.GClosure
	this1 = (*C.GClosure)(unsafe.Pointer(this0))
	C.g_closure_invalidate(this1)
}
func (this0 *Closure) Invoke(return_value0 *Value, param_values0 []Value, invocation_hint0 unsafe.Pointer) {
	var this1 *C.GClosure
	var return_value1 *C.GValue
	var param_values1 *C.GValue
	var n_param_values1 C.uint32_t
	var invocation_hint1 unsafe.Pointer
	this1 = (*C.GClosure)(unsafe.Pointer(this0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	param_values1 = (*C.GValue)(C.malloc(C.size_t(int(unsafe.Sizeof(*param_values1)) * len(param_values0))))
	defer C.free(unsafe.Pointer(param_values1))
	for i, e := range param_values0 {
		(*(*[999999]C.GValue)(unsafe.Pointer(param_values1)))[i] = *(*C.GValue)(unsafe.Pointer(&e))
	}
	n_param_values1 = C.uint32_t(len(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	C.g_closure_invoke(this1, return_value1, n_param_values1, param_values1, invocation_hint1)
}
func (this0 *Closure) Sink() {
	var this1 *C.GClosure
	this1 = (*C.GClosure)(unsafe.Pointer(this0))
	C.g_closure_sink(this1)
}
// blacklisted: ClosureMarshal (callback)
// blacklisted: ClosureNotify (callback)
type ClosureNotifyData struct {
	Data unsafe.Pointer
	Notify unsafe.Pointer
}
type ConnectFlags C.uint32_t
const (
	ConnectFlagsAfter ConnectFlags = 1
	ConnectFlagsSwapped ConnectFlags = 2
)
type EnumClass struct {
	GTypeClass TypeClass
	Minimum int32
	Maximum int32
	NValues uint32
	_ [4]byte
	Values *EnumValue
}
type EnumValue struct {
	Value int32
	_ [4]byte
	value_name0 *C.char
	value_nick0 *C.char
}
func (this0 *EnumValue) ValueName() string {
	var value_name1 string
	value_name1 = C.GoString(this0.value_name0)
	return value_name1
}
func (this0 *EnumValue) ValueNick() string {
	var value_nick1 string
	value_nick1 = C.GoString(this0.value_nick0)
	return value_nick1
}
type FlagsClass struct {
	GTypeClass TypeClass
	Mask uint32
	NValues uint32
	Values *FlagsValue
}
type FlagsValue struct {
	Value uint32
	_ [4]byte
	value_name0 *C.char
	value_nick0 *C.char
}
func (this0 *FlagsValue) ValueName() string {
	var value_name1 string
	value_name1 = C.GoString(this0.value_name0)
	return value_name1
}
func (this0 *FlagsValue) ValueNick() string {
	var value_nick1 string
	value_nick1 = C.GoString(this0.value_nick0)
	return value_nick1
}
type InitiallyUnownedLike interface {
	ObjectLike
	InheritedFromGInitiallyUnowned() *C.GInitiallyUnowned
}

type InitiallyUnowned struct {
	Object
	
}

func ToInitiallyUnowned(objlike ObjectLike) *InitiallyUnowned {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*InitiallyUnowned)(nil).GetStaticType()
	obj := ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*InitiallyUnowned)(obj)
	}
	panic("cannot cast to InitiallyUnowned")
}

func (this0 *InitiallyUnowned) InheritedFromGInitiallyUnowned() *C.GInitiallyUnowned {
	if this0 == nil {
		return nil
	}
	return (*C.GInitiallyUnowned)(this0.C)
}

func (this0 *InitiallyUnowned) GetStaticType() Type {
	return Type(C.g_initially_unowned_get_type())
}

func InitiallyUnownedGetType() Type {
	return (*InitiallyUnowned)(nil).GetStaticType()
}
// blacklisted: InstanceInitFunc (callback)
// blacklisted: InterfaceFinalizeFunc (callback)
type InterfaceInfo struct {
	InterfaceInit unsafe.Pointer
	InterfaceFinalize unsafe.Pointer
	InterfaceData unsafe.Pointer
}
// blacklisted: InterfaceInitFunc (callback)
type ObjectLike interface {
	
	InheritedFromGObject() *C.GObject
}

type Object struct {
	C unsafe.Pointer
	
}

func ToObject(objlike ObjectLike) *Object {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Object)(nil).GetStaticType()
	obj := ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Object)(obj)
	}
	panic("cannot cast to Object")
}

func (this0 *Object) InheritedFromGObject() *C.GObject {
	if this0 == nil {
		return nil
	}
	return (*C.GObject)(this0.C)
}

func (this0 *Object) GetStaticType() Type {
	return Type(C.g_object_get_type())
}

func ObjectGetType() Type {
	return (*Object)(nil).GetStaticType()
}
// blacklisted: Object.new (method)
// blacklisted: Object.compat_control (method)
// blacklisted: Object.interface_find_property (method)
// blacklisted: Object.interface_install_property (method)
// blacklisted: Object.interface_list_properties (method)
// blacklisted: Object.bind_property (method)
// blacklisted: Object.bind_property_full (method)
// blacklisted: Object.force_floating (method)
// blacklisted: Object.freeze_notify (method)
// blacklisted: Object.get_data (method)
// blacklisted: Object.get_property (method)
// blacklisted: Object.get_qdata (method)
// blacklisted: Object.is_floating (method)
// blacklisted: Object.notify (method)
// blacklisted: Object.notify_by_pspec (method)
// blacklisted: Object.ref (method)
// blacklisted: Object.ref_sink (method)
// blacklisted: Object.replace_data (method)
// blacklisted: Object.replace_qdata (method)
// blacklisted: Object.run_dispose (method)
// blacklisted: Object.set_data (method)
// blacklisted: Object.set_property (method)
// blacklisted: Object.steal_data (method)
// blacklisted: Object.steal_qdata (method)
// blacklisted: Object.thaw_notify (method)
// blacklisted: Object.unref (method)
// blacklisted: Object.watch_closure (method)
type ObjectConstructParam struct {
	pspec0 *C.GParamSpec
	Value *Value
}
func (this0 *ObjectConstructParam) Pspec() *ParamSpec {
	var pspec1 *ParamSpec
	pspec1 = (*ParamSpec)(ObjectWrap(unsafe.Pointer(this0.pspec0), true))
	return pspec1
}
// blacklisted: ObjectFinalizeFunc (callback)
// blacklisted: ObjectGetPropertyFunc (callback)
// blacklisted: ObjectSetPropertyFunc (callback)
const ParamMask = 255
const ParamReadwrite = 0
const ParamStaticStrings = 0
const ParamUserShift = 8
type ParamFlags C.uint32_t
const (
	ParamFlagsReadable ParamFlags = 1
	ParamFlagsWritable ParamFlags = 2
	ParamFlagsConstruct ParamFlags = 4
	ParamFlagsConstructOnly ParamFlags = 8
	ParamFlagsLaxValidation ParamFlags = 16
	ParamFlagsStaticName ParamFlags = 32
	ParamFlagsPrivate ParamFlags = 32
	ParamFlagsStaticNick ParamFlags = 64
	ParamFlagsStaticBlurb ParamFlags = 128
	ParamFlagsDeprecated ParamFlags = 2147483648
)
// blacklisted: ParamSpec (object)
// blacklisted: ParamSpecBoolean (object)
// blacklisted: ParamSpecBoxed (object)
// blacklisted: ParamSpecChar (object)
// blacklisted: ParamSpecDouble (object)
// blacklisted: ParamSpecEnum (object)
// blacklisted: ParamSpecFlags (object)
// blacklisted: ParamSpecFloat (object)
// blacklisted: ParamSpecGType (object)
// blacklisted: ParamSpecInt (object)
// blacklisted: ParamSpecInt64 (object)
// blacklisted: ParamSpecLong (object)
// blacklisted: ParamSpecObject (object)
// blacklisted: ParamSpecOverride (object)
// blacklisted: ParamSpecParam (object)
// blacklisted: ParamSpecPointer (object)
// blacklisted: ParamSpecPool (struct)
// blacklisted: ParamSpecString (object)
type ParamSpecTypeInfo struct {
	InstanceSize uint16
	NPreallocs uint16
	_ [4]byte
	InstanceInit unsafe.Pointer
	ValueType Type
	Finalize unsafe.Pointer
	ValueSetDefault unsafe.Pointer
	ValueValidate unsafe.Pointer
	ValuesCmp unsafe.Pointer
}
// blacklisted: ParamSpecUChar (object)
// blacklisted: ParamSpecUInt (object)
// blacklisted: ParamSpecUInt64 (object)
// blacklisted: ParamSpecULong (object)
// blacklisted: ParamSpecUnichar (object)
// blacklisted: ParamSpecValueArray (object)
// blacklisted: ParamSpecVariant (object)
type Parameter struct {
	name0 *C.char
	Value Value
}
func (this0 *Parameter) Name() string {
	var name1 string
	name1 = C.GoString(this0.name0)
	return name1
}
const SignalFlagsMask = 511
const SignalMatchMask = 63
// blacklisted: SignalAccumulator (callback)
// blacklisted: SignalEmissionHook (callback)
type SignalFlags C.uint32_t
const (
	SignalFlagsRunFirst SignalFlags = 1
	SignalFlagsRunLast SignalFlags = 2
	SignalFlagsRunCleanup SignalFlags = 4
	SignalFlagsNoRecurse SignalFlags = 8
	SignalFlagsDetailed SignalFlags = 16
	SignalFlagsAction SignalFlags = 32
	SignalFlagsNoHooks SignalFlags = 64
	SignalFlagsMustCollect SignalFlags = 128
	SignalFlagsDeprecated SignalFlags = 256
)
type SignalInvocationHint struct {
	SignalId uint32
	Detail uint32
	RunType SignalFlags
}
type SignalMatchType C.uint32_t
const (
	SignalMatchTypeId SignalMatchType = 1
	SignalMatchTypeDetail SignalMatchType = 2
	SignalMatchTypeClosure SignalMatchType = 4
	SignalMatchTypeFunc SignalMatchType = 8
	SignalMatchTypeData SignalMatchType = 16
	SignalMatchTypeUnblocked SignalMatchType = 32
)
type SignalQuery struct {
	SignalId uint32
	_ [4]byte
	signal_name0 *C.char
	Itype Type
	SignalFlags SignalFlags
	_ [4]byte
	ReturnType Type
	NParams uint32
	_ [4]byte
	param_types0 *C.GType
}
func (this0 *SignalQuery) SignalName() string {
	var signal_name1 string
	signal_name1 = C.GoString(this0.signal_name0)
	return signal_name1
}
func (this0 *SignalQuery) ParamTypes() []Type {
	var param_types1 []Type
	for i := range param_types1 {
		param_types1[i] = Type((*(*[999999]C.GType)(unsafe.Pointer(this0.param_types0)))[i])
	}
	return param_types1
}
const TypeFlagReservedIdBit = 0x1
const TypeFundamentalMax = 255
const TypeFundamentalShift = 2
const TypeReservedBseFirst = 32
const TypeReservedBseLast = 48
const TypeReservedGlibFirst = 22
const TypeReservedGlibLast = 31
const TypeReservedUserFirst = 49
// blacklisted: ToggleNotify (callback)
type TypeCValue struct {
	_data [8]byte
}
type TypeClass struct {
	GType Type
}
func (this0 *TypeClass) PeekParent() *TypeClass {
	var this1 *C.GTypeClass
	this1 = (*C.GTypeClass)(unsafe.Pointer(this0))
	ret1 := C.g_type_class_peek_parent(this1)
	var ret2 *TypeClass
	ret2 = (*TypeClass)(unsafe.Pointer(ret1))
	return ret2
}
func TypeClassAddPrivate(g_class0 unsafe.Pointer, private_size0 uint64) {
	var g_class1 unsafe.Pointer
	var private_size1 C.uint64_t
	g_class1 = unsafe.Pointer(g_class0)
	private_size1 = C.uint64_t(private_size0)
	C.g_type_class_add_private(g_class1, private_size1)
}
func TypeClassAdjustPrivateOffset(g_class0 unsafe.Pointer, private_size_or_offset0 *int32) {
	var g_class1 unsafe.Pointer
	var private_size_or_offset1 *C.int32_t
	g_class1 = unsafe.Pointer(g_class0)
	private_size_or_offset1 = (*C.int32_t)(unsafe.Pointer(private_size_or_offset0))
	C.g_type_class_adjust_private_offset(g_class1, private_size_or_offset1)
}
func TypeClassPeek(type0 Type) *TypeClass {
	var type1 C.GType
	type1 = C.GType(type0)
	ret1 := C.g_type_class_peek(type1)
	var ret2 *TypeClass
	ret2 = (*TypeClass)(unsafe.Pointer(ret1))
	return ret2
}
func TypeClassPeekStatic(type0 Type) *TypeClass {
	var type1 C.GType
	type1 = C.GType(type0)
	ret1 := C.g_type_class_peek_static(type1)
	var ret2 *TypeClass
	ret2 = (*TypeClass)(unsafe.Pointer(ret1))
	return ret2
}
// blacklisted: TypeClassCacheFunc (callback)
type TypeDebugFlags C.uint32_t
const (
	TypeDebugFlagsNone TypeDebugFlags = 0
	TypeDebugFlagsObjects TypeDebugFlags = 1
	TypeDebugFlagsSignals TypeDebugFlags = 2
	TypeDebugFlagsMask TypeDebugFlags = 3
)
type TypeFlags C.uint32_t
const (
	TypeFlagsAbstract TypeFlags = 16
	TypeFlagsValueAbstract TypeFlags = 32
)
type TypeFundamentalFlags C.uint32_t
const (
	TypeFundamentalFlagsClassed TypeFundamentalFlags = 1
	TypeFundamentalFlagsInstantiatable TypeFundamentalFlags = 2
	TypeFundamentalFlagsDerivable TypeFundamentalFlags = 4
	TypeFundamentalFlagsDeepDerivable TypeFundamentalFlags = 8
)
type TypeFundamentalInfo struct {
	TypeFlags TypeFundamentalFlags
}
type TypeInfo struct {
	ClassSize uint16
	_ [6]byte
	BaseInit unsafe.Pointer
	BaseFinalize unsafe.Pointer
	ClassInit unsafe.Pointer
	ClassFinalize unsafe.Pointer
	ClassData unsafe.Pointer
	InstanceSize uint16
	NPreallocs uint16
	_ [4]byte
	InstanceInit unsafe.Pointer
	ValueTable *TypeValueTable
}
type TypeInstance struct {
	GClass *TypeClass
}
type TypeInterface struct {
	GType Type
	GInstanceType Type
}
func (this0 *TypeInterface) PeekParent() *TypeInterface {
	var this1 *C.GTypeInterface
	this1 = (*C.GTypeInterface)(unsafe.Pointer(this0))
	ret1 := C.g_type_interface_peek_parent(this1)
	var ret2 *TypeInterface
	ret2 = (*TypeInterface)(unsafe.Pointer(ret1))
	return ret2
}
func TypeInterfaceAddPrerequisite(interface_type0 Type, prerequisite_type0 Type) {
	var interface_type1 C.GType
	var prerequisite_type1 C.GType
	interface_type1 = C.GType(interface_type0)
	prerequisite_type1 = C.GType(prerequisite_type0)
	C.g_type_interface_add_prerequisite(interface_type1, prerequisite_type1)
}
func TypeInterfaceGetPlugin(instance_type0 Type, interface_type0 Type) *TypePlugin {
	var instance_type1 C.GType
	var interface_type1 C.GType
	instance_type1 = C.GType(instance_type0)
	interface_type1 = C.GType(interface_type0)
	ret1 := C.g_type_interface_get_plugin(instance_type1, interface_type1)
	var ret2 *TypePlugin
	ret2 = (*TypePlugin)(ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func TypeInterfacePeek(instance_class0 *TypeClass, iface_type0 Type) *TypeInterface {
	var instance_class1 *C.GTypeClass
	var iface_type1 C.GType
	instance_class1 = (*C.GTypeClass)(unsafe.Pointer(instance_class0))
	iface_type1 = C.GType(iface_type0)
	ret1 := C.g_type_interface_peek(instance_class1, iface_type1)
	var ret2 *TypeInterface
	ret2 = (*TypeInterface)(unsafe.Pointer(ret1))
	return ret2
}
func TypeInterfacePrerequisites(interface_type0 Type) (uint32, []Type) {
	var interface_type1 C.GType
	var n_prerequisites1 C.uint32_t
	interface_type1 = C.GType(interface_type0)
	ret1 := C.g_type_interface_prerequisites(interface_type1, &n_prerequisites1)
	var n_prerequisites2 uint32
	var ret2 []Type
	n_prerequisites2 = uint32(n_prerequisites1)
	ret2 = make([]Type, n_prerequisites1)
	for i := range ret2 {
		ret2[i] = Type((*(*[999999]C.GType)(unsafe.Pointer(ret1)))[i])
	}
	return n_prerequisites2, ret2
}
// blacklisted: TypeInterfaceCheckFunc (callback)
// blacklisted: TypeModule (object)
type TypePluginLike interface {
	ImplementsGTypePlugin() *C.GTypePlugin
}

type TypePlugin struct {
	Object
	TypePluginImpl
}

func (*TypePlugin) GetStaticType() Type {
	return Type(C.g_type_plugin_get_type())
}


type TypePluginImpl struct {}

func ToTypePlugin(objlike ObjectLike) *TypePlugin {
	c := objlike.InheritedFromGObject()
	obj := ObjectGrabIfType(unsafe.Pointer(c), Type(C.g_type_plugin_get_type()))
	if obj != nil {
		return (*TypePlugin)(obj)
	}
	panic("cannot cast to TypePlugin")
}

func (this0 *TypePluginImpl) ImplementsGTypePlugin() *C.GTypePlugin {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GTypePlugin)((*Object)(unsafe.Pointer(obj)).C)
}
func (this0 *TypePluginImpl) CompleteInterfaceInfo(instance_type0 Type, interface_type0 Type, info0 *InterfaceInfo) {
	var this1 *C.GTypePlugin
	var instance_type1 C.GType
	var interface_type1 C.GType
	var info1 *C.GInterfaceInfo
	if this0 != nil {
		this1 = this0.ImplementsGTypePlugin()
	}
	instance_type1 = C.GType(instance_type0)
	interface_type1 = C.GType(interface_type0)
	info1 = (*C.GInterfaceInfo)(unsafe.Pointer(info0))
	C.g_type_plugin_complete_interface_info(this1, instance_type1, interface_type1, info1)
}
func (this0 *TypePluginImpl) CompleteTypeInfo(g_type0 Type, info0 *TypeInfo, value_table0 *TypeValueTable) {
	var this1 *C.GTypePlugin
	var g_type1 C.GType
	var info1 *C.GTypeInfo
	var value_table1 *C.GTypeValueTable
	if this0 != nil {
		this1 = this0.ImplementsGTypePlugin()
	}
	g_type1 = C.GType(g_type0)
	info1 = (*C.GTypeInfo)(unsafe.Pointer(info0))
	value_table1 = (*C.GTypeValueTable)(unsafe.Pointer(value_table0))
	C.g_type_plugin_complete_type_info(this1, g_type1, info1, value_table1)
}
func (this0 *TypePluginImpl) Unuse() {
	var this1 *C.GTypePlugin
	if this0 != nil {
		this1 = this0.ImplementsGTypePlugin()
	}
	C.g_type_plugin_unuse(this1)
}
func (this0 *TypePluginImpl) Use() {
	var this1 *C.GTypePlugin
	if this0 != nil {
		this1 = this0.ImplementsGTypePlugin()
	}
	C.g_type_plugin_use(this1)
}
type TypePluginClass struct {
	BaseIface TypeInterface
	UsePlugin unsafe.Pointer
	UnusePlugin unsafe.Pointer
	CompleteTypeInfo unsafe.Pointer
	CompleteInterfaceInfo unsafe.Pointer
}
// blacklisted: TypePluginCompleteInterfaceInfo (callback)
// blacklisted: TypePluginCompleteTypeInfo (callback)
// blacklisted: TypePluginUnuse (callback)
// blacklisted: TypePluginUse (callback)
type TypeQuery struct {
	Type Type
	type_name0 *C.char
	ClassSize uint32
	InstanceSize uint32
}
func (this0 *TypeQuery) TypeName() string {
	var type_name1 string
	type_name1 = C.GoString(this0.type_name0)
	return type_name1
}
type TypeValueTable struct {
	ValueInit unsafe.Pointer
	ValueFree unsafe.Pointer
	ValueCopy unsafe.Pointer
	ValuePeekPointer unsafe.Pointer
	collect_format0 *C.char
	CollectValue unsafe.Pointer
	lcopy_format0 *C.char
	LcopyValue unsafe.Pointer
}
func (this0 *TypeValueTable) CollectFormat() string {
	var collect_format1 string
	collect_format1 = C.GoString(this0.collect_format0)
	return collect_format1
}
func (this0 *TypeValueTable) LcopyFormat() string {
	var lcopy_format1 string
	lcopy_format1 = C.GoString(this0.lcopy_format0)
	return lcopy_format1
}
const ValueCollectFormatMaxLength = 8
const ValueNocopyContents = 134217728
type Value struct {
	GType Type
	Data [2]_Value__data__union
}
type ValueArray struct {
	NValues uint32
	_ [4]byte
	Values *Value
	NPrealloced uint32
	_ [4]byte
}
func NewValueArray(n_prealloced0 uint32) *ValueArray {
	var n_prealloced1 C.uint32_t
	n_prealloced1 = C.uint32_t(n_prealloced0)
	ret1 := C.g_value_array_new(n_prealloced1)
	var ret2 *ValueArray
	ret2 = (*ValueArray)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *ValueArray) Append(value0 *Value) *ValueArray {
	var this1 *C.GValueArray
	var value1 *C.GValue
	this1 = (*C.GValueArray)(unsafe.Pointer(this0))
	value1 = (*C.GValue)(unsafe.Pointer(value0))
	ret1 := C.g_value_array_append(this1, value1)
	var ret2 *ValueArray
	ret2 = (*ValueArray)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *ValueArray) Copy() *ValueArray {
	var this1 *C.GValueArray
	this1 = (*C.GValueArray)(unsafe.Pointer(this0))
	ret1 := C.g_value_array_copy(this1)
	var ret2 *ValueArray
	ret2 = (*ValueArray)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *ValueArray) Free() {
	var this1 *C.GValueArray
	this1 = (*C.GValueArray)(unsafe.Pointer(this0))
	C.g_value_array_free(this1)
}
func (this0 *ValueArray) GetNth(index_0 uint32) *Value {
	var this1 *C.GValueArray
	var index_1 C.uint32_t
	this1 = (*C.GValueArray)(unsafe.Pointer(this0))
	index_1 = C.uint32_t(index_0)
	ret1 := C.g_value_array_get_nth(this1, index_1)
	var ret2 *Value
	ret2 = (*Value)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *ValueArray) Insert(index_0 uint32, value0 *Value) *ValueArray {
	var this1 *C.GValueArray
	var index_1 C.uint32_t
	var value1 *C.GValue
	this1 = (*C.GValueArray)(unsafe.Pointer(this0))
	index_1 = C.uint32_t(index_0)
	value1 = (*C.GValue)(unsafe.Pointer(value0))
	ret1 := C.g_value_array_insert(this1, index_1, value1)
	var ret2 *ValueArray
	ret2 = (*ValueArray)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *ValueArray) Prepend(value0 *Value) *ValueArray {
	var this1 *C.GValueArray
	var value1 *C.GValue
	this1 = (*C.GValueArray)(unsafe.Pointer(this0))
	value1 = (*C.GValue)(unsafe.Pointer(value0))
	ret1 := C.g_value_array_prepend(this1, value1)
	var ret2 *ValueArray
	ret2 = (*ValueArray)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *ValueArray) Remove(index_0 uint32) *ValueArray {
	var this1 *C.GValueArray
	var index_1 C.uint32_t
	this1 = (*C.GValueArray)(unsafe.Pointer(this0))
	index_1 = C.uint32_t(index_0)
	ret1 := C.g_value_array_remove(this1, index_1)
	var ret2 *ValueArray
	ret2 = (*ValueArray)(unsafe.Pointer(ret1))
	return ret2
}
// blacklisted: ValueTransform (callback)
// blacklisted: WeakNotify (callback)
type WeakRef struct {}
type _Value__data__union struct {
	_data [8]byte
}
// blacklisted: boxed_copy (function)
// blacklisted: boxed_free (function)
// blacklisted: cclosure_marshal_BOOLEAN__BOXED_BOXED (function)
// blacklisted: cclosure_marshal_BOOLEAN__FLAGS (function)
// blacklisted: cclosure_marshal_STRING__OBJECT_POINTER (function)
// blacklisted: cclosure_marshal_VOID__BOOLEAN (function)
// blacklisted: cclosure_marshal_VOID__BOXED (function)
// blacklisted: cclosure_marshal_VOID__CHAR (function)
// blacklisted: cclosure_marshal_VOID__DOUBLE (function)
// blacklisted: cclosure_marshal_VOID__ENUM (function)
// blacklisted: cclosure_marshal_VOID__FLAGS (function)
// blacklisted: cclosure_marshal_VOID__FLOAT (function)
// blacklisted: cclosure_marshal_VOID__INT (function)
// blacklisted: cclosure_marshal_VOID__LONG (function)
// blacklisted: cclosure_marshal_VOID__OBJECT (function)
// blacklisted: cclosure_marshal_VOID__PARAM (function)
// blacklisted: cclosure_marshal_VOID__POINTER (function)
// blacklisted: cclosure_marshal_VOID__STRING (function)
// blacklisted: cclosure_marshal_VOID__UCHAR (function)
// blacklisted: cclosure_marshal_VOID__UINT (function)
// blacklisted: cclosure_marshal_VOID__UINT_POINTER (function)
// blacklisted: cclosure_marshal_VOID__ULONG (function)
// blacklisted: cclosure_marshal_VOID__VARIANT (function)
// blacklisted: cclosure_marshal_VOID__VOID (function)
// blacklisted: cclosure_marshal_generic (function)
// blacklisted: enum_complete_type_info (function)
// blacklisted: enum_get_value (function)
// blacklisted: enum_get_value_by_name (function)
// blacklisted: enum_get_value_by_nick (function)
// blacklisted: enum_register_static (function)
// blacklisted: flags_complete_type_info (function)
// blacklisted: flags_get_first_value (function)
// blacklisted: flags_get_value_by_name (function)
// blacklisted: flags_get_value_by_nick (function)
// blacklisted: flags_register_static (function)
// blacklisted: gtype_get_type (function)
// blacklisted: param_spec_boolean (function)
// blacklisted: param_spec_boxed (function)
// blacklisted: param_spec_char (function)
// blacklisted: param_spec_double (function)
// blacklisted: param_spec_enum (function)
// blacklisted: param_spec_flags (function)
// blacklisted: param_spec_float (function)
// blacklisted: param_spec_gtype (function)
// blacklisted: param_spec_int (function)
// blacklisted: param_spec_int64 (function)
// blacklisted: param_spec_long (function)
// blacklisted: param_spec_object (function)
// blacklisted: param_spec_param (function)
// blacklisted: param_spec_pointer (function)
// blacklisted: param_spec_pool_new (function)
// blacklisted: param_spec_string (function)
// blacklisted: param_spec_uchar (function)
// blacklisted: param_spec_uint (function)
// blacklisted: param_spec_uint64 (function)
// blacklisted: param_spec_ulong (function)
// blacklisted: param_spec_unichar (function)
// blacklisted: param_spec_variant (function)
// blacklisted: param_type_register_static (function)
// blacklisted: param_value_convert (function)
// blacklisted: param_value_defaults (function)
// blacklisted: param_value_set_default (function)
// blacklisted: param_value_validate (function)
// blacklisted: param_values_cmp (function)
// blacklisted: pointer_type_register_static (function)
// blacklisted: signal_accumulator_first_wins (function)
// blacklisted: signal_accumulator_true_handled (function)
// blacklisted: signal_add_emission_hook (function)
// blacklisted: signal_chain_from_overridden (function)
// blacklisted: signal_connect_closure (function)
// blacklisted: signal_connect_closure_by_id (function)
// blacklisted: signal_emitv (function)
// blacklisted: signal_get_invocation_hint (function)
// blacklisted: signal_handler_block (function)
// blacklisted: signal_handler_disconnect (function)
// blacklisted: signal_handler_find (function)
// blacklisted: signal_handler_is_connected (function)
// blacklisted: signal_handler_unblock (function)
// blacklisted: signal_handlers_block_matched (function)
// blacklisted: signal_handlers_destroy (function)
// blacklisted: signal_handlers_disconnect_matched (function)
// blacklisted: signal_handlers_unblock_matched (function)
// blacklisted: signal_has_handler_pending (function)
// blacklisted: signal_list_ids (function)
// blacklisted: signal_lookup (function)
// blacklisted: signal_name (function)
// blacklisted: signal_override_class_closure (function)
// blacklisted: signal_parse_name (function)
// blacklisted: signal_query (function)
// blacklisted: signal_remove_emission_hook (function)
// blacklisted: signal_set_va_marshaller (function)
// blacklisted: signal_stop_emission (function)
// blacklisted: signal_stop_emission_by_name (function)
// blacklisted: signal_type_cclosure_new (function)
// blacklisted: source_set_closure (function)
// blacklisted: source_set_dummy_callback (function)
// blacklisted: strdup_value_contents (function)
// blacklisted: type_add_class_private (function)
// blacklisted: type_add_instance_private (function)
// blacklisted: type_add_interface_dynamic (function)
// blacklisted: type_add_interface_static (function)
// blacklisted: type_check_class_is_a (function)
// blacklisted: type_check_instance (function)
// blacklisted: type_check_instance_is_a (function)
// blacklisted: type_check_is_value_type (function)
// blacklisted: type_check_value (function)
// blacklisted: type_check_value_holds (function)
// blacklisted: type_children (function)
// blacklisted: type_class_add_private (function)
// blacklisted: type_class_adjust_private_offset (function)
// blacklisted: type_class_peek (function)
// blacklisted: type_class_peek_static (function)
// blacklisted: type_class_ref (function)
// blacklisted: type_default_interface_peek (function)
// blacklisted: type_default_interface_ref (function)
// blacklisted: type_default_interface_unref (function)
// blacklisted: type_depth (function)
// blacklisted: type_ensure (function)
// blacklisted: type_free_instance (function)
// blacklisted: type_from_name (function)
// blacklisted: type_fundamental (function)
// blacklisted: type_fundamental_next (function)
// blacklisted: type_get_plugin (function)
// blacklisted: type_get_qdata (function)
// blacklisted: type_get_type_registration_serial (function)
// blacklisted: type_init (function)
// blacklisted: type_init_with_debug_flags (function)
// blacklisted: type_interface_add_prerequisite (function)
// blacklisted: type_interface_get_plugin (function)
// blacklisted: type_interface_peek (function)
// blacklisted: type_interface_prerequisites (function)
// blacklisted: type_interfaces (function)
// blacklisted: type_is_a (function)
// blacklisted: type_name (function)
// blacklisted: type_name_from_class (function)
// blacklisted: type_name_from_instance (function)
// blacklisted: type_next_base (function)
// blacklisted: type_parent (function)
// blacklisted: type_qname (function)
// blacklisted: type_query (function)
// blacklisted: type_register_dynamic (function)
// blacklisted: type_register_fundamental (function)
// blacklisted: type_register_static (function)
// blacklisted: type_set_qdata (function)
// blacklisted: type_test_flags (function)
// blacklisted: value_type_compatible (function)
// blacklisted: value_type_transformable (function)


//--------------------------------------------------------------
// Holder
//--------------------------------------------------------------
// holy crap, what am I doing here..

type holder_key [2]unsafe.Pointer
type holder_type map[holder_key]int

var Holder = holder_type(make(map[holder_key]int))

func (this holder_type) Grab(x interface{}) {
	if x == nil {
		return
	}

	key := *(*holder_key)(unsafe.Pointer(&x))
	count := this[key]
	this[key] = count + 1
}

func (this holder_type) Release(x interface{}) {
	if x == nil {
		return
	}

	key := *(*holder_key)(unsafe.Pointer(&x))
	count := this[key]
	if count <= 1 {
		delete(this, key)
	} else {
		this[key] = count - 1
	}
}

//--------------------------------------------------------------
// FinalizerQueue
//--------------------------------------------------------------

type finalizer_item struct {
	ptr unsafe.Pointer
	finalizer func(unsafe.Pointer)
}

type fqueue_type struct {
	sync.Mutex
	queue []finalizer_item
	exec_queue []finalizer_item
	tid uint32
}

var FQueue fqueue_type

func (this *fqueue_type) Start(interval int) {
	this.Lock()
	this.queue = make([]finalizer_item, 0, 50)
	this.exec_queue = make([]finalizer_item, 50)
	this.tid = uint32(C._g_timeout_add_fqueue(C.uint32_t(interval)))
	this.Unlock()
}

func (this *fqueue_type) Stop() {
	this.Lock()
	// TODO: we'll discard few items here at Stop, is it ok?
	this.queue = nil
	C.g_source_remove(C.uint32_t(this.tid))
	this.Unlock()
}

// returns true if the item was enqueued, thread safe
func (this *fqueue_type) Push(ptr unsafe.Pointer, finalizer func(unsafe.Pointer)) bool {
	this.Lock()
	if this.queue != nil {
		this.queue = append(this.queue, finalizer_item{ptr, finalizer})
		this.Unlock()
		return true
	}
	this.Unlock()
	return false
}

// exec is only thread safe if executed by a single thread
func (this *fqueue_type) exec() {
	// exec_queue is used for not holding the lock a lot
	this.Lock()
	// common case
	if len(this.queue) == 0 {
		this.Unlock()
		return
	}

	// non-empty queue, copy everything to exec_queue
	if len(this.queue) > len(this.exec_queue) {
		this.exec_queue = make([]finalizer_item, len(this.queue))
	}
	nitems := copy(this.exec_queue, this.queue)
	this.queue = this.queue[:0]
	this.Unlock()

	// then do our work
	for i := 0; i < nitems; i++ {
		this.exec_queue[i].finalizer(this.exec_queue[i].ptr)
		this.exec_queue[i] = finalizer_item{}
	}
}

//export fqueue_dispatcher
func fqueue_dispatcher(unused unsafe.Pointer) int32 {
	FQueue.exec()
	return 1
}

//--------------------------------------------------------------
// NilString
//--------------------------------------------------------------

// its value will stay the same forever, use the value directly if you like
const NilString = "\x00"

//--------------------------------------------------------------
// Quark
//
// TODO: probably it's a temporary place for this, quarks are
// from glib
//--------------------------------------------------------------

type Quark uint32

func NewQuarkFromString(s string) Quark {
	cs := C.CString(s)
	quark := C.g_quark_from_string(cs)
	C.free(unsafe.Pointer(cs))
	return Quark(quark)
}

// we use this one to store Go's representation of the GObject
// as user data in that GObject once it was allocated. For the
// sake of avoiding allocations.
var go_repr Quark

func init() {
	go_repr = NewQuarkFromString("go-representation")
}

//--------------------------------------------------------------
// ParamSpec utils
//--------------------------------------------------------------

// Let's implement these manually (not Object based and small amount of things
// to implement).

// First some utils
func param_spec_finalizer(pspec *ParamSpec) {
	if FQueue.Push(unsafe.Pointer(pspec), param_spec_finalizer2) {
		return
	}
	C.g_param_spec_unref((*C.GParamSpec)(pspec.C))
}

func param_spec_finalizer2(pspec_un unsafe.Pointer) {
	pspec := (*ParamSpec)(pspec_un)
	C.g_param_spec_unref((*C.GParamSpec)(pspec.C))
}

func set_param_spec_finalizer(pspec *ParamSpec) {
	runtime.SetFinalizer(pspec, param_spec_finalizer)
}

func ParamSpecGrabIfType(c unsafe.Pointer, t Type) unsafe.Pointer {
	if c == nil {
		return nil
	}
	obj := &ParamSpec{c}
	if obj.GetType().IsA(t) {
		C.g_param_spec_ref_sink((*C.GParamSpec)(obj.C))
		set_param_spec_finalizer(obj)
		return unsafe.Pointer(obj)
	}
	return nil
}

func ParamSpecWrap(c unsafe.Pointer, grab bool) unsafe.Pointer {
	if c == nil {
		return nil
	}
	obj := &ParamSpec{c}
	if grab {
		C.g_param_spec_ref_sink((*C.GParamSpec)(obj.C))
	}
	set_param_spec_finalizer(obj)
	return unsafe.Pointer(obj)
}

//--------------------------------------------------------------
// ParamSpec
//--------------------------------------------------------------

type ParamSpecLike interface {
	InheritedFromGParamSpec() *C.GParamSpec
}

type ParamSpec struct {
	C unsafe.Pointer
}

func ToParamSpec(pspeclike ParamSpecLike) *ParamSpec {
	t := (*ParamSpec)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpec()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpec)(obj)
	}
	panic("cannot cast to ParamSpec")
}

func (this *ParamSpec) InheritedFromGParamSpec() *C.GParamSpec {
	return (*C.GParamSpec)(this.C)
}

func (this *ParamSpec) GetStaticType() Type {
	return Type(C._g_type_param())
}

func (this *ParamSpec) GetType() Type {
	return Type(C._g_param_spec_type(this.InheritedFromGParamSpec()))
}

func (this *ParamSpec) GetValueType() Type {
	return Type(C._g_param_spec_value_type(this.InheritedFromGParamSpec()))
}

//--------------------------------------------------------------
// ParamSpecBoolean
//--------------------------------------------------------------

type ParamSpecBooleanLike interface {
	InheritedFromGParamSpecBoolean() *C.GParamSpecBoolean
}

type ParamSpecBoolean struct {
	ParamSpec
}

func ToParamSpecBoolean(pspeclike ParamSpecBooleanLike) *ParamSpecBoolean {
	t := (*ParamSpecBoolean)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecBoolean()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecBoolean)(obj)
	}
	panic("cannot cast to ParamSpecBoolean")
}

func (this *ParamSpecBoolean) InheritedFromGParamSpecBoolean() *C.GParamSpecBoolean {
	return (*C.GParamSpecBoolean)(this.C)
}

func (this *ParamSpecBoolean) GetStaticType() Type {
	return Type(C._g_type_param_boolean())
}

//--------------------------------------------------------------
// ParamSpecBoxed
//--------------------------------------------------------------

type ParamSpecBoxedLike interface {
	InheritedFromGParamSpecBoxed() *C.GParamSpecBoxed
}

type ParamSpecBoxed struct {
	ParamSpec
}

func ToParamSpecBoxed(pspeclike ParamSpecBoxedLike) *ParamSpecBoxed {
	t := (*ParamSpecBoxed)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecBoxed()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecBoxed)(obj)
	}
	panic("cannot cast to ParamSpecBoxed")
}

func (this *ParamSpecBoxed) InheritedFromGParamSpecBoxed() *C.GParamSpecBoxed {
	return (*C.GParamSpecBoxed)(this.C)
}

func (this *ParamSpecBoxed) GetStaticType() Type {
	return Type(C._g_type_param_boxed())
}

//--------------------------------------------------------------
// ParamSpecChar
//--------------------------------------------------------------

type ParamSpecCharLike interface {
	InheritedFromGParamSpecChar() *C.GParamSpecChar
}

type ParamSpecChar struct {
	ParamSpec
}

func ToParamSpecChar(pspeclike ParamSpecCharLike) *ParamSpecChar {
	t := (*ParamSpecChar)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecChar()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecChar)(obj)
	}
	panic("cannot cast to ParamSpecChar")
}

func (this *ParamSpecChar) InheritedFromGParamSpecChar() *C.GParamSpecChar {
	return (*C.GParamSpecChar)(this.C)
}

func (this *ParamSpecChar) GetStaticType() Type {
	return Type(C._g_type_param_char())
}

//--------------------------------------------------------------
// ParamSpecDouble
//--------------------------------------------------------------

type ParamSpecDoubleLike interface {
	InheritedFromGParamSpecDouble() *C.GParamSpecDouble
}

type ParamSpecDouble struct {
	ParamSpec
}

func ToParamSpecDouble(pspeclike ParamSpecDoubleLike) *ParamSpecDouble {
	t := (*ParamSpecDouble)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecDouble()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecDouble)(obj)
	}
	panic("cannot cast to ParamSpecDouble")
}

func (this *ParamSpecDouble) InheritedFromGParamSpecDouble() *C.GParamSpecDouble {
	return (*C.GParamSpecDouble)(this.C)
}

func (this *ParamSpecDouble) GetStaticType() Type {
	return Type(C._g_type_param_double())
}

//--------------------------------------------------------------
// ParamSpecEnum
//--------------------------------------------------------------

type ParamSpecEnumLike interface {
	InheritedFromGParamSpecEnum() *C.GParamSpecEnum
}

type ParamSpecEnum struct {
	ParamSpec
}

func ToParamSpecEnum(pspeclike ParamSpecEnumLike) *ParamSpecEnum {
	t := (*ParamSpecEnum)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecEnum()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecEnum)(obj)
	}
	panic("cannot cast to ParamSpecEnum")
}

func (this *ParamSpecEnum) InheritedFromGParamSpecEnum() *C.GParamSpecEnum {
	return (*C.GParamSpecEnum)(this.C)
}

func (this *ParamSpecEnum) GetStaticType() Type {
	return Type(C._g_type_param_enum())
}

//--------------------------------------------------------------
// ParamSpecFlags
//--------------------------------------------------------------

type ParamSpecFlagsLike interface {
	InheritedFromGParamSpecFlags() *C.GParamSpecFlags
}

type ParamSpecFlags struct {
	ParamSpec
}

func ToParamSpecFlags(pspeclike ParamSpecFlagsLike) *ParamSpecFlags {
	t := (*ParamSpecFlags)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecFlags()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecFlags)(obj)
	}
	panic("cannot cast to ParamSpecFlags")
}

func (this *ParamSpecFlags) InheritedFromGParamSpecFlags() *C.GParamSpecFlags {
	return (*C.GParamSpecFlags)(this.C)
}

func (this *ParamSpecFlags) GetStaticType() Type {
	return Type(C._g_type_param_flags())
}

//--------------------------------------------------------------
// ParamSpecFloat
//--------------------------------------------------------------

type ParamSpecFloatLike interface {
	InheritedFromGParamSpecFloat() *C.GParamSpecFloat
}

type ParamSpecFloat struct {
	ParamSpec
}

func ToParamSpecFloat(pspeclike ParamSpecFloatLike) *ParamSpecFloat {
	t := (*ParamSpecFloat)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecFloat()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecFloat)(obj)
	}
	panic("cannot cast to ParamSpecFloat")
}

func (this *ParamSpecFloat) InheritedFromGParamSpecFloat() *C.GParamSpecFloat {
	return (*C.GParamSpecFloat)(this.C)
}

func (this *ParamSpecFloat) GetStaticType() Type {
	return Type(C._g_type_param_float())
}

//--------------------------------------------------------------
// ParamSpecGType
//--------------------------------------------------------------

type ParamSpecGTypeLike interface {
	InheritedFromGParamSpecGType() *C.GParamSpecGType
}

type ParamSpecGType struct {
	ParamSpec
}

func ToParamSpecGType(pspeclike ParamSpecGTypeLike) *ParamSpecGType {
	t := (*ParamSpecGType)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecGType()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecGType)(obj)
	}
	panic("cannot cast to ParamSpecGType")
}

func (this *ParamSpecGType) InheritedFromGParamSpecGType() *C.GParamSpecGType {
	return (*C.GParamSpecGType)(this.C)
}

func (this *ParamSpecGType) GetStaticType() Type {
	return Type(C._g_type_param_gtype())
}

//--------------------------------------------------------------
// ParamSpecInt
//--------------------------------------------------------------

type ParamSpecIntLike interface {
	InheritedFromGParamSpecInt() *C.GParamSpecInt
}

type ParamSpecInt struct {
	ParamSpec
}

func ToParamSpecInt(pspeclike ParamSpecIntLike) *ParamSpecInt {
	t := (*ParamSpecInt)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecInt()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecInt)(obj)
	}
	panic("cannot cast to ParamSpecInt")
}

func (this *ParamSpecInt) InheritedFromGParamSpecInt() *C.GParamSpecInt {
	return (*C.GParamSpecInt)(this.C)
}

func (this *ParamSpecInt) GetStaticType() Type {
	return Type(C._g_type_param_int())
}

//--------------------------------------------------------------
// ParamSpecInt64
//--------------------------------------------------------------

type ParamSpecInt64Like interface {
	InheritedFromGParamSpecInt64() *C.GParamSpecInt64
}

type ParamSpecInt64 struct {
	ParamSpec
}

func ToParamSpecInt64(pspeclike ParamSpecInt64Like) *ParamSpecInt64 {
	t := (*ParamSpecInt64)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecInt64()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecInt64)(obj)
	}
	panic("cannot cast to ParamSpecInt64")
}

func (this *ParamSpecInt64) InheritedFromGParamSpecInt64() *C.GParamSpecInt64 {
	return (*C.GParamSpecInt64)(this.C)
}

func (this *ParamSpecInt64) GetStaticType() Type {
	return Type(C._g_type_param_int64())
}

//--------------------------------------------------------------
// ParamSpecLong
//--------------------------------------------------------------

type ParamSpecLongLike interface {
	InheritedFromGParamSpecLong() *C.GParamSpecLong
}

type ParamSpecLong struct {
	ParamSpec
}

func ToParamSpecLong(pspeclike ParamSpecLongLike) *ParamSpecLong {
	t := (*ParamSpecLong)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecLong()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecLong)(obj)
	}
	panic("cannot cast to ParamSpecLong")
}

func (this *ParamSpecLong) InheritedFromGParamSpecLong() *C.GParamSpecLong {
	return (*C.GParamSpecLong)(this.C)
}

func (this *ParamSpecLong) GetStaticType() Type {
	return Type(C._g_type_param_long())
}

//--------------------------------------------------------------
// ParamSpecObject
//--------------------------------------------------------------

type ParamSpecObjectLike interface {
	InheritedFromGParamSpecObject() *C.GParamSpecObject
}

type ParamSpecObject struct {
	ParamSpec
}

func ToParamSpecObject(pspeclike ParamSpecObjectLike) *ParamSpecObject {
	t := (*ParamSpecObject)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecObject()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecObject)(obj)
	}
	panic("cannot cast to ParamSpecObject")
}

func (this *ParamSpecObject) InheritedFromGParamSpecObject() *C.GParamSpecObject {
	return (*C.GParamSpecObject)(this.C)
}

func (this *ParamSpecObject) GetStaticType() Type {
	return Type(C._g_type_param_object())
}

//--------------------------------------------------------------
// ParamSpecOverride
//--------------------------------------------------------------

type ParamSpecOverrideLike interface {
	InheritedFromGParamSpecOverride() *C.GParamSpecOverride
}

type ParamSpecOverride struct {
	ParamSpec
}

func ToParamSpecOverride(pspeclike ParamSpecOverrideLike) *ParamSpecOverride {
	t := (*ParamSpecOverride)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecOverride()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecOverride)(obj)
	}
	panic("cannot cast to ParamSpecOverride")
}

func (this *ParamSpecOverride) InheritedFromGParamSpecOverride() *C.GParamSpecOverride {
	return (*C.GParamSpecOverride)(this.C)
}

func (this *ParamSpecOverride) GetStaticType() Type {
	return Type(C._g_type_param_override())
}

//--------------------------------------------------------------
// ParamSpecParam
//--------------------------------------------------------------

type ParamSpecParamLike interface {
	InheritedFromGParamSpecParam() *C.GParamSpecParam
}

type ParamSpecParam struct {
	ParamSpec
}

func ToParamSpecParam(pspeclike ParamSpecParamLike) *ParamSpecParam {
	t := (*ParamSpecParam)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecParam()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecParam)(obj)
	}
	panic("cannot cast to ParamSpecParam")
}

func (this *ParamSpecParam) InheritedFromGParamSpecParam() *C.GParamSpecParam {
	return (*C.GParamSpecParam)(this.C)
}

func (this *ParamSpecParam) GetStaticType() Type {
	return Type(C._g_type_param_param())
}

//--------------------------------------------------------------
// ParamSpecPointer
//--------------------------------------------------------------

type ParamSpecPointerLike interface {
	InheritedFromGParamSpecPointer() *C.GParamSpecPointer
}

type ParamSpecPointer struct {
	ParamSpec
}

func ToParamSpecPointer(pspeclike ParamSpecPointerLike) *ParamSpecPointer {
	t := (*ParamSpecPointer)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecPointer()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecPointer)(obj)
	}
	panic("cannot cast to ParamSpecPointer")
}

func (this *ParamSpecPointer) InheritedFromGParamSpecPointer() *C.GParamSpecPointer {
	return (*C.GParamSpecPointer)(this.C)
}

func (this *ParamSpecPointer) GetStaticType() Type {
	return Type(C._g_type_param_pointer())
}

//--------------------------------------------------------------
// ParamSpecString
//--------------------------------------------------------------

type ParamSpecStringLike interface {
	InheritedFromGParamSpecString() *C.GParamSpecString
}

type ParamSpecString struct {
	ParamSpec
}

func ToParamSpecString(pspeclike ParamSpecStringLike) *ParamSpecString {
	t := (*ParamSpecString)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecString()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecString)(obj)
	}
	panic("cannot cast to ParamSpecString")
}

func (this *ParamSpecString) InheritedFromGParamSpecString() *C.GParamSpecString {
	return (*C.GParamSpecString)(this.C)
}

func (this *ParamSpecString) GetStaticType() Type {
	return Type(C._g_type_param_string())
}

//--------------------------------------------------------------
// ParamSpecUChar
//--------------------------------------------------------------

type ParamSpecUCharLike interface {
	InheritedFromGParamSpecUChar() *C.GParamSpecUChar
}

type ParamSpecUChar struct {
	ParamSpec
}

func ToParamSpecUChar(pspeclike ParamSpecUCharLike) *ParamSpecUChar {
	t := (*ParamSpecUChar)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecUChar()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecUChar)(obj)
	}
	panic("cannot cast to ParamSpecUChar")
}

func (this *ParamSpecUChar) InheritedFromGParamSpecUChar() *C.GParamSpecUChar {
	return (*C.GParamSpecUChar)(this.C)
}

func (this *ParamSpecUChar) GetStaticType() Type {
	return Type(C._g_type_param_uchar())
}

//--------------------------------------------------------------
// ParamSpecUInt
//--------------------------------------------------------------

type ParamSpecUIntLike interface {
	InheritedFromGParamSpecUInt() *C.GParamSpecUInt
}

type ParamSpecUInt struct {
	ParamSpec
}

func ToParamSpecUInt(pspeclike ParamSpecUIntLike) *ParamSpecUInt {
	t := (*ParamSpecUInt)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecUInt()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecUInt)(obj)
	}
	panic("cannot cast to ParamSpecUInt")
}

func (this *ParamSpecUInt) InheritedFromGParamSpecUInt() *C.GParamSpecUInt {
	return (*C.GParamSpecUInt)(this.C)
}

func (this *ParamSpecUInt) GetStaticType() Type {
	return Type(C._g_type_param_uint())
}

//--------------------------------------------------------------
// ParamSpecUInt64
//--------------------------------------------------------------

type ParamSpecUInt64Like interface {
	InheritedFromGParamSpecUInt64() *C.GParamSpecUInt64
}

type ParamSpecUInt64 struct {
	ParamSpec
}

func ToParamSpecUInt64(pspeclike ParamSpecUInt64Like) *ParamSpecUInt64 {
	t := (*ParamSpecUInt64)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecUInt64()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecUInt64)(obj)
	}
	panic("cannot cast to ParamSpecUInt64")
}

func (this *ParamSpecUInt64) InheritedFromGParamSpecUInt64() *C.GParamSpecUInt64 {
	return (*C.GParamSpecUInt64)(this.C)
}

func (this *ParamSpecUInt64) GetStaticType() Type {
	return Type(C._g_type_param_uint64())
}

//--------------------------------------------------------------
// ParamSpecULong
//--------------------------------------------------------------

type ParamSpecULongLike interface {
	InheritedFromGParamSpecULong() *C.GParamSpecULong
}

type ParamSpecULong struct {
	ParamSpec
}

func ToParamSpecULong(pspeclike ParamSpecULongLike) *ParamSpecULong {
	t := (*ParamSpecULong)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecULong()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecULong)(obj)
	}
	panic("cannot cast to ParamSpecULong")
}

func (this *ParamSpecULong) InheritedFromGParamSpecULong() *C.GParamSpecULong {
	return (*C.GParamSpecULong)(this.C)
}

func (this *ParamSpecULong) GetStaticType() Type {
	return Type(C._g_type_param_ulong())
}

//--------------------------------------------------------------
// ParamSpecUnichar
//--------------------------------------------------------------

type ParamSpecUnicharLike interface {
	InheritedFromGParamSpecUnichar() *C.GParamSpecUnichar
}

type ParamSpecUnichar struct {
	ParamSpec
}

func ToParamSpecUnichar(pspeclike ParamSpecUnicharLike) *ParamSpecUnichar {
	t := (*ParamSpecUnichar)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecUnichar()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecUnichar)(obj)
	}
	panic("cannot cast to ParamSpecUnichar")
}

func (this *ParamSpecUnichar) InheritedFromGParamSpecUnichar() *C.GParamSpecUnichar {
	return (*C.GParamSpecUnichar)(this.C)
}

func (this *ParamSpecUnichar) GetStaticType() Type {
	return Type(C._g_type_param_unichar())
}

//--------------------------------------------------------------
// ParamSpecValueArray
//--------------------------------------------------------------

type ParamSpecValueArrayLike interface {
	InheritedFromGParamSpecValueArray() *C.GParamSpecValueArray
}

type ParamSpecValueArray struct {
	ParamSpec
}

func ToParamSpecValueArray(pspeclike ParamSpecValueArrayLike) *ParamSpecValueArray {
	t := (*ParamSpecValueArray)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecValueArray()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecValueArray)(obj)
	}
	panic("cannot cast to ParamSpecValueArray")
}

func (this *ParamSpecValueArray) InheritedFromGParamSpecValueArray() *C.GParamSpecValueArray {
	return (*C.GParamSpecValueArray)(this.C)
}

func (this *ParamSpecValueArray) GetStaticType() Type {
	return Type(C._g_type_param_value_array())
}

//--------------------------------------------------------------
// ParamSpecVariant
//--------------------------------------------------------------

type ParamSpecVariantLike interface {
	InheritedFromGParamSpecVariant() *C.GParamSpecVariant
}

type ParamSpecVariant struct {
	ParamSpec
}

func ToParamSpecVariant(pspeclike ParamSpecVariantLike) *ParamSpecVariant {
	t := (*ParamSpecVariant)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecVariant()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecVariant)(obj)
	}
	panic("cannot cast to ParamSpecVariant")
}

func (this *ParamSpecVariant) InheritedFromGParamSpecVariant() *C.GParamSpecVariant {
	return (*C.GParamSpecVariant)(this.C)
}

func (this *ParamSpecVariant) GetStaticType() Type {
	return Type(C._g_type_param_variant())
}

//--------------------------------------------------------------
// Object
//--------------------------------------------------------------

func object_finalizer(obj *Object) {
	if FQueue.Push(unsafe.Pointer(obj), object_finalizer2) {
		return
	}
	C.g_object_set_qdata((*C.GObject)(obj.C), C.uint32_t(go_repr), nil)
	C.g_object_unref((*C.GObject)(obj.C))
}

func object_finalizer2(obj_un unsafe.Pointer) {
	obj := (*Object)(obj_un)
	C.g_object_set_qdata((*C.GObject)(obj.C), C.uint32_t(go_repr), nil)
	C.g_object_unref((*C.GObject)(obj.C))
}

func set_object_finalizer(obj *Object) {
	runtime.SetFinalizer(obj, object_finalizer)
}

func ObjectWrap(c unsafe.Pointer, grab bool) unsafe.Pointer {
	if c == nil {
		return nil
	}
	obj := (*Object)(C.g_object_get_qdata((*C.GObject)(c), C.uint32_t(go_repr)))
	if obj != nil {
		return unsafe.Pointer(obj)
	}
	obj = &Object{c}
	if grab {
		C.g_object_ref_sink((*C.GObject)(obj.C))
	}
	set_object_finalizer(obj)
	C.g_object_set_qdata((*C.GObject)(obj.C),
		C.uint32_t(go_repr), unsafe.Pointer(obj))
	return unsafe.Pointer(obj)
}

func ObjectGrabIfType(c unsafe.Pointer, t Type) unsafe.Pointer {
	if c == nil {
		return nil
	}
	hasrepr := true
	obj := (*Object)(C.g_object_get_qdata((*C.GObject)(c), C.uint32_t(go_repr)))
	if obj == nil {
		obj = &Object{c}
		hasrepr = false
	}
	if obj.GetType().IsA(t) {
		if !hasrepr {
			C.g_object_ref_sink((*C.GObject)(obj.C))
			set_object_finalizer(obj)
			C.g_object_set_qdata((*C.GObject)(obj.C),
				C.uint32_t(go_repr), unsafe.Pointer(obj))
		}
		return unsafe.Pointer(obj)
	}
	return nil
}

func (this *Object) GetType() Type {
	return Type(C._g_object_type((*C.GObject)(this.C)))
}

func (this *Object) Connect(signal string, clo interface{}) {
	csignal := C.CString(signal)
	Holder.Grab(clo)
	goclosure := C.g_goclosure_new(unsafe.Pointer(&clo), nil)
	C.g_signal_connect_closure((*C.GObject)(this.C), csignal, (*C.GClosure)(unsafe.Pointer(goclosure)), 0)
	C.free(unsafe.Pointer(csignal))
}

func (this *Object) ConnectMethod(signal string, clo interface{}, recv interface{}) {
	csignal := C.CString(signal)
	Holder.Grab(clo)
	Holder.Grab(recv)
	goclosure := C.g_goclosure_new(unsafe.Pointer(&clo), unsafe.Pointer(&recv))
	C.g_signal_connect_closure((*C.GObject)(this.C), csignal, (*C.GClosure)(unsafe.Pointer(goclosure)), 0)
	C.free(unsafe.Pointer(csignal))

}

func (this *Object) FindProperty(name string) *ParamSpec {
	cname := C.CString(name)
	ret := C._g_object_find_property(this.InheritedFromGObject(), cname)
	C.free(unsafe.Pointer(cname))
	return (*ParamSpec)(ParamSpecWrap(unsafe.Pointer(ret), true))
}

func (this *Object) SetProperty(name string, value interface{}) {
	cname := C.CString(name)
	pspec := this.FindProperty(name)
	if pspec == nil {
		panic("Object has no property with that name: " + name)
	}
	var gvalue Value
	gvalue.Init(pspec.GetValueType())
	gvalue.SetGoInterface(value)
	C.g_object_set_property(this.InheritedFromGObject(), cname,
		(*C.GValue)(unsafe.Pointer(&gvalue)))
	gvalue.Unset()
	C.free(unsafe.Pointer(cname))
}

func (this *Object) GetProperty(name string, value interface{}) {
	cname := C.CString(name)
	pspec := this.FindProperty(name)
	if pspec == nil {
		panic("Object has no property with that name: " + name)
	}
	var gvalue Value
	gvalue.Init(pspec.GetValueType())
	C.g_object_get_property(this.InheritedFromGObject(), cname,
		(*C.GValue)(unsafe.Pointer(&gvalue)))
	gvalue.GetGoInterface(value)
	gvalue.Unset()
	C.free(unsafe.Pointer(cname))
}

func ObjectBindProperty(source ObjectLike, source_property string, target ObjectLike, target_property string, flags BindingFlags) *Binding {
	csource_property := C.CString(source_property)
	ctarget_property := C.CString(target_property)
	obj := C.g_object_bind_property(
		source.InheritedFromGObject(), csource_property,
		target.InheritedFromGObject(), ctarget_property,
		C.GBindingFlags(flags))
	C.free(unsafe.Pointer(csource_property))
	C.free(unsafe.Pointer(ctarget_property))
	return (*Binding)(ObjectWrap(unsafe.Pointer(obj), true))
}

func (this *Object) Unref() {
	runtime.SetFinalizer(this, nil)
	C.g_object_set_qdata((*C.GObject)(this.C), C.uint32_t(go_repr), nil)
	C.g_object_unref((*C.GObject)(this.C))
	this.C = nil
}

//--------------------------------------------------------------
// Closures
//--------------------------------------------------------------

//export g_goclosure_finalize_go
func g_goclosure_finalize_go(goclosure_up unsafe.Pointer) {
	goclosure := (*C.GGoClosure)(goclosure_up)
	clo := *(*interface{})(C.g_goclosure_get_func(goclosure))
	recv := *(*interface{})(C.g_goclosure_get_recv(goclosure))
	Holder.Release(clo)
	Holder.Release(recv)
}

//export g_goclosure_marshal_go
func g_goclosure_marshal_go(goclosure_up, ret_up unsafe.Pointer, nargs int32, args_up unsafe.Pointer) {
	var callargs [20]reflect.Value
	var recv reflect.Value
	goclosure := (*C.GGoClosure)(goclosure_up)
	ret := (*Value)(ret_up)
	args := (*(*[alot]Value)(args_up))[:nargs]
	f := reflect.ValueOf(*(*interface{})(C.g_goclosure_get_func(goclosure)))
	ft := f.Type()
	callargsn := ft.NumIn()

	recvi := *(*interface{})(C.g_goclosure_get_recv(goclosure))
	if recvi != nil {
		recv = reflect.ValueOf(recvi)
	}

	if callargsn >= 20 {
		panic("too many arguments in a closure")
	}

	for i, n := 0, callargsn; i < n; i++ {
		idx := i
		if recvi != nil {
			idx--
			if i == 0 {
				callargs[i] = recv
				continue
			}
		}

		in := ft.In(i)

		// use default value, if there is not enough args
		if len(args) <= idx {
			callargs[i] = reflect.New(in).Elem()
			continue
		}

		v := args[idx].GetGoValue(in)
		callargs[i] = v
	}

	out := f.Call(callargs[:callargsn])
	if len(out) == 1 {
		ret.SetGoValue(out[0])
	}
}

//--------------------------------------------------------------
// Go Interface boxed type
//--------------------------------------------------------------

//export g_go_interface_copy_go
func g_go_interface_copy_go(boxed unsafe.Pointer) unsafe.Pointer {
	Holder.Grab(*(*interface{})(boxed))
	newboxed := C.malloc(C.size_t(unsafe.Sizeof([2]unsafe.Pointer{})))
	C.memcpy(newboxed, boxed, C.size_t(unsafe.Sizeof([2]unsafe.Pointer{})))
	return newboxed
}

//export g_go_interface_free_go
func g_go_interface_free_go(boxed unsafe.Pointer) {
	Holder.Release(*(*interface{})(boxed))
	C.free(boxed)
}

//--------------------------------------------------------------
// Type
//--------------------------------------------------------------

type Type C.GType

func (this Type) IsA(other Type) bool {
	return C.g_type_is_a(C.GType(this), C.GType(other)) != 0
}

func (this Type) String() string {
	cname := C.g_type_name(C.GType(this))
	if cname == nil {
		return ""
	}
	return C.GoString(cname)
}

func (this Type) asC() C.GType {
	return C.GType(this)
}

var (
	Interface Type
	Char Type
	UChar Type
	Boolean Type
	Int Type
	UInt Type
	Long Type
	ULong Type
	Int64 Type
	UInt64 Type
	Enum Type
	Flags Type
	Float Type
	Double Type
	String Type
	Pointer Type
	Boxed Type
	Param Type
	GObject Type
	GType Type
	Variant Type
	GoInterface Type
)

func init() {
	C.g_type_init()

	Interface = Type(C._g_type_interface())
	Char = Type(C._g_type_char())
	UChar = Type(C._g_type_uchar())
	Boolean = Type(C._g_type_boolean())
	Int = Type(C._g_type_int())
	UInt = Type(C._g_type_uint())
	Long = Type(C._g_type_long())
	ULong = Type(C._g_type_ulong())
	Int64 = Type(C._g_type_int64())
	UInt64 = Type(C._g_type_uint64())
	Enum = Type(C._g_type_enum())
	Flags = Type(C._g_type_flags())
	Float = Type(C._g_type_float())
	Double = Type(C._g_type_double())
	String = Type(C._g_type_string())
	Pointer = Type(C._g_type_pointer())
	Boxed = Type(C._g_type_boxed())
	Param = Type(C._g_type_param())
	GObject = Type(C._g_type_object())
	GType = Type(C._g_type_gtype())
	Variant = Type(C._g_type_variant())
	GoInterface = Type(C._g_type_go_interface())
}

// Every GObject generated by this generator implements this interface
// and it must work even if the receiver is a nil value
type StaticTyper interface {
	GetStaticType() Type
}

//--------------------------------------------------------------
// Value
//--------------------------------------------------------------

func (this *Value) asC() *C.GValue {
	return (*C.GValue)(unsafe.Pointer(this))
}

// g_value_init
func (this *Value) Init(t Type) {
	C.g_value_init(this.asC(), t.asC())
}

// g_value_copy
func (this *Value) Set(src *Value) {
	C.g_value_copy(src.asC(), this.asC())
}

// g_value_reset
func (this *Value) Reset() {
	C.g_value_reset(this.asC())
}

// g_value_unset
func (this *Value) Unset() {
	C.g_value_unset(this.asC())
}

// G_VALUE_TYPE
func (this *Value) GetType() Type {
	return Type(C._g_value_type(this.asC()))
}

// g_value_type_compatible
func ValueTypeCompatible(src, dst Type) bool {
	return C.g_value_type_compatible(src.asC(), dst.asC()) != 0
}

// g_value_type_transformable
func ValueTypeTransformable(src, dst Type) bool {
	return C.g_value_type_transformable(src.asC(), dst.asC()) != 0
}

// g_value_transform
func (this *Value) Transform(src *Value) bool {
	return C.g_value_transform(src.asC(), this.asC()) != 0
}

// g_value_get_boolean
func (this *Value) GetBool() bool {
	return C.g_value_get_boolean(this.asC()) != 0
}

// g_value_set_boolean
func (this *Value) SetBool(v bool) {
	C.g_value_set_boolean(this.asC(), _GoBoolToCBool(v))
}

// g_value_get_int64
func (this *Value) GetInt() int64 {
	return int64(C.g_value_get_int64(this.asC()))
}

// g_value_set_int64
func (this *Value) SetInt(v int64) {
	C.g_value_set_int64(this.asC(), C.int64_t(v))
}

// g_value_get_uint64
func (this *Value) GetUint() uint64 {
	return uint64(C.g_value_get_uint64(this.asC()))
}

// g_value_set_uint64
func (this *Value) SetUint(v uint64) {
	C.g_value_set_uint64(this.asC(), C.uint64_t(v))
}

// g_value_get_double
func (this *Value) GetFloat() float64 {
	return float64(C.g_value_get_double(this.asC()))
}

// g_value_set_double
func (this *Value) SetFloat(v float64) {
	C.g_value_set_double(this.asC(), C.double(v))
}

// g_value_get_string
func (this *Value) GetString() string {
	return C.GoString(C.g_value_get_string(this.asC()))
}

// g_value_take_string
func (this *Value) SetString(v string) {
	cstr := C.CString(v)
	C.g_value_take_string(this.asC(), cstr)
	// not freeing, because GValue takes the ownership
}

// g_value_get_object
func (this *Value) GetObject() unsafe.Pointer {
	return unsafe.Pointer(C.g_value_get_object(this.asC()))
}

// g_value_set_object
func (this *Value) SetObject(x unsafe.Pointer) {
	C.g_value_set_object(this.asC(), (*C.GObject)(x))
}

// g_value_get_boxed
func (this *Value) GetBoxed() unsafe.Pointer {
	return C.g_value_get_boxed(this.asC())
}

// g_value_take_boxed
func (this *Value) SetBoxed(x unsafe.Pointer) {
	C.g_value_take_boxed(this.asC(), x)
}

func (this *Value) GetBoxedInterface() interface{} {
	return *(*interface{})(C.g_value_get_boxed(this.asC()))
}

func (this *Value) SetBoxedInterface(x interface{}) {
	Holder.Grab(x)
	newboxed := C.malloc(C.size_t(unsafe.Sizeof([2]unsafe.Pointer{})))
	C.memcpy(newboxed, unsafe.Pointer(&x), C.size_t(unsafe.Sizeof([2]unsafe.Pointer{})))
	C.g_value_take_boxed(this.asC(), newboxed)
}

//--------------------------------------------------------------
// A giant glue for connecting GType and Go's reflection
//--------------------------------------------------------------

var statictyper = reflect.TypeOf((*StaticTyper)(nil)).Elem()
var objectlike = reflect.TypeOf((*ObjectLike)(nil)).Elem()

func (this *Value) SetGoValue(v reflect.Value) {
	valuetype := this.GetType()
	var src Value

	if valuetype == GoInterface {
		// special case
		this.SetBoxedInterface(v.Interface())
		return
	}

	transform := func() {
		ok := this.Transform(&src)
		if !ok {
			panic("Go value (" + v.Type().String() + ") is not transformable to " + valuetype.String())
		}
	}

	switch v.Kind() {
	case reflect.Bool:
		src.Init(Boolean)
		src.SetBool(v.Bool())
		transform()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		src.Init(Int64)
		src.SetInt(v.Int())
		transform()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		src.Init(UInt64)
		src.SetUint(v.Uint())
		transform()
	case reflect.Float32, reflect.Float64:
		src.Init(Double)
		src.SetFloat(v.Float())
		transform()
	case reflect.String:
		src.Init(String)
		src.SetString(v.String())
		transform()
		src.Unset()
	case reflect.Ptr:
		gotype := v.Type()
		src.Init(GObject)
		if gotype.Implements(objectlike) {
			obj, ok := v.Interface().(ObjectLike)
			if !ok {
				panic(gotype.String() + " is not transformable to GValue")
			}

			src.SetObject(unsafe.Pointer(obj.InheritedFromGObject()))
			transform()
		}
		src.Unset()
	}
}

var CairoMarshaler func(*Value, reflect.Type) (reflect.Value, bool)

func (this *Value) GetGoValue(t reflect.Type) reflect.Value {
	var out reflect.Value
	var dst Value

	if (this.GetType() == GoInterface) {
		return reflect.ValueOf(this.GetBoxedInterface())
	}

	transform := func() {
		ok := dst.Transform(this)
		if !ok {
			panic("GValue is not transformable to " + t.String())
		}
	}

	switch t.Kind() {
	case reflect.Bool:
		dst.Init(Boolean)
		transform()
		out = reflect.New(t).Elem()
		out.SetBool(dst.GetBool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		dst.Init(Int64)
		transform()
		out = reflect.New(t).Elem()
		out.SetInt(dst.GetInt())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		dst.Init(UInt64)
		transform()
		out = reflect.New(t).Elem()
		out.SetUint(dst.GetUint())
	case reflect.Float32, reflect.Float64:
		dst.Init(Double)
		transform()
		out = reflect.New(t).Elem()
		out.SetFloat(dst.GetFloat())
	case reflect.String:
		dst.Init(String)
		transform()
		out = reflect.New(t).Elem()
		out.SetString(dst.GetString())
		dst.Unset() // need to clean up in this case
	case reflect.Ptr:
		if t.Implements(objectlike) {
			// at this point we're sure that this is a pointer to the ObjectLike
			out = reflect.New(t)
			st, ok := out.Elem().Interface().(StaticTyper)
			if !ok {
				panic("ObjectLike type must implement StaticTyper as well")
			}
			dst.Init(st.GetStaticType())
			transform()
			*(*unsafe.Pointer)(unsafe.Pointer(out.Pointer())) = ObjectWrap(dst.GetObject(), true)
			dst.Unset()
			out = out.Elem()
		} else {
			// cairo marshaler hook
			if CairoMarshaler != nil {
				var ok bool
				out, ok = CairoMarshaler(this, t)
				if ok {
					break
				}
			}

			// must be a struct then
			out = reflect.New(t)
			*(*unsafe.Pointer)(unsafe.Pointer(out.Pointer())) = this.GetBoxed()
			out = out.Elem()
		}
	}
	return out
}

func (this *Value) SetGoInterface(v interface{}) {
	this.SetGoValue(reflect.ValueOf(v))
}

func (this *Value) GetGoInterface(v interface{}) {
	vp := reflect.ValueOf(v)
	if vp.Kind() != reflect.Ptr {
		panic("a pointer to value is expected for Value.GetGoInterface")
	}
	vp.Elem().Set(this.GetGoValue(vp.Type().Elem()))
}