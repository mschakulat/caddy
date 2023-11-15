package osarch

import (
	"fmt"
	"os"
	"runtime"
)

func GetPlatform() string {
	var platform string

	currentPlatform := runtime.GOOS
	switch currentPlatform {
	case "linux":
		platform = "linux"
	case "darwin":
		platform = "darwin"
	case "windows":
		platform = "win"
	default:
		err := fmt.Sprintf("Platform [%s] not tools", currentPlatform)
		fmt.Println(err)
		os.Exit(1)
	}

	return platform
}

func GetArch() string {
	var arch string

	currentArch := runtime.GOARCH
	switch currentArch {
	case "x86_64":
	case "amd64":
		arch = "x64"
	case "armv":
		arch = "armv7l"
	case "arm64":
	case "aarch64":
		arch = "arm64"
	default:
		err := fmt.Sprintf("Architecture [%s] not tools", currentArch)
		fmt.Println(err)
		os.Exit(1)
	}

	return arch
}
