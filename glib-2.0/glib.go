package glib

/*
#include "glib.gen.h"
extern void g_key_file_free(GKeyFile*);
#cgo pkg-config: glib-2.0
*/
import "C"
import "unsafe"

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


const AnalyzerAnalyzing = 1
const AsciiDtostrBufSize = 39
// blacklisted: Array (struct)
type AsciiType C.uint32_t
const (
	AsciiTypeAlnum AsciiType = 1
	AsciiTypeAlpha AsciiType = 2
	AsciiTypeCntrl AsciiType = 4
	AsciiTypeDigit AsciiType = 8
	AsciiTypeGraph AsciiType = 16
	AsciiTypeLower AsciiType = 32
	AsciiTypePrint AsciiType = 64
	AsciiTypePunct AsciiType = 128
	AsciiTypeSpace AsciiType = 256
	AsciiTypeUpper AsciiType = 512
	AsciiTypeXdigit AsciiType = 1024
)
// blacklisted: AsyncQueue (struct)
const BigEndian = 4321
type BookmarkFile struct {}
func (this0 *BookmarkFile) AddApplication(uri0 string, name0 string, exec0 string) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var name1 *C.char
	var exec1 *C.char
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	exec1 = _GoStringToGString(exec0)
	defer C.free(unsafe.Pointer(exec1))
	C.g_bookmark_file_add_application(this1, uri1, name1, exec1)
}
func (this0 *BookmarkFile) AddGroup(uri0 string, group0 string) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var group1 *C.char
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	group1 = _GoStringToGString(group0)
	defer C.free(unsafe.Pointer(group1))
	C.g_bookmark_file_add_group(this1, uri1, group1)
}
func (this0 *BookmarkFile) Free() {
	var this1 *C.GBookmarkFile
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	C.g_bookmark_file_free(this1)
}
func (this0 *BookmarkFile) GetAdded(uri0 string) (int64, error) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var err1 *C.GError
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	ret1 := C.g_bookmark_file_get_added(this1, uri1, &err1)
	var ret2 int64
	var err2 error
	ret2 = int64(ret1)
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *BookmarkFile) GetAppInfo(uri0 string, name0 string) (string, uint32, int64, bool, error) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var name1 *C.char
	var exec1 *C.char
	var count1 C.uint32_t
	var stamp1 C.int64_t
	var err1 *C.GError
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	ret1 := C.g_bookmark_file_get_app_info(this1, uri1, name1, &exec1, &count1, &stamp1, &err1)
	var exec2 string
	var count2 uint32
	var stamp2 int64
	var ret2 bool
	var err2 error
	exec2 = C.GoString(exec1)
	C.g_free(unsafe.Pointer(exec1))
	count2 = uint32(count1)
	stamp2 = int64(stamp1)
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return exec2, count2, stamp2, ret2, err2
}
func (this0 *BookmarkFile) GetApplications(uri0 string) (uint64, []string, error) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var length1 C.uint64_t
	var err1 *C.GError
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	ret1 := C.g_bookmark_file_get_applications(this1, uri1, &length1, &err1)
	var length2 uint64
	var ret2 []string
	var err2 error
	length2 = uint64(length1)
	ret2 = make([]string, length1)
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return length2, ret2, err2
}
func (this0 *BookmarkFile) GetDescription(uri0 string) (string, error) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var err1 *C.GError
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	ret1 := C.g_bookmark_file_get_description(this1, uri1, &err1)
	var ret2 string
	var err2 error
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *BookmarkFile) GetGroups(uri0 string) (uint64, []string, error) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var length1 C.uint64_t
	var err1 *C.GError
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	ret1 := C.g_bookmark_file_get_groups(this1, uri1, &length1, &err1)
	var length2 uint64
	var ret2 []string
	var err2 error
	length2 = uint64(length1)
	ret2 = make([]string, length1)
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return length2, ret2, err2
}
func (this0 *BookmarkFile) GetIcon(uri0 string) (string, string, bool, error) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var href1 *C.char
	var mime_type1 *C.char
	var err1 *C.GError
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	ret1 := C.g_bookmark_file_get_icon(this1, uri1, &href1, &mime_type1, &err1)
	var href2 string
	var mime_type2 string
	var ret2 bool
	var err2 error
	href2 = C.GoString(href1)
	C.g_free(unsafe.Pointer(href1))
	mime_type2 = C.GoString(mime_type1)
	C.g_free(unsafe.Pointer(mime_type1))
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return href2, mime_type2, ret2, err2
}
func (this0 *BookmarkFile) GetIsPrivate(uri0 string) (bool, error) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var err1 *C.GError
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	ret1 := C.g_bookmark_file_get_is_private(this1, uri1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *BookmarkFile) GetMimeType(uri0 string) (string, error) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var err1 *C.GError
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	ret1 := C.g_bookmark_file_get_mime_type(this1, uri1, &err1)
	var ret2 string
	var err2 error
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *BookmarkFile) GetModified(uri0 string) (int64, error) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var err1 *C.GError
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	ret1 := C.g_bookmark_file_get_modified(this1, uri1, &err1)
	var ret2 int64
	var err2 error
	ret2 = int64(ret1)
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *BookmarkFile) GetSize() int32 {
	var this1 *C.GBookmarkFile
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	ret1 := C.g_bookmark_file_get_size(this1)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func (this0 *BookmarkFile) GetTitle(uri0 string) (string, error) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var err1 *C.GError
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	ret1 := C.g_bookmark_file_get_title(this1, uri1, &err1)
	var ret2 string
	var err2 error
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *BookmarkFile) GetUris() (uint64, []string) {
	var this1 *C.GBookmarkFile
	var length1 C.uint64_t
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	ret1 := C.g_bookmark_file_get_uris(this1, &length1)
	var length2 uint64
	var ret2 []string
	length2 = uint64(length1)
	ret2 = make([]string, length1)
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return length2, ret2
}
func (this0 *BookmarkFile) GetVisited(uri0 string) (int64, error) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var err1 *C.GError
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	ret1 := C.g_bookmark_file_get_visited(this1, uri1, &err1)
	var ret2 int64
	var err2 error
	ret2 = int64(ret1)
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *BookmarkFile) HasApplication(uri0 string, name0 string) (bool, error) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var name1 *C.char
	var err1 *C.GError
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	ret1 := C.g_bookmark_file_has_application(this1, uri1, name1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *BookmarkFile) HasGroup(uri0 string, group0 string) (bool, error) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var group1 *C.char
	var err1 *C.GError
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	group1 = _GoStringToGString(group0)
	defer C.free(unsafe.Pointer(group1))
	ret1 := C.g_bookmark_file_has_group(this1, uri1, group1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *BookmarkFile) HasItem(uri0 string) bool {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	ret1 := C.g_bookmark_file_has_item(this1, uri1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *BookmarkFile) LoadFromData(data0 string, length0 uint64) (bool, error) {
	var this1 *C.GBookmarkFile
	var data1 *C.char
	var length1 C.uint64_t
	var err1 *C.GError
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	data1 = _GoStringToGString(data0)
	defer C.free(unsafe.Pointer(data1))
	length1 = C.uint64_t(length0)
	ret1 := C.g_bookmark_file_load_from_data(this1, data1, length1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *BookmarkFile) LoadFromDataDirs(file0 string, full_path0 string) (bool, error) {
	var this1 *C.GBookmarkFile
	var file1 *C.char
	var full_path1 *C.char
	var err1 *C.GError
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	file1 = _GoStringToGString(file0)
	defer C.free(unsafe.Pointer(file1))
	full_path1 = _GoStringToGString(full_path0)
	defer C.free(unsafe.Pointer(full_path1))
	ret1 := C.g_bookmark_file_load_from_data_dirs(this1, file1, full_path1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *BookmarkFile) LoadFromFile(filename0 string) (bool, error) {
	var this1 *C.GBookmarkFile
	var filename1 *C.char
	var err1 *C.GError
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	filename1 = _GoStringToGString(filename0)
	defer C.free(unsafe.Pointer(filename1))
	ret1 := C.g_bookmark_file_load_from_file(this1, filename1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *BookmarkFile) MoveItem(old_uri0 string, new_uri0 string) (bool, error) {
	var this1 *C.GBookmarkFile
	var old_uri1 *C.char
	var new_uri1 *C.char
	var err1 *C.GError
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	old_uri1 = _GoStringToGString(old_uri0)
	defer C.free(unsafe.Pointer(old_uri1))
	new_uri1 = _GoStringToGString(new_uri0)
	defer C.free(unsafe.Pointer(new_uri1))
	ret1 := C.g_bookmark_file_move_item(this1, old_uri1, new_uri1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *BookmarkFile) RemoveApplication(uri0 string, name0 string) (bool, error) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var name1 *C.char
	var err1 *C.GError
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	ret1 := C.g_bookmark_file_remove_application(this1, uri1, name1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *BookmarkFile) RemoveGroup(uri0 string, group0 string) (bool, error) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var group1 *C.char
	var err1 *C.GError
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	group1 = _GoStringToGString(group0)
	defer C.free(unsafe.Pointer(group1))
	ret1 := C.g_bookmark_file_remove_group(this1, uri1, group1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *BookmarkFile) RemoveItem(uri0 string) (bool, error) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var err1 *C.GError
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	ret1 := C.g_bookmark_file_remove_item(this1, uri1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *BookmarkFile) SetAdded(uri0 string, added0 int64) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var added1 C.int64_t
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	added1 = C.int64_t(added0)
	C.g_bookmark_file_set_added(this1, uri1, added1)
}
func (this0 *BookmarkFile) SetAppInfo(uri0 string, name0 string, exec0 string, count0 int32, stamp0 int64) (bool, error) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var name1 *C.char
	var exec1 *C.char
	var count1 C.int32_t
	var stamp1 C.int64_t
	var err1 *C.GError
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	exec1 = _GoStringToGString(exec0)
	defer C.free(unsafe.Pointer(exec1))
	count1 = C.int32_t(count0)
	stamp1 = C.int64_t(stamp0)
	ret1 := C.g_bookmark_file_set_app_info(this1, uri1, name1, exec1, count1, stamp1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *BookmarkFile) SetDescription(uri0 string, description0 string) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var description1 *C.char
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	description1 = _GoStringToGString(description0)
	defer C.free(unsafe.Pointer(description1))
	C.g_bookmark_file_set_description(this1, uri1, description1)
}
func (this0 *BookmarkFile) SetGroups(uri0 string, groups0 string, length0 uint64) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var groups1 *C.char
	var length1 C.uint64_t
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	groups1 = _GoStringToGString(groups0)
	defer C.free(unsafe.Pointer(groups1))
	length1 = C.uint64_t(length0)
	C.g_bookmark_file_set_groups(this1, uri1, groups1, length1)
}
func (this0 *BookmarkFile) SetIcon(uri0 string, href0 string, mime_type0 string) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var href1 *C.char
	var mime_type1 *C.char
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	href1 = _GoStringToGString(href0)
	defer C.free(unsafe.Pointer(href1))
	mime_type1 = _GoStringToGString(mime_type0)
	defer C.free(unsafe.Pointer(mime_type1))
	C.g_bookmark_file_set_icon(this1, uri1, href1, mime_type1)
}
func (this0 *BookmarkFile) SetIsPrivate(uri0 string, is_private0 bool) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var is_private1 C.int
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	is_private1 = _GoBoolToCBool(is_private0)
	C.g_bookmark_file_set_is_private(this1, uri1, is_private1)
}
func (this0 *BookmarkFile) SetMimeType(uri0 string, mime_type0 string) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var mime_type1 *C.char
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	mime_type1 = _GoStringToGString(mime_type0)
	defer C.free(unsafe.Pointer(mime_type1))
	C.g_bookmark_file_set_mime_type(this1, uri1, mime_type1)
}
func (this0 *BookmarkFile) SetModified(uri0 string, modified0 int64) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var modified1 C.int64_t
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	modified1 = C.int64_t(modified0)
	C.g_bookmark_file_set_modified(this1, uri1, modified1)
}
func (this0 *BookmarkFile) SetTitle(uri0 string, title0 string) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var title1 *C.char
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	title1 = _GoStringToGString(title0)
	defer C.free(unsafe.Pointer(title1))
	C.g_bookmark_file_set_title(this1, uri1, title1)
}
func (this0 *BookmarkFile) SetVisited(uri0 string, visited0 int64) {
	var this1 *C.GBookmarkFile
	var uri1 *C.char
	var visited1 C.int64_t
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	uri1 = _GoStringToGString(uri0)
	defer C.free(unsafe.Pointer(uri1))
	visited1 = C.int64_t(visited0)
	C.g_bookmark_file_set_visited(this1, uri1, visited1)
}
func (this0 *BookmarkFile) ToData() (uint64, string, error) {
	var this1 *C.GBookmarkFile
	var length1 C.uint64_t
	var err1 *C.GError
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	ret1 := C.g_bookmark_file_to_data(this1, &length1, &err1)
	var length2 uint64
	var ret2 string
	var err2 error
	length2 = uint64(length1)
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return length2, ret2, err2
}
func (this0 *BookmarkFile) ToFile(filename0 string) (bool, error) {
	var this1 *C.GBookmarkFile
	var filename1 *C.char
	var err1 *C.GError
	this1 = (*C.GBookmarkFile)(unsafe.Pointer(this0))
	filename1 = _GoStringToGString(filename0)
	defer C.free(unsafe.Pointer(filename1))
	ret1 := C.g_bookmark_file_to_file(this1, filename1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func BookmarkFileErrorQuark() uint32 {
	ret1 := C.g_bookmark_file_error_quark()
	var ret2 uint32
	ret2 = uint32(ret1)
	return ret2
}
type BookmarkFileError C.uint32_t
const (
	BookmarkFileErrorInvalidUri BookmarkFileError = 0
	BookmarkFileErrorInvalidValue BookmarkFileError = 1
	BookmarkFileErrorAppNotRegistered BookmarkFileError = 2
	BookmarkFileErrorUriNotFound BookmarkFileError = 3
	BookmarkFileErrorRead BookmarkFileError = 4
	BookmarkFileErrorUnknownEncoding BookmarkFileError = 5
	BookmarkFileErrorWrite BookmarkFileError = 6
	BookmarkFileErrorFileNotFound BookmarkFileError = 7
)
// blacklisted: ByteArray (struct)
type Bytes struct {}
func NewBytes(data0 []uint8) *Bytes {
	var data1 *C.uint8_t
	var size1 C.uint64_t
	data1 = (*C.uint8_t)(C.malloc(C.size_t(int(unsafe.Sizeof(*data1)) * len(data0))))
	defer C.free(unsafe.Pointer(data1))
	for i, e := range data0 {
		(*(*[999999]C.uint8_t)(unsafe.Pointer(data1)))[i] = C.uint8_t(e)
	}
	size1 = C.uint64_t(len(data0))
	ret1 := C.g_bytes_new(data1, size1)
	var ret2 *Bytes
	ret2 = (*Bytes)(unsafe.Pointer(ret1))
	return ret2
}
func NewBytesTake(data0 []uint8) *Bytes {
	var data1 *C.uint8_t
	var size1 C.uint64_t
	data1 = (*C.uint8_t)(C.malloc(C.size_t(int(unsafe.Sizeof(*data1)) * len(data0))))
	defer C.free(unsafe.Pointer(data1))
	for i, e := range data0 {
		(*(*[999999]C.uint8_t)(unsafe.Pointer(data1)))[i] = C.uint8_t(e)
	}
	size1 = C.uint64_t(len(data0))
	ret1 := C.g_bytes_new_take(data1, size1)
	var ret2 *Bytes
	ret2 = (*Bytes)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Bytes) Compare(bytes20 *Bytes) int32 {
	var this1 *C.GBytes
	var bytes21 *C.GBytes
	this1 = (*C.GBytes)(unsafe.Pointer(this0))
	bytes21 = (*C.GBytes)(unsafe.Pointer(bytes20))
	ret1 := C.g_bytes_compare(this1, bytes21)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func (this0 *Bytes) Equal(bytes20 *Bytes) bool {
	var this1 *C.GBytes
	var bytes21 *C.GBytes
	this1 = (*C.GBytes)(unsafe.Pointer(this0))
	bytes21 = (*C.GBytes)(unsafe.Pointer(bytes20))
	ret1 := C.g_bytes_equal(this1, bytes21)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Bytes) GetSize() uint64 {
	var this1 *C.GBytes
	this1 = (*C.GBytes)(unsafe.Pointer(this0))
	ret1 := C.g_bytes_get_size(this1)
	var ret2 uint64
	ret2 = uint64(ret1)
	return ret2
}
func (this0 *Bytes) Hash() uint32 {
	var this1 *C.GBytes
	this1 = (*C.GBytes)(unsafe.Pointer(this0))
	ret1 := C.g_bytes_hash(this1)
	var ret2 uint32
	ret2 = uint32(ret1)
	return ret2
}
func (this0 *Bytes) NewFromBytes(offset0 uint64, length0 uint64) *Bytes {
	var this1 *C.GBytes
	var offset1 C.uint64_t
	var length1 C.uint64_t
	this1 = (*C.GBytes)(unsafe.Pointer(this0))
	offset1 = C.uint64_t(offset0)
	length1 = C.uint64_t(length0)
	ret1 := C.g_bytes_new_from_bytes(this1, offset1, length1)
	var ret2 *Bytes
	ret2 = (*Bytes)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Bytes) UnrefToData(size0 *uint64) {
	var this1 *C.GBytes
	var size1 *C.uint64_t
	this1 = (*C.GBytes)(unsafe.Pointer(this0))
	size1 = (*C.uint64_t)(unsafe.Pointer(size0))
	C.g_bytes_unref_to_data(this1, size1)
}
const CanInline = 1
// blacklisted: CSET_A_2_Z (constant)
const CsetDigits = "0123456789"
// blacklisted: CSET_a_2_z (constant)
// blacklisted: Checksum (struct)
type ChecksumType C.uint32_t
const (
	ChecksumTypeMd5 ChecksumType = 0
	ChecksumTypeSha1 ChecksumType = 1
	ChecksumTypeSha256 ChecksumType = 2
	ChecksumTypeSha512 ChecksumType = 3
)
// blacklisted: ChildWatchFunc (callback)
// blacklisted: CompareDataFunc (callback)
// blacklisted: CompareFunc (callback)
// blacklisted: Cond (struct)
type ConvertError C.uint32_t
const (
	ConvertErrorNoConversion ConvertError = 0
	ConvertErrorIllegalSequence ConvertError = 1
	ConvertErrorFailed ConvertError = 2
	ConvertErrorPartialInput ConvertError = 3
	ConvertErrorBadUri ConvertError = 4
	ConvertErrorNotAbsolutePath ConvertError = 5
	ConvertErrorNoMemory ConvertError = 6
)
const DatalistFlagsMask = 3
const DateBadDay = 0
const DateBadJulian = 0
const DateBadYear = 0
const DirSeparator = 92
const DirSeparatorS = "\\"
// blacklisted: Data (struct)
// blacklisted: DataForeachFunc (callback)
// blacklisted: Date (struct)
type DateDMY C.uint32_t
const (
	DateDMYDay DateDMY = 0
	DateDMYMonth DateDMY = 1
	DateDMYYear DateDMY = 2
)
type DateMonth C.uint32_t
const (
	DateMonthBadMonth DateMonth = 0
	DateMonthJanuary DateMonth = 1
	DateMonthFebruary DateMonth = 2
	DateMonthMarch DateMonth = 3
	DateMonthApril DateMonth = 4
	DateMonthMay DateMonth = 5
	DateMonthJune DateMonth = 6
	DateMonthJuly DateMonth = 7
	DateMonthAugust DateMonth = 8
	DateMonthSeptember DateMonth = 9
	DateMonthOctober DateMonth = 10
	DateMonthNovember DateMonth = 11
	DateMonthDecember DateMonth = 12
)
type DateTime struct {}
func NewDateTime(tz0 *TimeZone, year0 int32, month0 int32, day0 int32, hour0 int32, minute0 int32, seconds0 float64) *DateTime {
	var tz1 *C.GTimeZone
	var year1 C.int32_t
	var month1 C.int32_t
	var day1 C.int32_t
	var hour1 C.int32_t
	var minute1 C.int32_t
	var seconds1 C.double
	tz1 = (*C.GTimeZone)(unsafe.Pointer(tz0))
	year1 = C.int32_t(year0)
	month1 = C.int32_t(month0)
	day1 = C.int32_t(day0)
	hour1 = C.int32_t(hour0)
	minute1 = C.int32_t(minute0)
	seconds1 = C.double(seconds0)
	ret1 := C.g_date_time_new(tz1, year1, month1, day1, hour1, minute1, seconds1)
	var ret2 *DateTime
	ret2 = (*DateTime)(unsafe.Pointer(ret1))
	return ret2
}
func NewDateTimeFromTimevalLocal(tv0 *TimeVal) *DateTime {
	var tv1 *C.GTimeVal
	tv1 = (*C.GTimeVal)(unsafe.Pointer(tv0))
	ret1 := C.g_date_time_new_from_timeval_local(tv1)
	var ret2 *DateTime
	ret2 = (*DateTime)(unsafe.Pointer(ret1))
	return ret2
}
func NewDateTimeFromTimevalUtc(tv0 *TimeVal) *DateTime {
	var tv1 *C.GTimeVal
	tv1 = (*C.GTimeVal)(unsafe.Pointer(tv0))
	ret1 := C.g_date_time_new_from_timeval_utc(tv1)
	var ret2 *DateTime
	ret2 = (*DateTime)(unsafe.Pointer(ret1))
	return ret2
}
func NewDateTimeFromUnixLocal(t0 int64) *DateTime {
	var t1 C.int64_t
	t1 = C.int64_t(t0)
	ret1 := C.g_date_time_new_from_unix_local(t1)
	var ret2 *DateTime
	ret2 = (*DateTime)(unsafe.Pointer(ret1))
	return ret2
}
func NewDateTimeFromUnixUtc(t0 int64) *DateTime {
	var t1 C.int64_t
	t1 = C.int64_t(t0)
	ret1 := C.g_date_time_new_from_unix_utc(t1)
	var ret2 *DateTime
	ret2 = (*DateTime)(unsafe.Pointer(ret1))
	return ret2
}
func NewDateTimeLocal(year0 int32, month0 int32, day0 int32, hour0 int32, minute0 int32, seconds0 float64) *DateTime {
	var year1 C.int32_t
	var month1 C.int32_t
	var day1 C.int32_t
	var hour1 C.int32_t
	var minute1 C.int32_t
	var seconds1 C.double
	year1 = C.int32_t(year0)
	month1 = C.int32_t(month0)
	day1 = C.int32_t(day0)
	hour1 = C.int32_t(hour0)
	minute1 = C.int32_t(minute0)
	seconds1 = C.double(seconds0)
	ret1 := C.g_date_time_new_local(year1, month1, day1, hour1, minute1, seconds1)
	var ret2 *DateTime
	ret2 = (*DateTime)(unsafe.Pointer(ret1))
	return ret2
}
func NewDateTimeNow(tz0 *TimeZone) *DateTime {
	var tz1 *C.GTimeZone
	tz1 = (*C.GTimeZone)(unsafe.Pointer(tz0))
	ret1 := C.g_date_time_new_now(tz1)
	var ret2 *DateTime
	ret2 = (*DateTime)(unsafe.Pointer(ret1))
	return ret2
}
func NewDateTimeNowLocal() *DateTime {
	ret1 := C.g_date_time_new_now_local()
	var ret2 *DateTime
	ret2 = (*DateTime)(unsafe.Pointer(ret1))
	return ret2
}
func NewDateTimeNowUtc() *DateTime {
	ret1 := C.g_date_time_new_now_utc()
	var ret2 *DateTime
	ret2 = (*DateTime)(unsafe.Pointer(ret1))
	return ret2
}
func NewDateTimeUtc(year0 int32, month0 int32, day0 int32, hour0 int32, minute0 int32, seconds0 float64) *DateTime {
	var year1 C.int32_t
	var month1 C.int32_t
	var day1 C.int32_t
	var hour1 C.int32_t
	var minute1 C.int32_t
	var seconds1 C.double
	year1 = C.int32_t(year0)
	month1 = C.int32_t(month0)
	day1 = C.int32_t(day0)
	hour1 = C.int32_t(hour0)
	minute1 = C.int32_t(minute0)
	seconds1 = C.double(seconds0)
	ret1 := C.g_date_time_new_utc(year1, month1, day1, hour1, minute1, seconds1)
	var ret2 *DateTime
	ret2 = (*DateTime)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *DateTime) Add(timespan0 int64) *DateTime {
	var this1 *C.GDateTime
	var timespan1 C.int64_t
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	timespan1 = C.int64_t(timespan0)
	ret1 := C.g_date_time_add(this1, timespan1)
	var ret2 *DateTime
	ret2 = (*DateTime)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *DateTime) AddDays(days0 int32) *DateTime {
	var this1 *C.GDateTime
	var days1 C.int32_t
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	days1 = C.int32_t(days0)
	ret1 := C.g_date_time_add_days(this1, days1)
	var ret2 *DateTime
	ret2 = (*DateTime)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *DateTime) AddFull(years0 int32, months0 int32, days0 int32, hours0 int32, minutes0 int32, seconds0 float64) *DateTime {
	var this1 *C.GDateTime
	var years1 C.int32_t
	var months1 C.int32_t
	var days1 C.int32_t
	var hours1 C.int32_t
	var minutes1 C.int32_t
	var seconds1 C.double
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	years1 = C.int32_t(years0)
	months1 = C.int32_t(months0)
	days1 = C.int32_t(days0)
	hours1 = C.int32_t(hours0)
	minutes1 = C.int32_t(minutes0)
	seconds1 = C.double(seconds0)
	ret1 := C.g_date_time_add_full(this1, years1, months1, days1, hours1, minutes1, seconds1)
	var ret2 *DateTime
	ret2 = (*DateTime)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *DateTime) AddHours(hours0 int32) *DateTime {
	var this1 *C.GDateTime
	var hours1 C.int32_t
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	hours1 = C.int32_t(hours0)
	ret1 := C.g_date_time_add_hours(this1, hours1)
	var ret2 *DateTime
	ret2 = (*DateTime)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *DateTime) AddMinutes(minutes0 int32) *DateTime {
	var this1 *C.GDateTime
	var minutes1 C.int32_t
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	minutes1 = C.int32_t(minutes0)
	ret1 := C.g_date_time_add_minutes(this1, minutes1)
	var ret2 *DateTime
	ret2 = (*DateTime)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *DateTime) AddMonths(months0 int32) *DateTime {
	var this1 *C.GDateTime
	var months1 C.int32_t
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	months1 = C.int32_t(months0)
	ret1 := C.g_date_time_add_months(this1, months1)
	var ret2 *DateTime
	ret2 = (*DateTime)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *DateTime) AddSeconds(seconds0 float64) *DateTime {
	var this1 *C.GDateTime
	var seconds1 C.double
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	seconds1 = C.double(seconds0)
	ret1 := C.g_date_time_add_seconds(this1, seconds1)
	var ret2 *DateTime
	ret2 = (*DateTime)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *DateTime) AddWeeks(weeks0 int32) *DateTime {
	var this1 *C.GDateTime
	var weeks1 C.int32_t
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	weeks1 = C.int32_t(weeks0)
	ret1 := C.g_date_time_add_weeks(this1, weeks1)
	var ret2 *DateTime
	ret2 = (*DateTime)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *DateTime) AddYears(years0 int32) *DateTime {
	var this1 *C.GDateTime
	var years1 C.int32_t
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	years1 = C.int32_t(years0)
	ret1 := C.g_date_time_add_years(this1, years1)
	var ret2 *DateTime
	ret2 = (*DateTime)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *DateTime) Difference(begin0 *DateTime) int64 {
	var this1 *C.GDateTime
	var begin1 *C.GDateTime
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	begin1 = (*C.GDateTime)(unsafe.Pointer(begin0))
	ret1 := C.g_date_time_difference(this1, begin1)
	var ret2 int64
	ret2 = int64(ret1)
	return ret2
}
func (this0 *DateTime) Format(format0 string) string {
	var this1 *C.GDateTime
	var format1 *C.char
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	format1 = _GoStringToGString(format0)
	defer C.free(unsafe.Pointer(format1))
	ret1 := C.g_date_time_format(this1, format1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *DateTime) GetDayOfMonth() int32 {
	var this1 *C.GDateTime
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	ret1 := C.g_date_time_get_day_of_month(this1)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func (this0 *DateTime) GetDayOfWeek() int32 {
	var this1 *C.GDateTime
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	ret1 := C.g_date_time_get_day_of_week(this1)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func (this0 *DateTime) GetDayOfYear() int32 {
	var this1 *C.GDateTime
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	ret1 := C.g_date_time_get_day_of_year(this1)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func (this0 *DateTime) GetHour() int32 {
	var this1 *C.GDateTime
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	ret1 := C.g_date_time_get_hour(this1)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func (this0 *DateTime) GetMicrosecond() int32 {
	var this1 *C.GDateTime
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	ret1 := C.g_date_time_get_microsecond(this1)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func (this0 *DateTime) GetMinute() int32 {
	var this1 *C.GDateTime
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	ret1 := C.g_date_time_get_minute(this1)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func (this0 *DateTime) GetMonth() int32 {
	var this1 *C.GDateTime
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	ret1 := C.g_date_time_get_month(this1)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func (this0 *DateTime) GetSecond() int32 {
	var this1 *C.GDateTime
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	ret1 := C.g_date_time_get_second(this1)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func (this0 *DateTime) GetSeconds() float64 {
	var this1 *C.GDateTime
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	ret1 := C.g_date_time_get_seconds(this1)
	var ret2 float64
	ret2 = float64(ret1)
	return ret2
}
func (this0 *DateTime) GetTimezoneAbbreviation() string {
	var this1 *C.GDateTime
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	ret1 := C.g_date_time_get_timezone_abbreviation(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *DateTime) GetUtcOffset() int64 {
	var this1 *C.GDateTime
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	ret1 := C.g_date_time_get_utc_offset(this1)
	var ret2 int64
	ret2 = int64(ret1)
	return ret2
}
func (this0 *DateTime) GetWeekNumberingYear() int32 {
	var this1 *C.GDateTime
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	ret1 := C.g_date_time_get_week_numbering_year(this1)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func (this0 *DateTime) GetWeekOfYear() int32 {
	var this1 *C.GDateTime
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	ret1 := C.g_date_time_get_week_of_year(this1)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func (this0 *DateTime) GetYear() int32 {
	var this1 *C.GDateTime
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	ret1 := C.g_date_time_get_year(this1)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func (this0 *DateTime) GetYmd() (int32, int32, int32) {
	var this1 *C.GDateTime
	var year1 C.int32_t
	var month1 C.int32_t
	var day1 C.int32_t
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	C.g_date_time_get_ymd(this1, &year1, &month1, &day1)
	var year2 int32
	var month2 int32
	var day2 int32
	year2 = int32(year1)
	month2 = int32(month1)
	day2 = int32(day1)
	return year2, month2, day2
}
func (this0 *DateTime) IsDaylightSavings() bool {
	var this1 *C.GDateTime
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	ret1 := C.g_date_time_is_daylight_savings(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *DateTime) ToLocal() *DateTime {
	var this1 *C.GDateTime
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	ret1 := C.g_date_time_to_local(this1)
	var ret2 *DateTime
	ret2 = (*DateTime)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *DateTime) ToTimeval(tv0 *TimeVal) bool {
	var this1 *C.GDateTime
	var tv1 *C.GTimeVal
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	tv1 = (*C.GTimeVal)(unsafe.Pointer(tv0))
	ret1 := C.g_date_time_to_timeval(this1, tv1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *DateTime) ToTimezone(tz0 *TimeZone) *DateTime {
	var this1 *C.GDateTime
	var tz1 *C.GTimeZone
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	tz1 = (*C.GTimeZone)(unsafe.Pointer(tz0))
	ret1 := C.g_date_time_to_timezone(this1, tz1)
	var ret2 *DateTime
	ret2 = (*DateTime)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *DateTime) ToUnix() int64 {
	var this1 *C.GDateTime
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	ret1 := C.g_date_time_to_unix(this1)
	var ret2 int64
	ret2 = int64(ret1)
	return ret2
}
func (this0 *DateTime) ToUtc() *DateTime {
	var this1 *C.GDateTime
	this1 = (*C.GDateTime)(unsafe.Pointer(this0))
	ret1 := C.g_date_time_to_utc(this1)
	var ret2 *DateTime
	ret2 = (*DateTime)(unsafe.Pointer(ret1))
	return ret2
}
func DateTimeCompare(dt10 unsafe.Pointer, dt20 unsafe.Pointer) int32 {
	var dt11 unsafe.Pointer
	var dt21 unsafe.Pointer
	dt11 = unsafe.Pointer(dt10)
	dt21 = unsafe.Pointer(dt20)
	ret1 := C.g_date_time_compare(dt11, dt21)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func DateTimeEqual(dt10 unsafe.Pointer, dt20 unsafe.Pointer) bool {
	var dt11 unsafe.Pointer
	var dt21 unsafe.Pointer
	dt11 = unsafe.Pointer(dt10)
	dt21 = unsafe.Pointer(dt20)
	ret1 := C.g_date_time_equal(dt11, dt21)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func DateTimeHash(datetime0 unsafe.Pointer) uint32 {
	var datetime1 unsafe.Pointer
	datetime1 = unsafe.Pointer(datetime0)
	ret1 := C.g_date_time_hash(datetime1)
	var ret2 uint32
	ret2 = uint32(ret1)
	return ret2
}
type DateWeekday C.uint32_t
const (
	DateWeekdayBadWeekday DateWeekday = 0
	DateWeekdayMonday DateWeekday = 1
	DateWeekdayTuesday DateWeekday = 2
	DateWeekdayWednesday DateWeekday = 3
	DateWeekdayThursday DateWeekday = 4
	DateWeekdayFriday DateWeekday = 5
	DateWeekdaySaturday DateWeekday = 6
	DateWeekdaySunday DateWeekday = 7
)
// blacklisted: DebugKey (struct)
// blacklisted: DestroyNotify (callback)
// blacklisted: Dir (struct)
type DoubleIEEE754 struct {
	_data [8]byte
}
const E = 2.718282
// blacklisted: EqualFunc (callback)
// blacklisted: Error (struct)
type ErrorType C.uint32_t
const (
	ErrorTypeUnknown ErrorType = 0
	ErrorTypeUnexpEof ErrorType = 1
	ErrorTypeUnexpEofInString ErrorType = 2
	ErrorTypeUnexpEofInComment ErrorType = 3
	ErrorTypeNonDigitInConst ErrorType = 4
	ErrorTypeDigitRadix ErrorType = 5
	ErrorTypeFloatRadix ErrorType = 6
	ErrorTypeFloatMalformed ErrorType = 7
)
type FileError C.uint32_t
const (
	FileErrorExist FileError = 0
	FileErrorIsdir FileError = 1
	FileErrorAcces FileError = 2
	FileErrorNametoolong FileError = 3
	FileErrorNoent FileError = 4
	FileErrorNotdir FileError = 5
	FileErrorNxio FileError = 6
	FileErrorNodev FileError = 7
	FileErrorRofs FileError = 8
	FileErrorTxtbsy FileError = 9
	FileErrorFault FileError = 10
	FileErrorLoop FileError = 11
	FileErrorNospc FileError = 12
	FileErrorNomem FileError = 13
	FileErrorMfile FileError = 14
	FileErrorNfile FileError = 15
	FileErrorBadf FileError = 16
	FileErrorInval FileError = 17
	FileErrorPipe FileError = 18
	FileErrorAgain FileError = 19
	FileErrorIntr FileError = 20
	FileErrorIo FileError = 21
	FileErrorPerm FileError = 22
	FileErrorNosys FileError = 23
	FileErrorFailed FileError = 24
)
type FileTest C.uint32_t
const (
	FileTestIsRegular FileTest = 1
	FileTestIsSymlink FileTest = 2
	FileTestIsDir FileTest = 4
	FileTestIsExecutable FileTest = 8
	FileTestExists FileTest = 16
)
type FloatIEEE754 struct {
	_data [4]byte
}
type FormatSizeFlags C.uint32_t
const (
	FormatSizeFlagsDefault FormatSizeFlags = 0
	FormatSizeFlagsLongFormat FormatSizeFlags = 1
	FormatSizeFlagsIecUnits FormatSizeFlags = 2
)
// blacklisted: FreeFunc (callback)
// blacklisted: Func (callback)
const Gint16Format = "hi"
const Gint16Modifier = "h"
const Gint32Format = "i"
const Gint32Modifier = ""
const Gint64Format = "li"
const Gint64Modifier = "l"
const GintptrFormat = "li"
const GintptrModifier = "l"
const GnucFunction = ""
const GnucPrettyFunction = ""
const GsizeFormat = "lu"
const GsizeModifier = "l"
const GssizeFormat = "li"
const GssizeModifier = "l"
const Guint16Format = "hu"
const Guint32Format = "u"
const Guint64Format = "lu"
const GuintptrFormat = "lu"
const HaveGint64 = 1
const HaveGnucVarargs = 1
const HaveGnucVisibility = 1
const HaveGrowingStack = 0
const HaveInline = 1
const HaveIsoVarargs = 1
const Have__Inline = 1
const Have__Inline__ = 1
// blacklisted: HFunc (callback)
const HookFlagUserShift = 4
// blacklisted: HRFunc (callback)
// blacklisted: HashFunc (callback)
// blacklisted: HashTable (struct)
// blacklisted: HashTableIter (struct)
// blacklisted: Hmac (struct)
// blacklisted: Hook (struct)
// blacklisted: HookCheckFunc (callback)
// blacklisted: HookCheckMarshaller (callback)
// blacklisted: HookCompareFunc (callback)
// blacklisted: HookFinalizeFunc (callback)
// blacklisted: HookFindFunc (callback)
type HookFlagMask C.uint32_t
const (
	HookFlagMaskActive HookFlagMask = 1
	HookFlagMaskInCall HookFlagMask = 2
	HookFlagMaskMask HookFlagMask = 15
)
// blacklisted: HookFunc (callback)
// blacklisted: HookList (struct)
// blacklisted: HookMarshaller (callback)
// blacklisted: IConv (struct)
const Ieee754DoubleBias = 1023
const Ieee754FloatBias = 127
// blacklisted: IOChannel (struct)
type IOChannelError C.uint32_t
const (
	IOChannelErrorFbig IOChannelError = 0
	IOChannelErrorInval IOChannelError = 1
	IOChannelErrorIo IOChannelError = 2
	IOChannelErrorIsdir IOChannelError = 3
	IOChannelErrorNospc IOChannelError = 4
	IOChannelErrorNxio IOChannelError = 5
	IOChannelErrorOverflow IOChannelError = 6
	IOChannelErrorPipe IOChannelError = 7
	IOChannelErrorFailed IOChannelError = 8
)
type IOCondition C.uint32_t
const (
	IOConditionIn IOCondition = 1
	IOConditionOut IOCondition = 4
	IOConditionPri IOCondition = 2
	IOConditionErr IOCondition = 8
	IOConditionHup IOCondition = 16
	IOConditionNval IOCondition = 32
)
type IOError C.uint32_t
const (
	IOErrorNone IOError = 0
	IOErrorAgain IOError = 1
	IOErrorInval IOError = 2
	IOErrorUnknown IOError = 3
)
type IOFlags C.uint32_t
const (
	IOFlagsAppend IOFlags = 1
	IOFlagsNonblock IOFlags = 2
	IOFlagsIsReadable IOFlags = 4
	IOFlagsIsWritable IOFlags = 8
	IOFlagsIsWriteable IOFlags = 8
	IOFlagsIsSeekable IOFlags = 16
	IOFlagsMask IOFlags = 31
	IOFlagsGetMask IOFlags = 31
	IOFlagsSetMask IOFlags = 3
)
// blacklisted: IOFunc (callback)
// blacklisted: IOFuncs (struct)
type IOStatus C.uint32_t
const (
	IOStatusError IOStatus = 0
	IOStatusNormal IOStatus = 1
	IOStatusEof IOStatus = 2
	IOStatusAgain IOStatus = 3
)
const KeyFileDesktopGroup = "Desktop Entry"
const KeyFileDesktopKeyActions = "Actions"
const KeyFileDesktopKeyCategories = "Categories"
const KeyFileDesktopKeyComment = "Comment"
const KeyFileDesktopKeyDbusActivatable = "DBusActivatable"
const KeyFileDesktopKeyExec = "Exec"
const KeyFileDesktopKeyFullname = "X-GNOME-FullName"
const KeyFileDesktopKeyGenericName = "GenericName"
const KeyFileDesktopKeyGettextDomain = "X-GNOME-Gettext-Domain"
const KeyFileDesktopKeyHidden = "Hidden"
const KeyFileDesktopKeyIcon = "Icon"
const KeyFileDesktopKeyKeywords = "Keywords"
const KeyFileDesktopKeyMimeType = "MimeType"
const KeyFileDesktopKeyName = "Name"
const KeyFileDesktopKeyNotShowIn = "NotShowIn"
const KeyFileDesktopKeyNoDisplay = "NoDisplay"
const KeyFileDesktopKeyOnlyShowIn = "OnlyShowIn"
const KeyFileDesktopKeyPath = "Path"
const KeyFileDesktopKeyStartupNotify = "StartupNotify"
const KeyFileDesktopKeyStartupWmClass = "StartupWMClass"
const KeyFileDesktopKeyTerminal = "Terminal"
const KeyFileDesktopKeyTryExec = "TryExec"
const KeyFileDesktopKeyType = "Type"
const KeyFileDesktopKeyUrl = "URL"
const KeyFileDesktopKeyVersion = "Version"
const KeyFileDesktopTypeApplication = "Application"
const KeyFileDesktopTypeDirectory = "Directory"
const KeyFileDesktopTypeLink = "Link"
type KeyFile struct {}
func NewKeyFile() *KeyFile {
	ret1 := C.g_key_file_new()
	var ret2 *KeyFile
	ret2 = (*KeyFile)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *KeyFile) GetBoolean(group_name0 string, key0 string) (bool, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_get_boolean(this1, group_name1, key1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) GetBooleanList(group_name0 string, key0 string) (uint64, []bool, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var length1 C.uint64_t
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_get_boolean_list(this1, group_name1, key1, &length1, &err1)
	var length2 uint64
	var ret2 []bool
	var err2 error
	length2 = uint64(length1)
	ret2 = make([]bool, length1)
	for i := range ret2 {
		ret2[i] = (*(*[999999]C.int)(unsafe.Pointer(ret1)))[i] != 0
	}
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return length2, ret2, err2
}
func (this0 *KeyFile) GetComment(group_name0 string, key0 string) (string, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_get_comment(this1, group_name1, key1, &err1)
	var ret2 string
	var err2 error
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) GetDouble(group_name0 string, key0 string) (float64, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_get_double(this1, group_name1, key1, &err1)
	var ret2 float64
	var err2 error
	ret2 = float64(ret1)
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) GetDoubleList(group_name0 string, key0 string) (uint64, []float64, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var length1 C.uint64_t
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_get_double_list(this1, group_name1, key1, &length1, &err1)
	var length2 uint64
	var ret2 []float64
	var err2 error
	length2 = uint64(length1)
	ret2 = make([]float64, length1)
	for i := range ret2 {
		ret2[i] = float64((*(*[999999]C.double)(unsafe.Pointer(ret1)))[i])
	}
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return length2, ret2, err2
}
func (this0 *KeyFile) GetGroups() (uint64, []string) {
	var this1 *C.GKeyFile
	var length1 C.uint64_t
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	ret1 := C.g_key_file_get_groups(this1, &length1)
	var length2 uint64
	var ret2 []string
	length2 = uint64(length1)
	ret2 = make([]string, uint(C._array_length(unsafe.Pointer(ret1))))
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return length2, ret2
}
func (this0 *KeyFile) GetInt64(group_name0 string, key0 string) (int64, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_get_int64(this1, group_name1, key1, &err1)
	var ret2 int64
	var err2 error
	ret2 = int64(ret1)
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) GetInteger(group_name0 string, key0 string) (int32, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_get_integer(this1, group_name1, key1, &err1)
	var ret2 int32
	var err2 error
	ret2 = int32(ret1)
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) GetIntegerList(group_name0 string, key0 string) (uint64, []int32, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var length1 C.uint64_t
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_get_integer_list(this1, group_name1, key1, &length1, &err1)
	var length2 uint64
	var ret2 []int32
	var err2 error
	length2 = uint64(length1)
	ret2 = make([]int32, length1)
	for i := range ret2 {
		ret2[i] = int32((*(*[999999]C.int32_t)(unsafe.Pointer(ret1)))[i])
	}
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return length2, ret2, err2
}
func (this0 *KeyFile) GetKeys(group_name0 string) (uint64, []string, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var length1 C.uint64_t
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	ret1 := C.g_key_file_get_keys(this1, group_name1, &length1, &err1)
	var length2 uint64
	var ret2 []string
	var err2 error
	length2 = uint64(length1)
	ret2 = make([]string, uint(C._array_length(unsafe.Pointer(ret1))))
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return length2, ret2, err2
}
func (this0 *KeyFile) GetLocaleString(group_name0 string, key0 string, locale0 string) (string, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var locale1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	locale1 = _GoStringToGString(locale0)
	defer C.free(unsafe.Pointer(locale1))
	ret1 := C.g_key_file_get_locale_string(this1, group_name1, key1, locale1, &err1)
	var ret2 string
	var err2 error
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) GetLocaleStringList(group_name0 string, key0 string, locale0 string) (uint64, []string, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var locale1 *C.char
	var length1 C.uint64_t
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	locale1 = _GoStringToGString(locale0)
	defer C.free(unsafe.Pointer(locale1))
	ret1 := C.g_key_file_get_locale_string_list(this1, group_name1, key1, locale1, &length1, &err1)
	var length2 uint64
	var ret2 []string
	var err2 error
	length2 = uint64(length1)
	ret2 = make([]string, length1)
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return length2, ret2, err2
}
func (this0 *KeyFile) GetStartGroup() string {
	var this1 *C.GKeyFile
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	ret1 := C.g_key_file_get_start_group(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *KeyFile) GetString(group_name0 string, key0 string) (string, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_get_string(this1, group_name1, key1, &err1)
	var ret2 string
	var err2 error
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) GetStringList(group_name0 string, key0 string) (uint64, []string, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var length1 C.uint64_t
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_get_string_list(this1, group_name1, key1, &length1, &err1)
	var length2 uint64
	var ret2 []string
	var err2 error
	length2 = uint64(length1)
	ret2 = make([]string, length1)
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return length2, ret2, err2
}
func (this0 *KeyFile) GetUint64(group_name0 string, key0 string) (uint64, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_get_uint64(this1, group_name1, key1, &err1)
	var ret2 uint64
	var err2 error
	ret2 = uint64(ret1)
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) GetValue(group_name0 string, key0 string) (string, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_get_value(this1, group_name1, key1, &err1)
	var ret2 string
	var err2 error
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) HasGroup(group_name0 string) bool {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	ret1 := C.g_key_file_has_group(this1, group_name1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *KeyFile) LoadFromData(data0 string, length0 uint64, flags0 KeyFileFlags) (bool, error) {
	var this1 *C.GKeyFile
	var data1 *C.char
	var length1 C.uint64_t
	var flags1 C.GKeyFileFlags
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	data1 = _GoStringToGString(data0)
	defer C.free(unsafe.Pointer(data1))
	length1 = C.uint64_t(length0)
	flags1 = C.GKeyFileFlags(flags0)
	ret1 := C.g_key_file_load_from_data(this1, data1, length1, flags1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) LoadFromDataDirs(file0 string, flags0 KeyFileFlags) (string, bool, error) {
	var this1 *C.GKeyFile
	var file1 *C.char
	var flags1 C.GKeyFileFlags
	var full_path1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	file1 = _GoStringToGString(file0)
	defer C.free(unsafe.Pointer(file1))
	flags1 = C.GKeyFileFlags(flags0)
	ret1 := C.g_key_file_load_from_data_dirs(this1, file1, &full_path1, flags1, &err1)
	var full_path2 string
	var ret2 bool
	var err2 error
	full_path2 = C.GoString(full_path1)
	C.g_free(unsafe.Pointer(full_path1))
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return full_path2, ret2, err2
}
func (this0 *KeyFile) LoadFromDirs(file0 string, search_dirs0 []string, flags0 KeyFileFlags) (string, bool, error) {
	var this1 *C.GKeyFile
	var file1 *C.char
	var search_dirs1 **C.char
	var flags1 C.GKeyFileFlags
	var full_path1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	file1 = _GoStringToGString(file0)
	defer C.free(unsafe.Pointer(file1))
	search_dirs1 = (**C.char)(C.malloc(C.size_t(int(unsafe.Sizeof(*search_dirs1)) * (len(search_dirs0) + 1))))
	defer C.free(unsafe.Pointer(search_dirs1))
	for i, e := range search_dirs0 {
		(*(*[999999]*C.char)(unsafe.Pointer(search_dirs1)))[i] = _GoStringToGString(e)
		defer C.free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(search_dirs1)))[i]))
	}
	(*(*[999999]*C.char)(unsafe.Pointer(search_dirs1)))[len(search_dirs0)] = nil
	flags1 = C.GKeyFileFlags(flags0)
	ret1 := C.g_key_file_load_from_dirs(this1, file1, search_dirs1, &full_path1, flags1, &err1)
	var full_path2 string
	var ret2 bool
	var err2 error
	full_path2 = C.GoString(full_path1)
	C.g_free(unsafe.Pointer(full_path1))
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return full_path2, ret2, err2
}
func (this0 *KeyFile) LoadFromFile(file0 string, flags0 KeyFileFlags) (bool, error) {
	var this1 *C.GKeyFile
	var file1 *C.char
	var flags1 C.GKeyFileFlags
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	file1 = _GoStringToGString(file0)
	defer C.free(unsafe.Pointer(file1))
	flags1 = C.GKeyFileFlags(flags0)
	ret1 := C.g_key_file_load_from_file(this1, file1, flags1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) RemoveComment(group_name0 string, key0 string) (bool, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_remove_comment(this1, group_name1, key1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) RemoveGroup(group_name0 string) (bool, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	ret1 := C.g_key_file_remove_group(this1, group_name1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) RemoveKey(group_name0 string, key0 string) (bool, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.g_key_file_remove_key(this1, group_name1, key1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) SaveToFile(filename0 string) (bool, error) {
	var this1 *C.GKeyFile
	var filename1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	filename1 = _GoStringToGString(filename0)
	defer C.free(unsafe.Pointer(filename1))
	ret1 := C.g_key_file_save_to_file(this1, filename1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) SetBoolean(group_name0 string, key0 string, value0 bool) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var value1 C.int
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = _GoBoolToCBool(value0)
	C.g_key_file_set_boolean(this1, group_name1, key1, value1)
}
func (this0 *KeyFile) SetBooleanList(group_name0 string, key0 string, list0 []bool) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var list1 *C.int
	var length1 C.uint64_t
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	list1 = (*C.int)(C.malloc(C.size_t(int(unsafe.Sizeof(*list1)) * len(list0))))
	defer C.free(unsafe.Pointer(list1))
	for i, e := range list0 {
		(*(*[999999]C.int)(unsafe.Pointer(list1)))[i] = _GoBoolToCBool(e)
	}
	length1 = C.uint64_t(len(list0))
	C.g_key_file_set_boolean_list(this1, group_name1, key1, list1, length1)
}
func (this0 *KeyFile) SetComment(group_name0 string, key0 string, comment0 string) (bool, error) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var comment1 *C.char
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	comment1 = _GoStringToGString(comment0)
	defer C.free(unsafe.Pointer(comment1))
	ret1 := C.g_key_file_set_comment(this1, group_name1, key1, comment1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *KeyFile) SetDouble(group_name0 string, key0 string, value0 float64) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var value1 C.double
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = C.double(value0)
	C.g_key_file_set_double(this1, group_name1, key1, value1)
}
func (this0 *KeyFile) SetDoubleList(group_name0 string, key0 string, list0 []float64) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var list1 *C.double
	var length1 C.uint64_t
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	list1 = (*C.double)(C.malloc(C.size_t(int(unsafe.Sizeof(*list1)) * len(list0))))
	defer C.free(unsafe.Pointer(list1))
	for i, e := range list0 {
		(*(*[999999]C.double)(unsafe.Pointer(list1)))[i] = C.double(e)
	}
	length1 = C.uint64_t(len(list0))
	C.g_key_file_set_double_list(this1, group_name1, key1, list1, length1)
}
func (this0 *KeyFile) SetInt64(group_name0 string, key0 string, value0 int64) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var value1 C.int64_t
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = C.int64_t(value0)
	C.g_key_file_set_int64(this1, group_name1, key1, value1)
}
func (this0 *KeyFile) SetInteger(group_name0 string, key0 string, value0 int32) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var value1 C.int32_t
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = C.int32_t(value0)
	C.g_key_file_set_integer(this1, group_name1, key1, value1)
}
func (this0 *KeyFile) SetIntegerList(group_name0 string, key0 string, list0 []int32) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var list1 *C.int32_t
	var length1 C.uint64_t
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	list1 = (*C.int32_t)(C.malloc(C.size_t(int(unsafe.Sizeof(*list1)) * len(list0))))
	defer C.free(unsafe.Pointer(list1))
	for i, e := range list0 {
		(*(*[999999]C.int32_t)(unsafe.Pointer(list1)))[i] = C.int32_t(e)
	}
	length1 = C.uint64_t(len(list0))
	C.g_key_file_set_integer_list(this1, group_name1, key1, list1, length1)
}
func (this0 *KeyFile) SetListSeparator(separator0 int8) {
	var this1 *C.GKeyFile
	var separator1 C.int8_t
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	separator1 = C.int8_t(separator0)
	C.g_key_file_set_list_separator(this1, separator1)
}
func (this0 *KeyFile) SetLocaleString(group_name0 string, key0 string, locale0 string, string0 string) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var locale1 *C.char
	var string1 *C.char
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	locale1 = _GoStringToGString(locale0)
	defer C.free(unsafe.Pointer(locale1))
	string1 = _GoStringToGString(string0)
	defer C.free(unsafe.Pointer(string1))
	C.g_key_file_set_locale_string(this1, group_name1, key1, locale1, string1)
}
func (this0 *KeyFile) SetLocaleStringList(group_name0 string, key0 string, locale0 string, list0 []string) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var locale1 *C.char
	var list1 **C.char
	var length1 C.uint64_t
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	locale1 = _GoStringToGString(locale0)
	defer C.free(unsafe.Pointer(locale1))
	list1 = (**C.char)(C.malloc(C.size_t(int(unsafe.Sizeof(*list1)) * (len(list0) + 1))))
	defer C.free(unsafe.Pointer(list1))
	for i, e := range list0 {
		(*(*[999999]*C.char)(unsafe.Pointer(list1)))[i] = _GoStringToGString(e)
		defer C.free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(list1)))[i]))
	}
	(*(*[999999]*C.char)(unsafe.Pointer(list1)))[len(list0)] = nil
	length1 = C.uint64_t(len(list0))
	C.g_key_file_set_locale_string_list(this1, group_name1, key1, locale1, list1, length1)
}
func (this0 *KeyFile) SetString(group_name0 string, key0 string, string0 string) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var string1 *C.char
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	string1 = _GoStringToGString(string0)
	defer C.free(unsafe.Pointer(string1))
	C.g_key_file_set_string(this1, group_name1, key1, string1)
}
func (this0 *KeyFile) SetStringList(group_name0 string, key0 string, list0 []string) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var list1 **C.char
	var length1 C.uint64_t
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	list1 = (**C.char)(C.malloc(C.size_t(int(unsafe.Sizeof(*list1)) * (len(list0) + 1))))
	defer C.free(unsafe.Pointer(list1))
	for i, e := range list0 {
		(*(*[999999]*C.char)(unsafe.Pointer(list1)))[i] = _GoStringToGString(e)
		defer C.free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(list1)))[i]))
	}
	(*(*[999999]*C.char)(unsafe.Pointer(list1)))[len(list0)] = nil
	length1 = C.uint64_t(len(list0))
	C.g_key_file_set_string_list(this1, group_name1, key1, list1, length1)
}
func (this0 *KeyFile) SetUint64(group_name0 string, key0 string, value0 uint64) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var value1 C.uint64_t
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = C.uint64_t(value0)
	C.g_key_file_set_uint64(this1, group_name1, key1, value1)
}
func (this0 *KeyFile) SetValue(group_name0 string, key0 string, value0 string) {
	var this1 *C.GKeyFile
	var group_name1 *C.char
	var key1 *C.char
	var value1 *C.char
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	group_name1 = _GoStringToGString(group_name0)
	defer C.free(unsafe.Pointer(group_name1))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	value1 = _GoStringToGString(value0)
	defer C.free(unsafe.Pointer(value1))
	C.g_key_file_set_value(this1, group_name1, key1, value1)
}
func (this0 *KeyFile) ToData() (uint64, string, error) {
	var this1 *C.GKeyFile
	var length1 C.uint64_t
	var err1 *C.GError
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	ret1 := C.g_key_file_to_data(this1, &length1, &err1)
	var length2 uint64
	var ret2 string
	var err2 error
	length2 = uint64(length1)
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return length2, ret2, err2
}
func KeyFileErrorQuark() uint32 {
	ret1 := C.g_key_file_error_quark()
	var ret2 uint32
	ret2 = uint32(ret1)
	return ret2
}
type KeyFileError C.uint32_t
const (
	KeyFileErrorUnknownEncoding KeyFileError = 0
	KeyFileErrorParse KeyFileError = 1
	KeyFileErrorNotFound KeyFileError = 2
	KeyFileErrorKeyNotFound KeyFileError = 3
	KeyFileErrorGroupNotFound KeyFileError = 4
	KeyFileErrorInvalidValue KeyFileError = 5
)
type KeyFileFlags C.uint32_t
const (
	KeyFileFlagsNone KeyFileFlags = 0
	KeyFileFlagsKeepComments KeyFileFlags = 1
	KeyFileFlagsKeepTranslations KeyFileFlags = 2
)
const LittleEndian = 1234
const Ln10 = 2.302585
const Ln2 = 0.693147
const Log2Base10 = 0.30103
const LogDomain = 0
const LogFatalMask = 0
const LogLevelUserShift = 8
// blacklisted: List (struct)
// blacklisted: LogFunc (callback)
type LogLevelFlags C.int32_t
const (
	LogLevelFlagsFlagRecursion LogLevelFlags = 1
	LogLevelFlagsFlagFatal LogLevelFlags = 2
	LogLevelFlagsLevelError LogLevelFlags = 4
	LogLevelFlagsLevelCritical LogLevelFlags = 8
	LogLevelFlagsLevelWarning LogLevelFlags = 16
	LogLevelFlagsLevelMessage LogLevelFlags = 32
	LogLevelFlagsLevelInfo LogLevelFlags = 64
	LogLevelFlagsLevelDebug LogLevelFlags = 128
	LogLevelFlagsLevelMask LogLevelFlags = -4
)
const MajorVersion = 2
const Maxint16 = 32767
const Maxint32 = 2147483647
const Maxint64 = 9223372036854775807
const Maxint8 = 127
const Maxuint16 = 0xffff
const Maxuint32 = 0xffffffff
const Maxuint64 = 0xffffffffffffffff
const Maxuint8 = 0xff
const MicroVersion = 0
const Minint16 = -32768
const Minint32 = -2147483648
const Minint64 = -9223372036854775808
const Minint8 = -128
const MinorVersion = 40
const ModuleSuffix = "so"
// blacklisted: MainContext (struct)
// blacklisted: MainLoop (struct)
// blacklisted: MappedFile (struct)
type MarkupCollectType C.uint32_t
const (
	MarkupCollectTypeInvalid MarkupCollectType = 0
	MarkupCollectTypeString MarkupCollectType = 1
	MarkupCollectTypeStrdup MarkupCollectType = 2
	MarkupCollectTypeBoolean MarkupCollectType = 3
	MarkupCollectTypeTristate MarkupCollectType = 4
	MarkupCollectTypeOptional MarkupCollectType = 65536
)
type MarkupError C.uint32_t
const (
	MarkupErrorBadUtf8 MarkupError = 0
	MarkupErrorEmpty MarkupError = 1
	MarkupErrorParse MarkupError = 2
	MarkupErrorUnknownElement MarkupError = 3
	MarkupErrorUnknownAttribute MarkupError = 4
	MarkupErrorInvalidContent MarkupError = 5
	MarkupErrorMissingAttribute MarkupError = 6
)
// blacklisted: MarkupParseContext (struct)
type MarkupParseFlags C.uint32_t
const (
	MarkupParseFlagsDoNotUseThisUnsupportedFlag MarkupParseFlags = 1
	MarkupParseFlagsTreatCdataAsText MarkupParseFlags = 2
	MarkupParseFlagsPrefixErrorPosition MarkupParseFlags = 4
	MarkupParseFlagsIgnoreQualified MarkupParseFlags = 8
)
// blacklisted: MarkupParser (struct)
// blacklisted: MatchInfo (struct)
// blacklisted: MemVTable (struct)
type Mutex struct {
	_data [8]byte
}
func (this0 *Mutex) Clear() {
	var this1 *C.GMutex
	C.g_mutex_clear(this1)
}
func (this0 *Mutex) Init() {
	var this1 *C.GMutex
	C.g_mutex_init(this1)
}
func (this0 *Mutex) Lock() {
	var this1 *C.GMutex
	C.g_mutex_lock(this1)
}
func (this0 *Mutex) Trylock() bool {
	var this1 *C.GMutex
	ret1 := C.g_mutex_trylock(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Mutex) Unlock() {
	var this1 *C.GMutex
	C.g_mutex_unlock(this1)
}
// blacklisted: Node (struct)
// blacklisted: NodeForeachFunc (callback)
// blacklisted: NodeTraverseFunc (callback)
type NormalizeMode C.uint32_t
const (
	NormalizeModeDefault NormalizeMode = 0
	NormalizeModeNfd NormalizeMode = 0
	NormalizeModeDefaultCompose NormalizeMode = 1
	NormalizeModeNfc NormalizeMode = 1
	NormalizeModeAll NormalizeMode = 2
	NormalizeModeNfkd NormalizeMode = 2
	NormalizeModeAllCompose NormalizeMode = 3
	NormalizeModeNfkc NormalizeMode = 3
)
const OptionRemaining = ""
// blacklisted: Once (struct)
type OnceStatus C.uint32_t
const (
	OnceStatusNotcalled OnceStatus = 0
	OnceStatusProgress OnceStatus = 1
	OnceStatusReady OnceStatus = 2
)
type OptionArg C.uint32_t
const (
	OptionArgNone OptionArg = 0
	OptionArgString OptionArg = 1
	OptionArgInt OptionArg = 2
	OptionArgCallback OptionArg = 3
	OptionArgFilename OptionArg = 4
	OptionArgStringArray OptionArg = 5
	OptionArgFilenameArray OptionArg = 6
	OptionArgDouble OptionArg = 7
	OptionArgInt64 OptionArg = 8
)
// blacklisted: OptionArgFunc (callback)
// blacklisted: OptionContext (struct)
// blacklisted: OptionEntry (struct)
type OptionError C.uint32_t
const (
	OptionErrorUnknownOption OptionError = 0
	OptionErrorBadValue OptionError = 1
	OptionErrorFailed OptionError = 2
)
// blacklisted: OptionErrorFunc (callback)
type OptionFlags C.uint32_t
const (
	OptionFlagsHidden OptionFlags = 1
	OptionFlagsInMain OptionFlags = 2
	OptionFlagsReverse OptionFlags = 4
	OptionFlagsNoArg OptionFlags = 8
	OptionFlagsFilename OptionFlags = 16
	OptionFlagsOptionalArg OptionFlags = 32
	OptionFlagsNoalias OptionFlags = 64
)
// blacklisted: OptionGroup (struct)
// blacklisted: OptionParseFunc (callback)
const PdpEndian = 3412
const Pi = 3.141593
const Pi2 = 1.570796
const Pi4 = 0.785398
const PollfdFormat = "%#I64x"
const PriorityDefault = 0
const PriorityDefaultIdle = 200
const PriorityHigh = -100
const PriorityHighIdle = 100
const PriorityLow = 300
// blacklisted: PatternSpec (struct)
type PollFD struct {
	Fd int32
	Events uint16
	Revents uint16
}
// blacklisted: PollFunc (callback)
// blacklisted: PrintFunc (callback)
// blacklisted: Private (struct)
// blacklisted: PtrArray (struct)
// blacklisted: Queue (struct)
// blacklisted: RWLock (struct)
// blacklisted: Rand (struct)
// blacklisted: RecMutex (struct)
// blacklisted: Regex (struct)
type RegexCompileFlags C.uint32_t
const (
	RegexCompileFlagsCaseless RegexCompileFlags = 1
	RegexCompileFlagsMultiline RegexCompileFlags = 2
	RegexCompileFlagsDotall RegexCompileFlags = 4
	RegexCompileFlagsExtended RegexCompileFlags = 8
	RegexCompileFlagsAnchored RegexCompileFlags = 16
	RegexCompileFlagsDollarEndonly RegexCompileFlags = 32
	RegexCompileFlagsUngreedy RegexCompileFlags = 512
	RegexCompileFlagsRaw RegexCompileFlags = 2048
	RegexCompileFlagsNoAutoCapture RegexCompileFlags = 4096
	RegexCompileFlagsOptimize RegexCompileFlags = 8192
	RegexCompileFlagsFirstline RegexCompileFlags = 262144
	RegexCompileFlagsDupnames RegexCompileFlags = 524288
	RegexCompileFlagsNewlineCr RegexCompileFlags = 1048576
	RegexCompileFlagsNewlineLf RegexCompileFlags = 2097152
	RegexCompileFlagsNewlineCrlf RegexCompileFlags = 3145728
	RegexCompileFlagsNewlineAnycrlf RegexCompileFlags = 5242880
	RegexCompileFlagsBsrAnycrlf RegexCompileFlags = 8388608
	RegexCompileFlagsJavascriptCompat RegexCompileFlags = 33554432
)
type RegexError C.uint32_t
const (
	RegexErrorCompile RegexError = 0
	RegexErrorOptimize RegexError = 1
	RegexErrorReplace RegexError = 2
	RegexErrorMatch RegexError = 3
	RegexErrorInternal RegexError = 4
	RegexErrorStrayBackslash RegexError = 101
	RegexErrorMissingControlChar RegexError = 102
	RegexErrorUnrecognizedEscape RegexError = 103
	RegexErrorQuantifiersOutOfOrder RegexError = 104
	RegexErrorQuantifierTooBig RegexError = 105
	RegexErrorUnterminatedCharacterClass RegexError = 106
	RegexErrorInvalidEscapeInCharacterClass RegexError = 107
	RegexErrorRangeOutOfOrder RegexError = 108
	RegexErrorNothingToRepeat RegexError = 109
	RegexErrorUnrecognizedCharacter RegexError = 112
	RegexErrorPosixNamedClassOutsideClass RegexError = 113
	RegexErrorUnmatchedParenthesis RegexError = 114
	RegexErrorInexistentSubpatternReference RegexError = 115
	RegexErrorUnterminatedComment RegexError = 118
	RegexErrorExpressionTooLarge RegexError = 120
	RegexErrorMemoryError RegexError = 121
	RegexErrorVariableLengthLookbehind RegexError = 125
	RegexErrorMalformedCondition RegexError = 126
	RegexErrorTooManyConditionalBranches RegexError = 127
	RegexErrorAssertionExpected RegexError = 128
	RegexErrorUnknownPosixClassName RegexError = 130
	RegexErrorPosixCollatingElementsNotSupported RegexError = 131
	RegexErrorHexCodeTooLarge RegexError = 134
	RegexErrorInvalidCondition RegexError = 135
	RegexErrorSingleByteMatchInLookbehind RegexError = 136
	RegexErrorInfiniteLoop RegexError = 140
	RegexErrorMissingSubpatternNameTerminator RegexError = 142
	RegexErrorDuplicateSubpatternName RegexError = 143
	RegexErrorMalformedProperty RegexError = 146
	RegexErrorUnknownProperty RegexError = 147
	RegexErrorSubpatternNameTooLong RegexError = 148
	RegexErrorTooManySubpatterns RegexError = 149
	RegexErrorInvalidOctalValue RegexError = 151
	RegexErrorTooManyBranchesInDefine RegexError = 154
	RegexErrorDefineRepetion RegexError = 155
	RegexErrorInconsistentNewlineOptions RegexError = 156
	RegexErrorMissingBackReference RegexError = 157
	RegexErrorInvalidRelativeReference RegexError = 158
	RegexErrorBacktrackingControlVerbArgumentForbidden RegexError = 159
	RegexErrorUnknownBacktrackingControlVerb RegexError = 160
	RegexErrorNumberTooBig RegexError = 161
	RegexErrorMissingSubpatternName RegexError = 162
	RegexErrorMissingDigit RegexError = 163
	RegexErrorInvalidDataCharacter RegexError = 164
	RegexErrorExtraSubpatternName RegexError = 165
	RegexErrorBacktrackingControlVerbArgumentRequired RegexError = 166
	RegexErrorInvalidControlChar RegexError = 168
	RegexErrorMissingName RegexError = 169
	RegexErrorNotSupportedInClass RegexError = 171
	RegexErrorTooManyForwardReferences RegexError = 172
	RegexErrorNameTooLong RegexError = 175
	RegexErrorCharacterValueTooLarge RegexError = 176
)
// blacklisted: RegexEvalCallback (callback)
type RegexMatchFlags C.uint32_t
const (
	RegexMatchFlagsAnchored RegexMatchFlags = 16
	RegexMatchFlagsNotbol RegexMatchFlags = 128
	RegexMatchFlagsNoteol RegexMatchFlags = 256
	RegexMatchFlagsNotempty RegexMatchFlags = 1024
	RegexMatchFlagsPartial RegexMatchFlags = 32768
	RegexMatchFlagsNewlineCr RegexMatchFlags = 1048576
	RegexMatchFlagsNewlineLf RegexMatchFlags = 2097152
	RegexMatchFlagsNewlineCrlf RegexMatchFlags = 3145728
	RegexMatchFlagsNewlineAny RegexMatchFlags = 4194304
	RegexMatchFlagsNewlineAnycrlf RegexMatchFlags = 5242880
	RegexMatchFlagsBsrAnycrlf RegexMatchFlags = 8388608
	RegexMatchFlagsBsrAny RegexMatchFlags = 16777216
	RegexMatchFlagsPartialSoft RegexMatchFlags = 32768
	RegexMatchFlagsPartialHard RegexMatchFlags = 134217728
	RegexMatchFlagsNotemptyAtstart RegexMatchFlags = 268435456
)
const SearchpathSeparator = 59
const SearchpathSeparatorS = ";"
const SizeofLong = 8
const SizeofSizeT = 8
const SizeofSsizeT = 8
const SizeofVoidP = 8
// blacklisted: SList (struct)
const SourceContinue = true
const SourceRemove = false
const Sqrt2 = 1.414214
const StrDelimiters = "_-|> <."
const SysdefAfInet = 2
const SysdefAfInet6 = 10
const SysdefAfUnix = 1
const SysdefMsgDontroute = 4
const SysdefMsgOob = 1
const SysdefMsgPeek = 2
// blacklisted: Scanner (struct)
// blacklisted: ScannerConfig (struct)
// blacklisted: ScannerMsgFunc (callback)
type SeekType C.uint32_t
const (
	SeekTypeCur SeekType = 0
	SeekTypeSet SeekType = 1
	SeekTypeEnd SeekType = 2
)
// blacklisted: Sequence (struct)
// blacklisted: SequenceIter (struct)
// blacklisted: SequenceIterCompareFunc (callback)
type ShellError C.uint32_t
const (
	ShellErrorBadQuoting ShellError = 0
	ShellErrorEmptyString ShellError = 1
	ShellErrorFailed ShellError = 2
)
type SliceConfig C.uint32_t
const (
	SliceConfigAlwaysMalloc SliceConfig = 1
	SliceConfigBypassMagazines SliceConfig = 2
	SliceConfigWorkingSetMsecs SliceConfig = 3
	SliceConfigColorIncrement SliceConfig = 4
	SliceConfigChunkSizes SliceConfig = 5
	SliceConfigContentionCounter SliceConfig = 6
)
// blacklisted: Source (struct)
// blacklisted: SourceCallbackFuncs (struct)
// blacklisted: SourceDummyMarshal (callback)
// blacklisted: SourceFunc (callback)
// blacklisted: SourceFuncs (struct)
// blacklisted: SourcePrivate (struct)
// blacklisted: SpawnChildSetupFunc (callback)
type SpawnError C.uint32_t
const (
	SpawnErrorFork SpawnError = 0
	SpawnErrorRead SpawnError = 1
	SpawnErrorChdir SpawnError = 2
	SpawnErrorAcces SpawnError = 3
	SpawnErrorPerm SpawnError = 4
	SpawnErrorTooBig SpawnError = 5
	SpawnError2big SpawnError = 5
	SpawnErrorNoexec SpawnError = 6
	SpawnErrorNametoolong SpawnError = 7
	SpawnErrorNoent SpawnError = 8
	SpawnErrorNomem SpawnError = 9
	SpawnErrorNotdir SpawnError = 10
	SpawnErrorLoop SpawnError = 11
	SpawnErrorTxtbusy SpawnError = 12
	SpawnErrorIo SpawnError = 13
	SpawnErrorNfile SpawnError = 14
	SpawnErrorMfile SpawnError = 15
	SpawnErrorInval SpawnError = 16
	SpawnErrorIsdir SpawnError = 17
	SpawnErrorLibbad SpawnError = 18
	SpawnErrorFailed SpawnError = 19
)
type SpawnFlags C.uint32_t
const (
	SpawnFlagsDefault SpawnFlags = 0
	SpawnFlagsLeaveDescriptorsOpen SpawnFlags = 1
	SpawnFlagsDoNotReapChild SpawnFlags = 2
	SpawnFlagsSearchPath SpawnFlags = 4
	SpawnFlagsStdoutToDevNull SpawnFlags = 8
	SpawnFlagsStderrToDevNull SpawnFlags = 16
	SpawnFlagsChildInheritsStdin SpawnFlags = 32
	SpawnFlagsFileAndArgvZero SpawnFlags = 64
	SpawnFlagsSearchPathFromEnvp SpawnFlags = 128
	SpawnFlagsCloexecPipes SpawnFlags = 256
)
// blacklisted: StatBuf (struct)
// blacklisted: String (struct)
// blacklisted: StringChunk (struct)
const TimeSpanDay = 86400000000
const TimeSpanHour = 3600000000
const TimeSpanMillisecond = 1000
const TimeSpanMinute = 60000000
const TimeSpanSecond = 1000000
// blacklisted: TestCase (struct)
// blacklisted: TestConfig (struct)
// blacklisted: TestDataFunc (callback)
type TestFileType C.uint32_t
const (
	TestFileTypeDist TestFileType = 0
	TestFileTypeBuilt TestFileType = 1
)
// blacklisted: TestFixtureFunc (callback)
// blacklisted: TestFunc (callback)
// blacklisted: TestLogBuffer (struct)
// blacklisted: TestLogFatalFunc (callback)
// blacklisted: TestLogMsg (struct)
type TestLogType C.uint32_t
const (
	TestLogTypeNone TestLogType = 0
	TestLogTypeError TestLogType = 1
	TestLogTypeStartBinary TestLogType = 2
	TestLogTypeListCase TestLogType = 3
	TestLogTypeSkipCase TestLogType = 4
	TestLogTypeStartCase TestLogType = 5
	TestLogTypeStopCase TestLogType = 6
	TestLogTypeMinResult TestLogType = 7
	TestLogTypeMaxResult TestLogType = 8
	TestLogTypeMessage TestLogType = 9
	TestLogTypeStartSuite TestLogType = 10
	TestLogTypeStopSuite TestLogType = 11
)
type TestSubprocessFlags C.uint32_t
const (
	TestSubprocessFlagsStdin TestSubprocessFlags = 1
	TestSubprocessFlagsStdout TestSubprocessFlags = 2
	TestSubprocessFlagsStderr TestSubprocessFlags = 4
)
// blacklisted: TestSuite (struct)
type TestTrapFlags C.uint32_t
const (
	TestTrapFlagsSilenceStdout TestTrapFlags = 128
	TestTrapFlagsSilenceStderr TestTrapFlags = 256
	TestTrapFlagsInheritStdin TestTrapFlags = 512
)
// blacklisted: Thread (struct)
type ThreadError C.uint32_t
const (
	ThreadErrorThreadErrorAgain ThreadError = 0
)
// blacklisted: ThreadPool (struct)
type TimeType C.uint32_t
const (
	TimeTypeStandard TimeType = 0
	TimeTypeDaylight TimeType = 1
	TimeTypeUniversal TimeType = 2
)
type TimeVal struct {
	TvSec int64
	TvUsec int64
}
func (this0 *TimeVal) Add(microseconds0 int64) {
	var this1 *C.GTimeVal
	var microseconds1 C.int64_t
	this1 = (*C.GTimeVal)(unsafe.Pointer(this0))
	microseconds1 = C.int64_t(microseconds0)
	C.g_time_val_add(this1, microseconds1)
}
func (this0 *TimeVal) ToIso8601() string {
	var this1 *C.GTimeVal
	this1 = (*C.GTimeVal)(unsafe.Pointer(this0))
	ret1 := C.g_time_val_to_iso8601(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func TimeValFromIso8601(iso_date0 string) (TimeVal, bool) {
	var iso_date1 *C.char
	var time_1 C.GTimeVal
	iso_date1 = _GoStringToGString(iso_date0)
	defer C.free(unsafe.Pointer(iso_date1))
	ret1 := C.g_time_val_from_iso8601(iso_date1, &time_1)
	var time_2 TimeVal
	var ret2 bool
	time_2 = *(*TimeVal)(unsafe.Pointer(&time_1))
	ret2 = ret1 != 0
	return time_2, ret2
}
type TimeZone struct {}
func NewTimeZone(identifier0 string) *TimeZone {
	var identifier1 *C.char
	identifier1 = _GoStringToGString(identifier0)
	defer C.free(unsafe.Pointer(identifier1))
	ret1 := C.g_time_zone_new(identifier1)
	var ret2 *TimeZone
	ret2 = (*TimeZone)(unsafe.Pointer(ret1))
	return ret2
}
func NewTimeZoneLocal() *TimeZone {
	ret1 := C.g_time_zone_new_local()
	var ret2 *TimeZone
	ret2 = (*TimeZone)(unsafe.Pointer(ret1))
	return ret2
}
func NewTimeZoneUtc() *TimeZone {
	ret1 := C.g_time_zone_new_utc()
	var ret2 *TimeZone
	ret2 = (*TimeZone)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *TimeZone) AdjustTime(type0 TimeType, time_0 *int64) int32 {
	var this1 *C.GTimeZone
	var type1 C.GTimeType
	var time_1 *C.int64_t
	this1 = (*C.GTimeZone)(unsafe.Pointer(this0))
	type1 = C.GTimeType(type0)
	time_1 = (*C.int64_t)(unsafe.Pointer(time_0))
	ret1 := C.g_time_zone_adjust_time(this1, type1, time_1)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func (this0 *TimeZone) FindInterval(type0 TimeType, time_0 int64) int32 {
	var this1 *C.GTimeZone
	var type1 C.GTimeType
	var time_1 C.int64_t
	this1 = (*C.GTimeZone)(unsafe.Pointer(this0))
	type1 = C.GTimeType(type0)
	time_1 = C.int64_t(time_0)
	ret1 := C.g_time_zone_find_interval(this1, type1, time_1)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func (this0 *TimeZone) GetAbbreviation(interval0 int32) string {
	var this1 *C.GTimeZone
	var interval1 C.int32_t
	this1 = (*C.GTimeZone)(unsafe.Pointer(this0))
	interval1 = C.int32_t(interval0)
	ret1 := C.g_time_zone_get_abbreviation(this1, interval1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *TimeZone) GetOffset(interval0 int32) int32 {
	var this1 *C.GTimeZone
	var interval1 C.int32_t
	this1 = (*C.GTimeZone)(unsafe.Pointer(this0))
	interval1 = C.int32_t(interval0)
	ret1 := C.g_time_zone_get_offset(this1, interval1)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func (this0 *TimeZone) IsDst(interval0 int32) bool {
	var this1 *C.GTimeZone
	var interval1 C.int32_t
	this1 = (*C.GTimeZone)(unsafe.Pointer(this0))
	interval1 = C.int32_t(interval0)
	ret1 := C.g_time_zone_is_dst(this1, interval1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
// blacklisted: Timer (struct)
type TokenType C.uint32_t
const (
	TokenTypeEof TokenType = 0
	TokenTypeLeftParen TokenType = 40
	TokenTypeRightParen TokenType = 41
	TokenTypeLeftCurly TokenType = 123
	TokenTypeRightCurly TokenType = 125
	TokenTypeLeftBrace TokenType = 91
	TokenTypeRightBrace TokenType = 93
	TokenTypeEqualSign TokenType = 61
	TokenTypeComma TokenType = 44
	TokenTypeNone TokenType = 256
	TokenTypeError TokenType = 257
	TokenTypeChar TokenType = 258
	TokenTypeBinary TokenType = 259
	TokenTypeOctal TokenType = 260
	TokenTypeInt TokenType = 261
	TokenTypeHex TokenType = 262
	TokenTypeFloat TokenType = 263
	TokenTypeString TokenType = 264
	TokenTypeSymbol TokenType = 265
	TokenTypeIdentifier TokenType = 266
	TokenTypeIdentifierNull TokenType = 267
	TokenTypeCommentSingle TokenType = 268
	TokenTypeCommentMulti TokenType = 269
)
type TokenValue struct {
	_data [8]byte
}
// blacklisted: TranslateFunc (callback)
// blacklisted: TrashStack (struct)
type TraverseFlags C.uint32_t
const (
	TraverseFlagsLeaves TraverseFlags = 1
	TraverseFlagsNonLeaves TraverseFlags = 2
	TraverseFlagsAll TraverseFlags = 3
	TraverseFlagsMask TraverseFlags = 3
	TraverseFlagsLeafs TraverseFlags = 1
	TraverseFlagsNonLeafs TraverseFlags = 2
)
// blacklisted: TraverseFunc (callback)
type TraverseType C.uint32_t
const (
	TraverseTypeInOrder TraverseType = 0
	TraverseTypePreOrder TraverseType = 1
	TraverseTypePostOrder TraverseType = 2
	TraverseTypeLevelOrder TraverseType = 3
)
// blacklisted: Tree (struct)
const UnicharMaxDecompositionLength = 18
const UriReservedCharsGenericDelimiters = ":/?#[]@"
const UriReservedCharsSubcomponentDelimiters = "!$&'()*+,;="
const UsecPerSec = 1000000
type UnicodeBreakType C.uint32_t
const (
	UnicodeBreakTypeMandatory UnicodeBreakType = 0
	UnicodeBreakTypeCarriageReturn UnicodeBreakType = 1
	UnicodeBreakTypeLineFeed UnicodeBreakType = 2
	UnicodeBreakTypeCombiningMark UnicodeBreakType = 3
	UnicodeBreakTypeSurrogate UnicodeBreakType = 4
	UnicodeBreakTypeZeroWidthSpace UnicodeBreakType = 5
	UnicodeBreakTypeInseparable UnicodeBreakType = 6
	UnicodeBreakTypeNonBreakingGlue UnicodeBreakType = 7
	UnicodeBreakTypeContingent UnicodeBreakType = 8
	UnicodeBreakTypeSpace UnicodeBreakType = 9
	UnicodeBreakTypeAfter UnicodeBreakType = 10
	UnicodeBreakTypeBefore UnicodeBreakType = 11
	UnicodeBreakTypeBeforeAndAfter UnicodeBreakType = 12
	UnicodeBreakTypeHyphen UnicodeBreakType = 13
	UnicodeBreakTypeNonStarter UnicodeBreakType = 14
	UnicodeBreakTypeOpenPunctuation UnicodeBreakType = 15
	UnicodeBreakTypeClosePunctuation UnicodeBreakType = 16
	UnicodeBreakTypeQuotation UnicodeBreakType = 17
	UnicodeBreakTypeExclamation UnicodeBreakType = 18
	UnicodeBreakTypeIdeographic UnicodeBreakType = 19
	UnicodeBreakTypeNumeric UnicodeBreakType = 20
	UnicodeBreakTypeInfixSeparator UnicodeBreakType = 21
	UnicodeBreakTypeSymbol UnicodeBreakType = 22
	UnicodeBreakTypeAlphabetic UnicodeBreakType = 23
	UnicodeBreakTypePrefix UnicodeBreakType = 24
	UnicodeBreakTypePostfix UnicodeBreakType = 25
	UnicodeBreakTypeComplexContext UnicodeBreakType = 26
	UnicodeBreakTypeAmbiguous UnicodeBreakType = 27
	UnicodeBreakTypeUnknown UnicodeBreakType = 28
	UnicodeBreakTypeNextLine UnicodeBreakType = 29
	UnicodeBreakTypeWordJoiner UnicodeBreakType = 30
	UnicodeBreakTypeHangulLJamo UnicodeBreakType = 31
	UnicodeBreakTypeHangulVJamo UnicodeBreakType = 32
	UnicodeBreakTypeHangulTJamo UnicodeBreakType = 33
	UnicodeBreakTypeHangulLvSyllable UnicodeBreakType = 34
	UnicodeBreakTypeHangulLvtSyllable UnicodeBreakType = 35
	UnicodeBreakTypeCloseParanthesis UnicodeBreakType = 36
	UnicodeBreakTypeConditionalJapaneseStarter UnicodeBreakType = 37
	UnicodeBreakTypeHebrewLetter UnicodeBreakType = 38
	UnicodeBreakTypeRegionalIndicator UnicodeBreakType = 39
)
type UnicodeScript C.int32_t
const (
	UnicodeScriptInvalidCode UnicodeScript = -1
	UnicodeScriptCommon UnicodeScript = 0
	UnicodeScriptInherited UnicodeScript = 1
	UnicodeScriptArabic UnicodeScript = 2
	UnicodeScriptArmenian UnicodeScript = 3
	UnicodeScriptBengali UnicodeScript = 4
	UnicodeScriptBopomofo UnicodeScript = 5
	UnicodeScriptCherokee UnicodeScript = 6
	UnicodeScriptCoptic UnicodeScript = 7
	UnicodeScriptCyrillic UnicodeScript = 8
	UnicodeScriptDeseret UnicodeScript = 9
	UnicodeScriptDevanagari UnicodeScript = 10
	UnicodeScriptEthiopic UnicodeScript = 11
	UnicodeScriptGeorgian UnicodeScript = 12
	UnicodeScriptGothic UnicodeScript = 13
	UnicodeScriptGreek UnicodeScript = 14
	UnicodeScriptGujarati UnicodeScript = 15
	UnicodeScriptGurmukhi UnicodeScript = 16
	UnicodeScriptHan UnicodeScript = 17
	UnicodeScriptHangul UnicodeScript = 18
	UnicodeScriptHebrew UnicodeScript = 19
	UnicodeScriptHiragana UnicodeScript = 20
	UnicodeScriptKannada UnicodeScript = 21
	UnicodeScriptKatakana UnicodeScript = 22
	UnicodeScriptKhmer UnicodeScript = 23
	UnicodeScriptLao UnicodeScript = 24
	UnicodeScriptLatin UnicodeScript = 25
	UnicodeScriptMalayalam UnicodeScript = 26
	UnicodeScriptMongolian UnicodeScript = 27
	UnicodeScriptMyanmar UnicodeScript = 28
	UnicodeScriptOgham UnicodeScript = 29
	UnicodeScriptOldItalic UnicodeScript = 30
	UnicodeScriptOriya UnicodeScript = 31
	UnicodeScriptRunic UnicodeScript = 32
	UnicodeScriptSinhala UnicodeScript = 33
	UnicodeScriptSyriac UnicodeScript = 34
	UnicodeScriptTamil UnicodeScript = 35
	UnicodeScriptTelugu UnicodeScript = 36
	UnicodeScriptThaana UnicodeScript = 37
	UnicodeScriptThai UnicodeScript = 38
	UnicodeScriptTibetan UnicodeScript = 39
	UnicodeScriptCanadianAboriginal UnicodeScript = 40
	UnicodeScriptYi UnicodeScript = 41
	UnicodeScriptTagalog UnicodeScript = 42
	UnicodeScriptHanunoo UnicodeScript = 43
	UnicodeScriptBuhid UnicodeScript = 44
	UnicodeScriptTagbanwa UnicodeScript = 45
	UnicodeScriptBraille UnicodeScript = 46
	UnicodeScriptCypriot UnicodeScript = 47
	UnicodeScriptLimbu UnicodeScript = 48
	UnicodeScriptOsmanya UnicodeScript = 49
	UnicodeScriptShavian UnicodeScript = 50
	UnicodeScriptLinearB UnicodeScript = 51
	UnicodeScriptTaiLe UnicodeScript = 52
	UnicodeScriptUgaritic UnicodeScript = 53
	UnicodeScriptNewTaiLue UnicodeScript = 54
	UnicodeScriptBuginese UnicodeScript = 55
	UnicodeScriptGlagolitic UnicodeScript = 56
	UnicodeScriptTifinagh UnicodeScript = 57
	UnicodeScriptSylotiNagri UnicodeScript = 58
	UnicodeScriptOldPersian UnicodeScript = 59
	UnicodeScriptKharoshthi UnicodeScript = 60
	UnicodeScriptUnknown UnicodeScript = 61
	UnicodeScriptBalinese UnicodeScript = 62
	UnicodeScriptCuneiform UnicodeScript = 63
	UnicodeScriptPhoenician UnicodeScript = 64
	UnicodeScriptPhagsPa UnicodeScript = 65
	UnicodeScriptNko UnicodeScript = 66
	UnicodeScriptKayahLi UnicodeScript = 67
	UnicodeScriptLepcha UnicodeScript = 68
	UnicodeScriptRejang UnicodeScript = 69
	UnicodeScriptSundanese UnicodeScript = 70
	UnicodeScriptSaurashtra UnicodeScript = 71
	UnicodeScriptCham UnicodeScript = 72
	UnicodeScriptOlChiki UnicodeScript = 73
	UnicodeScriptVai UnicodeScript = 74
	UnicodeScriptCarian UnicodeScript = 75
	UnicodeScriptLycian UnicodeScript = 76
	UnicodeScriptLydian UnicodeScript = 77
	UnicodeScriptAvestan UnicodeScript = 78
	UnicodeScriptBamum UnicodeScript = 79
	UnicodeScriptEgyptianHieroglyphs UnicodeScript = 80
	UnicodeScriptImperialAramaic UnicodeScript = 81
	UnicodeScriptInscriptionalPahlavi UnicodeScript = 82
	UnicodeScriptInscriptionalParthian UnicodeScript = 83
	UnicodeScriptJavanese UnicodeScript = 84
	UnicodeScriptKaithi UnicodeScript = 85
	UnicodeScriptLisu UnicodeScript = 86
	UnicodeScriptMeeteiMayek UnicodeScript = 87
	UnicodeScriptOldSouthArabian UnicodeScript = 88
	UnicodeScriptOldTurkic UnicodeScript = 89
	UnicodeScriptSamaritan UnicodeScript = 90
	UnicodeScriptTaiTham UnicodeScript = 91
	UnicodeScriptTaiViet UnicodeScript = 92
	UnicodeScriptBatak UnicodeScript = 93
	UnicodeScriptBrahmi UnicodeScript = 94
	UnicodeScriptMandaic UnicodeScript = 95
	UnicodeScriptChakma UnicodeScript = 96
	UnicodeScriptMeroiticCursive UnicodeScript = 97
	UnicodeScriptMeroiticHieroglyphs UnicodeScript = 98
	UnicodeScriptMiao UnicodeScript = 99
	UnicodeScriptSharada UnicodeScript = 100
	UnicodeScriptSoraSompeng UnicodeScript = 101
	UnicodeScriptTakri UnicodeScript = 102
)
type UnicodeType C.uint32_t
const (
	UnicodeTypeControl UnicodeType = 0
	UnicodeTypeFormat UnicodeType = 1
	UnicodeTypeUnassigned UnicodeType = 2
	UnicodeTypePrivateUse UnicodeType = 3
	UnicodeTypeSurrogate UnicodeType = 4
	UnicodeTypeLowercaseLetter UnicodeType = 5
	UnicodeTypeModifierLetter UnicodeType = 6
	UnicodeTypeOtherLetter UnicodeType = 7
	UnicodeTypeTitlecaseLetter UnicodeType = 8
	UnicodeTypeUppercaseLetter UnicodeType = 9
	UnicodeTypeSpacingMark UnicodeType = 10
	UnicodeTypeEnclosingMark UnicodeType = 11
	UnicodeTypeNonSpacingMark UnicodeType = 12
	UnicodeTypeDecimalNumber UnicodeType = 13
	UnicodeTypeLetterNumber UnicodeType = 14
	UnicodeTypeOtherNumber UnicodeType = 15
	UnicodeTypeConnectPunctuation UnicodeType = 16
	UnicodeTypeDashPunctuation UnicodeType = 17
	UnicodeTypeClosePunctuation UnicodeType = 18
	UnicodeTypeFinalPunctuation UnicodeType = 19
	UnicodeTypeInitialPunctuation UnicodeType = 20
	UnicodeTypeOtherPunctuation UnicodeType = 21
	UnicodeTypeOpenPunctuation UnicodeType = 22
	UnicodeTypeCurrencySymbol UnicodeType = 23
	UnicodeTypeModifierSymbol UnicodeType = 24
	UnicodeTypeMathSymbol UnicodeType = 25
	UnicodeTypeOtherSymbol UnicodeType = 26
	UnicodeTypeLineSeparator UnicodeType = 27
	UnicodeTypeParagraphSeparator UnicodeType = 28
	UnicodeTypeSpaceSeparator UnicodeType = 29
)
// blacklisted: UnixFDSourceFunc (callback)
type UserDirectory C.uint32_t
const (
	UserDirectoryDirectoryDesktop UserDirectory = 0
	UserDirectoryDirectoryDocuments UserDirectory = 1
	UserDirectoryDirectoryDownload UserDirectory = 2
	UserDirectoryDirectoryMusic UserDirectory = 3
	UserDirectoryDirectoryPictures UserDirectory = 4
	UserDirectoryDirectoryPublicShare UserDirectory = 5
	UserDirectoryDirectoryTemplates UserDirectory = 6
	UserDirectoryDirectoryVideos UserDirectory = 7
	UserDirectoryNDirectories UserDirectory = 8
)
const VaCopyAsArray = 1
const VersionMinRequired = 2
type Variant struct {}
func NewVariantArray(child_type0 *VariantType, children0 []*Variant) *Variant {
	var child_type1 *C.GVariantType
	var children1 **C.GVariant
	var n_children1 C.uint64_t
	child_type1 = (*C.GVariantType)(unsafe.Pointer(child_type0))
	children1 = (**C.GVariant)(C.malloc(C.size_t(int(unsafe.Sizeof(*children1)) * len(children0))))
	defer C.free(unsafe.Pointer(children1))
	for i, e := range children0 {
		(*(*[999999]*C.GVariant)(unsafe.Pointer(children1)))[i] = (*C.GVariant)(unsafe.Pointer(e))
	}
	n_children1 = C.uint64_t(len(children0))
	ret1 := C.g_variant_new_array(child_type1, children1, n_children1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantBoolean(value0 bool) *Variant {
	var value1 C.int
	value1 = _GoBoolToCBool(value0)
	ret1 := C.g_variant_new_boolean(value1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantByte(value0 uint8) *Variant {
	var value1 C.uint8_t
	value1 = C.uint8_t(value0)
	ret1 := C.g_variant_new_byte(value1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantBytestringArray(strv0 []string) *Variant {
	var strv1 **C.char
	var length1 C.int64_t
	strv1 = (**C.char)(C.malloc(C.size_t(int(unsafe.Sizeof(*strv1)) * len(strv0))))
	defer C.free(unsafe.Pointer(strv1))
	for i, e := range strv0 {
		(*(*[999999]*C.char)(unsafe.Pointer(strv1)))[i] = _GoStringToGString(e)
		defer C.free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(strv1)))[i]))
	}
	length1 = C.int64_t(len(strv0))
	ret1 := C.g_variant_new_bytestring_array(strv1, length1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantDictEntry(key0 *Variant, value0 *Variant) *Variant {
	var key1 *C.GVariant
	var value1 *C.GVariant
	key1 = (*C.GVariant)(unsafe.Pointer(key0))
	value1 = (*C.GVariant)(unsafe.Pointer(value0))
	ret1 := C.g_variant_new_dict_entry(key1, value1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantDouble(value0 float64) *Variant {
	var value1 C.double
	value1 = C.double(value0)
	ret1 := C.g_variant_new_double(value1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantFixedArray(element_type0 *VariantType, elements0 unsafe.Pointer, n_elements0 uint64, element_size0 uint64) *Variant {
	var element_type1 *C.GVariantType
	var elements1 unsafe.Pointer
	var n_elements1 C.uint64_t
	var element_size1 C.uint64_t
	element_type1 = (*C.GVariantType)(unsafe.Pointer(element_type0))
	elements1 = unsafe.Pointer(elements0)
	n_elements1 = C.uint64_t(n_elements0)
	element_size1 = C.uint64_t(element_size0)
	ret1 := C.g_variant_new_fixed_array(element_type1, elements1, n_elements1, element_size1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantHandle(value0 int32) *Variant {
	var value1 C.int32_t
	value1 = C.int32_t(value0)
	ret1 := C.g_variant_new_handle(value1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantInt16(value0 int16) *Variant {
	var value1 C.int16_t
	value1 = C.int16_t(value0)
	ret1 := C.g_variant_new_int16(value1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantInt32(value0 int32) *Variant {
	var value1 C.int32_t
	value1 = C.int32_t(value0)
	ret1 := C.g_variant_new_int32(value1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantInt64(value0 int64) *Variant {
	var value1 C.int64_t
	value1 = C.int64_t(value0)
	ret1 := C.g_variant_new_int64(value1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantMaybe(child_type0 *VariantType, child0 *Variant) *Variant {
	var child_type1 *C.GVariantType
	var child1 *C.GVariant
	child_type1 = (*C.GVariantType)(unsafe.Pointer(child_type0))
	child1 = (*C.GVariant)(unsafe.Pointer(child0))
	ret1 := C.g_variant_new_maybe(child_type1, child1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantObjectPath(object_path0 string) *Variant {
	var object_path1 *C.char
	object_path1 = _GoStringToGString(object_path0)
	defer C.free(unsafe.Pointer(object_path1))
	ret1 := C.g_variant_new_object_path(object_path1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantObjv(strv0 []string) *Variant {
	var strv1 **C.char
	var length1 C.int64_t
	strv1 = (**C.char)(C.malloc(C.size_t(int(unsafe.Sizeof(*strv1)) * len(strv0))))
	defer C.free(unsafe.Pointer(strv1))
	for i, e := range strv0 {
		(*(*[999999]*C.char)(unsafe.Pointer(strv1)))[i] = _GoStringToGString(e)
		defer C.free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(strv1)))[i]))
	}
	length1 = C.int64_t(len(strv0))
	ret1 := C.g_variant_new_objv(strv1, length1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantSignature(signature0 string) *Variant {
	var signature1 *C.char
	signature1 = _GoStringToGString(signature0)
	defer C.free(unsafe.Pointer(signature1))
	ret1 := C.g_variant_new_signature(signature1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantString(string0 string) *Variant {
	var string1 *C.char
	string1 = _GoStringToGString(string0)
	defer C.free(unsafe.Pointer(string1))
	ret1 := C.g_variant_new_string(string1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantStrv(strv0 []string) *Variant {
	var strv1 **C.char
	var length1 C.int64_t
	strv1 = (**C.char)(C.malloc(C.size_t(int(unsafe.Sizeof(*strv1)) * len(strv0))))
	defer C.free(unsafe.Pointer(strv1))
	for i, e := range strv0 {
		(*(*[999999]*C.char)(unsafe.Pointer(strv1)))[i] = _GoStringToGString(e)
		defer C.free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(strv1)))[i]))
	}
	length1 = C.int64_t(len(strv0))
	ret1 := C.g_variant_new_strv(strv1, length1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantTuple(children0 []*Variant) *Variant {
	var children1 **C.GVariant
	var n_children1 C.uint64_t
	children1 = (**C.GVariant)(C.malloc(C.size_t(int(unsafe.Sizeof(*children1)) * len(children0))))
	defer C.free(unsafe.Pointer(children1))
	for i, e := range children0 {
		(*(*[999999]*C.GVariant)(unsafe.Pointer(children1)))[i] = (*C.GVariant)(unsafe.Pointer(e))
	}
	n_children1 = C.uint64_t(len(children0))
	ret1 := C.g_variant_new_tuple(children1, n_children1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantUint16(value0 uint16) *Variant {
	var value1 C.uint16_t
	value1 = C.uint16_t(value0)
	ret1 := C.g_variant_new_uint16(value1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantUint32(value0 uint32) *Variant {
	var value1 C.uint32_t
	value1 = C.uint32_t(value0)
	ret1 := C.g_variant_new_uint32(value1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantUint64(value0 uint64) *Variant {
	var value1 C.uint64_t
	value1 = C.uint64_t(value0)
	ret1 := C.g_variant_new_uint64(value1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantVariant(value0 *Variant) *Variant {
	var value1 *C.GVariant
	value1 = (*C.GVariant)(unsafe.Pointer(value0))
	ret1 := C.g_variant_new_variant(value1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Variant) Byteswap() *Variant {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_byteswap(this1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Variant) CheckFormatString(format_string0 string, copy_only0 bool) bool {
	var this1 *C.GVariant
	var format_string1 *C.char
	var copy_only1 C.int
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	format_string1 = _GoStringToGString(format_string0)
	defer C.free(unsafe.Pointer(format_string1))
	copy_only1 = _GoBoolToCBool(copy_only0)
	ret1 := C.g_variant_check_format_string(this1, format_string1, copy_only1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Variant) Classify() VariantClass {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_classify(this1)
	var ret2 VariantClass
	ret2 = VariantClass(ret1)
	return ret2
}
func (this0 *Variant) Compare(two0 *Variant) int32 {
	var this1 *C.GVariant
	var two1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	two1 = (*C.GVariant)(unsafe.Pointer(two0))
	ret1 := C.g_variant_compare(this1, two1)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func (this0 *Variant) DupBytestring() (uint64, []uint8) {
	var this1 *C.GVariant
	var length1 C.uint64_t
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_dup_bytestring(this1, &length1)
	var length2 uint64
	var ret2 []uint8
	length2 = uint64(length1)
	ret2 = make([]uint8, length1)
	for i := range ret2 {
		ret2[i] = uint8((*(*[999999]C.uint8_t)(unsafe.Pointer(ret1)))[i])
	}
	return length2, ret2
}
func (this0 *Variant) DupBytestringArray() (uint64, []string) {
	var this1 *C.GVariant
	var length1 C.uint64_t
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_dup_bytestring_array(this1, &length1)
	var length2 uint64
	var ret2 []string
	length2 = uint64(length1)
	ret2 = make([]string, length1)
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return length2, ret2
}
func (this0 *Variant) DupObjv() (uint64, []string) {
	var this1 *C.GVariant
	var length1 C.uint64_t
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_dup_objv(this1, &length1)
	var length2 uint64
	var ret2 []string
	length2 = uint64(length1)
	ret2 = make([]string, length1)
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return length2, ret2
}
func (this0 *Variant) DupString() (uint64, string) {
	var this1 *C.GVariant
	var length1 C.uint64_t
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_dup_string(this1, &length1)
	var length2 uint64
	var ret2 string
	length2 = uint64(length1)
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return length2, ret2
}
func (this0 *Variant) DupStrv() (uint64, []string) {
	var this1 *C.GVariant
	var length1 C.uint64_t
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_dup_strv(this1, &length1)
	var length2 uint64
	var ret2 []string
	length2 = uint64(length1)
	ret2 = make([]string, length1)
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return length2, ret2
}
func (this0 *Variant) Equal(two0 *Variant) bool {
	var this1 *C.GVariant
	var two1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	two1 = (*C.GVariant)(unsafe.Pointer(two0))
	ret1 := C.g_variant_equal(this1, two1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Variant) GetBoolean() bool {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_boolean(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Variant) GetByte() uint8 {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_byte(this1)
	var ret2 uint8
	ret2 = uint8(ret1)
	return ret2
}
func (this0 *Variant) GetBytestring() []uint8 {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_bytestring(this1)
	var ret2 []uint8
	ret2 = make([]uint8, uint(C._array_length(unsafe.Pointer(ret1))))
	for i := range ret2 {
		ret2[i] = uint8((*(*[999999]C.uint8_t)(unsafe.Pointer(ret1)))[i])
	}
	return ret2
}
func (this0 *Variant) GetBytestringArray() (uint64, []string) {
	var this1 *C.GVariant
	var length1 C.uint64_t
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_bytestring_array(this1, &length1)
	var length2 uint64
	var ret2 []string
	length2 = uint64(length1)
	ret2 = make([]string, length1)
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
	}
	return length2, ret2
}
func (this0 *Variant) GetChildValue(index_0 uint64) *Variant {
	var this1 *C.GVariant
	var index_1 C.uint64_t
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	index_1 = C.uint64_t(index_0)
	ret1 := C.g_variant_get_child_value(this1, index_1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Variant) GetData() {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	C.g_variant_get_data(this1)
}
func (this0 *Variant) GetDouble() float64 {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_double(this1)
	var ret2 float64
	ret2 = float64(ret1)
	return ret2
}
func (this0 *Variant) GetHandle() int32 {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_handle(this1)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func (this0 *Variant) GetInt16() int16 {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_int16(this1)
	var ret2 int16
	ret2 = int16(ret1)
	return ret2
}
func (this0 *Variant) GetInt32() int32 {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_int32(this1)
	var ret2 int32
	ret2 = int32(ret1)
	return ret2
}
func (this0 *Variant) GetInt64() int64 {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_int64(this1)
	var ret2 int64
	ret2 = int64(ret1)
	return ret2
}
func (this0 *Variant) GetMaybe() *Variant {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_maybe(this1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Variant) GetNormalForm() *Variant {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_normal_form(this1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Variant) GetObjv() (uint64, []string) {
	var this1 *C.GVariant
	var length1 C.uint64_t
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_objv(this1, &length1)
	var length2 uint64
	var ret2 []string
	length2 = uint64(length1)
	ret2 = make([]string, length1)
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
	}
	return length2, ret2
}
func (this0 *Variant) GetSize() uint64 {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_size(this1)
	var ret2 uint64
	ret2 = uint64(ret1)
	return ret2
}
func (this0 *Variant) GetString() (uint64, string) {
	var this1 *C.GVariant
	var length1 C.uint64_t
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_string(this1, &length1)
	var length2 uint64
	var ret2 string
	length2 = uint64(length1)
	ret2 = C.GoString(ret1)
	return length2, ret2
}
func (this0 *Variant) GetStrv() (uint64, []string) {
	var this1 *C.GVariant
	var length1 C.uint64_t
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_strv(this1, &length1)
	var length2 uint64
	var ret2 []string
	length2 = uint64(length1)
	ret2 = make([]string, length1)
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
	}
	return length2, ret2
}
func (this0 *Variant) GetType() *VariantType {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_type(this1)
	var ret2 *VariantType
	ret2 = (*VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Variant) GetTypeString() string {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_type_string(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *Variant) GetUint16() uint16 {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_uint16(this1)
	var ret2 uint16
	ret2 = uint16(ret1)
	return ret2
}
func (this0 *Variant) GetUint32() uint32 {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_uint32(this1)
	var ret2 uint32
	ret2 = uint32(ret1)
	return ret2
}
func (this0 *Variant) GetUint64() uint64 {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_uint64(this1)
	var ret2 uint64
	ret2 = uint64(ret1)
	return ret2
}
func (this0 *Variant) GetVariant() *Variant {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_get_variant(this1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Variant) Hash() uint32 {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_hash(this1)
	var ret2 uint32
	ret2 = uint32(ret1)
	return ret2
}
func (this0 *Variant) IsContainer() bool {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_is_container(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Variant) IsFloating() bool {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_is_floating(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Variant) IsNormalForm() bool {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_is_normal_form(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Variant) IsOfType(type0 *VariantType) bool {
	var this1 *C.GVariant
	var type1 *C.GVariantType
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	type1 = (*C.GVariantType)(unsafe.Pointer(type0))
	ret1 := C.g_variant_is_of_type(this1, type1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Variant) LookupValue(key0 string, expected_type0 *VariantType) *Variant {
	var this1 *C.GVariant
	var key1 *C.char
	var expected_type1 *C.GVariantType
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	expected_type1 = (*C.GVariantType)(unsafe.Pointer(expected_type0))
	ret1 := C.g_variant_lookup_value(this1, key1, expected_type1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Variant) NChildren() uint64 {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_n_children(this1)
	var ret2 uint64
	ret2 = uint64(ret1)
	return ret2
}
func (this0 *Variant) Print(type_annotate0 bool) string {
	var this1 *C.GVariant
	var type_annotate1 C.int
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	type_annotate1 = _GoBoolToCBool(type_annotate0)
	ret1 := C.g_variant_print(this1, type_annotate1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Variant) RefSink() *Variant {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_ref_sink(this1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Variant) Store(data0 unsafe.Pointer) {
	var this1 *C.GVariant
	var data1 unsafe.Pointer
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	data1 = unsafe.Pointer(data0)
	C.g_variant_store(this1, data1)
}
func (this0 *Variant) TakeRef() *Variant {
	var this1 *C.GVariant
	this1 = (*C.GVariant)(unsafe.Pointer(this0))
	ret1 := C.g_variant_take_ref(this1)
	var ret2 *Variant
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	return ret2
}
func VariantIsObjectPath(string0 string) bool {
	var string1 *C.char
	string1 = _GoStringToGString(string0)
	defer C.free(unsafe.Pointer(string1))
	ret1 := C.g_variant_is_object_path(string1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func VariantIsSignature(string0 string) bool {
	var string1 *C.char
	string1 = _GoStringToGString(string0)
	defer C.free(unsafe.Pointer(string1))
	ret1 := C.g_variant_is_signature(string1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func VariantParse(type0 *VariantType, text0 string, limit0 string, endptr0 string) (*Variant, error) {
	var type1 *C.GVariantType
	var text1 *C.char
	var limit1 *C.char
	var endptr1 *C.char
	var err1 *C.GError
	type1 = (*C.GVariantType)(unsafe.Pointer(type0))
	text1 = _GoStringToGString(text0)
	defer C.free(unsafe.Pointer(text1))
	limit1 = _GoStringToGString(limit0)
	defer C.free(unsafe.Pointer(limit1))
	endptr1 = _GoStringToGString(endptr0)
	defer C.free(unsafe.Pointer(endptr1))
	ret1 := C.g_variant_parse(type1, text1, limit1, endptr1, &err1)
	var ret2 *Variant
	var err2 error
	ret2 = (*Variant)(unsafe.Pointer(ret1))
	if err1 != nil {
		err2 = ((*_GError)(unsafe.Pointer(err1))).ToGError()
		C.g_error_free(err1)
	}
	return ret2, err2
}
func VariantParseErrorPrintContext(error0 error, source_str0 string) string {
	var error1 *C.GError
	var source_str1 *C.char
	//NOTEO: hasn't implemnt GSLIST/GHASH/GERROR convert.
	
	source_str1 = _GoStringToGString(source_str0)
	defer C.free(unsafe.Pointer(source_str1))
	ret1 := C.g_variant_parse_error_print_context(error1, source_str1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func VariantParseErrorQuark() uint32 {
	ret1 := C.g_variant_parse_error_quark()
	var ret2 uint32
	ret2 = uint32(ret1)
	return ret2
}
func VariantParserGetErrorQuark() uint32 {
	ret1 := C.g_variant_parser_get_error_quark()
	var ret2 uint32
	ret2 = uint32(ret1)
	return ret2
}
// blacklisted: VariantBuilder (struct)
type VariantClass C.uint32_t
const (
	VariantClassBoolean VariantClass = 98
	VariantClassByte VariantClass = 121
	VariantClassInt16 VariantClass = 110
	VariantClassUint16 VariantClass = 113
	VariantClassInt32 VariantClass = 105
	VariantClassUint32 VariantClass = 117
	VariantClassInt64 VariantClass = 120
	VariantClassUint64 VariantClass = 116
	VariantClassHandle VariantClass = 104
	VariantClassDouble VariantClass = 100
	VariantClassString VariantClass = 115
	VariantClassObjectPath VariantClass = 111
	VariantClassSignature VariantClass = 103
	VariantClassVariant VariantClass = 118
	VariantClassMaybe VariantClass = 109
	VariantClassArray VariantClass = 97
	VariantClassTuple VariantClass = 40
	VariantClassDictEntry VariantClass = 123
)
// blacklisted: VariantDict (struct)
type VariantParseError C.uint32_t
const (
	VariantParseErrorFailed VariantParseError = 0
	VariantParseErrorBasicTypeExpected VariantParseError = 1
	VariantParseErrorCannotInferType VariantParseError = 2
	VariantParseErrorDefiniteTypeExpected VariantParseError = 3
	VariantParseErrorInputNotAtEnd VariantParseError = 4
	VariantParseErrorInvalidCharacter VariantParseError = 5
	VariantParseErrorInvalidFormatString VariantParseError = 6
	VariantParseErrorInvalidObjectPath VariantParseError = 7
	VariantParseErrorInvalidSignature VariantParseError = 8
	VariantParseErrorInvalidTypeString VariantParseError = 9
	VariantParseErrorNoCommonType VariantParseError = 10
	VariantParseErrorNumberOutOfRange VariantParseError = 11
	VariantParseErrorNumberTooBig VariantParseError = 12
	VariantParseErrorTypeError VariantParseError = 13
	VariantParseErrorUnexpectedToken VariantParseError = 14
	VariantParseErrorUnknownKeyword VariantParseError = 15
	VariantParseErrorUnterminatedStringConstant VariantParseError = 16
	VariantParseErrorValueExpected VariantParseError = 17
)
type VariantType struct {}
func NewVariantType(type_string0 string) *VariantType {
	var type_string1 *C.char
	type_string1 = _GoStringToGString(type_string0)
	defer C.free(unsafe.Pointer(type_string1))
	ret1 := C.g_variant_type_new(type_string1)
	var ret2 *VariantType
	ret2 = (*VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantTypeArray(element0 *VariantType) *VariantType {
	var element1 *C.GVariantType
	element1 = (*C.GVariantType)(unsafe.Pointer(element0))
	ret1 := C.g_variant_type_new_array(element1)
	var ret2 *VariantType
	ret2 = (*VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantTypeDictEntry(key0 *VariantType, value0 *VariantType) *VariantType {
	var key1 *C.GVariantType
	var value1 *C.GVariantType
	key1 = (*C.GVariantType)(unsafe.Pointer(key0))
	value1 = (*C.GVariantType)(unsafe.Pointer(value0))
	ret1 := C.g_variant_type_new_dict_entry(key1, value1)
	var ret2 *VariantType
	ret2 = (*VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantTypeMaybe(element0 *VariantType) *VariantType {
	var element1 *C.GVariantType
	element1 = (*C.GVariantType)(unsafe.Pointer(element0))
	ret1 := C.g_variant_type_new_maybe(element1)
	var ret2 *VariantType
	ret2 = (*VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func NewVariantTypeTuple(items0 []*VariantType) *VariantType {
	var items1 **C.GVariantType
	var length1 C.int32_t
	items1 = (**C.GVariantType)(C.malloc(C.size_t(int(unsafe.Sizeof(*items1)) * len(items0))))
	defer C.free(unsafe.Pointer(items1))
	for i, e := range items0 {
		(*(*[999999]*C.GVariantType)(unsafe.Pointer(items1)))[i] = (*C.GVariantType)(unsafe.Pointer(e))
	}
	length1 = C.int32_t(len(items0))
	ret1 := C.g_variant_type_new_tuple(items1, length1)
	var ret2 *VariantType
	ret2 = (*VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *VariantType) Copy() *VariantType {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_copy(this1)
	var ret2 *VariantType
	ret2 = (*VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *VariantType) DupString() string {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_dup_string(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *VariantType) Element() *VariantType {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_element(this1)
	var ret2 *VariantType
	ret2 = (*VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *VariantType) Equal(type20 *VariantType) bool {
	var this1 *C.GVariantType
	var type21 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	type21 = (*C.GVariantType)(unsafe.Pointer(type20))
	ret1 := C.g_variant_type_equal(this1, type21)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *VariantType) First() *VariantType {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_first(this1)
	var ret2 *VariantType
	ret2 = (*VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *VariantType) Free() {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	C.g_variant_type_free(this1)
}
func (this0 *VariantType) GetStringLength() uint64 {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_get_string_length(this1)
	var ret2 uint64
	ret2 = uint64(ret1)
	return ret2
}
func (this0 *VariantType) Hash() uint32 {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_hash(this1)
	var ret2 uint32
	ret2 = uint32(ret1)
	return ret2
}
func (this0 *VariantType) IsArray() bool {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_is_array(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *VariantType) IsBasic() bool {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_is_basic(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *VariantType) IsContainer() bool {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_is_container(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *VariantType) IsDefinite() bool {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_is_definite(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *VariantType) IsDictEntry() bool {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_is_dict_entry(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *VariantType) IsMaybe() bool {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_is_maybe(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *VariantType) IsSubtypeOf(supertype0 *VariantType) bool {
	var this1 *C.GVariantType
	var supertype1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	supertype1 = (*C.GVariantType)(unsafe.Pointer(supertype0))
	ret1 := C.g_variant_type_is_subtype_of(this1, supertype1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *VariantType) IsTuple() bool {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_is_tuple(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *VariantType) IsVariant() bool {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_is_variant(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *VariantType) Key() *VariantType {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_key(this1)
	var ret2 *VariantType
	ret2 = (*VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *VariantType) NItems() uint64 {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_n_items(this1)
	var ret2 uint64
	ret2 = uint64(ret1)
	return ret2
}
func (this0 *VariantType) Next() *VariantType {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_next(this1)
	var ret2 *VariantType
	ret2 = (*VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *VariantType) Value() *VariantType {
	var this1 *C.GVariantType
	this1 = (*C.GVariantType)(unsafe.Pointer(this0))
	ret1 := C.g_variant_type_value(this1)
	var ret2 *VariantType
	ret2 = (*VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func VariantTypeChecked_(arg00 string) *VariantType {
	var arg01 *C.char
	arg01 = _GoStringToGString(arg00)
	defer C.free(unsafe.Pointer(arg01))
	ret1 := C.g_variant_type_checked_(arg01)
	var ret2 *VariantType
	ret2 = (*VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func VariantTypeStringIsValid(type_string0 string) bool {
	var type_string1 *C.char
	type_string1 = _GoStringToGString(type_string0)
	defer C.free(unsafe.Pointer(type_string1))
	ret1 := C.g_variant_type_string_is_valid(type_string1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func VariantTypeStringScan(string0 string, limit0 string) (string, bool) {
	var string1 *C.char
	var limit1 *C.char
	var endptr1 *C.char
	string1 = _GoStringToGString(string0)
	defer C.free(unsafe.Pointer(string1))
	limit1 = _GoStringToGString(limit0)
	defer C.free(unsafe.Pointer(limit1))
	ret1 := C.g_variant_type_string_scan(string1, limit1, &endptr1)
	var endptr2 string
	var ret2 bool
	endptr2 = C.GoString(endptr1)
	C.g_free(unsafe.Pointer(endptr1))
	ret2 = ret1 != 0
	return endptr2, ret2
}
// blacklisted: VoidFunc (callback)
const Win32MsgHandle = 19981206
// blacklisted: access (function)
// blacklisted: ascii_digit_value (function)
// blacklisted: ascii_dtostr (function)
// blacklisted: ascii_formatd (function)
// blacklisted: ascii_strcasecmp (function)
// blacklisted: ascii_strdown (function)
// blacklisted: ascii_strncasecmp (function)
// blacklisted: ascii_strtod (function)
// blacklisted: ascii_strtoll (function)
// blacklisted: ascii_strtoull (function)
// blacklisted: ascii_strup (function)
// blacklisted: ascii_tolower (function)
// blacklisted: ascii_toupper (function)
// blacklisted: ascii_xdigit_value (function)
// blacklisted: assert_warning (function)
// blacklisted: assertion_message (function)
// blacklisted: assertion_message_cmpstr (function)
// blacklisted: assertion_message_error (function)
// blacklisted: assertion_message_expr (function)
// blacklisted: atexit (function)
// blacklisted: atomic_int_add (function)
// blacklisted: atomic_int_and (function)
// blacklisted: atomic_int_compare_and_exchange (function)
// blacklisted: atomic_int_dec_and_test (function)
// blacklisted: atomic_int_exchange_and_add (function)
// blacklisted: atomic_int_get (function)
// blacklisted: atomic_int_inc (function)
// blacklisted: atomic_int_or (function)
// blacklisted: atomic_int_set (function)
// blacklisted: atomic_int_xor (function)
// blacklisted: atomic_pointer_add (function)
// blacklisted: atomic_pointer_and (function)
// blacklisted: atomic_pointer_compare_and_exchange (function)
// blacklisted: atomic_pointer_or (function)
// blacklisted: atomic_pointer_set (function)
// blacklisted: atomic_pointer_xor (function)
// blacklisted: base64_decode (function)
// blacklisted: base64_decode_inplace (function)
// blacklisted: base64_decode_step (function)
// blacklisted: base64_encode (function)
// blacklisted: base64_encode_close (function)
// blacklisted: base64_encode_step (function)
// blacklisted: basename (function)
// blacklisted: bit_lock (function)
// blacklisted: bit_nth_lsf (function)
// blacklisted: bit_nth_msf (function)
// blacklisted: bit_storage (function)
// blacklisted: bit_trylock (function)
// blacklisted: bit_unlock (function)
// blacklisted: bookmark_file_error_quark (function)
// blacklisted: build_filenamev (function)
// blacklisted: build_pathv (function)
// blacklisted: byte_array_free (function)
// blacklisted: byte_array_free_to_bytes (function)
// blacklisted: byte_array_new (function)
// blacklisted: byte_array_new_take (function)
// blacklisted: byte_array_unref (function)
// blacklisted: chdir (function)
// blacklisted: check_version (function)
// blacklisted: checksum_type_get_length (function)
// blacklisted: child_watch_add (function)
// blacklisted: child_watch_source_new (function)
// blacklisted: clear_error (function)
// blacklisted: close (function)
// blacklisted: compute_checksum_for_bytes (function)
// blacklisted: compute_checksum_for_data (function)
// blacklisted: compute_checksum_for_string (function)
// blacklisted: compute_hmac_for_data (function)
// blacklisted: compute_hmac_for_string (function)
// blacklisted: convert (function)
// blacklisted: convert_error_quark (function)
// blacklisted: convert_with_fallback (function)
// blacklisted: convert_with_iconv (function)
// blacklisted: datalist_clear (function)
// blacklisted: datalist_get_flags (function)
// blacklisted: datalist_id_replace_data (function)
// blacklisted: datalist_id_set_data_full (function)
// blacklisted: datalist_init (function)
// blacklisted: datalist_set_flags (function)
// blacklisted: datalist_unset_flags (function)
// blacklisted: dataset_destroy (function)
// blacklisted: dataset_id_set_data_full (function)
// blacklisted: date_get_days_in_month (function)
// blacklisted: date_get_monday_weeks_in_year (function)
// blacklisted: date_get_sunday_weeks_in_year (function)
// blacklisted: date_is_leap_year (function)
// blacklisted: date_strftime (function)
// blacklisted: date_time_compare (function)
// blacklisted: date_time_equal (function)
// blacklisted: date_time_hash (function)
// blacklisted: date_valid_day (function)
// blacklisted: date_valid_dmy (function)
// blacklisted: date_valid_julian (function)
// blacklisted: date_valid_month (function)
// blacklisted: date_valid_weekday (function)
// blacklisted: date_valid_year (function)
// blacklisted: dcgettext (function)
// blacklisted: dgettext (function)
// blacklisted: dir_make_tmp (function)
// blacklisted: direct_equal (function)
// blacklisted: direct_hash (function)
// blacklisted: dngettext (function)
// blacklisted: double_equal (function)
// blacklisted: double_hash (function)
// blacklisted: dpgettext (function)
// blacklisted: dpgettext2 (function)
// blacklisted: environ_getenv (function)
// blacklisted: environ_setenv (function)
// blacklisted: environ_unsetenv (function)
// blacklisted: file_error_from_errno (function)
// blacklisted: file_error_quark (function)
// blacklisted: file_get_contents (function)
// blacklisted: file_open_tmp (function)
// blacklisted: file_read_link (function)
// blacklisted: file_set_contents (function)
// blacklisted: file_test (function)
// blacklisted: filename_display_basename (function)
// blacklisted: filename_display_name (function)
// blacklisted: filename_from_uri (function)
// blacklisted: filename_from_utf8 (function)
// blacklisted: filename_to_uri (function)
// blacklisted: filename_to_utf8 (function)
// blacklisted: find_program_in_path (function)
// blacklisted: format_size (function)
// blacklisted: format_size_for_display (function)
// blacklisted: format_size_full (function)
// blacklisted: free (function)
// blacklisted: get_application_name (function)
// blacklisted: get_charset (function)
// blacklisted: get_codeset (function)
func GetCurrentDir() string {
	ret1 := C.g_get_current_dir()
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
// blacklisted: get_current_time (function)
// blacklisted: get_environ (function)
// blacklisted: get_filename_charsets (function)
func GetHomeDir() string {
	ret1 := C.g_get_home_dir()
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
// blacklisted: get_host_name (function)
// blacklisted: get_language_names (function)
// blacklisted: get_locale_variants (function)
// blacklisted: get_monotonic_time (function)
// blacklisted: get_num_processors (function)
// blacklisted: get_prgname (function)
// blacklisted: get_real_name (function)
// blacklisted: get_real_time (function)
func GetSystemConfigDirs() []string {
	ret1 := C.g_get_system_config_dirs()
	var ret2 []string
	ret2 = make([]string, uint(C._array_length(unsafe.Pointer(ret1))))
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
	}
	return ret2
}
func GetSystemDataDirs() []string {
	ret1 := C.g_get_system_data_dirs()
	var ret2 []string
	ret2 = make([]string, uint(C._array_length(unsafe.Pointer(ret1))))
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
	}
	return ret2
}
func GetTmpDir() string {
	ret1 := C.g_get_tmp_dir()
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func GetUserCacheDir() string {
	ret1 := C.g_get_user_cache_dir()
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func GetUserConfigDir() string {
	ret1 := C.g_get_user_config_dir()
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func GetUserDataDir() string {
	ret1 := C.g_get_user_data_dir()
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
// blacklisted: get_user_name (function)
func GetUserRuntimeDir() string {
	ret1 := C.g_get_user_runtime_dir()
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func GetUserSpecialDir(directory0 UserDirectory) string {
	var directory1 C.GUserDirectory
	directory1 = C.GUserDirectory(directory0)
	ret1 := C.g_get_user_special_dir(directory1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
// blacklisted: getenv (function)
// blacklisted: hash_table_add (function)
// blacklisted: hash_table_contains (function)
// blacklisted: hash_table_destroy (function)
// blacklisted: hash_table_insert (function)
// blacklisted: hash_table_lookup_extended (function)
// blacklisted: hash_table_remove (function)
// blacklisted: hash_table_remove_all (function)
// blacklisted: hash_table_replace (function)
// blacklisted: hash_table_size (function)
// blacklisted: hash_table_steal (function)
// blacklisted: hash_table_steal_all (function)
// blacklisted: hash_table_unref (function)
// blacklisted: hook_destroy (function)
// blacklisted: hook_destroy_link (function)
// blacklisted: hook_free (function)
// blacklisted: hook_insert_before (function)
// blacklisted: hook_prepend (function)
// blacklisted: hook_unref (function)
// blacklisted: hostname_is_ascii_encoded (function)
// blacklisted: hostname_is_ip_address (function)
// blacklisted: hostname_is_non_ascii (function)
// blacklisted: hostname_to_ascii (function)
// blacklisted: hostname_to_unicode (function)
// blacklisted: iconv (function)
// blacklisted: idle_add (function)
// blacklisted: idle_remove_by_data (function)
// blacklisted: idle_source_new (function)
// blacklisted: int64_equal (function)
// blacklisted: int64_hash (function)
// blacklisted: int_equal (function)
// blacklisted: int_hash (function)
// blacklisted: intern_static_string (function)
// blacklisted: intern_string (function)
// blacklisted: io_add_watch (function)
// blacklisted: io_channel_error_from_errno (function)
// blacklisted: io_channel_error_quark (function)
// blacklisted: io_create_watch (function)
// blacklisted: key_file_error_quark (function)
// blacklisted: listenv (function)
// blacklisted: locale_from_utf8 (function)
// blacklisted: locale_to_utf8 (function)
// blacklisted: log_default_handler (function)
// blacklisted: log_remove_handler (function)
// blacklisted: log_set_always_fatal (function)
// blacklisted: log_set_fatal_mask (function)
// blacklisted: main_context_default (function)
// blacklisted: main_context_get_thread_default (function)
// blacklisted: main_context_ref_thread_default (function)
// blacklisted: main_current_source (function)
// blacklisted: main_depth (function)
// blacklisted: markup_error_quark (function)
// blacklisted: markup_escape_text (function)
// blacklisted: mem_is_system_malloc (function)
// blacklisted: mem_profile (function)
// blacklisted: mem_set_vtable (function)
// blacklisted: mkdir_with_parents (function)
// blacklisted: mkdtemp (function)
// blacklisted: mkdtemp_full (function)
// blacklisted: mkstemp (function)
// blacklisted: mkstemp_full (function)
// blacklisted: nullify_pointer (function)
// blacklisted: on_error_query (function)
// blacklisted: on_error_stack_trace (function)
// blacklisted: once_init_enter (function)
// blacklisted: once_init_leave (function)
// blacklisted: option_error_quark (function)
// blacklisted: parse_debug_string (function)
// blacklisted: path_get_basename (function)
// blacklisted: path_get_dirname (function)
// blacklisted: path_is_absolute (function)
// blacklisted: path_skip_root (function)
// blacklisted: pattern_match (function)
// blacklisted: pattern_match_simple (function)
// blacklisted: pattern_match_string (function)
// blacklisted: pointer_bit_lock (function)
// blacklisted: pointer_bit_trylock (function)
// blacklisted: pointer_bit_unlock (function)
// blacklisted: poll (function)
// blacklisted: propagate_error (function)
// blacklisted: quark_from_static_string (function)
// blacklisted: quark_from_string (function)
// blacklisted: quark_to_string (function)
// blacklisted: quark_try_string (function)
// blacklisted: random_double (function)
// blacklisted: random_double_range (function)
// blacklisted: random_int (function)
// blacklisted: random_int_range (function)
// blacklisted: random_set_seed (function)
// blacklisted: regex_check_replacement (function)
// blacklisted: regex_error_quark (function)
// blacklisted: regex_escape_nul (function)
// blacklisted: regex_escape_string (function)
// blacklisted: regex_match_simple (function)
// blacklisted: regex_split_simple (function)
// blacklisted: reload_user_special_dirs_cache (function)
// blacklisted: return_if_fail_warning (function)
// blacklisted: rmdir (function)
// blacklisted: sequence_move (function)
// blacklisted: sequence_move_range (function)
// blacklisted: sequence_remove (function)
// blacklisted: sequence_remove_range (function)
// blacklisted: sequence_set (function)
// blacklisted: sequence_swap (function)
// blacklisted: set_application_name (function)
// blacklisted: set_error_literal (function)
// blacklisted: set_prgname (function)
// blacklisted: setenv (function)
// blacklisted: shell_error_quark (function)
// blacklisted: shell_parse_argv (function)
// blacklisted: shell_quote (function)
// blacklisted: shell_unquote (function)
// blacklisted: slice_free1 (function)
// blacklisted: slice_free_chain_with_offset (function)
// blacklisted: slice_get_config (function)
// blacklisted: slice_get_config_state (function)
// blacklisted: slice_set_config (function)
// blacklisted: source_remove (function)
// blacklisted: source_remove_by_funcs_user_data (function)
// blacklisted: source_remove_by_user_data (function)
// blacklisted: source_set_name_by_id (function)
// blacklisted: spaced_primes_closest (function)
// blacklisted: spawn_async (function)
// blacklisted: spawn_async_with_pipes (function)
// blacklisted: spawn_check_exit_status (function)
// blacklisted: spawn_close_pid (function)
// blacklisted: spawn_command_line_async (function)
// blacklisted: spawn_command_line_sync (function)
// blacklisted: spawn_error_quark (function)
// blacklisted: spawn_exit_error_quark (function)
// blacklisted: spawn_sync (function)
// blacklisted: stpcpy (function)
// blacklisted: str_equal (function)
// blacklisted: str_has_prefix (function)
// blacklisted: str_has_suffix (function)
// blacklisted: str_hash (function)
// blacklisted: str_is_ascii (function)
// blacklisted: str_match_string (function)
// blacklisted: str_to_ascii (function)
// blacklisted: str_tokenize_and_fold (function)
// blacklisted: strcanon (function)
// blacklisted: strcasecmp (function)
// blacklisted: strchomp (function)
// blacklisted: strchug (function)
// blacklisted: strcmp0 (function)
// blacklisted: strcompress (function)
// blacklisted: strdelimit (function)
// blacklisted: strdown (function)
// blacklisted: strdup (function)
// blacklisted: strerror (function)
// blacklisted: strescape (function)
// blacklisted: strfreev (function)
// blacklisted: string_new (function)
// blacklisted: string_new_len (function)
// blacklisted: string_sized_new (function)
// blacklisted: strip_context (function)
// blacklisted: strjoinv (function)
// blacklisted: strlcat (function)
// blacklisted: strlcpy (function)
// blacklisted: strncasecmp (function)
// blacklisted: strndup (function)
// blacklisted: strnfill (function)
// blacklisted: strreverse (function)
// blacklisted: strrstr (function)
// blacklisted: strrstr_len (function)
// blacklisted: strsignal (function)
// blacklisted: strstr_len (function)
// blacklisted: strtod (function)
// blacklisted: strup (function)
// blacklisted: strv_get_type (function)
// blacklisted: strv_length (function)
// blacklisted: test_add_data_func_full (function)
// blacklisted: test_assert_expected_messages_internal (function)
// blacklisted: test_bug (function)
// blacklisted: test_bug_base (function)
// blacklisted: test_expect_message (function)
// blacklisted: test_fail (function)
// blacklisted: test_failed (function)
// blacklisted: test_get_dir (function)
// blacklisted: test_incomplete (function)
// blacklisted: test_log_type_name (function)
// blacklisted: test_queue_destroy (function)
// blacklisted: test_queue_free (function)
// blacklisted: test_rand_double (function)
// blacklisted: test_rand_double_range (function)
// blacklisted: test_rand_int (function)
// blacklisted: test_rand_int_range (function)
// blacklisted: test_run (function)
// blacklisted: test_run_suite (function)
// blacklisted: test_set_nonfatal_assertions (function)
// blacklisted: test_skip (function)
// blacklisted: test_subprocess (function)
// blacklisted: test_timer_elapsed (function)
// blacklisted: test_timer_last (function)
// blacklisted: test_timer_start (function)
// blacklisted: test_trap_assertions (function)
// blacklisted: test_trap_fork (function)
// blacklisted: test_trap_has_passed (function)
// blacklisted: test_trap_reached_timeout (function)
// blacklisted: test_trap_subprocess (function)
// blacklisted: thread_error_quark (function)
// blacklisted: thread_exit (function)
// blacklisted: thread_pool_get_max_idle_time (function)
// blacklisted: thread_pool_get_max_unused_threads (function)
// blacklisted: thread_pool_get_num_unused_threads (function)
// blacklisted: thread_pool_set_max_idle_time (function)
// blacklisted: thread_pool_set_max_unused_threads (function)
// blacklisted: thread_pool_stop_unused_threads (function)
// blacklisted: thread_self (function)
// blacklisted: thread_yield (function)
// blacklisted: time_val_from_iso8601 (function)
// blacklisted: timeout_add (function)
// blacklisted: timeout_add_seconds (function)
// blacklisted: timeout_source_new (function)
// blacklisted: timeout_source_new_seconds (function)
// blacklisted: trash_stack_height (function)
// blacklisted: trash_stack_push (function)
// blacklisted: ucs4_to_utf16 (function)
// blacklisted: ucs4_to_utf8 (function)
// blacklisted: unichar_break_type (function)
// blacklisted: unichar_combining_class (function)
// blacklisted: unichar_compose (function)
// blacklisted: unichar_decompose (function)
// blacklisted: unichar_digit_value (function)
// blacklisted: unichar_fully_decompose (function)
// blacklisted: unichar_get_mirror_char (function)
// blacklisted: unichar_get_script (function)
// blacklisted: unichar_isalnum (function)
// blacklisted: unichar_isalpha (function)
// blacklisted: unichar_iscntrl (function)
// blacklisted: unichar_isdefined (function)
// blacklisted: unichar_isdigit (function)
// blacklisted: unichar_isgraph (function)
// blacklisted: unichar_islower (function)
// blacklisted: unichar_ismark (function)
// blacklisted: unichar_isprint (function)
// blacklisted: unichar_ispunct (function)
// blacklisted: unichar_isspace (function)
// blacklisted: unichar_istitle (function)
// blacklisted: unichar_isupper (function)
// blacklisted: unichar_iswide (function)
// blacklisted: unichar_iswide_cjk (function)
// blacklisted: unichar_isxdigit (function)
// blacklisted: unichar_iszerowidth (function)
// blacklisted: unichar_to_utf8 (function)
// blacklisted: unichar_tolower (function)
// blacklisted: unichar_totitle (function)
// blacklisted: unichar_toupper (function)
// blacklisted: unichar_type (function)
// blacklisted: unichar_validate (function)
// blacklisted: unichar_xdigit_value (function)
// blacklisted: unicode_canonical_decomposition (function)
// blacklisted: unicode_canonical_ordering (function)
// blacklisted: unicode_script_from_iso15924 (function)
// blacklisted: unicode_script_to_iso15924 (function)
// blacklisted: unix_error_quark (function)
// blacklisted: unix_fd_add_full (function)
// blacklisted: unix_fd_source_new (function)
// blacklisted: unix_open_pipe (function)
// blacklisted: unix_set_fd_nonblocking (function)
// blacklisted: unix_signal_add (function)
// blacklisted: unix_signal_source_new (function)
// blacklisted: unlink (function)
// blacklisted: unsetenv (function)
// blacklisted: uri_escape_string (function)
// blacklisted: uri_list_extract_uris (function)
// blacklisted: uri_parse_scheme (function)
// blacklisted: uri_unescape_segment (function)
// blacklisted: uri_unescape_string (function)
// blacklisted: usleep (function)
// blacklisted: utf16_to_ucs4 (function)
// blacklisted: utf16_to_utf8 (function)
// blacklisted: utf8_casefold (function)
// blacklisted: utf8_collate (function)
// blacklisted: utf8_collate_key (function)
// blacklisted: utf8_collate_key_for_filename (function)
// blacklisted: utf8_find_next_char (function)
// blacklisted: utf8_find_prev_char (function)
// blacklisted: utf8_get_char (function)
// blacklisted: utf8_get_char_validated (function)
// blacklisted: utf8_normalize (function)
// blacklisted: utf8_offset_to_pointer (function)
// blacklisted: utf8_pointer_to_offset (function)
// blacklisted: utf8_prev_char (function)
// blacklisted: utf8_strchr (function)
// blacklisted: utf8_strdown (function)
// blacklisted: utf8_strlen (function)
// blacklisted: utf8_strncpy (function)
// blacklisted: utf8_strrchr (function)
// blacklisted: utf8_strreverse (function)
// blacklisted: utf8_strup (function)
// blacklisted: utf8_substring (function)
// blacklisted: utf8_to_ucs4 (function)
// blacklisted: utf8_to_ucs4_fast (function)
// blacklisted: utf8_to_utf16 (function)
// blacklisted: utf8_validate (function)
// blacklisted: variant_get_gtype (function)
// blacklisted: variant_is_object_path (function)
// blacklisted: variant_is_signature (function)
// blacklisted: variant_parse (function)
// blacklisted: variant_parse_error_print_context (function)
// blacklisted: variant_parse_error_quark (function)
// blacklisted: variant_parser_get_error_quark (function)
// blacklisted: variant_type_checked_ (function)
// blacklisted: variant_type_string_is_valid (function)
// blacklisted: variant_type_string_scan (function)
// blacklisted: warn_message (function)




//workaround
func (this0 *KeyFile) Free() {
	var this1 *C.GKeyFile
	this1 = (*C.GKeyFile)(unsafe.Pointer(this0))
	C.g_key_file_free(this1)
}