/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package gzip

import (
	"archive/tar"
	"compress/gzip"
	"os"
	"pkg.deepin.io/lib/archive/utils"
)

func tarCompressFiles(files []string, dest string) error {
	dw, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer dw.Close()

	gw := gzip.NewWriter(dw)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	return utils.TarWriterCompressFiles(tw, files)
}

func tarExtracte(src, dest string) ([]string, error) {
	sr, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer sr.Close()

	gr, err := gzip.NewReader(sr)
	if err != nil {
		return nil, err
	}
	defer gr.Close()

	return utils.TarReaderExtracte(tar.NewReader(gr), dest)
}
