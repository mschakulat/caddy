package pnpm

import (
	"caddy/src/config"
	"fmt"
	"io"
	"os"
)

func Copy(src string, version string) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	}
	defer func(srcFile *os.File) {
		_ = srcFile.Close()
	}(srcFile)

	targetDir := fmt.Sprintf("%s/%s", config.SystemPaths.Pnpm, version)
	_ = os.MkdirAll(targetDir, 0755)

	dstFile, err := os.OpenFile(fmt.Sprintf("%s/pnpm", targetDir), os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println(err)
	}
	defer func(dstFile *os.File) {
		_ = dstFile.Close()
	}(dstFile)

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		fmt.Println(err)
	}
}
