/**
 * Copyright (c) 2011 ~ 2014 Deepin, Inc.
 *               2013 ~ 2014 jouyouyun
 *
 * Author:      jouyouyun <jouyouwen717@gmail.com>
 * Maintainer:  jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package gzip

import (
	"fmt"
)

const (
	ArchiveTypeTar int32 = iota + 1
	ArchiveTypeZip
	ArchiveTypeRar
)

func CompressDir(src, dest string, t int32) error {
	switch t {
	case ArchiveTypeTar:
		return tarCompressFiles([]string{src}, dest)
	case ArchiveTypeZip:
		return nil
	case ArchiveTypeRar:
		return nil
	default:
		return fmt.Errorf("Invalid archive type: %q", t)
	}

	return nil
}

func CompressFiles(files []string, dest string, t int32) error {
	switch t {
	case ArchiveTypeTar:
		return tarCompressFiles(files, dest)
	case ArchiveTypeZip:
		return nil
	case ArchiveTypeRar:
		return nil
	default:
		return fmt.Errorf("Invalid archive type: %q", t)
	}

	return nil
}

func Extracte(src, destDir string, t int32) ([]string, error) {
	switch t {
	case ArchiveTypeTar:
		return tarExtracte(src, destDir)
	case ArchiveTypeZip:
		return nil, nil
	case ArchiveTypeRar:
		return nil, nil
	default:
		return nil, fmt.Errorf("Invalid archive type: %q", t)
	}

	return nil, nil
}
