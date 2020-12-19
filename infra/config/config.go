package config

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	version = "v0.0.1"
)

// Configuration - Application Configuration
type Configuration struct {
	Version  string
	Database DatabaseConfiguration
	Server   ServerConfiguration
	Pubber   PubberConfiguration
}

// NewConfiguration - Read Application Configuration
func NewConfiguration() *Configuration {
	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err == nil {
		log.Info("Using config file:", viper.ConfigFileUsed())
	} else {
		log.Info("Config file not found")
		viper.SetEnvPrefix("PT")
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AutomaticEnv()
	}

	return &Configuration{
		Version: version,
		Database: DatabaseConfiguration{
			Host:         viper.GetString("database.host"),
			Port:         viper.GetInt("database.port"),
			User:         viper.GetString("database.user"),
			Password:     viper.GetString("database.password"),
			DatabaseName: viper.GetString("database.database"),
			SSLMode:      viper.GetString("database.sslmode"),
		},
		Server: ServerConfiguration{
			Port: viper.GetInt("server.port"),
		},
		Pubber: PubberConfiguration{
			UserAgent:               viper.GetString("pubber.user_agent"),
			MaxDeliveryDepth:        viper.GetInt("pubber.max_delivery_depth"),
			MaxInboxForwardingDepth: viper.GetInt("pubber.max_inbox_forwarding_depth"),
		},
	}
}
