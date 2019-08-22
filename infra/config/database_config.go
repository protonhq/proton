package config

import (
	"strconv"
	"strings"
)

// DatabaseConfiguration - Database Configuration
type DatabaseConfiguration struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
	SSLMode      string
}

// ConnectionString - Database Connection String
func (c *DatabaseConfiguration) ConnectionString() string {
	strs := []string{
		"host=" + c.Host,
		"port=" + strconv.Itoa(c.Port),
		"user=" + c.User,
		"password=" + c.Password,
		"dbname=" + c.DatabaseName,
		"sslmode=" + c.SSLMode,
	}
	result := strings.Join(strs, " ")
	return result
}
