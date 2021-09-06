package utils

import (
	"crypto/md5"
	"errors"
	"fmt"
)

func CheckMd5 (path string, md5sum string) error {
	data, err := ReadFile(path)

	if err != nil {
		return err
	}

	sum := md5.Sum(data)

	str := fmt.Sprintf("%x", sum)

	if str != md5sum {
		return errors.New("Md5 sum don't match")
	}

	return nil
}
