package config

import (
	"github.com/spf13/viper"
)

var App Config

type Config struct {
	AppName             string `mapstructure:"APP_NAME"`
	Port                string `mapstructure:"PORT"`
	JWTSecret           string `mapstructure:"JWT_SECRET"`
	JWTSigningAlgorithm string `mapstructure:"JWT_SIGNING_ALGORITHM"`
	JWTDuration         int    `mapstructure:"JWT_DURATION"`
	PostgresHost        string `mapstructure:"POSTGRES_HOST"`
	PostgresPort        int    `mapstructure:"POSTGRES_PORT"`
	PostgresUser        string `mapstructure:"POSTGRES_USER"`
	PostgresPassword    string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDb          string `mapstructure:"POSTGRES_DB"`
}

func Setup() {
	viper.SetDefault("APP_NAME", "crime-report")
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("JWT_SECRET", "")
	viper.SetDefault("JWT_SIGNING_ALGORITHM", "")
	viper.SetDefault("JWT_DURATION", 1440)
	viper.SetDefault("POSTGRES_HOST", "")
	viper.SetDefault("POSTGRES_PORT", 5432)
	viper.SetDefault("POSTGRES_USER", "")
	viper.SetDefault("POSTGRES_PASSWORD", "")
	viper.SetDefault("POSTGRES_DB", "")

	viper.AddConfigPath("./config")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.ReadInConfig()
	viper.AutomaticEnv()

	err := viper.Unmarshal(&App)
	if err != nil {
		panic("Could not unmarshal config")
	}
}
