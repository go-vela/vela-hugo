package main

import (
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
)

func TestConfig_Validate(t *testing.T) {

	tests := []struct {
		name    string
		setup   Config // the directories that should exist for the test
		config  Config // what flags are passed in the test
		wantErr bool
	}{
		{
			name:    "should pass - nothing set",
			setup:   Config{},
			config:  Config{},
			wantErr: false,
		},
		{
			name: "should pass - all values set and exist",
			setup: Config{
				CacheDirectory:   "/cache",
				ContentDirectory: "/content",
				Directory:        "/config",
				Environment:      "dev",
				File:             "config.toml",
				LayoutDirectory:  "/layout",
				OutputDirectory:  "/build",
				SourceDirectory:  "/source",
			},
			config: Config{
				CacheDirectory:   "/cache",
				ContentDirectory: "/content",
				Directory:        "/config",
				Environment:      "dev",
				File:             "config.toml",
				LayoutDirectory:  "/layout",
				OutputDirectory:  "/build",
				SourceDirectory:  "/source",
			},
			wantErr: false,
		},
		{
			name:  "should fail - cache dir doesn't exist",
			setup: Config{},
			config: Config{
				CacheDirectory: "/cache",
			},
			wantErr: true,
		},
		{
			name:  "should fail - content directory doesn't exist",
			setup: Config{},
			config: Config{
				ContentDirectory: "/content",
			},
			wantErr: true,
		},
		{
			name:  "should fail - config file doesn't exist",
			setup: Config{},
			config: Config{
				File: "config.toml",
			},
			wantErr: true,
		},
		{
			name:  "should fail - layout dir doesn't exist",
			setup: Config{},
			config: Config{
				LayoutDirectory: "/layout",
			},
			wantErr: true,
		},
		{
			name:  "should fail - source dir doesn't exist",
			setup: Config{},
			config: Config{
				SourceDirectory: "/source",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// setup in mem file system
			appFS = afero.NewMemMapFs()

			// prep that filesystem for the test with a set of directories that should already exist
			err := setupAppFS(tt.setup)
			if err != nil {
				t.Errorf("error %v while setting up file structure for test", err)
			}

			// run the validate with the test conditions and check if the reponse is expected
			if err := tt.config.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Config.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// build the folders/files requested
func setupAppFS(c Config) error {
	// if a cache directory was provided
	if len(c.CacheDirectory) > 0 {
		// create the cache directory
		err := appFS.MkdirAll(c.CacheDirectory, 0777)
		if err != nil {
			return err
		}
	}
	// if a content directory was provided
	if len(c.ContentDirectory) > 0 {
		// create the content directory
		err := appFS.MkdirAll(c.ContentDirectory, 0777)
		if err != nil {
			return err
		}
	}
	// if a config directory was provided
	if len(c.Directory) > 0 {
		// create the config directory
		err := appFS.MkdirAll(c.Directory, 0777)
		if err != nil {
			return err
		}
	}
	// if an environment was provided i.e. dev/prod
	if len(c.Environment) > 0 {
		// create the environment file in the config directory
		err := appFS.MkdirAll(filepath.Join(c.Directory, c.Environment), 0777)
		if err != nil {
			return err
		}
	}
	// if a config file was provided
	if len(c.File) > 0 {
		// create the config file
		_, err := appFS.Create(c.File)
		if err != nil {
			return err
		}
	}
	// if a layout directory was provided
	if len(c.LayoutDirectory) > 0 {
		// create the layout directory
		err := appFS.MkdirAll(c.LayoutDirectory, 0777)
		if err != nil {
			return err
		}
	}
	// if an output directory was provided
	if len(c.OutputDirectory) > 0 {
		// create the output directory
		err := appFS.MkdirAll(c.OutputDirectory, 0777)
		if err != nil {
			return err
		}
	}
	// if a source directory was provided
	if len(c.SourceDirectory) > 0 {
		// create the source directory
		err := appFS.MkdirAll(c.SourceDirectory, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}
