package cmd

import (
	"ai/config"
	"ai/openai"
	"fmt"

	"github.com/urfave/cli/v2"
)

func OpenAICommand(configPath string, context string) *cli.Command {
	cfg := config.LoadConfig(configPath)
	if cfg == nil {
		return &cli.Command{
			Name:  "ask",
			Usage: "Send a prompt using configured system roles",
			Action: func(c *cli.Context) error {
				return cli.Exit("Config file not found", 1)
			},
		}
	}

	var subCommands []*cli.Command
	for role, prompt := range cfg.OpenAI.Prompts {
		r := role
		p := prompt + " " + context

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
