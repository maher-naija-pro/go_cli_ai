
package main

import (
	"log"
	"os"
	"github.com/urfave/cli/v2"
	"ai/cmd"

)

func main() {
	
configPath := os.Getenv("CONFIG_PATH")
if configPath == "" {
	configPath = "config.yaml"
}
app := &cli.App{
	Name:  "cli",
	Usage: "An advanced AI CLI",
	Commands: []*cli.Command{
		cmd.OpenAICommand(configPath, ""),
	},
}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

