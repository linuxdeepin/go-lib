package main

import "dlib/dbus"
import "strings"
import "strconv"
import "log"

func getGoKeyword() map[string]bool {
	return map[string]bool{
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
		"uint8":       true,
		"uint16":      true,
		"uint32":      true,
		"uint64":      true,
		"int8":        true,
		"int16":       true,
		"int32":       true,
		"int64":       true,
		"float32":     true,
		"float64":     true,
		"complex64":   true,
		"complex128":  true,
		"byte":        true,
		"rune":        true,
		"uint":        true,
		"int":         true,
		"uintptr":     true,
		"string":      true,
		"bool":        true,
	}
}

func getPyQtKeyword() map[string]bool {
	return map[string]bool{}
}

func keywordFilter(v map[string]bool, old *string) (ret map[string]bool, hasHit bool) {
	*old = strings.Replace(*old, "-", "_", -1)
	if v[*old] {
		*old = *old + "_"
		hasHit = true
	}
	v[*old] = true
	ret = v
	return
}

func filterKeyWord(keyword func() map[string]bool, info *dbus.InterfaceInfo) {
	var hit bool
	keywordFilter(keyword(), &info.Name)

	methodKeyword := keyword()
	for i, _ := range info.Methods {
		method_name := &info.Methods[i].Name
		if methodKeyword, hit = keywordFilter(methodKeyword, method_name); hit {
			log.Printf("Method name(%s.%s) conflict: convert", info.Name, *method_name)
		}

		argKeyword := keyword()
		for j := 0; j < len(info.Methods[i].Args); j++ {
			name := &info.Methods[i].Args[j].Name
			if len(*name) == 0 {
				*name = "arg" + strconv.Itoa(j)
			}
			if argKeyword, hit = keywordFilter(argKeyword, name); hit {
				log.Printf("The %d arg of (%s.%s:%s) conflict", j, info.Name, *method_name, *name)
			}
		}
	}

	sigKeyword := keyword()
	for i, _ := range info.Signals {
		sig_name := &info.Signals[i].Name

		if sigKeyword, hit = keywordFilter(sigKeyword, sig_name); hit {
			log.Printf("Signal name(%s.%s) conflict", info.Name, *sig_name)
		}

		argKeyword := keyword()
		for j, _ := range info.Signals[i].Args {
			name := &info.Signals[i].Args[j].Name
			if len(*name) == 0 {
				*name = "arg" + strconv.Itoa(j)
			}
			if argKeyword, hit = keywordFilter(argKeyword, name); hit {
				log.Printf("The %d arg of (%s.%s:%s) conflict", j, info.Name, *sig_name, *name)
			}
		}
	}

	propKeyword := keyword()
	for i, _ := range info.Properties {
		prop_name := &info.Properties[i].Name
		if propKeyword, hit = keywordFilter(propKeyword, prop_name); hit {
			log.Printf("Property name(%s.%s) conflict: convert", info.Name, *prop_name)
		}
	}

	func(info *dbus.InterfaceInfo) {
		usedName := make(map[string]bool)
		for _, m := range info.Methods {
			usedName[m.Name] = true
		}
		for i, s := range info.Signals {
			sigName := "Connect" + s.Name
			if usedName[sigName] {
				newName := sigName + "_"
				info.Signals[i].Name = newName
				usedName[newName] = true
			}
		}
		for i, p := range info.Properties {
			if p.Access == "readwrite" {
				propName := "Set" + p.Name
				if usedName[propName] {
					newName := propName + "_"
					info.Properties[i].Name = newName
					usedName[newName] = true
				}
			}
			propName := "Get" + p.Name
			if usedName[propName] {
				newName := propName + "_"
				info.Properties[i].Name = newName
				usedName[newName] = true
			}
		}
	}(info) //solve name conflict
}
