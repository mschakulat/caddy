package commands

import (
	"caddy/src/osarch"
	"fmt"
	"github.com/urfave/cli/v2"
)

func Debug() *cli.Command {
	return &cli.Command{
		Name:  "debug",
		Usage: "Debug caddy",
		Action: func(cCtx *cli.Context) error {
			fmt.Println(osarch.GetRuntime("node"))
			fmt.Println(osarch.GetRuntime("pnpm"))

			return nil
		},
	}
}
