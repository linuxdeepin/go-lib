package gettext

/*
#include <locale.h>
#include <stdlib.h>
#include <libintl.h>
void _init_i18n() { setlocale(LC_ALL, ""); }
*/
import "C"
import "unsafe"
import "os"
import "strings"

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

// QueryLang return user lang.
// the rule is document at man gettext(3)
func QueryLang() string {
	return QueryLangs()[0]
}

// QueryLangs return array of user lang, split by ",".
// the rule is document at man gettext(3)
func QueryLangs() []string {
	LC_ALL := os.Getenv("LC_ALL")
	LC_MESSAGE := os.Getenv("LC_MESSAGE")
	LANGUAGE := os.Getenv("LANGUAGE")
	LANG := os.Getenv("LANG")

	if LC_ALL != "C" && LANGUAGE != "" {
		langs := strings.Split(LANGUAGE, ":")
		return langs
	}

	if LC_ALL != "" {
		return []string{LC_ALL}
	}
	if LC_MESSAGE != "" {
		return []string{LC_MESSAGE}
	}
	if LANG != "" {
		return []string{LANG}
	}
	return []string{""}
}
