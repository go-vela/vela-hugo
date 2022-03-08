// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"os/exec"
	"testing"
)

func Test_execCmd(t *testing.T) {
	tests := []struct {
		name    string
		command *exec.Cmd
		wantErr bool
	}{
		{
			name:    "should pass - valid command",
			command: exec.Command("echo", "hello"),
			wantErr: false,
		},
		{
			name:    "should fail - invalid command",
			command: exec.Command("foobar", "world"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// run the command and validate the results
			if err := execCmd(tt.command); (err != nil) != tt.wantErr {
				t.Errorf("execCmd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
