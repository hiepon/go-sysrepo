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

	sess, err := sr.SessionStart(conn, sr.DS_STARTUP, sr.SESS_DEFAULT)
	if err != nil {
		fmt.Printf("SessionStart error. %s\n", err)
		return
	}
	defer sess.Stop()

	fmt.Printf("SessuinStart ok.\n")

	xpath := "/ietf-interfaces:interfaces/interface"
	err = sess.Items(xpath, func(val *sr.Val) error {
		fmt.Printf("%s", val)
		return nil
	})

	if err != nil {
		fmt.Printf("GetItem error. %s %s\n", xpath, err)
		return
	}
}
