package setup

import (
	"caddy/src/config"
	"fmt"
	"os"
	"reflect"
)

func InitSystemPath() {
	systemPaths := reflect.ValueOf(config.SystemPaths)

	for i := 0; i < systemPaths.NumField(); i++ {
		dir := systemPaths.Field(i).String()
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			errDir := os.MkdirAll(dir, 0755)
			if errDir != nil {
				fmt.Println("Error creating directory:", errDir)
			}
		}
	}
}
