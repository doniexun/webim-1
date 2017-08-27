package server

import (
	"sync"

	"github.com/gorilla/websocket"
)

// WebSocketSession concurrency safe map
type WebSocketSession struct {
	wsMap map[string]*websocket.Conn
	mutex *sync.Mutex
}

// NewWebSocketSession return new pointer of WebSocketSession
func NewWebSocketSession() *WebSocketSession {
	wsSess := &WebSocketSession{
		wsMap: make(map[string]*websocket.Conn),
		mutex: &sync.Mutex{},
	}

	return wsSess
}

// Set set v of k
func (wsSess *WebSocketSession) Set(k string, v *websocket.Conn) {
	wsSess.mutex.Lock()
	defer wsSess.mutex.Unlock()
	wsSess.wsMap[k] = v
}

// Get return v of k
func (wsSess *WebSocketSession) Get(k string) *websocket.Conn {
	return wsSess.wsMap[k]
}

// Delete delete v of k
func (wsSess *WebSocketSession) Delete(k string) {
	wsSess.mutex.Lock()
	defer wsSess.mutex.Unlock()
	delete(wsSess.wsMap, k)
}

// HasKey if session has k, true has false or not
func (wsSess *WebSocketSession) HasKey(k string) bool {
	wsSess.mutex.Lock()
	defer wsSess.mutex.Unlock()
	_, ok := wsSess.wsMap[k]
	return ok
}
