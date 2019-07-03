// -*- coding: utf-8 -*-

package main

import (
	"fmt"

	sr "github.com/hiepon/go-sysrepo/sysrepo"
)

func main() {
	sr.LogStderr(sr.LogLevelDebug)

	conn, err := sr.Connect("test1", sr.CONN_DEFAULT)
	if err != nil {
		fmt.Printf("Connect error. %s\n", err)
		return
	}
	defer conn.Disconnect()

	fmt.Printf("Connect ok. %v\n", conn)

	sess, err := sr.SessionStart(conn, sr.DS_RUNNING, sr.SESS_DEFAULT)
	if err != nil {
		fmt.Printf("SessionStart error. %s\n", err)
		return
	}
	defer sess.Stop()

	fmt.Printf("SessuinStart ok.\n")

	modules := []string{"ietf-interfaces"}
	for _, module := range modules {
		fmt.Printf("=== module change subscr: %s ===\n", module)

		_, err := sess.ModuleChangeSubscribe(module, 0, sr.SUBSCR_DEFAULT|sr.SUBSCR_APPLY_ONLY, func(chgsess *sr.SessionContext, m string, ev sr.NotifyEvent) int {
			fmt.Printf("ModuleChanged: %s %s\n", m, ev)
			xpath := fmt.Sprintf("/%s:*", m)
			err := chgsess.ChangesRange(xpath, func(oper sr.ChangeOper, oldVal, newVal *sr.Val) error {
				fmt.Printf("%s %s %s\n", oper, oldVal, newVal)
				return nil
			})
			if err != nil {
				return int(sr.ERR_INTERNAL)
			}

			return int(sr.ERR_OK)
		})
		if err != nil {
			fmt.Printf("ModuleChangeSubscr  error. %s %s\n", module, err)
			return
		}
	}

	done := make(chan struct{})
	<-done
}
