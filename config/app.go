package config

import "github.com/spf13/viper"

type AppConfig struct {
	Name string
	Host string
	Port string
	DB   *Database
}

// LoadAppConfig loads the app configuration from the config file
func LoadAppConfig() *AppConfig {
	return &AppConfig{
		Name: viper.GetString("app.name"),
		Host: viper.GetString("app.host"),
		Port: viper.GetString("app.port"),
		DB:   getDatabaseConfig(),
	}
}
