/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

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
