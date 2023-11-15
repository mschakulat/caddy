package pipeline

import (
	"caddy/src/config"
	download "caddy/src/download"
	"caddy/src/tools/node"
	"github.com/fatih/color"
)

func Install(version string) {
	if download.HasVersion(config.CaddyTool.Node, version) {
		return
	}
	archive := download.FetchTool(
		download.GetLink(config.CaddyTool.Node, version), config.SystemPaths.Temp,
		download.Description("node", version),
	)
	err := node.Uncompress(archive, version)
	if err != nil {
		color.Red("Error while installing node", err)
	}
}
