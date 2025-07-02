package commands

import (
	"caddy/src/config"
	"caddy/src/setup"
	"context"
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/urfave/cli/v3"
)

func Setup() *cli.Command {
	return &cli.Command{
		Name:  "setup",
		Usage: "Enables Caddy for the current user and creates the necessary directories",
		Action: func(ctx context.Context, cCtx *cli.Command) error {
			setup.InitSystemPath()
			setup.CreateBin(config.CaddyTool.Node)
			setup.CreateBin(config.CaddyTool.Pnpm)
			setup.CreateBin(config.CaddyTool.Npm)
			setup.CreateBin(config.CaddyTool.Npx)
			setup.Enable()

			if !cCtx.Bool("skip-msg") {
				fmt.Println(aurora.Bold(aurora.BrightCyan("Setup complete")))
			}

			os.Exit(0)
			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "skip-msg",
				Usage: "Skip the setup message",
			},
		},
	}
}
