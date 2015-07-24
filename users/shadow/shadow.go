// Copyright (c) 2015 Deepin Ltd. All rights reserved.
// Use of this source is govered by General Public License that can be found
// in the LICENSE file.
package shadow

// #include <shadow.h>
import "C"

import (
	"fmt"
)

// spwd wraps up `spwd` (shadow password) used in kernel.
type Shadow struct {
	Name       string // login username
	Password   string // encrypted password
	LastChange int64  // Date of last change since 1970-01-01
	MinDays    int64  // Min # of days between changes
	MaxDays    int64  // Max # of days between changes
	Warn       int64  // # of dayd before password expires to warn user to change it
	Inactive   int64  // # of days after password expires until account is disabled
	Expire     int64  // Date when account expires since 1970-01-01
	Flag       uint64 // reserved
}

// shadowC2Go converts `spwd` struct to golang native struct
func shadowC2Go(shadowC *C.struct_spwd) *Shadow {
	return &Shadow{
		Name:       C.GoString(shadowC.sp_namp),
		Password:   C.GoString(shadowC.sp_pwdp),
		LastChange: int64(shadowC.sp_lstchg),
		MinDays:    int64(shadowC.sp_min),
		MaxDays:    int64(shadowC.sp_max),
		Warn:       int64(shadowC.sp_warn),
		Inactive:   int64(shadowC.sp_inact),
		Expire:     int64(shadowC.sp_expire),
		Flag:       uint64(shadowC.sp_flag),
	}
}

type UserNotFoundError struct {
	Name string
	Uid  uint32
}

func (err *UserNotFoundError) Error() string {
	if len(err.Name) > 0 {
		return fmt.Sprintf("User with name `%s` not found!", err.Name)
	} else {
		return fmt.Sprintf("User with uid `%d` not found!", err.Uid)
	}
}

// GetShadowByName wraps up `getspnam` system call.
// It returns the record in the shadow database file that matches the username
// @name
func GetShadowByName(name string) (*Shadow, error) {
	nameC := C.CString(name)
	shadowC, err := C.getspnam(nameC)
	if shadowC == nil {
		if err == nil {
			return nil, &UserNotFoundError{Name: name}
		} else {
			return nil, err
		}
	} else {
		return shadowC2Go(shadowC), nil
	}
}

// GetShadowEntry wraups up `getspent` system call.
// It performs sequential scans of the records in the shadow file.
func GetShadowEntry() []*Shadow {
	shadows := make([]*Shadow, 0)

	// Restart scanning from the begging of the shadow file.
	C.setspent()

	for shadowC, err := C.getspent(); shadowC != nil && err == nil; shadowC, err = C.getspent() {
		shadow := shadowC2Go(shadowC)
		shadows = append(shadows, shadow)
	}

	// Call endspent() is necessary so that any subsequent getspent() call will
	// reopen the shadow file and start from the beginning.
	C.endspent()

	return shadows
}
