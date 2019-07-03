// -*- coding: utf-8 -*-

package sysrepo

/*
#include <sysrepo.h>
#include "helper.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

//
// sr_conn_flag_t
//
type ConnFlag C.sr_conn_flag_t

func (v ConnFlag) C() C.sr_conn_flag_t {
	return (C.sr_conn_flag_t)(v)
}

const (
	CONN_DEFAULT         ConnFlag = C.SR_CONN_DEFAULT
	CONN_DAEMON_REQUIRED ConnFlag = C.SR_CONN_DAEMON_REQUIRED
	CONN_DAEMON_START    ConnFlag = C.SR_CONN_DAEMON_START
)

var srConnFlag_names = map[ConnFlag]string{
	CONN_DEFAULT:         "DEFAULT",
	CONN_DAEMON_REQUIRED: "DAEMON_REQUIRED",
	CONN_DAEMON_START:    "DAEMON_START",
}

var srConnFlag_values = map[string]ConnFlag{
	"DEFAULT":         CONN_DEFAULT,
	"DAEMON_REQUIRED": CONN_DAEMON_REQUIRED,
	"DAEMON_START":    CONN_DAEMON_START,
}

func (v ConnFlag) String() string {
	if s, ok := srConnFlag_names[v]; ok {
		return s
	}
	return fmt.Sprintf("ConnFlag(%d)", v)
}

func ParseConnFlag(s string) (ConnFlag, error) {
	if v, ok := srConnFlag_values[s]; ok {
		return v, nil
	}
	return CONN_DEFAULT, fmt.Errorf("Invalid ConnFlag. %s", s)
}

//
// sr_conn_options_t
//
type ConnOptions C.sr_conn_options_t

func (v ConnOptions) C() C.sr_conn_options_t {
	return (C.sr_conn_options_t)(v)
}

//
// sr_conn_ctx_t
//
type ConnContext C.sr_conn_ctx_t

func (v *ConnContext) C() *C.sr_conn_ctx_t {
	return (*C.sr_conn_ctx_t)(v)
}

func Connect(appName string, flag ConnFlag) (*ConnContext, error) {
	c_appname := C.CString(appName)
	defer C.free(unsafe.Pointer(c_appname))

	var ctx *C.sr_conn_ctx_t
	rc := C.sr_connect(c_appname, C.uint(flag.C()), &ctx)
	if err := ParseError(rc); err != nil {
		return nil, err
	}

	return (*ConnContext)(ctx), nil
}

func (v *ConnContext) Disconnect() {
	if v != nil {
		C.sr_disconnect(v.C())
	}
}
