package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

const maxLines = 10

func captureTerminalOutput(filename string) error {
	cmd := exec.Command("script", "-q", "-c", os.Getenv("SHELL"), filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func trimFile(filename string, maxLines int) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) <= maxLines {
		return nil
	}

	trimmed := lines[len(lines)-maxLines:]

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, line := range trimmed {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

var CaptureCommand = &cli.Command{
	Name:  "capture",
	Usage: "Capture terminal session and trim the output",
	Action: func(c *cli.Context) error {
		file := "/tmp/term_output.log"
		fmt.Println("Capturing terminal session, everything you type will be recorded...")

		if err := captureTerminalOutput(file); err != nil {
			return fmt.Errorf("capture failed: %w", err)
		}

		if err := trimFile(file, maxLines); err != nil {
			return fmt.Errorf("trim failed: %w", err)
		}

		fmt.Printf("Terminal session saved to: %s (trimmed to last %d lines)\n", file, maxLines)
		return nil
	},
}
