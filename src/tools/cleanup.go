package tools

import (
	"caddy/src/config"
	"github.com/fatih/color"
	"os"
)

func CleanupTmp() {
	err := os.RemoveAll(config.SystemPaths.Temp)
	if err != nil {
		color.Red("Error while cleaning up temporary files")
	}
	err = os.MkdirAll(config.SystemPaths.Temp, 0755)
	if err != nil {
		color.Red("Error while recreating temporary directory")
	}
}
