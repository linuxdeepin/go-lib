package operations

// #cgo pkg-config: gtk+-3.0 x11
// #cgo CFLAGS: -std=c99
// #include <X11/Xlib.h>
// #include <stdlib.h>
// int get_can_paste();
// void set_op_content_to_clipboard(int, char**, int);
// char** get_clipboard_content(int*);
// void clear_clipboard();
// void freeStrv(void*);
import "C"
import "unsafe"
import "pkg.deepin.io/lib/gio-2.0"
import "strings"

var (
	cChar    *C.char
	unitSize = unsafe.Sizeof(cChar)
)

const (
	OpCut  = "cut"
	OpCopy = "copy"
)

// !!! make sure XInitThreads is called before gtk_init.
// http://stackoverflow.com/questions/18647475/threading-problems-with-gtk.
func init() {
	C.XInitThreads()
}

func CanPasteClipboardContent() bool {
	return C.get_can_paste() == C.int(1)
}

func CanPasteInto(file string) bool {
	f := gio.FileNewForCommandlineArg(file)
	if f == nil {
		return false
	}
	defer f.Unref()

	info, _ := f.QueryInfo(strings.Join(
		[]string{
			gio.FileAttributeStandardType,
			gio.FileAttributeAccessCanWrite,
			gio.FileAttributeStandardTargetUri,
		}, ","), gio.FileQueryInfoFlagsNone, nil)
	if info == nil {
		return false
	}
	defer info.Unref()

	fileType := info.GetFileType()
	// TODO: check the readonly attribute for filesystem.
	if fileType == gio.FileTypeDirectory && info.GetAttributeBoolean(gio.FileAttributeAccessCanWrite) {
		return true
	}

	targetURI := info.GetAttributeString(gio.FileAttributeStandardTargetUri)
	if targetURI == "" {
		return false
	}

	targetFile := gio.FileNewForUri(targetURI)
	if targetFile == nil {
		return false
	}

	targetInfo, _ := targetFile.QueryInfo(strings.Join(
		[]string{
			gio.FileAttributeStandardType,
			gio.FileAttributeAccessCanWrite,
		}, "+"), gio.FileQueryInfoFlagsNone, nil)

	if targetInfo == nil {
		return false
	}
	defer targetInfo.Unref()

	targetFileType := targetInfo.GetFileType()
	// TODO: check the readonly attribute for filesystem.
	return targetFileType == gio.FileTypeUnknown ||
		(targetFileType == gio.FileTypeDirectory && targetInfo.GetAttributeBoolean(gio.FileAttributeAccessCanWrite))
}

func CanPaste(file string) bool {
	return CanPasteClipboardContent() && CanPasteInto(file)
}

func convertToCStrings(strs []string, strNum int) []*C.char {
	cStrings := make([]*C.char, strNum)
	for i, str := range strs {
		cs := C.CString(str)
		cStrings[i] = cs
	}
	return cStrings
}

func freeStrv(cStrs []*C.char) {
	for _, s := range cStrs {
		C.free(unsafe.Pointer(s))
	}
}

func setOpContentToClipboard(op string, files []string) {
	isCut := 0
	if op == OpCut {
		isCut = 1
	}

	num := len(files)
	cFiles := convertToCStrings(files, num)
	C.set_op_content_to_clipboard(C.int(isCut), (**C.char)(unsafe.Pointer(&cFiles[0])), C.int(num))
	freeStrv(cFiles)

}

func CutToClipboard(files []string) {
	setOpContentToClipboard(OpCut, files)
}

func CopyToClipboard(files []string) {
	setOpContentToClipboard(OpCopy, files)
}

func convertToGoStrings(ptr **C.char, n int) []string {
	goStrings := make([]string, n)
	for i := 0; i < n; i++ {
		cs := *(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + uintptr(i)*unitSize))
		goStrings[i] = C.GoString(cs)
	}

	return goStrings
}

func GetClipboardContents() []string {
	n := 0
	files := C.get_clipboard_content((*C.int)(unsafe.Pointer(&n)))
	defer C.freeStrv(unsafe.Pointer(files))

	if n <= 0 {
		return []string{}
	}

	// files is NULL terminal array, the last element is useless.
	// the first content is op.
	return convertToGoStrings(files, n)
}

func ClearClipboard() {
	C.clear_clipboard()
}
