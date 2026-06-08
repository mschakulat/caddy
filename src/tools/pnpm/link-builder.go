package pnpm

import (
	"fmt"
	"os"

	pnpmOsarch "caddy/src/tools/pnpm/osarch"

	"github.com/Masterminds/semver"
)

func GetLink(version string, arch string, platform string) string {
	platform, ext := normalizePlatformAndExt(version, platform)
	url := fmt.Sprintf("https://github.com/pnpm/pnpm/releases/download/v%s/pnpm-%s-%s", version, platform, arch)
	url += ext

	return url
}

func IsTarGz(version string) bool {
	platform := pnpmOsarch.GetPlatform()
	_, ext := normalizePlatformAndExt(version, platform)
	return ext == ".tar.gz"
}

/**
 * Starting from pnpm 11.0.0, macOS builds use "darwin" as platform with .tar.gz extension.
 */
func normalizePlatformAndExt(version string, platform string) (string, string) {
	v1, err := semver.NewVersion(version)
	if err != nil {
		fmt.Printf("Error parsing version: %s", err)
		os.Exit(0)
	}

	constraint, err := semver.NewConstraint(">= 11.0.0")
	if err != nil {
		fmt.Printf("Error parsing constraint: %s", err)
		os.Exit(0)
	}

	if constraint.Check(v1) && platform == "macos" {
		return "darwin", ".tar.gz"
	}

	if platform == "windows" {
		return platform, ".exe"
	}

	return platform, ""
}
