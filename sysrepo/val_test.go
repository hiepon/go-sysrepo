// -*- coding: utf-8 -*-

package sysrepo

import (
	"testing"
)

func TestVal_C(t *testing.T) {
	var err error

	val := Val{}

	if v, ok := val.StrData(); ok || v != "" {
		t.Errorf("Val.String unmatch. '%s'", v)
	}

	err = val.SetXPath("/interfaces/interface/eth0")
	if err != nil {
		t.Errorf("Val.SetXPath error. %s", err)
	}

	if v := val.XPath(); v != "/interfaces/interface/eth0" {
		t.Errorf("Val.String unmatch. '%s'", v)
	}

	err = val.SetStrData(STRING_T, "123")
	if err != nil {
		t.Errorf("Val.SetStrData error. %s", err)
	}
}
