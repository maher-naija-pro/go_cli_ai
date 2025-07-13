
package main

import (
	"log"
	"os"
	"github.com/urfave/cli/v2"
	"ai/cmd"
	"ai/context"
)

func main() {
	// Read and log context
	context_term, err := context.ReadTerminalContext("/tmp/term_output.log");
    if err != nil {
		log.Fatalf("Failed to read context: %v %s", err, context_term)
	}
	
app := &cli.App{
		Name:  "cli",
		Usage: "An advanced AI CLI",
		Commands: []*cli.Command{
			cmd.HelloCommand,
			cmd.OpenAICommand("config.yaml", context_term),
			cmd.CaptureCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

