package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`      //names as in .env file
	DBSource      string `mapstructure:"DB_SOURCE"`      //names as in .env file
	ServerAddress string `mapstructure:"SERVER_ADDRESS"` //names as in .env file
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app") // name of config file (without extension)
	viper.SetConfigType("env") // REQUIRED if the config file does not have the extension in the name
	viper.AutomaticEnv()       // read in environment variables that match
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.Unmarshal(&config)
	return
}
