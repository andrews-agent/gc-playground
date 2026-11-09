package main

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestLowThreshold(t *testing.T) {
	binary := filepath.Join(t.TempDir(), "inventory")
	if output, err := exec.Command("go", "build", "-o", binary, ".").CombinedOutput(); err != nil {
		t.Fatalf("build command: %v\n%s", err, output)
	}

	for _, tt := range []struct {
		name     string
		args     []string
		stdout   string
		stderr   string
		exitCode int
	}{
		{
			name:   "default",
			stdout: "GEAR-L\t0\tTbilisi\nGEAR-S\t5\tLisbon\nNUT-10\t4\tBerlin\n\n",
		},
		{
			name:   "zero",
			args:   []string{"0"},
			stdout: "GEAR-L\t0\tTbilisi\n\n",
		},
		{
			name:   "negative",
			args:   []string{"-1"},
			stdout: "\n",
		},
		{
			name:     "invalid",
			args:     []string{"invalid"},
			stderr:   "invalid threshold: invalid\n",
			exitCode: 2,
		},
		{
			name:     "overflow",
			args:     []string{"999999999999999999999999999999"},
			stderr:   "invalid threshold: 999999999999999999999999999999\n",
			exitCode: 2,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(binary, append([]string{"low"}, tt.args...)...)
			var stdout, stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr
			if err := cmd.Run(); err != nil {
				if _, ok := err.(*exec.ExitError); !ok {
					t.Fatalf("run command: %v", err)
				}
			}
			if got := cmd.ProcessState.ExitCode(); got != tt.exitCode {
				t.Errorf("exit code = %d, want %d", got, tt.exitCode)
			}
			if got := stdout.String(); got != tt.stdout {
				t.Errorf("stdout = %q, want %q", got, tt.stdout)
			}
			if got := stderr.String(); got != tt.stderr {
				t.Errorf("stderr = %q, want %q", got, tt.stderr)
			}
		})
	}
}
