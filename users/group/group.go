// Copyright (c) 2015 Deepin Ltd. All rights reserved.
// Use of this source is govered by General Public License that can be found
// in the LICENSE file.
package group

// #include <stdlib.h>
// #include <grp.h>
import "C"

import (
	"fmt"
	"unsafe"
)

// Group wraps up `group` struct used in kernel
type Group struct {
	Name    string
	Passwd  string // Password is encrypted if password-shadowing is on
	Gid     uint32
	Members *[]string // member list
}

var char *C.char

// stringArrayC2Go converts string arrray in C to string slice in Golang
func stringArrayC2Go(strArrC **C.char) *[]string {
	strArr := make([]string, 5)
	for offset := uintptr(0); uintptr(*(*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(strArrC)) + offset))) != uintptr(0); offset += unsafe.Sizeof(uintptr(0)) {
		str := C.GoString(*(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(strArrC)) + offset)))
		strArr = append(strArr, str)
	}
	return &strArr
}

// stringArrayC2Go_2 converts string array in C to string slice in Golang
// It is a revised version of @stringArrayC2Go
func stringArrayC2Go_2(strArrC **C.char) *[]string {
	strArr := make([]string, 0)
	arrIndex := (*[1 << 30]*C.char)(unsafe.Pointer(strArrC))
	for _, p := range arrIndex {
		if p == nil {
			break
		}
		strArr = append(strArr, C.GoString(p))
	}
	return &strArr
}

// groupC2Go converts `group` struct from C to golang native struct
func groupC2Go(groupC *C.struct_group) *Group {
	members := stringArrayC2Go_2(groupC.gr_mem)
	return &Group{
		Name:    C.GoString(groupC.gr_name),
		Passwd:  C.GoString(groupC.gr_passwd),
		Gid:     uint32(groupC.gr_gid),
		Members: members,
	}
}

type GroupNotFoundError struct {
	Name string
	Gid  uint32
}

func (err *GroupNotFoundError) Error() string {
	if len(err.Name) > 0 {
		return fmt.Sprintf("Group with name `%s` not found!", err.Name)
	} else {
		return fmt.Sprintf("Group with gid `%s` not found!", err.Gid)
	}
}

// GetGroupByName wraps up `getgrnam` system call.
// It retrieves group records from group file based on group name.
func GetGroupByName(name string) (*Group, error) {
	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	groupC, err := C.getgrnam(nameC)
	if groupC == nil {
		if err == nil {
			return nil, &GroupNotFoundError{Name: name}
		} else {
			return nil, err
		}
	} else {
		return groupC2Go(groupC), nil
	}
}

// GetGroupByGid wraps up `getgrgid` system call.
// It retrieves group records from group file based on gid.
func GetGroupByGid(gid uint32) (*Group, error) {
	gidC := C.__gid_t(gid)
	groupC, err := C.getgrgid(gidC)
	if groupC == nil {
		if err == nil {
			return nil, &GroupNotFoundError{Gid: gid}
		} else {
			return nil, err
		}
	} else {
		return groupC2Go(groupC), nil
	}
}

// GetGroupEntry wraps up `getgrent` system call.
// It performs sequential scans of the records in the group file.
func GetGroupEntry() []*Group {
	groups := make([]*Group, 0)

	// Restart scanning from the begging of the group file.
	C.setgrent()

	for groupC, err := C.getgrent(); groupC != nil && err == nil; groupC, err = C.getgrent() {
		group := groupC2Go(groupC)
		groups = append(groups, group)
	}

	// Call endgrent() is necessary so that any subsequent getgrent() call will
	// reopen the group file and start from the beginning.
	C.endgrent()

	return groups
}
