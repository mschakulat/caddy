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
		break
	case "darwin":
		platform = "darwin"
		break
	case "windows":
		platform = "win"
		break
	default:
		err := fmt.Sprintf("Platform [%s] not supported", currentPlatform)
		fmt.Println(err)
		os.Exit(1)
	}

	return platform
}

func GetArch() string {
	var arch string

	currentArch := runtime.GOARCH

	switch currentArch {
	case "x86_64", "amd64":
		arch = "x64"
		break
	case "armv":
		arch = "armv7l"
		break
	case "arm64", "aarch64":
		arch = "arm64"
		break
	default:
		err := fmt.Sprintf("Architecture [%s] not supported", currentArch)
		fmt.Println(err)
		os.Exit(1)
	}

	return arch
}
