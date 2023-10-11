// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

// Theme represents the plugin configuration for what Hugo theme(s) to use.
type Theme struct {
	// name of theme to use from theme directory
	Name string
	// filesystem path to theme directory
	Directory string
}

// Validate verifies the Theme is properly configured.
func (t *Theme) Validate() error {
	logrus.Trace("validating theme configuration")

	// use custom filesystem which enables us to test
	a := &afero.Afero{
		Fs: appFS,
	}

	// check if a theme is provided
	if len(t.Name) > 0 {
		// verify theme directory is provided
		if len(t.Directory) == 0 {
			return fmt.Errorf("no theme directory provided")
		}

		// check if theme directory exists
		_, err := a.Stat(t.Directory)
		if err != nil {
			// check if a not exist err was returned
			if os.IsNotExist(err) {
				return fmt.Errorf("no theme directory found @ %s", t.Directory)
			}

			return err
		}

		// create path to theme based off directory and name
		path := filepath.Join(t.Directory, t.Name)

		// check if theme path exists
		_, err = a.Stat(path)
		if err != nil {
			// check if a not exist err was returned
			if os.IsNotExist(err) {
				return fmt.Errorf("no theme found @ %s", path)
			}

			return err
		}
	}

	return nil
}
