// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package keyfile

import (
	"fmt"
	"regexp"
	"sync"
)

var LineBreak = "\n"

type KeyFile struct {
	mutex  sync.RWMutex
	keyReg *regexp.Regexp
	data   map[string]map[string]string // section -> key : value

	sectionList []string            // section name list
	keyList     map[string][]string // section -> key name list

	sectionComments map[string]string
	keyComments     map[string]map[string]string // keys comments.
	ListSeparator   byte
}

func NewKeyFile() *KeyFile {
	f := &KeyFile{
		data:            make(map[string]map[string]string),
		keyList:         make(map[string][]string),
		sectionComments: make(map[string]string),
		keyComments:     make(map[string]map[string]string),
		ListSeparator:   ';',
	}
	return f
}

func (f *KeyFile) SetKeyRegexp(keyReg *regexp.Regexp) {
	f.keyReg = keyReg
}

// it returns true if the key and value were inserted
// or return false if the value was overwritten.
func (f *KeyFile) SetValue(section, key, value string) bool {
	if section == "" || key == "" {
		return false
	}

	if f.keyReg != nil &&
		!f.keyReg.MatchString(key) {
		return false
	}

	f.mutex.Lock()
	defer f.mutex.Unlock()

	// Check if section exists
	if _, ok := f.data[section]; !ok {
		f.data[section] = make(map[string]string)
		f.sectionList = append(f.sectionList, section)
	}

	// check if key exists
	_, ok := f.data[section][key]
	f.data[section][key] = value
	if !ok {
		f.keyList[section] = append(f.keyList[section], key)
	}
	return !ok
}

func (f *KeyFile) DeleteKey(section, key string) bool {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	// Check if section exists
	if _, ok := f.data[section]; !ok {
		return false
	}

	// Check if key exists
	if _, ok := f.data[section][key]; ok {
		delete(f.data[section], key)
		f.SetKeyComments(section, key, "")
		i := 0
		for _, keyName := range f.keyList[section] {
			if keyName == key {
				break
			}
			i++
		}
		// Remove from key list
		keyList := f.keyList[section]
		f.keyList[section] = append(keyList[:i], keyList[i+1:]...)
		return true
	}
	return false
}

func (f *KeyFile) GetValue(section, key string) (string, error) {
	f.mutex.RLock()
	defer f.mutex.RUnlock()

	// Check if section exists
	if _, ok := f.data[section]; !ok {
		return "", SectionNotFoundError{section}
	}

	// Section exists
	// Check if key exists or empty value
	value, ok := f.data[section][key]
	if !ok {
		return "", KeyNotFoundError{key}
	}

	// Key exists
	return value, nil
}

func (f *KeyFile) GetBool(section, key string) (bool, error) {
	value, err := f.GetValue(section, key)
	if err != nil {
		return false, err
	}

	return parseValueAsBool(value)
}

func (f *KeyFile) GetSections() []string {
	return f.sectionList
}

func (f *KeyFile) GetKeys(section string) []string {
	if _, ok := f.data[section]; !ok {
		// Section not exists
		return nil
	}
	return f.keyList[section]
}

// It returns true if the section was deleted, and false if the section didn't exist
func (f *KeyFile) DeleteSection(section string) bool {
	if _, ok := f.data[section]; !ok {
		return false
	}

	delete(f.data, section)
	f.SetSectionComments(section, "")

	// Remove from sectionList
	i := 0
	for _, secName := range f.sectionList {
		if secName == section {
			break
		}
		i++
	}
	f.sectionList = append(f.sectionList[:i], f.sectionList[i+1:]...)
	delete(f.keyList, section)
	delete(f.keyComments, section)
	return true
}

func (f *KeyFile) GetSection(section string) (map[string]string, error) {
	if _, ok := f.data[section]; !ok {
		return nil, SectionNotFoundError{section}
	}

	secMap := f.data[section]
	return secMap, nil
}

func (f *KeyFile) SetSectionComments(section, comments string) bool {
	if len(comments) == 0 {
		delete(f.sectionComments, section)
		return true
	}

	// Check if comment exists
	_, ok := f.sectionComments[section]
	if comments[0] != '#' {
		comments = "# " + comments
	}
	f.sectionComments[section] = comments
	return !ok
}

func (f *KeyFile) SetKeyComments(section, key, comments string) bool {
	if _, ok := f.keyComments[section]; ok {
		if len(comments) == 0 {
			delete(f.keyComments[section], key)
			return true
		}
	} else {
		if len(comments) == 0 {
			return true
		} else {
			f.keyComments[section] = make(map[string]string)
		}
	}

	_, ok := f.keyComments[section][key]
	if comments[0] != '#' {
		comments = "# " + comments
	}
	f.keyComments[section][key] = comments
	return !ok
}

func (f *KeyFile) GetSectionComments(section string) string {
	return f.sectionComments[section]
}

func (f *KeyFile) GetKeyComments(section, key string) string {
	if _, ok := f.keyComments[section]; ok {
		return f.keyComments[section][key]
	}
	return ""
}

type SectionNotFoundError struct {
	Name string
}

func (err SectionNotFoundError) Error() string {
	return fmt.Sprintf("section %q not found", err.Name)
}

type KeyNotFoundError struct {
	Name string
}

func (err KeyNotFoundError) Error() string {
	return fmt.Sprintf("key %q not found", err.Name)
}

type InvalidValueError struct {
	Value string
}

func (err InvalidValueError) Error() string {
	return fmt.Sprintf("value %q is invalid", err.Value)
}
