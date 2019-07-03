// -*- coding: utf-8 -*-

package sysrepo

import (
	"testing"
)

const (
	TEST_XPATH_LEAF = "/example-module:container/list[key1='key=\"A'][key2='\"ke=yB\"']/leaf"
	TEST_XPATH_LIST = "/example-module:container/list[key1='key=\"A'][key2='\"ke=yB\"']"
	TEST_XPATH_AUG  = "/ietf-interfaces:interfaces/interface[name='eth0']/ietf-ip:ipv4/address[ip='192.168.2.100']/prefix-length"
)

func mapcmp(m1, m2 map[string]string) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k1, v1 := range m1 {
		v2, ok := m2[k1]
		if !ok {
			return false
		}
		if v1 != v2 {
			return false
		}
	}
	return true
}

func TestXPath_NextNode(t *testing.T) {
	xpath := ""
	ctx := NewXPathContext(TEST_XPATH_LEAF)
	defer ctx.Free()

	var s string
	var b bool

	// 'container'
	if s, b = ctx.NextNode(); !b {
		t.Errorf("XPath.NextNode error. %s %t", xpath, b)
	}
	if s != "container" {
		t.Errorf("XPath.NextNode unmatch. %s %s", xpath, s)
	}

	// 'list'
	if s, b = ctx.NextNode(); !b {
		t.Errorf("XPath.NextNode error. %s %t", xpath, b)
	}
	if s != "list" {
		t.Errorf("XPath.NextNode unmatch. %s %s", xpath, s)
	}

	// 'leaf'
	if s, b = ctx.NextNode(); !b {
		t.Errorf("XPath.NextNode error. %s %t", xpath, b)
	}
	if s != "leaf" {
		t.Errorf("XPath.NextNode unmatch. %s %s", xpath, s)
	}

	// no node
	if _, b = ctx.NextNode(); b {
		t.Errorf("XPath.NextNode must be error. %s %t", xpath, b)
	}
}

func TestXPath_NextNodeNS(t *testing.T) {
	xpath := ""
	ctx := NewXPathContext(TEST_XPATH_LEAF)
	defer ctx.Free()

	var s string
	var b bool

	// 'container'
	if s, b = ctx.NextNodeNS(); !b {
		t.Errorf("XPath.NextNode error. %s %t", xpath, b)
	}
	if s != "example-module:container" {
		t.Errorf("XPath.NextNode unmatch. '%s' '%s'", xpath, s)
	}

	// 'list'
	if s, b = ctx.NextNodeNS(); !b {
		t.Errorf("XPath.NextNode error. %s %t", xpath, b)
	}
	if s != "list" {
		t.Errorf("XPath.NextNode unmatch. %s %s", xpath, s)
	}

	// 'leaf'
	if s, b = ctx.NextNodeNS(); !b {
		t.Errorf("XPath.NextNode error. %s %t", xpath, b)
	}
	if s != "leaf" {
		t.Errorf("XPath.NextNode unmatch. %s %s", xpath, s)
	}

	// no node
	if _, b = ctx.NextNode(); b {
		t.Errorf("XPath.NextNode must be error. %s %t", xpath, b)
	}
}

func TestXPath_NextKey(t *testing.T) {
	xpath := ""
	ctx := NewXPathContext(TEST_XPATH_LEAF)
	defer ctx.Free()

	var s string
	var b bool

	// <root node>
	if s, b = ctx.NextAttrKey(); b {
		t.Errorf("XPath.NextKey must be error. %s %t", xpath, b)
	}

	// 'container'
	if _, b = ctx.NextNode(); !b {
		t.Errorf("XPath.NextNode(container) error. %s %t", xpath, b)
	}
	if _, b = ctx.NextAttrKey(); b {
		t.Errorf("XPath.NextKey must be error. %s %t", xpath, b)
	}

	// 'list...
	if _, b = ctx.NextNode(); !b {
		t.Errorf("XPath.NextNode(list) error. %s %t", xpath, b)
	}
	// key1=...
	if s, b = ctx.NextAttrKey(); !b {
		t.Errorf("XPath.NextKey error. %s %t", xpath, b)
	}
	if s != "key1" {
		t.Errorf("XPath.NextNode unmatch. %s %s", xpath, s)
	}
	// key2=...
	if s, b = ctx.NextAttrKey(); !b {
		t.Errorf("XPath.NextKey error. %s %t", xpath, b)
	}
	if s != "key2" {
		t.Errorf("XPath.NextNode unmatch. %s %s", xpath, s)
	}
	// no key
	if s, b = ctx.NextAttrKey(); b {
		t.Errorf("XPath.NextKey must be error. %s %t %s", xpath, b, s)
	}

	// 'leaf'
	if s, b = ctx.NextNodeNS(); !b {
		t.Errorf("XPath.NextNode error. %s %t", xpath, b)
	}
	if s != "leaf" {
		t.Errorf("XPath.NextNode unmatch. %s %s", xpath, s)
	}

	// no node
	if _, b = ctx.NextNode(); b {
		t.Errorf("XPath.NextNode must be error. %s %t", xpath, b)
	}
}

func TestXPath_NextAttrValue(t *testing.T) {
	xpath := ""
	ctx := NewXPathContext(TEST_XPATH_LEAF)
	defer ctx.Free()

	var s string
	var b bool

	// <root node>
	if s, b = ctx.NextAttrValue(); b {
		t.Errorf("XPath.NextKey must be error. %s %t", xpath, b)
	}

	// 'container'
	if _, b = ctx.NextNode(); !b {
		t.Errorf("XPath.NextNode(container) error. %s %t", xpath, b)
	}
	if _, b = ctx.NextAttrValue(); b {
		t.Errorf("XPath.NextKey must be error. %s %t", xpath, b)
	}

	// 'list...
	if _, b = ctx.NextNode(); !b {
		t.Errorf("XPath.NextNode(list) error. %s %t", xpath, b)
	}
	// key1=...
	if s, b = ctx.NextAttrValue(); !b {
		t.Errorf("XPath.NextKey error. %s %t", xpath, b)
	}
	if s != "key=\"A" {
		t.Errorf("XPath.NextNode unmatch. %s %s", xpath, s)
	}
	// key2=...
	if s, b = ctx.NextAttrValue(); !b {
		t.Errorf("XPath.NextKey error. %s %t", xpath, b)
	}
	if s != "\"ke=yB\"" {
		t.Errorf("XPath.NextNode unmatch. %s %s", xpath, s)
	}
	// no key
	if s, b = ctx.NextAttrValue(); b {
		t.Errorf("XPath.NextKey must be error. %s %t %s", xpath, b, s)
	}

	// 'leaf'
	if s, b = ctx.NextNodeNS(); !b {
		t.Errorf("XPath.NextNode error. %s %t", xpath, b)
	}
	if s != "leaf" {
		t.Errorf("XPath.NextNode unmatch. %s %s", xpath, s)
	}

	// no node
	if _, b = ctx.NextNode(); b {
		t.Errorf("XPath.NextNode must be error. %s %t", xpath, b)
	}
}

func TestXPath_Node(t *testing.T) {
	xpath := ""
	ctx := NewXPathContext(TEST_XPATH_LEAF)
	defer ctx.Free()

	var s string
	var b bool

	// "leaf"
	if s, b = ctx.Node("leaf"); !b {
		t.Errorf("XPath.Node error. %s %t", xpath, b)
	}
	if s != "leaf" {
		t.Errorf("XPath.Node unmatch. %s", s)
	}

	// 'container'
	if s, b = ctx.Node("container"); !b {
		t.Errorf("XPath.Node error. %s %t", xpath, b)
	}
	if s != "container" {
		t.Errorf("XPath.Node unmatch. %s", s)
	}

	// 'list'
	if s, b = ctx.Node("list"); !b {
		t.Errorf("XPath.Node error. %s %t", xpath, b)
	}
	if s != "list" {
		t.Errorf("XPath.Node unmatch. %s", s)
	}

	// 'unknown'
	if _, b = ctx.Node("unknown"); b {
		t.Errorf("XPath.Node must be error. %s %t", xpath, b)
	}
}

func TestParseXPath0(t *testing.T) {
	xpath := "/"
	cnt := 0
	ParseXPath(xpath, func(ns, name string, attrs map[string]string) {
		cnt += 1
		if cnt > 1 {
			t.Errorf("ParseXPath unmatch. cnt=%d", cnt)
		}
		if ns != "" || name != "" {
			t.Errorf("ParseXPath unmatch. node=%s:%s", ns, name)
		}
		if len(attrs) != 0 {
			t.Errorf("ParseXPath unmatch. attrs=%v", attrs)
		}
	})
}

func TestParseXPath0_ns(t *testing.T) {
	xpath := "/ns1:"
	cnt := 0
	ParseXPath(xpath, func(ns, name string, attrs map[string]string) {
		cnt += 1
		if cnt > 1 {
			t.Errorf("ParseXPath unmatch. cnt=%d", cnt)
		}
		if ns != "ns1" || name != "" {
			t.Errorf("ParseXPath unmatch. node=%s:%s", ns, name)
		}
		if len(attrs) != 0 {
			t.Errorf("ParseXPath unmatch. attrs=%v", attrs)
		}
	})
}

func TestParseXPath1(t *testing.T) {
	xpath := "/node1"
	cnt := 0
	ParseXPath(xpath, func(ns, name string, attrs map[string]string) {
		cnt += 1
		if cnt > 1 {
			t.Errorf("ParseXPath unmatch. cnt=%d", cnt)
		}
		if ns != "" || name != "node1" {
			t.Errorf("ParseXPath unmatch. node=%s:%s", ns, name)
		}
		if len(attrs) != 0 {
			t.Errorf("ParseXPath unmatch. attrs=%v", attrs)
		}
	})
}

func TestParseXPath1_ns(t *testing.T) {
	xpath := "/ns1:node1"
	cnt := 0
	ParseXPath(xpath, func(ns, name string, attrs map[string]string) {
		cnt += 1
		if cnt > 1 {
			t.Errorf("ParseXPath unmatch. cnt=%d", cnt)
		}
		if ns != "ns1" || name != "node1" {
			t.Errorf("ParseXPath unmatch. node=%s:%s", ns, name)
		}
		if len(attrs) != 0 {
			t.Errorf("ParseXPath unmatch. attrs=%v", attrs)
		}
	})
}

func TestParseXPath2_ns(t *testing.T) {
	xpath := "/node1/ns2:node2"
	nss := []string{"", "ns2"}
	names := []string{"node1", "node2"}
	cnt := 0
	ParseXPath(xpath, func(ns, name string, attrs map[string]string) {
		cnt += 1
		if cnt > 2 {
			t.Errorf("ParseXPath unmatch. cnt=%d", cnt)
		}
		if ns != nss[cnt-1] || name != names[cnt-1] {
			t.Errorf("ParseXPath unmatch. node=%s:%s", ns, name)
		}
		if len(attrs) != 0 {
			t.Errorf("ParseXPath unmatch. attrs=%v", attrs)
		}
	})
}

func TestParseXPath2_attr1(t *testing.T) {
	xpath := "/node1[key1='val1']/node2[key2='val2']"
	names := []string{"node1", "node2"}
	var attrss = []map[string]string{
		{"key1": "val1"},
		{"key2": "val2"},
	}
	cnt := 0
	ParseXPath(xpath, func(ns, name string, attrs map[string]string) {
		cnt += 1
		if cnt > 2 {
			t.Errorf("ParseXPath unmatch. cnt=%d", cnt)
		}
		if ns != "" || name != names[cnt-1] {
			t.Errorf("ParseXPath unmatch. node=%s:%s", ns, name)
		}
		if ok := mapcmp(attrs, attrss[cnt-1]); !ok {
			t.Errorf("ParseXPath unmatch. attrs=%v", attrs)
		}
	})
}

func TestParseXPath2_attr1_ns(t *testing.T) {
	xpath := "/node1[key1='val1']/ns2:node2[key2='val2']"
	nss := []string{"", "ns2"}
	names := []string{"node1", "node2"}
	var attrss = []map[string]string{
		{"key1": "val1"},
		{"key2": "val2"},
	}
	cnt := 0
	ParseXPath(xpath, func(ns, name string, attrs map[string]string) {
		cnt += 1
		if cnt > 2 {
			t.Errorf("ParseXPath unmatch. cnt=%d", cnt)
		}
		if ns != nss[cnt-1] || name != names[cnt-1] {
			t.Errorf("ParseXPath unmatch. node=%s:%s", ns, name)
		}
		if ok := mapcmp(attrs, attrss[cnt-1]); !ok {
			t.Errorf("ParseXPath unmatch. attrs=%v", attrs)
		}
	})
}

func TestParseXPath2_attr2(t *testing.T) {
	xpath := "/node1[key1='val1'][key11='val11']/node2[key2='val2']"
	names := []string{"node1", "node2"}
	var attrss = []map[string]string{
		{"key1": "val1", "key11": "val11"},
		{"key2": "val2"},
	}
	cnt := 0
	ParseXPath(xpath, func(ns, name string, attrs map[string]string) {
		cnt += 1
		if cnt > 2 {
			t.Errorf("ParseXPath unmatch. cnt=%d", cnt)
		}
		if ns != "" || name != names[cnt-1] {
			t.Errorf("ParseXPath unmatch. node=%s:%s", ns, name)
		}
		if ok := mapcmp(attrs, attrss[cnt-1]); !ok {
			t.Errorf("ParseXPath unmatch. attrs=%v", attrs)
		}
	})
}
