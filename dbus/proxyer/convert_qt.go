package main

import "strings"

var _sig2QType = map[byte]string{
	'y': "uchar",
	'b': "bool",
	'n': "short",
	'q': "ushort",
	'i': "int",
	'u': "uint",
	'x': "qlonglong",
	't': "qulonglong",
	'd': "double",
	's': "QString",
	'g': "QDBusOSignature",
	'o': "QDBusObjectPath",
	'v': "QDBusVariant",
}

var _convertQDBus = map[string]string{
	"o": "QVariant::fromValue(QDBusObjectPath({{.Name}}.value<QString>()))",
}

func normaliseQDBus(v string) (r string) {
	return //TODO:
	if result, ok := _convertQDBus[v]; ok {
		r = result
		/*return "huhu" //result*/
	}
	return
}

func getQType(sig string) string {
	if qtype, ok := _sig2QType[sig[0]]; ok {
		return qtype
	}
	switch sig[0] {
	case 'a':
		if sig[1] == '{' {
			i := strings.LastIndex(sig, "}")
			r := "QMap<"
			r += getQType(string(sig[2])) + ", "
			r += getQType(sig[3:i])
			return r + " >"
		} else {
			r := "QList<"
			r += getQType(sig[1:])
			return r + " >"
		}
	case '(':
		return "QVariant"
	}
	panic("Unknow Type" + sig)
}
