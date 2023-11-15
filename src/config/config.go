package config

import (
	"os"
	"path/filepath"
)

var Version = "0.2.0"

var ProjectName = "caddy"

type CaddyToolStruct struct {
	Node string
	Pnpm string
}

var CaddyTool = CaddyToolStruct{
	Node: "node",
	Pnpm: "pnpm",
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
