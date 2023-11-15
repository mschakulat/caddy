package pipeline

import (
	"caddy/src/config"
	"caddy/src/tools"
	"fmt"
	"github.com/fatih/color"
	"github.com/logrusorgru/aurora"
	"os"
)

func List() {
	versions, err := tools.ListVersions(config.SystemPaths.Pnpm)
	if err != nil {
		fmt.Println(err)
	}

	if len(versions) == 0 {
		color.Yellow("No pnpm runtimes installed")
		os.Exit(0)
	}

	fmt.Printf("⚡️ Pnpm runtimes in your toolchain:\n\n")

	defaultVersion := tools.GetDefaultVersion(config.CaddyTool.Pnpm)

	for _, version := range versions {
		hint := ""
		if version == *defaultVersion {
			hint = "(default)"
		}
		fmt.Printf("%s %s\n", version, aurora.BrightCyan(hint))
	}

	os.Exit(0)
}
