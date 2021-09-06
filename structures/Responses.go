package structures

import "time"

type ResponseInit struct {
	Status string
	Data DataInit
}

type DataInit struct {
	Name string
	Serial string
	Personal_key string
}

type ResponseAuth struct {
	Status string
	Data Token
}

type Token struct {
	Name string
	Created time.Time
	Token string
}

type ResponsePackages struct {
	Status string
	Data PackagesList
}

type ResponseSession struct {
	Status string
	Data DownloadSession
}

type DownloadSession struct {
	Md5 string
	SessionKey string `json:"session_key"`
}

type ResponseError struct {
	Status string
	Data string
}