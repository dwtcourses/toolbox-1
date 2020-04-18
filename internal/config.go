package internal

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigName("toolbox")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/toolbox/")
	viper.AddConfigPath("$HOME/.toolbox")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.AutomaticEnv()
}
