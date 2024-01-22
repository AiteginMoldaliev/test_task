package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	Dbdriver             string        `mapstructure:"DB_DRIVER"`
	Dbsource             string        `mapstructure:"DB_SOURCE"`
	GinServerAddress     string        `mapstructure:"GIN_SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
