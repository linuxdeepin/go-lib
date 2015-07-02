package operations_test

import (
	. "pkg.linuxdeepin.com/lib/operations"
	. "github.com/smartystreets/goconvey/convey"
	"net/url"
	"os/exec"
	"path/filepath"
	"pkg.linuxdeepin.com/lib/gio-2.0"
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

	createURL := func(path string) (*url.URL, error) {
		a, err := pathToURL(path)
		if err != nil {
			return nil, err
		}

		if a.Scheme == "" {
			a.Scheme = "file"
		}
		return a, nil
	}

	SkipConvey("trash one file get one file on trash", t, func() {
		NewEmptyTrashJob(false, skipMock).Execute()
		count, _ := getTrashCount()
		So(count, ShouldEqual, 0)

		a, _ := createURL("./testdata/trashfiles/dest/a")
		NewTrashJob([]*url.URL{a}, false, skipMock).Execute()

		// there is a delay, sleep a while.
		// time.Sleep(time.Second * 3)
		count, _ = getTrashCount()
		So(count, ShouldEqual, 1)
	})

	SkipConvey("trash multi files get multi files on trash", t, func() {
		NewEmptyTrashJob(false, skipMock).Execute()
		count, _ := getTrashCount()
		So(count, ShouldEqual, 0)

		b, _ := createURL("./testdata/trashfiles/dest/b")
		c, _ := createURL("./testdata/trashfiles/dest/c")
		NewTrashJob([]*url.URL{b, c}, false, skipMock).Execute()

		count, _ = getTrashCount()
		So(count, ShouldEqual, 2)
	})

	Convey("delete a file", t, func() {
		exec.Command("cp", "./testdata/delete/src/todelete.txt", "./testdata/delete/dest").Run()
		target, err := filepath.Abs("./testdata/delete/dest/todelete.txt")
		So(err, ShouldBeNil)
		exec.Command("touch", target).Run()
		urls, _ := url.Parse(target)
		job := NewDeleteJob([]*url.URL{urls}, false, nil)
		job.Execute()
		So(job.HasError(), ShouldBeFalse)
	})

	Convey("delete a dir", t, func() {
		exec.Command("cp", "-r", "./testdata/delete/src/todelete.dir", "./testdata/delete/dest").Run()
		target, _ := filepath.Abs("./testdata/delete/dest/todelete.dir")
		exec.Command("mkdir", target).Run()
		urls, _ := url.Parse(target)
		job := NewDeleteJob([]*url.URL{urls}, false, nil)
		job.Execute()
		So(job.HasError(), ShouldBeFalse)
		f := gio.FileNewForCommandlineArg(target)
		So(f.QueryExists(nil), ShouldBeFalse)
		f.Unref()
	})
}
