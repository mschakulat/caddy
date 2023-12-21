package commands

import (
	"caddy/src/config"
	"caddy/src/parser"
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/viper"
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
					if len(id) == 0 {
						fmt.Printf("%s %s\n", aurora.Bold(aurora.Cyan("Identifier")), parser.GetIdentifier())
						os.Exit(0)
					}

					if !isValidString(id) {
						fmt.Println(aurora.Yellow("Identifier not valid"))
						os.Exit(0)
					}

					config.InitConfig()
					viper.Set(config.CaddyConfigKeys.Identifier, id)

					config.WriteConfig()

					fmt.Printf("%s %s\n", aurora.Bold(aurora.Cyan("Identifier")), id)

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
