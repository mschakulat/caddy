package commands

import (
	"caddy/src/cmd/commands/command-helper"
	"caddy/src/config"
	nodePipeline "caddy/src/tools/node/pipeline"
	pnpmPipeline "caddy/src/tools/pnpm/pipeline"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

func Uninstall() *cli.Command {
	return &cli.Command{
		Name:  "uninstall",
		Usage: "Uninstalls a tool in your toolchain",
		Action: func(cCtx *cli.Context) error {
			tool, requestedVersion := command_helper.SplitToolAndVersion(cCtx.Args().Get(0))

			if len(requestedVersion) == 0 {
				color.Yellow("Version not specified")
				return nil
			}

			switch tool {
			case strings.ToLower(config.CaddyTool.Node):
				nodePipeline.Uninstall(requestedVersion)
				command_helper.PrintVersionRemovedNotification("node", requestedVersion)

			case strings.ToLower(config.CaddyTool.Pnpm):
				pnpmPipeline.Uninstall(requestedVersion)
				command_helper.PrintVersionRemovedNotification("pnpm", requestedVersion)

			default:
				color.Red("Tool not supported")
				return nil
			}

			os.Exit(0)
			return nil
		},
	}
}
