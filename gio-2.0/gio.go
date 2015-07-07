package gio

/*
#include "gio.gen.h"


GList* g_list_append(GList*, void*);
void g_list_free(GList*);

extern GObject *g_object_ref_sink(GObject*);
extern void g_object_unref(GObject*);
extern void g_error_free(GError*);
extern void g_free(void*);
#cgo pkg-config: gio-2.0
*/
import "C"
import "unsafe"

import (
	"pkg.deepin.io/lib/gobject-2.0"
	"pkg.deepin.io/lib/glib-2.0"
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

func (*Action) GetStaticType() gobject.Type {
	return gobject.Type(C.g_action_get_type())
}


type ActionImpl struct {}

func ToAction(objlike gobject.ObjectLike) *Action {
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), gobject.Type(C.g_action_get_type()))
	if obj != nil {
		return (*Action)(obj)
	}
	panic("cannot cast to Action")
}

func (this0 *ActionImpl) ImplementsGAction() *C.GAction {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GAction)((*gobject.Object)(unsafe.Pointer(obj)).C)
}
func ActionNameIsValid(action_name0 string) bool {
	var action_name1 *C.char
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	ret1 := C.g_action_name_is_valid(action_name1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func ActionParseDetailedName(detailed_name0 string) (string, *glib.Variant, bool, error) {
	var detailed_name1 *C.char
	var action_name1 *C.char
	var target_value1 *C.GVariant
	var err1 *C.GError
	detailed_name1 = _GoStringToGString(detailed_name0)
	defer C.free(unsafe.Pointer(detailed_name1))
	ret1 := C.g_action_parse_detailed_name(detailed_name1, &action_name1, &target_value1, &err1)
	var action_name2 string
	var target_value2 *glib.Variant
	var ret2 bool
	var err2 error
	action_name2 = C.GoString(action_name1)
	C.g_free(unsafe.Pointer(action_name1))
	target_value2 = (*glib.Variant)(unsafe.Pointer(target_value1))
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return action_name2, target_value2, ret2, err2
}
func ActionPrintDetailedName(action_name0 string, target_value0 *glib.Variant) string {
	var action_name1 *C.char
	var target_value1 *C.GVariant
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	target_value1 = (*C.GVariant)(unsafe.Pointer(target_value0))
	ret1 := C.g_action_print_detailed_name(action_name1, target_value1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
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

func (*ActionGroup) GetStaticType() gobject.Type {
	return gobject.Type(C.g_action_group_get_type())
}


type ActionGroupImpl struct {}

func ToActionGroup(objlike gobject.ObjectLike) *ActionGroup {
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), gobject.Type(C.g_action_group_get_type()))
	if obj != nil {
		return (*ActionGroup)(obj)
	}
	panic("cannot cast to ActionGroup")
}

func (this0 *ActionGroupImpl) ImplementsGActionGroup() *C.GActionGroup {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GActionGroup)((*gobject.Object)(unsafe.Pointer(obj)).C)
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

func (*ActionMap) GetStaticType() gobject.Type {
	return gobject.Type(C.g_action_map_get_type())
}


type ActionMapImpl struct {}

func ToActionMap(objlike gobject.ObjectLike) *ActionMap {
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), gobject.Type(C.g_action_map_get_type()))
	if obj != nil {
		return (*ActionMap)(obj)
	}
	panic("cannot cast to ActionMap")
}

func (this0 *ActionMapImpl) ImplementsGActionMap() *C.GActionMap {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GActionMap)((*gobject.Object)(unsafe.Pointer(obj)).C)
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

func (*AppInfo) GetStaticType() gobject.Type {
	return gobject.Type(C.g_app_info_get_type())
}


type AppInfoImpl struct {}

func ToAppInfo(objlike gobject.ObjectLike) *AppInfo {
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), gobject.Type(C.g_app_info_get_type()))
	if obj != nil {
		return (*AppInfo)(obj)
	}
	panic("cannot cast to AppInfo")
}

func (this0 *AppInfoImpl) ImplementsGAppInfo() *C.GAppInfo {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GAppInfo)((*gobject.Object)(unsafe.Pointer(obj)).C)
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
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
// blacklisted: AppInfo.get_all (method)
// blacklisted: AppInfo.get_all_for_type (method)
// blacklisted: AppInfo.get_default_for_type (method)
// blacklisted: AppInfo.get_default_for_uri_scheme (method)
// blacklisted: AppInfo.get_fallback_for_type (method)
// blacklisted: AppInfo.get_recommended_for_type (method)
// blacklisted: AppInfo.launch_default_for_uri (method)
// blacklisted: AppInfo.reset_type_associations (method)
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
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
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
	for _, e := range files0 {
		var s *C.GFile
		if e != nil {
			s = e.ImplementsGFile()
		}
		files1 = C.g_list_append(files1, unsafe.Pointer(s))
	}
	defer C.g_list_free(files1)
	if launch_context0 != nil {
		launch_context1 = (*C.GAppLaunchContext)(launch_context0.InheritedFromGAppLaunchContext())
	}
	ret1 := C.g_app_info_launch(this1, files1, launch_context1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
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
	for _, e := range uris0 {
		var s *C.char
		s = _GoStringToGString(e)
		defer C.free(unsafe.Pointer(s))
		uris1 = C.g_list_append(uris1, unsafe.Pointer(s))
	}
	defer C.g_list_free(uris1)
	if launch_context0 != nil {
		launch_context1 = (*C.GAppLaunchContext)(launch_context0.InheritedFromGAppLaunchContext())
	}
	ret1 := C.g_app_info_launch_uris(this1, uris1, launch_context1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
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
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
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
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
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
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
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
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
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
// blacklisted: AppInfoMonitor (object)
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
		this1 = (*C.GAppLaunchContext)(this0.InheritedFromGAppLaunchContext())
	}
	if info0 != nil {
		info1 = info0.ImplementsGAppInfo()
	}
	for _, e := range files0 {
		var s *C.GFile
		if e != nil {
			s = e.ImplementsGFile()
		}
		files1 = C.g_list_append(files1, unsafe.Pointer(s))
	}
	defer C.g_list_free(files1)
	ret1 := C.g_app_launch_context_get_display(this1, info1, files1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *AppLaunchContext) GetEnvironment() []string {
	var this1 *C.GAppLaunchContext
	if this0 != nil {
		this1 = (*C.GAppLaunchContext)(this0.InheritedFromGAppLaunchContext())
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
		this1 = (*C.GAppLaunchContext)(this0.InheritedFromGAppLaunchContext())
	}
	if info0 != nil {
		info1 = info0.ImplementsGAppInfo()
	}
	for _, e := range files0 {
		var s *C.GFile
		if e != nil {
			s = e.ImplementsGFile()
		}
		files1 = C.g_list_append(files1, unsafe.Pointer(s))
	}
	defer C.g_list_free(files1)
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
		this1 = (*C.GAppLaunchContext)(this0.InheritedFromGAppLaunchContext())
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
		this1 = (*C.GAppLaunchContext)(this0.InheritedFromGAppLaunchContext())
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
		this1 = (*C.GAppLaunchContext)(this0.InheritedFromGAppLaunchContext())
	}
	variable1 = _GoStringToGString(variable0)
	defer C.free(unsafe.Pointer(variable1))
	C.g_app_launch_context_unsetenv(this1, variable1)
}
// blacklisted: AppLaunchContextClass (struct)
// blacklisted: AppLaunchContextPrivate (struct)
// blacklisted: Application (object)
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

func (*AsyncResult) GetStaticType() gobject.Type {
	return gobject.Type(C.g_async_result_get_type())
}


type AsyncResultImpl struct {}

func ToAsyncResult(objlike gobject.ObjectLike) *AsyncResult {
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), gobject.Type(C.g_async_result_get_type()))
	if obj != nil {
		return (*AsyncResult)(obj)
	}
	panic("cannot cast to AsyncResult")
}

func (this0 *AsyncResultImpl) ImplementsGAsyncResult() *C.GAsyncResult {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GAsyncResult)((*gobject.Object)(unsafe.Pointer(obj)).C)
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
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
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
// blacklisted: BytesIcon (object)
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
		this1 = (*C.GCancellable)(this0.InheritedFromGCancellable())
	}
	C.g_cancellable_cancel(this1)
}
// blacklisted: Cancellable.connect (method)
func (this0 *Cancellable) Disconnect(handler_id0 uint64) {
	var this1 *C.GCancellable
	var handler_id1 C.uint64_t
	if this0 != nil {
		this1 = (*C.GCancellable)(this0.InheritedFromGCancellable())
	}
	handler_id1 = C.uint64_t(handler_id0)
	C.g_cancellable_disconnect(this1, handler_id1)
}
func (this0 *Cancellable) GetFd() int32 {
	var this1 *C.GCancellable
	if this0 != nil {
		this1 = (*C.GCancellable)(this0.InheritedFromGCancellable())
	}
	ret1 := C.g_cancellable_get_fd(this1)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func (this0 *Cancellable) IsCancelled() bool {
	var this1 *C.GCancellable
	if this0 != nil {
		this1 = (*C.GCancellable)(this0.InheritedFromGCancellable())
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
		this1 = (*C.GCancellable)(this0.InheritedFromGCancellable())
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
		this1 = (*C.GCancellable)(this0.InheritedFromGCancellable())
	}
	C.g_cancellable_pop_current(this1)
}
func (this0 *Cancellable) PushCurrent() {
	var this1 *C.GCancellable
	if this0 != nil {
		this1 = (*C.GCancellable)(this0.InheritedFromGCancellable())
	}
	C.g_cancellable_push_current(this1)
}
func (this0 *Cancellable) ReleaseFd() {
	var this1 *C.GCancellable
	if this0 != nil {
		this1 = (*C.GCancellable)(this0.InheritedFromGCancellable())
	}
	C.g_cancellable_release_fd(this1)
}
func (this0 *Cancellable) Reset() {
	var this1 *C.GCancellable
	if this0 != nil {
		this1 = (*C.GCancellable)(this0.InheritedFromGCancellable())
	}
	C.g_cancellable_reset(this1)
}
func (this0 *Cancellable) SetErrorIfCancelled() (bool, error) {
	var this1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GCancellable)(this0.InheritedFromGCancellable())
	}
	ret1 := C.g_cancellable_set_error_if_cancelled(this1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
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
	CredentialsTypeSolarisUcred CredentialsType = 4
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
	DBusProxyFlagsDoNotAutoStartAtConstruction DBusProxyFlags = 16
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
	DBusSignalFlagsMatchArg0Namespace DBusSignalFlags = 2
	DBusSignalFlagsMatchArg0Path DBusSignalFlags = 4
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
type DesktopAppInfoLike interface {
	gobject.ObjectLike
	InheritedFromGDesktopAppInfo() *C.GDesktopAppInfo
}

type DesktopAppInfo struct {
	gobject.Object
	AppInfoImpl
}

func ToDesktopAppInfo(objlike gobject.ObjectLike) *DesktopAppInfo {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*DesktopAppInfo)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*DesktopAppInfo)(obj)
	}
	panic("cannot cast to DesktopAppInfo")
}

func (this0 *DesktopAppInfo) InheritedFromGDesktopAppInfo() *C.GDesktopAppInfo {
	if this0 == nil {
		return nil
	}
	return (*C.GDesktopAppInfo)(this0.C)
}

func (this0 *DesktopAppInfo) GetStaticType() gobject.Type {
	return gobject.Type(C.g_desktop_app_info_get_type())
}

func DesktopAppInfoGetType() gobject.Type {
	return (*DesktopAppInfo)(nil).GetStaticType()
}
func NewDesktopAppInfo(desktop_id0 string) *DesktopAppInfo {
	var desktop_id1 *C.char
	desktop_id1 = _GoStringToGString(desktop_id0)
	defer C.free(unsafe.Pointer(desktop_id1))
	ret1 := C.g_desktop_app_info_new(desktop_id1)
	var ret2 *DesktopAppInfo
	ret2 = (*DesktopAppInfo)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func NewDesktopAppInfoFromFilename(filename0 string) *DesktopAppInfo {
	var filename1 *C.char
	filename1 = _GoStringToGString(filename0)
	defer C.free(unsafe.Pointer(filename1))
	ret1 := C.g_desktop_app_info_new_from_filename(filename1)
	var ret2 *DesktopAppInfo
	ret2 = (*DesktopAppInfo)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func NewDesktopAppInfoFromKeyfile(key_file0 *glib.KeyFile) *DesktopAppInfo {
	var key_file1 *C.GKeyFile
	key_file1 = (*C.GKeyFile)(unsafe.Pointer(key_file0))
	ret1 := C.g_desktop_app_info_new_from_keyfile(key_file1)
	var ret2 *DesktopAppInfo
	ret2 = (*DesktopAppInfo)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func DesktopAppInfoSearch(search_string0 string) [][]string {
	var search_string1 *C.char
	search_string1 = _GoStringToGString(search_string0)
	defer C.free(unsafe.Pointer(search_string1))
	ret1 := C.g_desktop_app_info_search(search_string1)
	var ret2 [][]string
	ret2 = make([][]string, uint(C._array_length(unsafe.Pointer(ret1))))
	for i := range ret2 {
		for i := range ret2[i] {
			ret2[i][i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer((*(*[999999]**C.char)(unsafe.Pointer(ret1)))[i])))[i])
			C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer((*(*[999999]**C.char)(unsafe.Pointer(ret1)))[i])))[i]))
		}
	}
	return ret2
}
func DesktopAppInfoSetDesktopEnv(desktop_env0 string) {
	var desktop_env1 *C.char
	desktop_env1 = _GoStringToGString(desktop_env0)
	defer C.free(unsafe.Pointer(desktop_env1))
	C.g_desktop_app_info_set_desktop_env(desktop_env1)
}
func (this0 *DesktopAppInfo) GetActionName(action_name0 string) string {
	var this1 *C.GDesktopAppInfo
	var action_name1 *C.char
	if this0 != nil {
		this1 = (*C.GDesktopAppInfo)(this0.InheritedFromGDesktopAppInfo())
	}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	ret1 := C.g_desktop_app_info_get_action_name(this1, action_name1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *DesktopAppInfo) GetBoolean(key0 string) bool {
	var this1 *C.GDesktopAppInfo
	var key1 *C.char
	if this0 != nil {
		this1 = (*C.GDesktopAppInfo)(this0.InheritedFromGDesktopAppInfo())
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_desktop_app_info_get_boolean(this1, key1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *DesktopAppInfo) GetCategories() string {
	var this1 *C.GDesktopAppInfo
	if this0 != nil {
		this1 = (*C.GDesktopAppInfo)(this0.InheritedFromGDesktopAppInfo())
	}
	ret1 := C.g_desktop_app_info_get_categories(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *DesktopAppInfo) GetFilename() string {
	var this1 *C.GDesktopAppInfo
	if this0 != nil {
		this1 = (*C.GDesktopAppInfo)(this0.InheritedFromGDesktopAppInfo())
	}
	ret1 := C.g_desktop_app_info_get_filename(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *DesktopAppInfo) GetGenericName() string {
	var this1 *C.GDesktopAppInfo
	if this0 != nil {
		this1 = (*C.GDesktopAppInfo)(this0.InheritedFromGDesktopAppInfo())
	}
	ret1 := C.g_desktop_app_info_get_generic_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *DesktopAppInfo) GetIsHidden() bool {
	var this1 *C.GDesktopAppInfo
	if this0 != nil {
		this1 = (*C.GDesktopAppInfo)(this0.InheritedFromGDesktopAppInfo())
	}
	ret1 := C.g_desktop_app_info_get_is_hidden(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *DesktopAppInfo) GetKeywords() []string {
	var this1 *C.GDesktopAppInfo
	if this0 != nil {
		this1 = (*C.GDesktopAppInfo)(this0.InheritedFromGDesktopAppInfo())
	}
	ret1 := C.g_desktop_app_info_get_keywords(this1)
	var ret2 []string
	ret2 = make([]string, uint(C._array_length(unsafe.Pointer(ret1))))
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
	}
	return ret2
}
func (this0 *DesktopAppInfo) GetNodisplay() bool {
	var this1 *C.GDesktopAppInfo
	if this0 != nil {
		this1 = (*C.GDesktopAppInfo)(this0.InheritedFromGDesktopAppInfo())
	}
	ret1 := C.g_desktop_app_info_get_nodisplay(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *DesktopAppInfo) GetShowIn(desktop_env0 string) bool {
	var this1 *C.GDesktopAppInfo
	var desktop_env1 *C.char
	if this0 != nil {
		this1 = (*C.GDesktopAppInfo)(this0.InheritedFromGDesktopAppInfo())
	}
	desktop_env1 = _GoStringToGString(desktop_env0)
	defer C.free(unsafe.Pointer(desktop_env1))
	ret1 := C.g_desktop_app_info_get_show_in(this1, desktop_env1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *DesktopAppInfo) GetStartupWmClass() string {
	var this1 *C.GDesktopAppInfo
	if this0 != nil {
		this1 = (*C.GDesktopAppInfo)(this0.InheritedFromGDesktopAppInfo())
	}
	ret1 := C.g_desktop_app_info_get_startup_wm_class(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *DesktopAppInfo) GetString(key0 string) string {
	var this1 *C.GDesktopAppInfo
	var key1 *C.char
	if this0 != nil {
		this1 = (*C.GDesktopAppInfo)(this0.InheritedFromGDesktopAppInfo())
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_desktop_app_info_get_string(this1, key1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *DesktopAppInfo) HasKey(key0 string) bool {
	var this1 *C.GDesktopAppInfo
	var key1 *C.char
	if this0 != nil {
		this1 = (*C.GDesktopAppInfo)(this0.InheritedFromGDesktopAppInfo())
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_desktop_app_info_has_key(this1, key1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *DesktopAppInfo) LaunchAction(action_name0 string, launch_context0 AppLaunchContextLike) {
	var this1 *C.GDesktopAppInfo
	var action_name1 *C.char
	var launch_context1 *C.GAppLaunchContext
	if this0 != nil {
		this1 = (*C.GDesktopAppInfo)(this0.InheritedFromGDesktopAppInfo())
	}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	if launch_context0 != nil {
		launch_context1 = (*C.GAppLaunchContext)(launch_context0.InheritedFromGAppLaunchContext())
	}
	C.g_desktop_app_info_launch_action(this1, action_name1, launch_context1)
}
// blacklisted: DesktopAppInfo.launch_uris_as_manager (method)
func (this0 *DesktopAppInfo) ListActions() []string {
	var this1 *C.GDesktopAppInfo
	if this0 != nil {
		this1 = (*C.GDesktopAppInfo)(this0.InheritedFromGDesktopAppInfo())
	}
	ret1 := C.g_desktop_app_info_list_actions(this1)
	var ret2 []string
	ret2 = make([]string, uint(C._array_length(unsafe.Pointer(ret1))))
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
	}
	return ret2
}
// blacklisted: DesktopAppInfoClass (struct)
// blacklisted: DesktopAppInfoLookup (interface)
// blacklisted: DesktopAppInfoLookupIface (struct)
type DesktopAppLaunchCallback func(appinfo *DesktopAppInfo, pid int32)
//export _GDesktopAppLaunchCallback_c_wrapper
func _GDesktopAppLaunchCallback_c_wrapper(appinfo0 unsafe.Pointer, pid0 int32, user_data0 unsafe.Pointer) {
	var appinfo1 *DesktopAppInfo
	var pid1 int32
	var user_data1 DesktopAppLaunchCallback
	appinfo1 = (*DesktopAppInfo)(gobject.ObjectWrap(unsafe.Pointer((*C.GDesktopAppInfo)(appinfo0)), true))
	pid1 = int32((C.int32_t)(pid0))
	user_data1 = *(*DesktopAppLaunchCallback)(user_data0)
	user_data1(appinfo1, pid1)
}
//export _GDesktopAppLaunchCallback_c_wrapper_once
func _GDesktopAppLaunchCallback_c_wrapper_once(appinfo0 unsafe.Pointer, pid0 int32, user_data0 unsafe.Pointer) {
	_GDesktopAppLaunchCallback_c_wrapper(appinfo0, pid0, user_data0)
	gobject.Holder.Release(user_data0)
}
type DriveLike interface {
	ImplementsGDrive() *C.GDrive
}

type Drive struct {
	gobject.Object
	DriveImpl
}

func (*Drive) GetStaticType() gobject.Type {
	return gobject.Type(C.g_drive_get_type())
}


type DriveImpl struct {}

func ToDrive(objlike gobject.ObjectLike) *Drive {
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), gobject.Type(C.g_drive_get_type()))
	if obj != nil {
		return (*Drive)(obj)
	}
	panic("cannot cast to Drive")
}

func (this0 *DriveImpl) ImplementsGDrive() *C.GDrive {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GDrive)((*gobject.Object)(unsafe.Pointer(obj)).C)
}
func (this0 *DriveImpl) CanEject() bool {
	var this1 *C.GDrive
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	ret1 := C.g_drive_can_eject(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *DriveImpl) CanPollForMedia() bool {
	var this1 *C.GDrive
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	ret1 := C.g_drive_can_poll_for_media(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *DriveImpl) CanStart() bool {
	var this1 *C.GDrive
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	ret1 := C.g_drive_can_start(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *DriveImpl) CanStartDegraded() bool {
	var this1 *C.GDrive
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	ret1 := C.g_drive_can_start_degraded(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *DriveImpl) CanStop() bool {
	var this1 *C.GDrive
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	ret1 := C.g_drive_can_stop(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *DriveImpl) Eject(flags0 MountUnmountFlags, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GDrive
	var flags1 C.GMountUnmountFlags
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	flags1 = C.GMountUnmountFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_drive_eject(this1, flags1, cancellable1, callback1)
}
func (this0 *DriveImpl) EjectFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GDrive
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_drive_eject_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *DriveImpl) EjectWithOperation(flags0 MountUnmountFlags, mount_operation0 MountOperationLike, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GDrive
	var flags1 C.GMountUnmountFlags
	var mount_operation1 *C.GMountOperation
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	flags1 = C.GMountUnmountFlags(flags0)
	if mount_operation0 != nil {
		mount_operation1 = (*C.GMountOperation)(mount_operation0.InheritedFromGMountOperation())
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_drive_eject_with_operation(this1, flags1, mount_operation1, cancellable1, callback1)
}
func (this0 *DriveImpl) EjectWithOperationFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GDrive
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_drive_eject_with_operation_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *DriveImpl) EnumerateIdentifiers() []string {
	var this1 *C.GDrive
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	ret1 := C.g_drive_enumerate_identifiers(this1)
	var ret2 []string
	ret2 = make([]string, uint(C._array_length(unsafe.Pointer(ret1))))
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return ret2
}
func (this0 *DriveImpl) GetIcon() *Icon {
	var this1 *C.GDrive
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	ret1 := C.g_drive_get_icon(this1)
	var ret2 *Icon
	ret2 = (*Icon)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *DriveImpl) GetIdentifier(kind0 string) string {
	var this1 *C.GDrive
	var kind1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	kind1 = _GoStringToGString(kind0)
	defer C.free(unsafe.Pointer(kind1))
	ret1 := C.g_drive_get_identifier(this1, kind1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *DriveImpl) GetName() string {
	var this1 *C.GDrive
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	ret1 := C.g_drive_get_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *DriveImpl) GetSortKey() string {
	var this1 *C.GDrive
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	ret1 := C.g_drive_get_sort_key(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *DriveImpl) GetStartStopType() DriveStartStopType {
	var this1 *C.GDrive
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	ret1 := C.g_drive_get_start_stop_type(this1)
	var ret2 DriveStartStopType
	ret2 = DriveStartStopType(ret1)
	return ret2
}
func (this0 *DriveImpl) GetSymbolicIcon() *Icon {
	var this1 *C.GDrive
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	ret1 := C.g_drive_get_symbolic_icon(this1)
	var ret2 *Icon
	ret2 = (*Icon)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *DriveImpl) GetVolumes() []*Volume {
	var this1 *C.GDrive
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	ret1 := C.g_drive_get_volumes(this1)
	var ret2 []*Volume
	for iter := (*_GList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt *Volume
		elt = (*Volume)(gobject.ObjectWrap(unsafe.Pointer((*C.GVolume)(iter.data)), false))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func (this0 *DriveImpl) HasMedia() bool {
	var this1 *C.GDrive
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	ret1 := C.g_drive_has_media(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *DriveImpl) HasVolumes() bool {
	var this1 *C.GDrive
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	ret1 := C.g_drive_has_volumes(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *DriveImpl) IsMediaCheckAutomatic() bool {
	var this1 *C.GDrive
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	ret1 := C.g_drive_is_media_check_automatic(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *DriveImpl) IsMediaRemovable() bool {
	var this1 *C.GDrive
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	ret1 := C.g_drive_is_media_removable(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *DriveImpl) PollForMedia(cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GDrive
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_drive_poll_for_media(this1, cancellable1, callback1)
}
func (this0 *DriveImpl) PollForMediaFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GDrive
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_drive_poll_for_media_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *DriveImpl) Start(flags0 DriveStartFlags, mount_operation0 MountOperationLike, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GDrive
	var flags1 C.GDriveStartFlags
	var mount_operation1 *C.GMountOperation
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	flags1 = C.GDriveStartFlags(flags0)
	if mount_operation0 != nil {
		mount_operation1 = (*C.GMountOperation)(mount_operation0.InheritedFromGMountOperation())
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_drive_start(this1, flags1, mount_operation1, cancellable1, callback1)
}
func (this0 *DriveImpl) StartFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GDrive
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_drive_start_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *DriveImpl) Stop(flags0 MountUnmountFlags, mount_operation0 MountOperationLike, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GDrive
	var flags1 C.GMountUnmountFlags
	var mount_operation1 *C.GMountOperation
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	flags1 = C.GMountUnmountFlags(flags0)
	if mount_operation0 != nil {
		mount_operation1 = (*C.GMountOperation)(mount_operation0.InheritedFromGMountOperation())
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_drive_stop(this1, flags1, mount_operation1, cancellable1, callback1)
}
func (this0 *DriveImpl) StopFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GDrive
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGDrive()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_drive_stop_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
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
		this1 = (*C.GEmblem)(this0.InheritedFromGEmblem())
	}
	ret1 := C.g_emblem_get_icon(this1)
	var ret2 *Icon
	ret2 = (*Icon)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Emblem) GetOrigin() EmblemOrigin {
	var this1 *C.GEmblem
	if this0 != nil {
		this1 = (*C.GEmblem)(this0.InheritedFromGEmblem())
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
		emblem1 = (*C.GEmblem)(emblem0.InheritedFromGEmblem())
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
		this1 = (*C.GEmblemedIcon)(this0.InheritedFromGEmblemedIcon())
	}
	if emblem0 != nil {
		emblem1 = (*C.GEmblem)(emblem0.InheritedFromGEmblem())
	}
	C.g_emblemed_icon_add_emblem(this1, emblem1)
}
func (this0 *EmblemedIcon) ClearEmblems() {
	var this1 *C.GEmblemedIcon
	if this0 != nil {
		this1 = (*C.GEmblemedIcon)(this0.InheritedFromGEmblemedIcon())
	}
	C.g_emblemed_icon_clear_emblems(this1)
}
func (this0 *EmblemedIcon) GetEmblems() []*Emblem {
	var this1 *C.GEmblemedIcon
	if this0 != nil {
		this1 = (*C.GEmblemedIcon)(this0.InheritedFromGEmblemedIcon())
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
		this1 = (*C.GEmblemedIcon)(this0.InheritedFromGEmblemedIcon())
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
const FileAttributeThumbnailIsValid = "thumbnail::is-valid"
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

func (*File) GetStaticType() gobject.Type {
	return gobject.Type(C.g_file_get_type())
}


type FileImpl struct {}

func ToFile(objlike gobject.ObjectLike) *File {
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), gobject.Type(C.g_file_get_type()))
	if obj != nil {
		return (*File)(obj)
	}
	panic("cannot cast to File")
}

func (this0 *FileImpl) ImplementsGFile() *C.GFile {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GFile)((*gobject.Object)(unsafe.Pointer(obj)).C)
}
// blacklisted: File.new_for_commandline_arg (method)
// blacklisted: File.new_for_commandline_arg_and_cwd (method)
// blacklisted: File.new_for_path (method)
// blacklisted: File.new_for_uri (method)
// blacklisted: File.new_tmp (method)
// blacklisted: File.parse_name (method)
func (this0 *FileImpl) AppendTo(flags0 FileCreateFlags, cancellable0 CancellableLike) (*FileOutputStream, error) {
	var this1 *C.GFile
	var flags1 C.GFileCreateFlags
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	flags1 = C.GFileCreateFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_append_to(this1, flags1, cancellable1, &err1)
	var ret2 *FileOutputStream
	var err2 error
	ret2 = (*FileOutputStream)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) AppendToAsync(flags0 FileCreateFlags, io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var flags1 C.GFileCreateFlags
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	flags1 = C.GFileCreateFlags(flags0)
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_append_to_async(this1, flags1, io_priority1, cancellable1, callback1)
}
func (this0 *FileImpl) AppendToFinish(res0 AsyncResultLike) (*FileOutputStream, error) {
	var this1 *C.GFile
	var res1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if res0 != nil {
		res1 = res0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_append_to_finish(this1, res1, &err1)
	var ret2 *FileOutputStream
	var err2 error
	ret2 = (*FileOutputStream)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) Copy(destination0 FileLike, flags0 FileCopyFlags, cancellable0 CancellableLike, progress_callback0 FileProgressCallback) (bool, error) {
	var this1 *C.GFile
	var destination1 *C.GFile
	var flags1 C.GFileCopyFlags
	var cancellable1 *C.GCancellable
	var progress_callback1 unsafe.Pointer
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if destination0 != nil {
		destination1 = destination0.ImplementsGFile()
	}
	flags1 = C.GFileCopyFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if progress_callback0 != nil {
		progress_callback1 = unsafe.Pointer(&progress_callback0)}
	ret1 := C._g_file_copy(this1, destination1, flags1, cancellable1, progress_callback1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) CopyAttributes(destination0 FileLike, flags0 FileCopyFlags, cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GFile
	var destination1 *C.GFile
	var flags1 C.GFileCopyFlags
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if destination0 != nil {
		destination1 = destination0.ImplementsGFile()
	}
	flags1 = C.GFileCopyFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_copy_attributes(this1, destination1, flags1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) CopyFinish(res0 AsyncResultLike) (bool, error) {
	var this1 *C.GFile
	var res1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if res0 != nil {
		res1 = res0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_copy_finish(this1, res1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) Create(flags0 FileCreateFlags, cancellable0 CancellableLike) (*FileOutputStream, error) {
	var this1 *C.GFile
	var flags1 C.GFileCreateFlags
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	flags1 = C.GFileCreateFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_create(this1, flags1, cancellable1, &err1)
	var ret2 *FileOutputStream
	var err2 error
	ret2 = (*FileOutputStream)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) CreateAsync(flags0 FileCreateFlags, io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var flags1 C.GFileCreateFlags
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	flags1 = C.GFileCreateFlags(flags0)
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_create_async(this1, flags1, io_priority1, cancellable1, callback1)
}
func (this0 *FileImpl) CreateFinish(res0 AsyncResultLike) (*FileOutputStream, error) {
	var this1 *C.GFile
	var res1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if res0 != nil {
		res1 = res0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_create_finish(this1, res1, &err1)
	var ret2 *FileOutputStream
	var err2 error
	ret2 = (*FileOutputStream)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) CreateReadwrite(flags0 FileCreateFlags, cancellable0 CancellableLike) (*FileIOStream, error) {
	var this1 *C.GFile
	var flags1 C.GFileCreateFlags
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	flags1 = C.GFileCreateFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_create_readwrite(this1, flags1, cancellable1, &err1)
	var ret2 *FileIOStream
	var err2 error
	ret2 = (*FileIOStream)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) CreateReadwriteAsync(flags0 FileCreateFlags, io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var flags1 C.GFileCreateFlags
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	flags1 = C.GFileCreateFlags(flags0)
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_create_readwrite_async(this1, flags1, io_priority1, cancellable1, callback1)
}
func (this0 *FileImpl) CreateReadwriteFinish(res0 AsyncResultLike) (*FileIOStream, error) {
	var this1 *C.GFile
	var res1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if res0 != nil {
		res1 = res0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_create_readwrite_finish(this1, res1, &err1)
	var ret2 *FileIOStream
	var err2 error
	ret2 = (*FileIOStream)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) Delete(cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GFile
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_delete(this1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) DeleteAsync(io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_delete_async(this1, io_priority1, cancellable1, callback1)
}
func (this0 *FileImpl) DeleteFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GFile
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_delete_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) Dup() *File {
	var this1 *C.GFile
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	ret1 := C.g_file_dup(this1)
	var ret2 *File
	ret2 = (*File)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *FileImpl) EjectMountable(flags0 MountUnmountFlags, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var flags1 C.GMountUnmountFlags
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	flags1 = C.GMountUnmountFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_eject_mountable(this1, flags1, cancellable1, callback1)
}
func (this0 *FileImpl) EjectMountableFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GFile
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_eject_mountable_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) EjectMountableWithOperation(flags0 MountUnmountFlags, mount_operation0 MountOperationLike, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var flags1 C.GMountUnmountFlags
	var mount_operation1 *C.GMountOperation
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	flags1 = C.GMountUnmountFlags(flags0)
	if mount_operation0 != nil {
		mount_operation1 = (*C.GMountOperation)(mount_operation0.InheritedFromGMountOperation())
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_eject_mountable_with_operation(this1, flags1, mount_operation1, cancellable1, callback1)
}
func (this0 *FileImpl) EjectMountableWithOperationFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GFile
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_eject_mountable_with_operation_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) EnumerateChildren(attributes0 string, flags0 FileQueryInfoFlags, cancellable0 CancellableLike) (*FileEnumerator, error) {
	var this1 *C.GFile
	var attributes1 *C.char
	var flags1 C.GFileQueryInfoFlags
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	attributes1 = _GoStringToGString(attributes0)
	defer C.free(unsafe.Pointer(attributes1))
	flags1 = C.GFileQueryInfoFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_enumerate_children(this1, attributes1, flags1, cancellable1, &err1)
	var ret2 *FileEnumerator
	var err2 error
	ret2 = (*FileEnumerator)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) EnumerateChildrenAsync(attributes0 string, flags0 FileQueryInfoFlags, io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var attributes1 *C.char
	var flags1 C.GFileQueryInfoFlags
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	attributes1 = _GoStringToGString(attributes0)
	defer C.free(unsafe.Pointer(attributes1))
	flags1 = C.GFileQueryInfoFlags(flags0)
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_enumerate_children_async(this1, attributes1, flags1, io_priority1, cancellable1, callback1)
}
func (this0 *FileImpl) EnumerateChildrenFinish(res0 AsyncResultLike) (*FileEnumerator, error) {
	var this1 *C.GFile
	var res1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if res0 != nil {
		res1 = res0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_enumerate_children_finish(this1, res1, &err1)
	var ret2 *FileEnumerator
	var err2 error
	ret2 = (*FileEnumerator)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) Equal(file20 FileLike) bool {
	var this1 *C.GFile
	var file21 *C.GFile
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if file20 != nil {
		file21 = file20.ImplementsGFile()
	}
	ret1 := C.g_file_equal(this1, file21)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *FileImpl) FindEnclosingMount(cancellable0 CancellableLike) (*Mount, error) {
	var this1 *C.GFile
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_find_enclosing_mount(this1, cancellable1, &err1)
	var ret2 *Mount
	var err2 error
	ret2 = (*Mount)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) FindEnclosingMountAsync(io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_find_enclosing_mount_async(this1, io_priority1, cancellable1, callback1)
}
func (this0 *FileImpl) FindEnclosingMountFinish(res0 AsyncResultLike) (*Mount, error) {
	var this1 *C.GFile
	var res1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if res0 != nil {
		res1 = res0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_find_enclosing_mount_finish(this1, res1, &err1)
	var ret2 *Mount
	var err2 error
	ret2 = (*Mount)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) GetBasename() string {
	var this1 *C.GFile
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	ret1 := C.g_file_get_basename(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *FileImpl) GetChild(name0 string) *File {
	var this1 *C.GFile
	var name1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	ret1 := C.g_file_get_child(this1, name1)
	var ret2 *File
	ret2 = (*File)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *FileImpl) GetChildForDisplayName(display_name0 string) (*File, error) {
	var this1 *C.GFile
	var display_name1 *C.char
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	display_name1 = _GoStringToGString(display_name0)
	defer C.free(unsafe.Pointer(display_name1))
	ret1 := C.g_file_get_child_for_display_name(this1, display_name1, &err1)
	var ret2 *File
	var err2 error
	ret2 = (*File)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) GetParent() *File {
	var this1 *C.GFile
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	ret1 := C.g_file_get_parent(this1)
	var ret2 *File
	ret2 = (*File)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *FileImpl) GetParseName() string {
	var this1 *C.GFile
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	ret1 := C.g_file_get_parse_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *FileImpl) GetPath() string {
	var this1 *C.GFile
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	ret1 := C.g_file_get_path(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *FileImpl) GetRelativePath(descendant0 FileLike) string {
	var this1 *C.GFile
	var descendant1 *C.GFile
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if descendant0 != nil {
		descendant1 = descendant0.ImplementsGFile()
	}
	ret1 := C.g_file_get_relative_path(this1, descendant1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *FileImpl) GetUri() string {
	var this1 *C.GFile
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	ret1 := C.g_file_get_uri(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *FileImpl) GetUriScheme() string {
	var this1 *C.GFile
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	ret1 := C.g_file_get_uri_scheme(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *FileImpl) HasParent(parent0 FileLike) bool {
	var this1 *C.GFile
	var parent1 *C.GFile
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if parent0 != nil {
		parent1 = parent0.ImplementsGFile()
	}
	ret1 := C.g_file_has_parent(this1, parent1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *FileImpl) HasPrefix(prefix0 FileLike) bool {
	var this1 *C.GFile
	var prefix1 *C.GFile
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if prefix0 != nil {
		prefix1 = prefix0.ImplementsGFile()
	}
	ret1 := C.g_file_has_prefix(this1, prefix1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *FileImpl) HasUriScheme(uri_scheme0 string) bool {
	var this1 *C.GFile
	var uri_scheme1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	uri_scheme1 = _GoStringToGString(uri_scheme0)
	defer C.free(unsafe.Pointer(uri_scheme1))
	ret1 := C.g_file_has_uri_scheme(this1, uri_scheme1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *FileImpl) Hash() uint32 {
	var this1 *C.GFile
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	ret1 := C.g_file_hash(this1)
	var ret2 uint32
	ret2 = uint32(ret1)
	return ret2
}
func (this0 *FileImpl) IsNative() bool {
	var this1 *C.GFile
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	ret1 := C.g_file_is_native(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *FileImpl) LoadContents(cancellable0 CancellableLike) ([]uint8, string, bool, error) {
	var this1 *C.GFile
	var cancellable1 *C.GCancellable
	var contents1 *C.uint8_t
	var length1 C.uint64_t
	var etag_out1 *C.char
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_load_contents(this1, cancellable1, &contents1, &length1, &etag_out1, &err1)
	var contents2 []uint8
	var etag_out2 string
	var ret2 bool
	var err2 error
	contents2 = make([]uint8, length1)
	for i := range contents2 {
		contents2[i] = uint8((*(*[999999]C.uint8_t)(unsafe.Pointer(contents1)))[i])
	}
	etag_out2 = C.GoString(etag_out1)
	C.g_free(unsafe.Pointer(etag_out1))
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return contents2, etag_out2, ret2, err2
}
func (this0 *FileImpl) LoadContentsAsync(cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_load_contents_async(this1, cancellable1, callback1)
}
func (this0 *FileImpl) LoadContentsFinish(res0 AsyncResultLike) ([]uint8, string, bool, error) {
	var this1 *C.GFile
	var res1 *C.GAsyncResult
	var contents1 *C.uint8_t
	var length1 C.uint64_t
	var etag_out1 *C.char
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if res0 != nil {
		res1 = res0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_load_contents_finish(this1, res1, &contents1, &length1, &etag_out1, &err1)
	var contents2 []uint8
	var etag_out2 string
	var ret2 bool
	var err2 error
	contents2 = make([]uint8, length1)
	for i := range contents2 {
		contents2[i] = uint8((*(*[999999]C.uint8_t)(unsafe.Pointer(contents1)))[i])
	}
	etag_out2 = C.GoString(etag_out1)
	C.g_free(unsafe.Pointer(etag_out1))
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return contents2, etag_out2, ret2, err2
}
func (this0 *FileImpl) LoadPartialContentsFinish(res0 AsyncResultLike) ([]uint8, string, bool, error) {
	var this1 *C.GFile
	var res1 *C.GAsyncResult
	var contents1 *C.uint8_t
	var length1 C.uint64_t
	var etag_out1 *C.char
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if res0 != nil {
		res1 = res0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_load_partial_contents_finish(this1, res1, &contents1, &length1, &etag_out1, &err1)
	var contents2 []uint8
	var etag_out2 string
	var ret2 bool
	var err2 error
	contents2 = make([]uint8, length1)
	for i := range contents2 {
		contents2[i] = uint8((*(*[999999]C.uint8_t)(unsafe.Pointer(contents1)))[i])
	}
	etag_out2 = C.GoString(etag_out1)
	C.g_free(unsafe.Pointer(etag_out1))
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return contents2, etag_out2, ret2, err2
}
func (this0 *FileImpl) MakeDirectory(cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GFile
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_make_directory(this1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) MakeDirectoryAsync(io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_make_directory_async(this1, io_priority1, cancellable1, callback1)
}
func (this0 *FileImpl) MakeDirectoryFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GFile
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_make_directory_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) MakeDirectoryWithParents(cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GFile
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_make_directory_with_parents(this1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) MakeSymbolicLink(symlink_value0 string, cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GFile
	var symlink_value1 *C.char
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	symlink_value1 = _GoStringToGString(symlink_value0)
	defer C.free(unsafe.Pointer(symlink_value1))
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_make_symbolic_link(this1, symlink_value1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) MeasureDiskUsageFinish(result0 AsyncResultLike) (uint64, uint64, uint64, bool, error) {
	var this1 *C.GFile
	var result1 *C.GAsyncResult
	var disk_usage1 C.uint64_t
	var num_dirs1 C.uint64_t
	var num_files1 C.uint64_t
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_measure_disk_usage_finish(this1, result1, &disk_usage1, &num_dirs1, &num_files1, &err1)
	var disk_usage2 uint64
	var num_dirs2 uint64
	var num_files2 uint64
	var ret2 bool
	var err2 error
	disk_usage2 = uint64(disk_usage1)
	num_dirs2 = uint64(num_dirs1)
	num_files2 = uint64(num_files1)
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return disk_usage2, num_dirs2, num_files2, ret2, err2
}
func (this0 *FileImpl) Monitor(flags0 FileMonitorFlags, cancellable0 CancellableLike) (*FileMonitor, error) {
	var this1 *C.GFile
	var flags1 C.GFileMonitorFlags
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	flags1 = C.GFileMonitorFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_monitor(this1, flags1, cancellable1, &err1)
	var ret2 *FileMonitor
	var err2 error
	ret2 = (*FileMonitor)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) MonitorDirectory(flags0 FileMonitorFlags, cancellable0 CancellableLike) (*FileMonitor, error) {
	var this1 *C.GFile
	var flags1 C.GFileMonitorFlags
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	flags1 = C.GFileMonitorFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_monitor_directory(this1, flags1, cancellable1, &err1)
	var ret2 *FileMonitor
	var err2 error
	ret2 = (*FileMonitor)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) MonitorFile(flags0 FileMonitorFlags, cancellable0 CancellableLike) (*FileMonitor, error) {
	var this1 *C.GFile
	var flags1 C.GFileMonitorFlags
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	flags1 = C.GFileMonitorFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_monitor_file(this1, flags1, cancellable1, &err1)
	var ret2 *FileMonitor
	var err2 error
	ret2 = (*FileMonitor)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) MountEnclosingVolume(flags0 MountMountFlags, mount_operation0 MountOperationLike, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var flags1 C.GMountMountFlags
	var mount_operation1 *C.GMountOperation
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	flags1 = C.GMountMountFlags(flags0)
	if mount_operation0 != nil {
		mount_operation1 = (*C.GMountOperation)(mount_operation0.InheritedFromGMountOperation())
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_mount_enclosing_volume(this1, flags1, mount_operation1, cancellable1, callback1)
}
func (this0 *FileImpl) MountEnclosingVolumeFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GFile
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_mount_enclosing_volume_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) MountMountable(flags0 MountMountFlags, mount_operation0 MountOperationLike, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var flags1 C.GMountMountFlags
	var mount_operation1 *C.GMountOperation
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	flags1 = C.GMountMountFlags(flags0)
	if mount_operation0 != nil {
		mount_operation1 = (*C.GMountOperation)(mount_operation0.InheritedFromGMountOperation())
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_mount_mountable(this1, flags1, mount_operation1, cancellable1, callback1)
}
func (this0 *FileImpl) MountMountableFinish(result0 AsyncResultLike) (*File, error) {
	var this1 *C.GFile
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_mount_mountable_finish(this1, result1, &err1)
	var ret2 *File
	var err2 error
	ret2 = (*File)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) Move(destination0 FileLike, flags0 FileCopyFlags, cancellable0 CancellableLike, progress_callback0 FileProgressCallback) (bool, error) {
	var this1 *C.GFile
	var destination1 *C.GFile
	var flags1 C.GFileCopyFlags
	var cancellable1 *C.GCancellable
	var progress_callback1 unsafe.Pointer
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if destination0 != nil {
		destination1 = destination0.ImplementsGFile()
	}
	flags1 = C.GFileCopyFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if progress_callback0 != nil {
		progress_callback1 = unsafe.Pointer(&progress_callback0)}
	ret1 := C._g_file_move(this1, destination1, flags1, cancellable1, progress_callback1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) OpenReadwrite(cancellable0 CancellableLike) (*FileIOStream, error) {
	var this1 *C.GFile
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_open_readwrite(this1, cancellable1, &err1)
	var ret2 *FileIOStream
	var err2 error
	ret2 = (*FileIOStream)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) OpenReadwriteAsync(io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_open_readwrite_async(this1, io_priority1, cancellable1, callback1)
}
func (this0 *FileImpl) OpenReadwriteFinish(res0 AsyncResultLike) (*FileIOStream, error) {
	var this1 *C.GFile
	var res1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if res0 != nil {
		res1 = res0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_open_readwrite_finish(this1, res1, &err1)
	var ret2 *FileIOStream
	var err2 error
	ret2 = (*FileIOStream)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) PollMountable(cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_poll_mountable(this1, cancellable1, callback1)
}
func (this0 *FileImpl) PollMountableFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GFile
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_poll_mountable_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) QueryDefaultHandler(cancellable0 CancellableLike) (*AppInfo, error) {
	var this1 *C.GFile
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_query_default_handler(this1, cancellable1, &err1)
	var ret2 *AppInfo
	var err2 error
	ret2 = (*AppInfo)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) QueryExists(cancellable0 CancellableLike) bool {
	var this1 *C.GFile
	var cancellable1 *C.GCancellable
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_query_exists(this1, cancellable1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *FileImpl) QueryFileType(flags0 FileQueryInfoFlags, cancellable0 CancellableLike) FileType {
	var this1 *C.GFile
	var flags1 C.GFileQueryInfoFlags
	var cancellable1 *C.GCancellable
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	flags1 = C.GFileQueryInfoFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_query_file_type(this1, flags1, cancellable1)
	var ret2 FileType
	ret2 = FileType(ret1)
	return ret2
}
func (this0 *FileImpl) QueryFilesystemInfo(attributes0 string, cancellable0 CancellableLike) (*FileInfo, error) {
	var this1 *C.GFile
	var attributes1 *C.char
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	attributes1 = _GoStringToGString(attributes0)
	defer C.free(unsafe.Pointer(attributes1))
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_query_filesystem_info(this1, attributes1, cancellable1, &err1)
	var ret2 *FileInfo
	var err2 error
	ret2 = (*FileInfo)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) QueryFilesystemInfoAsync(attributes0 string, io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var attributes1 *C.char
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	attributes1 = _GoStringToGString(attributes0)
	defer C.free(unsafe.Pointer(attributes1))
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_query_filesystem_info_async(this1, attributes1, io_priority1, cancellable1, callback1)
}
func (this0 *FileImpl) QueryFilesystemInfoFinish(res0 AsyncResultLike) (*FileInfo, error) {
	var this1 *C.GFile
	var res1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if res0 != nil {
		res1 = res0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_query_filesystem_info_finish(this1, res1, &err1)
	var ret2 *FileInfo
	var err2 error
	ret2 = (*FileInfo)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) QueryInfo(attributes0 string, flags0 FileQueryInfoFlags, cancellable0 CancellableLike) (*FileInfo, error) {
	var this1 *C.GFile
	var attributes1 *C.char
	var flags1 C.GFileQueryInfoFlags
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	attributes1 = _GoStringToGString(attributes0)
	defer C.free(unsafe.Pointer(attributes1))
	flags1 = C.GFileQueryInfoFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_query_info(this1, attributes1, flags1, cancellable1, &err1)
	var ret2 *FileInfo
	var err2 error
	ret2 = (*FileInfo)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) QueryInfoAsync(attributes0 string, flags0 FileQueryInfoFlags, io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var attributes1 *C.char
	var flags1 C.GFileQueryInfoFlags
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	attributes1 = _GoStringToGString(attributes0)
	defer C.free(unsafe.Pointer(attributes1))
	flags1 = C.GFileQueryInfoFlags(flags0)
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_query_info_async(this1, attributes1, flags1, io_priority1, cancellable1, callback1)
}
func (this0 *FileImpl) QueryInfoFinish(res0 AsyncResultLike) (*FileInfo, error) {
	var this1 *C.GFile
	var res1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if res0 != nil {
		res1 = res0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_query_info_finish(this1, res1, &err1)
	var ret2 *FileInfo
	var err2 error
	ret2 = (*FileInfo)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) QuerySettableAttributes(cancellable0 CancellableLike) (*FileAttributeInfoList, error) {
	var this1 *C.GFile
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_query_settable_attributes(this1, cancellable1, &err1)
	var ret2 *FileAttributeInfoList
	var err2 error
	ret2 = (*FileAttributeInfoList)(unsafe.Pointer(ret1))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) QueryWritableNamespaces(cancellable0 CancellableLike) (*FileAttributeInfoList, error) {
	var this1 *C.GFile
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_query_writable_namespaces(this1, cancellable1, &err1)
	var ret2 *FileAttributeInfoList
	var err2 error
	ret2 = (*FileAttributeInfoList)(unsafe.Pointer(ret1))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) Read(cancellable0 CancellableLike) (*FileInputStream, error) {
	var this1 *C.GFile
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_read(this1, cancellable1, &err1)
	var ret2 *FileInputStream
	var err2 error
	ret2 = (*FileInputStream)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) ReadAsync(io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_read_async(this1, io_priority1, cancellable1, callback1)
}
func (this0 *FileImpl) ReadFinish(res0 AsyncResultLike) (*FileInputStream, error) {
	var this1 *C.GFile
	var res1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if res0 != nil {
		res1 = res0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_read_finish(this1, res1, &err1)
	var ret2 *FileInputStream
	var err2 error
	ret2 = (*FileInputStream)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) Replace(etag0 string, make_backup0 bool, flags0 FileCreateFlags, cancellable0 CancellableLike) (*FileOutputStream, error) {
	var this1 *C.GFile
	var etag1 *C.char
	var make_backup1 C.int
	var flags1 C.GFileCreateFlags
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	etag1 = _GoStringToGString(etag0)
	defer C.free(unsafe.Pointer(etag1))
	make_backup1 = _GoBoolToCBool(make_backup0)
	flags1 = C.GFileCreateFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_replace(this1, etag1, make_backup1, flags1, cancellable1, &err1)
	var ret2 *FileOutputStream
	var err2 error
	ret2 = (*FileOutputStream)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) ReplaceAsync(etag0 string, make_backup0 bool, flags0 FileCreateFlags, io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var etag1 *C.char
	var make_backup1 C.int
	var flags1 C.GFileCreateFlags
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	etag1 = _GoStringToGString(etag0)
	defer C.free(unsafe.Pointer(etag1))
	make_backup1 = _GoBoolToCBool(make_backup0)
	flags1 = C.GFileCreateFlags(flags0)
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_replace_async(this1, etag1, make_backup1, flags1, io_priority1, cancellable1, callback1)
}
func (this0 *FileImpl) ReplaceContents(contents0 []uint8, etag0 string, make_backup0 bool, flags0 FileCreateFlags, cancellable0 CancellableLike) (string, bool, error) {
	var this1 *C.GFile
	var contents1 *C.uint8_t
	var length1 C.uint64_t
	var etag1 *C.char
	var make_backup1 C.int
	var flags1 C.GFileCreateFlags
	var cancellable1 *C.GCancellable
	var new_etag1 *C.char
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	contents1 = (*C.uint8_t)(C.malloc(C.size_t(int(unsafe.Sizeof(*contents1)) * len(contents0))))
	defer C.free(unsafe.Pointer(contents1))
	for i, e := range contents0 {
		(*(*[999999]C.uint8_t)(unsafe.Pointer(contents1)))[i] = C.uint8_t(e)
	}
	length1 = C.uint64_t(len(contents0))
	etag1 = _GoStringToGString(etag0)
	defer C.free(unsafe.Pointer(etag1))
	make_backup1 = _GoBoolToCBool(make_backup0)
	flags1 = C.GFileCreateFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_replace_contents(this1, contents1, length1, etag1, make_backup1, flags1, &new_etag1, cancellable1, &err1)
	var new_etag2 string
	var ret2 bool
	var err2 error
	new_etag2 = C.GoString(new_etag1)
	C.g_free(unsafe.Pointer(new_etag1))
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return new_etag2, ret2, err2
}
func (this0 *FileImpl) ReplaceContentsAsync(contents0 []uint8, etag0 string, make_backup0 bool, flags0 FileCreateFlags, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var contents1 *C.uint8_t
	var length1 C.uint64_t
	var etag1 *C.char
	var make_backup1 C.int
	var flags1 C.GFileCreateFlags
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	contents1 = (*C.uint8_t)(C.malloc(C.size_t(int(unsafe.Sizeof(*contents1)) * len(contents0))))
	defer C.free(unsafe.Pointer(contents1))
	for i, e := range contents0 {
		(*(*[999999]C.uint8_t)(unsafe.Pointer(contents1)))[i] = C.uint8_t(e)
	}
	length1 = C.uint64_t(len(contents0))
	etag1 = _GoStringToGString(etag0)
	defer C.free(unsafe.Pointer(etag1))
	make_backup1 = _GoBoolToCBool(make_backup0)
	flags1 = C.GFileCreateFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_replace_contents_async(this1, contents1, length1, etag1, make_backup1, flags1, cancellable1, callback1)
}
// blacklisted: File.replace_contents_bytes_async (method)
func (this0 *FileImpl) ReplaceContentsFinish(res0 AsyncResultLike) (string, bool, error) {
	var this1 *C.GFile
	var res1 *C.GAsyncResult
	var new_etag1 *C.char
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if res0 != nil {
		res1 = res0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_replace_contents_finish(this1, res1, &new_etag1, &err1)
	var new_etag2 string
	var ret2 bool
	var err2 error
	new_etag2 = C.GoString(new_etag1)
	C.g_free(unsafe.Pointer(new_etag1))
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return new_etag2, ret2, err2
}
func (this0 *FileImpl) ReplaceFinish(res0 AsyncResultLike) (*FileOutputStream, error) {
	var this1 *C.GFile
	var res1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if res0 != nil {
		res1 = res0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_replace_finish(this1, res1, &err1)
	var ret2 *FileOutputStream
	var err2 error
	ret2 = (*FileOutputStream)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) ReplaceReadwrite(etag0 string, make_backup0 bool, flags0 FileCreateFlags, cancellable0 CancellableLike) (*FileIOStream, error) {
	var this1 *C.GFile
	var etag1 *C.char
	var make_backup1 C.int
	var flags1 C.GFileCreateFlags
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	etag1 = _GoStringToGString(etag0)
	defer C.free(unsafe.Pointer(etag1))
	make_backup1 = _GoBoolToCBool(make_backup0)
	flags1 = C.GFileCreateFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_replace_readwrite(this1, etag1, make_backup1, flags1, cancellable1, &err1)
	var ret2 *FileIOStream
	var err2 error
	ret2 = (*FileIOStream)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) ReplaceReadwriteAsync(etag0 string, make_backup0 bool, flags0 FileCreateFlags, io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var etag1 *C.char
	var make_backup1 C.int
	var flags1 C.GFileCreateFlags
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	etag1 = _GoStringToGString(etag0)
	defer C.free(unsafe.Pointer(etag1))
	make_backup1 = _GoBoolToCBool(make_backup0)
	flags1 = C.GFileCreateFlags(flags0)
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_replace_readwrite_async(this1, etag1, make_backup1, flags1, io_priority1, cancellable1, callback1)
}
func (this0 *FileImpl) ReplaceReadwriteFinish(res0 AsyncResultLike) (*FileIOStream, error) {
	var this1 *C.GFile
	var res1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if res0 != nil {
		res1 = res0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_replace_readwrite_finish(this1, res1, &err1)
	var ret2 *FileIOStream
	var err2 error
	ret2 = (*FileIOStream)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) ResolveRelativePath(relative_path0 string) *File {
	var this1 *C.GFile
	var relative_path1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	relative_path1 = _GoStringToGString(relative_path0)
	defer C.free(unsafe.Pointer(relative_path1))
	ret1 := C.g_file_resolve_relative_path(this1, relative_path1)
	var ret2 *File
	ret2 = (*File)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *FileImpl) SetAttribute(attribute0 string, type0 FileAttributeType, value_p0 unsafe.Pointer, flags0 FileQueryInfoFlags, cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GFile
	var attribute1 *C.char
	var type1 C.GFileAttributeType
	var value_p1 unsafe.Pointer
	var flags1 C.GFileQueryInfoFlags
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	type1 = C.GFileAttributeType(type0)
	value_p1 = unsafe.Pointer(value_p0)
	flags1 = C.GFileQueryInfoFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_set_attribute(this1, attribute1, type1, value_p1, flags1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) SetAttributeByteString(attribute0 string, value0 string, flags0 FileQueryInfoFlags, cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GFile
	var attribute1 *C.char
	var value1 *C.char
	var flags1 C.GFileQueryInfoFlags
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	value1 = _GoStringToGString(value0)
	defer C.free(unsafe.Pointer(value1))
	flags1 = C.GFileQueryInfoFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_set_attribute_byte_string(this1, attribute1, value1, flags1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) SetAttributeInt32(attribute0 string, value0 int32, flags0 FileQueryInfoFlags, cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GFile
	var attribute1 *C.char
	var value1 C.int32_t
	var flags1 C.GFileQueryInfoFlags
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	value1 = C.int32_t(value0)
	flags1 = C.GFileQueryInfoFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_set_attribute_int32(this1, attribute1, value1, flags1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) SetAttributeInt64(attribute0 string, value0 int64, flags0 FileQueryInfoFlags, cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GFile
	var attribute1 *C.char
	var value1 C.int64_t
	var flags1 C.GFileQueryInfoFlags
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	value1 = C.int64_t(value0)
	flags1 = C.GFileQueryInfoFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_set_attribute_int64(this1, attribute1, value1, flags1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) SetAttributeString(attribute0 string, value0 string, flags0 FileQueryInfoFlags, cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GFile
	var attribute1 *C.char
	var value1 *C.char
	var flags1 C.GFileQueryInfoFlags
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	value1 = _GoStringToGString(value0)
	defer C.free(unsafe.Pointer(value1))
	flags1 = C.GFileQueryInfoFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_set_attribute_string(this1, attribute1, value1, flags1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) SetAttributeUint32(attribute0 string, value0 uint32, flags0 FileQueryInfoFlags, cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GFile
	var attribute1 *C.char
	var value1 C.uint32_t
	var flags1 C.GFileQueryInfoFlags
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	value1 = C.uint32_t(value0)
	flags1 = C.GFileQueryInfoFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_set_attribute_uint32(this1, attribute1, value1, flags1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) SetAttributeUint64(attribute0 string, value0 uint64, flags0 FileQueryInfoFlags, cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GFile
	var attribute1 *C.char
	var value1 C.uint64_t
	var flags1 C.GFileQueryInfoFlags
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	value1 = C.uint64_t(value0)
	flags1 = C.GFileQueryInfoFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_set_attribute_uint64(this1, attribute1, value1, flags1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) SetAttributesAsync(info0 FileInfoLike, flags0 FileQueryInfoFlags, io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var info1 *C.GFileInfo
	var flags1 C.GFileQueryInfoFlags
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if info0 != nil {
		info1 = (*C.GFileInfo)(info0.InheritedFromGFileInfo())
	}
	flags1 = C.GFileQueryInfoFlags(flags0)
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_set_attributes_async(this1, info1, flags1, io_priority1, cancellable1, callback1)
}
func (this0 *FileImpl) SetAttributesFinish(result0 AsyncResultLike) (*FileInfo, bool, error) {
	var this1 *C.GFile
	var result1 *C.GAsyncResult
	var info1 *C.GFileInfo
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_set_attributes_finish(this1, result1, &info1, &err1)
	var info2 *FileInfo
	var ret2 bool
	var err2 error
	info2 = (*FileInfo)(gobject.ObjectWrap(unsafe.Pointer(info1), false))
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return info2, ret2, err2
}
func (this0 *FileImpl) SetAttributesFromInfo(info0 FileInfoLike, flags0 FileQueryInfoFlags, cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GFile
	var info1 *C.GFileInfo
	var flags1 C.GFileQueryInfoFlags
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if info0 != nil {
		info1 = (*C.GFileInfo)(info0.InheritedFromGFileInfo())
	}
	flags1 = C.GFileQueryInfoFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_set_attributes_from_info(this1, info1, flags1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) SetDisplayName(display_name0 string, cancellable0 CancellableLike) (*File, error) {
	var this1 *C.GFile
	var display_name1 *C.char
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	display_name1 = _GoStringToGString(display_name0)
	defer C.free(unsafe.Pointer(display_name1))
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_set_display_name(this1, display_name1, cancellable1, &err1)
	var ret2 *File
	var err2 error
	ret2 = (*File)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) SetDisplayNameAsync(display_name0 string, io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var display_name1 *C.char
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	display_name1 = _GoStringToGString(display_name0)
	defer C.free(unsafe.Pointer(display_name1))
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_set_display_name_async(this1, display_name1, io_priority1, cancellable1, callback1)
}
func (this0 *FileImpl) SetDisplayNameFinish(res0 AsyncResultLike) (*File, error) {
	var this1 *C.GFile
	var res1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if res0 != nil {
		res1 = res0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_set_display_name_finish(this1, res1, &err1)
	var ret2 *File
	var err2 error
	ret2 = (*File)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) StartMountable(flags0 DriveStartFlags, start_operation0 MountOperationLike, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var flags1 C.GDriveStartFlags
	var start_operation1 *C.GMountOperation
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	flags1 = C.GDriveStartFlags(flags0)
	if start_operation0 != nil {
		start_operation1 = (*C.GMountOperation)(start_operation0.InheritedFromGMountOperation())
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_start_mountable(this1, flags1, start_operation1, cancellable1, callback1)
}
func (this0 *FileImpl) StartMountableFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GFile
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_start_mountable_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) StopMountable(flags0 MountUnmountFlags, mount_operation0 MountOperationLike, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var flags1 C.GMountUnmountFlags
	var mount_operation1 *C.GMountOperation
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	flags1 = C.GMountUnmountFlags(flags0)
	if mount_operation0 != nil {
		mount_operation1 = (*C.GMountOperation)(mount_operation0.InheritedFromGMountOperation())
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_stop_mountable(this1, flags1, mount_operation1, cancellable1, callback1)
}
func (this0 *FileImpl) StopMountableFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GFile
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_stop_mountable_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) SupportsThreadContexts() bool {
	var this1 *C.GFile
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	ret1 := C.g_file_supports_thread_contexts(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *FileImpl) Trash(cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GFile
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_trash(this1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) TrashAsync(io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_trash_async(this1, io_priority1, cancellable1, callback1)
}
func (this0 *FileImpl) TrashFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GFile
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_trash_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) UnmountMountable(flags0 MountUnmountFlags, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var flags1 C.GMountUnmountFlags
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	flags1 = C.GMountUnmountFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_unmount_mountable(this1, flags1, cancellable1, callback1)
}
func (this0 *FileImpl) UnmountMountableFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GFile
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_unmount_mountable_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileImpl) UnmountMountableWithOperation(flags0 MountUnmountFlags, mount_operation0 MountOperationLike, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFile
	var flags1 C.GMountUnmountFlags
	var mount_operation1 *C.GMountOperation
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	flags1 = C.GMountUnmountFlags(flags0)
	if mount_operation0 != nil {
		mount_operation1 = (*C.GMountOperation)(mount_operation0.InheritedFromGMountOperation())
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_unmount_mountable_with_operation(this1, flags1, mount_operation1, cancellable1, callback1)
}
func (this0 *FileImpl) UnmountMountableWithOperationFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GFile
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGFile()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_unmount_mountable_with_operation_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
type FileAttributeInfo struct {
	name0 *C.char
	Type FileAttributeType
	Flags FileAttributeInfoFlags
}
func (this0 *FileAttributeInfo) Name() string {
	var name1 string
	name1 = C.GoString(this0.name0)
	return name1
}
type FileAttributeInfoFlags C.uint32_t
const (
	FileAttributeInfoFlagsNone FileAttributeInfoFlags = 0
	FileAttributeInfoFlagsCopyWithFile FileAttributeInfoFlags = 1
	FileAttributeInfoFlagsCopyWhenMoved FileAttributeInfoFlags = 2
)
type FileAttributeInfoList struct {
	Infos *FileAttributeInfo
	NInfos int32
	_ [4]byte
}
func NewFileAttributeInfoList() *FileAttributeInfoList {
	ret1 := C.g_file_attribute_info_list_new()
	var ret2 *FileAttributeInfoList
	ret2 = (*FileAttributeInfoList)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *FileAttributeInfoList) Add(name0 string, type0 FileAttributeType, flags0 FileAttributeInfoFlags) {
	var this1 *C.GFileAttributeInfoList
	var name1 *C.char
	var type1 C.GFileAttributeType
	var flags1 C.GFileAttributeInfoFlags
	this1 = (*C.GFileAttributeInfoList)(unsafe.Pointer(this0))
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	type1 = C.GFileAttributeType(type0)
	flags1 = C.GFileAttributeInfoFlags(flags0)
	C.g_file_attribute_info_list_add(this1, name1, type1, flags1)
}
func (this0 *FileAttributeInfoList) Dup() *FileAttributeInfoList {
	var this1 *C.GFileAttributeInfoList
	this1 = (*C.GFileAttributeInfoList)(unsafe.Pointer(this0))
	ret1 := C.g_file_attribute_info_list_dup(this1)
	var ret2 *FileAttributeInfoList
	ret2 = (*FileAttributeInfoList)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *FileAttributeInfoList) Lookup(name0 string) *FileAttributeInfo {
	var this1 *C.GFileAttributeInfoList
	var name1 *C.char
	this1 = (*C.GFileAttributeInfoList)(unsafe.Pointer(this0))
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	ret1 := C.g_file_attribute_info_list_lookup(this1, name1)
	var ret2 *FileAttributeInfo
	ret2 = (*FileAttributeInfo)(unsafe.Pointer(ret1))
	return ret2
}
type FileAttributeMatcher struct {}
func NewFileAttributeMatcher(attributes0 string) *FileAttributeMatcher {
	var attributes1 *C.char
	attributes1 = _GoStringToGString(attributes0)
	defer C.free(unsafe.Pointer(attributes1))
	ret1 := C.g_file_attribute_matcher_new(attributes1)
	var ret2 *FileAttributeMatcher
	ret2 = (*FileAttributeMatcher)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *FileAttributeMatcher) EnumerateNamespace(ns0 string) bool {
	var this1 *C.GFileAttributeMatcher
	var ns1 *C.char
	this1 = (*C.GFileAttributeMatcher)(unsafe.Pointer(this0))
	ns1 = _GoStringToGString(ns0)
	defer C.free(unsafe.Pointer(ns1))
	ret1 := C.g_file_attribute_matcher_enumerate_namespace(this1, ns1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *FileAttributeMatcher) EnumerateNext() string {
	var this1 *C.GFileAttributeMatcher
	this1 = (*C.GFileAttributeMatcher)(unsafe.Pointer(this0))
	ret1 := C.g_file_attribute_matcher_enumerate_next(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *FileAttributeMatcher) Matches(attribute0 string) bool {
	var this1 *C.GFileAttributeMatcher
	var attribute1 *C.char
	this1 = (*C.GFileAttributeMatcher)(unsafe.Pointer(this0))
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	ret1 := C.g_file_attribute_matcher_matches(this1, attribute1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *FileAttributeMatcher) MatchesOnly(attribute0 string) bool {
	var this1 *C.GFileAttributeMatcher
	var attribute1 *C.char
	this1 = (*C.GFileAttributeMatcher)(unsafe.Pointer(this0))
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	ret1 := C.g_file_attribute_matcher_matches_only(this1, attribute1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *FileAttributeMatcher) Subtract(subtract0 *FileAttributeMatcher) *FileAttributeMatcher {
	var this1 *C.GFileAttributeMatcher
	var subtract1 *C.GFileAttributeMatcher
	this1 = (*C.GFileAttributeMatcher)(unsafe.Pointer(this0))
	subtract1 = (*C.GFileAttributeMatcher)(unsafe.Pointer(subtract0))
	ret1 := C.g_file_attribute_matcher_subtract(this1, subtract1)
	var ret2 *FileAttributeMatcher
	ret2 = (*FileAttributeMatcher)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *FileAttributeMatcher) ToString() string {
	var this1 *C.GFileAttributeMatcher
	this1 = (*C.GFileAttributeMatcher)(unsafe.Pointer(this0))
	ret1 := C.g_file_attribute_matcher_to_string(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
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
type FileEnumeratorLike interface {
	gobject.ObjectLike
	InheritedFromGFileEnumerator() *C.GFileEnumerator
}

type FileEnumerator struct {
	gobject.Object
	
}

func ToFileEnumerator(objlike gobject.ObjectLike) *FileEnumerator {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*FileEnumerator)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*FileEnumerator)(obj)
	}
	panic("cannot cast to FileEnumerator")
}

func (this0 *FileEnumerator) InheritedFromGFileEnumerator() *C.GFileEnumerator {
	if this0 == nil {
		return nil
	}
	return (*C.GFileEnumerator)(this0.C)
}

func (this0 *FileEnumerator) GetStaticType() gobject.Type {
	return gobject.Type(C.g_file_enumerator_get_type())
}

func FileEnumeratorGetType() gobject.Type {
	return (*FileEnumerator)(nil).GetStaticType()
}
func (this0 *FileEnumerator) Close(cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GFileEnumerator
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GFileEnumerator)(this0.InheritedFromGFileEnumerator())
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_enumerator_close(this1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileEnumerator) CloseAsync(io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFileEnumerator
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = (*C.GFileEnumerator)(this0.InheritedFromGFileEnumerator())
	}
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_enumerator_close_async(this1, io_priority1, cancellable1, callback1)
}
func (this0 *FileEnumerator) CloseFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GFileEnumerator
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GFileEnumerator)(this0.InheritedFromGFileEnumerator())
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_enumerator_close_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileEnumerator) GetChild(info0 FileInfoLike) *File {
	var this1 *C.GFileEnumerator
	var info1 *C.GFileInfo
	if this0 != nil {
		this1 = (*C.GFileEnumerator)(this0.InheritedFromGFileEnumerator())
	}
	if info0 != nil {
		info1 = (*C.GFileInfo)(info0.InheritedFromGFileInfo())
	}
	ret1 := C.g_file_enumerator_get_child(this1, info1)
	var ret2 *File
	ret2 = (*File)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *FileEnumerator) GetContainer() *File {
	var this1 *C.GFileEnumerator
	if this0 != nil {
		this1 = (*C.GFileEnumerator)(this0.InheritedFromGFileEnumerator())
	}
	ret1 := C.g_file_enumerator_get_container(this1)
	var ret2 *File
	ret2 = (*File)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *FileEnumerator) HasPending() bool {
	var this1 *C.GFileEnumerator
	if this0 != nil {
		this1 = (*C.GFileEnumerator)(this0.InheritedFromGFileEnumerator())
	}
	ret1 := C.g_file_enumerator_has_pending(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *FileEnumerator) IsClosed() bool {
	var this1 *C.GFileEnumerator
	if this0 != nil {
		this1 = (*C.GFileEnumerator)(this0.InheritedFromGFileEnumerator())
	}
	ret1 := C.g_file_enumerator_is_closed(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *FileEnumerator) NextFile(cancellable0 CancellableLike) (*FileInfo, error) {
	var this1 *C.GFileEnumerator
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GFileEnumerator)(this0.InheritedFromGFileEnumerator())
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_enumerator_next_file(this1, cancellable1, &err1)
	var ret2 *FileInfo
	var err2 error
	ret2 = (*FileInfo)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileEnumerator) NextFilesAsync(num_files0 int32, io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFileEnumerator
	var num_files1 C.int32_t
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = (*C.GFileEnumerator)(this0.InheritedFromGFileEnumerator())
	}
	num_files1 = C.int32_t(num_files0)
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_enumerator_next_files_async(this1, num_files1, io_priority1, cancellable1, callback1)
}
func (this0 *FileEnumerator) NextFilesFinish(result0 AsyncResultLike) ([]*FileInfo, error) {
	var this1 *C.GFileEnumerator
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GFileEnumerator)(this0.InheritedFromGFileEnumerator())
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_enumerator_next_files_finish(this1, result1, &err1)
	var ret2 []*FileInfo
	var err2 error
	for iter := (*_GList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt *FileInfo
		elt = (*FileInfo)(gobject.ObjectWrap(unsafe.Pointer((*C.GFileInfo)(iter.data)), false))
		ret2 = append(ret2, elt)
	}
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileEnumerator) SetPending(pending0 bool) {
	var this1 *C.GFileEnumerator
	var pending1 C.int
	if this0 != nil {
		this1 = (*C.GFileEnumerator)(this0.InheritedFromGFileEnumerator())
	}
	pending1 = _GoBoolToCBool(pending0)
	C.g_file_enumerator_set_pending(this1, pending1)
}
// blacklisted: FileEnumeratorClass (struct)
// blacklisted: FileEnumeratorPrivate (struct)
type FileIOStreamLike interface {
	IOStreamLike
	InheritedFromGFileIOStream() *C.GFileIOStream
}

type FileIOStream struct {
	IOStream
	SeekableImpl
}

func ToFileIOStream(objlike gobject.ObjectLike) *FileIOStream {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*FileIOStream)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*FileIOStream)(obj)
	}
	panic("cannot cast to FileIOStream")
}

func (this0 *FileIOStream) InheritedFromGFileIOStream() *C.GFileIOStream {
	if this0 == nil {
		return nil
	}
	return (*C.GFileIOStream)(this0.C)
}

func (this0 *FileIOStream) GetStaticType() gobject.Type {
	return gobject.Type(C.g_file_io_stream_get_type())
}

func FileIOStreamGetType() gobject.Type {
	return (*FileIOStream)(nil).GetStaticType()
}
func (this0 *FileIOStream) GetEtag() string {
	var this1 *C.GFileIOStream
	if this0 != nil {
		this1 = (*C.GFileIOStream)(this0.InheritedFromGFileIOStream())
	}
	ret1 := C.g_file_io_stream_get_etag(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *FileIOStream) QueryInfo(attributes0 string, cancellable0 CancellableLike) (*FileInfo, error) {
	var this1 *C.GFileIOStream
	var attributes1 *C.char
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GFileIOStream)(this0.InheritedFromGFileIOStream())
	}
	attributes1 = _GoStringToGString(attributes0)
	defer C.free(unsafe.Pointer(attributes1))
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_io_stream_query_info(this1, attributes1, cancellable1, &err1)
	var ret2 *FileInfo
	var err2 error
	ret2 = (*FileInfo)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileIOStream) QueryInfoAsync(attributes0 string, io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFileIOStream
	var attributes1 *C.char
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = (*C.GFileIOStream)(this0.InheritedFromGFileIOStream())
	}
	attributes1 = _GoStringToGString(attributes0)
	defer C.free(unsafe.Pointer(attributes1))
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_io_stream_query_info_async(this1, attributes1, io_priority1, cancellable1, callback1)
}
func (this0 *FileIOStream) QueryInfoFinish(result0 AsyncResultLike) (*FileInfo, error) {
	var this1 *C.GFileIOStream
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GFileIOStream)(this0.InheritedFromGFileIOStream())
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_io_stream_query_info_finish(this1, result1, &err1)
	var ret2 *FileInfo
	var err2 error
	ret2 = (*FileInfo)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
// blacklisted: FileIOStreamClass (struct)
// blacklisted: FileIOStreamPrivate (struct)
// blacklisted: FileIcon (object)
// blacklisted: FileIconClass (struct)
// blacklisted: FileIface (struct)
type FileInfoLike interface {
	gobject.ObjectLike
	InheritedFromGFileInfo() *C.GFileInfo
}

type FileInfo struct {
	gobject.Object
	
}

func ToFileInfo(objlike gobject.ObjectLike) *FileInfo {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*FileInfo)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*FileInfo)(obj)
	}
	panic("cannot cast to FileInfo")
}

func (this0 *FileInfo) InheritedFromGFileInfo() *C.GFileInfo {
	if this0 == nil {
		return nil
	}
	return (*C.GFileInfo)(this0.C)
}

func (this0 *FileInfo) GetStaticType() gobject.Type {
	return gobject.Type(C.g_file_info_get_type())
}

func FileInfoGetType() gobject.Type {
	return (*FileInfo)(nil).GetStaticType()
}
func NewFileInfo() *FileInfo {
	ret1 := C.g_file_info_new()
	var ret2 *FileInfo
	ret2 = (*FileInfo)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *FileInfo) ClearStatus() {
	var this1 *C.GFileInfo
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	C.g_file_info_clear_status(this1)
}
func (this0 *FileInfo) CopyInto(dest_info0 FileInfoLike) {
	var this1 *C.GFileInfo
	var dest_info1 *C.GFileInfo
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	if dest_info0 != nil {
		dest_info1 = (*C.GFileInfo)(dest_info0.InheritedFromGFileInfo())
	}
	C.g_file_info_copy_into(this1, dest_info1)
}
func (this0 *FileInfo) Dup() *FileInfo {
	var this1 *C.GFileInfo
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	ret1 := C.g_file_info_dup(this1)
	var ret2 *FileInfo
	ret2 = (*FileInfo)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *FileInfo) GetAttributeAsString(attribute0 string) string {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	ret1 := C.g_file_info_get_attribute_as_string(this1, attribute1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *FileInfo) GetAttributeBoolean(attribute0 string) bool {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	ret1 := C.g_file_info_get_attribute_boolean(this1, attribute1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *FileInfo) GetAttributeByteString(attribute0 string) string {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	ret1 := C.g_file_info_get_attribute_byte_string(this1, attribute1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *FileInfo) GetAttributeData(attribute0 string) (FileAttributeType, unsafe.Pointer, FileAttributeStatus, bool) {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	var type1 C.GFileAttributeType
	var value_pp1 unsafe.Pointer
	var status1 C.GFileAttributeStatus
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	ret1 := C.g_file_info_get_attribute_data(this1, attribute1, &type1, &value_pp1, &status1)
	var type2 FileAttributeType
	var value_pp2 unsafe.Pointer
	var status2 FileAttributeStatus
	var ret2 bool
	type2 = FileAttributeType(type1)
	value_pp2 = value_pp1
	status2 = FileAttributeStatus(status1)
	ret2 = ret1 != 0
	return type2, value_pp2, status2, ret2
}
func (this0 *FileInfo) GetAttributeInt32(attribute0 string) int32 {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	ret1 := C.g_file_info_get_attribute_int32(this1, attribute1)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func (this0 *FileInfo) GetAttributeInt64(attribute0 string) int64 {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	ret1 := C.g_file_info_get_attribute_int64(this1, attribute1)
	var ret2 int64
	ret2 = int64(ret1)
	return ret2
}
func (this0 *FileInfo) GetAttributeObject(attribute0 string) *gobject.Object {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	ret1 := C.g_file_info_get_attribute_object(this1, attribute1)
	var ret2 *gobject.Object
	ret2 = (*gobject.Object)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *FileInfo) GetAttributeStatus(attribute0 string) FileAttributeStatus {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	ret1 := C.g_file_info_get_attribute_status(this1, attribute1)
	var ret2 FileAttributeStatus
	ret2 = FileAttributeStatus(ret1)
	return ret2
}
func (this0 *FileInfo) GetAttributeString(attribute0 string) string {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	ret1 := C.g_file_info_get_attribute_string(this1, attribute1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *FileInfo) GetAttributeStringv(attribute0 string) []string {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	ret1 := C.g_file_info_get_attribute_stringv(this1, attribute1)
	var ret2 []string
	ret2 = make([]string, uint(C._array_length(unsafe.Pointer(ret1))))
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
	}
	return ret2
}
func (this0 *FileInfo) GetAttributeType(attribute0 string) FileAttributeType {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	ret1 := C.g_file_info_get_attribute_type(this1, attribute1)
	var ret2 FileAttributeType
	ret2 = FileAttributeType(ret1)
	return ret2
}
func (this0 *FileInfo) GetAttributeUint32(attribute0 string) uint32 {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	ret1 := C.g_file_info_get_attribute_uint32(this1, attribute1)
	var ret2 uint32
	ret2 = uint32(ret1)
	return ret2
}
func (this0 *FileInfo) GetAttributeUint64(attribute0 string) uint64 {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	ret1 := C.g_file_info_get_attribute_uint64(this1, attribute1)
	var ret2 uint64
	ret2 = uint64(ret1)
	return ret2
}
func (this0 *FileInfo) GetContentType() string {
	var this1 *C.GFileInfo
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	ret1 := C.g_file_info_get_content_type(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *FileInfo) GetDeletionDate() *glib.DateTime {
	var this1 *C.GFileInfo
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	ret1 := C.g_file_info_get_deletion_date(this1)
	var ret2 *glib.DateTime
	ret2 = (*glib.DateTime)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *FileInfo) GetDisplayName() string {
	var this1 *C.GFileInfo
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	ret1 := C.g_file_info_get_display_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *FileInfo) GetEditName() string {
	var this1 *C.GFileInfo
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	ret1 := C.g_file_info_get_edit_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *FileInfo) GetEtag() string {
	var this1 *C.GFileInfo
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	ret1 := C.g_file_info_get_etag(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *FileInfo) GetFileType() FileType {
	var this1 *C.GFileInfo
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	ret1 := C.g_file_info_get_file_type(this1)
	var ret2 FileType
	ret2 = FileType(ret1)
	return ret2
}
func (this0 *FileInfo) GetIcon() *Icon {
	var this1 *C.GFileInfo
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	ret1 := C.g_file_info_get_icon(this1)
	var ret2 *Icon
	ret2 = (*Icon)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *FileInfo) GetIsBackup() bool {
	var this1 *C.GFileInfo
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	ret1 := C.g_file_info_get_is_backup(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *FileInfo) GetIsHidden() bool {
	var this1 *C.GFileInfo
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	ret1 := C.g_file_info_get_is_hidden(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *FileInfo) GetIsSymlink() bool {
	var this1 *C.GFileInfo
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	ret1 := C.g_file_info_get_is_symlink(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *FileInfo) GetModificationTime() glib.TimeVal {
	var this1 *C.GFileInfo
	var result1 C.GTimeVal
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	C.g_file_info_get_modification_time(this1, &result1)
	var result2 glib.TimeVal
	result2 = *(*glib.TimeVal)(unsafe.Pointer(&result1))
	return result2
}
func (this0 *FileInfo) GetName() string {
	var this1 *C.GFileInfo
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	ret1 := C.g_file_info_get_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *FileInfo) GetSize() int64 {
	var this1 *C.GFileInfo
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	ret1 := C.g_file_info_get_size(this1)
	var ret2 int64
	ret2 = int64(ret1)
	return ret2
}
func (this0 *FileInfo) GetSortOrder() int32 {
	var this1 *C.GFileInfo
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	ret1 := C.g_file_info_get_sort_order(this1)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func (this0 *FileInfo) GetSymbolicIcon() *Icon {
	var this1 *C.GFileInfo
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	ret1 := C.g_file_info_get_symbolic_icon(this1)
	var ret2 *Icon
	ret2 = (*Icon)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *FileInfo) GetSymlinkTarget() string {
	var this1 *C.GFileInfo
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	ret1 := C.g_file_info_get_symlink_target(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *FileInfo) HasAttribute(attribute0 string) bool {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	ret1 := C.g_file_info_has_attribute(this1, attribute1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *FileInfo) HasNamespace(name_space0 string) bool {
	var this1 *C.GFileInfo
	var name_space1 *C.char
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	name_space1 = _GoStringToGString(name_space0)
	defer C.free(unsafe.Pointer(name_space1))
	ret1 := C.g_file_info_has_namespace(this1, name_space1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *FileInfo) ListAttributes(name_space0 string) []string {
	var this1 *C.GFileInfo
	var name_space1 *C.char
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	name_space1 = _GoStringToGString(name_space0)
	defer C.free(unsafe.Pointer(name_space1))
	ret1 := C.g_file_info_list_attributes(this1, name_space1)
	var ret2 []string
	ret2 = make([]string, uint(C._array_length(unsafe.Pointer(ret1))))
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return ret2
}
func (this0 *FileInfo) RemoveAttribute(attribute0 string) {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	C.g_file_info_remove_attribute(this1, attribute1)
}
func (this0 *FileInfo) SetAttribute(attribute0 string, type0 FileAttributeType, value_p0 unsafe.Pointer) {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	var type1 C.GFileAttributeType
	var value_p1 unsafe.Pointer
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	type1 = C.GFileAttributeType(type0)
	value_p1 = unsafe.Pointer(value_p0)
	C.g_file_info_set_attribute(this1, attribute1, type1, value_p1)
}
func (this0 *FileInfo) SetAttributeBoolean(attribute0 string, attr_value0 bool) {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	var attr_value1 C.int
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	attr_value1 = _GoBoolToCBool(attr_value0)
	C.g_file_info_set_attribute_boolean(this1, attribute1, attr_value1)
}
func (this0 *FileInfo) SetAttributeByteString(attribute0 string, attr_value0 string) {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	var attr_value1 *C.char
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	attr_value1 = _GoStringToGString(attr_value0)
	defer C.free(unsafe.Pointer(attr_value1))
	C.g_file_info_set_attribute_byte_string(this1, attribute1, attr_value1)
}
func (this0 *FileInfo) SetAttributeInt32(attribute0 string, attr_value0 int32) {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	var attr_value1 C.int32_t
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	attr_value1 = C.int32_t(attr_value0)
	C.g_file_info_set_attribute_int32(this1, attribute1, attr_value1)
}
func (this0 *FileInfo) SetAttributeInt64(attribute0 string, attr_value0 int64) {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	var attr_value1 C.int64_t
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	attr_value1 = C.int64_t(attr_value0)
	C.g_file_info_set_attribute_int64(this1, attribute1, attr_value1)
}
func (this0 *FileInfo) SetAttributeMask(mask0 *FileAttributeMatcher) {
	var this1 *C.GFileInfo
	var mask1 *C.GFileAttributeMatcher
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	mask1 = (*C.GFileAttributeMatcher)(unsafe.Pointer(mask0))
	C.g_file_info_set_attribute_mask(this1, mask1)
}
func (this0 *FileInfo) SetAttributeObject(attribute0 string, attr_value0 gobject.ObjectLike) {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	var attr_value1 *C.GObject
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	if attr_value0 != nil {
		attr_value1 = (*C.GObject)(attr_value0.InheritedFromGObject())
	}
	C.g_file_info_set_attribute_object(this1, attribute1, attr_value1)
}
func (this0 *FileInfo) SetAttributeStatus(attribute0 string, status0 FileAttributeStatus) bool {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	var status1 C.GFileAttributeStatus
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	status1 = C.GFileAttributeStatus(status0)
	ret1 := C.g_file_info_set_attribute_status(this1, attribute1, status1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *FileInfo) SetAttributeString(attribute0 string, attr_value0 string) {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	var attr_value1 *C.char
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	attr_value1 = _GoStringToGString(attr_value0)
	defer C.free(unsafe.Pointer(attr_value1))
	C.g_file_info_set_attribute_string(this1, attribute1, attr_value1)
}
func (this0 *FileInfo) SetAttributeStringv(attribute0 string, attr_value0 []string) {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	var attr_value1 **C.char
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	attr_value1 = (**C.char)(C.malloc(C.size_t(int(unsafe.Sizeof(*attr_value1)) * len(attr_value0))))
	defer C.free(unsafe.Pointer(attr_value1))
	for i, e := range attr_value0 {
		(*(*[999999]*C.char)(unsafe.Pointer(attr_value1)))[i] = _GoStringToGString(e)
		defer C.free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(attr_value1)))[i]))
	}
	C.g_file_info_set_attribute_stringv(this1, attribute1, attr_value1)
}
func (this0 *FileInfo) SetAttributeUint32(attribute0 string, attr_value0 uint32) {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	var attr_value1 C.uint32_t
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	attr_value1 = C.uint32_t(attr_value0)
	C.g_file_info_set_attribute_uint32(this1, attribute1, attr_value1)
}
func (this0 *FileInfo) SetAttributeUint64(attribute0 string, attr_value0 uint64) {
	var this1 *C.GFileInfo
	var attribute1 *C.char
	var attr_value1 C.uint64_t
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	attribute1 = _GoStringToGString(attribute0)
	defer C.free(unsafe.Pointer(attribute1))
	attr_value1 = C.uint64_t(attr_value0)
	C.g_file_info_set_attribute_uint64(this1, attribute1, attr_value1)
}
func (this0 *FileInfo) SetContentType(content_type0 string) {
	var this1 *C.GFileInfo
	var content_type1 *C.char
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	content_type1 = _GoStringToGString(content_type0)
	defer C.free(unsafe.Pointer(content_type1))
	C.g_file_info_set_content_type(this1, content_type1)
}
func (this0 *FileInfo) SetDisplayName(display_name0 string) {
	var this1 *C.GFileInfo
	var display_name1 *C.char
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	display_name1 = _GoStringToGString(display_name0)
	defer C.free(unsafe.Pointer(display_name1))
	C.g_file_info_set_display_name(this1, display_name1)
}
func (this0 *FileInfo) SetEditName(edit_name0 string) {
	var this1 *C.GFileInfo
	var edit_name1 *C.char
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	edit_name1 = _GoStringToGString(edit_name0)
	defer C.free(unsafe.Pointer(edit_name1))
	C.g_file_info_set_edit_name(this1, edit_name1)
}
func (this0 *FileInfo) SetFileType(type0 FileType) {
	var this1 *C.GFileInfo
	var type1 C.GFileType
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	type1 = C.GFileType(type0)
	C.g_file_info_set_file_type(this1, type1)
}
func (this0 *FileInfo) SetIcon(icon0 IconLike) {
	var this1 *C.GFileInfo
	var icon1 *C.GIcon
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	if icon0 != nil {
		icon1 = icon0.ImplementsGIcon()
	}
	C.g_file_info_set_icon(this1, icon1)
}
func (this0 *FileInfo) SetIsHidden(is_hidden0 bool) {
	var this1 *C.GFileInfo
	var is_hidden1 C.int
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	is_hidden1 = _GoBoolToCBool(is_hidden0)
	C.g_file_info_set_is_hidden(this1, is_hidden1)
}
func (this0 *FileInfo) SetIsSymlink(is_symlink0 bool) {
	var this1 *C.GFileInfo
	var is_symlink1 C.int
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	is_symlink1 = _GoBoolToCBool(is_symlink0)
	C.g_file_info_set_is_symlink(this1, is_symlink1)
}
func (this0 *FileInfo) SetModificationTime(mtime0 *glib.TimeVal) {
	var this1 *C.GFileInfo
	var mtime1 *C.GTimeVal
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	mtime1 = (*C.GTimeVal)(unsafe.Pointer(mtime0))
	C.g_file_info_set_modification_time(this1, mtime1)
}
func (this0 *FileInfo) SetName(name0 string) {
	var this1 *C.GFileInfo
	var name1 *C.char
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	C.g_file_info_set_name(this1, name1)
}
func (this0 *FileInfo) SetSize(size0 int64) {
	var this1 *C.GFileInfo
	var size1 C.int64_t
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	size1 = C.int64_t(size0)
	C.g_file_info_set_size(this1, size1)
}
func (this0 *FileInfo) SetSortOrder(sort_order0 int32) {
	var this1 *C.GFileInfo
	var sort_order1 C.int32_t
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	sort_order1 = C.int32_t(sort_order0)
	C.g_file_info_set_sort_order(this1, sort_order1)
}
func (this0 *FileInfo) SetSymbolicIcon(icon0 IconLike) {
	var this1 *C.GFileInfo
	var icon1 *C.GIcon
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	if icon0 != nil {
		icon1 = icon0.ImplementsGIcon()
	}
	C.g_file_info_set_symbolic_icon(this1, icon1)
}
func (this0 *FileInfo) SetSymlinkTarget(symlink_target0 string) {
	var this1 *C.GFileInfo
	var symlink_target1 *C.char
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	symlink_target1 = _GoStringToGString(symlink_target0)
	defer C.free(unsafe.Pointer(symlink_target1))
	C.g_file_info_set_symlink_target(this1, symlink_target1)
}
func (this0 *FileInfo) UnsetAttributeMask() {
	var this1 *C.GFileInfo
	if this0 != nil {
		this1 = (*C.GFileInfo)(this0.InheritedFromGFileInfo())
	}
	C.g_file_info_unset_attribute_mask(this1)
}
// blacklisted: FileInfoClass (struct)
type FileInputStreamLike interface {
	InputStreamLike
	InheritedFromGFileInputStream() *C.GFileInputStream
}

type FileInputStream struct {
	InputStream
	SeekableImpl
}

func ToFileInputStream(objlike gobject.ObjectLike) *FileInputStream {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*FileInputStream)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*FileInputStream)(obj)
	}
	panic("cannot cast to FileInputStream")
}

func (this0 *FileInputStream) InheritedFromGFileInputStream() *C.GFileInputStream {
	if this0 == nil {
		return nil
	}
	return (*C.GFileInputStream)(this0.C)
}

func (this0 *FileInputStream) GetStaticType() gobject.Type {
	return gobject.Type(C.g_file_input_stream_get_type())
}

func FileInputStreamGetType() gobject.Type {
	return (*FileInputStream)(nil).GetStaticType()
}
func (this0 *FileInputStream) QueryInfo(attributes0 string, cancellable0 CancellableLike) (*FileInfo, error) {
	var this1 *C.GFileInputStream
	var attributes1 *C.char
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GFileInputStream)(this0.InheritedFromGFileInputStream())
	}
	attributes1 = _GoStringToGString(attributes0)
	defer C.free(unsafe.Pointer(attributes1))
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_input_stream_query_info(this1, attributes1, cancellable1, &err1)
	var ret2 *FileInfo
	var err2 error
	ret2 = (*FileInfo)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileInputStream) QueryInfoAsync(attributes0 string, io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFileInputStream
	var attributes1 *C.char
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = (*C.GFileInputStream)(this0.InheritedFromGFileInputStream())
	}
	attributes1 = _GoStringToGString(attributes0)
	defer C.free(unsafe.Pointer(attributes1))
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_input_stream_query_info_async(this1, attributes1, io_priority1, cancellable1, callback1)
}
func (this0 *FileInputStream) QueryInfoFinish(result0 AsyncResultLike) (*FileInfo, error) {
	var this1 *C.GFileInputStream
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GFileInputStream)(this0.InheritedFromGFileInputStream())
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_input_stream_query_info_finish(this1, result1, &err1)
	var ret2 *FileInfo
	var err2 error
	ret2 = (*FileInfo)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
// blacklisted: FileInputStreamClass (struct)
// blacklisted: FileInputStreamPrivate (struct)
type FileMeasureFlags C.uint32_t
const (
	FileMeasureFlagsNone FileMeasureFlags = 0
	FileMeasureFlagsReportAnyError FileMeasureFlags = 2
	FileMeasureFlagsApparentSize FileMeasureFlags = 4
	FileMeasureFlagsNoXdev FileMeasureFlags = 8
)
// blacklisted: FileMeasureProgressCallback (callback)
type FileMonitorLike interface {
	gobject.ObjectLike
	InheritedFromGFileMonitor() *C.GFileMonitor
}

type FileMonitor struct {
	gobject.Object
	
}

func ToFileMonitor(objlike gobject.ObjectLike) *FileMonitor {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*FileMonitor)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*FileMonitor)(obj)
	}
	panic("cannot cast to FileMonitor")
}

func (this0 *FileMonitor) InheritedFromGFileMonitor() *C.GFileMonitor {
	if this0 == nil {
		return nil
	}
	return (*C.GFileMonitor)(this0.C)
}

func (this0 *FileMonitor) GetStaticType() gobject.Type {
	return gobject.Type(C.g_file_monitor_get_type())
}

func FileMonitorGetType() gobject.Type {
	return (*FileMonitor)(nil).GetStaticType()
}
func (this0 *FileMonitor) Cancel() bool {
	var this1 *C.GFileMonitor
	if this0 != nil {
		this1 = (*C.GFileMonitor)(this0.InheritedFromGFileMonitor())
	}
	ret1 := C.g_file_monitor_cancel(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *FileMonitor) EmitEvent(child0 FileLike, other_file0 FileLike, event_type0 FileMonitorEvent) {
	var this1 *C.GFileMonitor
	var child1 *C.GFile
	var other_file1 *C.GFile
	var event_type1 C.GFileMonitorEvent
	if this0 != nil {
		this1 = (*C.GFileMonitor)(this0.InheritedFromGFileMonitor())
	}
	if child0 != nil {
		child1 = child0.ImplementsGFile()
	}
	if other_file0 != nil {
		other_file1 = other_file0.ImplementsGFile()
	}
	event_type1 = C.GFileMonitorEvent(event_type0)
	C.g_file_monitor_emit_event(this1, child1, other_file1, event_type1)
}
func (this0 *FileMonitor) IsCancelled() bool {
	var this1 *C.GFileMonitor
	if this0 != nil {
		this1 = (*C.GFileMonitor)(this0.InheritedFromGFileMonitor())
	}
	ret1 := C.g_file_monitor_is_cancelled(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *FileMonitor) SetRateLimit(limit_msecs0 int32) {
	var this1 *C.GFileMonitor
	var limit_msecs1 C.int32_t
	if this0 != nil {
		this1 = (*C.GFileMonitor)(this0.InheritedFromGFileMonitor())
	}
	limit_msecs1 = C.int32_t(limit_msecs0)
	C.g_file_monitor_set_rate_limit(this1, limit_msecs1)
}
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
type FileOutputStreamLike interface {
	OutputStreamLike
	InheritedFromGFileOutputStream() *C.GFileOutputStream
}

type FileOutputStream struct {
	OutputStream
	SeekableImpl
}

func ToFileOutputStream(objlike gobject.ObjectLike) *FileOutputStream {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*FileOutputStream)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*FileOutputStream)(obj)
	}
	panic("cannot cast to FileOutputStream")
}

func (this0 *FileOutputStream) InheritedFromGFileOutputStream() *C.GFileOutputStream {
	if this0 == nil {
		return nil
	}
	return (*C.GFileOutputStream)(this0.C)
}

func (this0 *FileOutputStream) GetStaticType() gobject.Type {
	return gobject.Type(C.g_file_output_stream_get_type())
}

func FileOutputStreamGetType() gobject.Type {
	return (*FileOutputStream)(nil).GetStaticType()
}
func (this0 *FileOutputStream) GetEtag() string {
	var this1 *C.GFileOutputStream
	if this0 != nil {
		this1 = (*C.GFileOutputStream)(this0.InheritedFromGFileOutputStream())
	}
	ret1 := C.g_file_output_stream_get_etag(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *FileOutputStream) QueryInfo(attributes0 string, cancellable0 CancellableLike) (*FileInfo, error) {
	var this1 *C.GFileOutputStream
	var attributes1 *C.char
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GFileOutputStream)(this0.InheritedFromGFileOutputStream())
	}
	attributes1 = _GoStringToGString(attributes0)
	defer C.free(unsafe.Pointer(attributes1))
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_file_output_stream_query_info(this1, attributes1, cancellable1, &err1)
	var ret2 *FileInfo
	var err2 error
	ret2 = (*FileInfo)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *FileOutputStream) QueryInfoAsync(attributes0 string, io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GFileOutputStream
	var attributes1 *C.char
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = (*C.GFileOutputStream)(this0.InheritedFromGFileOutputStream())
	}
	attributes1 = _GoStringToGString(attributes0)
	defer C.free(unsafe.Pointer(attributes1))
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_file_output_stream_query_info_async(this1, attributes1, io_priority1, cancellable1, callback1)
}
func (this0 *FileOutputStream) QueryInfoFinish(result0 AsyncResultLike) (*FileInfo, error) {
	var this1 *C.GFileOutputStream
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GFileOutputStream)(this0.InheritedFromGFileOutputStream())
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_file_output_stream_query_info_finish(this1, result1, &err1)
	var ret2 *FileInfo
	var err2 error
	ret2 = (*FileInfo)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
// blacklisted: FileOutputStreamClass (struct)
// blacklisted: FileOutputStreamPrivate (struct)
type FileProgressCallback func(current_num_bytes int64, total_num_bytes int64)
//export _GFileProgressCallback_c_wrapper
func _GFileProgressCallback_c_wrapper(current_num_bytes0 int64, total_num_bytes0 int64, user_data0 unsafe.Pointer) {
	var current_num_bytes1 int64
	var total_num_bytes1 int64
	var user_data1 FileProgressCallback
	current_num_bytes1 = int64((C.int64_t)(current_num_bytes0))
	total_num_bytes1 = int64((C.int64_t)(total_num_bytes0))
	user_data1 = *(*FileProgressCallback)(user_data0)
	user_data1(current_num_bytes1, total_num_bytes1)
}
//export _GFileProgressCallback_c_wrapper_once
func _GFileProgressCallback_c_wrapper_once(current_num_bytes0 int64, total_num_bytes0 int64, user_data0 unsafe.Pointer) {
	_GFileProgressCallback_c_wrapper(current_num_bytes0, total_num_bytes0, user_data0)
	gobject.Holder.Release(user_data0)
}
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
type IOStreamLike interface {
	gobject.ObjectLike
	InheritedFromGIOStream() *C.GIOStream
}

type IOStream struct {
	gobject.Object
	
}

func ToIOStream(objlike gobject.ObjectLike) *IOStream {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*IOStream)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*IOStream)(obj)
	}
	panic("cannot cast to IOStream")
}

func (this0 *IOStream) InheritedFromGIOStream() *C.GIOStream {
	if this0 == nil {
		return nil
	}
	return (*C.GIOStream)(this0.C)
}

func (this0 *IOStream) GetStaticType() gobject.Type {
	return gobject.Type(C.g_io_stream_get_type())
}

func IOStreamGetType() gobject.Type {
	return (*IOStream)(nil).GetStaticType()
}
func IOStreamSpliceFinish(result0 AsyncResultLike) (bool, error) {
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_io_stream_splice_finish(result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *IOStream) ClearPending() {
	var this1 *C.GIOStream
	if this0 != nil {
		this1 = (*C.GIOStream)(this0.InheritedFromGIOStream())
	}
	C.g_io_stream_clear_pending(this1)
}
func (this0 *IOStream) Close(cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GIOStream
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GIOStream)(this0.InheritedFromGIOStream())
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_io_stream_close(this1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *IOStream) CloseAsync(io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GIOStream
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = (*C.GIOStream)(this0.InheritedFromGIOStream())
	}
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_io_stream_close_async(this1, io_priority1, cancellable1, callback1)
}
func (this0 *IOStream) CloseFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GIOStream
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GIOStream)(this0.InheritedFromGIOStream())
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_io_stream_close_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *IOStream) GetInputStream() *InputStream {
	var this1 *C.GIOStream
	if this0 != nil {
		this1 = (*C.GIOStream)(this0.InheritedFromGIOStream())
	}
	ret1 := C.g_io_stream_get_input_stream(this1)
	var ret2 *InputStream
	ret2 = (*InputStream)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *IOStream) GetOutputStream() *OutputStream {
	var this1 *C.GIOStream
	if this0 != nil {
		this1 = (*C.GIOStream)(this0.InheritedFromGIOStream())
	}
	ret1 := C.g_io_stream_get_output_stream(this1)
	var ret2 *OutputStream
	ret2 = (*OutputStream)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *IOStream) HasPending() bool {
	var this1 *C.GIOStream
	if this0 != nil {
		this1 = (*C.GIOStream)(this0.InheritedFromGIOStream())
	}
	ret1 := C.g_io_stream_has_pending(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *IOStream) IsClosed() bool {
	var this1 *C.GIOStream
	if this0 != nil {
		this1 = (*C.GIOStream)(this0.InheritedFromGIOStream())
	}
	ret1 := C.g_io_stream_is_closed(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *IOStream) SetPending() (bool, error) {
	var this1 *C.GIOStream
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GIOStream)(this0.InheritedFromGIOStream())
	}
	ret1 := C.g_io_stream_set_pending(this1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *IOStream) SpliceAsync(stream20 IOStreamLike, flags0 IOStreamSpliceFlags, io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GIOStream
	var stream21 *C.GIOStream
	var flags1 C.GIOStreamSpliceFlags
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = (*C.GIOStream)(this0.InheritedFromGIOStream())
	}
	if stream20 != nil {
		stream21 = (*C.GIOStream)(stream20.InheritedFromGIOStream())
	}
	flags1 = C.GIOStreamSpliceFlags(flags0)
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_io_stream_splice_async(this1, stream21, flags1, io_priority1, cancellable1, callback1)
}
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

func (*Icon) GetStaticType() gobject.Type {
	return gobject.Type(C.g_icon_get_type())
}


type IconImpl struct {}

func ToIcon(objlike gobject.ObjectLike) *Icon {
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), gobject.Type(C.g_icon_get_type()))
	if obj != nil {
		return (*Icon)(obj)
	}
	panic("cannot cast to Icon")
}

func (this0 *IconImpl) ImplementsGIcon() *C.GIcon {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GIcon)((*gobject.Object)(unsafe.Pointer(obj)).C)
}
func IconDeserialize(value0 *glib.Variant) *Icon {
	var value1 *C.GVariant
	value1 = (*C.GVariant)(unsafe.Pointer(value0))
	ret1 := C.g_icon_deserialize(value1)
	var ret2 *Icon
	ret2 = (*Icon)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func IconHash(icon0 unsafe.Pointer) uint32 {
	var icon1 unsafe.Pointer
	icon1 = unsafe.Pointer(icon0)
	ret1 := C.g_icon_hash(icon1)
	var ret2 uint32
	ret2 = uint32(ret1)
	return ret2
}
func IconNewForString(str0 string) (*Icon, error) {
	var str1 *C.char
	var err1 *C.GError
	str1 = _GoStringToGString(str0)
	defer C.free(unsafe.Pointer(str1))
	ret1 := C.g_icon_new_for_string(str1, &err1)
	var ret2 *Icon
	var err2 error
	ret2 = (*Icon)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *IconImpl) Equal(icon20 IconLike) bool {
	var this1 *C.GIcon
	var icon21 *C.GIcon
	if this0 != nil {
		this1 = this0.ImplementsGIcon()
	}
	if icon20 != nil {
		icon21 = icon20.ImplementsGIcon()
	}
	ret1 := C.g_icon_equal(this1, icon21)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *IconImpl) Serialize() *glib.Variant {
	var this1 *C.GIcon
	if this0 != nil {
		this1 = this0.ImplementsGIcon()
	}
	ret1 := C.g_icon_serialize(this1)
	var ret2 *glib.Variant
	ret2 = (*glib.Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *IconImpl) ToString() string {
	var this1 *C.GIcon
	if this0 != nil {
		this1 = this0.ImplementsGIcon()
	}
	ret1 := C.g_icon_to_string(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
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
type InputStreamLike interface {
	gobject.ObjectLike
	InheritedFromGInputStream() *C.GInputStream
}

type InputStream struct {
	gobject.Object
	
}

func ToInputStream(objlike gobject.ObjectLike) *InputStream {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*InputStream)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*InputStream)(obj)
	}
	panic("cannot cast to InputStream")
}

func (this0 *InputStream) InheritedFromGInputStream() *C.GInputStream {
	if this0 == nil {
		return nil
	}
	return (*C.GInputStream)(this0.C)
}

func (this0 *InputStream) GetStaticType() gobject.Type {
	return gobject.Type(C.g_input_stream_get_type())
}

func InputStreamGetType() gobject.Type {
	return (*InputStream)(nil).GetStaticType()
}
func (this0 *InputStream) ClearPending() {
	var this1 *C.GInputStream
	if this0 != nil {
		this1 = (*C.GInputStream)(this0.InheritedFromGInputStream())
	}
	C.g_input_stream_clear_pending(this1)
}
func (this0 *InputStream) Close(cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GInputStream
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GInputStream)(this0.InheritedFromGInputStream())
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_input_stream_close(this1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *InputStream) CloseAsync(io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GInputStream
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = (*C.GInputStream)(this0.InheritedFromGInputStream())
	}
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_input_stream_close_async(this1, io_priority1, cancellable1, callback1)
}
func (this0 *InputStream) CloseFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GInputStream
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GInputStream)(this0.InheritedFromGInputStream())
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_input_stream_close_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *InputStream) HasPending() bool {
	var this1 *C.GInputStream
	if this0 != nil {
		this1 = (*C.GInputStream)(this0.InheritedFromGInputStream())
	}
	ret1 := C.g_input_stream_has_pending(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *InputStream) IsClosed() bool {
	var this1 *C.GInputStream
	if this0 != nil {
		this1 = (*C.GInputStream)(this0.InheritedFromGInputStream())
	}
	ret1 := C.g_input_stream_is_closed(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *InputStream) Read(buffer0 []uint8, cancellable0 CancellableLike) (int64, error) {
	var this1 *C.GInputStream
	var buffer1 *C.uint8_t
	var count1 C.uint64_t
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GInputStream)(this0.InheritedFromGInputStream())
	}
	buffer1 = (*C.uint8_t)(C.malloc(C.size_t(int(unsafe.Sizeof(*buffer1)) * len(buffer0))))
	defer C.free(unsafe.Pointer(buffer1))
	for i, e := range buffer0 {
		(*(*[999999]C.uint8_t)(unsafe.Pointer(buffer1)))[i] = C.uint8_t(e)
	}
	count1 = C.uint64_t(len(buffer0))
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_input_stream_read(this1, buffer1, count1, cancellable1, &err1)
	var ret2 int64
	var err2 error
	ret2 = int64(ret1)
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *InputStream) ReadAll(buffer0 []uint8, cancellable0 CancellableLike) (uint64, bool, error) {
	var this1 *C.GInputStream
	var buffer1 *C.uint8_t
	var count1 C.uint64_t
	var cancellable1 *C.GCancellable
	var bytes_read1 C.uint64_t
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GInputStream)(this0.InheritedFromGInputStream())
	}
	buffer1 = (*C.uint8_t)(C.malloc(C.size_t(int(unsafe.Sizeof(*buffer1)) * len(buffer0))))
	defer C.free(unsafe.Pointer(buffer1))
	for i, e := range buffer0 {
		(*(*[999999]C.uint8_t)(unsafe.Pointer(buffer1)))[i] = C.uint8_t(e)
	}
	count1 = C.uint64_t(len(buffer0))
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_input_stream_read_all(this1, buffer1, count1, &bytes_read1, cancellable1, &err1)
	var bytes_read2 uint64
	var ret2 bool
	var err2 error
	bytes_read2 = uint64(bytes_read1)
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return bytes_read2, ret2, err2
}
func (this0 *InputStream) ReadAsync(buffer0 []uint8, io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GInputStream
	var buffer1 *C.uint8_t
	var count1 C.uint64_t
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = (*C.GInputStream)(this0.InheritedFromGInputStream())
	}
	buffer1 = (*C.uint8_t)(C.malloc(C.size_t(int(unsafe.Sizeof(*buffer1)) * len(buffer0))))
	defer C.free(unsafe.Pointer(buffer1))
	for i, e := range buffer0 {
		(*(*[999999]C.uint8_t)(unsafe.Pointer(buffer1)))[i] = C.uint8_t(e)
	}
	count1 = C.uint64_t(len(buffer0))
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_input_stream_read_async(this1, buffer1, count1, io_priority1, cancellable1, callback1)
}
func (this0 *InputStream) ReadBytes(count0 uint64, cancellable0 CancellableLike) (*glib.Bytes, error) {
	var this1 *C.GInputStream
	var count1 C.uint64_t
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GInputStream)(this0.InheritedFromGInputStream())
	}
	count1 = C.uint64_t(count0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_input_stream_read_bytes(this1, count1, cancellable1, &err1)
	var ret2 *glib.Bytes
	var err2 error
	ret2 = (*glib.Bytes)(unsafe.Pointer(ret1))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *InputStream) ReadBytesAsync(count0 uint64, io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GInputStream
	var count1 C.uint64_t
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = (*C.GInputStream)(this0.InheritedFromGInputStream())
	}
	count1 = C.uint64_t(count0)
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_input_stream_read_bytes_async(this1, count1, io_priority1, cancellable1, callback1)
}
func (this0 *InputStream) ReadBytesFinish(result0 AsyncResultLike) (*glib.Bytes, error) {
	var this1 *C.GInputStream
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GInputStream)(this0.InheritedFromGInputStream())
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_input_stream_read_bytes_finish(this1, result1, &err1)
	var ret2 *glib.Bytes
	var err2 error
	ret2 = (*glib.Bytes)(unsafe.Pointer(ret1))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *InputStream) ReadFinish(result0 AsyncResultLike) (int64, error) {
	var this1 *C.GInputStream
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GInputStream)(this0.InheritedFromGInputStream())
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_input_stream_read_finish(this1, result1, &err1)
	var ret2 int64
	var err2 error
	ret2 = int64(ret1)
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *InputStream) SetPending() (bool, error) {
	var this1 *C.GInputStream
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GInputStream)(this0.InheritedFromGInputStream())
	}
	ret1 := C.g_input_stream_set_pending(this1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *InputStream) Skip(count0 uint64, cancellable0 CancellableLike) (int64, error) {
	var this1 *C.GInputStream
	var count1 C.uint64_t
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GInputStream)(this0.InheritedFromGInputStream())
	}
	count1 = C.uint64_t(count0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_input_stream_skip(this1, count1, cancellable1, &err1)
	var ret2 int64
	var err2 error
	ret2 = int64(ret1)
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *InputStream) SkipAsync(count0 uint64, io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GInputStream
	var count1 C.uint64_t
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = (*C.GInputStream)(this0.InheritedFromGInputStream())
	}
	count1 = C.uint64_t(count0)
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_input_stream_skip_async(this1, count1, io_priority1, cancellable1, callback1)
}
func (this0 *InputStream) SkipFinish(result0 AsyncResultLike) (int64, error) {
	var this1 *C.GInputStream
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GInputStream)(this0.InheritedFromGInputStream())
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_input_stream_skip_finish(this1, result1, &err1)
	var ret2 int64
	var err2 error
	ret2 = int64(ret1)
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
// blacklisted: InputStreamClass (struct)
// blacklisted: InputStreamPrivate (struct)
// blacklisted: InputVector (struct)
// blacklisted: LoadableIcon (interface)
// blacklisted: LoadableIconIface (struct)
const MenuAttributeAction = "action"
const MenuAttributeActionNamespace = "action-namespace"
const MenuAttributeIcon = "icon"
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
		this1 = (*C.GMenuAttributeIter)(this0.InheritedFromGMenuAttributeIter())
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
		this1 = (*C.GMenuAttributeIter)(this0.InheritedFromGMenuAttributeIter())
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
		this1 = (*C.GMenuAttributeIter)(this0.InheritedFromGMenuAttributeIter())
	}
	ret1 := C.g_menu_attribute_iter_get_value(this1)
	var ret2 *glib.Variant
	ret2 = (*glib.Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *MenuAttributeIter) Next() bool {
	var this1 *C.GMenuAttributeIter
	if this0 != nil {
		this1 = (*C.GMenuAttributeIter)(this0.InheritedFromGMenuAttributeIter())
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
		this1 = (*C.GMenuLinkIter)(this0.InheritedFromGMenuLinkIter())
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
		this1 = (*C.GMenuLinkIter)(this0.InheritedFromGMenuLinkIter())
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
		this1 = (*C.GMenuLinkIter)(this0.InheritedFromGMenuLinkIter())
	}
	ret1 := C.g_menu_link_iter_get_value(this1)
	var ret2 *MenuModel
	ret2 = (*MenuModel)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *MenuLinkIter) Next() bool {
	var this1 *C.GMenuLinkIter
	if this0 != nil {
		this1 = (*C.GMenuLinkIter)(this0.InheritedFromGMenuLinkIter())
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
func (this0 *MenuModel) GetItemAttributeValue(item_index0 int32, attribute0 string, expected_type0 *glib.VariantType) *glib.Variant {
	var this1 *C.GMenuModel
	var item_index1 C.int32_t
	var attribute1 *C.char
	var expected_type1 *C.GVariantType
	if this0 != nil {
		this1 = (*C.GMenuModel)(this0.InheritedFromGMenuModel())
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
func (this0 *MenuModel) GetItemLink(item_index0 int32, link0 string) *MenuModel {
	var this1 *C.GMenuModel
	var item_index1 C.int32_t
	var link1 *C.char
	if this0 != nil {
		this1 = (*C.GMenuModel)(this0.InheritedFromGMenuModel())
	}
	item_index1 = C.int32_t(item_index0)
	link1 = _GoStringToGString(link0)
	defer C.free(unsafe.Pointer(link1))
	ret1 := C.g_menu_model_get_item_link(this1, item_index1, link1)
	var ret2 *MenuModel
	ret2 = (*MenuModel)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *MenuModel) GetNItems() int32 {
	var this1 *C.GMenuModel
	if this0 != nil {
		this1 = (*C.GMenuModel)(this0.InheritedFromGMenuModel())
	}
	ret1 := C.g_menu_model_get_n_items(this1)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func (this0 *MenuModel) IsMutable() bool {
	var this1 *C.GMenuModel
	if this0 != nil {
		this1 = (*C.GMenuModel)(this0.InheritedFromGMenuModel())
	}
	ret1 := C.g_menu_model_is_mutable(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *MenuModel) ItemsChanged(position0 int32, removed0 int32, added0 int32) {
	var this1 *C.GMenuModel
	var position1 C.int32_t
	var removed1 C.int32_t
	var added1 C.int32_t
	if this0 != nil {
		this1 = (*C.GMenuModel)(this0.InheritedFromGMenuModel())
	}
	position1 = C.int32_t(position0)
	removed1 = C.int32_t(removed0)
	added1 = C.int32_t(added0)
	C.g_menu_model_items_changed(this1, position1, removed1, added1)
}
func (this0 *MenuModel) IterateItemAttributes(item_index0 int32) *MenuAttributeIter {
	var this1 *C.GMenuModel
	var item_index1 C.int32_t
	if this0 != nil {
		this1 = (*C.GMenuModel)(this0.InheritedFromGMenuModel())
	}
	item_index1 = C.int32_t(item_index0)
	ret1 := C.g_menu_model_iterate_item_attributes(this1, item_index1)
	var ret2 *MenuAttributeIter
	ret2 = (*MenuAttributeIter)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *MenuModel) IterateItemLinks(item_index0 int32) *MenuLinkIter {
	var this1 *C.GMenuModel
	var item_index1 C.int32_t
	if this0 != nil {
		this1 = (*C.GMenuModel)(this0.InheritedFromGMenuModel())
	}
	item_index1 = C.int32_t(item_index0)
	ret1 := C.g_menu_model_iterate_item_links(this1, item_index1)
	var ret2 *MenuLinkIter
	ret2 = (*MenuLinkIter)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
// blacklisted: MenuModelClass (struct)
// blacklisted: MenuModelPrivate (struct)
type MountLike interface {
	ImplementsGMount() *C.GMount
}

type Mount struct {
	gobject.Object
	MountImpl
}

func (*Mount) GetStaticType() gobject.Type {
	return gobject.Type(C.g_mount_get_type())
}


type MountImpl struct {}

func ToMount(objlike gobject.ObjectLike) *Mount {
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), gobject.Type(C.g_mount_get_type()))
	if obj != nil {
		return (*Mount)(obj)
	}
	panic("cannot cast to Mount")
}

func (this0 *MountImpl) ImplementsGMount() *C.GMount {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GMount)((*gobject.Object)(unsafe.Pointer(obj)).C)
}
func (this0 *MountImpl) CanEject() bool {
	var this1 *C.GMount
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	ret1 := C.g_mount_can_eject(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *MountImpl) CanUnmount() bool {
	var this1 *C.GMount
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	ret1 := C.g_mount_can_unmount(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *MountImpl) Eject(flags0 MountUnmountFlags, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GMount
	var flags1 C.GMountUnmountFlags
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	flags1 = C.GMountUnmountFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_mount_eject(this1, flags1, cancellable1, callback1)
}
func (this0 *MountImpl) EjectFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GMount
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_mount_eject_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *MountImpl) EjectWithOperation(flags0 MountUnmountFlags, mount_operation0 MountOperationLike, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GMount
	var flags1 C.GMountUnmountFlags
	var mount_operation1 *C.GMountOperation
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	flags1 = C.GMountUnmountFlags(flags0)
	if mount_operation0 != nil {
		mount_operation1 = (*C.GMountOperation)(mount_operation0.InheritedFromGMountOperation())
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_mount_eject_with_operation(this1, flags1, mount_operation1, cancellable1, callback1)
}
func (this0 *MountImpl) EjectWithOperationFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GMount
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_mount_eject_with_operation_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *MountImpl) GetDefaultLocation() *File {
	var this1 *C.GMount
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	ret1 := C.g_mount_get_default_location(this1)
	var ret2 *File
	ret2 = (*File)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *MountImpl) GetDrive() *Drive {
	var this1 *C.GMount
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	ret1 := C.g_mount_get_drive(this1)
	var ret2 *Drive
	ret2 = (*Drive)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *MountImpl) GetIcon() *Icon {
	var this1 *C.GMount
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	ret1 := C.g_mount_get_icon(this1)
	var ret2 *Icon
	ret2 = (*Icon)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *MountImpl) GetName() string {
	var this1 *C.GMount
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	ret1 := C.g_mount_get_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *MountImpl) GetRoot() *File {
	var this1 *C.GMount
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	ret1 := C.g_mount_get_root(this1)
	var ret2 *File
	ret2 = (*File)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *MountImpl) GetSortKey() string {
	var this1 *C.GMount
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	ret1 := C.g_mount_get_sort_key(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *MountImpl) GetSymbolicIcon() *Icon {
	var this1 *C.GMount
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	ret1 := C.g_mount_get_symbolic_icon(this1)
	var ret2 *Icon
	ret2 = (*Icon)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *MountImpl) GetUuid() string {
	var this1 *C.GMount
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	ret1 := C.g_mount_get_uuid(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *MountImpl) GetVolume() *Volume {
	var this1 *C.GMount
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	ret1 := C.g_mount_get_volume(this1)
	var ret2 *Volume
	ret2 = (*Volume)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *MountImpl) GuessContentType(force_rescan0 bool, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GMount
	var force_rescan1 C.int
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	force_rescan1 = _GoBoolToCBool(force_rescan0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_mount_guess_content_type(this1, force_rescan1, cancellable1, callback1)
}
func (this0 *MountImpl) GuessContentTypeFinish(result0 AsyncResultLike) ([]string, error) {
	var this1 *C.GMount
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_mount_guess_content_type_finish(this1, result1, &err1)
	var ret2 []string
	var err2 error
	ret2 = make([]string, uint(C._array_length(unsafe.Pointer(ret1))))
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *MountImpl) GuessContentTypeSync(force_rescan0 bool, cancellable0 CancellableLike) ([]string, error) {
	var this1 *C.GMount
	var force_rescan1 C.int
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	force_rescan1 = _GoBoolToCBool(force_rescan0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_mount_guess_content_type_sync(this1, force_rescan1, cancellable1, &err1)
	var ret2 []string
	var err2 error
	ret2 = make([]string, uint(C._array_length(unsafe.Pointer(ret1))))
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *MountImpl) IsShadowed() bool {
	var this1 *C.GMount
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	ret1 := C.g_mount_is_shadowed(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *MountImpl) Remount(flags0 MountMountFlags, mount_operation0 MountOperationLike, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GMount
	var flags1 C.GMountMountFlags
	var mount_operation1 *C.GMountOperation
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	flags1 = C.GMountMountFlags(flags0)
	if mount_operation0 != nil {
		mount_operation1 = (*C.GMountOperation)(mount_operation0.InheritedFromGMountOperation())
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_mount_remount(this1, flags1, mount_operation1, cancellable1, callback1)
}
func (this0 *MountImpl) RemountFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GMount
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_mount_remount_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *MountImpl) Shadow() {
	var this1 *C.GMount
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	C.g_mount_shadow(this1)
}
func (this0 *MountImpl) Unmount(flags0 MountUnmountFlags, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GMount
	var flags1 C.GMountUnmountFlags
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	flags1 = C.GMountUnmountFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_mount_unmount(this1, flags1, cancellable1, callback1)
}
func (this0 *MountImpl) UnmountFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GMount
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_mount_unmount_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *MountImpl) UnmountWithOperation(flags0 MountUnmountFlags, mount_operation0 MountOperationLike, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GMount
	var flags1 C.GMountUnmountFlags
	var mount_operation1 *C.GMountOperation
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	flags1 = C.GMountUnmountFlags(flags0)
	if mount_operation0 != nil {
		mount_operation1 = (*C.GMountOperation)(mount_operation0.InheritedFromGMountOperation())
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_mount_unmount_with_operation(this1, flags1, mount_operation1, cancellable1, callback1)
}
func (this0 *MountImpl) UnmountWithOperationFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GMount
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_mount_unmount_with_operation_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *MountImpl) Unshadow() {
	var this1 *C.GMount
	if this0 != nil {
		this1 = this0.ImplementsGMount()
	}
	C.g_mount_unshadow(this1)
}
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
		this1 = (*C.GMountOperation)(this0.InheritedFromGMountOperation())
	}
	ret1 := C.g_mount_operation_get_anonymous(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *MountOperation) GetChoice() int32 {
	var this1 *C.GMountOperation
	if this0 != nil {
		this1 = (*C.GMountOperation)(this0.InheritedFromGMountOperation())
	}
	ret1 := C.g_mount_operation_get_choice(this1)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func (this0 *MountOperation) GetDomain() string {
	var this1 *C.GMountOperation
	if this0 != nil {
		this1 = (*C.GMountOperation)(this0.InheritedFromGMountOperation())
	}
	ret1 := C.g_mount_operation_get_domain(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *MountOperation) GetPassword() string {
	var this1 *C.GMountOperation
	if this0 != nil {
		this1 = (*C.GMountOperation)(this0.InheritedFromGMountOperation())
	}
	ret1 := C.g_mount_operation_get_password(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *MountOperation) GetPasswordSave() PasswordSave {
	var this1 *C.GMountOperation
	if this0 != nil {
		this1 = (*C.GMountOperation)(this0.InheritedFromGMountOperation())
	}
	ret1 := C.g_mount_operation_get_password_save(this1)
	var ret2 PasswordSave
	ret2 = PasswordSave(ret1)
	return ret2
}
func (this0 *MountOperation) GetUsername() string {
	var this1 *C.GMountOperation
	if this0 != nil {
		this1 = (*C.GMountOperation)(this0.InheritedFromGMountOperation())
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
		this1 = (*C.GMountOperation)(this0.InheritedFromGMountOperation())
	}
	result1 = C.GMountOperationResult(result0)
	C.g_mount_operation_reply(this1, result1)
}
func (this0 *MountOperation) SetAnonymous(anonymous0 bool) {
	var this1 *C.GMountOperation
	var anonymous1 C.int
	if this0 != nil {
		this1 = (*C.GMountOperation)(this0.InheritedFromGMountOperation())
	}
	anonymous1 = _GoBoolToCBool(anonymous0)
	C.g_mount_operation_set_anonymous(this1, anonymous1)
}
func (this0 *MountOperation) SetChoice(choice0 int32) {
	var this1 *C.GMountOperation
	var choice1 C.int32_t
	if this0 != nil {
		this1 = (*C.GMountOperation)(this0.InheritedFromGMountOperation())
	}
	choice1 = C.int32_t(choice0)
	C.g_mount_operation_set_choice(this1, choice1)
}
func (this0 *MountOperation) SetDomain(domain0 string) {
	var this1 *C.GMountOperation
	var domain1 *C.char
	if this0 != nil {
		this1 = (*C.GMountOperation)(this0.InheritedFromGMountOperation())
	}
	domain1 = _GoStringToGString(domain0)
	defer C.free(unsafe.Pointer(domain1))
	C.g_mount_operation_set_domain(this1, domain1)
}
func (this0 *MountOperation) SetPassword(password0 string) {
	var this1 *C.GMountOperation
	var password1 *C.char
	if this0 != nil {
		this1 = (*C.GMountOperation)(this0.InheritedFromGMountOperation())
	}
	password1 = _GoStringToGString(password0)
	defer C.free(unsafe.Pointer(password1))
	C.g_mount_operation_set_password(this1, password1)
}
func (this0 *MountOperation) SetPasswordSave(save0 PasswordSave) {
	var this1 *C.GMountOperation
	var save1 C.GPasswordSave
	if this0 != nil {
		this1 = (*C.GMountOperation)(this0.InheritedFromGMountOperation())
	}
	save1 = C.GPasswordSave(save0)
	C.g_mount_operation_set_password_save(this1, save1)
}
func (this0 *MountOperation) SetUsername(username0 string) {
	var this1 *C.GMountOperation
	var username1 *C.char
	if this0 != nil {
		this1 = (*C.GMountOperation)(this0.InheritedFromGMountOperation())
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
// blacklisted: Notification (object)
type OutputStreamLike interface {
	gobject.ObjectLike
	InheritedFromGOutputStream() *C.GOutputStream
}

type OutputStream struct {
	gobject.Object
	
}

func ToOutputStream(objlike gobject.ObjectLike) *OutputStream {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*OutputStream)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*OutputStream)(obj)
	}
	panic("cannot cast to OutputStream")
}

func (this0 *OutputStream) InheritedFromGOutputStream() *C.GOutputStream {
	if this0 == nil {
		return nil
	}
	return (*C.GOutputStream)(this0.C)
}

func (this0 *OutputStream) GetStaticType() gobject.Type {
	return gobject.Type(C.g_output_stream_get_type())
}

func OutputStreamGetType() gobject.Type {
	return (*OutputStream)(nil).GetStaticType()
}
func (this0 *OutputStream) ClearPending() {
	var this1 *C.GOutputStream
	if this0 != nil {
		this1 = (*C.GOutputStream)(this0.InheritedFromGOutputStream())
	}
	C.g_output_stream_clear_pending(this1)
}
func (this0 *OutputStream) Close(cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GOutputStream
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GOutputStream)(this0.InheritedFromGOutputStream())
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_output_stream_close(this1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *OutputStream) CloseAsync(io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GOutputStream
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = (*C.GOutputStream)(this0.InheritedFromGOutputStream())
	}
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_output_stream_close_async(this1, io_priority1, cancellable1, callback1)
}
func (this0 *OutputStream) CloseFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GOutputStream
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GOutputStream)(this0.InheritedFromGOutputStream())
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_output_stream_close_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *OutputStream) Flush(cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GOutputStream
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GOutputStream)(this0.InheritedFromGOutputStream())
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_output_stream_flush(this1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *OutputStream) FlushAsync(io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GOutputStream
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = (*C.GOutputStream)(this0.InheritedFromGOutputStream())
	}
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_output_stream_flush_async(this1, io_priority1, cancellable1, callback1)
}
func (this0 *OutputStream) FlushFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GOutputStream
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GOutputStream)(this0.InheritedFromGOutputStream())
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_output_stream_flush_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *OutputStream) HasPending() bool {
	var this1 *C.GOutputStream
	if this0 != nil {
		this1 = (*C.GOutputStream)(this0.InheritedFromGOutputStream())
	}
	ret1 := C.g_output_stream_has_pending(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *OutputStream) IsClosed() bool {
	var this1 *C.GOutputStream
	if this0 != nil {
		this1 = (*C.GOutputStream)(this0.InheritedFromGOutputStream())
	}
	ret1 := C.g_output_stream_is_closed(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *OutputStream) IsClosing() bool {
	var this1 *C.GOutputStream
	if this0 != nil {
		this1 = (*C.GOutputStream)(this0.InheritedFromGOutputStream())
	}
	ret1 := C.g_output_stream_is_closing(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *OutputStream) SetPending() (bool, error) {
	var this1 *C.GOutputStream
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GOutputStream)(this0.InheritedFromGOutputStream())
	}
	ret1 := C.g_output_stream_set_pending(this1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *OutputStream) Splice(source0 InputStreamLike, flags0 OutputStreamSpliceFlags, cancellable0 CancellableLike) (int64, error) {
	var this1 *C.GOutputStream
	var source1 *C.GInputStream
	var flags1 C.GOutputStreamSpliceFlags
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GOutputStream)(this0.InheritedFromGOutputStream())
	}
	if source0 != nil {
		source1 = (*C.GInputStream)(source0.InheritedFromGInputStream())
	}
	flags1 = C.GOutputStreamSpliceFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_output_stream_splice(this1, source1, flags1, cancellable1, &err1)
	var ret2 int64
	var err2 error
	ret2 = int64(ret1)
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *OutputStream) SpliceAsync(source0 InputStreamLike, flags0 OutputStreamSpliceFlags, io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GOutputStream
	var source1 *C.GInputStream
	var flags1 C.GOutputStreamSpliceFlags
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = (*C.GOutputStream)(this0.InheritedFromGOutputStream())
	}
	if source0 != nil {
		source1 = (*C.GInputStream)(source0.InheritedFromGInputStream())
	}
	flags1 = C.GOutputStreamSpliceFlags(flags0)
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_output_stream_splice_async(this1, source1, flags1, io_priority1, cancellable1, callback1)
}
func (this0 *OutputStream) SpliceFinish(result0 AsyncResultLike) (int64, error) {
	var this1 *C.GOutputStream
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GOutputStream)(this0.InheritedFromGOutputStream())
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_output_stream_splice_finish(this1, result1, &err1)
	var ret2 int64
	var err2 error
	ret2 = int64(ret1)
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *OutputStream) Write(buffer0 []uint8, cancellable0 CancellableLike) (int64, error) {
	var this1 *C.GOutputStream
	var buffer1 *C.uint8_t
	var count1 C.uint64_t
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GOutputStream)(this0.InheritedFromGOutputStream())
	}
	buffer1 = (*C.uint8_t)(C.malloc(C.size_t(int(unsafe.Sizeof(*buffer1)) * len(buffer0))))
	defer C.free(unsafe.Pointer(buffer1))
	for i, e := range buffer0 {
		(*(*[999999]C.uint8_t)(unsafe.Pointer(buffer1)))[i] = C.uint8_t(e)
	}
	count1 = C.uint64_t(len(buffer0))
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_output_stream_write(this1, buffer1, count1, cancellable1, &err1)
	var ret2 int64
	var err2 error
	ret2 = int64(ret1)
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *OutputStream) WriteAll(buffer0 []uint8, cancellable0 CancellableLike) (uint64, bool, error) {
	var this1 *C.GOutputStream
	var buffer1 *C.uint8_t
	var count1 C.uint64_t
	var cancellable1 *C.GCancellable
	var bytes_written1 C.uint64_t
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GOutputStream)(this0.InheritedFromGOutputStream())
	}
	buffer1 = (*C.uint8_t)(C.malloc(C.size_t(int(unsafe.Sizeof(*buffer1)) * len(buffer0))))
	defer C.free(unsafe.Pointer(buffer1))
	for i, e := range buffer0 {
		(*(*[999999]C.uint8_t)(unsafe.Pointer(buffer1)))[i] = C.uint8_t(e)
	}
	count1 = C.uint64_t(len(buffer0))
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_output_stream_write_all(this1, buffer1, count1, &bytes_written1, cancellable1, &err1)
	var bytes_written2 uint64
	var ret2 bool
	var err2 error
	bytes_written2 = uint64(bytes_written1)
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return bytes_written2, ret2, err2
}
func (this0 *OutputStream) WriteAsync(buffer0 []uint8, io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GOutputStream
	var buffer1 *C.uint8_t
	var count1 C.uint64_t
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = (*C.GOutputStream)(this0.InheritedFromGOutputStream())
	}
	buffer1 = (*C.uint8_t)(C.malloc(C.size_t(int(unsafe.Sizeof(*buffer1)) * len(buffer0))))
	defer C.free(unsafe.Pointer(buffer1))
	for i, e := range buffer0 {
		(*(*[999999]C.uint8_t)(unsafe.Pointer(buffer1)))[i] = C.uint8_t(e)
	}
	count1 = C.uint64_t(len(buffer0))
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_output_stream_write_async(this1, buffer1, count1, io_priority1, cancellable1, callback1)
}
func (this0 *OutputStream) WriteBytes(bytes0 *glib.Bytes, cancellable0 CancellableLike) (int64, error) {
	var this1 *C.GOutputStream
	var bytes1 *C.GBytes
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GOutputStream)(this0.InheritedFromGOutputStream())
	}
	bytes1 = (*C.GBytes)(unsafe.Pointer(bytes0))
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_output_stream_write_bytes(this1, bytes1, cancellable1, &err1)
	var ret2 int64
	var err2 error
	ret2 = int64(ret1)
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *OutputStream) WriteBytesAsync(bytes0 *glib.Bytes, io_priority0 int32, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GOutputStream
	var bytes1 *C.GBytes
	var io_priority1 C.int32_t
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = (*C.GOutputStream)(this0.InheritedFromGOutputStream())
	}
	bytes1 = (*C.GBytes)(unsafe.Pointer(bytes0))
	io_priority1 = C.int32_t(io_priority0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_output_stream_write_bytes_async(this1, bytes1, io_priority1, cancellable1, callback1)
}
func (this0 *OutputStream) WriteBytesFinish(result0 AsyncResultLike) (int64, error) {
	var this1 *C.GOutputStream
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GOutputStream)(this0.InheritedFromGOutputStream())
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_output_stream_write_bytes_finish(this1, result1, &err1)
	var ret2 int64
	var err2 error
	ret2 = int64(ret1)
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *OutputStream) WriteFinish(result0 AsyncResultLike) (int64, error) {
	var this1 *C.GOutputStream
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = (*C.GOutputStream)(this0.InheritedFromGOutputStream())
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_output_stream_write_finish(this1, result1, &err1)
	var ret2 int64
	var err2 error
	ret2 = int64(ret1)
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
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
		this1 = (*C.GPermission)(this0.InheritedFromGPermission())
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_permission_acquire(this1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *Permission) AcquireAsync(cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GPermission
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = (*C.GPermission)(this0.InheritedFromGPermission())
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
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
		this1 = (*C.GPermission)(this0.InheritedFromGPermission())
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_permission_acquire_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *Permission) GetAllowed() bool {
	var this1 *C.GPermission
	if this0 != nil {
		this1 = (*C.GPermission)(this0.InheritedFromGPermission())
	}
	ret1 := C.g_permission_get_allowed(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Permission) GetCanAcquire() bool {
	var this1 *C.GPermission
	if this0 != nil {
		this1 = (*C.GPermission)(this0.InheritedFromGPermission())
	}
	ret1 := C.g_permission_get_can_acquire(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Permission) GetCanRelease() bool {
	var this1 *C.GPermission
	if this0 != nil {
		this1 = (*C.GPermission)(this0.InheritedFromGPermission())
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
		this1 = (*C.GPermission)(this0.InheritedFromGPermission())
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
		this1 = (*C.GPermission)(this0.InheritedFromGPermission())
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_permission_release(this1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *Permission) ReleaseAsync(cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GPermission
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = (*C.GPermission)(this0.InheritedFromGPermission())
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
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
		this1 = (*C.GPermission)(this0.InheritedFromGPermission())
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_permission_release_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
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
// blacklisted: PropertyAction (object)
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
type SeekableLike interface {
	ImplementsGSeekable() *C.GSeekable
}

type Seekable struct {
	gobject.Object
	SeekableImpl
}

func (*Seekable) GetStaticType() gobject.Type {
	return gobject.Type(C.g_seekable_get_type())
}


type SeekableImpl struct {}

func ToSeekable(objlike gobject.ObjectLike) *Seekable {
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), gobject.Type(C.g_seekable_get_type()))
	if obj != nil {
		return (*Seekable)(obj)
	}
	panic("cannot cast to Seekable")
}

func (this0 *SeekableImpl) ImplementsGSeekable() *C.GSeekable {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GSeekable)((*gobject.Object)(unsafe.Pointer(obj)).C)
}
func (this0 *SeekableImpl) CanSeek() bool {
	var this1 *C.GSeekable
	if this0 != nil {
		this1 = this0.ImplementsGSeekable()
	}
	ret1 := C.g_seekable_can_seek(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *SeekableImpl) CanTruncate() bool {
	var this1 *C.GSeekable
	if this0 != nil {
		this1 = this0.ImplementsGSeekable()
	}
	ret1 := C.g_seekable_can_truncate(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *SeekableImpl) Seek(offset0 int64, type0 glib.SeekType, cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GSeekable
	var offset1 C.int64_t
	var type1 C.GSeekType
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGSeekable()
	}
	offset1 = C.int64_t(offset0)
	type1 = C.GSeekType(type0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_seekable_seek(this1, offset1, type1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *SeekableImpl) Tell() int64 {
	var this1 *C.GSeekable
	if this0 != nil {
		this1 = this0.ImplementsGSeekable()
	}
	ret1 := C.g_seekable_tell(this1)
	var ret2 int64
	ret2 = int64(ret1)
	return ret2
}
func (this0 *SeekableImpl) Truncate(offset0 int64, cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GSeekable
	var offset1 C.int64_t
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGSeekable()
	}
	offset1 = C.int64_t(offset0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	ret1 := C.g_seekable_truncate(this1, offset1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
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
func NewSettingsWithBackend(schema_id0 string, backend0 *SettingsBackend) *Settings {
	var schema_id1 *C.char
	var backend1 *C.GSettingsBackend
	schema_id1 = _GoStringToGString(schema_id0)
	defer C.free(unsafe.Pointer(schema_id1))
	backend1 = (*C.GSettingsBackend)(unsafe.Pointer(backend0))
	ret1 := C.g_settings_new_with_backend(schema_id1, backend1)
	var ret2 *Settings
	ret2 = (*Settings)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func NewSettingsWithBackendAndPath(schema_id0 string, backend0 *SettingsBackend, path0 string) *Settings {
	var schema_id1 *C.char
	var backend1 *C.GSettingsBackend
	var path1 *C.char
	schema_id1 = _GoStringToGString(schema_id0)
	defer C.free(unsafe.Pointer(schema_id1))
	backend1 = (*C.GSettingsBackend)(unsafe.Pointer(backend0))
	path1 = _GoStringToGString(path0)
	defer C.free(unsafe.Pointer(path1))
	ret1 := C.g_settings_new_with_backend_and_path(schema_id1, backend1, path1)
	var ret2 *Settings
	ret2 = (*Settings)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
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
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
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
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	if object0 != nil {
		object1 = (*C.GObject)(object0.InheritedFromGObject())
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
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	if object0 != nil {
		object1 = (*C.GObject)(object0.InheritedFromGObject())
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
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
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
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
	}
	C.g_settings_delay(this1)
}
func (this0 *Settings) GetBoolean(key0 string) bool {
	var this1 *C.GSettings
	var key1 *C.char
	if this0 != nil {
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
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
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
	}
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	ret1 := C.g_settings_get_child(this1, name1)
	var ret2 *Settings
	ret2 = (*Settings)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Settings) GetDefaultValue(key0 string) *glib.Variant {
	var this1 *C.GSettings
	var key1 *C.char
	if this0 != nil {
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_settings_get_default_value(this1, key1)
	var ret2 *glib.Variant
	ret2 = (*glib.Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Settings) GetDouble(key0 string) float64 {
	var this1 *C.GSettings
	var key1 *C.char
	if this0 != nil {
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_settings_get_double(this1, key1)
	var ret2 float64
	ret2 = float64(ret1)
	return ret2
}
func (this0 *Settings) GetEnum(key0 string) int32 {
	var this1 *C.GSettings
	var key1 *C.char
	if this0 != nil {
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_settings_get_enum(this1, key1)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func (this0 *Settings) GetFlags(key0 string) uint32 {
	var this1 *C.GSettings
	var key1 *C.char
	if this0 != nil {
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_settings_get_flags(this1, key1)
	var ret2 uint32
	ret2 = uint32(ret1)
	return ret2
}
func (this0 *Settings) GetHasUnapplied() bool {
	var this1 *C.GSettings
	if this0 != nil {
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
	}
	ret1 := C.g_settings_get_has_unapplied(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Settings) GetInt(key0 string) int32 {
	var this1 *C.GSettings
	var key1 *C.char
	if this0 != nil {
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_settings_get_int(this1, key1)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
// blacklisted: Settings.get_mapped (method)
func (this0 *Settings) GetRange(key0 string) *glib.Variant {
	var this1 *C.GSettings
	var key1 *C.char
	if this0 != nil {
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
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
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
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
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
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
func (this0 *Settings) GetUint(key0 string) uint32 {
	var this1 *C.GSettings
	var key1 *C.char
	if this0 != nil {
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_settings_get_uint(this1, key1)
	var ret2 uint32
	ret2 = uint32(ret1)
	return ret2
}
func (this0 *Settings) GetUserValue(key0 string) *glib.Variant {
	var this1 *C.GSettings
	var key1 *C.char
	if this0 != nil {
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_settings_get_user_value(this1, key1)
	var ret2 *glib.Variant
	ret2 = (*glib.Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Settings) GetValue(key0 string) *glib.Variant {
	var this1 *C.GSettings
	var key1 *C.char
	if this0 != nil {
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
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
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
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
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
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
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
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
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
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
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	C.g_settings_reset(this1, key1)
}
func (this0 *Settings) Revert() {
	var this1 *C.GSettings
	if this0 != nil {
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
	}
	C.g_settings_revert(this1)
}
func (this0 *Settings) SetBoolean(key0 string, value0 bool) bool {
	var this1 *C.GSettings
	var key1 *C.char
	var value1 C.int
	if this0 != nil {
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
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
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = C.double(value0)
	ret1 := C.g_settings_set_double(this1, key1, value1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Settings) SetEnum(key0 string, value0 int32) bool {
	var this1 *C.GSettings
	var key1 *C.char
	var value1 C.int32_t
	if this0 != nil {
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = C.int32_t(value0)
	ret1 := C.g_settings_set_enum(this1, key1, value1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Settings) SetFlags(key0 string, value0 uint32) bool {
	var this1 *C.GSettings
	var key1 *C.char
	var value1 C.uint32_t
	if this0 != nil {
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = C.uint32_t(value0)
	ret1 := C.g_settings_set_flags(this1, key1, value1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Settings) SetInt(key0 string, value0 int32) bool {
	var this1 *C.GSettings
	var key1 *C.char
	var value1 C.int32_t
	if this0 != nil {
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
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
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
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
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
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
func (this0 *Settings) SetUint(key0 string, value0 uint32) bool {
	var this1 *C.GSettings
	var key1 *C.char
	var value1 C.uint32_t
	if this0 != nil {
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
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
		this1 = (*C.GSettings)(this0.InheritedFromGSettings())
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = (*C.GVariant)(unsafe.Pointer(value0))
	ret1 := C.g_settings_set_value(this1, key1, value1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
type SettingsBackend struct {}
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
// blacklisted: SettingsSchemaKey (struct)
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
// blacklisted: Subprocess (object)
type SubprocessFlags C.uint32_t
const (
	SubprocessFlagsNone SubprocessFlags = 0
	SubprocessFlagsStdinPipe SubprocessFlags = 1
	SubprocessFlagsStdinInherit SubprocessFlags = 2
	SubprocessFlagsStdoutPipe SubprocessFlags = 4
	SubprocessFlagsStdoutSilence SubprocessFlags = 8
	SubprocessFlagsStderrPipe SubprocessFlags = 16
	SubprocessFlagsStderrSilence SubprocessFlags = 32
	SubprocessFlagsStderrMerge SubprocessFlags = 64
	SubprocessFlagsInheritFds SubprocessFlags = 128
)
// blacklisted: SubprocessLauncher (object)
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
type TlsCertificateRequestFlags C.uint32_t
const (
	TlsCertificateRequestFlagsNone TlsCertificateRequestFlags = 0
)
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
type VolumeLike interface {
	ImplementsGVolume() *C.GVolume
}

type Volume struct {
	gobject.Object
	VolumeImpl
}

func (*Volume) GetStaticType() gobject.Type {
	return gobject.Type(C.g_volume_get_type())
}


type VolumeImpl struct {}

func ToVolume(objlike gobject.ObjectLike) *Volume {
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), gobject.Type(C.g_volume_get_type()))
	if obj != nil {
		return (*Volume)(obj)
	}
	panic("cannot cast to Volume")
}

func (this0 *VolumeImpl) ImplementsGVolume() *C.GVolume {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GVolume)((*gobject.Object)(unsafe.Pointer(obj)).C)
}
func (this0 *VolumeImpl) CanEject() bool {
	var this1 *C.GVolume
	if this0 != nil {
		this1 = this0.ImplementsGVolume()
	}
	ret1 := C.g_volume_can_eject(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *VolumeImpl) CanMount() bool {
	var this1 *C.GVolume
	if this0 != nil {
		this1 = this0.ImplementsGVolume()
	}
	ret1 := C.g_volume_can_mount(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *VolumeImpl) Eject(flags0 MountUnmountFlags, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GVolume
	var flags1 C.GMountUnmountFlags
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGVolume()
	}
	flags1 = C.GMountUnmountFlags(flags0)
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_volume_eject(this1, flags1, cancellable1, callback1)
}
func (this0 *VolumeImpl) EjectFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GVolume
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGVolume()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_volume_eject_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *VolumeImpl) EjectWithOperation(flags0 MountUnmountFlags, mount_operation0 MountOperationLike, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GVolume
	var flags1 C.GMountUnmountFlags
	var mount_operation1 *C.GMountOperation
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGVolume()
	}
	flags1 = C.GMountUnmountFlags(flags0)
	if mount_operation0 != nil {
		mount_operation1 = (*C.GMountOperation)(mount_operation0.InheritedFromGMountOperation())
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_volume_eject_with_operation(this1, flags1, mount_operation1, cancellable1, callback1)
}
func (this0 *VolumeImpl) EjectWithOperationFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GVolume
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGVolume()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_volume_eject_with_operation_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *VolumeImpl) EnumerateIdentifiers() []string {
	var this1 *C.GVolume
	if this0 != nil {
		this1 = this0.ImplementsGVolume()
	}
	ret1 := C.g_volume_enumerate_identifiers(this1)
	var ret2 []string
	ret2 = make([]string, uint(C._array_length(unsafe.Pointer(ret1))))
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return ret2
}
func (this0 *VolumeImpl) GetActivationRoot() *File {
	var this1 *C.GVolume
	if this0 != nil {
		this1 = this0.ImplementsGVolume()
	}
	ret1 := C.g_volume_get_activation_root(this1)
	var ret2 *File
	ret2 = (*File)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *VolumeImpl) GetDrive() *Drive {
	var this1 *C.GVolume
	if this0 != nil {
		this1 = this0.ImplementsGVolume()
	}
	ret1 := C.g_volume_get_drive(this1)
	var ret2 *Drive
	ret2 = (*Drive)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *VolumeImpl) GetIcon() *Icon {
	var this1 *C.GVolume
	if this0 != nil {
		this1 = this0.ImplementsGVolume()
	}
	ret1 := C.g_volume_get_icon(this1)
	var ret2 *Icon
	ret2 = (*Icon)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *VolumeImpl) GetIdentifier(kind0 string) string {
	var this1 *C.GVolume
	var kind1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsGVolume()
	}
	kind1 = _GoStringToGString(kind0)
	defer C.free(unsafe.Pointer(kind1))
	ret1 := C.g_volume_get_identifier(this1, kind1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *VolumeImpl) GetMount() *Mount {
	var this1 *C.GVolume
	if this0 != nil {
		this1 = this0.ImplementsGVolume()
	}
	ret1 := C.g_volume_get_mount(this1)
	var ret2 *Mount
	ret2 = (*Mount)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *VolumeImpl) GetName() string {
	var this1 *C.GVolume
	if this0 != nil {
		this1 = this0.ImplementsGVolume()
	}
	ret1 := C.g_volume_get_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *VolumeImpl) GetSortKey() string {
	var this1 *C.GVolume
	if this0 != nil {
		this1 = this0.ImplementsGVolume()
	}
	ret1 := C.g_volume_get_sort_key(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *VolumeImpl) GetSymbolicIcon() *Icon {
	var this1 *C.GVolume
	if this0 != nil {
		this1 = this0.ImplementsGVolume()
	}
	ret1 := C.g_volume_get_symbolic_icon(this1)
	var ret2 *Icon
	ret2 = (*Icon)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *VolumeImpl) GetUuid() string {
	var this1 *C.GVolume
	if this0 != nil {
		this1 = this0.ImplementsGVolume()
	}
	ret1 := C.g_volume_get_uuid(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *VolumeImpl) Mount(flags0 MountMountFlags, mount_operation0 MountOperationLike, cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GVolume
	var flags1 C.GMountMountFlags
	var mount_operation1 *C.GMountOperation
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.ImplementsGVolume()
	}
	flags1 = C.GMountMountFlags(flags0)
	if mount_operation0 != nil {
		mount_operation1 = (*C.GMountOperation)(mount_operation0.InheritedFromGMountOperation())
	}
	if cancellable0 != nil {
		cancellable1 = (*C.GCancellable)(cancellable0.InheritedFromGCancellable())
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_volume_mount(this1, flags1, mount_operation1, cancellable1, callback1)
}
func (this0 *VolumeImpl) MountFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GVolume
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.ImplementsGVolume()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()
	}
	ret1 := C.g_volume_mount_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *VolumeImpl) ShouldAutomount() bool {
	var this1 *C.GVolume
	if this0 != nil {
		this1 = this0.ImplementsGVolume()
	}
	ret1 := C.g_volume_should_automount(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
// blacklisted: VolumeIface (struct)
type VolumeMonitorLike interface {
	gobject.ObjectLike
	InheritedFromGVolumeMonitor() *C.GVolumeMonitor
}

type VolumeMonitor struct {
	gobject.Object
	
}

func ToVolumeMonitor(objlike gobject.ObjectLike) *VolumeMonitor {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*VolumeMonitor)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*VolumeMonitor)(obj)
	}
	panic("cannot cast to VolumeMonitor")
}

func (this0 *VolumeMonitor) InheritedFromGVolumeMonitor() *C.GVolumeMonitor {
	if this0 == nil {
		return nil
	}
	return (*C.GVolumeMonitor)(this0.C)
}

func (this0 *VolumeMonitor) GetStaticType() gobject.Type {
	return gobject.Type(C.g_volume_monitor_get_type())
}

func VolumeMonitorGetType() gobject.Type {
	return (*VolumeMonitor)(nil).GetStaticType()
}
func VolumeMonitorAdoptOrphanMount(mount0 MountLike) *Volume {
	var mount1 *C.GMount
	if mount0 != nil {
		mount1 = mount0.ImplementsGMount()
	}
	ret1 := C.g_volume_monitor_adopt_orphan_mount(mount1)
	var ret2 *Volume
	ret2 = (*Volume)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func VolumeMonitorGet() *VolumeMonitor {
	ret1 := C.g_volume_monitor_get()
	var ret2 *VolumeMonitor
	ret2 = (*VolumeMonitor)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *VolumeMonitor) GetConnectedDrives() []*Drive {
	var this1 *C.GVolumeMonitor
	if this0 != nil {
		this1 = (*C.GVolumeMonitor)(this0.InheritedFromGVolumeMonitor())
	}
	ret1 := C.g_volume_monitor_get_connected_drives(this1)
	var ret2 []*Drive
	for iter := (*_GList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt *Drive
		elt = (*Drive)(gobject.ObjectWrap(unsafe.Pointer((*C.GDrive)(iter.data)), false))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func (this0 *VolumeMonitor) GetMountForUuid(uuid0 string) *Mount {
	var this1 *C.GVolumeMonitor
	var uuid1 *C.char
	if this0 != nil {
		this1 = (*C.GVolumeMonitor)(this0.InheritedFromGVolumeMonitor())
	}
	uuid1 = _GoStringToGString(uuid0)
	defer C.free(unsafe.Pointer(uuid1))
	ret1 := C.g_volume_monitor_get_mount_for_uuid(this1, uuid1)
	var ret2 *Mount
	ret2 = (*Mount)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *VolumeMonitor) GetMounts() []*Mount {
	var this1 *C.GVolumeMonitor
	if this0 != nil {
		this1 = (*C.GVolumeMonitor)(this0.InheritedFromGVolumeMonitor())
	}
	ret1 := C.g_volume_monitor_get_mounts(this1)
	var ret2 []*Mount
	for iter := (*_GList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt *Mount
		elt = (*Mount)(gobject.ObjectWrap(unsafe.Pointer((*C.GMount)(iter.data)), false))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func (this0 *VolumeMonitor) GetVolumeForUuid(uuid0 string) *Volume {
	var this1 *C.GVolumeMonitor
	var uuid1 *C.char
	if this0 != nil {
		this1 = (*C.GVolumeMonitor)(this0.InheritedFromGVolumeMonitor())
	}
	uuid1 = _GoStringToGString(uuid0)
	defer C.free(unsafe.Pointer(uuid1))
	ret1 := C.g_volume_monitor_get_volume_for_uuid(this1, uuid1)
	var ret2 *Volume
	ret2 = (*Volume)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *VolumeMonitor) GetVolumes() []*Volume {
	var this1 *C.GVolumeMonitor
	if this0 != nil {
		this1 = (*C.GVolumeMonitor)(this0.InheritedFromGVolumeMonitor())
	}
	ret1 := C.g_volume_monitor_get_volumes(this1)
	var ret2 []*Volume
	for iter := (*_GList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt *Volume
		elt = (*Volume)(gobject.ObjectWrap(unsafe.Pointer((*C.GVolume)(iter.data)), false))
		ret2 = append(ret2, elt)
	}
	return ret2
}
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
// blacklisted: action_name_is_valid (function)
// blacklisted: action_parse_detailed_name (function)
// blacklisted: action_print_detailed_name (function)
// blacklisted: app_info_create_from_commandline (function)
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
		launch_context1 = (*C.GAppLaunchContext)(launch_context0.InheritedFromGAppLaunchContext())
	}
	ret1 := C.g_app_info_launch_default_for_uri(uri1, launch_context1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
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
func FileNewForCommandlineArg(arg0 string) *File {
	var arg1 *C.char
	arg1 = _GoStringToGString(arg0)
	defer C.free(unsafe.Pointer(arg1))
	ret1 := C.g_file_new_for_commandline_arg(arg1)
	var ret2 *File
	ret2 = (*File)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func FileNewForCommandlineArgAndCwd(arg0 string, cwd0 string) *File {
	var arg1 *C.char
	var cwd1 *C.char
	arg1 = _GoStringToGString(arg0)
	defer C.free(unsafe.Pointer(arg1))
	cwd1 = _GoStringToGString(cwd0)
	defer C.free(unsafe.Pointer(cwd1))
	ret1 := C.g_file_new_for_commandline_arg_and_cwd(arg1, cwd1)
	var ret2 *File
	ret2 = (*File)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func FileNewForPath(path0 string) *File {
	var path1 *C.char
	path1 = _GoStringToGString(path0)
	defer C.free(unsafe.Pointer(path1))
	ret1 := C.g_file_new_for_path(path1)
	var ret2 *File
	ret2 = (*File)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func FileNewForUri(uri0 string) *File {
	var uri1 *C.char
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	ret1 := C.g_file_new_for_uri(uri1)
	var ret2 *File
	ret2 = (*File)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func FileNewTmp(tmpl0 string) (*FileIOStream, *File, error) {
	var tmpl1 *C.char
	var iostream1 *C.GFileIOStream
	var err1 *C.GError
	tmpl1 = _GoStringToGString(tmpl0)
	defer C.free(unsafe.Pointer(tmpl1))
	ret1 := C.g_file_new_tmp(tmpl1, &iostream1, &err1)
	var iostream2 *FileIOStream
	var ret2 *File
	var err2 error
	iostream2 = (*FileIOStream)(gobject.ObjectWrap(unsafe.Pointer(iostream1), false))
	ret2 = (*File)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return iostream2, ret2, err2
}
func FileParseName(parse_name0 string) *File {
	var parse_name1 *C.char
	parse_name1 = _GoStringToGString(parse_name0)
	defer C.free(unsafe.Pointer(parse_name1))
	ret1 := C.g_file_parse_name(parse_name1)
	var ret2 *File
	ret2 = (*File)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
// blacklisted: icon_deserialize (function)
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
// blacklisted: resources_register (function)
// blacklisted: resources_unregister (function)
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
