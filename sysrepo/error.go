// -*- coding: utf-8 -*-

package sysrepo

/*
#include <sysrepo.h>
*/
import "C"

import (
	"fmt"
)

//
// sr_error_t
//
type Error int

const (
	ERR_OK                Error = C.SR_ERR_OK
	ERR_INVAL_ARG         Error = C.SR_ERR_INVAL_ARG
	ERR_NOMEM             Error = C.SR_ERR_NOMEM
	ERR_NOT_FOUND         Error = C.SR_ERR_NOT_FOUND
	ERR_INTERNAL          Error = C.SR_ERR_INTERNAL
	ERR_INIT_FAILED       Error = C.SR_ERR_INIT_FAILED
	ERR_IO                Error = C.SR_ERR_IO
	ERR_DISCONNECT        Error = C.SR_ERR_DISCONNECT
	ERR_MALFORMED_MSG     Error = C.SR_ERR_MALFORMED_MSG
	ERR_UNSUPPORTED       Error = C.SR_ERR_UNSUPPORTED
	ERR_UNKNOWN_MODEL     Error = C.SR_ERR_UNKNOWN_MODEL
	ERR_BAD_ELEMENT       Error = C.SR_ERR_BAD_ELEMENT
	ERR_VALIDATION_FAILED Error = C.SR_ERR_VALIDATION_FAILED
	ERR_OPERATION_FAILED  Error = C.SR_ERR_OPERATION_FAILED
	ERR_DATA_EXISTS       Error = C.SR_ERR_DATA_EXISTS
	ERR_DATA_MISSING      Error = C.SR_ERR_DATA_MISSING
	ERR_UNAUTHORIZED      Error = C.SR_ERR_UNAUTHORIZED
	ERR_INVAL_USER        Error = C.SR_ERR_INVAL_USER
	ERR_LOCKED            Error = C.SR_ERR_LOCKED
	ERR_TIME_OUT          Error = C.SR_ERR_TIME_OUT
	ERR_RESTART_NEEDED    Error = C.SR_ERR_RESTART_NEEDED
	ERR_VERSION_MISMATCH  Error = C.SR_ERR_VERSION_MISMATCH
)

var srError_names = map[Error]string{
	ERR_OK:                "ERR_OK",
	ERR_INVAL_ARG:         "ERR_INVAL_ARG",
	ERR_NOMEM:             "ERR_NOMEM",
	ERR_NOT_FOUND:         "ERR_NOT_FOUND",
	ERR_INTERNAL:          "ERR_INTERNAL",
	ERR_INIT_FAILED:       "ERR_INIT_FAILED",
	ERR_IO:                "ERR_IO",
	ERR_DISCONNECT:        "ERR_DISCONNECT",
	ERR_MALFORMED_MSG:     "ERR_MALFORMED_MSG",
	ERR_UNSUPPORTED:       "ERR_UNSUPPORTED",
	ERR_UNKNOWN_MODEL:     "ERR_UNKNOWN_MODEL",
	ERR_BAD_ELEMENT:       "ERR_BAD_ELEMENT",
	ERR_VALIDATION_FAILED: "ERR_VALIDATION_FAILED",
	ERR_OPERATION_FAILED:  "ERR_OPERATION_FAILED",
	ERR_DATA_EXISTS:       "ERR_DATA_EXISTS",
	ERR_DATA_MISSING:      "ERR_DATA_MISSING",
	ERR_UNAUTHORIZED:      "ERR_UNAUTHORIZED",
	ERR_INVAL_USER:        "ERR_INVAL_USER",
	ERR_LOCKED:            "ERR_LOCKED",
	ERR_TIME_OUT:          "ERR_TIME_OUT",
	ERR_RESTART_NEEDED:    "ERR_RESTART_NEEDED",
	ERR_VERSION_MISMATCH:  "ERR_VERSION_MISMATCH",
}

var srError_values = map[string]Error{
	"ERR_OK":                ERR_OK,
	"ERR_INVAL_ARG":         ERR_INVAL_ARG,
	"ERR_NOMEM":             ERR_NOMEM,
	"ERR_NOT_FOUND":         ERR_NOT_FOUND,
	"ERR_INTERNAL":          ERR_INTERNAL,
	"ERR_INIT_FAILED":       ERR_INIT_FAILED,
	"ERR_IO":                ERR_IO,
	"ERR_DISCONNECT":        ERR_DISCONNECT,
	"ERR_MALFORMED_MSG":     ERR_MALFORMED_MSG,
	"ERR_UNSUPPORTED":       ERR_UNSUPPORTED,
	"ERR_UNKNOWN_MODEL":     ERR_UNKNOWN_MODEL,
	"ERR_BAD_ELEMENT":       ERR_BAD_ELEMENT,
	"ERR_VALIDATION_FAILED": ERR_VALIDATION_FAILED,
	"ERR_OPERATION_FAILED":  ERR_OPERATION_FAILED,
	"ERR_DATA_EXISTS":       ERR_DATA_EXISTS,
	"ERR_DATA_MISSING":      ERR_DATA_MISSING,
	"ERR_UNAUTHORIZED":      ERR_UNAUTHORIZED,
	"ERR_INVAL_USER":        ERR_INVAL_USER,
	"ERR_LOCKED":            ERR_LOCKED,
	"ERR_TIME_OUT":          ERR_TIME_OUT,
	"ERR_RESTART_NEEDED":    ERR_RESTART_NEEDED,
	"ERR_VERSION_MISMATCH":  ERR_VERSION_MISMATCH,
}

func (v Error) String() string {
	if s, ok := srError_names[v]; ok {
		return s
	}
	return fmt.Sprintf("Error(%d)", v)
}

func ParseError(rc C.int) error {
	if rc == C.SR_ERR_OK {
		return nil
	}

	return fmt.Errorf("%s", Error(rc))
}

//
// sr_error_info_t
//
type ErrorInfo C.sr_error_info_t

func (v *ErrorInfo) C() *C.sr_error_info_t {
	return (*C.sr_error_info_t)(v)
}

func (v *ErrorInfo) Message() string {
	return C.GoString(v.message)
}

func (v *ErrorInfo) XPath() string {
	return C.GoString(v.xpath)
}
