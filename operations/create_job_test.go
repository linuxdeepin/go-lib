package operations_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"os/exec"
	"path/filepath"
	. "pkg.deepin.io/lib/operations"
	"testing"
)

func TestCreateDirectory(t *testing.T) {
	SkipConvey("create directory on /tmp", t, func() {
		destDir := "./testdata/create"
		So(NewCreateDirectoryJob(destDir, "", skipMock).Execute(), ShouldBeNil)
		So(NewCreateDirectoryJob(destDir, "", skipMock).Execute(), ShouldBeNil)
		// TODO: filter keep_me
		files, _ := filepath.Glob(destDir + "/*")
		exec.Command("rmdir", files...).Run()
	})
}

func TestCreateFile(t *testing.T) {
	SkipConvey("create a file without a specific name", t, func() {
		destDir := "./testdata/create"
		So(NewCreateFileJob(destDir, "", []byte{}, skipMock).Execute(), ShouldBeNil)
		So(NewCreateFileJob(destDir, "", []byte{}, skipMock).Execute(), ShouldBeNil)
		files, _ := filepath.Glob(destDir + "/*")
		exec.Command("rm", files...).Run()
	})

	SkipConvey("create a file with a specific name", t, func() {
		destDir := "./testdata/create"
		So(NewCreateFileJob(destDir, "xxxxx", []byte{}, skipMock).Execute(), ShouldBeNil)
		So(NewCreateFileJob(destDir, "xxxxx", []byte{}, skipMock).Execute(), ShouldBeNil)
		files, _ := filepath.Glob(destDir + "/xxxxx*")
		exec.Command("rm", files...).Run()
	})

	SkipConvey("create a file with some init content", t, func() {
		destDir := "./testdata/create"
		So(NewCreateFileJob(destDir, "xxxxx", []byte("xxxxxxx"), skipMock).Execute(), ShouldBeNil)
		files, _ := filepath.Glob(destDir + "/xxxxx*")
		exec.Command("rm", files...).Run()
	})
}

func TestCreateFileFromTemplate(t *testing.T) {
	SkipConvey("create a file from template", t, func() {
		destDir := "./testdata/create"
		templateURL := "/home/liliqiang/Templates/newPowerPoint.ppt"
		So(NewCreateFileFromTemplateJob(destDir, templateURL, skipMock).Execute(), ShouldBeNil)
		So(NewCreateFileFromTemplateJob(destDir, templateURL, skipMock).Execute(), ShouldBeNil)
		files, _ := filepath.Glob(destDir + "/*")
		exec.Command("rm", files...).Run()
	})
}
