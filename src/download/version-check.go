package download

import (
	"caddy/src/config"
	"fmt"
	"os"
)

func HasVersion(tool string, version string) bool {
	var path string
	if tool == config.CaddyTool.Node {
		path = fmt.Sprintf("%s/%s", config.SystemPaths.Node, version)
	}

	if tool == config.CaddyTool.Pnpm {
		path = fmt.Sprintf("%s/%s", config.SystemPaths.Pnpm, version)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}
