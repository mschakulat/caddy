package commands

import (
	nodePipeline "caddy/src/tools/node/pipeline"
	pmpmPipeline "caddy/src/tools/pnpm/pipeline"
	"context"

	"github.com/urfave/cli/v3"
)

func List() *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "Lists all installed tools and their versions",
		Commands: []*cli.Command{
			{
				Name: "node",
				Action: func(ctx context.Context, cCtx *cli.Command) error {
					nodePipeline.List()
					return nil
				},
			},
			{
				Name: "pnpm",
				Action: func(ctx context.Context, cCtx *cli.Command) error {
					pmpmPipeline.List()
					return nil
				},
			},
		},
	}
}
