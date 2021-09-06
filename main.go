package main

import (
	"RDMS_Client/dbus"
	"RDMS_Client/handler"
	. "RDMS_Client/logging"
	"RDMS_Client/rest"
	"RDMS_Client/utils"
	"github.com/joho/godotenv"
	"time"
)

const SleepTime = 15 * time.Minute

func main() {
	//TODO:Забить фиксированные адреса
	err := godotenv.Load()
	if err != nil {
		Error.Println("Can't load env. ", err)
	}

	go dbus.StartServer()

	//Проверка на инициализацию
	if status, err := utils.CheckInitStatus(); !status {
		Info.Println("Initialize on server. ")
		response, err := rest.Initialize()
		if err != nil {
			Error.Println("Can't initialize. ", err)
			return
		}
		err = utils.WritePersonalKey(response.Data.Personal_key)

		if err != nil {
			Error.Fatal("Can't Write key. ", err)
		}

		err = utils.ChangeInitStatus()

		if err != nil {
			Error.Fatal("Can't write init status. ", err)
		}

		Info.Println("Successfully registered on the server. ")
	} else if err != nil {
		Error.Fatal("Can't check init status. ", err)
	} else {
		Info.Println("Already registered. ")
	}


	for {
		wa := &rest.WorkstationAgent
		err := wa.Init()

		if err != nil {
			Error.Fatal("Can't init WorkstationAgent. ", err)
		}

		err = rest.SignIn()

		if err != nil {
			Error.Fatal("SignIn error. ", err)
		}

		Info.Println("Successfully logged in. ")

		packages, err := rest.GetPackagesList()

		if err != nil {
			Warning.Println("There was an error getting the package list. ", err)
		} else {
			Info.Println("Package list received successfully. ")
		}

		handler.HandlePackages(packages)
		Info.Println("Cycle is done. Wait 15 minutes. ")
		time.Sleep(SleepTime)
	}

}
