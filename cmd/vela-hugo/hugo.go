package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/hashicorp/go-getter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

const (
	_hugo     = "/bin/hugo"
	_hugoTmp  = "/bin/download"
	_download = "https://github.com/gohugoio/hugo/releases/download/v%s/%s_%s_%s-%s.tar.gz"
)

func install(extendedBinary bool, customVer, defaultVer string) error {
	// use custom filesystem which enables us to test
	a := &afero.Afero{
		Fs: appFS,
	}

	// setup vars for building the _download url
	//   based off of https://github.com/gohugoio/hugo/releases for the naming convention
	var (
		binary   string = "hugo"
		osName   string
		archType string
	)

	switch runtime.GOOS {
	case "darwin":
		osName = "macOS"
	case "linux":
		osName = "Linux"
	case "windows":
		osName = "Windows"
	default:
		osName = "unsupported"
	}

	switch runtime.GOARCH {
	case "amd64":
		archType = "64bit"
	case "arm64":
		archType = "arm64"
	case "arm":
		archType = "arm"
	case "386":
		archType = "32bit"
	default:
		archType = "unsupported"
	}

	versionMatch := strings.EqualFold(customVer, defaultVer)

	// check if the custom version matches the default version and/or the extneded binary is requested
	if versionMatch && !extendedBinary {
		// the hugo versions match and using the base hugo binary so no action is required
		return nil
	}

	if !versionMatch {
		logrus.Infof("custom version does not match default: %s", defaultVer)
	}

	if extendedBinary {
		logrus.Infof("using extended hugo binary")
		binary = "hugo_extended"
	}

	// rename the old hugo binary since we can't overwrite it for now
	//
	// https://github.com/hashicorp/go-getter/issues/219
	err := a.Rename(_hugo, fmt.Sprintf("%s.default", _hugo))
	if err != nil {
		return err
	}

	// create the download URL to install hugo - https://github.com/gohugoio/hugo/releases
	url := fmt.Sprintf(_download, customVer, binary, customVer, osName, archType)

	logrus.Infof("downloading hugo version from: %s", url)
	// send the HTTP request to install hugo
	err = getter.Get(_hugoTmp, url, []getter.ClientOption{}...)
	if err != nil {
		return err
	}

	// getter installed a directory of files, move the binary from that to the _hugo location
	err = a.Rename(_hugoTmp+"/hugo", _hugo)
	if err != nil {
		return err
	}

	logrus.Debugf("changing ownership of file: %s", _hugo)
	// ensure the hugo binary is executable
	err = a.Chmod(_hugo, 0700)
	if err != nil {
		return err
	}

	return nil
}
