// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package libc

/*
#define _POSIX_SOURCE
#include <stdlib.h>
#include <time.h>
*/
import "C"

import (
	"time"
	"unsafe"
)

type Tm struct {
	core C.struct_tm
}

func (tm Tm) native() *C.struct_tm {
	return &tm.core
}

func NewTm(t time.Time) Tm {
	var tm C.struct_tm
	tm.tm_sec = C.int(t.Second())
	tm.tm_min = C.int(t.Minute())
	tm.tm_hour = C.int(t.Hour())
	tm.tm_mday = C.int(t.Day())
	tm.tm_mon = C.int(t.Month() - 1)
	tm.tm_year = C.int(t.Year() - 1900)
	tm.tm_wday = C.int(t.Weekday())
	tm.tm_yday = C.int(t.YearDay() - 1)
	tm.tm_isdst = -1
	return Tm{
		core: tm,
	}
}

const initialBufSize = 256

func Strftime(format string, tm Tm) (s string) {
	if format == "" {
		return
	}

	fmt := C.CString(format)
	defer C.free(unsafe.Pointer(fmt))

	for size := initialBufSize; ; size *= 2 {
		buf := (*C.char)(C.malloc(C.size_t(size))) // can panic
		defer C.free(unsafe.Pointer(buf))
		n := C.strftime(buf, C.size_t(size), fmt, tm.native())
		if n == 0 {
			// strftime(3), unhelpfully: "Note that the return value 0 does not
			// necessarily indicate an error; for example, in many locales %p
			// yields an empty string." This leaves no definite way to
			// distinguish between the cases where the value doesn't fit and
			// where it does because the string is empty. In the latter case,
			// allocating increasingly larger buffers will never change the
			// result, so we need some heuristic for bailing out.
			//
			// Since a single 2-byte conversion sequence should not produce an
			// output longer than about 24 bytes, we conservatively allow the
			// buffer size to grow up to 20 times larger than the format string
			// before giving up.
			if size > 20*len(format) {
				return
			}
		} else if int(n) < size {
			s = C.GoStringN(buf, C.int(n))
			return
		}
	}
}
