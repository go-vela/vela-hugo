package main

import (
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
)

const ThemesDirectory = "themes"

func TestTheme_Validate(t *testing.T) {
	tests := []struct {
		name    string
		theme   Theme
		wantErr bool
	}{
		{
			name: "should pass - provided theme, default directory exists",
			theme: Theme{
				Name:      "docsy",
				Directory: "",
			},
			wantErr: false,
		},
		{
			name: "should pass - provided theme, custom directory exists",
			theme: Theme{
				Name:      "docsy",
				Directory: "themes",
			},
			wantErr: false,
		},
		{
			name: "should fail - provided theme, default directory doesn't exist",
			theme: Theme{
				Name:      "docsy",
				Directory: "",
			},
			wantErr: true,
		},
		{
			name: "should fail - provided theme, custom directory doesn't exist",
			theme: Theme{
				Name:      "docsy",
				Directory: "foo",
			},
			wantErr: true,
		},
		{
			name: "should pass - no themes",
			theme: Theme{
				Name:      "",
				Directory: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup in mem file system
			appFS = afero.NewMemMapFs()

			// setup the themes directory
			if len(tt.theme.Directory) > 0 {
				// create the specified directory
				err := appFS.MkdirAll(tt.theme.Directory, 0777)
				if err != nil {
					t.Errorf("error = %v, unable to create %s", err, tt.theme.Directory)
				}
			} else {
				// only create the default if the test should pass
				if !tt.wantErr {
					// create the default themes directory
					err := appFS.MkdirAll(ThemesDirectory, 0777)
					if err != nil {
						t.Errorf("error = %v, unable to create %s", err, ThemesDirectory)
					}
				}
			}

			// setup the theme within that directory if provided
			if len(tt.theme.Name) > 0 {
				// if they provided a themes directory
				if len(tt.theme.Directory) > 0 {
					// only create the file if the directory it's in should exist
					if !tt.wantErr {

						// create the theme in the provided directory
						_, err := appFS.Create(filepath.Join(tt.theme.Directory, tt.theme.Name))
						if err != nil {
							t.Errorf("error = %v, unable to create %s", err, filepath.Join(tt.theme.Directory, tt.theme.Name))
						}
					}
				} else {
					// only create the file if the directory it's in should exist
					if !tt.wantErr {

						// create the theme in the default directory
						_, err := appFS.Create(filepath.Join(ThemesDirectory, tt.theme.Name))
						if err != nil {
							t.Errorf("error = %v, unable to create %s", err, filepath.Join(ThemesDirectory, tt.theme.Name))
						}
					}
				}
			}

			// run the validate with the test conditions and check if the reponse is expected
			if err := tt.theme.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Theme.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
