package rest

import (
	"RDMS_Client/structures"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func GetPackagesList() (*structures.PackagesList, error) {
	wa := WorkstationAgent
	data := &structures.PackagesList{}
	req, err := http.NewRequest("GET", addr + "/private/packages/getPackagesList", nil)

	if err != nil {
		return data, err
	}

	req.Header.Add("Content-Type", ContentType)
	req.Header.Add("Authorization", wa.Token)
	req.Header.Add("Workstation_name", wa.Name)

	resp, err := client.Do(req)

	if err != nil{
		return data, err
	}

	if resp.StatusCode == 200 {
		var response structures.ResponsePackages
		err = json.NewDecoder(resp.Body).Decode(&response)
		data = &response.Data
		return data, err
	} else {
		var response structures.ResponseError
		err = json.NewDecoder(resp.Body).Decode(&response)

		return data, errors.New(fmt.Sprintf("Status Code: %d. Data: %s.\nJsonEncode: %s", resp.StatusCode,
			response.Data, err))
	}
}

func CreateDownloadSession(pkg structures.Package) (session_key string, err error) {
	wa := WorkstationAgent
	body := bytes.NewBuffer([]byte{})

	err = json.NewEncoder(body).Encode(pkg.DownloadSessionData())

	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", addr + "/private/packages/getSessionKey", body)

	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", ContentType)
	req.Header.Add("Authorization", wa.Token)
	req.Header.Add("Workstation_name", wa.Name)

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusCreated {
		data := structures.ResponseError{}

		err := json.NewDecoder(resp.Body).Decode(&data)

		if err != nil {
			return "", err
		}

		message := fmt.Sprintf("Status code: %s;  Message: %s", data.Status, data.Data)

		return "", errors.New(message)
	}

	data := structures.ResponseSession{}

	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		return "", err
	}

	return data.Data.SessionKey, nil
}