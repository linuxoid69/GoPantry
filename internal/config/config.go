package config

import (
	"fmt"
	"path"

	"github.com/spf13/viper"
)

func LoadConfig(confFilePath string) {
	cfgDir, cfgFile := path.Split(confFilePath)
	viper.SetConfigFile(cfgFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(cfgDir)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
