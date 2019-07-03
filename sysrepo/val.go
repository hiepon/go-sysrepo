// -*- coding: utf-8 -*-

package sysrepo

/*
#include <stdio.h>
#include <sysrepo.h>
#include <sysrepo/values.h>
*/
import "C"

import (
	"fmt"
	"io"
	"strings"
	"unsafe"
)

//
// Val
//
type Val C.sr_val_t

func (v *Val) C() *C.sr_val_t {
	return (*C.sr_val_t)(v)
}

func (v *Val) Free() {
	C.sr_free_val(v.C())
}

func (v *Val) Type() Type {
	return Type(v._type)
}

func (v *Val) XPath() string {
	return C.GoString(v.xpath)
}

func (v *Val) SetXPath(xpath string) error {
	c_xpath := C.CString(xpath)
	defer C.free(unsafe.Pointer(c_xpath))

	rc := C.sr_val_set_xpath(v.C(), c_xpath)
	return ParseError(rc)
}

func (v *Val) SetDefault(dflt bool) {
	v.dflt = C.bool(dflt)
}

func (v *Val) Default() bool {
	return bool(v.dflt)
}

func (v *Val) Data() *Data {
	return (*Data)(&v.data)
}

func (v *Val) SetStrData(t Type, strVal string) error {
	c_strval := C.CString(strVal)
	defer C.free(unsafe.Pointer(c_strval))

	rc := C.sr_val_set_str_data(v.C(), t.C(), c_strval)
	return ParseError(rc)
}

func (v *Val) StrData() string {
	s, _ := v.StrDataSafe()
	return s
}

func (v *Val) StrDataSafe() (string, bool) {
	if v == nil {
		return "", false
	}

	c_str := C.sr_val_to_str(v.C())
	if c_str == nil {
		return "", false
	}

	defer C.free(unsafe.Pointer(c_str))
	return C.GoString(c_str), true
}

func (v *Val) InitStrData(xpath string, strdata string, t Type) error {
	if err := v.SetXPath(xpath); err != nil {
		return err
	}
	return v.SetStrData(t, strdata)
}

func (v *Val) String() string {
	if v == nil {
		return "<nil>"
	}

	var mem *C.char = nil
	rc := C.sr_print_val_mem(&mem, v.C())

	if err := ParseError(rc); err != nil {
		return fmt.Sprintf("%s", err)
	}

	defer C.free(unsafe.Pointer(mem))
	return C.GoString(mem)
}

func (v *Val) WriteTo(w io.Writer) (int64, error) {
	c_str := C.sr_val_to_str(v.C())
	if c_str == nil {
		return 0, fmt.Errorf("sr_val_to_str error.")
	}
	C.free(unsafe.Pointer(c_str))

	s := C.GoString(c_str)
	return strings.NewReader(s).WriteTo(w)
}

//
// sr_val_iter_t
//
type ValIter C.sr_val_iter_t

func (v *ValIter) C() *C.sr_val_iter_t {
	return (*C.sr_val_iter_t)(v)
}

func (v *ValIter) Free() {
	C.sr_free_val_iter(v.C())
}

//
// GoVal
//
type GoVal struct {
	t     Type
	xpath string
	data  string
	dflt  bool
}

func NewGoValArray(n int) []*GoVal {
	vals := make([]*GoVal, n)
	for index := 0; index < n; index++ {
		vals[index] = &GoVal{}
	}
	return vals
}

func (v *GoVal) Type() Type {
	return v.t
}

func (v *GoVal) SetXPath(xpath string) error {
	v.xpath = xpath
	return nil
}

func (v *GoVal) XPath() string {
	return v.xpath
}

func (v *GoVal) SetStrData(t Type, data string) error {
	v.t = t
	v.data = data
	return nil
}

func (v *GoVal) StrData() string {
	return v.data
}

func (v *GoVal) SetDefault(dflt bool) {
	v.dflt = dflt
}

func (v *GoVal) Default() bool {
	return v.dflt
}

func (v *GoVal) InitStrData(xpath string, data string, t Type) error {
	v.SetXPath(xpath)
	v.SetStrData(t, data)
	return nil
}

func (v *GoVal) String() string {
	return fmt.Sprintf("%s = %s", v.xpath, v.data)
}

//
// Val interface
//
type ValInterface interface {
	Type() Type
	SetXPath(xpath string) error
	XPath() string
	SetStrData(t Type, data string) error
	StrData() string
	SetDefault(dflt bool)
	Default() bool
}

func CopyVal(dst, src ValInterface) error {
	if err := dst.SetXPath(src.XPath()); err != nil {
		return err
	}

	if err := dst.SetStrData(src.Type(), src.StrData()); err != nil {
		return err
	}

	dst.SetDefault(src.Default())
	return nil
}
