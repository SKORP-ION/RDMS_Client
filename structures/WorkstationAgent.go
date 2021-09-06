package structures

import (
	"io/ioutil"
	"os"
)

type WorkstationAgent struct {
	Name string
	Personal_key string
	Token string
}

func (wa WorkstationAgent) Auth() map[string]interface{} {
	return map[string]interface{} {"name": wa.Name, "personal_key": wa.Personal_key}
}

func (wa *WorkstationAgent) Init() error {
	err := wa.readHostname()
	if err != nil {
		return err
	}
	err = wa.readPersonalKey()
	if err != nil {
		return err
	}
	return nil
}

func (wa *WorkstationAgent) readHostname() error {
	/*

	file, err := os.Open("/etc/hostname")

	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	wa.Name = string(data)
	return nil
	*/
	wa.Name = "DevelopWS"
	return nil

}

func (wa *WorkstationAgent) readPersonalKey() error {
	file, err := os.Open("/etc/personal_key")

	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	wa.Personal_key = string(data)
	return nil
}