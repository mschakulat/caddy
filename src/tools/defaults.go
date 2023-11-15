package tools

import (
	"caddy/src/config"
	"caddy/src/fileinfo"
	"encoding/gob"
	"fmt"
	"os"
)

type DefaultVersion struct {
	Node *string
	Pnpm *string
}

func GetDefaultVersion(tool string) *string {
	createFileIfNotExists()
	if tool == config.CaddyTool.Node {
		return getVersions().Node
	} else if tool == config.CaddyTool.Pnpm {
		return getVersions().Pnpm
	}
	return nil
}

func SetDefaultVersion(tool string, version string) {
	createFileIfNotExists()
	versions := getVersions()

	if tool == config.CaddyTool.Node {
		versions.Node = &version
	} else if tool == config.CaddyTool.Pnpm {
		versions.Pnpm = &version
	} else {
		panic("Unsupported tool")
	}

	setVersions(versions)
}

func createFileIfNotExists() {
	if !fileinfo.Exist(config.DefaultsPath) {
		defaults := DefaultVersion{}
		setVersions(defaults)
	}
}

func setVersions(defaults DefaultVersion) {
	file, err := os.Create(config.DefaultsPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(defaults)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getVersions() DefaultVersion {
	file, err := os.Open(config.DefaultsPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var versions DefaultVersion

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&versions)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	return versions
}
