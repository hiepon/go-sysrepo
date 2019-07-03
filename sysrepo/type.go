// -*- coding: utf-8 -*-

package sysrepo

/*
#include <sysrepo.h>
*/
import "C"

import (
	"fmt"
)

//
// sr_type_t
//
type Type C.sr_type_t

func (v Type) C() C.sr_type_t {
	return C.sr_type_t(v)
}

const (
	UNKNOWN_T            Type = C.SR_UNKNOWN_T
	TREE_ITERATOR_T      Type = C.SR_TREE_ITERATOR_T
	LIST_T               Type = C.SR_LIST_T
	CONTAINER_T          Type = C.SR_CONTAINER_T
	CONTAINER_PRESENCE_T Type = C.SR_CONTAINER_PRESENCE_T
	LEAF_EMPTY_T         Type = C.SR_LEAF_EMPTY_T
	BINARY_T             Type = C.SR_BINARY_T
	BITS_T               Type = C.SR_BITS_T
	BOOL_T               Type = C.SR_BOOL_T
	DECIMAL64_T          Type = C.SR_DECIMAL64_T
	ENUM_T               Type = C.SR_ENUM_T
	IDENTITYREF_T        Type = C.SR_IDENTITYREF_T
	INSTANCEID_T         Type = C.SR_INSTANCEID_T
	INT8_T               Type = C.SR_INT8_T
	INT16_T              Type = C.SR_INT16_T
	INT32_T              Type = C.SR_INT32_T
	INT64_T              Type = C.SR_INT64_T
	STRING_T             Type = C.SR_STRING_T
	UINT8_T              Type = C.SR_UINT8_T
	UINT16_T             Type = C.SR_UINT16_T
	UINT32_T             Type = C.SR_UINT32_T
	UINT64_T             Type = C.SR_UINT64_T
	ANYXML_T             Type = C.SR_ANYXML_T
	ANYDATA_T            Type = C.SR_ANYDATA_T
)

var srTypeValues = map[string]Type{
	"UNKNOWN_T":            UNKNOWN_T,
	"TREE_ITERATOR_T":      TREE_ITERATOR_T,
	"LIST_T":               LIST_T,
	"CONTAINER_T":          CONTAINER_T,
	"CONTAINER_PRESENCE_T": CONTAINER_PRESENCE_T,
	"LEAF_EMPTY_T":         LEAF_EMPTY_T,
	"BINARY_T":             BINARY_T,
	"BITS_T":               BITS_T,
	"BOOL_T":               BOOL_T,
	"DECIMAL64_T":          DECIMAL64_T,
	"ENUM_T":               ENUM_T,
	"IDENTITYREF_T":        IDENTITYREF_T,
	"INSTANCEID_T":         INSTANCEID_T,
	"INT8_T":               INT8_T,
	"INT16_T":              INT16_T,
	"INT32_T":              INT32_T,
	"INT64_T":              INT64_T,
	"STRING_T":             STRING_T,
	"UINT8_T":              UINT8_T,
	"UINT16_T":             UINT16_T,
	"UINT32_T":             UINT32_T,
	"UINT64_T":             UINT64_T,
	"ANYXML_T":             ANYXML_T,
	"ANYDATA_T":            ANYDATA_T,
}
var srTypeNames = map[Type]string{
	UNKNOWN_T:            "UNKNOWN_T",
	TREE_ITERATOR_T:      "TREE_ITERATOR_T",
	LIST_T:               "LIST_T",
	CONTAINER_T:          "CONTAINER_T",
	CONTAINER_PRESENCE_T: "CONTAINER_PRESENCE_T",
	LEAF_EMPTY_T:         "LEAF_EMPTY_T",
	BINARY_T:             "BINARY_T",
	BITS_T:               "BITS_T",
	BOOL_T:               "BOOL_T",
	DECIMAL64_T:          "DECIMAL64_T",
	ENUM_T:               "ENUM_T",
	IDENTITYREF_T:        "IDENTITYREF_T",
	INSTANCEID_T:         "INSTANCEID_T",
	INT8_T:               "INT8_T",
	INT16_T:              "INT16_T",
	INT32_T:              "INT32_T",
	INT64_T:              "INT64_T",
	STRING_T:             "STRING_T",
	UINT8_T:              "UINT8_T",
	UINT16_T:             "UINT16_T",
	UINT32_T:             "UINT32_T",
	UINT64_T:             "UINT64_T",
	ANYXML_T:             "ANYXML_T",
	ANYDATA_T:            "ANYDATA_T",
}

func (v Type) String() string {
	if s, ok := srTypeNames[v]; ok {
		return s
	}
	return fmt.Sprintf("Type(%d)", v)
}
