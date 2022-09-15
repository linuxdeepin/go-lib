// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

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
