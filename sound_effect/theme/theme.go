// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package theme

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"github.com/linuxdeepin/go-lib/keyfile"
	"github.com/linuxdeepin/go-lib/strv"
	"github.com/linuxdeepin/go-lib/xdg/basedir"
	"strings"
	"sync"
)

type Theme struct {
	InternalName string
	Path         string
	Inherits     []string
	Directories  []string
	SubDirs      []SubDir

	//Name string
	//Comment string
	//Hidden bool
	//Example string
}

type SubDir struct {
	Name          string
	OutputProfile string
	//Context string
}

type Finder struct {
	dataDirs     []string
	loadedThemes map[string]*Theme
	mu           sync.Mutex
}

func NewFinder() *Finder {
	const soundsDir = "sounds"
	dataDirs := strv.Strv{filepath.Join(basedir.GetUserDataDir(), soundsDir)}
	for _, dir := range basedir.GetSystemDataDirs() {
		dataDirs, _ = dataDirs.Add(filepath.Join(dir, soundsDir))
	}
	return &Finder{
		dataDirs: dataDirs,
	}
}

func (finder *Finder) GetTheme(name string) (*Theme, error) {
	finder.mu.Lock()
	defer finder.mu.Unlock()

	if theme, ok := finder.loadedThemes[name]; ok {
		return theme, nil
	}

	for _, dataDir := range finder.dataDirs {
		fileInfos, err := ioutil.ReadDir(dataDir)
		if err != nil {
			continue
		}
		for _, fileInfo := range fileInfos {
			if !fileInfo.IsDir() {
				continue
			}

			if fileInfo.Name() != name {
				continue
			}

			// ex. /usr/share/sounds/ freedesktop index.theme
			theme, err := LoadTheme(filepath.Join(dataDir, fileInfo.Name(), "index.theme"))
			if err != nil {
				continue
			}
			// save theme in cache
			if finder.loadedThemes == nil {
				finder.loadedThemes = make(map[string]*Theme)
			}
			finder.loadedThemes[theme.InternalName] = theme
			return theme, nil
		}
	}

	return nil, errors.New("not found theme " + name)
}

func LoadTheme(filename string) (*Theme, error) {
	kf := keyfile.NewKeyFile()
	kf.ListSeparator = ','
	err := kf.LoadFromFile(filename)
	if err != nil {
		return nil, err
	}
	const mainSection = "Sound Theme"
	_, err = kf.GetSection(mainSection)
	if err != nil {
		return nil, err
	}
	inherits, _ := kf.GetStringList(mainSection, "Inherits")
	directories, _ := kf.GetStringList(mainSection, "Directories")

	var subDirs []SubDir
	for _, section := range directories {
		outputProfile, _ := kf.GetString(section, "OutputProfile")
		if outputProfile != "" {
			subDirs = append(subDirs, SubDir{
				OutputProfile: outputProfile,
				Name:          section,
			})
		}
	}

	path := filepath.Dir(filename)
	t := &Theme{
		Path:         path,
		InternalName: filepath.Base(path),
		Inherits:     inherits,
		Directories:  directories,
		SubDirs:      subDirs,
	}
	return t, nil
}

func isNameOk(name string) bool {
	if strings.ContainsRune(name, '/') || name == "" ||
		name == "." || name == ".." {
		return false
	}
	return true
}

func (finder *Finder) Find(reqTheme, reqOutputProfile, reqName string) string {
	if !(isNameOk(reqTheme) && isNameOk(reqName)) {
		return ""
	}

	filename := finder.find(reqTheme, reqOutputProfile, reqName)
	if filename != "" {
		return filename
	}

	const fallbackTheme = "freedesktop"
	if reqTheme != fallbackTheme {
		// fallback
		filename = finder.find(fallbackTheme, reqOutputProfile, reqName)
		if filename != "" {
			return filename
		}
	}
	return ""
}

func (finder *Finder) find(reqTheme, reqOutputProfile, reqName string) string {
	theme, err := finder.GetTheme(reqTheme)
	if err != nil {
		return ""
	}

	filename := finder.lookup(theme, reqOutputProfile, reqName)
	if filename != "" {
		return filename
	}

	// recursively find from parents
	for _, parent := range theme.Inherits {
		parentTheme, err := finder.GetTheme(parent)
		if err != nil {
			continue
		}

		filename = finder.find(parentTheme.InternalName, reqOutputProfile, reqName)
		if filename != "" {
			return filename
		}
	}
	return ""
}

func (finder *Finder) lookup(theme *Theme, reqOutputProfile, reqName string) string {
	profiles := strv.Strv{reqOutputProfile}
	profiles, _ = profiles.Add("stereo")

	for _, profile := range profiles {
		for _, subDir := range theme.SubDirs {
			if subDir.OutputProfile == profile {
				for _, directory := range finder.dataDirs {
					for _, name := range getNames(reqName, nil) {
						for _, ext := range []string{"oga", "ogg", "wav"} {
							filename := filepath.Join(directory, theme.InternalName, subDir.Name, name+"."+ext)
							statInfo, err := os.Stat(filename)
							if err != nil {
								continue
							}

							if statInfo.IsDir() {
								continue
							}
							return filename
						}
					}
				}
			}
		}
	}
	return ""
}

func getNames(name string, acc []string) []string {
	idx := strings.LastIndex(name, "-")
	if idx < 0 {
		return append(acc, name)
	} else {
		return getNames(name[:idx], append(acc, name))
	}
}
