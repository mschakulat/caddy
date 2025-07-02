package commands

import (
	"caddy/src/osarch"
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func Debug() *cli.Command {
	return &cli.Command{
		Name:  "debug",
		Usage: "Debug caddy",
		Action: func(ctx context.Context, cCtx *cli.Command) error {
			fmt.Println(osarch.GetRuntime("node"))
			fmt.Println(osarch.GetRuntime("pnpm"))

			return nil
		},
	}
}
