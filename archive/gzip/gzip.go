// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

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
}
