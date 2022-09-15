// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package pam

/*
#cgo LDFLAGS: -lpam
#include <security/pam_appl.h>
#include <stdlib.h>

void init_pam_conv(struct pam_conv *conv, long c);
*/
import "C"
import (
	"fmt"
	"strings"
	"unsafe"
)

// Style is the type of message that the conversation handler should display.
type Style int

// Coversation handler style types.
const (
	// PromptEchoOff indicates the conversation handler should obtain a
	// string without echoing any text.
	PromptEchoOff Style = C.PAM_PROMPT_ECHO_OFF
	// PromptEchoOn indicates the conversation handler should obtain a
	// string while echoing text.
	PromptEchoOn = C.PAM_PROMPT_ECHO_ON
	// ErrorMsg indicates the conversation handler should display an
	// error message.
	ErrorMsg = C.PAM_ERROR_MSG
	// TextInfo indicates the conversation handler should display some
	// text.
	TextInfo = C.PAM_TEXT_INFO
)

// return values
// const (
// 	Success          = C.PAM_SUCCESS
// 	SystemErr        = C.PAM_SYSTEM_ERR
// 	BadItem          = C.PAM_BAD_ITEM
// 	BufErr           = C.PAM_BUF_ERR
// 	PermDenied       = C.PAM_PERM_DENIED
// 	Abort            = C.PAM_ABORT
// 	AuthErr          = C.PAM_AUTH_ERR
// 	CredInsufficient = C.PAM_CRED_INSUFFICIENT
// 	// AuthInfoUnavailable   = C.PAM_AUTHINFO_UNVAIL
// 	MaxTries              = C.PAM_MAXTRIES
// 	UserUnknown           = C.PAM_USER_UNKNOWN
// 	CredErr               = C.PAM_CRED_ERR
// 	CredExpired           = C.PAM_CRED_EXPIRED
// 	CredUnavailable       = C.PAM_CRED_UNAVAIL
// 	AccountExpired        = C.PAM_ACCT_EXPIRED
// 	NewAuthTokenRequired  = C.PAM_NEW_AUTHTOK_REQD
// 	AuthTokenErr          = C.PAM_AUTHTOK_ERR
// 	AuthTokenRecoveryErr  = C.PAM_AUTHTOK_RECOVERY_ERR
// 	AuthTokenLockBusy     = C.PAM_AUTHTOK_LOCK_BUSY
// 	AuthTokenDisableAging = C.PAM_AUTHTOK_DISABLE_AGING
// 	TryAgain              = C.PAM_TRY_AGAIN
// 	SessionErr            = C.PAM_SESSION_ERR
// )

type Error struct {
	Code int
	Msg  string
}

func (err Error) Error() string {
	return fmt.Sprintf("pam error(%d): %s", err.Code, err.Msg)
}

type Transaction struct {
	ptr       *C.pam_handle_t
	handlerId handlerId
	status    C.int
}

func Start(service, user string, handler ConversationHandler) (*Transaction, error) {
	t := &Transaction{
		handlerId: addHandler(handler),
	}
	service0 := C.CString(service)
	var user0 *C.char
	if user != "" {
		user0 = C.CString(user)
	}

	var conv C.struct_pam_conv
	C.init_pam_conv(&conv, C.long(t.handlerId))
	ret := C.pam_start(service0, user0, &conv, &t.ptr)

	// clean
	C.free(unsafe.Pointer(service0))
	if user0 != nil {
		C.free(unsafe.Pointer(user0))
	}

	err := t.toErr(ret)
	if err != nil {
		deleteHandler(t.handlerId)
		return nil, err
	}
	return t, nil
}

func StartFunc(service, user string, handler func(Style, string) (string, error)) (*Transaction, error) {
	return Start(service, user, ConversationFunc(handler))
}

func (h *Transaction) End(lastStatus int) error {
	deleteHandler(h.handlerId)
	ret := C.pam_end(h.ptr, C.int(lastStatus))
	return h.toErr(ret)
}

func (h *Transaction) LastStatus() int {
	return int(h.status)
}

func (h *Transaction) toErr(errNum C.int) error {
	if errNum == C.PAM_SUCCESS {
		return nil
	}

	msg := C.GoString(C.pam_strerror(h.ptr, errNum))
	return Error{
		Code: int(errNum),
		Msg:  msg,
	}
}

// Item is a an PAM information type.
type Item int

// PAM Item types.
const (
	// Service is the name which identifies the PAM stack.
	Service Item = C.PAM_SERVICE
	// User identifies the username identity used by a service.
	User = C.PAM_USER
	// Tty is the terminal name.
	Tty = C.PAM_TTY
	// Rhost is the requesting host name.
	Rhost = C.PAM_RHOST
	// Authtok is the currently active authentication token.
	Authtok = C.PAM_AUTHTOK
	// Oldauthtok is the old authentication token.
	Oldauthtok = C.PAM_OLDAUTHTOK
	// Ruser is the requesting user name.
	Ruser = C.PAM_RUSER
	// UserPrompt is the string use to prompt for a username.
	UserPrompt = C.PAM_USER_PROMPT
)

func (h *Transaction) SetItemStr(itemType Item, item string) error {
	item0 := unsafe.Pointer(C.CString(item))
	ret := C.pam_set_item(h.ptr, C.int(itemType), item0)
	C.free(item0)
	return h.toErr(ret)
}

func (h *Transaction) GetItemStr(itemType Item) (string, error) {
	var item0 unsafe.Pointer
	ret := C.pam_get_item(h.ptr, C.int(itemType), &item0)
	err := h.toErr(ret)
	if err != nil {
		return "", err
	}
	item := C.GoString((*C.char)(item0))
	return item, nil
}

// Flags are inputs to various PAM functions than be combined with a bitwise
// or. Refer to the official PAM documentation for which flags are accepted
// by which functions.
type Flags int

// PAM Flag types.
const (
	// Silent indicates that no messages should be emitted.
	Silent Flags = C.PAM_SILENT
	// DisallowNullAuthtok indicates that authorization should fail
	// if the user does not have a registered authentication token.
	DisallowNullAuthtok = C.PAM_DISALLOW_NULL_AUTHTOK
	// EstablishCred indicates that credentials should be established
	// for the user.
	EstablishCred = C.PAM_ESTABLISH_CRED
	// DeleteCred inidicates that credentials should be deleted.
	DeleteCred = C.PAM_DELETE_CRED
	// ReinitializeCred indicates that credentials should be fully
	// reinitialized.
	ReinitializeCred = C.PAM_REINITIALIZE_CRED
	// RefreshCred indicates that the lifetime of existing credentials
	// should be extended.
	RefreshCred = C.PAM_REFRESH_CRED
	// ChangeExpiredAuthtok indicates that the authentication token
	// should be changed if it has expired.
	ChangeExpiredAuthtok = C.PAM_CHANGE_EXPIRED_AUTHTOK
)

func (h *Transaction) Authenticate(flags Flags) error {
	ret := C.pam_authenticate(h.ptr, C.int(flags))
	h.status = ret
	return h.toErr(ret)
}

func (h *Transaction) SetCredentials(flags Flags) error {
	ret := C.pam_setcred(h.ptr, C.int(flags))
	h.status = ret
	return h.toErr(ret)
}

func (h *Transaction) AccountValidManagement(flags Flags) error {
	ret := C.pam_acct_mgmt(h.ptr, C.int(flags))
	h.status = ret
	return h.toErr(ret)
}

func (h *Transaction) ChangeAuthToken(flags Flags) error {
	ret := C.pam_chauthtok(h.ptr, C.int(flags))
	h.status = ret
	return h.toErr(ret)
}

func (h *Transaction) OpenSession(flags Flags) error {
	ret := C.pam_open_session(h.ptr, C.int(flags))
	h.status = ret
	return h.toErr(ret)
}

func (h *Transaction) CloseSession(flags Flags) error {
	ret := C.pam_close_session(h.ptr, C.int(flags))
	h.status = ret
	return h.toErr(ret)
}

func (h *Transaction) PutEnv(name, value string) error {
	env := C.CString(name + "=" + value)
	ret := C.pam_putenv(h.ptr, env)
	C.free(unsafe.Pointer(env))
	h.status = ret
	return h.toErr(ret)
}

func (h *Transaction) DeleteEnv(name string) error {
	name0 := C.CString(name)
	ret := C.pam_putenv(h.ptr, name0)
	C.free(unsafe.Pointer(name0))
	h.status = ret
	return h.toErr(ret)
}

func (h *Transaction) GetEnv(name string) string {
	name0 := C.CString(name)
	ret := C.pam_getenv(h.ptr, name0)
	C.free(unsafe.Pointer(name0))
	return C.GoString(ret)
}

func next(p **C.char) **C.char {
	return (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Sizeof(p)))
}

// GetEnvList returns a copy of the PAM environment as a map.
func (h *Transaction) GetEnvList() map[string]string {
	envMap := make(map[string]string)
	p := C.pam_getenvlist(h.ptr)
	if p == nil {
		return nil
	}
	for q := p; *q != nil; q = next(q) {
		chunks := strings.SplitN(C.GoString(*q), "=", 2)
		if len(chunks) == 2 {
			envMap[chunks[0]] = chunks[1]
		}
		C.free(unsafe.Pointer(*q))
	}
	C.free(unsafe.Pointer(p))
	return envMap
}

// cbPAMConv is a wrapper for the conversation callback function.
//export cbPAMConv
func cbPAMConv(style C.int, msg *C.char, id C.long, resp **C.char) C.int {
	var r string
	var err error
	handler := getHandler(handlerId(id))
	r, err = handler.RespondPAM(Style(style), C.GoString(msg))
	if err != nil {
		return C.PAM_CONV_ERR
	}
	*resp = C.CString(r)
	return C.PAM_SUCCESS
}
