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
