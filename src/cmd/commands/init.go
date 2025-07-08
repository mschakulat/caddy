package commands

import (
	commandhelper "caddy/src/cmd/commands/command-helper"
	"caddy/src/config"
	"caddy/src/parser"
	"caddy/src/tools"
	"context"
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/urfave/cli/v3"
	"os"
)

func Init() *cli.Command {
	skipEngines := "skip-engines"
	skipNode := "skip-node"
	skipPnpm := "skip-pnpm"

	return &cli.Command{
		Name:  "init",
		Usage: "Init caddy",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  skipNode,
				Usage: "Skip pinning node version",
			},
			&cli.BoolFlag{
				Name:  skipPnpm,
				Usage: "Skip pinning pnpm version",
			},
			&cli.BoolFlag{
				Name:  skipEngines,
				Usage: "Skip pinning engine versions",
			},
		},
		Action: func(ctx context.Context, cCtx *cli.Command) error {
			if !cCtx.Bool(skipNode) {
				initNode(cCtx.Bool(skipEngines))
			}

			if !cCtx.Bool(skipPnpm) {
				initPnpm(cCtx.Bool(skipEngines))
			}

			fmt.Println(aurora.Bold(aurora.BrightCyan("Caddy initialized successfully!")))

			return nil
		},
	}
}

func initNode(skipEngines bool) {
	path := []string{parser.GetIdentifier(), "node"}
	versions, _ := tools.ListVersions(config.SystemPaths.Node)

	if len(versions) == 0 {
		fmt.Println(aurora.Bold(aurora.Yellow("No node versions found. Please install a node version first.")))
		os.Exit(0)
	}

	latestVersion := versions[len(versions)-1]
	commandhelper.WriteToJSON(path, latestVersion)

	if !skipEngines {
		pathEngine := []string{"engines", "node"}
		commandhelper.WriteToJSON(pathEngine, latestVersion)
	}
}

func initPnpm(skipEngines bool) {
	path := []string{parser.GetIdentifier(), "pnpm"}
	versions, _ := tools.ListVersions(config.SystemPaths.Pnpm)

	if len(versions) == 0 {
		fmt.Println(aurora.Bold(aurora.Yellow("No pnpm versions found. Please install a pnpm version first.")))
		os.Exit(0)
	}

	latestVersion := versions[len(versions)-1]
	commandhelper.WriteToJSON(path, versions[len(versions)-1])

	if !skipEngines {
		pathEngine := []string{"engines", "pnpm"}
		commandhelper.WriteToJSON(pathEngine, latestVersion)
	}
}
