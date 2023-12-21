package commands

import (
	"caddy/src/shell"
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
)

func Update() *cli.Command {
	return &cli.Command{
		Name:  "update",
		Usage: "Update caddy to the latest version",
		Action: func(cCtx *cli.Context) error {
			cmd := exec.Command(shell.GetShell(), "-c", `curl -s https://raw.githubusercontent.com/mschakulat/caddy/main/ci/install.sh | bash`)
			err := cmd.Run()

			if err != nil {
				fmt.Println(aurora.Yellow("Error during update"))
				os.Exit(0)
			}

			fmt.Println(aurora.Cyan("⛳️ You now have the latest version of Caddy"))
			
			return nil
		},
	}
}
