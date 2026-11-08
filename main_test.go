package main

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestCommand(t *testing.T) {
	binary := filepath.Join(t.TempDir(), "gc-playground")
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
			name:   "numbers",
			args:   []string{"1", "2", "3"},
			stdout: "sum 6\nmean 2\n",
		},
		{
			name:   "negative and fractional numbers",
			args:   []string{"-2.5", "1.5"},
			stdout: "sum -1\nmean -0.5\n",
		},
		{
			name:   "no arguments",
			stdout: "sum 0\nmean 0\n",
		},
		{
			name:     "invalid argument after a number",
			args:     []string{"1", "invalid", "3"},
			stderr:   "not a number: invalid\n",
			exitCode: 1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(binary, tt.args...)
			var stdout, stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr
			err := cmd.Run()
			if err != nil {
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
