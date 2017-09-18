/*
 * Copyright (C) 2014 ~ 2017 Deepin Technology Co., Ltd.
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

package introspect

import (
	"encoding/xml"
	"io"
)

const (
	ExtendFieldI18nDir    = "com.deepin.DBus.I18n.Dir"
	ExtendFieldI18nDomain = "com.deepin.DBus.I18n.Domain"
	ExtendFieldNoReply    = "org.freedesktop.DBus.Method.NoReply"
)

func Parse(reader io.Reader) (*NodeInfo, error) {
	decoder := xml.NewDecoder(reader)
	obj := &NodeInfo{}
	err := decoder.Decode(&obj)
	return obj, err
}

type AnnotationInfos []AnnotationInfo

func (annos AnnotationInfos) FieldValue(field string) (string, bool) {
	for _, ano := range annos {
		if ano.Name == field {
			return ano.Value, true
		}
	}
	return "", false
}
func (annos AnnotationInfos) I18nInfo() (string, string, bool) {
	dir, ok := annos.FieldValue(ExtendFieldI18nDir)
	domain, ok := annos.FieldValue(ExtendFieldI18nDomain)
	return dir, domain, ok
}

func extendFieldValue(annos []AnnotationInfo, field string) (string, bool) {
	for _, ano := range annos {
		if ano.Name == field {
			return ano.Value, true
		}
	}
	return "", false
}

func (arg ArgInfo) I18nInfo() (string, string, bool) {
	return AnnotationInfos(arg.Annotations).I18nInfo()
}

func (prop PropertyInfo) I18nInfo() (string, string, bool) {
	return AnnotationInfos(prop.Annotations).I18nInfo()
}

func (m MethodInfo) NoReply() bool {
	value, ok := extendFieldValue(m.Annotations, ExtendFieldNoReply)
	if !ok {
		return false
	}
	return boolValue(value)
}

func boolValue(v string) bool {
	if v == "true" {
		return true
	} else {
		return false
	}
}
