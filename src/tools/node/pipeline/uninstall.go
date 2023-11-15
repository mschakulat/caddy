package pipeline

import (
	"caddy/src/config"
	"caddy/src/download"
	"fmt"
	"github.com/fatih/color"
	"os"
)

func Uninstall(version string) {
	if !download.HasVersion(config.CaddyTool.Node, version) {
		msg := fmt.Sprintf("Version %s not installed\n", version)
		color.Cyan(msg)
		return
	}
	err := os.RemoveAll(fmt.Sprintf("%s/%s", config.SystemPaths.Node, version))
	if err != nil {
		color.Red("Error while uninstalling node", err)
	}
}
