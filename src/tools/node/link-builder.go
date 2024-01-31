package node

import (
	"fmt"
	"os"

	"github.com/Masterminds/semver"
)

func GetLink(version string, arch string, platform string) string {
	arch = fallbackToAmd(version, arch)
	return fmt.Sprintf("https://nodejs.org/dist/v%s/node-v%s-%s-%s.tar.gz", version, version, platform, arch)
}

/**
 * Before Node.js 16.0.0 there is no arm64 build for macOS
 */
func fallbackToAmd(version string, arch string) string {
	v1, err := semver.NewVersion(version)
	if err != nil {
		fmt.Printf("Error parsing version: %s", err)
		os.Exit(0)
	}

	constraint, err := semver.NewConstraint("< 16.0.0")
	if err != nil {
		fmt.Printf("Error parsing constraint: %s", err)
		os.Exit(0)
	}

	fallback := arch
	if constraint.Check(v1) {
		fallback = "x64"
	}

	return fallback
}
