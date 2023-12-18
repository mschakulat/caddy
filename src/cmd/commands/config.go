package commands

import (
	"caddy/src/config"
	"caddy/src/shell"
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/urfave/cli/v2"
	"os"
	"regexp"
)

func Config() *cli.Command {
	return &cli.Command{
		Name:  "config",
		Usage: "Configure caddy",
		Subcommands: []*cli.Command{
			{
				Name:  "id",
				Usage: "Configure the identifier in the package.json (defaults to 'caddy')",
				Action: func(cCtx *cli.Context) error {
					id := cCtx.Args().Get(0)
					if len(id) == 0 || !isValidString(id) {
						fmt.Println(aurora.Yellow("Please provide an identifier"))
						os.Exit(0)
					}

					configFile := shell.GetShellConfig()
					identifier := os.Getenv(config.CaddyEnvId)
					if len(identifier) != 0 {
						fmt.Println(aurora.Cyan("Identifier already set: " + identifier))
						fmt.Printf("To change or remove it, please edit your shell config manually (%s)\n", configFile)
						os.Exit(0)
					}

					shell.AppendToFile(configFile, "\nexport "+config.CaddyEnvId+"=\""+id+"\"")

					fmt.Println(aurora.Cyan("Identifier added to your env variables"))
					fmt.Println("Please reload your shell config or open a new terminal")

					return nil
				},
			},
		},
	}
}

func isValidString(str string) bool {
	reg, _ := regexp.Compile("^[a-zA-Z0-9-_]*$")
	return reg.MatchString(str)
}
