package parser

import (
	"caddy/src/fileinfo"
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/tidwall/gjson"
	"os"
	"path/filepath"
)

type CaddyStruct struct {
	Node    string
	Pnpm    string
	Extends string
}

func GetVersionsFromPackage(extend *string, currentDir *string) CaddyStruct {
	var dir string

	if extend == nil {
		dir, _ = os.Getwd()
	} else {
		dir = *currentDir
	}

	path := filepath.Join(dir, "package.json")

	if extend != nil {
		path = filepath.Join(dir, *extend)
	}

	fileContent, _ := os.ReadFile(path)

	node := gjson.Get(string(fileContent), "caddy.node")
	pnpm := gjson.Get(string(fileContent), "caddy.pnpm")
	extends := gjson.Get(string(fileContent), "caddy.extends")

	if len(extends.Str) > 0 {
		if !fileinfo.Exist(filepath.Join(dir, extends.Str)) {
			fmt.Println(aurora.Yellow("Tried to extend a non-existing package.json file"))
			os.Exit(0)
		}
		var current string
		if extend != nil {
			current = filepath.Dir(filepath.Join(dir, extends.Str))
		}
		return GetVersionsFromPackage(&extends.Str, &current)
	}

	return CaddyStruct{
		Node:    node.Str,
		Pnpm:    pnpm.Str,
		Extends: extends.Str,
	}
}

func HasPackageJson() bool {
	currentDir, _ := os.Getwd()
	path := filepath.Join(currentDir, "package.json")
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
