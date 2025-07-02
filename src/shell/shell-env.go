package shell

import (
	"caddy/src/config"
	"fmt"
	"os"
	"path/filepath"
)

func HasSupportedShell() bool {
	shell := os.Getenv("SHELL")
	return shell == "/bin/bash" || shell == "/bin/zsh"
}

func GetShellConfig() string {
	shell := os.Getenv("SHELL")
	if shell == "/bin/bash" {
		return filepath.Join(config.HomeDir, ".bash_profile")
	} else if shell == "/bin/zsh" {
		return filepath.Join(config.HomeDir, ".zshrc")
	}
	return ""
}

func GetShell() string {
	shell := os.Getenv("SHELL")
	return filepath.Base(shell)
}

func AppendToFile(file string, content string) {
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