package operations_test

import (
	. "github.com/smartystreets/goconvey/convey"
	. "pkg.deepin.io/lib/operations"
	"testing"
)

func TestGetLaunchAppInfo(t *testing.T) {
	// FIXME: how to make a stable test???
	// Convey("get launch app info", t, func() {
	// 	uri, _ := pathToURL("./testdata/launchapp/test.c")
	// 	job := NewLaunchAppJob(uri)
	// 	appInfo := job.Execute()
	// 	So(job.HasError(), ShouldBeFalse)
	// 	t.Log(appInfo)
	// })
}

func TestSetLaunchAppInfo(t *testing.T) {
	// FIXME: how to make a stable test???
	SkipConvey("set launch app info", t, func() {
		mimeType := "text/html"
		job := NewSetDefaultLaunchAppJob("google-chrome.desktop", mimeType)
		job.Execute()
		So(job.HasError(), ShouldBeFalse)
	})
}
