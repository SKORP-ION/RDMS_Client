package structures

import "fmt"

var Actions = map[uint8]string {
	0: "skip",
	1: "install",
	2: "remove"}

type Package struct {
	Name 		string
	Version 	string
	Ord 		uint8
	OnServer 	bool
	Md5			string
	Action 		uint8
}

func (pkg Package) String() string {
	response := fmt.Sprintf("%s - %s. %s in order %d.", pkg.Name, pkg.Version, Actions[pkg.Action], pkg.Ord)

	if pkg.OnServer == true {
		response = fmt.Sprintf("%s Package is on server", response)
	} else {
		response = fmt.Sprintf("%s Packages is in repository", response)
	}
	return response
}

func (pkg Package) DownloadSessionData() map[string]string {
	return map[string]string {"name": pkg.Name, "version": pkg.Version}
}

type PackagesList []Package

func (pl PackagesList) String() string {
	response := ""
	for i := 0; i < len(pl); i++ {
		response = fmt.Sprintf("%s%s\n", response, pl[i].String())
	}
	return response
}