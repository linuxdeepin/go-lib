package operations_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"path/filepath"
	"pkg.deepin.io/lib/gio-2.0"
	. "pkg.deepin.io/lib/operations"
	"testing"
)

// TODO: clean test directories
func TestCopyJob(t *testing.T) {
	SkipConvey("copy a file", t, func() {
		os.Setenv("LANGUAGE", "en_US")
		srcFilePath, _ := filepath.Abs("./testdata/copy/src/afile")
		destPath, _ := filepath.Abs("./testdata/copy/dest")

		job := NewCopyJob([]string{srcFilePath}, destPath, "", gio.FileCopyFlagsNone, nil)
		job.Execute()

		copyedFileURL, _ := filepath.Abs("./testdata/copy/dest/afile")
		copyedFile := gio.FileNewForCommandlineArg(copyedFileURL)
		So(copyedFile.QueryExists(nil), ShouldBeTrue)
	})

	SkipConvey("copy a exists file", t, func() {
		os.Setenv("LANGUAGE", "en_US")
		srcFilePath, _ := filepath.Abs("./testdata/copy/src/exsitfile")
		destPath, _ := filepath.Abs("./testdata/copy/dest")

		job := NewCopyJob([]string{srcFilePath}, destPath, "", gio.FileCopyFlagsNone, renameMock)
		job.Execute()

		copyedFileURL, _ := filepath.Abs("./testdata/copy/dest/exsitfile")
		copyedFile := gio.FileNewForCommandlineArg(copyedFileURL)
		So(copyedFile.QueryExists(nil), ShouldBeTrue)
	})

	SkipConvey("copy a dir", t, func() {
		os.Setenv("LANGUAGE", "en_US")
		srcFilePath, _ := filepath.Abs("./testdata/copy/src/adir")
		destPath, _ := filepath.Abs("./testdata/copy/dest")

		job := NewCopyJob([]string{srcFilePath}, destPath, "", gio.FileCopyFlagsNone, nil)
		job.Execute()

		copyedFileURL, _ := filepath.Abs("./testdata/copy/dest/adir")
		copyedFile := gio.FileNewForCommandlineArg(copyedFileURL)
		So(copyedFile.QueryExists(nil), ShouldBeTrue)
	})

	SkipConvey("copy a exists dir", t, func() {
		os.Setenv("LANGUAGE", "en_US")
		srcFilePath, _ := filepath.Abs("./testdata/copy/src/exsitdir")
		destPath, _ := filepath.Abs("./testdata/copy/dest")

		job := NewCopyJob([]string{srcFilePath}, destPath, "", gio.FileCopyFlagsNone, renameMock)
		job.Execute()

		copyedFileURL, _ := filepath.Abs("./testdata/copy/dest/adir")
		copyedFile := gio.FileNewForCommandlineArg(copyedFileURL)
		So(copyedFile.QueryExists(nil), ShouldBeTrue)
	})
}

// TODO
func TestMoveJob(t *testing.T) {
}
