package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Func to initialize configuration, mainly viper in server
func Initialize() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}