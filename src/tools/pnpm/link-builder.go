package pnpm

import "fmt"

func GetLink(version string, arch string, platform string) string {
	url := fmt.Sprintf("https://github.com/pnpm/pnpm/releases/download/v%s/pnpm-%s-%s", version, platform, arch)

	if "windows" == platform {
		url += ".exe"
	}

	return url
}
