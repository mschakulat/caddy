package setup

import (
	"caddy/src/config"
	"fmt"
	"github.com/logrusorgru/aurora"
	"os"
	"path/filepath"
)

type EnableCommands struct {
	Home string
	Path string
}

func Enable() {
	commands := EnableCommands{
		Home: "export CADDY_HOME=\"$HOME/.caddy\"",
		Path: "export PATH=\"$CADDY_HOME/bin:$PATH\"",
	}

	if !hasSupportedShell() {
		fmt.Println(aurora.Bold(aurora.Yellow("Unsupported shell")))
		fmt.Println("Only bash and zsh are supported at the moment")
		fmt.Println("Please add the following lines at the end your shell config file:")
		fmt.Println(aurora.Bold(aurora.BrightCyan(commands.Home)))
		fmt.Println(aurora.Bold(aurora.BrightCyan(commands.Path)))
		os.Exit(0)
	}

	configFile := getShellConfig()
	appendToFile(configFile, "\n# Caddy")
	appendToFile(configFile, "\n"+commands.Home)
	appendToFile(configFile, "\n"+commands.Path)
}

func appendToFile(file string, content string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	if _, err := f.WriteString(content); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func hasSupportedShell() bool {
	shell := os.Getenv("SHELL")
	return shell == "/bin/bash" || shell == "/bin/zsh"
}

func getShellConfig() string {
	shell := os.Getenv("SHELL")
	if shell == "/bin/bash" {
		return filepath.Join(config.HomeDir, ".bashrc")
	} else if shell == "/bin/zsh" {
		return filepath.Join(config.HomeDir, ".zshrc")
	}
	return ""
}
