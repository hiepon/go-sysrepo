// -*- coding: utf-8 -*-

package main

import (
	"flag"
	"fmt"

	sr "github.com/hiepon/go-sysrepo/sysrepo"
	log "github.com/sirupsen/logrus"
)

func dataProvider(session *sr.SessionContext, xpath string) error {
	_, err := session.DpGetItemsSubscribe(xpath, sr.SUBSCR_DEFAULT, func(getXPath string) ([]*sr.GoVal, sr.Error) {
		nodeName, _ := sr.XPathNodeName(getXPath)

		log.Debugf("DpGetItemsSubscribe Callback: %s (%s)", getXPath, nodeName)

		switch nodeName {
		case "interface":
			vals := sr.NewGoValArray(2)
			vals[0].InitStrData(
				fmt.Sprintf("%s[name='%s']/type", getXPath, "eth0"),
				"iana-if-type:ethernetCsmacd",
				sr.IDENTITYREF_T,
			)
			vals[1].InitStrData(
				fmt.Sprintf("%s[name='%s']/oper-status", getXPath, "eth0"),
				"down",
				sr.ENUM_T,
			)

			return vals, sr.ERR_OK

		case "statistics":
			vals := sr.NewGoValArray(1)
			vals[0].InitStrData(
				fmt.Sprintf("%s/%s", getXPath, "discontinuity-time"),
				"2016-10-06T15:12:50.52Z",
				sr.STRING_T,
			)

			return vals, sr.ERR_OK

		default:
			return sr.NewGoValArray(0), sr.ERR_OK
		}
	})

	if err != nil {
		log.Errorf("DpGetItemsSubscribe error. %s", err)
		return err
	}
	return nil
}

func dataRequester(session *sr.SessionContext, xpath string) error {
	err := session.ItemsRange(xpath, func(val *sr.Val) error {
		log.Debugf("IterItems: %v", val)
		return nil
	})

	if err != nil {
		log.Errorf("IterItems error. %s", err)
		return err
	}
	return nil
}

func main() {

	var provider bool
	flag.BoolVar(&provider, "provider", false, "run as provider")
	flag.Parse()

	// sr.LogStderr(sr.LogLevelDebug)
	log.SetLevel(log.DebugLevel)

	conn, err := sr.Connect("test1", sr.CONN_DEFAULT)
	if err != nil {
		fmt.Printf("Connect error. %s\n", err)
		return
	}
	defer conn.Disconnect()

	log.Debugf("Connect ok.")

	sess, err := sr.SessionStart(conn, sr.DS_RUNNING, sr.SESS_DEFAULT)
	if err != nil {
		log.Debugf("SessionStart error. %s", err)
		return
	}
	defer sess.Stop()

	log.Debugf("SessuinStart ok.")

	if provider {
		dataProvider(sess, "/ietf-interfaces:interfaces-state")
		done := make(chan struct{})
		<-done
	} else {
		dataRequester(sess, "/ietf-interfaces:interfaces-state/interface//*")
	}
}
