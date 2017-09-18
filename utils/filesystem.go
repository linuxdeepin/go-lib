/*
 * Copyright (C) 2017 ~ 2017 Deepin Technology Co., Ltd.
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

package utils
/*
#include <stdlib.h>
#include <sys/statvfs.h>
*/
import "C"
import "unsafe"
import "errors"

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
	if (C.statvfs(cpath, &buf) != 0){
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
