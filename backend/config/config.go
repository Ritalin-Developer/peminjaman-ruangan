package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Version          string `mapstructure:"VERSION"`
	Port             string `mapstructure:"PORT"`
	Environment      string `mapstructure:"ENVIRONMENT"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDatabase string `mapstructure:"POSTGRES_DATABASE"`
	GORMLog          bool   `mapstructure:"GORM_LOG"`
	PostgresMinConn  int    `mapstructure:"POSTGRES_MIN_CONN"`
	PostgresMaxConn  int    `mapstructure:"POSTGRES_MAX_CONN"`
	AllowedOrigin    string `mapstructure:"ALLOWERD_ORIGIN"`
	SecretKey        string `mapstructure:"SECRET_KEY"`
	TokenLifetimeMin int    `mapstructure:"TOKEN_LIFETIME_MIN"`
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
