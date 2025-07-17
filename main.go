
package main

import (
	"log"
	"os"
	"github.com/urfave/cli/v2"
	"ai/cmd"

)

func main() {
	
app := &cli.App{
		Name:  "cli",
		Usage: "An advanced AI CLI",
		Commands: []*cli.Command{
			cmd.OpenAICommand("config.yaml", ""),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

