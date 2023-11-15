package node

import "fmt"

func GetLink(version string, arch string, platform string) string {
	return fmt.Sprintf("https://nodejs.org/dist/v%s/node-v%s-%s-%s.tar.gz", version, version, platform, arch)
}
