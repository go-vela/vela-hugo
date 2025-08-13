// SPDX-License-Identifier: Apache-2.0

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
			command: exec.CommandContext(t.Context(), "echo", "hello"),
			wantErr: false,
		},
		{
			name:    "should fail - invalid command",
			command: exec.CommandContext(t.Context(), "foobar", "world"),
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
