package commands

import (
	"caddy/src/config"
	"context"
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/urfave/cli/v3"
	"os"
)

func Cleanup() *cli.Command {
	skipNode := "skip-node"
	skipPnpm := "skip-pnpm"

	return &cli.Command{
		Name:  "cleanup",
		Usage: "Cleanup tool versions",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  skipNode,
				Usage: "Skip cleanup node versions",
			},
			&cli.BoolFlag{
				Name:  skipPnpm,
				Usage: "Skip cleanup pnpm versions",
			},
		},
		Action: func(ctx context.Context, cCtx *cli.Command) error {
			if !cCtx.Bool(skipPnpm) {
				fmt.Println(aurora.Bold(aurora.BrightCyan("Cleanup pnpm versions...")))
				err := recreateDirectory(config.SystemPaths.Pnpm)
				if err != nil {
					fmt.Printf("Error cleaning up directory: %s", err)
					os.Exit(1)
				}
			}

			if !cCtx.Bool(skipNode) {
				fmt.Println(aurora.Bold(aurora.BrightCyan("Cleanup node versions...")))
				err := recreateDirectory(config.SystemPaths.Node)
				if err != nil {
					fmt.Printf("Error cleaning up directory: %s", err)
					os.Exit(1)
				}
			}

			return nil
		},
	}
}

func recreateDirectory(dir string) error {
	if err := os.RemoveAll(dir); err != nil {
		return err
	}
	return os.MkdirAll(dir, 0755)
}
