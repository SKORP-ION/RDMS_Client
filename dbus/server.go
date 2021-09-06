package dbus

import (
	. "RDMS_Client/logging"
	"github.com/godbus/dbus/v5"
)

var Status chan bool = make(chan bool)
var Response string
var Conn *dbus.Conn

func StartServer() {
	var err error
	Conn, err = dbus.ConnectSystemBus()

	if err != nil {
		Error.Fatal(err)
	}

	err = ExportInterfaces()

	if err != nil {
		Error.Fatal(err)
	}

	select {}
}

