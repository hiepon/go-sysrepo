// -*- coding: utf-8 -*-

package sysrepo

/*
#include <stdio.h>
#include <sysrepo.h>
#include <sysrepo/values.h>
#include <sysrepo/trees.h>
#include "helper.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

//
// sr_node_t
//
type Node C.sr_node_t

func (n *Node) C() *C.sr_node_t {
	return (*C.sr_node_t)(n)
}

func (n *Node) CVal() *C.sr_val_t {
	return C._sr_node_to_val(n.C())
}

func (n *Node) Free() {
	C.sr_free_tree(n.C())
}

func (n *Node) Type() Type {
	return Type(n._type)
}

func (n *Node) Name() string {
	return C.GoString(n.name)
}

func (n *Node) SetName(name string) error {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	rc := C.sr_node_set_name(n.C(), c_name)
	return ParseError(rc)
}

func (n *Node) XPath() string {
	return n.Name()
}

func (n *Node) SetXPath(name string) error {
	return n.SetName(name)
}

func (v *Node) SetDefault(dflt bool) {
	v.dflt = C.bool(dflt)
}

func (n *Node) Default() bool {
	return bool(n.dflt)
}

func (n *Node) Data() *Data {
	return (*Data)(&n.data)
}

func (n *Node) SetModuleName(name string) error {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	rc := C.sr_node_set_module(n.C(), c_name)
	return ParseError(rc)
}

func (n *Node) ModuleName() string {
	if n.module_name == nil {
		return ""
	}

	return C.GoString(n.module_name)
}

func (n *Node) SetStrData(t Type, data string) error {
	c_data := C.CString(data)
	defer C.free(unsafe.Pointer(c_data))

	rc := C.sr_node_set_str_data(n.C(), t.C(), c_data)
	return ParseError(rc)
}

func (n *Node) StrData() string {
	s, _ := n.StrDataSafe()
	return s
}

func (n *Node) StrDataSafe() (string, bool) {
	if n == nil {
		return "", false
	}

	c_str := C.sr_val_to_str(n.CVal())
	if c_str == nil {
		return "", false
	}

	defer C.free(unsafe.Pointer(c_str))
	return C.GoString(c_str), true
}

func (n *Node) InitStrData(name string, data string, t Type) error {
	if err := n.SetName(name); err != nil {
		return err
	}
	return n.SetStrData(t, data)
}

func (n *Node) String() string {
	if n == nil {
		return "<nil>"
	}

	var mem *C.char = nil
	rc := C.sr_print_val_mem(&mem, n.CVal())

	if err := ParseError(rc); err != nil {
		return fmt.Sprintf("%s", err)
	}

	defer C.free(unsafe.Pointer(mem))
	return C.GoString(mem)
}

func (n *Node) FirstChild() *Node {
	return (*Node)(n.first_child)
}

func (n *Node) LastChild() *Node {
	return (*Node)(n.last_child)
}

func (n *Node) Next() *Node {
	return (*Node)(n.next)
}

func (n *Node) Parent() *Node {
	return (*Node)(n.parent)
}

func (n *Node) Prev() *Node {
	return (*Node)(n.prev)
}

func (n *Node) Children(f func(*Node) error) error {
	child := n.FirstChild()
	for child != nil {
		if err := f(child); err != nil {
			return err
		}
		child = child.Next()
	}

	return nil
}

//
// API
//
func NodeIterChild(session *SessionContext, node *Node) *Node {
	c_node := C.sr_node_get_child(session.C(), node.C())
	return (*Node)(c_node)
}

func (n *Node) IterChild(session *SessionContext) *Node {
	return NodeIterChild(session, n)
}

func NodeIterNextSibling(session *SessionContext, node *Node) *Node {
	c_node := C.sr_node_get_next_sibling(session.C(), node.C())
	return (*Node)(c_node)
}

func (n *Node) IterNextSibling(session *SessionContext) *Node {
	return NodeIterNextSibling(session, n)
}

func NodeIterChildren(session *SessionContext, node *Node, f func(*Node) error) error {
	child := node.IterChild(session)
	for child != nil {
		if err := f(child); err != nil {
			return err
		}

		child = child.IterNextSibling(session)
	}

	return nil
}

func (n *Node) IterChildren(session *SessionContext, f func(*Node) error) error {
	return NodeIterChildren(session, n, f)
}

func NodeGetParent(session *SessionContext, node *Node) *Node {
	c_node := C.sr_node_get_parent(session.C(), node.C())
	return (*Node)(c_node)
}

func (n *Node) GetParent(session *SessionContext) *Node {
	return NodeGetParent(session, n)
}

//
// NodeInterface
//
type NodeInterface interface {
	ValInterface
	ModuleName() string
	SetModuleName(string) error
}

//
// GoNode
//
type GoNode struct {
	GoVal
	module string
}

func (n *GoNode) Name() string {
	return n.XPath()
}

func (n *GoNode) SetName(name string) error {
	return n.SetXPath(name)
}

func (n *GoNode) ModuleName() string {
	return n.module
}

func (n *GoNode) SetModuleName(module string) error {
	n.module = module
	return nil
}

func CopyNode(dst, src NodeInterface) error {
	if err := CopyVal(dst, src); err != nil {
		return err
	}

	return dst.SetModuleName(src.ModuleName())
}
