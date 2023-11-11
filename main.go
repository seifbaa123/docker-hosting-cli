package main

import (
	"docker-hosting-cli/subcommands"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "dh-cli"
	app.Usage = "CLI tool for docker-hosting website"

	create := cli.Command{
		Name:    "create",
		Aliases: []string{},
		Usage:   "Create a new project",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Usage:    "Specify the project name",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			projectName := c.String("name")
			subcommands.Create(projectName)
			return nil
		},
	}

	update := cli.Command{
		Name:    "update",
		Aliases: []string{},
		Usage:   "Update the project",
		Action: func(c *cli.Context) error {
			subcommands.Update()
			return nil
		},
	}

	delete := cli.Command{
		Name:    "delete",
		Aliases: []string{},
		Usage:   "Delete the project",
		Action: func(c *cli.Context) error {
			subcommands.Delete()
			return nil
		},
	}

	app.Commands = []cli.Command{create, update, delete}
	app.Run(os.Args)
}
