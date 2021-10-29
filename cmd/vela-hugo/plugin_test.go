// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"
	"os/exec"
	"reflect"
	"testing"
)

func TestPlugin_Command(t *testing.T) {
	// setup tests
	tests := []struct {
		name   string
		plugin Plugin
		want   *exec.Cmd
	}{
		{
			name: "full plugin object with all flags",
			plugin: Plugin{
				Build: &Build{
					BaseURL: "BaseURL",
					Draft:   true,
					Expired: true,
					Future:  true,
				},
				Config: &Config{
					CacheDirectory:   "/cache",
					ContentDirectory: "/content",
					Directory:        "/config",
					Environment:      "dev",
					File:             "config.toml",
					LayoutDirectory:  "/layout",
					OutputDirectory:  "/build",
					SourceDirectory:  "/source",
				},
				Theme: &Theme{
					Name:      "docsy",
					Directory: "themes",
				},
			},
			want: exec.Command(
				_hugo,
				fmt.Sprintf("--baseURL=%s", "BaseURL"),
				"--buildDrafts",
				"--buildExpired",
				"--buildFuture",
				fmt.Sprintf("--cacheDir=%s", "CacheDirectory"),
				fmt.Sprintf("--config=%s", "ConfigFile"),
				fmt.Sprintf("--configDir=%s", "ConfigDirectory"),
				fmt.Sprintf("--contentDir=%s", "ContentDirectory"),
				fmt.Sprintf("--environment=%s", "Environment"),
				fmt.Sprintf("--layoutDir=%s", "LayoutDirectory"),
				fmt.Sprintf("--destination=%s", "OutputDirectory"),
				fmt.Sprintf("--source=%s", "SourceDirectory"),
				fmt.Sprintf("--theme=%s", "Name1"),
				fmt.Sprintf("--themesDir=%s", "ThemeDirectory"),
			),
		},
	}

	// run tests
	for _, test := range tests {
		got := test.plugin.Command()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s Command is %v, want %v", test.name, got, test.want)
		}
	}
}

func TestPlugin_Exec(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		name    string
		plugin  Plugin
	}{
		{
			failure: true,
			name:    "empty plugin object provided",
			plugin: Plugin{
				Build:  &Build{},
				Config: &Config{},
				Theme:  &Theme{},
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.plugin.Exec()

		if test.failure {
			if err == nil {
				t.Errorf("%s Exec should have returned err", test.name)
			}

			continue
		}

		if err != nil {
			t.Errorf("%s Exec returned err: %v", test.name, err)
		}
	}
}
