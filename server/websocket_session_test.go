package server

import (
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestNewWebSocketSession(t *testing.T) {
	wsSess := NewWebSocketSession()
	assert.NotNil(t, wsSess)
	assert.NotNil(t, wsSess.mutex)
	assert.NotNil(t, wsSess.wsMap)
}

func TestSet(t *testing.T) {
	wsSess := NewWebSocketSession()
	ws := &websocket.Conn{}
	wsSess.Set("test", ws)
	assert.NotNil(t, wsSess.Get("test"))
	assert.Equal(t, wsSess.HasKey("test"), true)
	assert.Equal(t, wsSess.Get("test"), ws)
}

func TestGet(t *testing.T) {
	wsSess := NewWebSocketSession()
	ws := &websocket.Conn{}
	wsSess.Set("test", ws)
	assert.NotNil(t, wsSess.Get("test"))
	assert.Equal(t, wsSess.Get("test"), ws)
}

func TestDelete(t *testing.T) {
	wsSess := NewWebSocketSession()
	ws := &websocket.Conn{}
	wsSess.Set("test", ws)
	assert.NotNil(t, wsSess.Get("test"))
	assert.Equal(t, wsSess.HasKey("test"), true)
	wsSess.Delete("test")
	assert.Equal(t, wsSess.HasKey("test"), false)
}

func TestHasKey(t *testing.T) {
	wsSess := NewWebSocketSession()
	ws := &websocket.Conn{}
	wsSess.Set("test", ws)
	assert.Equal(t, wsSess.HasKey("test"), true)
	assert.Equal(t, wsSess.HasKey("ws"), false)
}
