package commands

import (
	commandhelper "caddy/src/cmd/commands/command-helper"
	"caddy/src/parser"
	"context"
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/urfave/cli/v3"
	"os"
)

func Pin() *cli.Command {
	return &cli.Command{
		Name:  "pin",
		Usage: "Pin a tool to a specific version",
		Action: func(ctx context.Context, cCtx *cli.Command) error {
			tool, requestedVersion := commandhelper.SplitToolAndVersion(cCtx.Args().Get(0))

			checkValidVersion(requestedVersion)

			path := []string{parser.GetIdentifier(), tool}
			commandhelper.WriteToJSON(path, requestedVersion)

			fmt.Printf("%s %s@%s\n", aurora.Bold(aurora.Cyan("Pinned")), tool, requestedVersion)

			return nil
		},
	}
}

func checkValidVersion(version string) {
	if len(version) == 0 {
		fmt.Println(aurora.Yellow("Version not specified"))
		os.Exit(0)
	}

	if !commandhelper.IsVersionFq(version) {
		fmt.Println(aurora.Yellow("Only fully qualified versions are supported"))
		os.Exit(0)
	}
}
