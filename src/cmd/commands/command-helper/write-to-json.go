package command_helper

import (
	"github.com/tidwall/pretty"
	"github.com/tidwall/sjson"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func WriteToJSON(path []string, value string) {
	dir, _ := os.Getwd()
	packageJson := filepath.Join(dir, "package.json")

	fileContent, _ := os.ReadFile(packageJson)

	newJsonString, _ := sjson.Set(
		string(fileContent), strings.Join(path, "."), value,
	)

	prettyJson := pretty.Pretty([]byte(newJsonString))

	err := os.WriteFile(packageJson, prettyJson, 0644)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
}
