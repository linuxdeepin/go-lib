// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package utils

/*
#include <stdlib.h>
#include <sys/statvfs.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

type FilesystemInfo struct {
	TotalSize uint64 // byte
	FreeSize  uint64
	AvailSize uint64
	UsedSize  uint64
}

func QueryFilesytemInfo(path string) (*FilesystemInfo, error) {
	buf := C.struct_statvfs{}
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	if C.statvfs(cpath, &buf) != 0 {
		return nil, errors.New("Statvfs error.")
	}

	total := uint64(buf.f_blocks) * uint64(buf.f_frsize)
	free := uint64(buf.f_bfree) * uint64(buf.f_frsize)
	//Get real avail size instead of free size.
	avail := uint64(buf.f_bavail) * uint64(buf.f_frsize)
	return &FilesystemInfo{
		TotalSize: total,
		FreeSize:  free,
		AvailSize: avail,
		UsedSize:  total - free,
	}, nil
}
