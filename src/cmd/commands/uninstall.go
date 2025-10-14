package commands

import (
	"context"
	"os"
	"strings"

	commandHelper "caddy/src/cmd/commands/command-helper"
	"caddy/src/config"
	nodePipeline "caddy/src/tools/node/pipeline"
	pnpmPipeline "caddy/src/tools/pnpm/pipeline"

	"github.com/fatih/color"
	"github.com/urfave/cli/v3"
)

func Uninstall() *cli.Command {
	return &cli.Command{
		Name:  "uninstall",
		Usage: "Uninstalls a tool with a specific version",
		Action: func(ctx context.Context, cCtx *cli.Command) error {
			tool, requestedVersion := commandHelper.SplitToolAndVersion(cCtx.Args().Get(0))

			if len(requestedVersion) == 0 {
				color.Yellow("Version not specified")
				return nil
			}

			switch tool {
			case strings.ToLower(config.CaddyTool.Node):
				nodePipeline.Uninstall(requestedVersion)
				commandHelper.PrintVersionRemovedNotification("node", requestedVersion)

			case strings.ToLower(config.CaddyTool.Pnpm):
				pnpmPipeline.Uninstall(requestedVersion)
				commandHelper.PrintVersionRemovedNotification("pnpm", requestedVersion)

			default:
				color.Red("Tool not supported")
				return nil
			}

			os.Exit(0)
			return nil
		},
	}
}
