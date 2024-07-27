package config

import (
	"github.com/nawafilhusnul/NAWNAW-API/common/vars"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Name   string
	Host   string
	Port   string
	DB     *Database
	Secret *Secret
}

type Secret struct {
	JWT        string
	Encrypt    string
	AccessExp  int
	RefreshExp int
}

// LoadAppConfig loads the app configuration from the config file
// It reads the configuration values using viper and sets them in the AppConfig struct
// It also sets the global variables in the vars package with the secret values
// Usage example:
//
//	appConfig := LoadAppConfig()
func LoadAppConfig() *AppConfig {
	appCfg := &AppConfig{
		Name: viper.GetString("app.name"),
		Host: viper.GetString("app.host"),
		Port: viper.GetString("app.port"),
		DB:   getDatabaseConfig(),
		Secret: &Secret{
			JWT:        viper.GetString("secret.jwt"),
			Encrypt:    viper.GetString("secret.encrypt"),
			AccessExp:  viper.GetInt("secret.access_expired"),
			RefreshExp: viper.GetInt("secret.refresh_expired"),
		},
	}

	vars.JWT_SECRET = appCfg.Secret.JWT
	vars.ENCRYPT_SECRET = appCfg.Secret.Encrypt
	vars.ACCESS_EXPIRED = appCfg.Secret.AccessExp
	vars.REFRESH_EXPIRED = appCfg.Secret.RefreshExp

	return appCfg
}
