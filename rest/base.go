package rest

import (
	. "RDMS_Client/logging"
	"RDMS_Client/structures"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"time"
)

const ContentType = "application/json"

var client *http.Client
var addr string
var WorkstationAgent structures.WorkstationAgent

func init() {
	err := godotenv.Load()

	if err != nil {
		Error.Fatal("Can't load .env file", err)
	}

	client = &http.Client{}
	client.Timeout = 10 * time.Second
	addr = os.Getenv("server_host") + ":" + os.Getenv("server_port")
}
