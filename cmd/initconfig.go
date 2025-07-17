package cmd

import (
	"fmt"
	"os"
	"ai/config"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

func InitConfigCommand() *cli.Command {
	return &cli.Command{
		Name:  "init-config",
		Usage: "Generate a default config.yaml file in the current directory",
		Action: func(c *cli.Context) error {
			cfg := config.DefaultConfig()
			file, err := os.Create("config.yaml")
			if err != nil {
				return cli.Exit("Failed to create config.yaml: "+err.Error(), 1)
			}
			defer file.Close()

			enc := yaml.NewEncoder(file)
			defer enc.Close()
			if err := enc.Encode(cfg); err != nil {
				return cli.Exit("Failed to write config.yaml: "+err.Error(), 1)
			}
			fmt.Println("Default config.yaml generated.")
			return nil
		},
	}
} 