// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
)

func TestConfig_Validate(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		name    string
		config  Config
	}{
		{
			failure: false,
			name:    "no config paths provided",
			config: Config{
				CacheDirectory:   "",
				ContentDirectory: "",
				Directory:        "",
				Environment:      "",
				File:             "",
				LayoutDirectory:  "",
				OutputDirectory:  "",
				SourceDirectory:  "",
			},
		},
		{
			failure: false,
			name:    "all config paths provided",
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
		},
		{
			failure: true,
			name:    "nonexistent cache directory provided",
			config: Config{
				CacheDirectory:   "/cache",
				ContentDirectory: "",
				Directory:        "",
				Environment:      "",
				File:             "",
				LayoutDirectory:  "",
				OutputDirectory:  "",
				SourceDirectory:  "",
			},
		},
		{
			failure: true,
			name:    "nonexistent content directory provided",
			config: Config{
				CacheDirectory:   "",
				ContentDirectory: "/content",
				Directory:        "",
				Environment:      "",
				File:             "",
				LayoutDirectory:  "",
				OutputDirectory:  "",
				SourceDirectory:  "",
			},
		},
		{
			failure: true,
			name:    "config file with no config directory provided",
			config: Config{
				CacheDirectory:   "",
				ContentDirectory: "",
				Directory:        "",
				Environment:      "",
				File:             "config.toml",
				LayoutDirectory:  "",
				OutputDirectory:  "",
				SourceDirectory:  "",
			},
		},
		{
			failure: true,
			name:    "nonexistent config file provided",
			config: Config{
				CacheDirectory:   "",
				ContentDirectory: "",
				Directory:        "/config",
				Environment:      "",
				File:             "config.toml",
				LayoutDirectory:  "",
				OutputDirectory:  "",
				SourceDirectory:  "",
			},
		},
		{
			failure: true,
			name:    "nonexistent layout directory provided",
			config: Config{
				CacheDirectory:   "",
				ContentDirectory: "",
				Directory:        "",
				Environment:      "",
				File:             "",
				LayoutDirectory:  "/layout",
				OutputDirectory:  "",
				SourceDirectory:  "",
			},
		},
		{
			failure: true,
			name:    "nonexistent source directory provided",
			config: Config{
				CacheDirectory:   "",
				ContentDirectory: "",
				Directory:        "",
				Environment:      "",
				File:             "",
				LayoutDirectory:  "",
				OutputDirectory:  "",
				SourceDirectory:  "/source",
			},
		},
	}

	// run tests
	for _, test := range tests {
		// setup in mem file system
		appFS = afero.NewMemMapFs()

		// check if a cache directory was provided
		if len(test.config.CacheDirectory) > 0 {
			// check if the test is supposed to fail
			if !test.failure {
				// create the cache directory
				err := appFS.MkdirAll(test.config.CacheDirectory, 0777)
				if err != nil {
					t.Errorf("unable to create cache directory %s: %v", test.config.CacheDirectory, err)
				}
			}
		}

		// check if a content directory was provided
		if len(test.config.ContentDirectory) > 0 {
			// check if the test is supposed to fail
			if !test.failure {
				// create the content directory
				err := appFS.MkdirAll(test.config.ContentDirectory, 0777)
				if err != nil {
					t.Errorf("unable to create content directory %s: %v", test.config.ContentDirectory, err)
				}
			}
		}

		// check if a config directory was provided
		if len(test.config.Directory) > 0 {
			// check if the test is supposed to fail
			if !test.failure {
				// create the config directory
				err := appFS.MkdirAll(test.config.Directory, 0777)
				if err != nil {
					t.Errorf("unable to create config directory %s: %v", test.config.Directory, err)
				}
			}
		}

		// check if an environment was provided
		if len(test.config.Environment) > 0 {
			// check if the test is supposed to fail
			if !test.failure {
				// create full path to environment in config directory
				path := filepath.Join(test.config.Directory, test.config.Environment)

				// create the environment in the config directory
				err := appFS.MkdirAll(path, 0777)
				if err != nil {
					t.Errorf("unable to create environment %s: %v", path, err)
				}
			}
		}

		// check if a config file was provided
		if len(test.config.File) > 0 {
			// check if the test is supposed to fail
			if !test.failure {
				// create full path to config file in config directory
				path := filepath.Join(test.config.Directory, test.config.File)

				// create the config file
				_, err := appFS.Create(path)
				if err != nil {
					t.Errorf("unable to create config file %s: %v", path, err)
				}
			}
		}

		// check if a layout directory was provided
		if len(test.config.LayoutDirectory) > 0 {
			// check if the test is supposed to fail
			if !test.failure {
				// create the layout directory
				err := appFS.MkdirAll(test.config.LayoutDirectory, 0777)
				if err != nil {
					t.Errorf("unable to create output directory %s: %v", test.config.LayoutDirectory, err)
				}
			}
		}

		// check if an output directory was provided
		if len(test.config.OutputDirectory) > 0 {
			// check if the test is supposed to fail
			if !test.failure {
				// create the output directory
				err := appFS.MkdirAll(test.config.OutputDirectory, 0777)
				if err != nil {
					t.Errorf("unable to create output directory %s: %v", test.config.OutputDirectory, err)
				}
			}
		}

		// check if a source directory was provided
		if len(test.config.SourceDirectory) > 0 {
			// check if the test is supposed to fail
			if !test.failure {
				// create the source directory
				err := appFS.MkdirAll(test.config.SourceDirectory, 0777)
				if err != nil {
					t.Errorf("unable to create source directory %s: %v", test.config.SourceDirectory, err)
				}
			}
		}

		err := test.config.Validate()

		if test.failure {
			if err == nil {
				t.Errorf("%s Validate should have returned err", test.name)
			}

			continue
		}

		if err != nil {
			t.Errorf("%s Validate returned err: %v", test.name, err)
		}
	}
}
