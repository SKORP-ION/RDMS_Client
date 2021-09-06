package utils

import "io/ioutil"

func CheckInitStatus() (bool, error) {
	data, err := ReadFile("/etc/init_status") //TODO:Исправить путь перед продом

	if err != nil {
		return false, err
	}

	str := string(data)

	if str == "true" {
		return true, nil
	} else {
		return false, nil
	}
}

func ChangeInitStatus() (error) {
	var data []byte
	data = []byte("true")
	err := ioutil.WriteFile("/etc/init_status", data, 664) //TODO:Исправить путь перед продом

	if err != nil {
		return err
	}
	return nil
}