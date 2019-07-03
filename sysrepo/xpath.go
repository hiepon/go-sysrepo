// -*- coding: utf-8 -*-

package sysrepo

/*
#include <stdlib.h>
#include <sysrepo/xpath.h>
*/
import "C"

import (
	"strings"
	"unsafe"
)

type XPathContext struct {
	ctx   C.sr_xpath_ctx_t
	xpath *C.char
	iter  *C.char
}

func (v *XPathContext) Ctx() *C.sr_xpath_ctx_t {
	return &v.ctx
}

func NewXPathContext(xpath string) *XPathContext {
	ctx := &XPathContext{
		xpath: C.CString(xpath),
	}
	ctx.Init()
	return ctx
}

func (v *XPathContext) Init() {
	v.iter = v.xpath
}

func (v *XPathContext) Recover() {
	C.sr_xpath_recover(v.Ctx())
	v.Init()
}

func (v *XPathContext) Free() {
	v.iter = nil
	if v.xpath != nil {
		C.free(unsafe.Pointer(v.xpath))
		v.xpath = nil
	}
}

func charToString(s *C.char) (string, bool) {
	if s == nil {
		return "", false
	}
	return C.GoString(s), true
}

func (v *XPathContext) charToString(s *C.char) (string, bool) {
	ss, b := charToString(s)
	if b {
		v.iter = nil
	}
	return ss, b
}

func (v *XPathContext) String() string {
	if v.xpath == nil {
		return ""
	}
	return C.GoString(v.xpath)
}

func (v *XPathContext) NextNode() (string, bool) {
	c_next := C.sr_xpath_next_node(v.iter, v.Ctx())
	return v.charToString(c_next)
}

func (v *XPathContext) NextNodeNS() (string, bool) {
	c_next := C.sr_xpath_next_node_with_ns(v.iter, v.Ctx())
	return v.charToString(c_next)
}

func (v *XPathContext) LastNode() (string, bool) {
	c_last := C.sr_xpath_last_node(v.iter, v.Ctx())
	return v.charToString(c_last)
}

func (v *XPathContext) NextAttrKey() (string, bool) {
	c_key := C.sr_xpath_next_key_name(v.iter, v.Ctx())
	return v.charToString(c_key)
}

func (v *XPathContext) NextAttrValue() (string, bool) {
	c_key := C.sr_xpath_next_key_value(v.iter, v.Ctx())
	return v.charToString(c_key)
}

func (v *XPathContext) Node(name string) (string, bool) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_key := C.sr_xpath_node(v.iter, c_name, v.Ctx())
	return v.charToString(c_key)
}

func (v *XPathContext) NodeRel(name string) (string, bool) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_node := C.sr_xpath_node_rel(v.iter, c_name, v.Ctx())
	return v.charToString(c_node)
}

func (v *XPathContext) NodeIndex(index uint32) (string, bool) {
	c_node := C.sr_xpath_node_idx(v.iter, C.size_t(index), v.Ctx())
	return v.charToString(c_node)
}

func (v *XPathContext) NodeIndexRel(index uint32) (string, bool) {
	c_node := C.sr_xpath_node_idx_rel(v.iter, C.size_t(index), v.Ctx())
	return v.charToString(c_node)
}

func (v *XPathContext) NodeAttrValue(key string) (string, bool) {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_node := C.sr_xpath_node_key_value(v.iter, c_key, v.Ctx())
	return v.charToString(c_node)
}

func (v *XPathContext) NodeAttrValueIndex(index uint32) (string, bool) {
	c_node := C.sr_xpath_node_key_value_idx(v.iter, C.size_t(index), v.Ctx())
	return v.charToString(c_node)
}

func (v *XPathContext) AttrValue(node string, key string) (string, bool) {
	c_node := C.CString(node)
	defer C.free(unsafe.Pointer(c_node))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := C.sr_xpath_key_value(v.iter, c_node, c_key, v.Ctx())
	return v.charToString(c_value)
}

func (v *XPathContext) AttrValueIndex(nodeIndex uint32, keyIndex uint32) (string, bool) {
	c_value := C.sr_xpath_key_value_idx(v.iter, C.size_t(nodeIndex), C.size_t(keyIndex), v.Ctx())
	return v.charToString(c_value)
}

func XPathNodeName(xpath string) (string, bool) {
	c_xpath := C.CString(xpath)
	defer C.free(unsafe.Pointer(c_xpath))

	c_node := C.sr_xpath_node_name(c_xpath)
	return charToString(c_node)
}

func ParseXPathName(name string) (string, string) {
	ss := strings.SplitN(name, ":", 2)
	switch len(ss) {
	case 2:
		return ss[0], ss[1]
	default:
		return "", name
	}
}

func ParseXPath(xpath string, f func(string, string, map[string]string)) {
	c_xpath := C.CString(xpath)
	defer C.free(unsafe.Pointer(c_xpath))

	var state C.sr_xpath_ctx_t

	c_node := C.sr_xpath_next_node_with_ns(c_xpath, &state)
	for c_node != nil {
		ns, name := ParseXPathName(C.GoString(c_node))

		attrs := map[string]string{}
		for {
			c_attr_key := C.sr_xpath_next_key_name(nil, &state)
			if c_attr_key == nil {
				break
			}

			attr_key := C.GoString(c_attr_key)
			attr_val := func() string {
				c_attr_val := C.sr_xpath_next_key_value(nil, &state)
				if c_attr_val == nil {
					return ""
				}
				return C.GoString(c_attr_val)
			}()

			attrs[attr_key] = attr_val
		}

		f(ns, name, attrs)

		c_node = C.sr_xpath_next_node_with_ns(nil, &state)
	}

	C.sr_xpath_recover(&state)
}
