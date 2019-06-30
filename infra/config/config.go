package config

import "github.com/spf13/viper"

const (
	version = "v0.0.1"
)

// Configuration - Application Configuration
type Configuration struct {
	Version  string
	Database DatabaseConfiguration
	Server   ServerConfiguration
}

// NewConfiguration - Read Application Configuration
func NewConfiguration() *Configuration {
	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return &Configuration{
		Version: version,
		Database: DatabaseConfiguration{
			Host:         viper.GetString("database.host"),
			Port:         viper.GetInt("database.port"),
			User:         viper.GetString("database.user"),
			Password:     viper.GetString("database.password"),
			DatabaseName: viper.GetString("database.database"),
		},
		Server: ServerConfiguration{
			Port: viper.GetInt("server.port"),
		},
	}
}
