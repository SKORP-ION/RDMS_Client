package dbus

import (
	"RDMS_Client/structures"
	"RDMS_Client/utils"
	"errors"
	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
)

const intro = `
<node>
	<interface name="rtc.rdms.daemon">
		<method name="Put">
			<arg direction="in" type="s"/>
			<arg direction="out" type="b"/>
			<arg direction="out" type="s"/>
		</method>
        <method name="GetInitStatus">
			<arg direction="out" type="b"/>
		</method>
	</interface>` + introspect.IntrospectDataString  +`</node>`

var Ws = Workstation{}

type Workstation struct {
	Name string
	status bool
	response string
}

func (ws Workstation) Put(name string) (bool, string, *dbus.Error ){
	structures.InputName <- name
	st := <-Status
	return st, Response, nil
}

func (ws Workstation) GetInitStatus() (bool, *dbus.Error) {
	status, _ := utils.CheckInitStatus()
	return status, nil
}


func ExportInterfaces() error {

	err := Conn.Export(Ws, "/rtc/rdms/daemon", "rtc.rdms.daemon")

	if err != nil {
		return err
	}

	err = Conn.Export(Ws, "/rtc/rdms/daemon", "rtc.rdms.daemon")

	if err != nil {
		return err
	}


	err = Conn.Export(introspect.Introspectable(intro), "/rtc/rdms/daemon",
		"org.freedesktop.DBus.Introspectable")
	if err != nil {
		return err
	}

	reply, err := Conn.RequestName("rtc.rdms.daemon", dbus.NameFlagDoNotQueue)

	if err != nil {
		return err
	} else if reply != dbus.RequestNameReplyPrimaryOwner {
		return errors.New("Name already taken")
	}
	return nil
}
