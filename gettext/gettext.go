// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package gettext

/*
#include <locale.h>
#include <stdlib.h>
#include <libintl.h>
void _init_i18n() { setlocale(LC_ALL, ""); }
*/
import "C"

import (
	"os"
	"strings"
	"unsafe"
)

var (
	// LcAll is for all of the locale.
	LcAll = int(C.LC_ALL)

	// LcCollate is for regular expression matching (it determines the meaning of
	// range expressions and equivalence classes) and string collation.
	LcCollate = int(C.LC_COLLATE)

	// LcCtype is for regular expression matching, character classification,
	// conversion, case-sensitive comparison, and wide character functions.
	LcCtype = int(C.LC_CTYPE)

	// LcMessages is for localizable natural-language messages.
	LcMessages = int(C.LC_MESSAGES)

	// LcMonetary is for monetary formatting.
	LcMonetary = int(C.LC_MONETARY)

	// LcNumeric is for number formatting (such as the decimal point and the
	// thousands separator).
	LcNumeric = int(C.LC_NUMERIC)

	// LcTime is for time and date formatting.
	LcTime = int(C.LC_TIME)
)

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

func SetLocale(category int, locale string) string {
	cLocale := C.CString(locale)
	defer C.free(unsafe.Pointer(cLocale))
	return C.GoString(C.setlocale(C.int(category), cLocale))
}

func Bindtextdomain(domain, dirname string) string {
	_domain := C.CString(domain)
	_dirname := C.CString(dirname)
	defer C.free(unsafe.Pointer(_domain))
	defer C.free(unsafe.Pointer(_dirname))
	return C.GoString(C.bindtextdomain(_domain, _dirname))
}

func BindTextdomainCodeset(domain, codeset string) string {
	_domain := C.CString(domain)
	_codeset := C.CString(codeset)
	defer C.free(unsafe.Pointer(_domain))
	defer C.free(unsafe.Pointer(_codeset))
	return C.GoString(C.bind_textdomain_codeset(_domain, _codeset))
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

// QueryLangs return array of user lang, split by ":".
// the rule is document at man gettext(3)
func QueryLangs() []string {
	LC_ALL := os.Getenv("LC_ALL")
	LC_MESSAGE := os.Getenv("LC_MESSAGE")
	LANGUAGE := os.Getenv("LANGUAGE")
	LANG := os.Getenv("LANG")

	cutoff := func(s string) string {
		for i, c := range s {
			if c == '.' {
				return s[:i]
			}
		}
		return s
	}

	if LC_ALL != "C" && LANGUAGE != "" {
		var r []string
		for _, lang := range strings.Split(LANGUAGE, ":") {
			r = append(r, cutoff(lang))
		}
		return r
	}

	if LC_ALL != "" {
		return []string{cutoff(LC_ALL)}
	}
	if LC_MESSAGE != "" {
		return []string{cutoff(LC_MESSAGE)}
	}
	if LANG != "" {
		return []string{cutoff(LANG)}
	}
	return []string{""}
}
