package mysql

import (
	"testing"

	"github.com/adolphlwq/webim/config"
	"github.com/stretchr/testify/assert"
)

var configPath = "../test.properties"

func TestNewMySQLConfig(t *testing.T) {
	MySQLConfig := NewMySQLConfig(configPath)
	assert.NotNil(t, MySQLConfig)
	assert.Equal(t, MySQLConfig.Database, "test_webim")
}

func TestNewMySQLConfigFromConfig(t *testing.T) {
	props := config.ReadProps(configPath)
	MySQLConfig := NewMySQLConfigFromConfig(props)
	assert.NotNil(t, MySQLConfig)
	assert.Equal(t, MySQLConfig.Database, "test_webim")
}
