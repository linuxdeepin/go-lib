package main

import "fmt"
import "strings"
import "log"

func guessTypeQML(sig, con string) (r2 string, obj string) {
	return
}

func guessTypeGo(sig, con string) (r2 string, obj string) {
	switch sig[0] {
	case 'a':
		if sig[1] == '{' {
			r2 += "map["
			t2, t3 := guessTypeGo(sig[2:], con[2:])
			r2 += t2
			obj = t3

			r2 += "]"

			t2, obj = guessTypeGo(sig[3:], con[3:])
			r2 += t2
		} else if sig[1] == 'a' {
			r2 += "[]"
			t2, t3 := guessTypeGo(sig[1:], con[1:])
			r2 += t2
			obj = t3
		} else {
			r2 += "[]"
			t2, t3 := guessTypeGo(sig[1:], con[1:])
			r2 += t2
			obj = t3
		}
	case 's':
		r2 += "string"
	case 'o':
		i := strings.Index(con[1:], "|")
		obj = interfaceToObjectName(con[1 : i+1])
		r2 += "*" + obj
	}
	return
}

func tryConvertObjectPathQML(sig, con string) (r string) {
	return
}

func tryConvertObjectPathGo(sig, con string) (r string) {
	if !strings.Contains(sig, "o") {
		return ""
	}
	if strings.ContainsAny(sig, "()") || strings.Count(sig, "o") != 1 {
		log.Printf("Didn't support struct Object convert (%s)", sig)
		return ""
	}
	r2, obj := guessTypeGo(sig, con)
	n := strings.Count(sig, "a")
	r = fmt.Sprintf("after := %s{}\n", r2)
	for i := 0; i < n; i++ {
		index := ""
		for j := 0; j < i; j++ {
			index += fmt.Sprintf("[i%d]", j)
		}
		r += fmt.Sprintf("%sfor i%d, _ := range before%s {\n", strings.Repeat("\t", i), i, index)
		if i == n-1 {
			t := ""
			for j := 0; j < n; j++ {
				t += fmt.Sprintf("[i%d]", j)
			}
			r += fmt.Sprintf("%safter%s = Get%s(string(before%s))\n", strings.Repeat("\t", i+1), t, obj, t)
		}
	}
	for i := n - 1; i >= 0; i-- {
		r += strings.Repeat("\t", i) + "}\n"
	}
	return
}
