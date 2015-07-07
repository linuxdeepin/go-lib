package operations

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"pkg.deepin.io/lib/gio-2.0"
	"pkg.deepin.io/lib/glib-2.0"
	"pkg.deepin.io/lib/utils"
	"strings"
)

const (
	ErrorInvalidFileName = iota
	ErrorSameFileName
)

type RenameError struct {
	Code int
	Name string
}

func (e *RenameError) Error() string {
	switch e.Code {
	case ErrorInvalidFileName:
		return fmt.Sprintf("invalid name %q", e.Name)
	case ErrorSameFileName:
		return fmt.Sprintf("the new name(%q) is the same as the old", e.Name)
	default:
		return "Unknown error"
	}
}

func newRenameError(code int, name string) error {
	return &RenameError{Code: code, Name: name}
}

const (
	_RenameJobSignalOldName string = "old-name"
	_RenameJobSignalNewFile string = "new-file"
)

// Rename is just working for current directory.
type RenameJob struct {
	*CommonJob

	file    *gio.File
	newName string

	userDir string // like XDG_VIDEOS_DIR, store on $XDG_CONFIG_HOME/user-dirs.dirs.
}

func (job *RenameJob) emitOldName(oldName string) error {
	return job.Emit(_RenameJobSignalOldName, oldName)
}

func (job *RenameJob) emitNewFile(newFileURL string) error {
	return job.Emit(_RenameJobSignalNewFile, newFileURL)
}

func (job *RenameJob) ListenOldName(fn func(string)) {
	job.ListenSignal(_RenameJobSignalOldName, fn)
}

func (job *RenameJob) ListenNewFile(fn func(string)) {
	job.ListenSignal(_RenameJobSignalNewFile, fn)
}

func (job *RenameJob) checkLocale(keyFile *glib.KeyFile) string {
	locale := ""
	names := GetLanguageNames()
	for _, localeName := range names {
		name, err := keyFile.GetLocaleString(glib.KeyFileDesktopGroup, glib.KeyFileDesktopKeyName, localeName)
		if name != "" && err == nil {
			locale = localeName
			break
		}
	}
	return locale
}

func getUserConfigDir() string {
	h := os.Getenv("XDG_CONFIG_HOME")
	if h == "" {
		home := os.Getenv("HOME")
		h = filepath.Join(home, ".config")
	}
	return h
}

func getUserDirsPath() string {
	// userConfigDir := glib.GetUserConfigDir() // cannot reload
	userConfigDir := getUserConfigDir()
	return filepath.Join(userConfigDir, "user-dirs.dirs")
}

func (job *RenameJob) checkUserDirs() {
	userDirs := getUserDirsPath()

	fileContent, err := ioutil.ReadFile(userDirs)
	if err != nil || string(fileContent) == "" {
		return
	}

	fileReader := bytes.NewReader(fileContent)
	scanner := bufio.NewScanner(fileReader)
	for scanner.Scan() {
		lineText := strings.TrimSpace(scanner.Text())
		if lineText == "" || lineText[0] == '#' {
			continue
		}

		values := strings.SplitN(lineText, "=", 2)
		if len(values) != 2 {
			continue
		}
		userDir := strings.TrimSpace(values[0])
		value := strings.TrimSpace(values[1])

		filePath := job.file.GetPath()
		value = os.ExpandEnv(value)
		if value == filePath {
			job.userDir = userDir
			break
		}
	}
}

func (job *RenameJob) changeUserDir() error {
	buffer := bytes.NewBuffer([]byte{})
	writer := bufio.NewWriter(buffer)
	userDirs := getUserDirsPath()
	fileContent, err := ioutil.ReadFile(userDirs)
	if err != nil {
		return err
	}

	fileReader := bytes.NewReader(fileContent)
	scanner := bufio.NewScanner(fileReader)
	for scanner.Scan() {
		originLineText := scanner.Text()

		lineText := strings.TrimSpace(originLineText)
		if len(lineText) == 0 || lineText[0] == '#' {
			writer.WriteString(originLineText)
			writer.WriteString("\n")
			continue
		}

		values := strings.SplitN(lineText, "=", 2)
		userDir := strings.TrimSpace(values[0])

		if userDir == job.userDir {
			writer.WriteString(userDir)
			writer.WriteString("=\"$HOME")
			writer.WriteRune(filepath.Separator)
			writer.WriteString(job.newName)
			writer.WriteString("\"\n")
			break
		}
	}
	writer.Flush()
	stat, err := os.Stat(userDirs)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(userDirs, buffer.Bytes(), stat.Mode())
}

func (job *RenameJob) setDesktopName() error {
	keyFile := glib.NewKeyFile()
	defer keyFile.Free()
	filePath := job.file.GetPath()
	_, err := keyFile.LoadFromFile(filePath, glib.KeyFileFlagsNone)
	if err != nil {
		return err
	}

	appInfo := gio.NewDesktopAppInfoFromKeyfile(keyFile)
	job.emitOldName(appInfo.GetDisplayName())
	appInfo.Unref()

	locale := job.checkLocale(keyFile)
	if locale != "" {
		keyFile.SetLocaleString(glib.KeyFileDesktopGroup, glib.KeyFileDesktopKeyName, locale, job.newName)
	} else {
		keyFile.SetString(glib.KeyFileDesktopGroup, glib.KeyFileDesktopKeyName, job.newName)
	}

	_, content, err := keyFile.ToData()
	if err != nil {
		return err
	}

	utils.WriteStringToKeyFile(filePath, content)
	return nil
}

func (job *RenameJob) init() {
	job.checkUserDirs()
}

func (job *RenameJob) finalize() {
	defer job.CommonJob.finalize()
	job.file.Unref()
}

func (job *RenameJob) isValidName(name string) bool {
	return name == "." || name == ".." || strings.ContainsRune(name, filepath.Separator)
}

func (job *RenameJob) Execute() {
	defer job.finalize()
	defer job.emitDone()

	if job.isValidName(job.newName) {
		job.setError(newRenameError(ErrorInvalidFileName, job.newName))
		return
	}

	info, err := job.file.QueryInfo(strings.Join(
		[]string{
			gio.FileAttributeStandardContentType,
			gio.FileAttributeStandardDisplayName,
			gio.FileAttributeAccessCanExecute,
		}, ","),
		gio.FileQueryInfoFlagsNofollowSymlinks, nil)
	if err != nil {
		job.setError(err)
		return
	}
	defer info.Unref()

	displayName := info.GetDisplayName()
	if displayName == "" {
		displayName = job.file.GetBasename()
	}

	if displayName == job.newName {
		job.setError(newRenameError(ErrorSameFileName, job.newName))
		return
	}

	mimeType := info.GetContentType()
	if mimeType == _DesktopMIMEType &&
		info.GetAttributeBoolean(gio.FileAttributeAccessCanExecute) {
		err = job.setDesktopName()
		if err != nil {
			job.setError(err)
		}
	} else {
		job.emitOldName(displayName)
		newFile, err := job.file.SetDisplayName(job.newName, job.cancellable)
		if newFile != nil {
			job.emitNewFile(newFile.GetUri())
			newFile.Unref()
		}
		if err != nil {
			job.setError(err)
			return
		}

		if job.userDir != "" {
			err = job.changeUserDir()
			if err != nil {
				job.setError(err)
			}
		}
	}
}

func newRenameJob(file *gio.File, newName string) *RenameJob {
	job := &RenameJob{
		CommonJob: newCommon(nil),
		file:      file,
		newName:   newName,
	}
	job.init()
	return job
}

func NewRenameJob(file string, newName string) *RenameJob {
	gfile := gio.FileNewForCommandlineArg(file)
	return newRenameJob(gfile, newName)
}
