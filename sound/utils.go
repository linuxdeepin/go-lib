package sound

import (
	"fmt"
	"os"
	"path"
	"pkg.deepin.io/lib/utils"
)

var (
	errInvalidEvent = fmt.Errorf("invalid theme event")
)

func findThemeFile(theme, event string) (string, error) {
	var home = os.Getenv("HOME")
	var themeDirs = []string{
		path.Join(home, ".local/share/sounds"),
		"/usr/local/share/sounds",
		"/usr/share/sounds",
	}

	for _, dir := range themeDirs {
		// TODO: fix non ogg/oga event
		// TODO: handle sound theme 'index.theme'
		file := path.Join(dir, theme, "stereo", event+".ogg")
		if utils.IsFileExist(file) {
			return file, nil
		}
		file = path.Join(dir, theme, "stereo", event+".oga")
		if utils.IsFileExist(file) {
			return file, nil
		}
	}
	return "", errInvalidEvent
}
