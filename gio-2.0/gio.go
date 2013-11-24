package gio

/*
#include "gio.gen.h"

extern GObject *g_object_ref_sink(GObject*);
extern void g_object_unref(GObject*);
extern void g_error_free(GError*);
extern void g_free(void*);
#cgo pkg-config: gio-2.0
*/
import "C"
import "unsafe"
import "errors"

import (
	"dlib/gobject-2.0"
	"dlib/glib-2.0"
)

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


//export _Gio_go_callback_cleanup
func _Gio_go_callback_cleanup(gofunc unsafe.Pointer) {
	gobject.Holder.Release(gofunc)
}


type ActionLike interface {
	ImplementsGAction() *C.GAction
}

type Action struct {
	gobject.Object
	ActionImpl
}

type ActionImpl struct {}

func ToAction(objlike gobject.ObjectLike) *Action {
	t := (*ActionImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Action)(obj)
	}
	panic("cannot cast to Action")
}

func (this0 *ActionImpl) ImplementsGAction() *C.GAction {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GAction)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *ActionImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.g_action_get_type())
}

func ActionGetType() gobject.Type {
	return (*ActionImpl)(nil).GetStaticType()
}
func (this0 *ActionImpl) Activate(parameter0 *glib.Variant) {
	var this1 *C.GAction
	var parameter1 *C.GVariant
	if this0 != nil {
		this1 = this0.ImplementsGAction()
	}
	parameter1 = (*C.GVariant)(unsafe.Pointer(parameter0))
	C.g_action_activate(this1, parameter1)
}
func (this0 *ActionImpl) ChangeState(value0 *glib.Variant) {
	var this1 *C.GAction
	var value1 *C.GVariant
	if this0 != nil {
		this1 = this0.ImplementsGAction()
	}
	value1 = (*C.GVariant)(unsafe.Pointer(value0))
	C.g_action_change_state(this1, value1)
}
func (this0 *ActionImpl) GetEnabled() bool {
	var this1 *C.GAction
	if this0 != nil {
		this1 = this0.ImplementsGAction()
	}
	ret1 := C.g_action_get_enabled(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *ActionImpl) GetName() string {
	var this1 *C.GAction
	if this0 != nil {
		this1 = this0.ImplementsGAction()
	}
	ret1 := C.g_action_get_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *ActionImpl) GetParameterType() *glib.VariantType {
	var this1 *C.GAction
	if this0 != nil {
		this1 = this0.ImplementsGAction()
	}
	ret1 := C.g_action_get_parameter_type(this1)
	var ret2 *glib.VariantType
	ret2 = (*glib.VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *ActionImpl) GetState() *glib.Variant {
	var this1 *C.GAction
	if this0 != nil {
		this1 = this0.ImplementsGAction()
	}
	ret1 := C.g_action_get_state(this1)
	var ret2 *glib.Variant
	ret2 = (*glib.Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *ActionImpl) GetStateHint() *glib.Variant {
	var this1 *C.GAction
	if this0 != nil {
		this1 = this0.ImplementsGAction()
	}
	ret1 := C.g_action_get_state_hint(this1)
	var ret2 *glib.Variant
	ret2 = (*glib.Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *ActionImpl) GetStateType() *glib.VariantType {
	var this1 *C.GAction
	if this0 != nil {
		this1 = this0.ImplementsGAction()
	}
	ret1 := C.g_action_get_state_type(this1)
	var ret2 *glib.VariantType
	ret2 = (*glib.VariantType)(unsafe.Pointer(ret1))
	return ret2
}
type ActionEntry struct {
	name0 *C.char
	Activate unsafe.Pointer
	parameter_type0 *C.char
	state0 *C.char
	ChangeState unsafe.Pointer
	Padding [3]uint64
}
func (this0 *ActionEntry) Name() string {
	var name1 string
	name1 = C.GoString(this0.name0)
	return name1
}
func (this0 *ActionEntry) ParameterType() string {
	var parameter_type1 string
	parameter_type1 = C.GoString(this0.parameter_type0)
	return parameter_type1
}
func (this0 *ActionEntry) State() string {
	var state1 string
	state1 = C.GoString(this0.state0)
	return state1
}
type ActionGroupLike interface {
	ImplementsGActionGroup() *C.GActionGroup
}

type ActionGroup struct {
	gobject.Object
	ActionGroupImpl
}

type ActionGroupImpl struct {}

func ToActionGroup(objlike gobject.ObjectLike) *ActionGroup {
	t := (*ActionGroupImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ActionGroup)(obj)
	}
	panic("cannot cast to ActionGroup")
}

func (this0 *ActionGroupImpl) ImplementsGActionGroup() *C.GActionGroup {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GActionGroup)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *ActionGroupImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.g_action_group_get_type())
}

func ActionGroupGetType() gobject.Type {
	return (*ActionGroupImpl)(nil).GetStaticType()
}
func (this0 *ActionGroupImpl) ActionAdded(action_name0 string) {
	var this1 *C.GActionGroup
	var action_name1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()
	}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	C.g_action_group_action_added(this1, action_name1)
}
func (this0 *ActionGroupImpl) ActionEnabledChanged(action_name0 string, enabled0 bool) {
	var this1 *C.GActionGroup
	var action_name1 *C.char
	var enabled1 C.int
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()
	}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	enabled1 = _GoBoolToCBool(enabled0)
	C.g_action_group_action_enabled_changed(this1, action_name1, enabled1)
}
func (this0 *ActionGroupImpl) ActionRemoved(action_name0 string) {
	var this1 *C.GActionGroup
	var action_name1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()
	}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	C.g_action_group_action_removed(this1, action_name1)
}
func (this0 *ActionGroupImpl) ActionStateChanged(action_name0 string, state0 *glib.Variant) {
	var this1 *C.GActionGroup
	var action_name1 *C.char
	var state1 *C.GVariant
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()
	}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	state1 = (*C.GVariant)(unsafe.Pointer(state0))
	C.g_action_group_action_state_changed(this1, action_name1, state1)
}
func (this0 *ActionGroupImpl) ActivateAction(action_name0 string, parameter0 *glib.Variant) {
	var this1 *C.GActionGroup
	var action_name1 *C.char
	var parameter1 *C.GVariant
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()
	}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	parameter1 = (*C.GVariant)(unsafe.Pointer(parameter0))
	C.g_action_group_activate_action(this1, action_name1, parameter1)
}
func (this0 *ActionGroupImpl) ChangeActionState(action_name0 string, value0 *glib.Variant) {
	var this1 *C.GActionGroup
	var action_name1 *C.char
	var value1 *C.GVariant
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()
	}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	value1 = (*C.GVariant)(unsafe.Pointer(value0))
	C.g_action_group_change_action_state(this1, action_name1, value1)
}
func (this0 *ActionGroupImpl) GetActionEnabled(action_name0 string) bool {
	var this1 *C.GActionGroup
	var action_name1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()
	}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	ret1 := C.g_action_group_get_action_enabled(this1, action_name1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *ActionGroupImpl) GetActionParameterType(action_name0 string) *glib.VariantType {
	var this1 *C.GActionGroup
	var action_name1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()
	}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	ret1 := C.g_action_group_get_action_parameter_type(this1, action_name1)
	var ret2 *glib.VariantType
	ret2 = (*glib.VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *ActionGroupImpl) GetActionState(action_name0 string) *glib.Variant {
	var this1 *C.GActionGroup
	var action_name1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()
	}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	ret1 := C.g_action_group_get_action_state(this1, action_name1)
	var ret2 *glib.Variant
	ret2 = (*glib.Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *ActionGroupImpl) GetActionStateHint(action_name0 string) *glib.Variant {
	var this1 *C.GActionGroup
	var action_name1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()
	}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	ret1 := C.g_action_group_get_action_state_hint(this1, action_name1)
	var ret2 *glib.Variant
	ret2 = (*glib.Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *ActionGroupImpl) GetActionStateType(action_name0 string) *glib.VariantType {
	var this1 *C.GActionGroup
	var action_name1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()
	}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	ret1 := C.g_action_group_get_action_state_type(this1, action_name1)
	var ret2 *glib.VariantType
	ret2 = (*glib.VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *ActionGroupImpl) HasAction(action_name0 string) bool {
	var this1 *C.GActionGroup
	var action_name1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()
	}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	ret1 := C.g_action_group_has_action(this1, action_name1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *ActionGroupImpl) ListActions() []string {
	var this1 *C.GActionGroup
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()
	}
	ret1 := C.g_action_group_list_actions(this1)
	var ret2 []string
	ret2 = make([]string, uint(C._array_length(unsafe.Pointer(ret1))))
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return ret2
}
func (this0 *ActionGroupImpl) QueryAction(action_name0 string) (bool, *glib.VariantType, *glib.VariantType, *glib.Variant, *glib.Variant, bool) {
	var this1 *C.GActionGroup
	var action_name1 *C.char
	var enabled1 C.int
	var parameter_type1 *C.GVariantType
	var state_type1 *C.GVariantType
	var state_hint1 *C.GVariant
	var state1 *C.GVariant
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()
	}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	ret1 := C.g_action_group_query_action(this1, action_name1, &enabled1, &parameter_type1, &state_type1, &state_hint1, &state1)
	var enabled2 bool
	var parameter_type2 *glib.VariantType
	var state_type2 *glib.VariantType
	var state_hint2 *glib.Variant
	var state2 *glib.Variant
	var ret2 bool
	enabled2 = enabled1 != 0
	parameter_type2 = (*glib.VariantType)(unsafe.Pointer(parameter_type1))
	state_type2 = (*glib.VariantType)(unsafe.Pointer(state_type1))
	state_hint2 = (*glib.Variant)(unsafe.Pointer(state_hint1))
	state2 = (*glib.Variant)(unsafe.Pointer(state1))
	ret2 = ret1 != 0
	return enabled2, parameter_type2, state_type2, state_hint2, state2, ret2
}
// blacklisted: ActionGroupInterface (struct)
// blacklisted: ActionInterface (struct)
type ActionMapLike interface {
	ImplementsGActionMap() *C.GActionMap
}

type ActionMap struct {
	gobject.Object
	ActionMapImpl
}

type ActionMapImpl struct {}

func ToActionMap(objlike gobject.ObjectLike) *ActionMap {
	t := (*ActionMapImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ActionMap)(obj)
	}
	panic("cannot cast to ActionMap")
}

func (this0 *ActionMapImpl) ImplementsGActionMap() *C.GActionMap {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GActionMap)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *ActionMapImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.g_action_map_get_type())
}

func ActionMapGetType() gobject.Type {
	return (*ActionMapImpl)(nil).GetStaticType()
}
func (this0 *ActionMapImpl) AddAction(action0 ActionLike) {
	var this1 *C.GActionMap
	var action1 *C.GAction
	if this0 != nil {
		this1 = this0.ImplementsGActionMap()
	}
	if action0 != nil {
		action1 = action0.ImplementsGAction()
	}
	C.g_action_map_add_action(this1, action1)
}
func (this0 *ActionMapImpl) AddActionEntries(entries0 []ActionEntry, user_data0 unsafe.Pointer) {
	var this1 *C.GActionMap
	var entries1 *C.GActionEntry
	var n_entries1 C.int32_t
	var user_data1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGActionMap()
	}
	entries1 = (*C.GActionEntry)(C.malloc(C.size_t(int(unsafe.Sizeof(*entries1)) * len(entries0))))
	defer C.free(unsafe.Pointer(entries1))
	for i, e := range entries0 {
		(*(*[999999]C.GActionEntry)(unsafe.Pointer(entries1)))[i] = *(*C.GActionEntry)(unsafe.Pointer(&e))
	}
	n_entries1 = C.int32_t(len(entries0))
	user_data1 = unsafe.Pointer(user_data0)
	C.g_action_map_add_action_entries(this1, entries1, n_entries1, user_data1)
}
func (this0 *ActionMapImpl) LookupAction(action_name0 string) *Action {
	var this1 *C.GActionMap
	var action_name1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsGActionMap()
	}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	ret1 := C.g_action_map_lookup_action(this1, action_name1)
	var ret2 *Action
	ret2 = (*Action)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *ActionMapImpl) RemoveAction(action_name0 string) {
	var this1 *C.GActionMap
	var action_name1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsGActionMap()
	}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	C.g_action_map_remove_action(this1, action_name1)
}
// blacklisted: ActionMapInterface (struct)
type AppInfoLike interface {
	ImplementsGAppInfo() *C.GAppInfo
}

type AppInfo struct {
	gobject.Object
	AppInfoImpl
}

type AppInfoImpl struct {}

func ToAppInfo(objlike gobject.ObjectLike) *AppInfo {
	t := (*AppInfoImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*AppInfo)(obj)
	}
	panic("cannot cast to AppInfo")
}

func (this0 *AppInfoImpl) ImplementsGAppInfo() *C.GAppInfo {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GAppInfo)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *AppInfoImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.g_app_info_get_type())
}

func AppInfoGetType() gobject.Type {
	return (*AppInfoImpl)(nil).GetStaticType()
}
func AppInfoCreateFromCommandline(commandline0 string, application_name0 string, flags0 AppInfoCreateFlags) (*AppInfo, error) {
	var commandline1 *C.char
	var application_name1 *C.char
	var flags1 C.GAppInfoCreateFlags
	var err1 *C.GError
	commandline1 = _GoStringToGString(commandline0)
	defer C.free(unsafe.Pointer(commandline1))
	application_name1 = _GoStringToGString(application_name0)
	defer C.free(unsafe.Pointer(application_name1))
	flags1 = C.GAppInfoCreateFlags(flags0)
	ret1 := C.g_app_info_create_from_commandline(commandline1, application_name1, flags1, &err1)
	var ret2 *AppInfo
	var err2 error
	ret2 = (*AppInfo)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func AppInfoGetAll() []*AppInfo {
	ret1 := C.g_app_info_get_all()
	var ret2 []*AppInfo
	for iter := (*_GList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt *AppInfo
		elt = (*AppInfo)(gobject.ObjectWrap(unsafe.Pointer((*C.GAppInfo)(iter.data)), false))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func AppInfoGetAllForType(content_type0 string) []*AppInfo {
	var content_type1 *C.char
	content_type1 = _GoStringToGString(content_type0)
	defer C.free(unsafe.Pointer(content_type1))
	ret1 := C.g_app_info_get_all_for_type(content_type1)
	var ret2 []*AppInfo
	for iter := (*_GList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt *AppInfo
		elt = (*AppInfo)(gobject.ObjectWrap(unsafe.Pointer((*C.GAppInfo)(iter.data)), false))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func AppInfoGetDefaultForType(content_type0 string, must_support_uris0 bool) *AppInfo {
	var content_type1 *C.char
	var must_support_uris1 C.int
	content_type1 = _GoStringToGString(content_type0)
	defer C.free(unsafe.Pointer(content_type1))
	must_support_uris1 = _GoBoolToCBool(must_support_uris0)
	ret1 := C.g_app_info_get_default_for_type(content_type1, must_support_uris1)
	var ret2 *AppInfo
	ret2 = (*AppInfo)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func AppInfoGetDefaultForUriScheme(uri_scheme0 string) *AppInfo {
	var uri_scheme1 *C.char
	uri_scheme1 = _GoStringToGString(uri_scheme0)
	defer C.free(unsafe.Pointer(uri_scheme1))
	ret1 := C.g_app_info_get_default_for_uri_scheme(uri_scheme1)
	var ret2 *AppInfo
	ret2 = (*AppInfo)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func AppInfoGetFallbackForType(content_type0 string) []*AppInfo {
	var content_type1 *C.char
	content_type1 = _GoStringToGString(content_type0)
	defer C.free(unsafe.Pointer(content_type1))
	ret1 := C.g_app_info_get_fallback_for_type(content_type1)
	var ret2 []*AppInfo
	for iter := (*_GList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt *AppInfo
		elt = (*AppInfo)(gobject.ObjectWrap(unsafe.Pointer((*C.GAppInfo)(iter.data)), false))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func AppInfoGetRecommendedForType(content_type0 string) []*AppInfo {
	var content_type1 *C.char
	content_type1 = _GoStringToGString(content_type0)
	defer C.free(unsafe.Pointer(content_type1))
	ret1 := C.g_app_info_get_recommended_for_type(content_type1)
	var ret2 []*AppInfo
	for iter := (*_GList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt *AppInfo
		elt = (*AppInfo)(gobject.ObjectWrap(unsafe.Pointer((*C.GAppInfo)(iter.data)), false))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func AppInfoLaunchDefaultForUri(uri0 string, launch_context0 AppLaunchContextLike) (bool, error) {
	var uri1 *C.char
	var launch_context1 *C.GAppLaunchContext
	var err1 *C.GError
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	if launch_context0 != nil {
		launch_context1 = launch_context0.InheritedFromGAppLaunchContext()
	}
	ret1 := C.g_app_info_launch_default_for_uri(uri1, launch_context1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func AppInfoResetTypeAssociations(content_type0 string) {
	var content_type1 *C.char
	content_type1 = _GoStringToGString(content_type0)
	defer C.free(unsafe.Pointer(content_type1))
	C.g_app_info_reset_type_associations(content_type1)
}
func (this0 *AppInfoImpl) AddSupportsType(content_type0 string) (bool, error) {
	var this1 *C.GAppInfo
	var content_type1 *C.char
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGAppInfo()
	}
	content_type1 = _GoStringToGString(content_type0)
	defer C.free(unsafe.Pointer(content_type1))
	ret1 := C.g_app_info_add_supports_type(this1, content_type1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *AppInfoImpl) CanDelete() bool {
	var this1 *C.GAppInfo
	if this0 != nil {
		this1 = this0.ImplementsGAppInfo()
	}
	ret1 := C.g_app_info_can_delete(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *AppInfoImpl) CanRemoveSupportsType() bool {
	var this1 *C.GAppInfo
	if this0 != nil {
		this1 = this0.ImplementsGAppInfo()
	}
	ret1 := C.g_app_info_can_remove_supports_type(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *AppInfoImpl) Delete() bool {
	var this1 *C.GAppInfo
	if this0 != nil {
		this1 = this0.ImplementsGAppInfo()
	}
	ret1 := C.g_app_info_delete(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *AppInfoImpl) Dup() *AppInfo {
	var this1 *C.GAppInfo
	if this0 != nil {
		this1 = this0.ImplementsGAppInfo()
	}
	ret1 := C.g_app_info_dup(this1)
	var ret2 *AppInfo
	ret2 = (*AppInfo)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *AppInfoImpl) Equal(appinfo20 AppInfoLike) bool {
	var this1 *C.GAppInfo
	var appinfo21 *C.GAppInfo
	if this0 != nil {
		this1 = this0.ImplementsGAppInfo()
	}
	if appinfo20 != nil {
		appinfo21 = appinfo20.ImplementsGAppInfo()
	}
	ret1 := C.g_app_info_equal(this1, appinfo21)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *AppInfoImpl) GetCommandline() string {
	var this1 *C.GAppInfo
	if this0 != nil {
		this1 = this0.ImplementsGAppInfo()
	}
	ret1 := C.g_app_info_get_commandline(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *AppInfoImpl) GetDescription() string {
	var this1 *C.GAppInfo
	if this0 != nil {
		this1 = this0.ImplementsGAppInfo()
	}
	ret1 := C.g_app_info_get_description(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *AppInfoImpl) GetDisplayName() string {
	var this1 *C.GAppInfo
	if this0 != nil {
		this1 = this0.ImplementsGAppInfo()
	}
	ret1 := C.g_app_info_get_display_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *AppInfoImpl) GetExecutable() string {
	var this1 *C.GAppInfo
	if this0 != nil {
		this1 = this0.ImplementsGAppInfo()
	}
	ret1 := C.g_app_info_get_executable(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *AppInfoImpl) GetIcon() *Icon {
	var this1 *C.GAppInfo
	if this0 != nil {
		this1 = this0.ImplementsGAppInfo()
	}
	ret1 := C.g_app_info_get_icon(this1)
	var ret2 *Icon
	ret2 = (*Icon)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *AppInfoImpl) GetId() string {
	var this1 *C.GAppInfo
	if this0 != nil {
		this1 = this0.ImplementsGAppInfo()
	}
	ret1 := C.g_app_info_get_id(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *AppInfoImpl) GetName() string {
	var this1 *C.GAppInfo
	if this0 != nil {
		this1 = this0.ImplementsGAppInfo()
	}
	ret1 := C.g_app_info_get_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *AppInfoImpl) GetSupportedTypes() []string {
	var this1 *C.GAppInfo
	if this0 != nil {
		this1 = this0.ImplementsGAppInfo()
	}
	ret1 := C.g_app_info_get_supported_types(this1)
	var ret2 []string
	ret2 = make([]string, uint(C._array_length(unsafe.Pointer(ret1))))
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
	}
	return ret2
}
func (this0 *AppInfoImpl) Launch(files0 []*File, launch_context0 AppLaunchContextLike) (bool, error) {
	var this1 *C.GAppInfo
	var files1 *C.GList
	var launch_context1 *C.GAppLaunchContext
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGAppInfo()
	}
	if launch_context0 != nil {
		launch_context1 = launch_context0.InheritedFromGAppLaunchContext()
	}
	ret1 := C.g_app_info_launch(this1, files1, launch_context1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *AppInfoImpl) LaunchUris(uris0 []string, launch_context0 AppLaunchContextLike) (bool, error) {
	var this1 *C.GAppInfo
	var uris1 *C.GList
	var launch_context1 *C.GAppLaunchContext
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGAppInfo()
	}
	if launch_context0 != nil {
		launch_context1 = launch_context0.InheritedFromGAppLaunchContext()
	}
	ret1 := C.g_app_info_launch_uris(this1, uris1, launch_context1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *AppInfoImpl) RemoveSupportsType(content_type0 string) (bool, error) {
	var this1 *C.GAppInfo
	var content_type1 *C.char
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGAppInfo()
	}
	content_type1 = _GoStringToGString(content_type0)
	defer C.free(unsafe.Pointer(content_type1))
	ret1 := C.g_app_info_remove_supports_type(this1, content_type1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *AppInfoImpl) SetAsDefaultForExtension(extension0 string) (bool, error) {
	var this1 *C.GAppInfo
	var extension1 *C.char
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGAppInfo()
	}
	extension1 = _GoStringToGString(extension0)
	defer C.free(unsafe.Pointer(extension1))
	ret1 := C.g_app_info_set_as_default_for_extension(this1, extension1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *AppInfoImpl) SetAsDefaultForType(content_type0 string) (bool, error) {
	var this1 *C.GAppInfo
	var content_type1 *C.char
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGAppInfo()
	}
	content_type1 = _GoStringToGString(content_type0)
	defer C.free(unsafe.Pointer(content_type1))
	ret1 := C.g_app_info_set_as_default_for_type(this1, content_type1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *AppInfoImpl) SetAsLastUsedForType(content_type0 string) (bool, error) {
	var this1 *C.GAppInfo
	var content_type1 *C.char
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGAppInfo()
	}
	content_type1 = _GoStringToGString(content_type0)
	defer C.free(unsafe.Pointer(content_type1))
	ret1 := C.g_app_info_set_as_last_used_for_type(this1, content_type1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *AppInfoImpl) ShouldShow() bool {
	var this1 *C.GAppInfo
	if this0 != nil {
		this1 = this0.ImplementsGAppInfo()
	}
	ret1 := C.g_app_info_should_show(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *AppInfoImpl) SupportsFiles() bool {
	var this1 *C.GAppInfo
	if this0 != nil {
		this1 = this0.ImplementsGAppInfo()
	}
	ret1 := C.g_app_info_supports_files(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *AppInfoImpl) SupportsUris() bool {
	var this1 *C.GAppInfo
	if this0 != nil {
		this1 = this0.ImplementsGAppInfo()
	}
	ret1 := C.g_app_info_supports_uris(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
type AppInfoCreateFlags C.uint32_t
const (
	AppInfoCreateFlagsNone AppInfoCreateFlags = 0
	AppInfoCreateFlagsNeedsTerminal AppInfoCreateFlags = 1
	AppInfoCreateFlagsSupportsUris AppInfoCreateFlags = 2
	AppInfoCreateFlagsSupportsStartupNotification AppInfoCreateFlags = 4
)
// blacklisted: AppInfoIface (struct)
type AppLaunchContextLike interface {
	gobject.ObjectLike
	InheritedFromGAppLaunchContext() *C.GAppLaunchContext
}

type AppLaunchContext struct {
	gobject.Object
	
}

func ToAppLaunchContext(objlike gobject.ObjectLike) *AppLaunchContext {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*AppLaunchContext)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*AppLaunchContext)(obj)
	}
	panic("cannot cast to AppLaunchContext")
}

func (this0 *AppLaunchContext) InheritedFromGAppLaunchContext() *C.GAppLaunchContext {
	if this0 == nil {
		return nil
	}
	return (*C.GAppLaunchContext)(this0.C)
}

func (this0 *AppLaunchContext) GetStaticType() gobject.Type {
	return gobject.Type(C.g_app_launch_context_get_type())
}

func AppLaunchContextGetType() gobject.Type {
	return (*AppLaunchContext)(nil).GetStaticType()
}
func NewAppLaunchContext() *AppLaunchContext {
	ret1 := C.g_app_launch_context_new()
	var ret2 *AppLaunchContext
	ret2 = (*AppLaunchContext)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *AppLaunchContext) GetDisplay(info0 AppInfoLike, files0 []*File) string {
	var this1 *C.GAppLaunchContext
	var info1 *C.GAppInfo
	var files1 *C.GList
	if this0 != nil {
		this1 = this0.InheritedFromGAppLaunchContext()
	}
	if info0 != nil {
		info1 = info0.ImplementsGAppInfo()
	}
	ret1 := C.g_app_launch_context_get_display(this1, info1, files1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *AppLaunchContext) GetEnvironment() []string {
	var this1 *C.GAppLaunchContext
	if this0 != nil {
		this1 = this0.InheritedFromGAppLaunchContext()
	}
	ret1 := C.g_app_launch_context_get_environment(this1)
	var ret2 []string
	ret2 = make([]string, uint(C._array_length(unsafe.Pointer(ret1))))
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return ret2
}
func (this0 *AppLaunchContext) GetStartupNotifyId(info0 AppInfoLike, files0 []*File) string {
	var this1 *C.GAppLaunchContext
	var info1 *C.GAppInfo
	var files1 *C.GList
	if this0 != nil {
		this1 = this0.InheritedFromGAppLaunchContext()
	}
	if info0 != nil {
		info1 = info0.ImplementsGAppInfo()
	}
	ret1 := C.g_app_launch_context_get_startup_notify_id(this1, info1, files1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *AppLaunchContext) LaunchFailed(startup_notify_id0 string) {
	var this1 *C.GAppLaunchContext
	var startup_notify_id1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGAppLaunchContext()
	}
	startup_notify_id1 = _GoStringToGString(startup_notify_id0)
	defer C.free(unsafe.Pointer(startup_notify_id1))
	C.g_app_launch_context_launch_failed(this1, startup_notify_id1)
}
func (this0 *AppLaunchContext) Setenv(variable0 string, value0 string) {
	var this1 *C.GAppLaunchContext
	var variable1 *C.char
	var value1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGAppLaunchContext()
	}
	variable1 = _GoStringToGString(variable0)
	defer C.free(unsafe.Pointer(variable1))
	value1 = _GoStringToGString(value0)
	defer C.free(unsafe.Pointer(value1))
	C.g_app_launch_context_setenv(this1, variable1, value1)
}
func (this0 *AppLaunchContext) Unsetenv(variable0 string) {
	var this1 *C.GAppLaunchContext
	var variable1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGAppLaunchContext()
	}
	variable1 = _GoStringToGString(variable0)
	defer C.free(unsafe.Pointer(variable1))
	C.g_app_launch_context_unsetenv(this1, variable1)
}
// blacklisted: AppLaunchContextClass (struct)
// blacklisted: AppLaunchContextPrivate (struct)
type ApplicationLike interface {
	gobject.ObjectLike
	InheritedFromGApplication() *C.GApplication
}

type Application struct {
	gobject.Object
	ActionGroupImpl
	ActionMapImpl
}

func ToApplication(objlike gobject.ObjectLike) *Application {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Application)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Application)(obj)
	}
	panic("cannot cast to Application")
}

func (this0 *Application) InheritedFromGApplication() *C.GApplication {
	if this0 == nil {
		return nil
	}
	return (*C.GApplication)(this0.C)
}

func (this0 *Application) GetStaticType() gobject.Type {
	return gobject.Type(C.g_application_get_type())
}

func ApplicationGetType() gobject.Type {
	return (*Application)(nil).GetStaticType()
}
func NewApplication(application_id0 string, flags0 ApplicationFlags) *Application {
	var application_id1 *C.char
	var flags1 C.GApplicationFlags
	application_id1 = _GoStringToGString(application_id0)
	defer C.free(unsafe.Pointer(application_id1))
	flags1 = C.GApplicationFlags(flags0)
	ret1 := C.g_application_new(application_id1, flags1)
	var ret2 *Application
	ret2 = (*Application)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func ApplicationGetDefault() *Application {
	ret1 := C.g_application_get_default()
	var ret2 *Application
	ret2 = (*Application)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func ApplicationIdIsValid(application_id0 string) bool {
	var application_id1 *C.char
	application_id1 = _GoStringToGString(application_id0)
	defer C.free(unsafe.Pointer(application_id1))
	ret1 := C.g_application_id_is_valid(application_id1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Application) Activate() {
	var this1 *C.GApplication
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	C.g_application_activate(this1)
}
func (this0 *Application) GetApplicationId() string {
	var this1 *C.GApplication
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	ret1 := C.g_application_get_application_id(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
// blacklisted: Application.get_dbus_connection (method)
func (this0 *Application) GetDbusObjectPath() string {
	var this1 *C.GApplication
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	ret1 := C.g_application_get_dbus_object_path(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *Application) GetFlags() ApplicationFlags {
	var this1 *C.GApplication
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	ret1 := C.g_application_get_flags(this1)
	var ret2 ApplicationFlags
	ret2 = ApplicationFlags(ret1)
	return ret2
}
func (this0 *Application) GetInactivityTimeout() int {
	var this1 *C.GApplication
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	ret1 := C.g_application_get_inactivity_timeout(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Application) GetIsRegistered() bool {
	var this1 *C.GApplication
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	ret1 := C.g_application_get_is_registered(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Application) GetIsRemote() bool {
	var this1 *C.GApplication
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	ret1 := C.g_application_get_is_remote(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Application) Hold() {
	var this1 *C.GApplication
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	C.g_application_hold(this1)
}
func (this0 *Application) Open(files0 []FileLike, hint0 string) {
	var this1 *C.GApplication
	var files1 **C.GFile
	var n_files1 C.int32_t
	var hint1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	files1 = (**C.GFile)(C.malloc(C.size_t(int(unsafe.Sizeof(*files1)) * len(files0))))
	defer C.free(unsafe.Pointer(files1))
	for i, e := range files0 {
		if e != nil {
			(*(*[999999]*C.GFile)(unsafe.Pointer(files1)))[i] = e.ImplementsGFile()
		}
	}
	n_files1 = C.int32_t(len(files0))
	hint1 = _GoStringToGString(hint0)
	defer C.free(unsafe.Pointer(hint1))
	C.g_application_open(this1, files1, n_files1, hint1)
}
func (this0 *Application) Quit() {
	var this1 *C.GApplication
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	C.g_application_quit(this1)
}
func (this0 *Application) Register(cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GApplication
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	if cancellable0 != nil {
		cancellable1 = cancellable0.InheritedFromGCancellable()
	}
	ret1 := C.g_application_register(this1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *Application) Release() {
	var this1 *C.GApplication
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	C.g_application_release(this1)
}
func (this0 *Application) Run(argv0 []string) int {
	var this1 *C.GApplication
	var argv1 **C.char
	var argc1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	argv1 = (**C.char)(C.malloc(C.size_t(int(unsafe.Sizeof(*argv1)) * len(argv0))))
	defer C.free(unsafe.Pointer(argv1))
	for i, e := range argv0 {
		(*(*[999999]*C.char)(unsafe.Pointer(argv1)))[i] = _GoStringToGString(e)
		defer C.free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(argv1)))[i]))
	}
	argc1 = C.int32_t(len(argv0))
	ret1 := C.g_application_run(this1, argc1, argv1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Application) SetActionGroup(action_group0 ActionGroupLike) {
	var this1 *C.GApplication
	var action_group1 *C.GActionGroup
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	if action_group0 != nil {
		action_group1 = action_group0.ImplementsGActionGroup()
	}
	C.g_application_set_action_group(this1, action_group1)
}
func (this0 *Application) SetApplicationId(application_id0 string) {
	var this1 *C.GApplication
	var application_id1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	application_id1 = _GoStringToGString(application_id0)
	defer C.free(unsafe.Pointer(application_id1))
	C.g_application_set_application_id(this1, application_id1)
}
func (this0 *Application) SetDefault() {
	var this1 *C.GApplication
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	C.g_application_set_default(this1)
}
func (this0 *Application) SetFlags(flags0 ApplicationFlags) {
	var this1 *C.GApplication
	var flags1 C.GApplicationFlags
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	flags1 = C.GApplicationFlags(flags0)
	C.g_application_set_flags(this1, flags1)
}
func (this0 *Application) SetInactivityTimeout(inactivity_timeout0 int) {
	var this1 *C.GApplication
	var inactivity_timeout1 C.uint32_t
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	inactivity_timeout1 = C.uint32_t(inactivity_timeout0)
	C.g_application_set_inactivity_timeout(this1, inactivity_timeout1)
}
// blacklisted: ApplicationClass (struct)
// blacklisted: ApplicationCommandLine (object)
// blacklisted: ApplicationCommandLineClass (struct)
// blacklisted: ApplicationCommandLinePrivate (struct)
type ApplicationFlags C.uint32_t
const (
	ApplicationFlagsFlagsNone ApplicationFlags = 0
	ApplicationFlagsIsService ApplicationFlags = 1
	ApplicationFlagsIsLauncher ApplicationFlags = 2
	ApplicationFlagsHandlesOpen ApplicationFlags = 4
	ApplicationFlagsHandlesCommandLine ApplicationFlags = 8
	ApplicationFlagsSendEnvironment ApplicationFlags = 16
	ApplicationFlagsNonUnique ApplicationFlags = 32
)
// blacklisted: ApplicationPrivate (struct)
type AskPasswordFlags C.uint32_t
const (
	AskPasswordFlagsNeedPassword AskPasswordFlags = 1
	AskPasswordFlagsNeedUsername AskPasswordFlags = 2
	AskPasswordFlagsNeedDomain AskPasswordFlags = 4
	AskPasswordFlagsSavingSupported AskPasswordFlags = 8
	AskPasswordFlagsAnonymousSupported AskPasswordFlags = 16
)
// blacklisted: AsyncInitable (interface)
// blacklisted: AsyncInitableIface (struct)
type AsyncReadyCallback func(source_object *gobject.Object, res *AsyncResult)
//export _GAsyncReadyCallback_c_wrapper
func _GAsyncReadyCallback_c_wrapper(source_object0 unsafe.Pointer, res0 unsafe.Pointer, user_data0 unsafe.Pointer) {
	var source_object1 *gobject.Object
	var res1 *AsyncResult
	var user_data1 AsyncReadyCallback
	source_object1 = (*gobject.Object)(gobject.ObjectWrap(unsafe.Pointer((*C.GObject)(source_object0)), true))
	res1 = (*AsyncResult)(gobject.ObjectWrap(unsafe.Pointer((*C.GAsyncResult)(res0)), true))
	user_data1 = *(*AsyncReadyCallback)(user_data0)
	user_data1(source_object1, res1)
}
//export _GAsyncReadyCallback_c_wrapper_once
func _GAsyncReadyCallback_c_wrapper_once(source_object0 unsafe.Pointer, res0 unsafe.Pointer, user_data0 unsafe.Pointer) {
	_GAsyncReadyCallback_c_wrapper(source_object0, res0, user_data0)
	gobject.Holder.Release(user_data0)
}
type AsyncResultLike interface {
	ImplementsGAsyncResult() *C.GAsyncResult
}

type AsyncResult struct {
	gobject.Object
	AsyncResultImpl
}

type AsyncResultImpl struct {}

func ToAsyncResult(objlike gobject.ObjectLike) *AsyncResult {
	t := (*AsyncResultImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*AsyncResult)(obj)
	}
	panic("cannot cast to AsyncResult")
}

func (this0 *AsyncResultImpl) ImplementsGAsyncResult() *C.GAsyncResult {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GAsyncResult)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *AsyncResultImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.g_async_result_get_type())
}

func AsyncResultGetType() gobject.Type {
	return (*AsyncResultImpl)(nil).GetStaticType()
}
func (this0 *AsyncResultImpl) GetSourceObject() *gobject.Object {
	var this1 *C.GAsyncResult
	if this0 != nil {
		this1 = this0.ImplementsGAsyncResult()
	}
	ret1 := C.g_async_result_get_source_object(this1)
	var ret2 *gobject.Object
	ret2 = (*gobject.Object)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *AsyncResultImpl) GetUserData() {
	var this1 *C.GAsyncResult
	if this0 != nil {
		this1 = this0.ImplementsGAsyncResult()
	}
	C.g_async_result_get_user_data(this1)
}
func (this0 *AsyncResultImpl) IsTagged(source_tag0 unsafe.Pointer) bool {
	var this1 *C.GAsyncResult
	var source_tag1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGAsyncResult()
	}
	source_tag1 = unsafe.Pointer(source_tag0)
	ret1 := C.g_async_result_is_tagged(this1, source_tag1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *AsyncResultImpl) LegacyPropagateError() (bool, error) {
	var this1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGAsyncResult()
	}
	ret1 := C.g_async_result_legacy_propagate_error(this1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
// blacklisted: AsyncResultIface (struct)
// blacklisted: BufferedInputStream (object)
// blacklisted: BufferedInputStreamClass (struct)
// blacklisted: BufferedInputStreamPrivate (struct)
// blacklisted: BufferedOutputStream (object)
// blacklisted: BufferedOutputStreamClass (struct)
// blacklisted: BufferedOutputStreamPrivate (struct)
// blacklisted: BusAcquiredCallback (callback)
// blacklisted: BusNameAcquiredCallback (callback)
// blacklisted: BusNameAppearedCallback (callback)
// blacklisted: BusNameLostCallback (callback)
type BusNameOwnerFlags C.uint32_t
const (
	BusNameOwnerFlagsNone BusNameOwnerFlags = 0
	BusNameOwnerFlagsAllowReplacement BusNameOwnerFlags = 1
	BusNameOwnerFlagsReplace BusNameOwnerFlags = 2
)
// blacklisted: BusNameVanishedCallback (callback)
type BusNameWatcherFlags C.uint32_t
const (
	BusNameWatcherFlagsNone BusNameWatcherFlags = 0
	BusNameWatcherFlagsAutoStart BusNameWatcherFlags = 1
)
type BusType C.int32_t
const (
	BusTypeStarter BusType = -1
	BusTypeNone BusType = 0
	BusTypeSystem BusType = 1
	BusTypeSession BusType = 2
)
type CancellableLike interface {
	gobject.ObjectLike
	InheritedFromGCancellable() *C.GCancellable
}

type Cancellable struct {
	gobject.Object
	
}

func ToCancellable(objlike gobject.ObjectLike) *Cancellable {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Cancellable)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Cancellable)(obj)
	}
	panic("cannot cast to Cancellable")
}

func (this0 *Cancellable) InheritedFromGCancellable() *C.GCancellable {
	if this0 == nil {
		return nil
	}
	return (*C.GCancellable)(this0.C)
}

func (this0 *Cancellable) GetStaticType() gobject.Type {
	return gobject.Type(C.g_cancellable_get_type())
}

func CancellableGetType() gobject.Type {
	return (*Cancellable)(nil).GetStaticType()
}
func NewCancellable() *Cancellable {
	ret1 := C.g_cancellable_new()
	var ret2 *Cancellable
	ret2 = (*Cancellable)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func CancellableGetCurrent() *Cancellable {
	ret1 := C.g_cancellable_get_current()
	var ret2 *Cancellable
	ret2 = (*Cancellable)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Cancellable) Cancel() {
	var this1 *C.GCancellable
	if this0 != nil {
		this1 = this0.InheritedFromGCancellable()
	}
	C.g_cancellable_cancel(this1)
}
// blacklisted: Cancellable.connect (method)
func (this0 *Cancellable) Disconnect(handler_id0 uint64) {
	var this1 *C.GCancellable
	var handler_id1 C.uint64_t
	if this0 != nil {
		this1 = this0.InheritedFromGCancellable()
	}
	handler_id1 = C.uint64_t(handler_id0)
	C.g_cancellable_disconnect(this1, handler_id1)
}
func (this0 *Cancellable) GetFd() int {
	var this1 *C.GCancellable
	if this0 != nil {
		this1 = this0.InheritedFromGCancellable()
	}
	ret1 := C.g_cancellable_get_fd(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Cancellable) IsCancelled() bool {
	var this1 *C.GCancellable
	if this0 != nil {
		this1 = this0.InheritedFromGCancellable()
	}
	ret1 := C.g_cancellable_is_cancelled(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Cancellable) MakePollfd(pollfd0 *glib.PollFD) bool {
	var this1 *C.GCancellable
	var pollfd1 *C.GPollFD
	if this0 != nil {
		this1 = this0.InheritedFromGCancellable()
	}
	pollfd1 = (*C.GPollFD)(unsafe.Pointer(pollfd0))
	ret1 := C.g_cancellable_make_pollfd(this1, pollfd1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Cancellable) PopCurrent() {
	var this1 *C.GCancellable
	if this0 != nil {
		this1 = this0.InheritedFromGCancellable()
	}
	C.g_cancellable_pop_current(this1)
}
func (this0 *Cancellable) PushCurrent() {
	var this1 *C.GCancellable
	if this0 != nil {
		this1 = this0.InheritedFromGCancellable()
	}
	C.g_cancellable_push_current(this1)
}
func (this0 *Cancellable) ReleaseFd() {
	var this1 *C.GCancellable
	if this0 != nil {
		this1 = this0.InheritedFromGCancellable()
	}
	C.g_cancellable_release_fd(this1)
}
func (this0 *Cancellable) Reset() {
	var this1 *C.GCancellable
	if this0 != nil {
		this1 = this0.InheritedFromGCancellable()
	}
	C.g_cancellable_reset(this1)
}
func (this0 *Cancellable) SetErrorIfCancelled() (bool, error) {
	var this1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.InheritedFromGCancellable()
	}
	ret1 := C.g_cancellable_set_error_if_cancelled(this1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
// blacklisted: CancellableClass (struct)
// blacklisted: CancellablePrivate (struct)
// blacklisted: CancellableSourceFunc (callback)
// blacklisted: CharsetConverter (object)
// blacklisted: CharsetConverterClass (struct)
// blacklisted: Converter (interface)
type ConverterFlags C.uint32_t
const (
	ConverterFlagsNone ConverterFlags = 0
	ConverterFlagsInputAtEnd ConverterFlags = 1
	ConverterFlagsFlush ConverterFlags = 2
)
// blacklisted: ConverterIface (struct)
// blacklisted: ConverterInputStream (object)
// blacklisted: ConverterInputStreamClass (struct)
// blacklisted: ConverterInputStreamPrivate (struct)
// blacklisted: ConverterOutputStream (object)
// blacklisted: ConverterOutputStreamClass (struct)
// blacklisted: ConverterOutputStreamPrivate (struct)
type ConverterResult C.uint32_t
const (
	ConverterResultError ConverterResult = 0
	ConverterResultConverted ConverterResult = 1
	ConverterResultFinished ConverterResult = 2
	ConverterResultFlushed ConverterResult = 3
)
// blacklisted: Credentials (object)
// blacklisted: CredentialsClass (struct)
type CredentialsType C.uint32_t
const (
	CredentialsTypeInvalid CredentialsType = 0
	CredentialsTypeLinuxUcred CredentialsType = 1
	CredentialsTypeFreebsdCmsgcred CredentialsType = 2
	CredentialsTypeOpenbsdSockpeercred CredentialsType = 3
)
// blacklisted: DBusActionGroup (object)
// blacklisted: DBusAnnotationInfo (struct)
// blacklisted: DBusArgInfo (struct)
// blacklisted: DBusAuthObserver (object)
type DBusCallFlags C.uint32_t
const (
	DBusCallFlagsNone DBusCallFlags = 0
	DBusCallFlagsNoAutoStart DBusCallFlags = 1
)
type DBusCapabilityFlags C.uint32_t
const (
	DBusCapabilityFlagsNone DBusCapabilityFlags = 0
	DBusCapabilityFlagsUnixFdPassing DBusCapabilityFlags = 1
)
// blacklisted: DBusConnection (object)
type DBusConnectionFlags C.uint32_t
const (
	DBusConnectionFlagsNone DBusConnectionFlags = 0
	DBusConnectionFlagsAuthenticationClient DBusConnectionFlags = 1
	DBusConnectionFlagsAuthenticationServer DBusConnectionFlags = 2
	DBusConnectionFlagsAuthenticationAllowAnonymous DBusConnectionFlags = 4
	DBusConnectionFlagsMessageBusConnection DBusConnectionFlags = 8
	DBusConnectionFlagsDelayMessageProcessing DBusConnectionFlags = 16
)
type DBusError C.uint32_t
const (
	DBusErrorFailed DBusError = 0
	DBusErrorNoMemory DBusError = 1
	DBusErrorServiceUnknown DBusError = 2
	DBusErrorNameHasNoOwner DBusError = 3
	DBusErrorNoReply DBusError = 4
	DBusErrorIoError DBusError = 5
	DBusErrorBadAddress DBusError = 6
	DBusErrorNotSupported DBusError = 7
	DBusErrorLimitsExceeded DBusError = 8
	DBusErrorAccessDenied DBusError = 9
	DBusErrorAuthFailed DBusError = 10
	DBusErrorNoServer DBusError = 11
	DBusErrorTimeout DBusError = 12
	DBusErrorNoNetwork DBusError = 13
	DBusErrorAddressInUse DBusError = 14
	DBusErrorDisconnected DBusError = 15
	DBusErrorInvalidArgs DBusError = 16
	DBusErrorFileNotFound DBusError = 17
	DBusErrorFileExists DBusError = 18
	DBusErrorUnknownMethod DBusError = 19
	DBusErrorTimedOut DBusError = 20
	DBusErrorMatchRuleNotFound DBusError = 21
	DBusErrorMatchRuleInvalid DBusError = 22
	DBusErrorSpawnExecFailed DBusError = 23
	DBusErrorSpawnForkFailed DBusError = 24
	DBusErrorSpawnChildExited DBusError = 25
	DBusErrorSpawnChildSignaled DBusError = 26
	DBusErrorSpawnFailed DBusError = 27
	DBusErrorSpawnSetupFailed DBusError = 28
	DBusErrorSpawnConfigInvalid DBusError = 29
	DBusErrorSpawnServiceInvalid DBusError = 30
	DBusErrorSpawnServiceNotFound DBusError = 31
	DBusErrorSpawnPermissionsInvalid DBusError = 32
	DBusErrorSpawnFileInvalid DBusError = 33
	DBusErrorSpawnNoMemory DBusError = 34
	DBusErrorUnixProcessIdUnknown DBusError = 35
	DBusErrorInvalidSignature DBusError = 36
	DBusErrorInvalidFileContent DBusError = 37
	DBusErrorSelinuxSecurityContextUnknown DBusError = 38
	DBusErrorAdtAuditDataUnknown DBusError = 39
	DBusErrorObjectPathInUse DBusError = 40
)
// blacklisted: DBusErrorEntry (struct)
// blacklisted: DBusInterface (interface)
// blacklisted: DBusInterfaceGetPropertyFunc (callback)
// blacklisted: DBusInterfaceIface (struct)
// blacklisted: DBusInterfaceInfo (struct)
// blacklisted: DBusInterfaceMethodCallFunc (callback)
// blacklisted: DBusInterfaceSetPropertyFunc (callback)
// blacklisted: DBusInterfaceSkeleton (object)
// blacklisted: DBusInterfaceSkeletonClass (struct)
type DBusInterfaceSkeletonFlags C.uint32_t
const (
	DBusInterfaceSkeletonFlagsNone DBusInterfaceSkeletonFlags = 0
	DBusInterfaceSkeletonFlagsHandleMethodInvocationsInThread DBusInterfaceSkeletonFlags = 1
)
// blacklisted: DBusInterfaceSkeletonPrivate (struct)
// blacklisted: DBusInterfaceVTable (struct)
// blacklisted: DBusMenuModel (object)
// blacklisted: DBusMessage (object)
type DBusMessageByteOrder C.uint32_t
const (
	DBusMessageByteOrderBigEndian DBusMessageByteOrder = 66
	DBusMessageByteOrderLittleEndian DBusMessageByteOrder = 108
)
// blacklisted: DBusMessageFilterFunction (callback)
type DBusMessageFlags C.uint32_t
const (
	DBusMessageFlagsNone DBusMessageFlags = 0
	DBusMessageFlagsNoReplyExpected DBusMessageFlags = 1
	DBusMessageFlagsNoAutoStart DBusMessageFlags = 2
)
type DBusMessageHeaderField C.uint32_t
const (
	DBusMessageHeaderFieldInvalid DBusMessageHeaderField = 0
	DBusMessageHeaderFieldPath DBusMessageHeaderField = 1
	DBusMessageHeaderFieldInterface DBusMessageHeaderField = 2
	DBusMessageHeaderFieldMember DBusMessageHeaderField = 3
	DBusMessageHeaderFieldErrorName DBusMessageHeaderField = 4
	DBusMessageHeaderFieldReplySerial DBusMessageHeaderField = 5
	DBusMessageHeaderFieldDestination DBusMessageHeaderField = 6
	DBusMessageHeaderFieldSender DBusMessageHeaderField = 7
	DBusMessageHeaderFieldSignature DBusMessageHeaderField = 8
	DBusMessageHeaderFieldNumUnixFds DBusMessageHeaderField = 9
)
type DBusMessageType C.uint32_t
const (
	DBusMessageTypeInvalid DBusMessageType = 0
	DBusMessageTypeMethodCall DBusMessageType = 1
	DBusMessageTypeMethodReturn DBusMessageType = 2
	DBusMessageTypeError DBusMessageType = 3
	DBusMessageTypeSignal DBusMessageType = 4
)
// blacklisted: DBusMethodInfo (struct)
// blacklisted: DBusMethodInvocation (object)
// blacklisted: DBusNodeInfo (struct)
// blacklisted: DBusObject (interface)
// blacklisted: DBusObjectIface (struct)
// blacklisted: DBusObjectManager (interface)
// blacklisted: DBusObjectManagerClient (object)
// blacklisted: DBusObjectManagerClientClass (struct)
type DBusObjectManagerClientFlags C.uint32_t
const (
	DBusObjectManagerClientFlagsNone DBusObjectManagerClientFlags = 0
	DBusObjectManagerClientFlagsDoNotAutoStart DBusObjectManagerClientFlags = 1
)
// blacklisted: DBusObjectManagerClientPrivate (struct)
// blacklisted: DBusObjectManagerIface (struct)
// blacklisted: DBusObjectManagerServer (object)
// blacklisted: DBusObjectManagerServerClass (struct)
// blacklisted: DBusObjectManagerServerPrivate (struct)
// blacklisted: DBusObjectProxy (object)
// blacklisted: DBusObjectProxyClass (struct)
// blacklisted: DBusObjectProxyPrivate (struct)
// blacklisted: DBusObjectSkeleton (object)
// blacklisted: DBusObjectSkeletonClass (struct)
// blacklisted: DBusObjectSkeletonPrivate (struct)
// blacklisted: DBusPropertyInfo (struct)
type DBusPropertyInfoFlags C.uint32_t
const (
	DBusPropertyInfoFlagsNone DBusPropertyInfoFlags = 0
	DBusPropertyInfoFlagsReadable DBusPropertyInfoFlags = 1
	DBusPropertyInfoFlagsWritable DBusPropertyInfoFlags = 2
)
// blacklisted: DBusProxy (object)
// blacklisted: DBusProxyClass (struct)
type DBusProxyFlags C.uint32_t
const (
	DBusProxyFlagsNone DBusProxyFlags = 0
	DBusProxyFlagsDoNotLoadProperties DBusProxyFlags = 1
	DBusProxyFlagsDoNotConnectSignals DBusProxyFlags = 2
	DBusProxyFlagsDoNotAutoStart DBusProxyFlags = 4
	DBusProxyFlagsGetInvalidatedProperties DBusProxyFlags = 8
)
// blacklisted: DBusProxyPrivate (struct)
// blacklisted: DBusProxyTypeFunc (callback)
type DBusSendMessageFlags C.uint32_t
const (
	DBusSendMessageFlagsNone DBusSendMessageFlags = 0
	DBusSendMessageFlagsPreserveSerial DBusSendMessageFlags = 1
)
// blacklisted: DBusServer (object)
type DBusServerFlags C.uint32_t
const (
	DBusServerFlagsNone DBusServerFlags = 0
	DBusServerFlagsRunInThread DBusServerFlags = 1
	DBusServerFlagsAuthenticationAllowAnonymous DBusServerFlags = 2
)
// blacklisted: DBusSignalCallback (callback)
type DBusSignalFlags C.uint32_t
const (
	DBusSignalFlagsNone DBusSignalFlags = 0
	DBusSignalFlagsNoMatchRule DBusSignalFlags = 1
)
// blacklisted: DBusSignalInfo (struct)
// blacklisted: DBusSubtreeDispatchFunc (callback)
type DBusSubtreeFlags C.uint32_t
const (
	DBusSubtreeFlagsNone DBusSubtreeFlags = 0
	DBusSubtreeFlagsDispatchToUnenumeratedNodes DBusSubtreeFlags = 1
)
// blacklisted: DBusSubtreeIntrospectFunc (callback)
// blacklisted: DBusSubtreeVTable (struct)
const DesktopAppInfoLookupExtensionPointName = "gio-desktop-app-info-lookup"
// blacklisted: DataInputStream (object)
// blacklisted: DataInputStreamClass (struct)
// blacklisted: DataInputStreamPrivate (struct)
// blacklisted: DataOutputStream (object)
// blacklisted: DataOutputStreamClass (struct)
// blacklisted: DataOutputStreamPrivate (struct)
type DataStreamByteOrder C.uint32_t
const (
	DataStreamByteOrderBigEndian DataStreamByteOrder = 0
	DataStreamByteOrderLittleEndian DataStreamByteOrder = 1
	DataStreamByteOrderHostEndian DataStreamByteOrder = 2
)
type DataStreamNewlineType C.uint32_t
const (
	DataStreamNewlineTypeLf DataStreamNewlineType = 0
	DataStreamNewlineTypeCr DataStreamNewlineType = 1
	DataStreamNewlineTypeCrLf DataStreamNewlineType = 2
	DataStreamNewlineTypeAny DataStreamNewlineType = 3
)
// blacklisted: DesktopAppInfo (object)
// blacklisted: DesktopAppInfoClass (struct)
// blacklisted: DesktopAppInfoLookup (interface)
// blacklisted: DesktopAppInfoLookupIface (struct)
// blacklisted: DesktopAppLaunchCallback (callback)
// blacklisted: Drive (interface)
// blacklisted: DriveIface (struct)
type DriveStartFlags C.uint32_t
const (
	DriveStartFlagsNone DriveStartFlags = 0
)
type DriveStartStopType C.uint32_t
const (
	DriveStartStopTypeUnknown DriveStartStopType = 0
	DriveStartStopTypeShutdown DriveStartStopType = 1
	DriveStartStopTypeNetwork DriveStartStopType = 2
	DriveStartStopTypeMultidisk DriveStartStopType = 3
	DriveStartStopTypePassword DriveStartStopType = 4
)
type EmblemLike interface {
	gobject.ObjectLike
	InheritedFromGEmblem() *C.GEmblem
}

type Emblem struct {
	gobject.Object
	IconImpl
}

func ToEmblem(objlike gobject.ObjectLike) *Emblem {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Emblem)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Emblem)(obj)
	}
	panic("cannot cast to Emblem")
}

func (this0 *Emblem) InheritedFromGEmblem() *C.GEmblem {
	if this0 == nil {
		return nil
	}
	return (*C.GEmblem)(this0.C)
}

func (this0 *Emblem) GetStaticType() gobject.Type {
	return gobject.Type(C.g_emblem_get_type())
}

func EmblemGetType() gobject.Type {
	return (*Emblem)(nil).GetStaticType()
}
func NewEmblem(icon0 IconLike) *Emblem {
	var icon1 *C.GIcon
	if icon0 != nil {
		icon1 = icon0.ImplementsGIcon()
	}
	ret1 := C.g_emblem_new(icon1)
	var ret2 *Emblem
	ret2 = (*Emblem)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func NewEmblemWithOrigin(icon0 IconLike, origin0 EmblemOrigin) *Emblem {
	var icon1 *C.GIcon
	var origin1 C.GEmblemOrigin
	if icon0 != nil {
		icon1 = icon0.ImplementsGIcon()
	}
	origin1 = C.GEmblemOrigin(origin0)
	ret1 := C.g_emblem_new_with_origin(icon1, origin1)
	var ret2 *Emblem
	ret2 = (*Emblem)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Emblem) GetIcon() *Icon {
	var this1 *C.GEmblem
	if this0 != nil {
		this1 = this0.InheritedFromGEmblem()
	}
	ret1 := C.g_emblem_get_icon(this1)
	var ret2 *Icon
	ret2 = (*Icon)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Emblem) GetOrigin() EmblemOrigin {
	var this1 *C.GEmblem
	if this0 != nil {
		this1 = this0.InheritedFromGEmblem()
	}
	ret1 := C.g_emblem_get_origin(this1)
	var ret2 EmblemOrigin
	ret2 = EmblemOrigin(ret1)
	return ret2
}
// blacklisted: EmblemClass (struct)
type EmblemOrigin C.uint32_t
const (
	EmblemOriginUnknown EmblemOrigin = 0
	EmblemOriginDevice EmblemOrigin = 1
	EmblemOriginLivemetadata EmblemOrigin = 2
	EmblemOriginTag EmblemOrigin = 3
)
type EmblemedIconLike interface {
	gobject.ObjectLike
	InheritedFromGEmblemedIcon() *C.GEmblemedIcon
}

type EmblemedIcon struct {
	gobject.Object
	IconImpl
}

func ToEmblemedIcon(objlike gobject.ObjectLike) *EmblemedIcon {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*EmblemedIcon)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*EmblemedIcon)(obj)
	}
	panic("cannot cast to EmblemedIcon")
}

func (this0 *EmblemedIcon) InheritedFromGEmblemedIcon() *C.GEmblemedIcon {
	if this0 == nil {
		return nil
	}
	return (*C.GEmblemedIcon)(this0.C)
}

func (this0 *EmblemedIcon) GetStaticType() gobject.Type {
	return gobject.Type(C.g_emblemed_icon_get_type())
}

func EmblemedIconGetType() gobject.Type {
	return (*EmblemedIcon)(nil).GetStaticType()
}
func NewEmblemedIcon(icon0 IconLike, emblem0 EmblemLike) *EmblemedIcon {
	var icon1 *C.GIcon
	var emblem1 *C.GEmblem
	if icon0 != nil {
		icon1 = icon0.ImplementsGIcon()
	}
	if emblem0 != nil {
		emblem1 = emblem0.InheritedFromGEmblem()
	}
	ret1 := C.g_emblemed_icon_new(icon1, emblem1)
	var ret2 *EmblemedIcon
	ret2 = (*EmblemedIcon)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *EmblemedIcon) AddEmblem(emblem0 EmblemLike) {
	var this1 *C.GEmblemedIcon
	var emblem1 *C.GEmblem
	if this0 != nil {
		this1 = this0.InheritedFromGEmblemedIcon()
	}
	if emblem0 != nil {
		emblem1 = emblem0.InheritedFromGEmblem()
	}
	C.g_emblemed_icon_add_emblem(this1, emblem1)
}
func (this0 *EmblemedIcon) ClearEmblems() {
	var this1 *C.GEmblemedIcon
	if this0 != nil {
		this1 = this0.InheritedFromGEmblemedIcon()
	}
	C.g_emblemed_icon_clear_emblems(this1)
}
func (this0 *EmblemedIcon) GetEmblems() []*Emblem {
	var this1 *C.GEmblemedIcon
	if this0 != nil {
		this1 = this0.InheritedFromGEmblemedIcon()
	}
	ret1 := C.g_emblemed_icon_get_emblems(this1)
	var ret2 []*Emblem
	for iter := (*_GList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt *Emblem
		elt = (*Emblem)(gobject.ObjectWrap(unsafe.Pointer((*C.GEmblem)(iter.data)), true))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func (this0 *EmblemedIcon) GetIcon() *Icon {
	var this1 *C.GEmblemedIcon
	if this0 != nil {
		this1 = this0.InheritedFromGEmblemedIcon()
	}
	ret1 := C.g_emblemed_icon_get_icon(this1)
	var ret2 *Icon
	ret2 = (*Icon)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
// blacklisted: EmblemedIconClass (struct)
// blacklisted: EmblemedIconPrivate (struct)
const FileAttributeAccessCanDelete = "access::can-delete"
const FileAttributeAccessCanExecute = "access::can-execute"
const FileAttributeAccessCanRead = "access::can-read"
const FileAttributeAccessCanRename = "access::can-rename"
const FileAttributeAccessCanTrash = "access::can-trash"
const FileAttributeAccessCanWrite = "access::can-write"
const FileAttributeDosIsArchive = "dos::is-archive"
const FileAttributeDosIsSystem = "dos::is-system"
const FileAttributeEtagValue = "etag::value"
const FileAttributeFilesystemFree = "filesystem::free"
const FileAttributeFilesystemReadonly = "filesystem::readonly"
const FileAttributeFilesystemSize = "filesystem::size"
const FileAttributeFilesystemType = "filesystem::type"
const FileAttributeFilesystemUsed = "filesystem::used"
const FileAttributeFilesystemUsePreview = "filesystem::use-preview"
const FileAttributeGvfsBackend = "gvfs::backend"
const FileAttributeIdFile = "id::file"
const FileAttributeIdFilesystem = "id::filesystem"
const FileAttributeMountableCanEject = "mountable::can-eject"
const FileAttributeMountableCanMount = "mountable::can-mount"
const FileAttributeMountableCanPoll = "mountable::can-poll"
const FileAttributeMountableCanStart = "mountable::can-start"
const FileAttributeMountableCanStartDegraded = "mountable::can-start-degraded"
const FileAttributeMountableCanStop = "mountable::can-stop"
const FileAttributeMountableCanUnmount = "mountable::can-unmount"
const FileAttributeMountableHalUdi = "mountable::hal-udi"
const FileAttributeMountableIsMediaCheckAutomatic = "mountable::is-media-check-automatic"
const FileAttributeMountableStartStopType = "mountable::start-stop-type"
const FileAttributeMountableUnixDevice = "mountable::unix-device"
const FileAttributeMountableUnixDeviceFile = "mountable::unix-device-file"
const FileAttributeOwnerGroup = "owner::group"
const FileAttributeOwnerUser = "owner::user"
const FileAttributeOwnerUserReal = "owner::user-real"
const FileAttributePreviewIcon = "preview::icon"
const FileAttributeSelinuxContext = "selinux::context"
const FileAttributeStandardAllocatedSize = "standard::allocated-size"
const FileAttributeStandardContentType = "standard::content-type"
const FileAttributeStandardCopyName = "standard::copy-name"
const FileAttributeStandardDescription = "standard::description"
const FileAttributeStandardDisplayName = "standard::display-name"
const FileAttributeStandardEditName = "standard::edit-name"
const FileAttributeStandardFastContentType = "standard::fast-content-type"
const FileAttributeStandardIcon = "standard::icon"
const FileAttributeStandardIsBackup = "standard::is-backup"
const FileAttributeStandardIsHidden = "standard::is-hidden"
const FileAttributeStandardIsSymlink = "standard::is-symlink"
const FileAttributeStandardIsVirtual = "standard::is-virtual"
const FileAttributeStandardName = "standard::name"
const FileAttributeStandardSize = "standard::size"
const FileAttributeStandardSortOrder = "standard::sort-order"
const FileAttributeStandardSymbolicIcon = "standard::symbolic-icon"
const FileAttributeStandardSymlinkTarget = "standard::symlink-target"
const FileAttributeStandardTargetUri = "standard::target-uri"
const FileAttributeStandardType = "standard::type"
const FileAttributeThumbnailingFailed = "thumbnail::failed"
const FileAttributeThumbnailPath = "thumbnail::path"
const FileAttributeTimeAccess = "time::access"
const FileAttributeTimeAccessUsec = "time::access-usec"
const FileAttributeTimeChanged = "time::changed"
const FileAttributeTimeChangedUsec = "time::changed-usec"
const FileAttributeTimeCreated = "time::created"
const FileAttributeTimeCreatedUsec = "time::created-usec"
const FileAttributeTimeModified = "time::modified"
const FileAttributeTimeModifiedUsec = "time::modified-usec"
const FileAttributeTrashDeletionDate = "trash::deletion-date"
const FileAttributeTrashItemCount = "trash::item-count"
const FileAttributeTrashOrigPath = "trash::orig-path"
const FileAttributeUnixBlocks = "unix::blocks"
const FileAttributeUnixBlockSize = "unix::block-size"
const FileAttributeUnixDevice = "unix::device"
const FileAttributeUnixGid = "unix::gid"
const FileAttributeUnixInode = "unix::inode"
const FileAttributeUnixIsMountpoint = "unix::is-mountpoint"
const FileAttributeUnixMode = "unix::mode"
const FileAttributeUnixNlink = "unix::nlink"
const FileAttributeUnixRdev = "unix::rdev"
const FileAttributeUnixUid = "unix::uid"
type FileLike interface {
	ImplementsGFile() *C.GFile
}

type File struct {
	gobject.Object
	FileImpl
}

type FileImpl struct {}

func ToFile(objlike gobject.ObjectLike) *File {
	t := (*FileImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*File)(obj)
	}
	panic("cannot cast to File")
}

func (this0 *FileImpl) ImplementsGFile() *C.GFile {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GFile)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *FileImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.g_file_get_type())
}

func FileGetType() gobject.Type {
	return (*FileImpl)(nil).GetStaticType()
}
// blacklisted: File.new_for_commandline_arg (method)
// blacklisted: File.new_for_commandline_arg_and_cwd (method)
// blacklisted: File.new_for_path (method)
// blacklisted: File.new_for_uri (method)
// blacklisted: File.new_tmp (method)
// blacklisted: File.parse_name (method)
// blacklisted: File.append_to (method)
// blacklisted: File.append_to_async (method)
// blacklisted: File.append_to_finish (method)
// blacklisted: File.copy (method)
// blacklisted: File.copy_attributes (method)
// blacklisted: File.copy_finish (method)
// blacklisted: File.create (method)
// blacklisted: File.create_async (method)
// blacklisted: File.create_finish (method)
// blacklisted: File.create_readwrite (method)
// blacklisted: File.create_readwrite_async (method)
// blacklisted: File.create_readwrite_finish (method)
// blacklisted: File.delete (method)
// blacklisted: File.delete_async (method)
// blacklisted: File.delete_finish (method)
// blacklisted: File.dup (method)
// blacklisted: File.eject_mountable (method)
// blacklisted: File.eject_mountable_finish (method)
// blacklisted: File.eject_mountable_with_operation (method)
// blacklisted: File.eject_mountable_with_operation_finish (method)
// blacklisted: File.enumerate_children (method)
// blacklisted: File.enumerate_children_async (method)
// blacklisted: File.enumerate_children_finish (method)
// blacklisted: File.equal (method)
// blacklisted: File.find_enclosing_mount (method)
// blacklisted: File.find_enclosing_mount_async (method)
// blacklisted: File.find_enclosing_mount_finish (method)
// blacklisted: File.get_basename (method)
// blacklisted: File.get_child (method)
// blacklisted: File.get_child_for_display_name (method)
// blacklisted: File.get_parent (method)
// blacklisted: File.get_parse_name (method)
// blacklisted: File.get_path (method)
// blacklisted: File.get_relative_path (method)
// blacklisted: File.get_uri (method)
// blacklisted: File.get_uri_scheme (method)
// blacklisted: File.has_parent (method)
// blacklisted: File.has_prefix (method)
// blacklisted: File.has_uri_scheme (method)
// blacklisted: File.hash (method)
// blacklisted: File.is_native (method)
// blacklisted: File.load_contents (method)
// blacklisted: File.load_contents_async (method)
// blacklisted: File.load_contents_finish (method)
// blacklisted: File.load_partial_contents_finish (method)
// blacklisted: File.make_directory (method)
// blacklisted: File.make_directory_with_parents (method)
// blacklisted: File.make_symbolic_link (method)
// blacklisted: File.monitor (method)
// blacklisted: File.monitor_directory (method)
// blacklisted: File.monitor_file (method)
// blacklisted: File.mount_enclosing_volume (method)
// blacklisted: File.mount_enclosing_volume_finish (method)
// blacklisted: File.mount_mountable (method)
// blacklisted: File.mount_mountable_finish (method)
// blacklisted: File.move (method)
// blacklisted: File.open_readwrite (method)
// blacklisted: File.open_readwrite_async (method)
// blacklisted: File.open_readwrite_finish (method)
// blacklisted: File.poll_mountable (method)
// blacklisted: File.poll_mountable_finish (method)
// blacklisted: File.query_default_handler (method)
// blacklisted: File.query_exists (method)
// blacklisted: File.query_file_type (method)
// blacklisted: File.query_filesystem_info (method)
// blacklisted: File.query_filesystem_info_async (method)
// blacklisted: File.query_filesystem_info_finish (method)
// blacklisted: File.query_info (method)
// blacklisted: File.query_info_async (method)
// blacklisted: File.query_info_finish (method)
// blacklisted: File.query_settable_attributes (method)
// blacklisted: File.query_writable_namespaces (method)
// blacklisted: File.read (method)
// blacklisted: File.read_async (method)
// blacklisted: File.read_finish (method)
// blacklisted: File.replace (method)
// blacklisted: File.replace_async (method)
// blacklisted: File.replace_contents (method)
// blacklisted: File.replace_contents_async (method)
// blacklisted: File.replace_contents_finish (method)
// blacklisted: File.replace_finish (method)
// blacklisted: File.replace_readwrite (method)
// blacklisted: File.replace_readwrite_async (method)
// blacklisted: File.replace_readwrite_finish (method)
// blacklisted: File.resolve_relative_path (method)
// blacklisted: File.set_attribute (method)
// blacklisted: File.set_attribute_byte_string (method)
// blacklisted: File.set_attribute_int32 (method)
// blacklisted: File.set_attribute_int64 (method)
// blacklisted: File.set_attribute_string (method)
// blacklisted: File.set_attribute_uint32 (method)
// blacklisted: File.set_attribute_uint64 (method)
// blacklisted: File.set_attributes_async (method)
// blacklisted: File.set_attributes_finish (method)
// blacklisted: File.set_attributes_from_info (method)
// blacklisted: File.set_display_name (method)
// blacklisted: File.set_display_name_async (method)
// blacklisted: File.set_display_name_finish (method)
// blacklisted: File.start_mountable (method)
// blacklisted: File.start_mountable_finish (method)
// blacklisted: File.stop_mountable (method)
// blacklisted: File.stop_mountable_finish (method)
// blacklisted: File.supports_thread_contexts (method)
// blacklisted: File.trash (method)
// blacklisted: File.unmount_mountable (method)
// blacklisted: File.unmount_mountable_finish (method)
// blacklisted: File.unmount_mountable_with_operation (method)
// blacklisted: File.unmount_mountable_with_operation_finish (method)
// blacklisted: FileAttributeInfo (struct)
type FileAttributeInfoFlags C.uint32_t
const (
	FileAttributeInfoFlagsNone FileAttributeInfoFlags = 0
	FileAttributeInfoFlagsCopyWithFile FileAttributeInfoFlags = 1
	FileAttributeInfoFlagsCopyWhenMoved FileAttributeInfoFlags = 2
)
// blacklisted: FileAttributeInfoList (struct)
// blacklisted: FileAttributeMatcher (struct)
type FileAttributeStatus C.uint32_t
const (
	FileAttributeStatusUnset FileAttributeStatus = 0
	FileAttributeStatusSet FileAttributeStatus = 1
	FileAttributeStatusErrorSetting FileAttributeStatus = 2
)
type FileAttributeType C.uint32_t
const (
	FileAttributeTypeInvalid FileAttributeType = 0
	FileAttributeTypeString FileAttributeType = 1
	FileAttributeTypeByteString FileAttributeType = 2
	FileAttributeTypeBoolean FileAttributeType = 3
	FileAttributeTypeUint32 FileAttributeType = 4
	FileAttributeTypeInt32 FileAttributeType = 5
	FileAttributeTypeUint64 FileAttributeType = 6
	FileAttributeTypeInt64 FileAttributeType = 7
	FileAttributeTypeObject FileAttributeType = 8
	FileAttributeTypeStringv FileAttributeType = 9
)
type FileCopyFlags C.uint32_t
const (
	FileCopyFlagsNone FileCopyFlags = 0
	FileCopyFlagsOverwrite FileCopyFlags = 1
	FileCopyFlagsBackup FileCopyFlags = 2
	FileCopyFlagsNofollowSymlinks FileCopyFlags = 4
	FileCopyFlagsAllMetadata FileCopyFlags = 8
	FileCopyFlagsNoFallbackForMove FileCopyFlags = 16
	FileCopyFlagsTargetDefaultPerms FileCopyFlags = 32
)
type FileCreateFlags C.uint32_t
const (
	FileCreateFlagsNone FileCreateFlags = 0
	FileCreateFlagsPrivate FileCreateFlags = 1
	FileCreateFlagsReplaceDestination FileCreateFlags = 2
)
// blacklisted: FileDescriptorBased (interface)
// blacklisted: FileDescriptorBasedIface (struct)
// blacklisted: FileEnumerator (object)
// blacklisted: FileEnumeratorClass (struct)
// blacklisted: FileEnumeratorPrivate (struct)
// blacklisted: FileIOStream (object)
// blacklisted: FileIOStreamClass (struct)
// blacklisted: FileIOStreamPrivate (struct)
// blacklisted: FileIcon (object)
// blacklisted: FileIconClass (struct)
// blacklisted: FileIface (struct)
// blacklisted: FileInfo (object)
// blacklisted: FileInfoClass (struct)
// blacklisted: FileInputStream (object)
// blacklisted: FileInputStreamClass (struct)
// blacklisted: FileInputStreamPrivate (struct)
// blacklisted: FileMonitor (object)
// blacklisted: FileMonitorClass (struct)
type FileMonitorEvent C.uint32_t
const (
	FileMonitorEventChanged FileMonitorEvent = 0
	FileMonitorEventChangesDoneHint FileMonitorEvent = 1
	FileMonitorEventDeleted FileMonitorEvent = 2
	FileMonitorEventCreated FileMonitorEvent = 3
	FileMonitorEventAttributeChanged FileMonitorEvent = 4
	FileMonitorEventPreUnmount FileMonitorEvent = 5
	FileMonitorEventUnmounted FileMonitorEvent = 6
	FileMonitorEventMoved FileMonitorEvent = 7
)
type FileMonitorFlags C.uint32_t
const (
	FileMonitorFlagsNone FileMonitorFlags = 0
	FileMonitorFlagsWatchMounts FileMonitorFlags = 1
	FileMonitorFlagsSendMoved FileMonitorFlags = 2
	FileMonitorFlagsWatchHardLinks FileMonitorFlags = 4
)
// blacklisted: FileMonitorPrivate (struct)
// blacklisted: FileOutputStream (object)
// blacklisted: FileOutputStreamClass (struct)
// blacklisted: FileOutputStreamPrivate (struct)
// blacklisted: FileProgressCallback (callback)
type FileQueryInfoFlags C.uint32_t
const (
	FileQueryInfoFlagsNone FileQueryInfoFlags = 0
	FileQueryInfoFlagsNofollowSymlinks FileQueryInfoFlags = 1
)
// blacklisted: FileReadMoreCallback (callback)
type FileType C.uint32_t
const (
	FileTypeUnknown FileType = 0
	FileTypeRegular FileType = 1
	FileTypeDirectory FileType = 2
	FileTypeSymbolicLink FileType = 3
	FileTypeSpecial FileType = 4
	FileTypeShortcut FileType = 5
	FileTypeMountable FileType = 6
)
// blacklisted: FilenameCompleter (object)
// blacklisted: FilenameCompleterClass (struct)
type FilesystemPreviewType C.uint32_t
const (
	FilesystemPreviewTypeIfAlways FilesystemPreviewType = 0
	FilesystemPreviewTypeIfLocal FilesystemPreviewType = 1
	FilesystemPreviewTypeNever FilesystemPreviewType = 2
)
// blacklisted: FilterInputStream (object)
// blacklisted: FilterInputStreamClass (struct)
// blacklisted: FilterOutputStream (object)
// blacklisted: FilterOutputStreamClass (struct)
type IOErrorEnum C.uint32_t
const (
	IOErrorEnumFailed IOErrorEnum = 0
	IOErrorEnumNotFound IOErrorEnum = 1
	IOErrorEnumExists IOErrorEnum = 2
	IOErrorEnumIsDirectory IOErrorEnum = 3
	IOErrorEnumNotDirectory IOErrorEnum = 4
	IOErrorEnumNotEmpty IOErrorEnum = 5
	IOErrorEnumNotRegularFile IOErrorEnum = 6
	IOErrorEnumNotSymbolicLink IOErrorEnum = 7
	IOErrorEnumNotMountableFile IOErrorEnum = 8
	IOErrorEnumFilenameTooLong IOErrorEnum = 9
	IOErrorEnumInvalidFilename IOErrorEnum = 10
	IOErrorEnumTooManyLinks IOErrorEnum = 11
	IOErrorEnumNoSpace IOErrorEnum = 12
	IOErrorEnumInvalidArgument IOErrorEnum = 13
	IOErrorEnumPermissionDenied IOErrorEnum = 14
	IOErrorEnumNotSupported IOErrorEnum = 15
	IOErrorEnumNotMounted IOErrorEnum = 16
	IOErrorEnumAlreadyMounted IOErrorEnum = 17
	IOErrorEnumClosed IOErrorEnum = 18
	IOErrorEnumCancelled IOErrorEnum = 19
	IOErrorEnumPending IOErrorEnum = 20
	IOErrorEnumReadOnly IOErrorEnum = 21
	IOErrorEnumCantCreateBackup IOErrorEnum = 22
	IOErrorEnumWrongEtag IOErrorEnum = 23
	IOErrorEnumTimedOut IOErrorEnum = 24
	IOErrorEnumWouldRecurse IOErrorEnum = 25
	IOErrorEnumBusy IOErrorEnum = 26
	IOErrorEnumWouldBlock IOErrorEnum = 27
	IOErrorEnumHostNotFound IOErrorEnum = 28
	IOErrorEnumWouldMerge IOErrorEnum = 29
	IOErrorEnumFailedHandled IOErrorEnum = 30
	IOErrorEnumTooManyOpenFiles IOErrorEnum = 31
	IOErrorEnumNotInitialized IOErrorEnum = 32
	IOErrorEnumAddressInUse IOErrorEnum = 33
	IOErrorEnumPartialInput IOErrorEnum = 34
	IOErrorEnumInvalidData IOErrorEnum = 35
	IOErrorEnumDbusError IOErrorEnum = 36
	IOErrorEnumHostUnreachable IOErrorEnum = 37
	IOErrorEnumNetworkUnreachable IOErrorEnum = 38
	IOErrorEnumConnectionRefused IOErrorEnum = 39
	IOErrorEnumProxyFailed IOErrorEnum = 40
	IOErrorEnumProxyAuthFailed IOErrorEnum = 41
	IOErrorEnumProxyNeedAuth IOErrorEnum = 42
	IOErrorEnumProxyNotAllowed IOErrorEnum = 43
	IOErrorEnumBrokenPipe IOErrorEnum = 44
)
// blacklisted: IOExtension (struct)
// blacklisted: IOExtensionPoint (struct)
// blacklisted: IOModule (object)
// blacklisted: IOModuleClass (struct)
// blacklisted: IOModuleScope (struct)
type IOModuleScopeFlags C.uint32_t
const (
	IOModuleScopeFlagsNone IOModuleScopeFlags = 0
	IOModuleScopeFlagsBlockDuplicates IOModuleScopeFlags = 1
)
// blacklisted: IOSchedulerJob (struct)
// blacklisted: IOSchedulerJobFunc (callback)
// blacklisted: IOStream (object)
// blacklisted: IOStreamAdapter (struct)
// blacklisted: IOStreamClass (struct)
// blacklisted: IOStreamPrivate (struct)
type IOStreamSpliceFlags C.uint32_t
const (
	IOStreamSpliceFlagsNone IOStreamSpliceFlags = 0
	IOStreamSpliceFlagsCloseStream1 IOStreamSpliceFlags = 1
	IOStreamSpliceFlagsCloseStream2 IOStreamSpliceFlags = 2
	IOStreamSpliceFlagsWaitForBoth IOStreamSpliceFlags = 4
)
type IconLike interface {
	ImplementsGIcon() *C.GIcon
}

type Icon struct {
	gobject.Object
	IconImpl
}

type IconImpl struct {}

func ToIcon(objlike gobject.ObjectLike) *Icon {
	t := (*IconImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Icon)(obj)
	}
	panic("cannot cast to Icon")
}

func (this0 *IconImpl) ImplementsGIcon() *C.GIcon {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GIcon)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *IconImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.g_icon_get_type())
}

func IconGetType() gobject.Type {
	return (*IconImpl)(nil).GetStaticType()
}
// blacklisted: Icon.hash (method)
// blacklisted: Icon.new_for_string (method)
// blacklisted: Icon.equal (method)
// blacklisted: Icon.to_string (method)
// blacklisted: IconIface (struct)
// blacklisted: InetAddress (object)
// blacklisted: InetAddressClass (struct)
// blacklisted: InetAddressMask (object)
// blacklisted: InetAddressMaskClass (struct)
// blacklisted: InetAddressMaskPrivate (struct)
// blacklisted: InetAddressPrivate (struct)
// blacklisted: InetSocketAddress (object)
// blacklisted: InetSocketAddressClass (struct)
// blacklisted: InetSocketAddressPrivate (struct)
// blacklisted: Initable (interface)
// blacklisted: InitableIface (struct)
// blacklisted: InputStream (object)
// blacklisted: InputStreamClass (struct)
// blacklisted: InputStreamPrivate (struct)
// blacklisted: InputVector (struct)
// blacklisted: LoadableIcon (interface)
// blacklisted: LoadableIconIface (struct)
const MenuAttributeAction = "action"
const MenuAttributeActionNamespace = "action-namespace"
const MenuAttributeLabel = "label"
const MenuAttributeTarget = "target"
const MenuLinkSection = "section"
const MenuLinkSubmenu = "submenu"
// blacklisted: MemoryInputStream (object)
// blacklisted: MemoryInputStreamClass (struct)
// blacklisted: MemoryInputStreamPrivate (struct)
// blacklisted: MemoryOutputStream (object)
// blacklisted: MemoryOutputStreamClass (struct)
// blacklisted: MemoryOutputStreamPrivate (struct)
// blacklisted: Menu (object)
type MenuAttributeIterLike interface {
	gobject.ObjectLike
	InheritedFromGMenuAttributeIter() *C.GMenuAttributeIter
}

type MenuAttributeIter struct {
	gobject.Object
	
}

func ToMenuAttributeIter(objlike gobject.ObjectLike) *MenuAttributeIter {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*MenuAttributeIter)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*MenuAttributeIter)(obj)
	}
	panic("cannot cast to MenuAttributeIter")
}

func (this0 *MenuAttributeIter) InheritedFromGMenuAttributeIter() *C.GMenuAttributeIter {
	if this0 == nil {
		return nil
	}
	return (*C.GMenuAttributeIter)(this0.C)
}

func (this0 *MenuAttributeIter) GetStaticType() gobject.Type {
	return gobject.Type(C.g_menu_attribute_iter_get_type())
}

func MenuAttributeIterGetType() gobject.Type {
	return (*MenuAttributeIter)(nil).GetStaticType()
}
func (this0 *MenuAttributeIter) GetName() string {
	var this1 *C.GMenuAttributeIter
	if this0 != nil {
		this1 = this0.InheritedFromGMenuAttributeIter()
	}
	ret1 := C.g_menu_attribute_iter_get_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *MenuAttributeIter) GetNext() (string, *glib.Variant, bool) {
	var this1 *C.GMenuAttributeIter
	var out_name1 *C.char
	var value1 *C.GVariant
	if this0 != nil {
		this1 = this0.InheritedFromGMenuAttributeIter()
	}
	ret1 := C.g_menu_attribute_iter_get_next(this1, &out_name1, &value1)
	var out_name2 string
	var value2 *glib.Variant
	var ret2 bool
	out_name2 = C.GoString(out_name1)
	value2 = (*glib.Variant)(unsafe.Pointer(value1))
	ret2 = ret1 != 0
	return out_name2, value2, ret2
}
func (this0 *MenuAttributeIter) GetValue() *glib.Variant {
	var this1 *C.GMenuAttributeIter
	if this0 != nil {
		this1 = this0.InheritedFromGMenuAttributeIter()
	}
	ret1 := C.g_menu_attribute_iter_get_value(this1)
	var ret2 *glib.Variant
	ret2 = (*glib.Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *MenuAttributeIter) Next() bool {
	var this1 *C.GMenuAttributeIter
	if this0 != nil {
		this1 = this0.InheritedFromGMenuAttributeIter()
	}
	ret1 := C.g_menu_attribute_iter_next(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
// blacklisted: MenuAttributeIterClass (struct)
// blacklisted: MenuAttributeIterPrivate (struct)
// blacklisted: MenuItem (object)
type MenuLinkIterLike interface {
	gobject.ObjectLike
	InheritedFromGMenuLinkIter() *C.GMenuLinkIter
}

type MenuLinkIter struct {
	gobject.Object
	
}

func ToMenuLinkIter(objlike gobject.ObjectLike) *MenuLinkIter {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*MenuLinkIter)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*MenuLinkIter)(obj)
	}
	panic("cannot cast to MenuLinkIter")
}

func (this0 *MenuLinkIter) InheritedFromGMenuLinkIter() *C.GMenuLinkIter {
	if this0 == nil {
		return nil
	}
	return (*C.GMenuLinkIter)(this0.C)
}

func (this0 *MenuLinkIter) GetStaticType() gobject.Type {
	return gobject.Type(C.g_menu_link_iter_get_type())
}

func MenuLinkIterGetType() gobject.Type {
	return (*MenuLinkIter)(nil).GetStaticType()
}
func (this0 *MenuLinkIter) GetName() string {
	var this1 *C.GMenuLinkIter
	if this0 != nil {
		this1 = this0.InheritedFromGMenuLinkIter()
	}
	ret1 := C.g_menu_link_iter_get_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *MenuLinkIter) GetNext() (string, *MenuModel, bool) {
	var this1 *C.GMenuLinkIter
	var out_link1 *C.char
	var value1 *C.GMenuModel
	if this0 != nil {
		this1 = this0.InheritedFromGMenuLinkIter()
	}
	ret1 := C.g_menu_link_iter_get_next(this1, &out_link1, &value1)
	var out_link2 string
	var value2 *MenuModel
	var ret2 bool
	out_link2 = C.GoString(out_link1)
	value2 = (*MenuModel)(gobject.ObjectWrap(unsafe.Pointer(value1), false))
	ret2 = ret1 != 0
	return out_link2, value2, ret2
}
func (this0 *MenuLinkIter) GetValue() *MenuModel {
	var this1 *C.GMenuLinkIter
	if this0 != nil {
		this1 = this0.InheritedFromGMenuLinkIter()
	}
	ret1 := C.g_menu_link_iter_get_value(this1)
	var ret2 *MenuModel
	ret2 = (*MenuModel)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *MenuLinkIter) Next() bool {
	var this1 *C.GMenuLinkIter
	if this0 != nil {
		this1 = this0.InheritedFromGMenuLinkIter()
	}
	ret1 := C.g_menu_link_iter_next(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
// blacklisted: MenuLinkIterClass (struct)
// blacklisted: MenuLinkIterPrivate (struct)
type MenuModelLike interface {
	gobject.ObjectLike
	InheritedFromGMenuModel() *C.GMenuModel
}

type MenuModel struct {
	gobject.Object
	
}

func ToMenuModel(objlike gobject.ObjectLike) *MenuModel {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*MenuModel)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*MenuModel)(obj)
	}
	panic("cannot cast to MenuModel")
}

func (this0 *MenuModel) InheritedFromGMenuModel() *C.GMenuModel {
	if this0 == nil {
		return nil
	}
	return (*C.GMenuModel)(this0.C)
}

func (this0 *MenuModel) GetStaticType() gobject.Type {
	return gobject.Type(C.g_menu_model_get_type())
}

func MenuModelGetType() gobject.Type {
	return (*MenuModel)(nil).GetStaticType()
}
func (this0 *MenuModel) GetItemAttributeValue(item_index0 int, attribute0 string, expected_type0 *glib.VariantType) *glib.Variant {
	var this1 *C.GMenuModel
	var item_index1 C.int32_t
	var attribute1 *C.char
	var expected_type1 *C.GVariantType
	if this0 != nil {
		this1 = this0.InheritedFromGMenuModel()
	}
	item_index1 = C.int32_t(item_index0)
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	expected_type1 = (*C.GVariantType)(unsafe.Pointer(expected_type0))
	ret1 := C.g_menu_model_get_item_attribute_value(this1, item_index1, attribute1, expected_type1)
	var ret2 *glib.Variant
	ret2 = (*glib.Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *MenuModel) GetItemLink(item_index0 int, link0 string) *MenuModel {
	var this1 *C.GMenuModel
	var item_index1 C.int32_t
	var link1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGMenuModel()
	}
	item_index1 = C.int32_t(item_index0)
	link1 = _GoStringToGString(link0)
	defer C.free(unsafe.Pointer(link1))
	ret1 := C.g_menu_model_get_item_link(this1, item_index1, link1)
	var ret2 *MenuModel
	ret2 = (*MenuModel)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *MenuModel) GetNItems() int {
	var this1 *C.GMenuModel
	if this0 != nil {
		this1 = this0.InheritedFromGMenuModel()
	}
	ret1 := C.g_menu_model_get_n_items(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *MenuModel) IsMutable() bool {
	var this1 *C.GMenuModel
	if this0 != nil {
		this1 = this0.InheritedFromGMenuModel()
	}
	ret1 := C.g_menu_model_is_mutable(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *MenuModel) ItemsChanged(position0 int, removed0 int, added0 int) {
	var this1 *C.GMenuModel
	var position1 C.int32_t
	var removed1 C.int32_t
	var added1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGMenuModel()
	}
	position1 = C.int32_t(position0)
	removed1 = C.int32_t(removed0)
	added1 = C.int32_t(added0)
	C.g_menu_model_items_changed(this1, position1, removed1, added1)
}
func (this0 *MenuModel) IterateItemAttributes(item_index0 int) *MenuAttributeIter {
	var this1 *C.GMenuModel
	var item_index1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGMenuModel()
	}
	item_index1 = C.int32_t(item_index0)
	ret1 := C.g_menu_model_iterate_item_attributes(this1, item_index1)
	var ret2 *MenuAttributeIter
	ret2 = (*MenuAttributeIter)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *MenuModel) IterateItemLinks(item_index0 int) *MenuLinkIter {
	var this1 *C.GMenuModel
	var item_index1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGMenuModel()
	}
	item_index1 = C.int32_t(item_index0)
	ret1 := C.g_menu_model_iterate_item_links(this1, item_index1)
	var ret2 *MenuLinkIter
	ret2 = (*MenuLinkIter)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
// blacklisted: MenuModelClass (struct)
// blacklisted: MenuModelPrivate (struct)
// blacklisted: Mount (interface)
// blacklisted: MountIface (struct)
type MountMountFlags C.uint32_t
const (
	MountMountFlagsNone MountMountFlags = 0
)
type MountOperationLike interface {
	gobject.ObjectLike
	InheritedFromGMountOperation() *C.GMountOperation
}

type MountOperation struct {
	gobject.Object
	
}

func ToMountOperation(objlike gobject.ObjectLike) *MountOperation {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*MountOperation)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*MountOperation)(obj)
	}
	panic("cannot cast to MountOperation")
}

func (this0 *MountOperation) InheritedFromGMountOperation() *C.GMountOperation {
	if this0 == nil {
		return nil
	}
	return (*C.GMountOperation)(this0.C)
}

func (this0 *MountOperation) GetStaticType() gobject.Type {
	return gobject.Type(C.g_mount_operation_get_type())
}

func MountOperationGetType() gobject.Type {
	return (*MountOperation)(nil).GetStaticType()
}
func NewMountOperation() *MountOperation {
	ret1 := C.g_mount_operation_new()
	var ret2 *MountOperation
	ret2 = (*MountOperation)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *MountOperation) GetAnonymous() bool {
	var this1 *C.GMountOperation
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	ret1 := C.g_mount_operation_get_anonymous(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *MountOperation) GetChoice() int {
	var this1 *C.GMountOperation
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	ret1 := C.g_mount_operation_get_choice(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *MountOperation) GetDomain() string {
	var this1 *C.GMountOperation
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	ret1 := C.g_mount_operation_get_domain(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *MountOperation) GetPassword() string {
	var this1 *C.GMountOperation
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	ret1 := C.g_mount_operation_get_password(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *MountOperation) GetPasswordSave() PasswordSave {
	var this1 *C.GMountOperation
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	ret1 := C.g_mount_operation_get_password_save(this1)
	var ret2 PasswordSave
	ret2 = PasswordSave(ret1)
	return ret2
}
func (this0 *MountOperation) GetUsername() string {
	var this1 *C.GMountOperation
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	ret1 := C.g_mount_operation_get_username(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *MountOperation) Reply(result0 MountOperationResult) {
	var this1 *C.GMountOperation
	var result1 C.GMountOperationResult
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	result1 = C.GMountOperationResult(result0)
	C.g_mount_operation_reply(this1, result1)
}
func (this0 *MountOperation) SetAnonymous(anonymous0 bool) {
	var this1 *C.GMountOperation
	var anonymous1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	anonymous1 = _GoBoolToCBool(anonymous0)
	C.g_mount_operation_set_anonymous(this1, anonymous1)
}
func (this0 *MountOperation) SetChoice(choice0 int) {
	var this1 *C.GMountOperation
	var choice1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	choice1 = C.int32_t(choice0)
	C.g_mount_operation_set_choice(this1, choice1)
}
func (this0 *MountOperation) SetDomain(domain0 string) {
	var this1 *C.GMountOperation
	var domain1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	domain1 = _GoStringToGString(domain0)
	defer C.free(unsafe.Pointer(domain1))
	C.g_mount_operation_set_domain(this1, domain1)
}
func (this0 *MountOperation) SetPassword(password0 string) {
	var this1 *C.GMountOperation
	var password1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	password1 = _GoStringToGString(password0)
	defer C.free(unsafe.Pointer(password1))
	C.g_mount_operation_set_password(this1, password1)
}
func (this0 *MountOperation) SetPasswordSave(save0 PasswordSave) {
	var this1 *C.GMountOperation
	var save1 C.GPasswordSave
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	save1 = C.GPasswordSave(save0)
	C.g_mount_operation_set_password_save(this1, save1)
}
func (this0 *MountOperation) SetUsername(username0 string) {
	var this1 *C.GMountOperation
	var username1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	username1 = _GoStringToGString(username0)
	defer C.free(unsafe.Pointer(username1))
	C.g_mount_operation_set_username(this1, username1)
}
// blacklisted: MountOperationClass (struct)
// blacklisted: MountOperationPrivate (struct)
type MountOperationResult C.uint32_t
const (
	MountOperationResultHandled MountOperationResult = 0
	MountOperationResultAborted MountOperationResult = 1
	MountOperationResultUnhandled MountOperationResult = 2
)
type MountUnmountFlags C.uint32_t
const (
	MountUnmountFlagsNone MountUnmountFlags = 0
	MountUnmountFlagsForce MountUnmountFlags = 1
)
const NativeVolumeMonitorExtensionPointName = "gio-native-volume-monitor"
const NetworkMonitorExtensionPointName = "gio-network-monitor"
// blacklisted: NativeVolumeMonitor (object)
// blacklisted: NativeVolumeMonitorClass (struct)
// blacklisted: NetworkAddress (object)
// blacklisted: NetworkAddressClass (struct)
// blacklisted: NetworkAddressPrivate (struct)
// blacklisted: NetworkMonitor (interface)
// blacklisted: NetworkMonitorInterface (struct)
// blacklisted: NetworkService (object)
// blacklisted: NetworkServiceClass (struct)
// blacklisted: NetworkServicePrivate (struct)
// blacklisted: OutputStream (object)
// blacklisted: OutputStreamClass (struct)
// blacklisted: OutputStreamPrivate (struct)
type OutputStreamSpliceFlags C.uint32_t
const (
	OutputStreamSpliceFlagsNone OutputStreamSpliceFlags = 0
	OutputStreamSpliceFlagsCloseSource OutputStreamSpliceFlags = 1
	OutputStreamSpliceFlagsCloseTarget OutputStreamSpliceFlags = 2
)
// blacklisted: OutputVector (struct)
const ProxyExtensionPointName = "gio-proxy"
const ProxyResolverExtensionPointName = "gio-proxy-resolver"
type PasswordSave C.uint32_t
const (
	PasswordSaveNever PasswordSave = 0
	PasswordSaveForSession PasswordSave = 1
	PasswordSavePermanently PasswordSave = 2
)
type PermissionLike interface {
	gobject.ObjectLike
	InheritedFromGPermission() *C.GPermission
}

type Permission struct {
	gobject.Object
	
}

func ToPermission(objlike gobject.ObjectLike) *Permission {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Permission)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Permission)(obj)
	}
	panic("cannot cast to Permission")
}

func (this0 *Permission) InheritedFromGPermission() *C.GPermission {
	if this0 == nil {
		return nil
	}
	return (*C.GPermission)(this0.C)
}

func (this0 *Permission) GetStaticType() gobject.Type {
	return gobject.Type(C.g_permission_get_type())
}

func PermissionGetType() gobject.Type {
	return (*Permission)(nil).GetStaticType()
}
func (this0 *Permission) Acquire(cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GPermission
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.InheritedFromGPermission()
	}
	if cancellable0 != nil {
		cancellable1 = cancellable0.InheritedFromGCancellable()
	}
	ret1 := C.g_permission_acquire(this1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *Permission) AcquireAsync(cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GPermission
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.InheritedFromGPermission()
	}
	if cancellable0 != nil {
		cancellable1 = cancellable0.InheritedFromGCancellable()
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_permission_acquire_async(this1, cancellable1, callback1)
}
func (this0 *Permission) AcquireFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GPermission
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.InheritedFromGPermission()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_permission_acquire_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *Permission) GetAllowed() bool {
	var this1 *C.GPermission
	if this0 != nil {
		this1 = this0.InheritedFromGPermission()
	}
	ret1 := C.g_permission_get_allowed(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Permission) GetCanAcquire() bool {
	var this1 *C.GPermission
	if this0 != nil {
		this1 = this0.InheritedFromGPermission()
	}
	ret1 := C.g_permission_get_can_acquire(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Permission) GetCanRelease() bool {
	var this1 *C.GPermission
	if this0 != nil {
		this1 = this0.InheritedFromGPermission()
	}
	ret1 := C.g_permission_get_can_release(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Permission) ImplUpdate(allowed0 bool, can_acquire0 bool, can_release0 bool) {
	var this1 *C.GPermission
	var allowed1 C.int
	var can_acquire1 C.int
	var can_release1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGPermission()
	}
	allowed1 = _GoBoolToCBool(allowed0)
	can_acquire1 = _GoBoolToCBool(can_acquire0)
	can_release1 = _GoBoolToCBool(can_release0)
	C.g_permission_impl_update(this1, allowed1, can_acquire1, can_release1)
}
func (this0 *Permission) Release(cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GPermission
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.InheritedFromGPermission()
	}
	if cancellable0 != nil {
		cancellable1 = cancellable0.InheritedFromGCancellable()
	}
	ret1 := C.g_permission_release(this1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *Permission) ReleaseAsync(cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GPermission
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.InheritedFromGPermission()
	}
	if cancellable0 != nil {
		cancellable1 = cancellable0.InheritedFromGCancellable()
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_permission_release_async(this1, cancellable1, callback1)
}
func (this0 *Permission) ReleaseFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GPermission
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.InheritedFromGPermission()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_permission_release_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
// blacklisted: PermissionClass (struct)
// blacklisted: PermissionPrivate (struct)
// blacklisted: PollableInputStream (interface)
// blacklisted: PollableInputStreamInterface (struct)
// blacklisted: PollableOutputStream (interface)
// blacklisted: PollableOutputStreamInterface (struct)
// blacklisted: PollableSourceFunc (callback)
// blacklisted: Proxy (interface)
// blacklisted: ProxyAddress (object)
// blacklisted: ProxyAddressClass (struct)
// blacklisted: ProxyAddressEnumerator (object)
// blacklisted: ProxyAddressEnumeratorClass (struct)
// blacklisted: ProxyAddressEnumeratorPrivate (struct)
// blacklisted: ProxyAddressPrivate (struct)
// blacklisted: ProxyInterface (struct)
// blacklisted: ProxyResolver (interface)
// blacklisted: ProxyResolverInterface (struct)
// blacklisted: RemoteActionGroup (interface)
// blacklisted: RemoteActionGroupInterface (struct)
// blacklisted: Resolver (object)
// blacklisted: ResolverClass (struct)
type ResolverError C.uint32_t
const (
	ResolverErrorNotFound ResolverError = 0
	ResolverErrorTemporaryFailure ResolverError = 1
	ResolverErrorInternal ResolverError = 2
)
// blacklisted: ResolverPrivate (struct)
type ResolverRecordType C.uint32_t
const (
	ResolverRecordTypeSrv ResolverRecordType = 1
	ResolverRecordTypeMx ResolverRecordType = 2
	ResolverRecordTypeTxt ResolverRecordType = 3
	ResolverRecordTypeSoa ResolverRecordType = 4
	ResolverRecordTypeNs ResolverRecordType = 5
)
// blacklisted: Resource (struct)
type ResourceError C.uint32_t
const (
	ResourceErrorNotFound ResourceError = 0
	ResourceErrorInternal ResourceError = 1
)
type ResourceFlags C.uint32_t
const (
	ResourceFlagsNone ResourceFlags = 0
	ResourceFlagsCompressed ResourceFlags = 1
)
type ResourceLookupFlags C.uint32_t
const (
	ResourceLookupFlagsNone ResourceLookupFlags = 0
)
// blacklisted: Seekable (interface)
// blacklisted: SeekableIface (struct)
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
	var schema_id1 *C.char
	schema_id1 = _GoStringToGString(schema_id0)
	defer C.free(unsafe.Pointer(schema_id1))
	ret1 := C.g_settings_new(schema_id1)
	var ret2 *Settings
	ret2 = (*Settings)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
// blacklisted: Settings.new_full (method)
// blacklisted: Settings.new_with_backend (method)
// blacklisted: Settings.new_with_backend_and_path (method)
func NewSettingsWithPath(schema_id0 string, path0 string) *Settings {
	var schema_id1 *C.char
	var path1 *C.char
	schema_id1 = _GoStringToGString(schema_id0)
	defer C.free(unsafe.Pointer(schema_id1))
	path1 = _GoStringToGString(path0)
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
func SettingsUnbind(object0 unsafe.Pointer, property0 string) {
	var object1 unsafe.Pointer
	var property1 *C.char
	object1 = unsafe.Pointer(object0)
	property1 = _GoStringToGString(property0)
	defer C.free(unsafe.Pointer(property1))
	C.g_settings_unbind(object1, property1)
}
func (this0 *Settings) Apply() {
	var this1 *C.GSettings
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	C.g_settings_apply(this1)
}
func (this0 *Settings) Bind(key0 string, object0 gobject.ObjectLike, property0 string, flags0 SettingsBindFlags) {
	var this1 *C.GSettings
	var key1 *C.char
	var object1 *C.GObject
	var property1 *C.char
	var flags1 C.GSettingsBindFlags
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	if object0 != nil {
		object1 = object0.InheritedFromGObject()
	}
	property1 = _GoStringToGString(property0)
	defer C.free(unsafe.Pointer(property1))
	flags1 = C.GSettingsBindFlags(flags0)
	C.g_settings_bind(this1, key1, object1, property1, flags1)
}
func (this0 *Settings) BindWritable(key0 string, object0 gobject.ObjectLike, property0 string, inverted0 bool) {
	var this1 *C.GSettings
	var key1 *C.char
	var object1 *C.GObject
	var property1 *C.char
	var inverted1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	if object0 != nil {
		object1 = object0.InheritedFromGObject()
	}
	property1 = _GoStringToGString(property0)
	defer C.free(unsafe.Pointer(property1))
	inverted1 = _GoBoolToCBool(inverted0)
	C.g_settings_bind_writable(this1, key1, object1, property1, inverted1)
}
func (this0 *Settings) CreateAction(key0 string) *Action {
	var this1 *C.GSettings
	var key1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_settings_create_action(this1, key1)
	var ret2 *Action
	ret2 = (*Action)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
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
	var key1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_settings_get_boolean(this1, key1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Settings) GetChild(name0 string) *Settings {
	var this1 *C.GSettings
	var name1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	ret1 := C.g_settings_get_child(this1, name1)
	var ret2 *Settings
	ret2 = (*Settings)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Settings) GetDouble(key0 string) float64 {
	var this1 *C.GSettings
	var key1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_settings_get_double(this1, key1)
	var ret2 float64
	ret2 = float64(ret1)
	return ret2
}
func (this0 *Settings) GetEnum(key0 string) int {
	var this1 *C.GSettings
	var key1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_settings_get_enum(this1, key1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Settings) GetFlags(key0 string) int {
	var this1 *C.GSettings
	var key1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 = _GoStringToGString(key0)
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
	var key1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_settings_get_int(this1, key1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
// blacklisted: Settings.get_mapped (method)
func (this0 *Settings) GetRange(key0 string) *glib.Variant {
	var this1 *C.GSettings
	var key1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_settings_get_range(this1, key1)
	var ret2 *glib.Variant
	ret2 = (*glib.Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Settings) GetString(key0 string) string {
	var this1 *C.GSettings
	var key1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_settings_get_string(this1, key1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Settings) GetStrv(key0 string) []string {
	var this1 *C.GSettings
	var key1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_settings_get_strv(this1, key1)
	var ret2 []string
	ret2 = make([]string, uint(C._array_length(unsafe.Pointer(ret1))))
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return ret2
}
func (this0 *Settings) GetUint(key0 string) int {
	var this1 *C.GSettings
	var key1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_settings_get_uint(this1, key1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Settings) GetValue(key0 string) *glib.Variant {
	var this1 *C.GSettings
	var key1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_settings_get_value(this1, key1)
	var ret2 *glib.Variant
	ret2 = (*glib.Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Settings) IsWritable(name0 string) bool {
	var this1 *C.GSettings
	var name1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	name1 = _GoStringToGString(name0)
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
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
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
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return ret2
}
func (this0 *Settings) RangeCheck(key0 string, value0 *glib.Variant) bool {
	var this1 *C.GSettings
	var key1 *C.char
	var value1 *C.GVariant
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = (*C.GVariant)(unsafe.Pointer(value0))
	ret1 := C.g_settings_range_check(this1, key1, value1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Settings) Reset(key0 string) {
	var this1 *C.GSettings
	var key1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 = _GoStringToGString(key0)
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
	var key1 *C.char
	var value1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = _GoBoolToCBool(value0)
	ret1 := C.g_settings_set_boolean(this1, key1, value1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Settings) SetDouble(key0 string, value0 float64) bool {
	var this1 *C.GSettings
	var key1 *C.char
	var value1 C.double
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = C.double(value0)
	ret1 := C.g_settings_set_double(this1, key1, value1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Settings) SetEnum(key0 string, value0 int) bool {
	var this1 *C.GSettings
	var key1 *C.char
	var value1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = C.int32_t(value0)
	ret1 := C.g_settings_set_enum(this1, key1, value1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Settings) SetFlags(key0 string, value0 int) bool {
	var this1 *C.GSettings
	var key1 *C.char
	var value1 C.uint32_t
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = C.uint32_t(value0)
	ret1 := C.g_settings_set_flags(this1, key1, value1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Settings) SetInt(key0 string, value0 int) bool {
	var this1 *C.GSettings
	var key1 *C.char
	var value1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = C.int32_t(value0)
	ret1 := C.g_settings_set_int(this1, key1, value1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Settings) SetString(key0 string, value0 string) bool {
	var this1 *C.GSettings
	var key1 *C.char
	var value1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = _GoStringToGString(value0)
	defer C.free(unsafe.Pointer(value1))
	ret1 := C.g_settings_set_string(this1, key1, value1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Settings) SetStrv(key0 string, value0 []string) bool {
	var this1 *C.GSettings
	var key1 *C.char
	var value1 **C.char
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = (**C.char)(C.malloc(C.size_t(int(unsafe.Sizeof(*value1)) * (len(value0) + 1))))
	defer C.free(unsafe.Pointer(value1))
	for i, e := range value0 {
		(*(*[999999]*C.char)(unsafe.Pointer(value1)))[i] = _GoStringToGString(e)
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
	var key1 *C.char
	var value1 C.uint32_t
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = C.uint32_t(value0)
	ret1 := C.g_settings_set_uint(this1, key1, value1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Settings) SetValue(key0 string, value0 *glib.Variant) bool {
	var this1 *C.GSettings
	var key1 *C.char
	var value1 *C.GVariant
	if this0 != nil {
		this1 = this0.InheritedFromGSettings()
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = (*C.GVariant)(unsafe.Pointer(value0))
	ret1 := C.g_settings_set_value(this1, key1, value1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
// blacklisted: SettingsBackend (struct)
type SettingsBindFlags C.uint32_t
const (
	SettingsBindFlagsDefault SettingsBindFlags = 0
	SettingsBindFlagsGet SettingsBindFlags = 1
	SettingsBindFlagsSet SettingsBindFlags = 2
	SettingsBindFlagsNoSensitivity SettingsBindFlags = 4
	SettingsBindFlagsGetNoChanges SettingsBindFlags = 8
	SettingsBindFlagsInvertBoolean SettingsBindFlags = 16
)
// blacklisted: SettingsBindGetMapping (callback)
// blacklisted: SettingsBindSetMapping (callback)
// blacklisted: SettingsClass (struct)
// blacklisted: SettingsGetMapping (callback)
// blacklisted: SettingsPrivate (struct)
// blacklisted: SettingsSchema (struct)
// blacklisted: SettingsSchemaSource (struct)
// blacklisted: SimpleAction (object)
// blacklisted: SimpleActionGroup (object)
// blacklisted: SimpleActionGroupClass (struct)
// blacklisted: SimpleActionGroupPrivate (struct)
// blacklisted: SimpleAsyncResult (object)
// blacklisted: SimpleAsyncResultClass (struct)
// blacklisted: SimpleAsyncThreadFunc (callback)
// blacklisted: SimplePermission (object)
// blacklisted: SimpleProxyResolver (object)
// blacklisted: SimpleProxyResolverClass (struct)
// blacklisted: SimpleProxyResolverPrivate (struct)
// blacklisted: Socket (object)
// blacklisted: SocketAddress (object)
// blacklisted: SocketAddressClass (struct)
// blacklisted: SocketAddressEnumerator (object)
// blacklisted: SocketAddressEnumeratorClass (struct)
// blacklisted: SocketClass (struct)
// blacklisted: SocketClient (object)
// blacklisted: SocketClientClass (struct)
type SocketClientEvent C.uint32_t
const (
	SocketClientEventResolving SocketClientEvent = 0
	SocketClientEventResolved SocketClientEvent = 1
	SocketClientEventConnecting SocketClientEvent = 2
	SocketClientEventConnected SocketClientEvent = 3
	SocketClientEventProxyNegotiating SocketClientEvent = 4
	SocketClientEventProxyNegotiated SocketClientEvent = 5
	SocketClientEventTlsHandshaking SocketClientEvent = 6
	SocketClientEventTlsHandshaked SocketClientEvent = 7
	SocketClientEventComplete SocketClientEvent = 8
)
// blacklisted: SocketClientPrivate (struct)
// blacklisted: SocketConnectable (interface)
// blacklisted: SocketConnectableIface (struct)
// blacklisted: SocketConnection (object)
// blacklisted: SocketConnectionClass (struct)
// blacklisted: SocketConnectionPrivate (struct)
// blacklisted: SocketControlMessage (object)
// blacklisted: SocketControlMessageClass (struct)
// blacklisted: SocketControlMessagePrivate (struct)
type SocketFamily C.uint32_t
const (
	SocketFamilyInvalid SocketFamily = 0
	SocketFamilyUnix SocketFamily = 1
	SocketFamilyIpv4 SocketFamily = 2
	SocketFamilyIpv6 SocketFamily = 10
)
// blacklisted: SocketListener (object)
// blacklisted: SocketListenerClass (struct)
// blacklisted: SocketListenerPrivate (struct)
type SocketMsgFlags C.uint32_t
const (
	SocketMsgFlagsNone SocketMsgFlags = 0
	SocketMsgFlagsOob SocketMsgFlags = 1
	SocketMsgFlagsPeek SocketMsgFlags = 2
	SocketMsgFlagsDontroute SocketMsgFlags = 4
)
// blacklisted: SocketPrivate (struct)
type SocketProtocol C.int32_t
const (
	SocketProtocolUnknown SocketProtocol = -1
	SocketProtocolDefault SocketProtocol = 0
	SocketProtocolTcp SocketProtocol = 6
	SocketProtocolUdp SocketProtocol = 17
	SocketProtocolSctp SocketProtocol = 132
)
// blacklisted: SocketService (object)
// blacklisted: SocketServiceClass (struct)
// blacklisted: SocketServicePrivate (struct)
// blacklisted: SocketSourceFunc (callback)
type SocketType C.uint32_t
const (
	SocketTypeInvalid SocketType = 0
	SocketTypeStream SocketType = 1
	SocketTypeDatagram SocketType = 2
	SocketTypeSeqpacket SocketType = 3
)
// blacklisted: SrvTarget (struct)
// blacklisted: StaticResource (struct)
const TlsBackendExtensionPointName = "gio-tls-backend"
const TlsDatabasePurposeAuthenticateClient = "1.3.6.1.5.5.7.3.2"
const TlsDatabasePurposeAuthenticateServer = "1.3.6.1.5.5.7.3.1"
// blacklisted: Task (object)
// blacklisted: TaskClass (struct)
// blacklisted: TaskThreadFunc (callback)
// blacklisted: TcpConnection (object)
// blacklisted: TcpConnectionClass (struct)
// blacklisted: TcpConnectionPrivate (struct)
// blacklisted: TcpWrapperConnection (object)
// blacklisted: TcpWrapperConnectionClass (struct)
// blacklisted: TcpWrapperConnectionPrivate (struct)
// blacklisted: TestDBus (object)
type TestDBusFlags C.uint32_t
const (
	TestDBusFlagsNone TestDBusFlags = 0
)
// blacklisted: ThemedIcon (object)
// blacklisted: ThemedIconClass (struct)
// blacklisted: ThreadedSocketService (object)
// blacklisted: ThreadedSocketServiceClass (struct)
// blacklisted: ThreadedSocketServicePrivate (struct)
type TlsAuthenticationMode C.uint32_t
const (
	TlsAuthenticationModeNone TlsAuthenticationMode = 0
	TlsAuthenticationModeRequested TlsAuthenticationMode = 1
	TlsAuthenticationModeRequired TlsAuthenticationMode = 2
)
// blacklisted: TlsBackend (interface)
// blacklisted: TlsBackendInterface (struct)
// blacklisted: TlsCertificate (object)
// blacklisted: TlsCertificateClass (struct)
type TlsCertificateFlags C.uint32_t
const (
	TlsCertificateFlagsUnknownCa TlsCertificateFlags = 1
	TlsCertificateFlagsBadIdentity TlsCertificateFlags = 2
	TlsCertificateFlagsNotActivated TlsCertificateFlags = 4
	TlsCertificateFlagsExpired TlsCertificateFlags = 8
	TlsCertificateFlagsRevoked TlsCertificateFlags = 16
	TlsCertificateFlagsInsecure TlsCertificateFlags = 32
	TlsCertificateFlagsGenericError TlsCertificateFlags = 64
	TlsCertificateFlagsValidateAll TlsCertificateFlags = 127
)
// blacklisted: TlsCertificatePrivate (struct)
// blacklisted: TlsClientConnection (interface)
// blacklisted: TlsClientConnectionInterface (struct)
// blacklisted: TlsConnection (object)
// blacklisted: TlsConnectionClass (struct)
// blacklisted: TlsConnectionPrivate (struct)
// blacklisted: TlsDatabase (object)
// blacklisted: TlsDatabaseClass (struct)
type TlsDatabaseLookupFlags C.uint32_t
const (
	TlsDatabaseLookupFlagsNone TlsDatabaseLookupFlags = 0
	TlsDatabaseLookupFlagsKeypair TlsDatabaseLookupFlags = 1
)
// blacklisted: TlsDatabasePrivate (struct)
type TlsDatabaseVerifyFlags C.uint32_t
const (
	TlsDatabaseVerifyFlagsNone TlsDatabaseVerifyFlags = 0
)
type TlsError C.uint32_t
const (
	TlsErrorUnavailable TlsError = 0
	TlsErrorMisc TlsError = 1
	TlsErrorBadCertificate TlsError = 2
	TlsErrorNotTls TlsError = 3
	TlsErrorHandshake TlsError = 4
	TlsErrorCertificateRequired TlsError = 5
	TlsErrorEof TlsError = 6
)
// blacklisted: TlsFileDatabase (interface)
// blacklisted: TlsFileDatabaseInterface (struct)
// blacklisted: TlsInteraction (object)
// blacklisted: TlsInteractionClass (struct)
// blacklisted: TlsInteractionPrivate (struct)
type TlsInteractionResult C.uint32_t
const (
	TlsInteractionResultUnhandled TlsInteractionResult = 0
	TlsInteractionResultHandled TlsInteractionResult = 1
	TlsInteractionResultFailed TlsInteractionResult = 2
)
// blacklisted: TlsPassword (object)
// blacklisted: TlsPasswordClass (struct)
type TlsPasswordFlags C.uint32_t
const (
	TlsPasswordFlagsNone TlsPasswordFlags = 0
	TlsPasswordFlagsRetry TlsPasswordFlags = 2
	TlsPasswordFlagsManyTries TlsPasswordFlags = 4
	TlsPasswordFlagsFinalTry TlsPasswordFlags = 8
)
// blacklisted: TlsPasswordPrivate (struct)
type TlsRehandshakeMode C.uint32_t
const (
	TlsRehandshakeModeNever TlsRehandshakeMode = 0
	TlsRehandshakeModeSafely TlsRehandshakeMode = 1
	TlsRehandshakeModeUnsafely TlsRehandshakeMode = 2
)
// blacklisted: TlsServerConnection (interface)
// blacklisted: TlsServerConnectionInterface (struct)
// blacklisted: UnixConnection (object)
// blacklisted: UnixConnectionClass (struct)
// blacklisted: UnixConnectionPrivate (struct)
// blacklisted: UnixCredentialsMessage (object)
// blacklisted: UnixCredentialsMessageClass (struct)
// blacklisted: UnixCredentialsMessagePrivate (struct)
// blacklisted: UnixFDList (object)
// blacklisted: UnixFDListClass (struct)
// blacklisted: UnixFDListPrivate (struct)
// blacklisted: UnixFDMessage (object)
// blacklisted: UnixFDMessageClass (struct)
// blacklisted: UnixFDMessagePrivate (struct)
// blacklisted: UnixInputStream (object)
// blacklisted: UnixInputStreamClass (struct)
// blacklisted: UnixInputStreamPrivate (struct)
// blacklisted: UnixMountEntry (struct)
// blacklisted: UnixMountMonitor (object)
// blacklisted: UnixMountMonitorClass (struct)
// blacklisted: UnixMountPoint (struct)
// blacklisted: UnixOutputStream (object)
// blacklisted: UnixOutputStreamClass (struct)
// blacklisted: UnixOutputStreamPrivate (struct)
// blacklisted: UnixSocketAddress (object)
// blacklisted: UnixSocketAddressClass (struct)
// blacklisted: UnixSocketAddressPrivate (struct)
type UnixSocketAddressType C.uint32_t
const (
	UnixSocketAddressTypeInvalid UnixSocketAddressType = 0
	UnixSocketAddressTypeAnonymous UnixSocketAddressType = 1
	UnixSocketAddressTypePath UnixSocketAddressType = 2
	UnixSocketAddressTypeAbstract UnixSocketAddressType = 3
	UnixSocketAddressTypeAbstractPadded UnixSocketAddressType = 4
)
const VfsExtensionPointName = "gio-vfs"
const VolumeIdentifierKindClass = "class"
const VolumeIdentifierKindHalUdi = "hal-udi"
const VolumeIdentifierKindLabel = "label"
const VolumeIdentifierKindNfsMount = "nfs-mount"
const VolumeIdentifierKindUnixDevice = "unix-device"
const VolumeIdentifierKindUuid = "uuid"
const VolumeMonitorExtensionPointName = "gio-volume-monitor"
// blacklisted: Vfs (object)
// blacklisted: VfsClass (struct)
// blacklisted: Volume (interface)
// blacklisted: VolumeIface (struct)
// blacklisted: VolumeMonitor (object)
// blacklisted: VolumeMonitorClass (struct)
// blacklisted: ZlibCompressor (object)
// blacklisted: ZlibCompressorClass (struct)
type ZlibCompressorFormat C.uint32_t
const (
	ZlibCompressorFormatZlib ZlibCompressorFormat = 0
	ZlibCompressorFormatGzip ZlibCompressorFormat = 1
	ZlibCompressorFormatRaw ZlibCompressorFormat = 2
)
// blacklisted: ZlibDecompressor (object)
// blacklisted: ZlibDecompressorClass (struct)
// blacklisted: app_info_create_from_commandline (function)
// blacklisted: app_info_get_all (function)
// blacklisted: app_info_get_all_for_type (function)
// blacklisted: app_info_get_default_for_type (function)
// blacklisted: app_info_get_default_for_uri_scheme (function)
// blacklisted: app_info_get_fallback_for_type (function)
// blacklisted: app_info_get_recommended_for_type (function)
// blacklisted: app_info_launch_default_for_uri (function)
// blacklisted: app_info_reset_type_associations (function)
// blacklisted: async_initable_newv_async (function)
// blacklisted: bus_get (function)
// blacklisted: bus_get_finish (function)
// blacklisted: bus_get_sync (function)
// blacklisted: bus_own_name_on_connection (function)
// blacklisted: bus_own_name (function)
// blacklisted: bus_unown_name (function)
// blacklisted: bus_unwatch_name (function)
// blacklisted: bus_watch_name_on_connection (function)
// blacklisted: bus_watch_name (function)
// blacklisted: content_type_can_be_executable (function)
// blacklisted: content_type_equals (function)
// blacklisted: content_type_from_mime_type (function)
// blacklisted: content_type_get_description (function)
// blacklisted: content_type_get_generic_icon_name (function)
// blacklisted: content_type_get_icon (function)
// blacklisted: content_type_get_mime_type (function)
// blacklisted: content_type_get_symbolic_icon (function)
// blacklisted: content_type_guess (function)
// blacklisted: content_type_guess_for_tree (function)
// blacklisted: content_type_is_a (function)
// blacklisted: content_type_is_unknown (function)
// blacklisted: content_types_get_registered (function)
// blacklisted: dbus_address_escape_value (function)
// blacklisted: dbus_address_get_for_bus_sync (function)
// blacklisted: dbus_address_get_stream (function)
// blacklisted: dbus_address_get_stream_finish (function)
// blacklisted: dbus_address_get_stream_sync (function)
// blacklisted: dbus_annotation_info_lookup (function)
// blacklisted: dbus_error_encode_gerror (function)
// blacklisted: dbus_error_get_remote_error (function)
// blacklisted: dbus_error_is_remote_error (function)
// blacklisted: dbus_error_new_for_dbus_error (function)
// blacklisted: dbus_error_quark (function)
// blacklisted: dbus_error_register_error (function)
// blacklisted: dbus_error_register_error_domain (function)
// blacklisted: dbus_error_strip_remote_error (function)
// blacklisted: dbus_error_unregister_error (function)
// blacklisted: dbus_generate_guid (function)
// blacklisted: dbus_gvalue_to_gvariant (function)
// blacklisted: dbus_gvariant_to_gvalue (function)
// blacklisted: dbus_is_address (function)
// blacklisted: dbus_is_guid (function)
// blacklisted: dbus_is_interface_name (function)
// blacklisted: dbus_is_member_name (function)
// blacklisted: dbus_is_name (function)
// blacklisted: dbus_is_supported_address (function)
// blacklisted: dbus_is_unique_name (function)
// blacklisted: file_new_for_commandline_arg (function)
// blacklisted: file_new_for_commandline_arg_and_cwd (function)
// blacklisted: file_new_for_path (function)
// blacklisted: file_new_for_uri (function)
// blacklisted: file_new_tmp (function)
// blacklisted: file_parse_name (function)
// blacklisted: icon_hash (function)
// blacklisted: icon_new_for_string (function)
// blacklisted: initable_newv (function)
// blacklisted: io_error_from_errno (function)
// blacklisted: io_error_quark (function)
// blacklisted: io_extension_point_implement (function)
// blacklisted: io_extension_point_lookup (function)
// blacklisted: io_extension_point_register (function)
// blacklisted: io_modules_load_all_in_directory (function)
// blacklisted: io_modules_load_all_in_directory_with_scope (function)
// blacklisted: io_modules_scan_all_in_directory (function)
// blacklisted: io_modules_scan_all_in_directory_with_scope (function)
// blacklisted: io_scheduler_cancel_all_jobs (function)
// blacklisted: io_scheduler_push_job (function)
// blacklisted: network_monitor_get_default (function)
// blacklisted: networking_init (function)
// blacklisted: pollable_source_new (function)
// blacklisted: pollable_source_new_full (function)
// blacklisted: pollable_stream_read (function)
// blacklisted: pollable_stream_write (function)
// blacklisted: pollable_stream_write_all (function)
// blacklisted: proxy_get_default_for_protocol (function)
// blacklisted: proxy_resolver_get_default (function)
// blacklisted: resolver_error_quark (function)
// blacklisted: resource_error_quark (function)
// blacklisted: resource_load (function)
// blacklisted: resources_enumerate_children (function)
// blacklisted: resources_get_info (function)
// blacklisted: resources_lookup_data (function)
// blacklisted: resources_open_stream (function)
// blacklisted: settings_schema_source_get_default (function)
// blacklisted: simple_async_report_gerror_in_idle (function)
// blacklisted: tls_backend_get_default (function)
// blacklisted: tls_client_connection_new (function)
// blacklisted: tls_error_quark (function)
// blacklisted: tls_file_database_new (function)
// blacklisted: tls_server_connection_new (function)
// blacklisted: unix_is_mount_path_system_internal (function)
// blacklisted: unix_mount_compare (function)
// blacklisted: unix_mount_free (function)
// blacklisted: unix_mount_get_device_path (function)
// blacklisted: unix_mount_get_fs_type (function)
// blacklisted: unix_mount_get_mount_path (function)
// blacklisted: unix_mount_guess_can_eject (function)
// blacklisted: unix_mount_guess_icon (function)
// blacklisted: unix_mount_guess_name (function)
// blacklisted: unix_mount_guess_should_display (function)
// blacklisted: unix_mount_guess_symbolic_icon (function)
// blacklisted: unix_mount_is_readonly (function)
// blacklisted: unix_mount_is_system_internal (function)
// blacklisted: unix_mount_points_changed_since (function)
// blacklisted: unix_mounts_changed_since (function)
