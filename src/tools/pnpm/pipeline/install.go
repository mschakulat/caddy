package pipeline

import (
	"caddy/src/config"
	"caddy/src/download"
	"caddy/src/tools/pnpm"
)

func Install(version string) {
	if download.HasVersion(config.CaddyTool.Pnpm, version) {
		return
	}
	pnpmPath := download.FetchTool(
		download.GetLink(config.CaddyTool.Pnpm, version), config.SystemPaths.Temp,
		download.Description("pnpm", version),
	)
	pnpm.Copy(pnpmPath, version)
}
