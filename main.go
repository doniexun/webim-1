package main

import (
	"flag"
	"net"

	"github.com/adolphlwq/webim/mysql"
	"github.com/adolphlwq/webim/server"
)

func main() {
	var (
		configPath string
		h          string
		p          string
	)
	flag.StringVar(&configPath, "config", "./default.properties", "config file path")
	flag.StringVar(&h, "h", "0.0.0.0", "host server bind")
	flag.StringVar(&p, "p", "9066", "port server bind")

	addr := net.JoinHostPort(h, p)
	mysqlClient := mysql.NewMySQLClient(configPath)
	appService := &server.ServiceProvider{
		MysqlClient: mysqlClient,
	}

	server.Start(addr, appService)
}
