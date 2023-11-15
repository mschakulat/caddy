package tools

import (
	"os"
)

func ListVersions(path string) ([]string, error) {
	entries, err := os.ReadDir(path)

	if err != nil {
		return nil, err
	}

	var dirs []string
	for _, entry := range entries {
		if entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				return nil, err
			}
			dirs = append(dirs, info.Name())
		}
	}

	return dirs, nil
}
