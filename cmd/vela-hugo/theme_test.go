package main

import (
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
)

func TestTheme_Validate(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		name    string
		theme   Theme
	}{
		{
			failure: false,
			name:    "no theme or theme directory provided",
			theme: Theme{
				Name:      "",
				Directory: "",
			},
		},
		{
			failure: false,
			name:    "theme and theme directory provided",
			theme: Theme{
				Name:      "docsy",
				Directory: "themes",
			},
		},
		{
			failure: true,
			name:    "theme with no theme directory provided",
			theme: Theme{
				Name:      "docsy",
				Directory: "",
			},
		},
		{
			failure: true,
			name:    "theme with nonexistent theme directory provided",
			theme: Theme{
				Name:      "docsy",
				Directory: "foo",
			},
		},
	}

	// run tests
	for _, test := range tests {
		// setup in mem file system
		appFS = afero.NewMemMapFs()

		// check if a theme directory was provided
		if len(test.theme.Directory) > 0 {
			// check if the test is supposed to fail
			if !test.failure {
				// create the theme directory
				err := appFS.MkdirAll(test.theme.Directory, 0777)
				if err != nil {
					t.Errorf("unable to create theme directory %s: %v", test.theme.Directory, err)
				}
			}
		}

		// check if a theme name was provided
		if len(test.theme.Name) > 0 {
			// check if the test is supposed to fail
			if !test.failure {
				// create full path to theme in theme directory
				path := filepath.Join(test.theme.Directory, test.theme.Name)

				// create the theme in the provided directory
				_, err := appFS.Create(path)
				if err != nil {
					t.Errorf("unable to create theme %s: %v", path, err)
				}
			}
		}

		err := test.theme.Validate()

		if test.failure {
			if err == nil {
				t.Errorf("Validate should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Validate returned err: %v", err)
		}
	}
}
