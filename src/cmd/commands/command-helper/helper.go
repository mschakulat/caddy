package command_helper

import (
	"fmt"
	"github.com/Masterminds/semver"
	"github.com/fatih/color"
	"github.com/logrusorgru/aurora"
	"os"
	"strings"
)

func SplitToolAndVersion(toolAndVersion string) (string, string) {
	parts := strings.Split(toolAndVersion, "@")
	tool := strings.ToLower(parts[0])

	if len(tool) == 0 {
		color.Red("Tool not specified")
		os.Exit(0)
	}

	version := ""
	if len(parts) == 2 {
		version = parts[1]
	}

	if !isVersionValid(version) {
		color.Red("Invalid version")
		os.Exit(0)
	}

	return tool, version
}

func PrintVersionRemovedNotification(tool string, version string) {
	fmt.Printf("%s %s@%s\n", aurora.Bold(aurora.Cyan("Removed")), tool, version)
}

func IsSemver(version string) bool {
	_, err := semver.NewVersion(version)
	if err != nil {
		return false
	}

	return true
}

func isVersionValid(version string) bool {
	if len(version) == 0 {
		return true
	}

	_, err := semver.NewVersion(version)
	if err != nil {
		return false
	}
	return true
}

func IsVersionFq(version string) bool {
	if !isVersionValid(version) {
		return false
	}

	return len(strings.Split(version, ".")) == 3
}
