// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

// Config represents the plugin configuration for Hugo information.
type Config struct {
	// filesystem path to cache directory
	CacheDirectory string
	// filesystem path to content directory
	ContentDirectory string
	// filesystem path to config directory
	Directory string
	// targeted build environment in config directory
	Environment string
	// name of file in config directory
	File string
	// filesystem path to layout directory
	LayoutDirectory string
	// filesystem path to write files to
	OutputDirectory string
	// filesystem path to read files from
	SourceDirectory string
}

// Validate verifies the Config is properly configured.
func (c *Config) Validate() error {
	logrus.Trace("validating config configuration")

	// use custom filesystem which enables us to test
	a := &afero.Afero{
		Fs: appFS,
	}

	// check if a cache directory is provided
	if len(c.CacheDirectory) > 0 {
		// check if cache directory exists
		_, err := a.Stat(c.CacheDirectory)
		if err != nil {
			// check if a not exist err was returned
			if os.IsNotExist(err) {
				return fmt.Errorf("no cache directory found @ %s", c.CacheDirectory)
			}

			return err
		}
	}

	// check if a config file is provided
	if len(c.File) > 0 {
		// verify config directory is provided
		if len(c.Directory) == 0 {
			return fmt.Errorf("no config directory provided")
		}

		// check if config directory exists
		_, err := a.Stat(c.Directory)
		if err != nil {
			// check if a not exist err was returned
			if os.IsNotExist(err) {
				return fmt.Errorf("no config directory found @ %s", c.Directory)
			}

			return err
		}

		// create path to config based off directory and name
		path := filepath.Join(c.Directory, c.File)

		// validate that the config file exists
		_, err = a.Stat(path)
		if err != nil {
			// check if a not exist err was returned
			if os.IsNotExist(err) {
				return fmt.Errorf("no config found @ %s", path)
			}

			return err
		}
	}

	// check if a content directory is provided
	if len(c.ContentDirectory) > 0 {
		// check if content directory exists
		_, err := a.Stat(c.ContentDirectory)
		if err != nil {
			// check if a not exist err was returned
			if os.IsNotExist(err) {
				return fmt.Errorf("no content directory found @ %s", c.ContentDirectory)
			}

			return err
		}
	}

	// check if a layout directory is provided
	if len(c.LayoutDirectory) > 0 {
		// check if layout directory exists
		_, err := a.Stat(c.LayoutDirectory)
		if err != nil {
			// check if a not exist err was returned
			if os.IsNotExist(err) {
				return fmt.Errorf("no layout directory found @ %s", c.LayoutDirectory)
			}

			return err
		}
	}

	// check if a source directory is provided
	if len(c.SourceDirectory) > 0 {
		// check if source directory exists
		_, err := a.Stat(c.SourceDirectory)
		if err != nil {
			// check if a not exist err was returned
			if os.IsNotExist(err) {
				return fmt.Errorf("no source directory found @ %s", c.SourceDirectory)
			}

			return err
		}
	}

	return nil
}
