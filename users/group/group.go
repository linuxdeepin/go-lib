// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

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

// stringArrayC2Go_2 converts string array in C to string slice in Golang
// It is a revised version of @stringArrayC2Go
func stringArrayC2Go_2(strArrC **C.char) *[]string {
	strArr := make([]string, 0)
	arrIndex := (*[1 << 10]*C.char)(unsafe.Pointer(strArrC))
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
		return fmt.Sprintf("Group with gid `%d` not found!", err.Gid)
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
	var groups []*Group

	// Restart scanning from the begging of the group file.
	C.setgrent()

	for {
		groupC := C.getgrent()
		if groupC == nil {
			break
		}

		group := groupC2Go(groupC)
		groups = append(groups, group)
	}

	// Call endgrent() is necessary so that any subsequent getgrent() call will
	// reopen the group file and start from the beginning.
	C.endgrent()

	return groups
}
