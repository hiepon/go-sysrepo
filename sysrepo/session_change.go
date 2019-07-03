// -*- coding: utf-8 -*-

package sysrepo

/*
#include <stdio.h>
#include <sysrepo.h>
#include "helper.h"
*/
import "C"

import (
	"fmt"
	"sync"
	"unsafe"
)

type SubscriptionContext C.sr_subscription_ctx_t

func (v *SubscriptionContext) C() *C.sr_subscription_ctx_t {
	return (*C.sr_subscription_ctx_t)(v)
}

//
// ModuleChangeCallback
//
type ModuleChangeCallback func(sess *SessionContext, module string, event NotifyEvent) int

var moduleChangeCallbacks = sync.Map{}

func moduleChangeCallbackRegister(module string, callback ModuleChangeCallback) error {
	if _, loaded := moduleChangeCallbacks.LoadOrStore(module, callback); loaded {
		return fmt.Errorf("Module %s already exists.", module)
	}
	return nil
}

func moduleChangeCallbackUnregister(module string) {
	moduleChangeCallbacks.Delete(module)
}

func moduleChangeCallbackGet(module string) (ModuleChangeCallback, bool) {
	if callback, ok := moduleChangeCallbacks.Load(module); ok {
		return callback.(ModuleChangeCallback), true
	}
	return nil, false
}

//export go_sr_module_change_cb
func go_sr_module_change_cb(c_sess *C.sr_session_ctx_t, c_module *C.char, c_event C.sr_notif_event_t, c_data unsafe.Pointer) int {
	module := C.GoString(c_module)

	if callback, ok := moduleChangeCallbackGet(module); ok {
		return callback((*SessionContext)(c_sess), module, NotifyEvent(c_event))
	}

	return int(ERR_INTERNAL)
}

func (v *SessionContext) ModuleChangeSubscribe(module string, priority uint32, flag SubscrFlag, cb ModuleChangeCallback) (*SubscriptionContext, error) {
	if err := moduleChangeCallbackRegister(module, cb); err != nil {
		return nil, err
	}

	c_module := C.CString(module)
	defer C.free(unsafe.Pointer(c_module))

	var ctx *C.sr_subscription_ctx_t
	rc := C.sr_module_change_subscribe(v.C(), c_module, C.sr_module_change_cb(C._sr_module_change_cb), nil, C.uint32_t(priority), C.uint(flag), &ctx)
	if err := ParseError(rc); err != nil {
		moduleChangeCallbackUnregister(module)
		return nil, err
	}

	return (*SubscriptionContext)(ctx), nil
}

func (v *SessionContext) ModuleChangeUnsubscribe(module string, ctx *SubscriptionContext) error {
	moduleChangeCallbackUnregister(module)
	rc := C.sr_unsubscribe(v.C(), ctx.C())
	return ParseError(rc)
}

//
// SubtreeChangeCallback
//
type SubtreeChangeCallback func(sess *SessionContext, xpath string, event NotifyEvent) int

var subtreeChangeCallbacks = sync.Map{}

func subtreeChangeCallbackRegister(xpath string, cb SubtreeChangeCallback) error {
	if _, loaded := subtreeChangeCallbacks.LoadOrStore(xpath, cb); loaded {
		return fmt.Errorf("XPath %s already exists.", xpath)
	}
	return nil

}

func subtreeChangeCallbackUnregister(xpath string) {
	subtreeChangeCallbacks.Delete(xpath)
}

func subtreeChangeCallbackGet(xpath string) (SubtreeChangeCallback, bool) {
	if cb, ok := subtreeChangeCallbacks.Load(xpath); ok {
		return cb.(SubtreeChangeCallback), true
	}
	return nil, false
}

//export go_sr_subtree_change_cb
func go_sr_subtree_change_cb(c_sess *C.sr_session_ctx_t, c_xpath *C.char, c_event C.sr_notif_event_t, c_data unsafe.Pointer) int {
	xpath := C.GoString(c_xpath)

	if cb, ok := subtreeChangeCallbackGet(xpath); ok {
		return cb((*SessionContext)(c_sess), xpath, NotifyEvent(c_event))
	}

	return int(ERR_INTERNAL)
}

func (v *SessionContext) SubtreeChangeSubscribe(xpath string, priority uint32, flag SubscrFlag, cb SubtreeChangeCallback) (*SubscriptionContext, error) {
	if err := subtreeChangeCallbackRegister(xpath, cb); err != nil {
		return nil, err
	}

	c_xpath := C.CString(xpath)
	defer C.free(unsafe.Pointer(c_xpath))

	var ctx *C.sr_subscription_ctx_t
	rc := C.sr_subtree_change_subscribe(v.C(), c_xpath, C.sr_subtree_change_cb(C._sr_subtree_change_cb), nil, C.uint32_t(priority), C.uint(flag), &ctx)
	if err := ParseError(rc); err != nil {
		subtreeChangeCallbackUnregister(xpath)
		return nil, err
	}

	return (*SubscriptionContext)(ctx), nil
}

func (v *SessionContext) SubtreeChangeUnsubscribe(xpath string, ctx *SubscriptionContext) error {
	subtreeChangeCallbackUnregister(xpath)
	rc := C.sr_unsubscribe(v.C(), ctx.C())
	return ParseError(rc)
}

//
// ModuleInstallCallback
//
type ModuleInstallCallback func(module string, rev string, state ModuleState)

var moduleInstallCallbacks ModuleInstallCallback
var moduleInstallCallbackMutex = sync.Mutex{}

func moduleInstallCallbackRegister(cb ModuleInstallCallback) error {
	moduleInstallCallbackMutex.Lock()
	defer moduleInstallCallbackMutex.Unlock()

	if moduleInstallCallbacks != nil {
		return fmt.Errorf("Module Install Cakkback already exist.")
	}

	moduleInstallCallbacks = cb
	return nil
}

func moduleInstallCallbackUnregister() {
	moduleInstallCallbackMutex.Lock()
	defer moduleInstallCallbackMutex.Unlock()

	moduleInstallCallbacks = nil
}

func moduleInstallCallbackGet() (ModuleInstallCallback, bool) {
	moduleInstallCallbackMutex.Lock()
	defer moduleInstallCallbackMutex.Unlock()

	if moduleInstallCallbacks == nil {
		return nil, false
	}

	return moduleInstallCallbacks, true
}

//export go_sr_module_install_cb
func go_sr_module_install_cb(c_module *C.char, c_rev *C.char, c_state C.sr_module_state_t, c_data unsafe.Pointer) {
	if cb, ok := moduleInstallCallbackGet(); ok {
		module := C.GoString(c_module)
		rev := C.GoString(c_rev)

		cb(module, rev, ModuleState(c_state))
	}
}

func (v *SessionContext) ModuleInstallSubscribe(flag SubscrFlag, cb ModuleInstallCallback) (*SubscriptionContext, error) {
	if err := moduleInstallCallbackRegister(cb); err != nil {
		return nil, err
	}

	var ctx *C.sr_subscription_ctx_t
	rc := C.sr_module_install_subscribe(v.C(), C.sr_module_install_cb(C._sr_module_install_cb), nil, C.uint(flag), &ctx)
	if err := ParseError(rc); err != nil {
		moduleInstallCallbackUnregister()
		return nil, err
	}

	return (*SubscriptionContext)(ctx), nil
}

func (v *SessionContext) ModuleInstallUnsubscribe(ctx *SubscriptionContext) error {
	moduleInstallCallbackUnregister()
	rc := C.sr_unsubscribe(v.C(), ctx.C())
	return ParseError(rc)
}

//
// FeatureEnableCallback
//
type FeatureEnableCallback func(module string, feature string, enabled bool)

var featureEnableCallbacks FeatureEnableCallback
var featureEnableCallbackMutex = sync.Mutex{}

func featureEnableCallbackRegister(cb FeatureEnableCallback) error {
	featureEnableCallbackMutex.Lock()
	defer featureEnableCallbackMutex.Unlock()

	if featureEnableCallbacks != nil {
		return fmt.Errorf("Feature Enable Cakkback already exists.")
	}
	featureEnableCallbacks = cb
	return nil
}

func featureEnableCallbackUnregister() {
	featureEnableCallbackMutex.Lock()
	defer featureEnableCallbackMutex.Unlock()

	featureEnableCallbacks = nil
}

func featureEnableCallbackGet(module string) (FeatureEnableCallback, bool) {
	featureEnableCallbackMutex.Lock()
	defer featureEnableCallbackMutex.Unlock()

	if featureEnableCallbacks == nil {
		return nil, false
	}

	return featureEnableCallbacks, true
}

//export go_sr_feature_enable_cb
func go_sr_feature_enable_cb(c_module *C.char, c_feature *C.char, c_enabled bool, c_data unsafe.Pointer) {
	module := C.GoString(c_module)
	feature := C.GoString(c_feature)

	if cb, ok := featureEnableCallbackGet(module); ok {
		cb(module, feature, bool(c_enabled))
	}
}

func (v *SessionContext) FeatureEnableSubscribe(flag SubscrFlag, cb FeatureEnableCallback) (*SubscriptionContext, error) {
	if err := featureEnableCallbackRegister(cb); err != nil {
		return nil, err
	}

	var ctx *C.sr_subscription_ctx_t
	rc := C.sr_feature_enable_subscribe(v.C(), C.sr_feature_enable_cb(C._sr_feature_enable_cb), nil, C.uint(flag), &ctx)
	if err := ParseError(rc); err != nil {
		featureEnableCallbackUnregister()
		return nil, err
	}

	return (*SubscriptionContext)(ctx), nil
}

func (v *SessionContext) FeatureEnableUnsubscribe(ctx *SubscriptionContext) error {
	featureEnableCallbackUnregister()
	rc := C.sr_unsubscribe(v.C(), ctx.C())
	return ParseError(rc)
}

func (v *SessionContext) ChangesRange(xpath string, cb func(ChangeOper, *Val, *Val) error) error {
	c_xpath := C.CString(xpath)
	defer C.free(unsafe.Pointer(c_xpath))

	var c_iter *C.sr_change_iter_t
	rc := C.sr_get_changes_iter(v.C(), c_xpath, &c_iter)
	if err := ParseError(rc); err != nil {
		return err
	}

	C.sr_free_change_iter(c_iter)

	for {
		var c_oper C.sr_change_oper_t
		var c_oldVal *C.sr_val_t
		var c_newVal *C.sr_val_t

		rc := C.sr_get_change_next(v.C(), c_iter, &c_oper, &c_oldVal, &c_newVal)
		if rc == C.SR_ERR_NOT_FOUND {
			return nil
		}
		if err := ParseError(rc); err != nil {
			return err
		}

		err := cb((ChangeOper)(c_oper), (*Val)(c_oldVal), (*Val)(c_newVal))

		if c_oldVal != nil {
			C.sr_free_val(c_oldVal)
		}
		if c_newVal != nil {
			C.sr_free_val(c_newVal)
		}

		if err != nil {
			return err
		}
	}
}
