package commands

import (
	"caddy/src/cmd/commands/command-helper"
	"caddy/src/config"
	"caddy/src/tools"
	"caddy/src/tools/node"
	nodePipeline "caddy/src/tools/node/pipeline"
	"caddy/src/tools/pnpm"
	pnpmPipeline "caddy/src/tools/pnpm/pipeline"
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

func Install() *cli.Command {
	return &cli.Command{
		Name:  "install",
		Usage: "Installs a tool in your toolchain",
		Action: func(cCtx *cli.Context) error {
			tool, requestedVersion := command_helper.SplitToolAndVersion(cCtx.Args().Get(0))

			switch tool {
			case strings.ToLower(config.CaddyTool.Node):
				version := node.VersionByConstraint(requestedVersion)
				nodePipeline.Install(version)
				fmt.Println()
				setToDefault(config.CaddyTool.Node, version)
				tools.CleanupTmp()

			case strings.ToLower(config.CaddyTool.Pnpm):
				version := pnpm.VersionByConstraint(requestedVersion)
				pnpmPipeline.Install(version)
				fmt.Println()
				setToDefault(config.CaddyTool.Pnpm, version)
				tools.CleanupTmp()

			default:
				color.Red("Tool not supported")
			}

			os.Exit(0)
			return nil
		},
	}
}
