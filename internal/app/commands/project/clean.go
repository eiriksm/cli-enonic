package project

import (
	"github.com/urfave/cli"
)

var Clean = cli.Command{
	Name:  "clean",
	Usage: "Clean current project",
	Action: func(c *cli.Context) error {

		return nil
	},
}