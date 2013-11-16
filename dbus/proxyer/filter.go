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

func keywordFilter(v map[string]bool, old *string) map[string]bool {
	*old = strings.Replace(*old, "-", "_", -1)
	if v[*old] {
		log.Println("Name conflict:", *old, " convent to:", *old+"_")
		*old = *old + "_"
	}
	v[*old] = true
	return v
}

func filterKeyWord(keyword func() map[string]bool, info *dbus.InterfaceInfo) {
	keywordFilter(keyword(), &info.Name)

	methodKeyword := keyword()
	for i, _ := range info.Methods {
		methodKeyword = keywordFilter(methodKeyword, &info.Methods[i].Name)

		argKeyword := keyword()
		for j := 0; j < len(info.Methods[i].Args); j++ {
			name := &info.Methods[i].Args[j].Name
			if len(*name) == 0 {
				*name = "arg" + strconv.Itoa(j)
			}
			argKeyword = keywordFilter(argKeyword, name)
		}
	}

	sigKeyword := keyword()
	for i, _ := range info.Signals {
		sigKeyword = keywordFilter(sigKeyword, &info.Signals[i].Name)

		argKeyword := keyword()
		for j, _ := range info.Signals[i].Args {
			argKeyword = keywordFilter(argKeyword, &info.Signals[i].Args[j].Name)
		}
	}

	propKeyword := keyword()
	for i, _ := range info.Properties {
		propKeyword = keywordFilter(propKeyword, &info.Properties[i].Name)
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
			propName := "Set" + p.Name
			if usedName[propName] {
				newName := propName + "_"
				info.Properties[i].Name = newName
				usedName[newName] = true
			}
			propName = "Get" + p.Name
			if usedName[propName] {
				newName := propName + "_"
				info.Properties[i].Name = newName
				usedName[newName] = true
			}
		}
	}(info) //solve name conflict
}
