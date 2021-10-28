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

func TestPlugin_Exec(t *testing.T) {
	tests := []struct {
		name    string
		plugin  Plugin
		wantErr bool
	}{
		{
			name: "should fail - p.Command returns error, no hugo binary found",
			plugin: Plugin{
				Build:  &Build{},
				Config: &Config{},
				Theme:  &Theme{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// run the exec with the test conditions and check if the reponse is expected
			if err := tt.plugin.Exec(); (err != nil) != tt.wantErr {
				t.Errorf("Plugin.Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPlugin_Command(t *testing.T) {
	tests := []struct {
		name   string
		plugin Plugin
		want   *exec.Cmd
	}{
		{
			name: "should pass - all flags correctly applied",
			plugin: Plugin{
				Build: &Build{
					BaseURL: "BaseURL",
					Draft:   true,
					Expired: true,
					Future:  true,
				},
				Config: &Config{
					CacheDirectory:   "CacheDirectory",
					ContentDirectory: "ContentDirectory",
					Directory:        "ConfigDirectory",
					Environment:      "Environment",
					File:             "ConfigFile",
					LayoutDirectory:  "LayoutDirectory",
					OutputDirectory:  "OutputDirectory",
					SourceDirectory:  "SourceDirectory",
				},
				Theme: &Theme{
					Name:      "Name1",
					Directory: "ThemeDirectory",
				}},
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// run the hugo plugin with the provided flags
			if got := tt.plugin.Command(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Plugin.Command() = %v, want %v", got, tt.want)
			}
		})
	}
}
