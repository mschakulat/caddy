package pipeline

import (
	"caddy/src/config"
	"caddy/src/download"
	"fmt"
	"github.com/fatih/color"
	"os"
)

func Uninstall(version string) {
	if !download.HasVersion(config.CaddyTool.Pnpm, version) {
		color.Yellow("Version not installed")
		return
	}
	err := os.RemoveAll(fmt.Sprintf("%s/%s", config.SystemPaths.Pnpm, version))
	if err != nil {
		color.Red("Error while uninstalling pnpm", err)
	}
}
