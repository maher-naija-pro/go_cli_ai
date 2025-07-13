
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
	if str,err := context.ReadTerminalContext("/tmp/term_output.log"); err != nil {
		log.Fatalf("Failed to read context: %v %s", err, str)
	}

	app := &cli.App{
		Name:  "cli",
		Usage: "An advanced AI CLI",
		Commands: []*cli.Command{
			cmd.HelloCommand,
			cmd.OpenAICommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

