package cmd

import (
	"fmt"
	"ai/config"
	"ai/openai"

	"github.com/urfave/cli/v2"
)

func OpenAICommand() *cli.Command {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		panic("failed to load config.yaml: " + err.Error())
	}

	var subCommands []*cli.Command
	for role, prompt := range cfg.OpenAI.Prompts {
		r := role
		p := prompt

		subCommands = append(subCommands, &cli.Command{
			Name:  r,
			Usage: fmt.Sprintf("Send a prompt using '%s' system prompt", r),
			Action: func(c *cli.Context) error {
				userInput := c.Args().First()
				if userInput == "" {
					return cli.Exit("Prompt is required", 1)
				}
				resp, err := openai.Ask(cfg, p, userInput)
				if err != nil {
					return cli.Exit("OpenAI error: "+err.Error(), 1)
				}
				fmt.Println(resp)
				return nil
			},
		})
	}

	return &cli.Command{
		Name:        "ask",
		Usage:       "Send a prompt using configured system roles",
		Subcommands: subCommands,
	}
}

