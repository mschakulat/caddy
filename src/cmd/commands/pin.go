package commands

import (
	commandhelper "caddy/src/cmd/commands/command-helper"
	"caddy/src/parser"
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/tidwall/sjson"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Pin() *cli.Command {
	return &cli.Command{
		Name:  "pin",
		Usage: "Pin a tool to a specific version",
		Action: func(cCtx *cli.Context) error {
			tool, requestedVersion := commandhelper.SplitToolAndVersion(cCtx.Args().Get(0))

			checkValidVersion(requestedVersion)

			dir, _ := os.Getwd()
			packageJson := filepath.Join(dir, "package.json")

			fileContent, _ := os.ReadFile(packageJson)

			newJsonString, _ := sjson.Set(
				string(fileContent), strings.Join([]string{parser.GetIdentifier(), tool}, "."), requestedVersion,
			)

			err := os.WriteFile(packageJson, []byte(newJsonString), 0644)
			if err != nil {
				log.Fatalf("failed writing to file: %s", err)
			}

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
