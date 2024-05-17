package config

import (
	"os"
	"path/filepath"
)

var Version = "0.3.16"

var ProjectName = "caddy"

var PackageJsonIdentifier = "caddy"

type CaddyToolStruct struct {
	Node string
	Pnpm string
	Npm  string
	Npx  string
}

var CaddyTool = CaddyToolStruct{
	Node: "node",
	Pnpm: "pnpm",
	Npm:  "npm",
	Npx:  "npx",
}

type SystemPathStructure struct {
	Base  string
	Bin   string
	Tools string
	Pnpm  string
	Node  string
	Temp  string
	Cache string
}

var HomeDir, _ = os.UserHomeDir()
var SystemPaths = SystemPathStructure{
	Base:  filepath.Join(HomeDir, "."+ProjectName),
	Bin:   filepath.Join(HomeDir, "."+ProjectName, "/bin"),
	Tools: filepath.Join(HomeDir, "."+ProjectName, "/tools"),
	Pnpm:  filepath.Join(HomeDir, "."+ProjectName, "/tools/pnpm"),
	Node:  filepath.Join(HomeDir, "."+ProjectName, "/tools/node"),
	Temp:  filepath.Join(HomeDir, "."+ProjectName, "/tmp"),
	Cache: filepath.Join(HomeDir, "."+ProjectName, "/cache"),
}

var DefaultsPath = filepath.Join(SystemPaths.Tools, "defaults.gob")
