package apt

import (
	"errors"
	"fmt"
	vers "github.com/hashicorp/go-version"
	"os"
	"os/exec"
	"strings"
)

const distr = "packages" //TODO:Подставить нормальное имя

func InstallFromRepo(name string, version string) (out string, err error) {
	var cmd *exec.Cmd

	if version == "" {
		cmd = exec.Command("apt", "install", name, "-y")
	} else {
		cmd = exec.Command("apt", "install", fmt.Sprintf("%s=%s", name, version))
	}

	byteOut, err := cmd.CombinedOutput()

	return string(byteOut), err
}

func InstallFromDeb(filename string) (out string, err error) {
	var cmd *exec.Cmd

	cmd = exec.Command("apt", "install", fmt.Sprintf("./%s/%s", distr, filename), "-y")

	byteOut, err := cmd.CombinedOutput()

	if err != nil {
		return string(byteOut), err
	}

	err = removeDistrib(filename)

	if err != nil {
		return "Can't delete distrib", err
	}

	return string(byteOut), err
}

func removeDistrib(filename string) error {
	return  os.Remove(fmt.Sprintf("%s/%s", distr, filename))
}

func RemovePackage(name string) error {
	cmd := exec.Command("apt", "remove", name, "-y")


	out, err := cmd.CombinedOutput()

	if err != nil {
		return errors.New(string(out))
	}
	return nil
}

func IsInstalled(name string, version string) (status bool, err error) {
	cmd := exec.Command("dpkg-query", "-W", "-f=${Status}", name)

	out, err := cmd.CombinedOutput()

	InstalledStatus := string(out)

	if err != nil && !strings.Contains(InstalledStatus, "не соответствует ни один пакет") {
		return false, err
	} else if strings.Contains(InstalledStatus, "не соответствует ни один пакет") {
		return false, nil
	}

	if version == "" && strings.Contains(InstalledStatus, "installed") {
		return true, nil
	}

	cmd = exec.Command("dpkg-query", "-W", "-f=${Version}", name)

	out, err = cmd.CombinedOutput()

	if err != nil {
		return false, err
	}
	str := string(out)
	InstalledVersion, err := vers.NewVersion(str)
	NewVersion, _ := vers.NewVersion(version)

	if strings.Contains(InstalledStatus, "installed") && NewVersion.LessThanOrEqual(InstalledVersion) {
		return true, nil
	}

	return false, nil
}