/**
 * Copyright (C) 2017 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package utils

import (
	"syscall"
)

type FilesystemInfo struct {
	Type      int64
	TotalSize uint64 // byte
	FreeSize  uint64
	AvailSize uint64
	UsedSize  uint64
}

func QueryFilesytemInfo(path string) (*FilesystemInfo, error) {
	var buf syscall.Statfs_t
	err := syscall.Statfs(path, &buf)
	if err != nil {
		return nil, err
	}
	total := buf.Blocks * uint64(buf.Bsize)
	free := buf.Bfree * uint64(buf.Bsize)
	avail := buf.Bavail * uint64(buf.Bsize)
	return &FilesystemInfo{
		Type:      int64(buf.Type), // if in i386 it's int32
		TotalSize: total,
		FreeSize:  free,
		AvailSize: avail,
		UsedSize:  total - free,
	}, nil
}
