package setup

import (
	"caddy/src/shell"
	"fmt"
	"github.com/logrusorgru/aurora"
	"os"
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

	if !shell.HasSupportedShell() {
		fmt.Println(aurora.Bold(aurora.Yellow("Unsupported shell")))
		fmt.Println("Only bash and zsh are supported at the moment")
		fmt.Println("Please add the following lines at the end your shell config file:")
		fmt.Println(aurora.Bold(aurora.BrightCyan(commands.Home)))
		fmt.Println(aurora.Bold(aurora.BrightCyan(commands.Path)))
		os.Exit(0)
	}

	configFile := shell.GetShellConfig()
	shell.AppendToFile(configFile, "\n# Caddy")
	shell.AppendToFile(configFile, "\n"+commands.Home)
	shell.AppendToFile(configFile, "\n"+commands.Path)
}