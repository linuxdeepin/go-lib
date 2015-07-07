package operations_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"os/exec"
	"path/filepath"
	"pkg.deepin.io/lib/gio-2.0"
	. "pkg.deepin.io/lib/operations"
	"testing"
)

func TestDeleteJob(t *testing.T) {
	getTrashCount := func() (uint32, error) {
		trash := gio.FileNewForUri("trash:")
		info, err := trash.QueryInfo(gio.FileAttributeTrashItemCount, gio.FileQueryInfoFlagsNone, nil)
		if err != nil {
			return 0, err
		}
		count := info.GetAttributeUint32(gio.FileAttributeTrashItemCount)
		info.Unref()
		trash.Unref()
		return count, nil
	}

	Convey("trash should be empty", t, func() {
		NewEmptyTrashJob(false, skipMock).Execute()

		count, err := getTrashCount()
		if err != nil {
			SkipSo()
		}

		So(count, ShouldEqual, 0)
	})

	SkipConvey("trash one file get one file on trash", t, func() {
		NewEmptyTrashJob(false, skipMock).Execute()
		count, _ := getTrashCount()
		So(count, ShouldEqual, 0)

		a := "./testdata/trashfiles/dest/a"
		NewTrashJob([]string{a}, false, skipMock).Execute()

		// there is a delay, sleep a while.
		// time.Sleep(time.Second * 3)
		count, _ = getTrashCount()
		So(count, ShouldEqual, 1)
	})

	SkipConvey("trash multi files get multi files on trash", t, func() {
		NewEmptyTrashJob(false, skipMock).Execute()
		count, _ := getTrashCount()
		So(count, ShouldEqual, 0)

		b := "./testdata/trashfiles/dest/b"
		c := "./testdata/trashfiles/dest/c"
		NewTrashJob([]string{b, c}, false, skipMock).Execute()

		count, _ = getTrashCount()
		So(count, ShouldEqual, 2)
	})

	Convey("delete a file", t, func() {
		exec.Command("cp", "./testdata/delete/src/todelete.txt", "./testdata/delete/dest").Run()
		target, err := filepath.Abs("./testdata/delete/dest/todelete.txt")
		So(err, ShouldBeNil)
		exec.Command("touch", target).Run()
		job := NewDeleteJob([]string{target}, false, nil)
		job.Execute()
		So(job.HasError(), ShouldBeFalse)
	})

	Convey("delete a dir", t, func() {
		exec.Command("cp", "-r", "./testdata/delete/src/todelete.dir", "./testdata/delete/dest").Run()
		target, _ := filepath.Abs("./testdata/delete/dest/todelete.dir")
		exec.Command("mkdir", target).Run()
		job := NewDeleteJob([]string{target}, false, nil)
		job.Execute()
		So(job.HasError(), ShouldBeFalse)
		f := gio.FileNewForCommandlineArg(target)
		So(f.QueryExists(nil), ShouldBeFalse)
		f.Unref()
	})
}
