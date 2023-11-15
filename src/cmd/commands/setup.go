package commands

import (
	"caddy/src/config"
	"caddy/src/setup"
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/urfave/cli/v2"
	"os"
)

func Setup() *cli.Command {
	return &cli.Command{
		Name:  "setup",
		Usage: "Enables Caddy for the current user and creates the necessary directories",
		Action: func(cCtx *cli.Context) error {
			setup.InitSystemPath()
			setup.CreateBin(config.CaddyTool.Node)
			setup.CreateBin(config.CaddyTool.Pnpm)
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
