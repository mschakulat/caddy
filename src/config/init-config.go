package config

import (
	"github.com/logrusorgru/aurora"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type caddyConfigKeysStruct struct {
	Identifier string
}

var CaddyConfigKeys = caddyConfigKeysStruct{
	Identifier: "caddy.identifier",
}

var CaddyConfigFilePath = filepath.Join(SystemPaths.Base, "caddy.yaml")

func InitConfig() {
	viper.SetConfigName("caddy")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(SystemPaths.Base)
	
	viper.SetDefault(CaddyConfigKeys.Identifier, PackageJsonIdentifier)
	
	
	
	if err := viper.ReadInConfig(); err != nil {
		err = os.WriteFile(CaddyConfigFilePath, nil, 0644)
		if err != nil {
			println(aurora.Yellow("Could not create config file"))
			return
		}
		
		WriteConfig()
	}
}

func WriteConfig() {
	c := viper.AllSettings()
	
	bs, err := yaml.Marshal(c)
	if err != nil {
		println(aurora.Yellow("Unable to marshal config to YAML"))
		os.Exit(0)
	}
	
	err = os.WriteFile(CaddyConfigFilePath, bs, 0644)
	if err != nil {
		println(aurora.Yellow("Unable to write config to file"))
		os.Exit(0)
	}
}