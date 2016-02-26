/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package archive

import (
	"fmt"
	"pkg.deepin.io/lib/archive/gzip"
	"strings"
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

	return nil
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

	return nil
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

	return nil, nil
}

func getCompressType(file string) int32 {
	switch {
	case strings.HasSuffix(file, "tar"):
		return CompressTypeTar
	case strings.HasSuffix(file, "tar.gz"):
		return CompressTypeTarGz
	case strings.HasSuffix(file, "tar.ba2"):
		return CompressTypeTarBz2
	case strings.HasSuffix(file, "zip"):
		return CompressTypeZip
	}

	return CompressTypeUnkown
}
