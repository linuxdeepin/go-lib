// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package locale

import (
	"bufio"
	"bytes"
	"os"
	"regexp"
	"strings"
	"sync"
)

const (
	charUscore = '_'
	charDot    = '.'
	charAt     = '@'
)

/*
 * Compute all interesting variants for a given locale name -
 * by stripping off different components of the value.
 *
 * For simplicity, we assume that the locale is in
 * X/Open format: language[_territory][.codeset][@modifier]
 */
func GetLocaleVariants(locale string) (variants []string) {
	cs := ExplodeLocale(locale)
	mask := cs.Mask
	for j := uint(0); j <= mask; j++ {
		i := mask - j
		if (i & ^mask) == 0 {
			var buf bytes.Buffer
			buf.WriteString(cs.Language)
			if i&ComponentTerritory != 0 {
				buf.WriteByte(charUscore)
				buf.WriteString(cs.Territory)
			}
			if i&ComponentCodeset != 0 {
				buf.WriteByte(charDot)
				buf.WriteString(cs.Codeset)
			}
			if i&ComponentModifier != 0 {
				buf.WriteByte(charAt)
				buf.WriteString(cs.Modifier)
			}
			variants = append(variants, buf.String())
		}
	}
	return
}

type Components struct {
	Language  string
	Territory string
	Codeset   string
	Modifier  string
	Mask      uint
}

const (
	ComponentCodeset = 1 << iota
	ComponentTerritory
	ComponentModifier
)

// ExplodeLocale Break an X/Open style locale specification into components
func ExplodeLocale(locale string) *Components {
	var c Components
	atIdx := strings.IndexByte(locale, charAt)
	if atIdx != -1 {
		c.Modifier = locale[atIdx+1:]
		locale = locale[:atIdx]
		c.Mask |= ComponentModifier
	}

	dotIdx := strings.IndexByte(locale, charDot)
	if dotIdx != -1 {
		c.Codeset = locale[dotIdx+1:]
		locale = locale[:dotIdx]
		c.Mask |= ComponentCodeset
	}

	uscoreIdx := strings.IndexByte(locale, charUscore)
	if uscoreIdx != -1 {
		c.Territory = locale[uscoreIdx+1:]
		locale = locale[:uscoreIdx]
		c.Mask |= ComponentTerritory
	}
	c.Language = locale
	return &c
}

func guessCategoryValue(categoryName string) (ret string) {
	// The highest priority value is the 'LANGUAGE' environment
	// variable.  This is a GNU extension.
	ret = os.Getenv("LANGUAGE")
	if ret != "" {
		return
	}

	// Setting of LC_ALL overwrites all other.
	ret = os.Getenv("LC_ALL")
	if ret != "" {
		return
	}

	// Next comes the name of the desired category.
	ret = os.Getenv(categoryName)
	if ret != "" {
		return
	}

	// Last possibility is the LANG environment variable.
	ret = os.Getenv("LANG")
	if ret != "" {
		return
	}
	return "C"
}

type _LanguageNamesCache struct {
	Language string
	Names    []string
	mutex    sync.Mutex
}

var languageNameCache _LanguageNamesCache

func GetLanguageNames() []string {
	value := guessCategoryValue("LC_MESSAGES")
	if value == "" {
		return []string{"C"}
	}

	languageNameCache.mutex.Lock()
	if languageNameCache.Language != value {
		languageNameCache.Language = value
		langs := strings.Split(value, ":")
		var names []string
		for _, lang := range langs {
			variants := GetLocaleVariants(unaliasLang(lang))
			names = append(names, variants...)
		}
		names = append(names, "C")
		languageNameCache.Names = names
	}
	languageNameCache.mutex.Unlock()

	return languageNameCache.Names
}

var splitor = regexp.MustCompile(`\s+|:`)

func readAliases(filename string) (aliasTable map[string]string) {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	aliasTable = make(map[string]string)
	scanner := bufio.NewScanner(bufio.NewReader(file))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		parts := splitor.Split(line, -1)
		if len(parts) == 2 {
			aliasTable[parts[0]] = parts[1]
		}
	}
	return
}

var aliases map[string]string
var aliasesOnce sync.Once
var aliasFile = "/usr/share/locale/locale.alias"

func unaliasLang(lang string) string {
	aliasesOnce.Do(func() {
		// init Aliases
		aliases = readAliases(aliasFile)
	})

	if lang1, ok := aliases[lang]; ok {
		return lang1
	}
	return lang
}
