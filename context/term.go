package internal

import (
	"bytes"
	"io"
	"os"
	"os/exec"
)

// ReadTerminalContent returns the current terminal (tty) content as a string.
func ReadTerminalContent() (string, error) {
	// Try to read the tty device directly
	tty, err := os.Open("/dev/tty")
	if err != nil {
		return "", err
	}
	defer tty.Close()

	// Use "script" to capture the screen content
	cmd := exec.Command("script", "-q", "/dev/null", "-c", "cat /dev/tty")
	cmd.Stdin = tty
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = io.Discard

	if err := cmd.Run(); err != nil {
		return "", err
	}
	return out.String(), nil
}

