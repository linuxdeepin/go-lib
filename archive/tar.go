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
	"archive/tar"
	"os"
	"pkg.deepin.io/lib/archive/utils"
)

func tarCompressFiles(files []string, dest string) error {
	dw, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer dw.Close()

	tw := tar.NewWriter(dw)
	defer tw.Close()

	return utils.TarWriterCompressFiles(tw, files)
}

func tarExtracteFile(src, dest string) ([]string, error) {
	sr, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer sr.Close()

	return utils.TarReaderExtracte(tar.NewReader(sr), dest)
}
