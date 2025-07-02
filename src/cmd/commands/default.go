package commands

import (
	commandhelper "caddy/src/cmd/commands/command-helper"
	"caddy/src/config"
	"caddy/src/tools"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/logrusorgru/aurora"
	"github.com/urfave/cli/v3"
)

func Default() *cli.Command {
	return &cli.Command{
		Name:  "default",
		Usage: "Set a default version of a tool",
		Action: func(ctx context.Context, cCtx *cli.Command) error {
			tool, requestedVersion := commandhelper.SplitToolAndVersion(cCtx.Args().Get(0))
			if !commandhelper.IsSemver(requestedVersion) {
				color.Yellow("Version not valid")
				os.Exit(0)
			}

			if tool == strings.ToLower(config.CaddyTool.Node) {
				setToDefault(config.CaddyTool.Node, requestedVersion)
			} else if tool == strings.ToLower(config.CaddyTool.Pnpm) {
				setToDefault(config.CaddyTool.Pnpm, requestedVersion)
			} else {
				color.Red("Tool not supported\n")
				os.Exit(0)
			}
			return nil
		},
	}
}

func hasDownloadedVersion(tool string, version string) bool {
	var versions []string

	if tool == config.CaddyTool.Node {
		versions, _ = tools.ListVersions(config.SystemPaths.Node)
	} else if tool == config.CaddyTool.Pnpm {
		versions, _ = tools.ListVersions(config.SystemPaths.Pnpm)
	} else {
		color.Red("Tool not supported\n")
		os.Exit(0)
	}

	for _, v := range versions {
		if v == version {
			return true
		}
	}
	return false
}

func setToDefault(tool string, version string) {
	if !hasDownloadedVersion(tool, version) {
		fmt.Println(aurora.Yellow("Version not installed"))
		os.Exit(0)
	}

	if tool == config.CaddyTool.Node {
		tools.SetDefaultVersion(config.CaddyTool.Node, version)
	} else if tool == config.CaddyTool.Pnpm {
		tools.SetDefaultVersion(config.CaddyTool.Pnpm, version)
	} else {
		color.Red("Tool not supported\n")
		os.Exit(0)
	}

	fmt.Printf("%s %s@%s\n", aurora.Bold(aurora.Cyan("Default")), strings.ToLower(tool), version)
}
