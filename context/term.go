package context

import (
    "log"
    "os"
    "fmt"
)

// Constant marker to start context from
const contextStartMarker = "===START==="
func ReadTerminalContext(filename string) (string, error) {
    const marker = "---CONTEXT START---"

    //log.Printf("Reading file: %s", filename)
    content, err := os.ReadFile(filename)
    if err != nil {
        log.Printf("Failed to read file: %s, error: %v", filename, err)
        return "", fmt.Errorf("failed to read file: %w", err)
    }

    //log.Printf("File content: %s", string(content))
    text := string(content)
    log.Println("READ CONTEXT OK")

    return text, nil
} 