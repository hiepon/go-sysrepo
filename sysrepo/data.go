// -*- coding: utf-8 -*-

package sysrepo

/*
#include <sysrepo.h>
#include "helper.h"
*/
import "C"

//
// sr_data_u
//
type Data C.sr_data_t

func (v *Data) C() *C.sr_data_t {
	return (*C.sr_data_t)(v)
}

func (v *Data) Bool() bool {
	return bool(C._sr_data_bool_get(v.C()))
}

func (v *Data) SeBool(b bool) {
	C._sr_data_bool_set(v.C(), C.bool(b))
}

func (v *Data) Decimal64() int64 {
	return int64(C._sr_data_decimal64_get(v.C()))
}

func (v *Data) SetDecimal64(d int64) {
	C._sr_data_decimal64_set(v.C(), C.double(d))
}

func (v *Data) Int8() int8 {
	return int8(C._sr_data_int8_get(v.C()))
}

func (v *Data) SetInt8(i int8) {
	C._sr_data_int8_set(v.C(), C.int8_t(i))
}

func (v *Data) Int16() int16 {
	return int16(C._sr_data_int16_get(v.C()))
}

func (v *Data) SetInt16(i int16) {
	C._sr_data_int16_set(v.C(), C.int16_t(i))
}

func (v *Data) Int32() int32 {
	return int32(C._sr_data_int32_get(v.C()))
}

func (v *Data) SetInt32(i int32) {
	C._sr_data_int32_set(v.C(), C.int32_t(i))
}

func (v *Data) Int64() int64 {
	return int64(C._sr_data_int64_get(v.C()))
}

func (v *Data) SetInt64(i int64) {
	C._sr_data_int64_set(v.C(), C.int64_t(i))
}

func (v *Data) Uint8() uint8 {
	return uint8(C._sr_data_uint8_get(v.C()))
}

func (v *Data) SetUint8(i int8) {
	C._sr_data_uint8_set(v.C(), C.uint8_t(i))
}

func (v *Data) Uint16() uint16 {
	return uint16(C._sr_data_uint16_get(v.C()))
}

func (v *Data) SetUint16(i int16) {
	C._sr_data_uint16_set(v.C(), C.uint16_t(i))
}

func (v *Data) Uint32() uint32 {
	return uint32(C._sr_data_uint32_get(v.C()))
}

func (v *Data) SetUint32(i int32) {
	C._sr_data_uint32_set(v.C(), C.uint32_t(i))
}

func (v *Data) Uint64() uint64 {
	return uint64(C._sr_data_uint64_get(v.C()))
}

func (v *Data) SetUint64(i int8) {
	C._sr_data_uint64_set(v.C(), C.uint64_t(i))
}
