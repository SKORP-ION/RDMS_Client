package rest

import (
	"RDMS_Client/structures"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func SignIn() error {

	body := bytes.NewBuffer([]byte{})

	err := json.NewEncoder(body).Encode(WorkstationAgent.Auth())

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", addr + "/public/authorization", body)

	if err != nil {
		return err
	}

	response, err := client.Do(req)

	if err != nil {
		return err
	}

	if response.StatusCode == http.StatusOK {
		data := &structures.ResponseAuth{}
		err = json.NewDecoder(response.Body).Decode(data)

		if err != nil {
			return err
		}

		WorkstationAgent.Token = data.Data.Token
		return nil
	} else {
		data := &structures.ResponseError{}
		err := json.NewDecoder(response.Body).Decode(data)

		if err != nil {
			return err
		}

		return errors.New(data.Data)
	}
}
