// SPDX-License-Identifier: Apache-2.0

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
					BaseURL: "http://hugo.example.com/",
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
			//nolint:gosec // ignore for testing
			want: exec.Command(
				_hugo,
				fmt.Sprintf("--baseURL=%s", "http://hugo.example.com/"),
				"--buildDrafts",
				"--buildExpired",
				"--buildFuture",
				fmt.Sprintf("--cacheDir=%s", "/cache"),
				fmt.Sprintf("--config=%s", "config.toml"),
				fmt.Sprintf("--configDir=%s", "/config"),
				fmt.Sprintf("--contentDir=%s", "/content"),
				fmt.Sprintf("--environment=%s", "dev"),
				fmt.Sprintf("--layoutDir=%s", "/layout"),
				fmt.Sprintf("--destination=%s", "/build"),
				fmt.Sprintf("--source=%s", "/source"),
				fmt.Sprintf("--theme=%s", "docsy"),
				fmt.Sprintf("--themesDir=%s", "themes"),
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
