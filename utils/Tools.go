package utils

import (
	"io/ioutil"
)

func ReadFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	return data, err
}
