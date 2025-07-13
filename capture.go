package main

import (
    "fmt"
    "os"
    "os/exec"
)

func captureTerminalOutput(filename string) error {
    cmd := exec.Command("script", "-q", "-c", os.Getenv("SHELL"), filename)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Stdin = os.Stdin
    return cmd.Run()
}

func main() {
    file := "/tmp/term_output.log"
    fmt.Println("Capturing terminal session, everything you type will be recorded...")
    err := captureTerminalOutput(file)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Terminal session saved to: %s\n", file)
    }
}

