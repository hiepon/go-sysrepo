// -*- coding: utf-8 -*-

package sysrepo

/*
#include <sysrepo.h>
*/
import "C"

import (
	"fmt"
)

type LogLevel C.sr_log_level_t

func (v LogLevel) C() C.sr_log_level_t {
	return (C.sr_log_level_t)(v)
}

const (
	LogLevelNone  LogLevel = C.SR_LL_NONE
	LogLevelError LogLevel = C.SR_LL_ERR
	LogLevelWarn  LogLevel = C.SR_LL_WRN
	LogLevelInfo  LogLevel = C.SR_LL_INF
	LogLevelDebug LogLevel = C.SR_LL_DBG
)

var logLevel_names = map[LogLevel]string{
	LogLevelNone:  "None",
	LogLevelError: "Error",
	LogLevelWarn:  "Warn",
	LogLevelInfo:  "Info",
	LogLevelDebug: "Debug",
}

var logLevel_values = map[string]LogLevel{
	"None":  LogLevelNone,
	"Error": LogLevelError,
	"Warn":  LogLevelWarn,
	"Info":  LogLevelInfo,
	"Debug": LogLevelDebug,
}

func (v LogLevel) String() string {
	if s, ok := logLevel_names[v]; ok {
		return s
	}
	return fmt.Sprintf("LogLevel(%d)", v)
}

func ParseLogLevel(s string) (LogLevel, error) {
	if v, ok := logLevel_values[s]; ok {
		return v, nil
	}
	return LogLevelNone, fmt.Errorf("Invalid LogLevel. %s", s)
}

//
// API
//
func LogStderr(ll LogLevel) {
	C.sr_log_stderr(ll.C())
}

func LogSyslog(ll LogLevel) {
	C.sr_log_syslog(ll.C())
}
