package commands

import (
	"context"
	"fmt"

	commandHelper "caddy/src/cmd/commands/command-helper"
	"caddy/src/parser"
	nodePipeline "caddy/src/tools/node/pipeline"
	pmpmPipeline "caddy/src/tools/pnpm/pipeline"

	"github.com/logrusorgru/aurora"
	"github.com/urfave/cli/v3"
)

func Repair() *cli.Command {
	return &cli.Command{
		Name:  "repair",
		Usage: "Repair binaries in case of corruption (e.g. exec format error)",
		Action: func(ctx context.Context, cCtx *cli.Command) error {
			versions := parser.GetVersionsFromPackage(nil, nil)

			fmt.Println(aurora.Bold(aurora.BrightCyan("Reinstalling node " + versions.Node)))
			nodePipeline.Uninstall(versions.Node)
			commandHelper.PrintVersionRemovedNotification("node", versions.Node)
			nodePipeline.Install(versions.Node)
			fmt.Println(aurora.Bold(aurora.BrightCyan("Reinstalling node completed")))

			fmt.Println()

			fmt.Println(aurora.Bold(aurora.BrightCyan("Reinstalling Node " + versions.Node)))
			pmpmPipeline.Uninstall(versions.Pnpm)
			commandHelper.PrintVersionRemovedNotification("pnpm", versions.Pnpm)
			pmpmPipeline.Install(versions.Pnpm)
			fmt.Println(aurora.Bold(aurora.BrightCyan("Reinstalling pnpm completed")))

			return nil
		},
	}
}
