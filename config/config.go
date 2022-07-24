package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Server ServerConfig
	Mongo  MongoConfig
}

type ServerConfig struct {
	Port int
}

type MongoConfig struct {
	URI         string
	Database    string
	Collections MongoCollections
}

type MongoCollections struct {
	Talents string
}

func GetConfig(configFilePath string) (Config, error) {
	viper.SetConfigFile(configFilePath)

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	config := Config{}
	return config, viper.Unmarshal(&config)
}
