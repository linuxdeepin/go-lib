package introspect

import (
	"encoding/xml"
	"io"
)

const (
	ExtendFieldI18n    = "com.deepin.DBus.I18n"
	ExtendFieldNoReply = "org.freedesktop.DBus.Method.NoReply"
)

func Parse(reader io.Reader) (*NodeInfo, error) {
	decoder := xml.NewDecoder(reader)
	obj := &NodeInfo{}
	err := decoder.Decode(&obj)
	return obj, err
}

func extendFieldValue(annos []AnnotationInfo, field string) (string, bool) {
	for _, ano := range annos {
		if ano.Name == field {
			return ano.Value, true
		}
	}
	return "", false
}

func (arg ArgInfo) I18nField() bool {
	value, ok := extendFieldValue(arg.Annotations, ExtendFieldI18n)
	if !ok {
		return false
	}
	return boolValue(value)
}

func (prop PropertyInfo) I18nField() bool {
	value, ok := extendFieldValue(prop.Annotations, ExtendFieldI18n)
	if !ok {
		return false
	}
	return boolValue(value)
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
