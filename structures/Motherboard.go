package structures

import (
	"bufio"
	"io/ioutil"
	"os"
)

type Motherboard struct {
	Name string
	Serial string
	Vendor string
}

func (mb *Motherboard) ReadAllInfo() error {
	if err := mb.getName(); err != nil {
		return err
	} else if err := mb.GetSerial(); err != nil {
		return err
	} else if err := mb.getVendor(); err != nil {
		return err
	}
	return nil
}

func (mb *Motherboard) getName() (error) {
	data, err := ioutil.ReadFile("/sys/devices/virtual/dmi/id/board_name")

	if err != nil {
		return err
	}

	mb.Name = string(data)
	return nil

}

func (mb *Motherboard) GetSerial() (error) {
	data, err := ioutil.ReadFile("/sys/devices/virtual/dmi/id/board_serial")

	if err != nil {
		return err
	}

	mb.Serial = string(data)
	return nil

}

func (mb *Motherboard) getVendor() (error) {
	f, err := os.Open("/sys/devices/virtual/dmi/id/product_vendor")
	defer f.Close()

	if err != nil {
		return err
	}

	r := bufio.NewReader(f)

	data := []byte{}

	_, err = r.Read(data)

	if err != nil {
		return err
	}

	mb.Vendor = string(data)

	return nil

}