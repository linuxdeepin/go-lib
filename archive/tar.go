// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package archive

import (
	"archive/tar"
	"os"
	"github.com/linuxdeepin/go-lib/archive/utils"
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
