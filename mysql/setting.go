package mysql

import (
	"github.com/adolphlwq/webim/config"
	"github.com/magiconair/properties"
)

// MySQLConfig wrape mysql setting
type MySQLConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

// NewMySQLConfigFromConfig return new mysql setting from properties.Properties
func NewMySQLConfigFromConfig(props *properties.Properties) *MySQLConfig {
	ms := &MySQLConfig{
		Host:     props.MustGetString("mysql.host"),
		Port:     props.MustGetString("mysql.port"),
		Database: props.MustGetString("mysql.database"),
		User:     props.MustGetString("mysql.user"),
		Password: props.MustGetString("mysql.password"),
	}
	return ms
}

// NewMySQLConfig return new mysql setting from config file path
func NewMySQLConfig(configPath string) *MySQLConfig {
	props := config.ReadProps(configPath)
	return NewMySQLConfigFromConfig(props)
}
