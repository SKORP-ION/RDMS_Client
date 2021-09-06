package handler

import (
	"RDMS_Client/apt"
	. "RDMS_Client/logging"
	"RDMS_Client/rest"
	"RDMS_Client/structures"
	"RDMS_Client/tcp"
	"fmt"
)

func HandlePackages(pl *structures.PackagesList) {
	for i := 0; i < len(*pl); i ++ {
		pkg := (*pl)[i]
		//Установка
		if pkg.Action == 1 {
			if status, err := apt.IsInstalled(pkg.Name, pkg.Version); err != nil {
				Error.Println(pkg.Name, "installing error:", err.Error())
				continue
			} else if status {
				Info.Println("Package", pkg.Name, "is already installed, no update required.")
				continue
			}
			if pkg.OnServer == false {
				out, err := apt.InstallFromRepo(pkg.Name, pkg.Version)
				if err != nil {
					Error.Println(out, err)
				} else {
					Info.Printf("Succkessfully installed %s\n", pkg.Name)
				}
			} else if pkg.OnServer == true {
				sessionKey, err := rest.CreateDownloadSession(pkg)

				if err != nil {
					Error.Println("Can't get session key for download", pkg.Name, "errror: ", err)
					continue
				}

				conn := tcp.Connection{}
				err = conn.Connect()
				if err != nil {
					Error.Println("Connect error", err)
					continue
				}

				err = conn.ReceivePackage(&pkg, sessionKey)

				if err != nil {
					Error.Println("Can't get or install .deb package", pkg.Name, "error: ", err)
					continue
				}
				Info.Printf("Successfully downloaded package %s from server", pkg.Name)

				filename := fmt.Sprintf("%s_%s.deb", pkg.Name, pkg.Version)

				out, err := apt.InstallFromDeb(filename)

				if err != nil {
					Error.Println("Can't install package", pkg.Name, "Error: ", err, out)
					continue
				}

				Info.Println("Successfully installed", pkg.Name, "from .deb package")
			}
		//Пропуск установки
		} else if pkg.Action == 0 {
			Info.Printf("%s package skipped\n", pkg.Name)
		//Удаление пакета
		} else if status, err := apt.IsInstalled(pkg.Name, pkg.Version); status && err != nil && pkg.Action == 2{
			err := apt.RemovePackage(pkg.Name)
			if err != nil {
				Error.Println(err)
			} else {
				Info.Println("Package", pkg.Name, "successfully removed")
			}
		}
	}
}