package commands

import (
	nodePipeline "caddy/src/tools/node/pipeline"
	pmpmPipeline "caddy/src/tools/pnpm/pipeline"
	"github.com/urfave/cli/v2"
)

func List() *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "Lists all installed tools in your toolchain",
		Subcommands: []*cli.Command{
			{
				Name: "node",
				Action: func(*cli.Context) error {
					nodePipeline.List()
					return nil
				},
			},
			{
				Name: "pnpm",
				Action: func(*cli.Context) error {
					pmpmPipeline.List()
					return nil
				},
			},
		},
	}
}
