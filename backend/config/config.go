package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Version         string `mapstructure:"VERSION"`
	Port            string `mapstructure:"PORT"`
	Environment     string `mapstructure:"ENVIRONMENT"`
	MongoHost       string `mapstructure:"MONGO_HOST"`
	MongoUser       string `mapstructure:"MONGO_USER"`
	MongoPassword   string `mapstructure:"MONGO_PASSWORD"`
	MongoDatabase   string `mapstructure:"MONGO_DATABASE"`
	MongoCollection string `mapstructure:"MONGO_COLLECTION"`
	AllowedOrigin   string `mapstructure:"ALLOWERD_ORIGIN"`
}

var Conf Config

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	Conf = Config{}
	// Read file path
	viper.AddConfigPath(path)
	// set config file and path
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	// watching changes in app.env
	viper.AutomaticEnv()
	// reading the config file
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
