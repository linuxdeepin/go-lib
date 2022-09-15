// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package desktopappinfo

import (
	"os"
	"path/filepath"
	"strings"
	//"fmt"
)

type AppsDir struct {
	Path     string
	AppNames map[string]int
}

func getDirs(skipDirs map[string][]string) []AppsDir {
	var dirs []AppsDir
	for _, path := range xdgAppDirs {
		appsDir := AppsDir{
			Path: path,
		}

		appsDir.AppNames = getAppNames(path, skipDirs[path])
		dirs = append(dirs, appsDir)
	}

	return dirs
}

func GetAll(skipDirs map[string][]string) []*DesktopAppInfo {
	var infos []*DesktopAppInfo
	dirs := getDirs(skipDirs)

	infoMap := make(map[string]int)
	var count int
	for _, dir := range dirs {
		for appName := range dir.AppNames {

			if count > 0 {
				if _, ok := infoMap[appName]; ok {
					// appName masked
					continue
				}
			}

			filename := filepath.Join(dir.Path, appName)
			ai, err := NewDesktopAppInfoFromFile(filename)
			if err != nil {
				continue
			}
			infoMap[appName] = 0
			if ai.ShouldShow() {
				infos = append(infos, ai)
			}
		}

		count++
	}

	return infos
}

func getAppNames(root string, skipDirs []string) map[string]int {
	ret := make(map[string]int)
	Walk(root, func(name string, info os.FileInfo) bool {
		if info.IsDir() {
			return stringSliceContains(skipDirs, name)
		}

		if strings.HasSuffix(name, desktopExt) {
			ret[name] = 0
		}

		return false
	})
	return ret
}

func readDirNames(dirname string) ([]string, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	names, err := f.Readdirnames(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	return names, nil
}

func Walk(root string, walkFn WalkFunc) {
	info, err := os.Stat(root)
	if err != nil {
		return
	}
	walk(root, ".", info, walkFn)
}

func walk(root, name0 string, info os.FileInfo, walkFn WalkFunc) {
	//fmt.Println("root:", root, "name0:", name0)
	skipDir := walkFn(name0, info)
	if skipDir {
		//fmt.Println("skip ", name0)
		return
	}
	if !info.IsDir() {
		return
	}
	path := filepath.Join(root, name0)
	names, err := readDirNames(path)
	if err != nil {
		return
	}
	for _, name := range names {
		filename := filepath.Join(path, name)
		fileInfo, err := os.Lstat(filename)
		if err != nil {
			continue
		}
		walk(root, filepath.Join(name0, name), fileInfo, walkFn)
	}
}

type WalkFunc func(name string, info os.FileInfo) bool

func stringSliceContains(slice []string, str string) bool {
	for _, v := range slice {
		if str == v {
			return true
		}
	}
	return false
}
