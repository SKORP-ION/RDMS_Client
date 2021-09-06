package utils

import "RDMS_Client/structures"

func GetMBSerial() (string, error) {
	mb := structures.Motherboard{}

	err := mb.GetSerial()
	if err != nil {
		return "", err
	}

	return mb.Serial, nil
}