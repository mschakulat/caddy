package setup

import (
	"caddy/src/config"
	"fmt"
	"os"
)

func CreateBin(tool string) {
	targetBin := fmt.Sprintf("%s/%s", config.SystemPaths.Bin, tool)
	err := os.Symlink("caddy-shim", targetBin)
	if err != nil {
		fmt.Println(err)
	}
}
