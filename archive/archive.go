// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package archive

import (
	"fmt"
	"strings"

	"github.com/linuxdeepin/go-lib/archive/gzip"
)

const (
	CompressTypeUnkown int32 = iota
	CompressTypeTar
	CompressTypeTarGz
	CompressTypeTarBz2
	CompressTypeZip
)

func CompressDir(src, dest string) error {
	switch getCompressType(dest) {
	case CompressTypeTar:
		return tarCompressFiles([]string{src}, dest)
	case CompressTypeTarGz:
		return gzip.CompressDir(src, dest, gzip.ArchiveTypeTar)
	case CompressTypeTarBz2:
		return nil
	case CompressTypeZip:
		return nil
	default:
		return fmt.Errorf("Invalid archive compress type")
	}
}

func CompressFiles(files []string, dest string) error {
	switch getCompressType(dest) {
	case CompressTypeTar:
		return tarCompressFiles(files, dest)
	case CompressTypeTarGz:
		return gzip.CompressFiles(files, dest, gzip.ArchiveTypeTar)
	case CompressTypeTarBz2:
		return nil
	case CompressTypeZip:
		return nil
	default:
		return fmt.Errorf("Invalid archive compress type")
	}
}

func Extracte(src, dest string) ([]string, error) {
	switch getCompressType(src) {
	case CompressTypeTar:
		return tarExtracteFile(src, dest)
	case CompressTypeTarGz:
		return gzip.Extracte(src, dest, gzip.ArchiveTypeTar)
	case CompressTypeTarBz2:
		return nil, nil
	case CompressTypeZip:
		return nil, nil
	default:
		return nil, fmt.Errorf("Invalid archive compress type")
	}
}

func getCompressType(file string) int32 {
	switch {
	case strings.HasSuffix(file, "tar"):
		return CompressTypeTar
	case strings.HasSuffix(file, "tar.gz"):
		return CompressTypeTarGz
	case strings.HasSuffix(file, "tar.bz2"):
		return CompressTypeTarBz2
	case strings.HasSuffix(file, "zip"):
		return CompressTypeZip
	}

	return CompressTypeUnkown
}
