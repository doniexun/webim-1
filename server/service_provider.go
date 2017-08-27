package server

import (
	"github.com/adolphlwq/webim/entity"
	"github.com/adolphlwq/webim/mysql"
	"github.com/gorilla/websocket"
)

// ServiceProvider object hold service which server need
type ServiceProvider struct {
	MysqlClient       *mysql.Client
	WebSocketUpgrader *websocket.Upgrader
	WebSocketSession  *WebSocketSession
	MessageChan       chan entity.Message
}
