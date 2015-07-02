package operations

import (
	"fmt"
)

// type DeletionType int
//
// const (
// 	DELETE DeletionType = iota
// 	TRASH
// 	EMPTY_TRASH
// )
//
// type ConfirmationType int
//
// const (
// 	DefaultConfirmation ConfirmationType = iota
// 	ForceConfirmation
// )

// ResponseCode is a type for the response of UIDelegate.
type ResponseCode int32

// the code for response of UIDelegate.
const (
	ResponseCancel ResponseCode = 1 << iota
	ResponseSkip
	ResponseRetry
	ResponseDelete
	ResponseOverwrite
	ResponseAutoRename // auto rename the conflict file/directory
	ResponseYes
)

// String returns a human readable string for ResponseCode.
func (code ResponseCode) String() string {
	switch code {
	case ResponseCancel:
		return "Cancel"
	case ResponseSkip:
		return "Skip"
	case ResponseRetry:
		return "Retry"
	case ResponseDelete:
		return "Delete"
	case ResponseYes:
		return "Yes"
	case ResponseOverwrite:
		return "Overwrite"
	case ResponseAutoRename:
		return "AutoRename"
	}

	return fmt.Sprintf("Unknow code: %d", int32(code))
}

// Response stores the response relavant information like ResponseCode.
type Response struct {
	code       int32
	applyToAll bool
	userData   string
}

// NewResponse creates a new Response from response code and apply to all.
func NewResponse(code ResponseCode, applyToAll bool) Response {
	return Response{
		code:       int32(code),
		applyToAll: applyToAll,
	}
}

// String returns a human readable string for debuging or something like it.
func (response Response) String() string {
	str := ResponseCode(response.code).String()
	if response.applyToAll {
		str += " to all"
	}

	return str
}

// Code returns response code.
func (response Response) Code() ResponseCode {
	return ResponseCode(response.code)
}

// ApplyToAll returns whether apply to all.
func (response Response) ApplyToAll() bool {
	return response.applyToAll
}

// UserData returns some extra data.
func (response Response) UserData() string {
	return response.userData
}

type UIFlags int32

// UIFlags
const (
	UIFlagsNone UIFlags = iota << 0
	UIFlagsRetry
	UIFlagsMulti
)

func (flags UIFlags) String() string {
	strFlags := "UIFlags("
	if flags&UIFlagsRetry != 0 {
		strFlags += "retry"
	}

	if flags&UIFlagsMulti != 0 {
		if strFlags != "UIFlags(" {
			strFlags += "|"
		}
		strFlags += "multi"
	}

	if strFlags == "UIFlags(" {
		strFlags += "NONE"
	}

	return strFlags + ")"
}

// IUIDelegate is the interface for ui delegate.
type IUIDelegate interface {
	// TODO: using this internally, give a simpler interface, like kio,
	// a.k.a: AskDeleteConfirmation(urls, deleteType, confirmationType))
	// if necessary, ask user to confirm whether to delete or trash files.
	AskDeleteConfirmation(primaryText string, secondaryText string, detailText string) bool

	AskDelete(string, string, string, UIFlags) Response
	AskSkip(primaryText string, secondaryText string, detailText string, uiFlags UIFlags) Response
	AskRetry(primaryText string, secondaryText string, detailText string) Response

	// TODO: decide arguments
	ConflictDialog() Response
}

type _DefaultUIDelegate struct{}

func (*_DefaultUIDelegate) AskRetry(primaryText string, secondaryText string, detailText string) Response {
	return NewResponse(ResponseCancel, false)
}

func (*_DefaultUIDelegate) AskDeleteConfirmation(primaryText string, secondaryText string, detailText string) bool {
	return true
}

func (*_DefaultUIDelegate) AskDelete(primaryText string, secondaryText string, detailText string, flags UIFlags) Response {
	return NewResponse(ResponseSkip, true)
}

func (*_DefaultUIDelegate) AskSkip(primaryText string, secondaryText string, detailText string, flags UIFlags) Response {
	fmt.Printf("AskSkip:\n\tprimaryText: %s\n\tsecondaryText: %s\n\tdetailText: %s\n\tflags: %v\n", primaryText, secondaryText, detailText, flags)
	// TODO:
	return NewResponse(ResponseSkip, true)
}

func (*_DefaultUIDelegate) ConflictDialog() Response {
	return NewResponse(ResponseSkip, true)
}

var _deafultUIDelegate = &_DefaultUIDelegate{}
