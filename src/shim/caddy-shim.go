package main

import (
	"caddy/src/config"
	"caddy/src/download"
	"caddy/src/parser"
	"caddy/src/tools"
	nodePipeline "caddy/src/tools/node/pipeline"
	pnpmPipeline "caddy/src/tools/pnpm/pipeline"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"syscall"
)

func main() {
	tool := os.Args[0]
	args := os.Args[1:]

	if isDevMode() && len(args) <= 2 {
		tool = os.Args[1]
		args = os.Args[2:]
	}

	tool = path.Base(tool)

	var version string
	var versions parser.CaddyStruct
	defaultVersion := tools.GetDefaultVersion(tool)

	if parser.HasPackageJson() || defaultVersion == nil {
		versions = parser.GetVersionsFromPackage(nil, nil)
		if tool == "node" {
			version = installNode(versions.Node)
		} else if tool == "pnpm" {
			installNode(versions.Node)
			version = installPnpm(versions.Pnpm)
		} else {
			fmt.Printf("No default %s version set\n", tool)
			os.Exit(0)
		}
		if defaultVersion == nil {
			tools.SetDefaultVersion(tool, version)
		}
		tools.CleanupTmp()
		fmt.Println()
	}

	if len(version) == 0 && defaultVersion != nil {
		version = *defaultVersion
	}

	path, err := exec.LookPath(tools.ToolBin(tool, version))
	if err != nil {
		fmt.Printf("Could not find %s version %s\n", tool, version)
		os.Exit(0)
	}

	// Prepare arguments: first argument needs to be the program name
	args = append([]string{tool}, args...)

	// Replace this process with the given command
	err = syscall.Exec(path, args, os.Environ())

	// If we reach this point, an error occured
	if err != nil {
		log.Fatal(err)
	}
}

func isDevMode() bool {
	return os.Getenv("CADDY_DEV") == "true"
}

func installNode(version string) string {
	if len(version) != 0 {
		if !download.HasVersion(config.CaddyTool.Node, version) {
			nodePipeline.Install(version)
		}
		return version
	}
	return ""
}

func installPnpm(version string) string {
	if len(version) != 0 {
		if !download.HasVersion(config.CaddyTool.Pnpm, version) {
			pnpmPipeline.Install(version)
		}
		return version
	}
	return ""
}
