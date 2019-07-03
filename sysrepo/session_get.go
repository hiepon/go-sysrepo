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

func (v *SessionContext) Item(xpath string, cb func(*Val) error) error {
	var c_val *C.sr_val_t

	c_xpath := C.CString(xpath)
	defer C.free(unsafe.Pointer(c_xpath))

	rc := C.sr_get_item(v.C(), c_xpath, &c_val)
	if err := ParseError(rc); err != nil {
		return err
	}

	defer C.sr_free_val(c_val)

	return cb((*Val)(c_val))
}

func (v *SessionContext) Items(xpath string, cb func(*Val) error) error {
	var c_val *C.sr_val_t
	var c_cnt C.size_t

	c_xpath := C.CString(xpath)
	defer C.free(unsafe.Pointer(c_xpath))

	rc := C.sr_get_items(v.C(), c_xpath, &c_val, &c_cnt)
	if err := ParseError(rc); err != nil {
		return err
	}

	defer C.sr_free_values(c_val, c_cnt)

	var index C.size_t
	for index = 0; index < c_cnt; index++ {
		val := C._sr_val_array_get(c_val, index)
		if err := cb((*Val)(val)); err != nil {
			return err
		}
	}

	return nil
}

func (v *SessionContext) ItemsRange(xpath string, cb func(*Val) error) error {
	c_xpath := C.CString(xpath)
	defer C.free(unsafe.Pointer(c_xpath))

	var c_iter *C.sr_val_iter_t
	rc := C.sr_get_items_iter(v.C(), c_xpath, &c_iter)
	if err := ParseError(rc); err != nil {
		return err
	}

	defer C.sr_free_val_iter(c_iter)

	for {
		var c_val *C.sr_val_t
		rc := C.sr_get_item_next(v.C(), c_iter, &c_val)
		if rc == C.SR_ERR_NOT_FOUND {
			return nil
		}
		if err := ParseError(rc); err != nil {
			return err
		}

		err := cb((*Val)(c_val))
		C.sr_free_val(c_val)

		if err != nil {
			return err
		}
	}
}

func (v *SessionContext) Subtree(xpath string, flag GetSubtreeFlag, cb func(*Node) error) error {
	c_xpath := C.CString(xpath)
	defer C.free(unsafe.Pointer(c_xpath))

	var c_subtree *C.sr_node_t
	rc := C.sr_get_subtree(v.C(), c_xpath, C.uint(flag), &c_subtree)
	if err := ParseError(rc); err != nil {
		return err
	}

	defer C.sr_free_tree(c_subtree)

	return cb((*Node)(c_subtree))
}

func (v *SessionContext) Subtrees(xpath string, flag GetSubtreeFlag, cb func(*Node) error) error {
	c_xpath := C.CString(xpath)
	defer C.free(unsafe.Pointer(c_xpath))

	var c_subtrees *C.sr_node_t
	var c_cnt C.size_t
	rc := C.sr_get_subtrees(v.C(), c_xpath, C.uint(flag), &c_subtrees, &c_cnt)
	if err := ParseError(rc); err != nil {
		return err
	}

	defer C.sr_free_trees(c_subtrees, c_cnt)

	var c_index C.size_t
	for c_index = 0; c_index < c_cnt; c_index++ {
		c_subtree := C._sr_node_array_get(c_subtrees, c_index)
		if err := cb((*Node)(c_subtree)); err != nil {
			return err
		}
	}

	return nil
}
