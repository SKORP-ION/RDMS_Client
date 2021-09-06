package main

import "RDMS_Client/dbus"

func main() {
	go dbus.StartServer()
	<- dbus.Status
}
