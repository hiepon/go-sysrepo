// -*- coding: utf-8 -*-

package main

import (
	"fmt"
	"strings"

	sr "github.com/hiepon/go-sysrepo/sysrepo"
)

func traverseTree(session *sr.SessionContext, node *sr.Node, depth int) {
	indent := strings.Repeat(" ", depth*2)
	fmt.Printf("%s+-- %s", indent, node)
	node.IterChildren(session, func(child *sr.Node) error {
		traverseTree(session, child, depth+1)
		return nil
	})
}

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

	xpath := "/ietf-interfaces:interfaces/interface[name='eth0']"
	err = sess.Subtree(xpath, sr.GET_SUBTREE_ITERATIVE, func(node *sr.Node) error {
		traverseTree(sess, node, 0)
		return nil
	})

	if err != nil {
		fmt.Printf("GetItem error. %s %s\n", xpath, err)
		return
	}
}
