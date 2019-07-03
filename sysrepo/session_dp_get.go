// -*- coding: utf-8 -*-

package sysrepo

/*
#include <stdio.h>
#include <sysrepo.h>
#include <sysrepo/values.h>
#include "helper.h"
*/
import "C"

import (
	"fmt"
	"strings"
	"sync"
	"unsafe"
)

type DpGetItemsCallback func(xpath string) ([]*GoVal, Error)

var dpGetItemsCallbacks = sync.Map{}

func dpGetItemsCallbackKey(xpath string) string {
	return fmt.Sprintf("%s/", xpath)
}

func dpGetItemsCallbackRegister(xpath string, callback DpGetItemsCallback) error {
	key := dpGetItemsCallbackKey(xpath)
	if _, loaded := dpGetItemsCallbacks.LoadOrStore(key, callback); loaded {
		return fmt.Errorf("DpGetItemsCallback(%s) already exists.", xpath)
	}
	return nil
}

func dpGetItemsCallbackUnregister(xpath string) {
	key := dpGetItemsCallbackKey(xpath)
	dpGetItemsCallbacks.Delete(key)
}

func dpGetItemsCallbackGet(xpath string) (DpGetItemsCallback, bool) {
	var callback DpGetItemsCallback
	xpathCmp := dpGetItemsCallbackKey(xpath)
	dpGetItemsCallbacks.Range(func(key, value interface{}) bool {
		if ok := strings.HasPrefix(xpathCmp, key.(string)); !ok {
			return true
		}

		callback = value.(DpGetItemsCallback)
		return false
	})

	return callback, (callback != nil)
}

//export go_sr_dp_get_items_cb
func go_sr_dp_get_items_cb(c_xpath *C.char, c_vals_pp **C.sr_val_t, c_cnt_p *C.size_t, c_userdata unsafe.Pointer) int {
	xpath := C.GoString(c_xpath)
	if callback, ok := dpGetItemsCallbackGet(xpath); ok {
		vals, rc := callback(xpath)
		if rc != ERR_OK {
			return int(rc)
		}

		vals_n := len(vals)
		if len(vals) == 0 {
			*c_vals_pp = nil
			*c_cnt_p = 0
			return int(ERR_OK)
		}

		var c_vals_p *C.sr_val_t
		if rc := C.sr_new_values(C.size_t(vals_n), &c_vals_p); rc != C.SR_ERR_OK {
			return int(rc)
		}

		for index := 0; index < vals_n; index++ {
			c_val_p := C._sr_val_array_get(c_vals_p, C.size_t(index))
			CopyVal((*Val)(c_val_p), vals[index])
		}

		*c_vals_pp = c_vals_p
		*c_cnt_p = C.size_t(vals_n)

		return int(ERR_OK)
	}

	return int(ERR_INTERNAL)
}

func (v *SessionContext) DpGetItemsSubscribe(xpath string, flag SubscrFlag, callback DpGetItemsCallback) (*SubscriptionContext, error) {
	if err := dpGetItemsCallbackRegister(xpath, callback); err != nil {
		return nil, err
	}

	c_xpath := C.CString(xpath)
	defer C.free(unsafe.Pointer(c_xpath))

	var ctx *C.sr_subscription_ctx_t
	rc := C.sr_dp_get_items_subscribe(v.C(), c_xpath, C.sr_dp_get_items_cb(C._sr_dp_get_items_cb), nil, C.uint(flag), &ctx)
	if err := ParseError(rc); err != nil {
		dpGetItemsCallbackUnregister(xpath)
		return nil, err
	}

	return (*SubscriptionContext)(ctx), nil
}

func (v *SessionContext) DpGetItemsUnsubscribe(xpath string, ctx *SubscriptionContext) error {
	dpGetItemsCallbackUnregister(xpath)
	rc := C.sr_unsubscribe(v.C(), ctx.C())
	return ParseError(rc)
}
