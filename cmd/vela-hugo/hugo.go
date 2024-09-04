// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/hashicorp/go-getter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

const (
	_hugo     = "/bin/hugo"
	_hugoTmp  = "/bin/download"
	_download = "https://github.com/gohugoio/hugo/releases/download/v%s/%s_%s_%s-%s.tar.gz"
	_checksum = "https://github.com/gohugoio/hugo/releases/download/v%s/%s_%s_checksums.txt"
)

func install(extendedBinary bool, customVer, defaultVer string) error {
	// use custom filesystem which enables us to test
	a := &afero.Afero{
		Fs: appFS,
	}

	// setup vars for building the _download url
	//   based off of https://github.com/gohugoio/hugo/releases for the naming convention
	binary := "hugo"
	osName := runtime.GOOS
	archType := runtime.GOARCH

	// change the binary file name
	// if the extended version for Sass/SCSS support
	// has been requested
	if extendedBinary {
		logrus.Infof("using extended hugo binary")

		binary = "hugo_extended"
	}

	// use default version if no custom version
	// was requested
	if len(customVer) == 0 {
		customVer = defaultVer
	}

	// try to parse the version
	// into semantic version struct
	ver, err := semver.NewVersion(customVer)
	if err != nil {
		return fmt.Errorf("not a valid version: %s", customVer)
	}

	// get the version without leading "v",
	// if it was supplied
	verWithoutV := ver.String()

	// check if the custom version requested
	// is the default version
	isDefaultVersion := strings.EqualFold(verWithoutV, defaultVer)

	// are we using the included default
	// (non-extended) version?
	// if so, no need to download anything
	if isDefaultVersion && !extendedBinary {
		return nil
	}

	// let user know that a custom version
	// was requested
	if !isDefaultVersion {
		logrus.Infof("custom version requested (default is: %s): %s", defaultVer, verWithoutV)
	}

	// special handling for macOS.
	// starting with 0.102, hugo supplies
	// a "fat" universal binary
	//
	// see notes here: https://github.com/gohugoio/hugo/releases/tag/v0.102.0
	if osName == "darwin" && ver.Minor() > uint64(101) {
		archType = "universal"
	}

	// rename the old hugo binary since we can't overwrite it for now
	//
	// https://github.com/hashicorp/go-getter/issues/219
	err = a.Rename(_hugo, fmt.Sprintf("%s.default", _hugo))
	if err != nil {
		return err
	}

	// create the download URL to install hugo - https://github.com/gohugoio/hugo/releases
	url := fmt.Sprintf(_download, verWithoutV, binary, verWithoutV, osName, archType)
	checksumURL := fmt.Sprintf(_checksum, verWithoutV, binary, verWithoutV)
	fullURL := fmt.Sprintf("%s?checksum=file:%s", url, checksumURL)

	logrus.Infof("downloading hugo version from: %s", fullURL)
	// send the HTTP request to install hugo
	err = getter.Get(_hugoTmp, fullURL, []getter.ClientOption{}...)
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
