package main

import "dlib/dbus"
import "strings"

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
	}
}

func keywordFilter(v map[string]bool, old *string) map[string]bool {
	*old = strings.Replace(*old, "-", "_", -1)
	if v[*old] {
		*old = *old + "_"
		v[*old] = true
	}
	return v
}

func filterGoKeyWord(info *dbus.InterfaceInfo) {
	keywordFilter(getGoKeyword(), &info.Name)

	methodKeyword := getGoKeyword()
	for i, _ := range info.Methods {
		methodKeyword = keywordFilter(methodKeyword, &info.Methods[i].Name)

		argKeyword := getGoKeyword()
		/*for j, _ := range info.Methods[i].Args {*/
		/*argKeyword = keywordFilter(argKeyword, &info.Methods[i].Args[j].Name)*/
		/*}*/

		for j := 0; j < len(info.Methods[i].Args); j++ {
			argKeyword = keywordFilter(argKeyword, &info.Methods[i].Args[j].Name)
		}
	}

	sigKeyword := getGoKeyword()
	for i, _ := range info.Signals {
		sigKeyword = keywordFilter(sigKeyword, &info.Signals[i].Name)

		argKeyword := getGoKeyword()
		for j, _ := range info.Signals[i].Args {
			argKeyword = keywordFilter(argKeyword, &info.Signals[i].Args[j].Name)
		}
	}

	propKeyword := getGoKeyword()
	for i, _ := range info.Properties {
		propKeyword = keywordFilter(propKeyword, &info.Properties[i].Name)
	}
	solveNameConfilict(info)
}

func solveNameConfilict(info *dbus.InterfaceInfo) {
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
}
