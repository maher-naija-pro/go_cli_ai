package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
)

const maxLines = 10 // nombre maximum de lignes à conserver


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
        return nil // rien à faire
    }

    // Conserver uniquement les dernières lignes
    trimmed := lines[len(lines)-maxLines:]

    // Réécrire le fichier
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

func main() {
    file := "/tmp/term_output.log"
    fmt.Println("Capturing terminal session, everything you type will be recorded...")
    err := captureTerminalOutput(file)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    // Nettoyage après capture
    err = trimFile(file, maxLines)
    if err != nil {
        fmt.Printf("Error trimming file: %v\n", err)
        return
    }

    fmt.Printf("Terminal session saved to: %s (trimmed to last %d lines)\n", file, maxLines)
}

