package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"github.com/urfave/cli/v2"
)

const (
	maxLines   = 10
	logFile    = "/tmp/term_output.log"
	pidFile    = "/tmp/term_capture.pid"
)


func captureTerminalOutput(filename string) error {
	cmd := exec.Command("script", "-q", "-c", os.Getenv("SHELL"), filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := os.WriteFile(pidFile, []byte(fmt.Sprintf("%d", cmd.Process.Pid)), 0644)
	if err != nil {
		return err
	}
	fmt.Printf("Recording started. PID: %d\n", cmd.Process.Pid)
	return cmd.Run()
}

func stopCapture() error {
	data, err := os.ReadFile(pidFile)
	if err != nil {
		return fmt.Errorf("failed to read pid file: %w", err)
	}
	pidStr := strings.TrimSpace(string(data))
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return fmt.Errorf("invalid PID: %w", err)
	}

	// Send SIGTERM
	if err := syscall.Kill(pid, syscall.SIGTERM); err != nil {
		return fmt.Errorf("failed to stop process: %w", err)
	}

	_ = os.Remove(pidFile)
	fmt.Printf("Stopped terminal capture with PID %d\n", pid)
	return nil
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
		Subcommands: []*cli.Command{
		{
			Name:  "start",
			Usage: "Start capturing the terminal session",
			Action: func(c *cli.Context) error {
				fmt.Println("Captur
				
				ing terminal session...")
				if err := captureTerminalOutput(logFile); err != nil {
					return fmt.Errorf("capture failed: %w", err)
				}
				if err := trimFile(logFile, maxLines); err != nil {
					return fmt.Errorf("trim failed: %w", err)
				}
				fmt.Printf("Session saved to: %s\n", logFile)
				return nil
			},
		},
		{
			Name:  "stop",
			Usage: "Stop the terminal capture process",
			Action: func(c *cli.Context) error {
				return stopCapture()
			},
		},
	},

	
}
