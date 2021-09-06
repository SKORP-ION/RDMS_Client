package utils

import (
	"io/ioutil"
)

func WritePersonalKey(key string) error {
	err := ioutil.WriteFile("/etc/personal_key", []byte(key), 744) //TODO:Перед продом выставить правильный путь
	return err
}
