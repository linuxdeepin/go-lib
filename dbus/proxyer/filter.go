package main

import "dlib/dbus"

func _filter(old *string) {
	v := map[string]bool{
		"break":       true,
		"default":     true,
		"func":        true,
		"interface":   true,
		"select":      true,
		"case":        true,
		"defer":       true,
		"go":          true,
		"map":         true,
		"struct":      true,
		"chan":        true,
		"else":        true,
		"goto":        true,
		"package":     true,
		"switch":      true,
		"const":       true,
		"fallthrough": true,
		"if":          true,
		"range":       true,
		"type":        true,
		"continue":    true,
		"for":         true,
		"import":      true,
		"return":      true,
		"var":         true,
	}
	if v[*old] {
		*old = *old + "_"
	}
}

func filterGoKeyWord(info *dbus.InterfaceInfo) {
	_filter(&info.Name)
	for i, _ := range info.Methods {
		_filter(&info.Methods[i].Name)
		for j, _ := range info.Methods[i].Args {
			_filter(&info.Methods[i].Args[j].Name)
		}
	}
	for i, _ := range info.Signals {
		_filter(&info.Signals[i].Name)
		for j, _ := range info.Signals[i].Args {
			_filter(&info.Signals[i].Args[j].Name)
		}
	}
	for i, _ := range info.Properties {
		_filter(&info.Properties[i].Name)
	}
}
