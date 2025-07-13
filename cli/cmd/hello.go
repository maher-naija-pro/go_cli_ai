package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var HelloCommand = &cli.Command{
	Name:  "hello",
	Usage: "Say Hello",
	Action: func(c *cli.Context) error {
		fmt.Println("Hello world!")
		return nil
	},
}

