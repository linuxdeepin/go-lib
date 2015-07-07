package operations_test

import (
	"bufio"
	"bytes"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"pkg.deepin.io/lib/gio-2.0"
	. "pkg.deepin.io/lib/operations"
	"strings"
	"testing"
)

func TestRenameFile(t *testing.T) {
	var renameSrcDir = filepath.Join(testdataDir, "rename", "src")
	var renameTestDir = filepath.Join(testdataDir, "rename", "test")

	Convey("rename a file", t, func() {
		src := filepath.Join(renameSrcDir, "afile")
		err := exec.Command("cp", src, renameTestDir).Run()
		So(err, ShouldBeNil)

		targetPath := filepath.Join(renameTestDir, "afile")
		defer exec.Command("rm", targetPath).Run()
		newName := "newFileName"
		job := NewRenameJob(targetPath, newName)
		job.Execute()
		f, err := os.Open(targetPath)
		if f != nil {
			f.Close()
		}
		So(err, ShouldNotBeNil)
		So(os.IsNotExist(err), ShouldBeTrue)

		newTargetPath := filepath.Join(renameTestDir, newName)
		defer exec.Command("rm", newTargetPath).Run()

		stat, err := os.Stat(newTargetPath)
		So(err, ShouldBeNil)
		So(stat.Name(), ShouldEqual, newName)
	})

	Convey("rename a dir", t, func() {
		src := filepath.Join(renameSrcDir, "adir")
		err := exec.Command("cp", "-r", src, renameTestDir).Run()
		So(err, ShouldBeNil)

		targetPath := filepath.Join(renameTestDir, "adir")
		defer exec.Command("rm", "-r", targetPath).Run()
		newName := "newDirName"
		job := NewRenameJob(targetPath, newName)
		job.Execute()
		f, err := os.Open(targetPath)
		if f != nil {
			f.Close()
		}
		So(err, ShouldNotBeNil)
		So(os.IsNotExist(err), ShouldBeTrue)

		newTargetPath := filepath.Join(renameTestDir, newName)
		defer exec.Command("rm", "-r", newTargetPath).Run()

		stat, err := os.Stat(newTargetPath)
		So(err, ShouldBeNil)
		So(stat.Name(), ShouldEqual, newName)
	})

	Convey("rename a desktop dir", t, func() {
		xdgConfigDir := filepath.Join(testdataDir, "rename", "config")
		os.Setenv("HOME", renameTestDir)
		os.Setenv("XDG_CONFIG_HOME", xdgConfigDir)
		userDirsConfigFile := filepath.Join(xdgConfigDir, "user-dirs.dirs")

		ioutil.WriteFile(userDirsConfigFile, []byte("XDG_DESKTOP_DIR=$HOME/Desktop"), 0664)
		defer exec.Command("rm", userDirsConfigFile).Run()

		src := filepath.Join(renameSrcDir, "Desktop")
		targetPath := filepath.Join(renameTestDir, "Desktop")
		defer exec.Command("rm", "-r", targetPath).Run()
		exec.Command("cp", "-r", src, targetPath).Run()

		newName := "newDesktop"
		job := NewRenameJob(targetPath, newName)
		job.Execute()

		f, err := os.Open(targetPath)
		if f != nil {
			f.Close()
		}
		So(err, ShouldNotBeNil)
		So(os.IsNotExist(err), ShouldBeTrue)

		newTargetPath := filepath.Join(renameTestDir, newName)
		defer exec.Command("rm", "-r", newTargetPath).Run()

		stat, err := os.Stat(newTargetPath)
		So(err, ShouldBeNil)
		So(stat.Name(), ShouldEqual, newName)

		fileContent, err := ioutil.ReadFile(userDirsConfigFile)
		So(err, ShouldBeNil)
		So(string(fileContent), ShouldNotEqual, "")

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

			value = os.ExpandEnv(value) // value will contain double quote(")
			if userDir == "XDG_DESKTOP_DIR" {
				So(value, ShouldEqual, fmt.Sprintf("%q", newTargetPath))
				return
			}
		}
		panic("no reach here")
	})

	Convey("rename a non-executable desktop file", t, func() {
		src := filepath.Join(renameSrcDir, "nonexecutable.desktop")
		targetPath := filepath.Join(renameTestDir, "nonexecutable.desktop")
		exec.Command("cp", src, targetPath).Run()
		defer exec.Command("rm", targetPath).Run()

		p := gio.FileNewForPath(targetPath)
		So(p, ShouldNotBeNil)
		info, _ := p.QueryInfo(gio.FileAttributeStandardDisplayName, gio.FileQueryInfoFlagsNofollowSymlinks, nil)
		So(info, ShouldNotBeNil)
		So(info.GetDisplayName(), ShouldEqual, "nonexecutable.desktop")
		info.Unref()
		p.Unref()

		newName := "newDesktopName.desktop"
		job := NewRenameJob(targetPath, newName)
		job.Execute()

		newTargetPath := filepath.Join(renameTestDir, newName)
		defer exec.Command("rm", newTargetPath).Run()
		a := gio.FileNewForPath(newTargetPath)
		So(a, ShouldNotBeNil)
		info, _ = a.QueryInfo(gio.FileAttributeStandardDisplayName, gio.FileQueryInfoFlagsNofollowSymlinks, nil)
		So(info, ShouldNotBeNil)
		So(info.GetDisplayName(), ShouldEqual, newName)
		info.Unref()
		a.Unref()
	})

	// TODO: this test won't be pass on jenkins.
	SkipConvey("rename a executable desktop file", t, func() {
		Convey("change with en_US locale", func() {
			os.Setenv("LANGUAGE", "en_US")
			newName := "enName"

			src := filepath.Join(renameSrcDir, "executable.desktop")
			targetPath := filepath.Join(renameTestDir, "enexecutable.desktop")
			exec.Command("cp", src, targetPath).Run()
			defer exec.Command("rm", targetPath).Run()

			p := gio.NewDesktopAppInfoFromFilename(targetPath)
			So(p, ShouldNotBeNil)
			So(p.GetDisplayName(), ShouldEqual, "Executable")
			p.Unref()

			job := NewRenameJob(targetPath, newName)
			job.Execute()

			a := gio.NewDesktopAppInfoFromFilename(targetPath)
			So(a, ShouldNotBeNil)
			So(a.GetDisplayName(), ShouldEqual, newName)
			a.Unref()
		})

		Convey("change with zh_CN locale", func() {
			os.Setenv("LANGUAGE", "zh_CN")
			newName := "中文名"

			src := filepath.Join(renameSrcDir, "executable.desktop")
			targetPath := filepath.Join(renameTestDir, "zhexecutable.desktop")
			exec.Command("cp", src, targetPath).Run()
			defer exec.Command("rm", targetPath).Run()

			p := gio.NewDesktopAppInfoFromFilename(targetPath)
			So(p, ShouldNotBeNil)
			So(p.GetDisplayName(), ShouldEqual, "可执行")
			p.Unref()

			job := NewRenameJob(targetPath, newName)
			job.Execute()

			a := gio.NewDesktopAppInfoFromFilename(targetPath)
			So(a, ShouldNotBeNil)
			So(a.GetDisplayName(), ShouldEqual, newName)
			a.Unref()
		})
	})
}
