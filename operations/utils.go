package operations

// #cgo CFLAGS: -std=c99
// #cgo pkg-config: gio-unix-2.0 glib-2.0
// #include "utils.c"
import "C"

import (
	"net/url"
	"pkg.deepin.io/lib/gettext"
	"pkg.deepin.io/lib/gio-2.0"
	"pkg.deepin.io/lib/gobject-2.0"
	"unicode/utf8"
	"unsafe"
)

// Tr is a alias for gettext.Tr, which avoids to use dot import.
var Tr = gettext.Tr

// NTr is used for translating with words which has plural.
func NTr(singular, plural string, num uint64) string {
	if num&1 != 0 {
		return Tr(singular)
	}
	return Tr(plural)
}

func dummy(...interface{}) {
}

func uriToGFile(uri *url.URL) *gio.File {
	if uri.Scheme == "" {
		uri.Scheme = "file"

	}
	return gio.FileNewForUri(uri.String())
}

func locationListFromUriList(uris []string) []*gio.File {
	files := make([]*gio.File, len(uris))
	for i, uri := range uris {
		files[i] = gio.FileNewForCommandlineArg(uri)
	}
	return files
}

func locationListFromUrlList(uris []*url.URL) []*gio.File {
	files := []*gio.File{}
	for _, uri := range uris {
		files = append(files, uriToGFile(uri))
	}
	return files
}

func getMaxNameLength(fileDir *gio.File) int {
	return int(C.get_max_name_length((*C.struct__GFile)(fileDir.C)))
}

func fatStrReplace(str string, replacement rune) (string, bool) {
	cstr := C.CString(str)
	ok := C.fat_str_replace(cstr, C.char(replacement)) == 1
	newStr := C.GoString(cstr)
	C.free(unsafe.Pointer(cstr))
	return newStr, ok
}

func makeFileNameValidForDestFs(filename string, fsType string) (string, bool) {
	cname := C.CString(filename)
	defer C.free(unsafe.Pointer(cname))
	cFsType := C.CString(fsType)
	defer C.free(unsafe.Pointer(cFsType))

	ok := C.make_file_name_valid_for_dest_fs(cname, cFsType) == 1
	return C.GoString(cname), ok
}

func queryFsType(file *gio.File, cancellable *gio.Cancellable) string {
	fsinfo, _ := file.QueryFilesystemInfo(gio.FileAttributeFilesystemType, cancellable)
	fsType := ""
	if fsinfo != nil {
		fsType = fsinfo.GetAttributeString(gio.FileAttributeFilesystemType)
		fsinfo.Unref()
	}

	return fsType
}

// FilenameGetExtensionOffset is a C function wrap which return the offset of the extension.
func FilenameGetExtensionOffset(name string) int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cOffset := C.get_filename_extension_offset(cname)
	offset := int(cOffset)
	return offset
}

// FilenameStripExtension returns the basename without extension name.
func FilenameStripExtension(name string) string {
	filename := name

	offset := FilenameGetExtensionOffset(filename)

	if offset != -1 && filename[offset:] != filename {
		filename = filename[:offset]
	}

	return filename
}

func getUtf8FistValidChar(str []rune) int {
	for i, c := range str {
		if utf8.ValidRune(c) {
			return i
		}
	}

	return -1
}

func shortenUtf8Rune(str []rune, reduceByNumBytes int) []rune {
	if reduceByNumBytes <= 0 {
		return str
	}

	baseLen := len(str)
	baseLen -= reduceByNumBytes

	if baseLen <= 0 {
		return []rune("")
	}

	p := 0
	next := -1
	for baseLen != 0 {
		next = getUtf8FistValidChar(str[p:]) + p
		if next == -1 || next-p > baseLen {
			break
		}

		baseLen -= next + 1 - p
		p = next + 1
	}

	return str[:p]
}

// ShortenUtf8String shortens a utf8 string according to the reduceByNumBytes.
func ShortenUtf8String(str string, reduceByNumBytes int) string {
	if reduceByNumBytes <= 0 {
		return str
	}
	return string(shortenUtf8Rune([]rune(str), reduceByNumBytes))
}

func isReadonlyFileSystem(sourceDir *gio.File) bool {
	if sourceDir != nil {
		defer sourceDir.Unref()
		info, _ := sourceDir.QueryFilesystemInfo(gio.FileAttributeFilesystemReadonly, nil)
		if info != nil {
			defer info.Unref()
			return info.GetAttributeBoolean(gio.FileAttributeFilesystemReadonly)
		}
	}
	return false
}

// HasFsID checks whether the file has the expected filesystem ID.
func HasFsID(file *gio.File, expectedID string) bool {
	info, _ := file.QueryInfo(gio.FileAttributeIdFilesystem, gio.FileQueryInfoFlagsNofollowSymlinks, nil)
	if info != nil {
		defer info.Unref()
		id := info.GetAttributeString(gio.FileAttributeIdFilesystem)
		return id == expectedID
	}
	return false
}

// DirIsParentOf checks whether the root is the parent directory of child.
func DirIsParentOf(root *gio.File, child *gio.File) bool {
	f := child.Dup()
	for f != nil {
		if f.Equal(root) {
			f.Unref()
			return true
		}

		tmp := f
		f = f.GetParent()
		tmp.Unref()
	}

	return false
}

func isDir(file *gio.File) bool {
	info, _ := file.QueryInfo(gio.FileAttributeStandardType, gio.FileQueryInfoFlagsNofollowSymlinks, nil)
	if info != nil {
		defer info.Unref()
		return info.GetFileType() == gio.FileTypeDirectory
	}
	return false
}

func getUniqueTargetFile(src *gio.File, destDir *gio.File, sameFs bool, destFsType string, count int) *gio.File {
	cDestType := C.CString(destFsType)
	defer C.free(unsafe.Pointer(cDestType))
	cSameFs := 0
	if sameFs {
		cSameFs = 1
	}
	gfile := C.get_unique_target_file((*C.struct__GFile)(src.ImplementsGFile()), (*C.struct__GFile)(destDir.ImplementsGFile()), C.gboolean(cSameFs), cDestType, C.int(count))
	file := (*gio.File)(gobject.ObjectWrap(unsafe.Pointer(gfile), false))
	return file
}
