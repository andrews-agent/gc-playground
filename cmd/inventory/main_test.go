package main

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestCommand(t *testing.T) {
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
			name: "stock",
			args: []string{"stock"},
			stdout: "BELT-1\tDrive belt\t31\tLisbon\n" +
				"BOLT-10\tBolt 10 mm\t120\tBerlin\n" +
				"GEAR-L\tLarge gear\t0\tTbilisi\n" +
				"GEAR-S\tSmall gear\t5\tLisbon\n" +
				"NUT-10\tNut 10 mm\t4\tBerlin\n\n",
		},
		{
			name:   "show lowercase SKU",
			args:   []string{"show", "gear-s"},
			stdout: "GEAR-S Small gear: 5 in Lisbon\n",
		},
		{
			name:     "show unknown SKU",
			args:     []string{"show", "MISSING"},
			stderr:   "no such item or warehouse\n",
			exitCode: 1,
		},
		{
			name:     "show missing SKU",
			args:     []string{"show"},
			stderr:   "usage: inventory show <sku>\n",
			exitCode: 2,
		},
		{
			name:     "no command",
			stderr:   "usage: inventory stock | low [threshold] | show <sku>\n",
			exitCode: 2,
		},
		{
			name:     "unknown command",
			args:     []string{"unknown"},
			stderr:   "unknown command \"unknown\"\n",
			exitCode: 2,
		},
		{
			name:   "default",
			args:   []string{"low"},
			stdout: "GEAR-L\t0\tTbilisi\nGEAR-S\t5\tLisbon\nNUT-10\t4\tBerlin\n\n",
		},
		{
			name:   "zero",
			args:   []string{"low", "0"},
			stdout: "GEAR-L\t0\tTbilisi\n\n",
		},
		{
			name:   "negative",
			args:   []string{"low", "-1"},
			stdout: "\n",
		},
		{
			name:     "invalid",
			args:     []string{"low", "invalid"},
			stderr:   "invalid threshold: invalid\n",
			exitCode: 2,
		},
		{
			name:     "overflow",
			args:     []string{"low", "999999999999999999999999999999"},
			stderr:   "invalid threshold: 999999999999999999999999999999\n",
			exitCode: 2,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(binary, tt.args...)
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
