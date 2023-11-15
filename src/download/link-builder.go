package download

import (
	"caddy/src/config"
	"caddy/src/osarch"
	"caddy/src/tools/node"
	"caddy/src/tools/pnpm"
)

func GetLink(tool string, version string) string {
	platform, arch := osarch.GetRuntime(tool)

	if tool == config.CaddyTool.Node {
		return node.GetLink(version, arch, platform)
	}

	if tool == config.CaddyTool.Pnpm {
		return pnpm.GetLink(version, arch, platform)
	}

	panic("Unsupported tool")
}
