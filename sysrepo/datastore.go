// -*- coding: utf-8 -*-

package sysrepo

/*
#include <sysrepo.h>
#include "helper.h"
*/
import "C"

import (
	"fmt"
)

//
// sr_datastore_t
//
type Datastore C.sr_datastore_t

func (v Datastore) C() C.sr_datastore_t {
	return (C.sr_datastore_t)(v)
}

const (
	DS_STARTUP   Datastore = C.SR_DS_STARTUP
	DS_RUNNING   Datastore = C.SR_DS_RUNNING
	DS_CANDIDATE Datastore = C.SR_DS_CANDIDATE
)

var srDatastore_names = map[Datastore]string{
	DS_STARTUP:   "STARTUP",
	DS_RUNNING:   "RUNNING",
	DS_CANDIDATE: "CANDIDATE",
}

var srDatastore_values = map[string]Datastore{
	"STARTUP":   DS_STARTUP,
	"RUNNING":   DS_RUNNING,
	"CANDIDATE": DS_CANDIDATE,
}

func (v Datastore) String() string {
	if s, ok := srDatastore_names[v]; ok {
		return s
	}
	return fmt.Sprintf("Datastore(%d)", v)
}

func ParseDatastore(s string) (Datastore, error) {
	if v, ok := srDatastore_values[s]; ok {
		return v, nil
	}
	return DS_STARTUP, fmt.Errorf("Invalid Datastore. %s", s)
}
