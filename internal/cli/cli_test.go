package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	var stdout, stderr bytes.Buffer
	if code := Run(nil, &stdout, &stderr); code != 0 {
		t.Fatalf("help exit code = %d, want 0", code)
	}
	help := stdout.String()
	for _, command := range commands {
		if !strings.Contains(help, command) {
			t.Errorf("help missing %q", command)
		}
	}
	if stderr.Len() != 0 {
		t.Fatalf("help wrote to stderr: %q", stderr.String())
	}

	for _, command := range commands {
		stdout.Reset()
		stderr.Reset()
		code := Run([]string{command}, &stdout, &stderr)
		if code == 0 {
			t.Fatalf("%s exit code = 0, want non-zero", command)
		}
		if strings.TrimSpace(stdout.String()) != NotImplemented {
			t.Fatalf("%s stdout = %q, want %q", command, stdout.String(), NotImplemented)
		}
	}
}
