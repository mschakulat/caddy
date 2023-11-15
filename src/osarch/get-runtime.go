package osarch

import (
	"caddy/src/config"
	nodeOsarch "caddy/src/tools/node/osarch"
	pnpmOsarch "caddy/src/tools/pnpm/osarch"
)

func GetRuntime(tool string) (string, string) {
	if tool == config.CaddyTool.Node {
		return nodeOsarch.GetPlatform(), nodeOsarch.GetArch()
	} else if tool == config.CaddyTool.Pnpm {
		return pnpmOsarch.GetPlatform(), pnpmOsarch.GetArch()
	} else {
		panic("Unsupported tool")
	}
}
