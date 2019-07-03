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

func (v *SessionContext) SetItem(xpath string, val *Val, flag EditFlag) error {
	c_xpath := C.CString(xpath)
	defer C.free(unsafe.Pointer(c_xpath))

	rc := C.sr_set_item(v.C(), c_xpath, val.C(), C.uint(flag))
	return ParseError(rc)
}

func (v *SessionContext) SetItemStr(xpath string, val string, flag EditFlag) error {
	c_xpath := C.CString(xpath)
	defer C.free(unsafe.Pointer(c_xpath))

	c_val := C.CString(val)
	defer C.free(unsafe.Pointer(c_val))

	rc := C.sr_set_item_str(v.C(), c_xpath, c_val, C.uint(flag))
	return ParseError(rc)
}

func (v *SessionContext) DeleteItem(xpath string, flag EditFlag) error {
	c_xpath := C.CString(xpath)
	defer C.free(unsafe.Pointer(c_xpath))

	rc := C.sr_delete_item(v.C(), c_xpath, C.uint(flag))
	return ParseError(rc)
}

func (v *SessionContext) MoveItem(xpath string, pos MovePosition, relativeItem string) error {
	c_xpath := C.CString(xpath)
	defer C.free(unsafe.Pointer(c_xpath))

	c_relative := C.CString(relativeItem)
	defer C.free(unsafe.Pointer(c_relative))

	rc := C.sr_move_item(v.C(), c_xpath, pos.C(), c_relative)
	return ParseError(rc)
}

func (v *SessionContext) Validate() error {
	rc := C.sr_validate(v.C())
	return ParseError(rc)
}

func (v *SessionContext) Commit() error {
	rc := C.sr_commit(v.C())
	return ParseError(rc)
}

func (v *SessionContext) DiscardChanges() error {
	rc := C.sr_discard_changes(v.C())
	return ParseError(rc)
}

func (v *SessionContext) CopyConfig(module string, src, dst Datastore) error {
	c_module := C.CString(module)
	defer C.free(unsafe.Pointer(c_module))

	rc := C.sr_copy_config(v.C(), c_module, src.C(), dst.C())
	return ParseError(rc)
}
