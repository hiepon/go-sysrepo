// -*- coding: utf-8 -*-

package sysrepo

/*
#include <stdio.h>
#include <sysrepo.h>
#include "helper.h"
*/
import "C"

import (
	"unsafe"
)

func (v *SessionContext) LockDatastore() error {
	rc := C.sr_lock_datastore(v.C())
	return ParseError(rc)
}
func (v *SessionContext) UnlockDatastore() error {
	rc := C.sr_unlock_datastore(v.C())
	return ParseError(rc)
}

func (v *SessionContext) LockModule(module string) error {
	c_module := C.CString(module)
	defer C.free(unsafe.Pointer(c_module))

	rc := C.sr_lock_module(v.C(), c_module)
	return ParseError(rc)
}

func (v *SessionContext) UnlockModule(module string) error {
	c_module := C.CString(module)
	defer C.free(unsafe.Pointer(c_module))

	rc := C.sr_unlock_module(v.C(), c_module)
	return ParseError(rc)
}
