package gettext

/*
#include <locale.h>
#include <stdlib.h>
#include <libintl.h>
void _init_i18n() { setlocale(LC_ALL, ""); }
*/
import "C"
import "unsafe"

func Tr(id string) string {
	_id := C.CString(id)
	defer C.free(unsafe.Pointer(_id))
	return C.GoString(C.gettext(_id))
}

func DGettext(domain, id string) string {
	_id := C.CString(id)
	_d := C.CString(domain)
	defer C.free(unsafe.Pointer(_id))
	defer C.free(unsafe.Pointer(_d))
	return C.GoString(C.dgettext(_d, _id))
}

func Textdomain(domain string) {
	_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(_domain))
	C.textdomain(_domain)
}

func InitI18n() {
	C._init_i18n()
}

func Bindtextdomain(domain, dirname string) string {
	_domain := C.CString(domain)
	_dirname := C.CString(dirname)
	defer C.free(unsafe.Pointer(_domain))
	defer C.free(unsafe.Pointer(_dirname))
	return C.GoString(C.bindtextdomain(_domain, _dirname))
}

func NTr(msgid, plural string, n int) string {
	cMsgid := C.CString(msgid)
	defer C.free(unsafe.Pointer(cMsgid))

	cPlural := C.CString(plural)
	defer C.free(unsafe.Pointer(cPlural))

	return C.GoString(C.ngettext(cMsgid, cPlural, C.ulong(n)))
}

func DNGettext(domain, msgid, plural string, n int) string {
	cDomain := C.CString(domain)
	defer C.free(unsafe.Pointer(cDomain))

	cMsgid := C.CString(msgid)
	defer C.free(unsafe.Pointer(cMsgid))

	cPlural := C.CString(plural)
	defer C.free(unsafe.Pointer(cPlural))

	return C.GoString(C.dngettext(cDomain, cMsgid, cPlural, C.ulong(n)))
}
