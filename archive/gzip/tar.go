// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package gzip

import (
	"archive/tar"
	"compress/gzip"
	"os"
	"github.com/linuxdeepin/go-lib/archive/utils"
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
