package tools

import (
	"caddy/src/config"
	"fmt"
)

func ToolBin(tool string, version string) string {
	if tool == "node" {
		return fmt.Sprintf("%s/%s/%s", config.SystemPaths.Node, version, "bin/node")
	} else if tool == "pnpm" {
		return fmt.Sprintf("%s/%s/%s", config.SystemPaths.Pnpm, version, "pnpm")
	} else if tool == "npm" {
		return fmt.Sprintf("%s/%s/%s", config.SystemPaths.Node, version, "lib/node_modules/npm/bin/npm-cli.js")
	} else if tool == "npx" {
		return fmt.Sprintf("%s/%s/%s", config.SystemPaths.Node, version, "lib/node_modules/npm/bin/npx-cli.js")
	} else {
		panic("Unsupported tool")
	}
}
