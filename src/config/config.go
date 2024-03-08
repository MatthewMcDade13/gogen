package config

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

const (
	CONFIG_NAME         = "gogen_conf"
	CONFIG_PATH         = "$HOME/." + CONFIG_NAME
	CONFIG_FIELD_PREFIX = "GoModPrefix"
)

var default_config = fmt.Sprint("{\n", "    \"", CONFIG_FIELD_PREFIX, "\": ", `""`, "\n", "}")

func init() {
	viper.SetDefault(CONFIG_FIELD_PREFIX, "")

	viper.SetConfigName(CONFIG_NAME)
	viper.SetConfigType("json")
	viper.AddConfigPath(CONFIG_PATH)
	viper.AddConfigPath(".")

	// if _, err := os.Stat(CONFIG_PATH); err != nil {
	// 	// config file does not exist
	// 	if err := os.WriteFile(CONFIG_PATH, []byte(default_config), 0666); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			path := os.ExpandEnv(CONFIG_PATH)
			if err := os.WriteFile(path, []byte(default_config), 0666); err != nil {
				log.Fatal(err)
			}

		} else {
			log.Fatal(err)
		}

	}
}

func GoModPrefix() string {
	return viper.Get(CONFIG_FIELD_PREFIX).(string)
}

func Set[T any](field string, val T) {
	viper.Set(field, val)
}

func Save() {
	viper.WriteConfig()
}
