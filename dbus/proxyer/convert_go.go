package main

import (
	"strings"
)

func normalizeMethodName(name string) string {
	words := strings.Split(name, "_")
	normalized := ""
	for _, w := range words {
		normalized += upper(w)
	}
	return upper(normalized)
}
