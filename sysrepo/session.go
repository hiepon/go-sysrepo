// -*- coding: utf-8 -*-

package sysrepo

/*
#include <sysrepo.h>
*/
import "C"

import (
	"unsafe"
)

//
// sr_session_ctx_t
//
type SessionContext C.sr_session_ctx_t

func (v *SessionContext) C() *C.sr_session_ctx_t {
	return (*C.sr_session_ctx_t)(v)
}

func SessionStart(conn *ConnContext, datastore Datastore, flag SessionFlag) (*SessionContext, error) {
	var ctxt *C.sr_session_ctx_t
	rc := C.sr_session_start(conn.C(), datastore.C(), C.uint(flag.C()), &ctxt)
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return (*SessionContext)(ctxt), nil
}

func (v *SessionContext) Stop() error {
	rc := C.sr_session_stop(v.C())
	return ParseError(rc)
}

func (v *SessionContext) Refresh() error {
	rc := C.sr_session_refresh(v.C())
	return ParseError(rc)
}

func (v *SessionContext) Check() error {
	rc := C.sr_session_check(v.C())
	return ParseError(rc)
}

func (v *SessionContext) SwitchDatastore(datastore Datastore) error {
	rc := C.sr_session_switch_ds(v.C(), datastore.C())
	return ParseError(rc)
}

func (v *SessionContext) SetOptions(flag SessionFlag) error {
	rc := C.sr_session_set_options(v.C(), C.uint(flag.C()))
	return ParseError(rc)
}

func (v *SessionContext) LastError() (*ErrorInfo, error) {
	var c_errinfo *C.sr_error_info_t

	rc := C.sr_get_last_error(v.C(), &c_errinfo)
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return (*ErrorInfo)(c_errinfo), nil
}

func (v *SessionContext) SetError(message string, xpath string) error {
	c_message := C.CString(message)
	defer C.free(unsafe.Pointer(c_message))

	c_xpath := C.CString(xpath)
	defer C.free(unsafe.Pointer(c_xpath))

	rc := C.sr_set_error(v.C(), c_message, c_xpath)
	return ParseError(rc)
}

func (v *SessionContext) ID() uint32 {
	rc := C.sr_session_get_id(v.C())
	return (uint32)(rc)
}
