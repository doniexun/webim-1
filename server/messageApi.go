package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/adolphlwq/webim/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// GetUnreadMsg
func GetUnreadMsg(c *gin.Context) {
	receiver := c.Query("receiver")
	if receiver == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	msgList, err := im.GetUnreadMsg(receiver)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"data":   err.Error()})
		return
	}

	// update msg state: msg_cache --> msg_done
	// active when prod mode
	// im.UpdateMsgListState(msgList, "msg_done")
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   *msgList})
}

// init variable for further usage
var (
	wsupgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	msgChan = make(chan service.Message, 20)
)

// WSMsgHandler
func WSMsgHandler(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	ws, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logrus.Warnf("Failed to set websocket upgrade: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	logrus.Infof("set up ws between user: %s and server", username)

	// register username and ws
	im.UserWSMap.Set(username, ws)
	logrus.Infof("messageApi: current %s -> ws map is %v", username, ws)

	go im.HandleMsgFromWS(ws, msgChan, username)
	go im.HandleMsgFromMsgChan(msgChan)
}
