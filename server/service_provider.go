package server

import (
	"github.com/adolphlwq/webim/mysql"
)

// ServiceProvider object hold service which server need
type ServiceProvider struct {
	MysqlClient *mysql.Client
}
