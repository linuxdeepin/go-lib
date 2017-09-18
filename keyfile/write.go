/*
 * Copyright (C) 2017 ~ 2017 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package keyfile

import (
	"bytes"
	"io"
	"os"
)

func (f *KeyFile) SaveToWriter(w io.Writer) error {
	const equalSign = "="
	var buf bytes.Buffer
	for _, section := range f.sectionList {
		// write section comments
		sectionComments := f.GetSectionComments(section)
		if len(sectionComments) > 0 {
			if _, err := buf.WriteString(sectionComments); err != nil {
				return err
			}
			if _, err := buf.WriteString(LineBreak); err != nil {
				return err
			}
		}

		// write section name
		if err := buf.WriteByte('['); err != nil {
			return err
		}
		if _, err := buf.WriteString(section); err != nil {
			return err
		}
		if err := buf.WriteByte(']'); err != nil {
			return err
		}
		if _, err := buf.WriteString(LineBreak); err != nil {
			return err
		}

		// write keys
		for _, key := range f.keyList[section] {
			// ignore empty key
			if key == "" {
				continue
			}

			// write key comments
			keyComments := f.GetKeyComments(section, key)
			if len(keyComments) > 0 {
				if _, err := buf.WriteString(keyComments); err != nil {
					return err
				}
				if _, err := buf.WriteString(LineBreak); err != nil {
					return err
				}
			}

			// write key and value
			if _, err := buf.WriteString(key); err != nil {
				return err
			}
			if _, err := buf.WriteString(equalSign); err != nil {
				return err
			}
			value := f.data[section][key]
			if _, err := buf.WriteString(value); err != nil {
				return err
			}
			if _, err := buf.WriteString(LineBreak); err != nil {
				return err
			}

		}

		// Put a line between sections.
		if _, err := buf.WriteString(LineBreak); err != nil {
			return err
		}

	}

	_, err := buf.WriteTo(w)
	return err
}

func (kf *KeyFile) SaveToFile(file string) error {
	fh, err := os.Create(file)
	if err != nil {
		return err
	}

	defer fh.Close()
	return kf.SaveToWriter(fh)
}
