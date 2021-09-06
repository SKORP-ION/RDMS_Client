package rest

import (
	. "RDMS_Client/logging"
	"RDMS_Client/structures"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"RDMS_Client/dbus"
)

func Initialize() (*structures.ResponseInit, error) {
	initStruct := structures.InitStruct{}
	err := initStruct.Initialize()
	if err != nil {
		Error.Fatal("Can't read data for initialization", err)
	}
	//initStruct.Name = <- dbus.InputName
	body := bytes.NewBuffer([]byte{})
	err = json.NewEncoder(body).Encode(&initStruct)
	if err != nil {
		Error.Fatal("Can't marshal json", err)
	}

	req, err := http.NewRequest("POST", addr + "/public/workstations/registerWS", body)

	if err != nil {
		Error.Fatal(err)
	}

	req.Header.Set("Contenct-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		dbus.Response = "Connection error"
		dbus.Status <- false
		Error.Fatal(err)
	}

	switch resp.StatusCode {
	case http.StatusOK: {
		data := &structures.ResponseInit{}
		err = json.NewDecoder(resp.Body).Decode(data)
		if err != nil {
			dbus.Response = "Error"
			dbus.Status <- false
			return data, err
		}
		dbus.Response = "OK"
		dbus.Status <- true
		return data, nil
	}
	default: {
		data := &structures.ResponseError{}
		err = json.NewDecoder(resp.Body).Decode(data)
		if err != nil {
			dbus.Response = "Error"
			dbus.Status <- false
			return &structures.ResponseInit{}, err
		}
		dbus.Response = data.Data
		dbus.Status <- false
		return &structures.ResponseInit{}, errors.New(data.Data)
	}
	}
}
