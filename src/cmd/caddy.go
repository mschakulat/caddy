package main

import (
	"caddy/src/cmd/commands"
	"caddy/src/config"
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	config.InitConfig()

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Show current version",
	}

	app := &cli.Command{
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
			commands.Update(),
			commands.Debug(),
			commands.Init(),
			commands.Cleanup(),
		},
	}

	ctx := context.Background()
	if err := app.Run(ctx, os.Args); err != nil {
		log.Fatal(err)
	}
}
