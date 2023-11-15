package tools

import (
	"github.com/fatih/color"
	"os"
	"strings"
)

func VersionByConstraint(versions []string, req string) string {
	var returnVersion string
	versionParts := strings.Split(req, ".")

	if len(req) == 0 {
		return versions[0]
	}

OuterLoop:
	for _, version := range versions {
		currentParts := strings.Split(version, ".")

		switch len(versionParts) {
		case 1:
			if currentParts[0] == versionParts[0] {
				returnVersion = version
				break OuterLoop
			}
		case 2:
			if currentParts[0] == versionParts[0] && currentParts[1] == versionParts[1] {
				returnVersion = version
				break OuterLoop
			}
		case 3:
			if version == req {
				return req
			}
		}
	}

	if len(returnVersion) == 0 {
		color.Red("No version match found")
		os.Exit(0)
	}

	return returnVersion
}
