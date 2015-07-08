// Theme scanner
package scanner

import (
	"fmt"
	"os"
	"path"
	"pkg.deepin.io/lib/theme/checker"
	dutils "pkg.deepin.io/lib/utils"
)

// uri: ex "file:///usr/share/themes"
func ListGtkTheme(uri string) ([]string, error) {
	return doListTheme(uri, checker.IsGtkTheme)
}

// uri: ex "file:///usr/share/icons"
func ListIconTheme(uri string) ([]string, error) {
	return doListTheme(uri, checker.IsIconTheme)
}

// uri: ex "file:///usr/share/icons"
func ListCursorTheme(uri string) ([]string, error) {
	return doListTheme(uri, checker.IsCursorTheme)
}

func doListTheme(uri string, filter func(string) (bool, error)) ([]string, error) {
	dir := dutils.DecodeURI(uri)
	subDirs, err := listSubDir(dir)
	if err != nil {
		return nil, err
	}

	var themes []string
	for _, subDir := range subDirs {
		tmp := path.Join(subDir, "index.theme")
		ok, _ := filter(tmp)
		if !ok {
			continue
		}
		themes = append(themes, subDir)
	}
	return themes, nil
}

func listSubDir(dir string) ([]string, error) {
	if !dutils.IsDir(dir) {
		return nil, fmt.Errorf("'%s' not a dir", dir)
	}

	fr, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer fr.Close()

	names, err := fr.Readdirnames(-1)
	if err != nil {
		return nil, err
	}

	var subDirs []string
	for _, name := range names {
		tmp := path.Join(dir, name)
		if !dutils.IsDir(tmp) {
			continue
		}

		subDirs = append(subDirs, tmp)
	}
	return subDirs, nil

}
