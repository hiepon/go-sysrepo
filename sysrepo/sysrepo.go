// -*- coding: utf-8 -*-

package sysrepo

/*
#cgo pkg-config: libsysrepo
#include <sysrepo.h>
#include "helper.h"
*/
import "C"

import (
	"fmt"
)

//
// sr_session_flag_t
//
type SessionFlag C.sr_session_flag_t

func (v SessionFlag) C() C.sr_session_flag_t {
	return (C.sr_session_flag_t)(v)
}

const (
	SESS_DEFAULT      SessionFlag = C.SR_SESS_DEFAULT
	SESS_CONFIG_ONLY  SessionFlag = C.SR_SESS_CONFIG_ONLY
	SESS_ENABLE_NACM  SessionFlag = C.SR_SESS_ENABLE_NACM
	SESS_MUTABLE_OPTS SessionFlag = C.SR_SESS_MUTABLE_OPTS
)

var srSessionFlag_names = map[SessionFlag]string{
	SESS_DEFAULT:      "DEFAULT",
	SESS_CONFIG_ONLY:  "CONFIG_ONLY",
	SESS_ENABLE_NACM:  "ENABLE_NACM",
	SESS_MUTABLE_OPTS: "MUTABLE_OPTS",
}

var srSessionFlag_values = map[string]SessionFlag{
	"DEFAULT":      SESS_DEFAULT,
	"CONFIG_ONLY":  SESS_CONFIG_ONLY,
	"ENABLE_NACM":  SESS_ENABLE_NACM,
	"MUTABLE_OPTS": SESS_MUTABLE_OPTS,
}

func (v SessionFlag) String() string {
	if s, ok := srSessionFlag_names[v]; ok {
		return s
	}
	return fmt.Sprintf("SessionFlag(%d)", v)
}

func ParseSessionFlag(s string) (SessionFlag, error) {
	if v, ok := srSessionFlag_values[s]; ok {
		return v, nil
	}
	return SESS_DEFAULT, fmt.Errorf("Invalid SessionFlag. %s", s)
}

//
// sr_sess_options_t
//
type SessionOptions C.sr_sess_options_t

func (v SessionOptions) C() C.sr_sess_options_t {
	return (C.sr_sess_options_t)(v)
}

//
// sr_get_subtree_flag_t
//
type GetSubtreeFlag C.sr_get_subtree_flag_t

func (v GetSubtreeFlag) C() C.sr_get_subtree_flag_t {
	return (C.sr_get_subtree_flag_t)(v)
}

const (
	GET_SUBTREE_DEFAULT   GetSubtreeFlag = C.SR_GET_SUBTREE_DEFAULT
	GET_SUBTREE_ITERATIVE GetSubtreeFlag = C.SR_GET_SUBTREE_ITERATIVE
)

var getSubtreeFlag_names = map[GetSubtreeFlag]string{
	GET_SUBTREE_DEFAULT:   "default",
	GET_SUBTREE_ITERATIVE: "iterative",
}

var getSubtreeFlag_values = map[string]GetSubtreeFlag{
	"default":   GET_SUBTREE_DEFAULT,
	"iterative": GET_SUBTREE_ITERATIVE,
}

func (v GetSubtreeFlag) String() string {
	if s, ok := getSubtreeFlag_names[v]; ok {
		return s
	}
	return fmt.Sprintf("GetSubtreeFlag(%d)", v)
}

func ParseGetSubtreeFlag(s string) (GetSubtreeFlag, error) {
	if v, ok := getSubtreeFlag_values[s]; ok {
		return v, nil
	}
	return GET_SUBTREE_DEFAULT, fmt.Errorf("Invalid GetSubtreeFlag. %s", s)
}

//
// sr_edit_flag_t
//
type EditFlag C.sr_edit_flag_t

func (v EditFlag) C() C.sr_edit_flag_t {
	return (C.sr_edit_flag_t)(v)
}

const (
	EDIT_DEFAULT       EditFlag = C.SR_EDIT_DEFAULT
	EDIT_NON_RECURSIVE EditFlag = C.SR_EDIT_NON_RECURSIVE
	EDIT_STRICT        EditFlag = C.SR_EDIT_STRICT
)

var editFlag_names = map[EditFlag]string{
	EDIT_DEFAULT:       "default",
	EDIT_NON_RECURSIVE: "non-recursive",
	EDIT_STRICT:        "strict",
}

var editFlag_values = map[string]EditFlag{
	"default":       EDIT_DEFAULT,
	"non-recursive": EDIT_NON_RECURSIVE,
	"strict":        EDIT_STRICT,
}

func (v EditFlag) String() string {
	if s, ok := editFlag_names[v]; ok {
		return s
	}
	return fmt.Sprintf("EditFlag(%d)", v)
}

func ParseEditFlag(s string) (EditFlag, error) {
	if v, ok := editFlag_values[s]; ok {
		return v, nil
	}
	return EDIT_DEFAULT, fmt.Errorf("Invalid EditFlag. %s", s)
}

//
// sr_move_position_t
//
type MovePosition C.sr_move_position_t

func (v MovePosition) C() C.sr_move_position_t {
	return (C.sr_move_position_t)(v)
}

const (
	MOVE_BEFORE MovePosition = C.SR_MOVE_BEFORE
	MOVE_AFTER  MovePosition = C.SR_MOVE_AFTER
	MOVE_FIRST  MovePosition = C.SR_MOVE_FIRST
	MOVE_LAST   MovePosition = C.SR_MOVE_LAST
)

var movePosition_names = map[MovePosition]string{
	MOVE_BEFORE: "before",
	MOVE_AFTER:  "after",
	MOVE_FIRST:  "first",
	MOVE_LAST:   "last",
}

var movePosition_values = map[string]MovePosition{
	"before": MOVE_BEFORE,
	"after":  MOVE_AFTER,
	"first":  MOVE_FIRST,
	"last":   MOVE_LAST,
}

func (v MovePosition) String() string {
	if s, ok := movePosition_names[v]; ok {
		return s
	}
	return fmt.Sprintf("MovePosition(%d)", v)
}

func ParseMovePosition(s string) (MovePosition, error) {
	if v, ok := movePosition_values[s]; ok {
		return v, nil
	}
	return MOVE_BEFORE, fmt.Errorf("Invalid MovePosition. %s", s)
}

//
// sr_mem_ctx_t
//
type MemContext C.sr_mem_ctx_t

func (v MemContext) C() C.sr_mem_ctx_t {
	return C.sr_mem_ctx_t(v)
}

//
// sr_subscr_flag_t
//
type SubscrFlag C.sr_subscr_flag_t

func (v SubscrFlag) C() C.sr_subscr_flag_t {
	return (C.sr_subscr_flag_t)(v)
}

const (
	SUBSCR_DEFAULT                  SubscrFlag = C.SR_SUBSCR_DEFAULT
	SUBSCR_CTX_REUSE                SubscrFlag = C.SR_SUBSCR_CTX_REUSE
	SUBSCR_PASSIVE                  SubscrFlag = C.SR_SUBSCR_PASSIVE
	SUBSCR_APPLY_ONLY               SubscrFlag = C.SR_SUBSCR_APPLY_ONLY
	SUBSCR_EV_ENABLED               SubscrFlag = C.SR_SUBSCR_EV_ENABLED
	SUBSCR_NO_ABORT_FOR_REFUSED_CFG SubscrFlag = C.SR_SUBSCR_NO_ABORT_FOR_REFUSED_CFG
	SUBSCR_NOTIF_REPLAY_FIRST       SubscrFlag = C.SR_SUBSCR_NOTIF_REPLAY_FIRST
)

var subscrFlag_names = map[SubscrFlag]string{
	SUBSCR_DEFAULT:                  "default",
	SUBSCR_CTX_REUSE:                "ctx-reuse",
	SUBSCR_PASSIVE:                  "passive",
	SUBSCR_APPLY_ONLY:               "apply-only",
	SUBSCR_EV_ENABLED:               "ev-enable",
	SUBSCR_NO_ABORT_FOR_REFUSED_CFG: "no-abort-for-refused-cfg",
	SUBSCR_NOTIF_REPLAY_FIRST:       "notify-replay-first",
}

var subscrFlag_values = map[string]SubscrFlag{
	"default":                  SUBSCR_DEFAULT,
	"ctx-reuse":                SUBSCR_CTX_REUSE,
	"passive":                  SUBSCR_PASSIVE,
	"apply-only":               SUBSCR_APPLY_ONLY,
	"ev-enable":                SUBSCR_EV_ENABLED,
	"no-abort-for-refused-cfg": SUBSCR_NO_ABORT_FOR_REFUSED_CFG,
	"notify-replay-first":      SUBSCR_NOTIF_REPLAY_FIRST,
}

func (v SubscrFlag) String() string {
	if s, ok := subscrFlag_names[v]; ok {
		return s
	}
	return fmt.Sprintf("SubscrFlag(%d)", v)
}

func ParseSubscrFlag(s string) (SubscrFlag, error) {
	if v, ok := subscrFlag_values[s]; ok {
		return v, nil
	}
	return SUBSCR_DEFAULT, fmt.Errorf("Invalid SubscrFlag. %s", s)
}

//
// sr_notif_event_t
//
type NotifyEvent C.sr_notif_event_t

func (v NotifyEvent) C() C.sr_notif_event_t {
	return (C.sr_notif_event_t)(v)
}

const (
	EV_VERIFY  NotifyEvent = C.SR_EV_VERIFY
	EV_APPLY   NotifyEvent = C.SR_EV_APPLY
	EV_ABORT   NotifyEvent = C.SR_EV_ABORT
	EV_ENABLED NotifyEvent = C.SR_EV_ENABLED
)

var notifyEvent_names = map[NotifyEvent]string{
	EV_VERIFY:  "verify",
	EV_APPLY:   "apply",
	EV_ABORT:   "abort",
	EV_ENABLED: "enabled",
}

var notifyEvent_values = map[string]NotifyEvent{
	"verify":  EV_VERIFY,
	"apply":   EV_APPLY,
	"abort":   EV_ABORT,
	"enabled": EV_ENABLED,
}

func (v NotifyEvent) String() string {
	if s, ok := notifyEvent_names[v]; ok {
		return s
	}
	return fmt.Sprintf("NotifyEvent(%d)", v)
}

func ParseNotifyEvent(s string) (NotifyEvent, error) {
	if v, ok := notifyEvent_values[s]; ok {
		return v, nil
	}
	return EV_VERIFY, fmt.Errorf("Invalid NotifyEvent. %s", s)
}

//
// sr_change_oper_t
//
type ChangeOper C.sr_change_oper_t

func (v ChangeOper) C() C.sr_change_oper_t {
	return (C.sr_change_oper_t)(v)
}

const (
	OP_CREATED  ChangeOper = C.SR_OP_CREATED
	OP_MODIFIED ChangeOper = C.SR_OP_MODIFIED
	OP_DELETED  ChangeOper = C.SR_OP_DELETED
	OP_MOVED    ChangeOper = C.SR_OP_MOVED
)

var changeOper_names = map[ChangeOper]string{
	OP_CREATED:  "created",
	OP_MODIFIED: "modified",
	OP_DELETED:  "deleted",
	OP_MOVED:    "moved",
}

var changeOper_values = map[string]ChangeOper{
	"created":  OP_CREATED,
	"modified": OP_MODIFIED,
	"deleted":  OP_DELETED,
	"moved":    OP_MOVED,
}

func (v ChangeOper) String() string {
	if s, ok := changeOper_names[v]; ok {
		return s
	}
	return fmt.Sprintf("ChangeOper(%d)", v)
}

func ParseChangeOper(s string) (ChangeOper, error) {
	if v, ok := changeOper_values[s]; ok {
		return v, nil
	}
	return OP_CREATED, fmt.Errorf("Invalid ChangeOper. %s", s)
}

//
// sr_module_state_t
//
type ModuleState C.sr_module_state_t

func (v ModuleState) C() C.sr_module_state_t {
	return (C.sr_module_state_t)(v)
}

const (
	MS_UNINSTALLED ModuleState = C.SR_MS_UNINSTALLED
	MS_IMPORTED    ModuleState = C.SR_MS_IMPORTED
	MS_IMPLEMENTED ModuleState = C.SR_MS_IMPLEMENTED
)

var moduleState_names = map[ModuleState]string{
	MS_UNINSTALLED: "uninstalled",
	MS_IMPORTED:    "imported",
	MS_IMPLEMENTED: "implemented",
}

var moduleState_values = map[string]ModuleState{
	"uninstalled": MS_UNINSTALLED,
	"imported":    MS_IMPORTED,
	"implemented": MS_IMPLEMENTED,
}

func (v ModuleState) String() string {
	if s, ok := moduleState_names[v]; ok {
		return s
	}
	return fmt.Sprintf("ModuleState(%d)", v)
}

func ParseModuleState(s string) (ModuleState, error) {
	if v, ok := moduleState_values[s]; ok {
		return v, nil
	}
	return MS_UNINSTALLED, fmt.Errorf("Invalid ModuleState. %s", s)
}
