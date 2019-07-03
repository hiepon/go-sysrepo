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

	if err = sess.SetItemStr(
		"/ietf-interfaces:interfaces/interface[name='gigaeth0']/type",
		"iana-if-type:ethernetCsmacd",
		sr.EDIT_DEFAULT,
	); err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	if err := sess.SetItemStr(
		"/ietf-interfaces:interfaces/interface[name='gigaeth0']/ietf-ip:ipv6/address[ip='fe80::ab8']/prefix-length",
		"128",
		sr.EDIT_DEFAULT,
	); err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	if err := sess.Commit(); err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}
