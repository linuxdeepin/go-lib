package operations

// Port from glib(gkeyfile.c and gcharset.c).
// because we cannot pass NULL to glib.KeyFile.GetLocaleString like C, so the
// locale must be passed expclitly.

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

const (
	_ComponentCodeset = 1 << iota
	_ComponentModifier
	_ComponentTerritory
)

var (
	aliasTable map[string]string
	saidBefore = false
	splitor    = regexp.MustCompile(`\s+|:`)
)

func readAliases(filename string) {
	if aliasTable == nil {
		aliasTable = make(map[string]string, 0)
	}

	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		lineContent := strings.TrimSpace(s.Text())
		if len(lineContent) == 0 || lineContent == "#" {
			continue
		}

		content := splitor.Split(lineContent, -1)
		if len(content) != 2 {
			continue
		}
		aliasTable[content[0]] = aliasTable[content[1]]
	}
}

func unaliasLang(lang string) string {
	if aliasTable == nil {
		readAliases("/usr/share/locale/locale.alias")
	}

	i := 0
	for p, ok := aliasTable[lang]; ok && p == lang; i++ {
		lang = p
		if i == 30 {
			if !saidBefore {
				// Warning("Too many alias levels for a locale, may indicate a loop")
			}
			saidBefore = true
			return lang
		}
	}

	return lang
}

func explodeLocale(locale string) (mask uint, language string, territory string, codeset string, modifier string) {
	mask = uint(0)
	uscorePos := strings.IndexRune(locale, '_')
	dotPos := strings.IndexRune(locale, '.')
	atPos := strings.IndexRune(locale, '@')

	if atPos != -1 {
		mask |= _ComponentModifier
		modifier = locale[atPos:]
	} else {
		atPos = len(locale)
	}

	if dotPos != -1 {
		mask |= _ComponentCodeset
		codeset = locale[dotPos:atPos]
	} else {
		dotPos = atPos
	}

	if uscorePos != -1 {
		mask |= _ComponentTerritory
		territory = locale[uscorePos:dotPos]
	} else {
		uscorePos = dotPos
	}

	language = locale[:uscorePos]
	return
}

// GetLocaleVariants returns locale variants
func GetLocaleVariants(locale string) []string {
	var array []string
	mask, language, territory, codeset, modifier := explodeLocale(locale)

	for j := uint(0); j <= mask; j++ {
		i := mask - j
		if (i & ^mask) == 0 {
			val := language
			if (i & _ComponentTerritory) != 0 {
				val = val + territory
			}

			if (i & _ComponentCodeset) != 0 {
				val = val + codeset
			}

			if (i & _ComponentModifier) != 0 {
				val = val + modifier
			}

			array = append(array, val)
		}
	}

	return array
}

func guessCategoryValue(categoryName string) (retval string) {
	retval = os.Getenv("LANGUAGE")
	if retval != "" {
		return
	}

	retval = os.Getenv("LC_ALL")
	if retval != "" {
		return
	}

	retval = os.Getenv(categoryName)
	if retval != "" {
		return
	}

	retval = os.Getenv("LANG")
	if retval != "" {
		return
	}

	retval = "C"
	return
}

// GetLanguageNames returns all language names.
func GetLanguageNames() []string {
	val := guessCategoryValue("LC_MESSAGES")
	langs := strings.Split(val, ":")
	var array []string
	for _, lang := range langs {
		array = append(array, GetLocaleVariants(unaliasLang(lang))...)
	}

	array = append(array, "C")

	return array
}
