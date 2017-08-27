package main

import (
	"flag"
	"net"
	"net/http"

	"github.com/adolphlwq/webim/entity"
	"github.com/adolphlwq/webim/mysql"
	"github.com/adolphlwq/webim/server"
	"github.com/gorilla/websocket"
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
	flag.Parse()

	addr := net.JoinHostPort(h, p)
	mysqlClient := mysql.NewMySQLClient(configPath)
	wsUpgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	appService := &server.ServiceProvider{
		MysqlClient:       mysqlClient,
		WebSocketUpgrader: wsUpgrader,
		WebSocketSession:  server.NewWebSocketSession(),
		MessageChan:       make(chan entity.Message, 20),
	}

	server.Start(addr, appService)
}
