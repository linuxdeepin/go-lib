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
