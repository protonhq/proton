package config

import "strings"

// DatabaseConfiguration - Database Configuration
type DatabaseConfiguration struct {
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string
}

// ConnectionString - Database Connection String
func (c DatabaseConfiguration) ConnectionString string {
	strs := []string{
		"host=" + c.Host,
		"port=" + c.Port,
		"user=" + c.User,
		"password=" + c.Password,
		"dbname=" + c.DatabaseName,
	}
	return strings.Join(strs, " ")
}