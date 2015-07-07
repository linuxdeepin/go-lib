package operations_test

import (
	. "pkg.deepin.io/lib/operations"
	"net/url"
	"path/filepath"
)

var testdataDir, _ = filepath.Abs("./testdata")

type UIMock struct {
	skip bool
}

func (*UIMock) AskRetry(primaryText string, secondaryText string, detailText string) Response {
	return NewResponse(ResponseSkip, false)
}
func (*UIMock) AskDeleteConfirmation(primaryText string, secondaryText string, detailText string) bool {
	return true
}

func (*UIMock) AskDelete(primaryText string, secondaryText string, detailText string, flags UIFlags) Response {
	return NewResponse(ResponseSkip, true)
}

func (mock *UIMock) ConflictDialog() Response {
	code := ResponseAutoRename
	if mock.skip {
		code = ResponseSkip
	}
	return NewResponse(code, true)
}

func (*UIMock) AskSkip(primaryText string, secondaryText string, detailText string, flags UIFlags) Response {
	// TODO:
	return NewResponse(ResponseSkip, true)
}

func NewUIMock(skip bool) *UIMock {
	return &UIMock{
		skip: skip,
	}
}

var skipMock = NewUIMock(true)
var renameMock = NewUIMock(false)

func pathToURL(path string) (*url.URL, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	return url.Parse(absPath)
}
