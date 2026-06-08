package pipeline

import (
	"fmt"

	"caddy/src/config"
	"caddy/src/download"
	"caddy/src/tools/pnpm"

	"github.com/fatih/color"
)

func Install(version string) {
	if download.HasVersion(config.CaddyTool.Pnpm, version) {
		return
	}
	pnpmPath := download.FetchTool(
		download.GetLink(config.CaddyTool.Pnpm, version), config.SystemPaths.Temp,
		download.Description("pnpm", version),
	)

	if pnpm.IsTarGz(version) {
		fmt.Printf("pnpm %s: %s is a tar.gz archive → extracting\n", version, pnpmPath)
		err := pnpm.Uncompress(pnpmPath, version)
		if err != nil {
			color.Red("Error while installing pnpm: %s", err)
		}
	} else {
		fmt.Printf("pnpm %s: %s is a plain binary → copying\n", version, pnpmPath)
		pnpm.Copy(pnpmPath, version)
	}
}
