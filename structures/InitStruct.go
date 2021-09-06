package structures

import (
	. "RDMS_Client/logging"
)

var InputName chan string = make(chan string)

type InitStruct struct {
	Name string `json:"name"`
	Serial string `json:"serial"`
}

func (is *InitStruct) Initialize() error {
	serial, err := is.getMBSerial()
	if err != nil{
		return err
	}
	is.Serial = serial
	Info.Println("Wait for a name from dbus")
	is.Name = <- InputName //InputName записывается по dbus
	Info.Println("Name", is.Name, "received")
	return nil
}

func (is *InitStruct) getMBSerial() (string, error) {
	mb := Motherboard{}

	err := mb.GetSerial()
	if err != nil {
		return "", err
	}

	return mb.Serial, nil
}