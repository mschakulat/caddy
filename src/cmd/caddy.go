package main

import (
	"caddy/src/cmd/commands"
	"caddy/src/config"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	config.InitConfig()
	
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Show current version",
	}

	app := &cli.App{
		Name:    config.ProjectName,
		Version: config.Version,
		Usage:   "JavaScript tool manager",
		Commands: []*cli.Command{
			commands.Setup(),
			commands.Install(),
			commands.Uninstall(),
			commands.List(),
			commands.Default(),
			commands.Pin(),
			commands.Config(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
