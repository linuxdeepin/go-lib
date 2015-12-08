/**
 * Copyright (c) 2011 ~ 2015 Deepin, Inc.
 *               2013 ~ 2015 jouyouyun
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

package utils

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path"
	dutils "pkg.deepin.io/lib/utils"
)

func TarWriterCompressFiles(writer *tar.Writer, files []string) error {
	if writer == nil {
		return fmt.Errorf("Invalid tar writer")
	}

	for _, file := range files {
		var err error
		if dutils.IsDir(file) {
			err = tarWriterCompressDir(writer, file, "")
		} else {
			err = tarWriterCompressFile(writer, file, "")
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func TarReaderExtracte(reader *tar.Reader, dest string) ([]string, error) {
	if reader == nil {
		return nil, fmt.Errorf("Invalid tar reader")
	}

	var files []string
	for {
		h, err := reader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		tmp := path.Join(dest, h.Name)
		if h.FileInfo().IsDir() {
			err = os.MkdirAll(tmp, 0755)
			if err != nil {
				return nil, err
			}
			continue
		}
		err = os.MkdirAll(path.Dir(tmp), 0755)
		if err != nil {
			return nil, err
		}

		fw, err := os.OpenFile(tmp, os.O_CREATE|os.O_WRONLY,
			os.FileMode(h.Mode))
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(fw, reader)
		fw.Close()
		if err != nil {
			return nil, err
		}

		files = append(files, tmp)
	}

	return files, nil
}

func tarWriterCompressDir(tw *tar.Writer, dir, parent string) error {
	dr, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer dr.Close()

	finfos, err := dr.Readdir(0)
	if err != nil {
		return err
	}

	parent = path.Join(parent, path.Base(dir))
	for _, finfo := range finfos {
		file := path.Join(dir, finfo.Name())

		var err error
		if finfo.IsDir() {
			err = tarWriterCompressDir(tw, file, parent)
		} else {
			err = tarWriterCompressFile(tw, file, parent)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func tarWriterCompressFile(tw *tar.Writer, file, parent string) error {
	fr, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fr.Close()

	finfo, err := fr.Stat()
	if err != nil {
		return err
	}

	h := new(tar.Header)
	h.Name = path.Join(parent, finfo.Name())
	h.Size = finfo.Size()
	h.Mode = 0644
	//h.ModTime = finfo.ModTime()

	err = tw.WriteHeader(h)
	if err != nil {
		return err
	}

	_, err = io.Copy(tw, fr)
	if err != nil {
		return err
	}

	return nil
}
