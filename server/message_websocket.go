package server

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/adolphlwq/webim/entity"
	webimjwt "github.com/adolphlwq/webim/server/jwt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// MessageWebSocket set connection on websocket and handle message
func MessageWebSocket(c *gin.Context, appService *ServiceProvider) {
	tokenString := c.Param("token")
	if tokenString == "" {
		logrus.Warnln(ErrNeedUserToken)
		c.JSON(http.StatusUnprocessableEntity, CommonResponse{
			Message: ErrNeedUserToken,
		})
		return
	}

	// parse and valid token
	token := webimjwt.ParseToken(tokenString)
	if !token.Valid {
		logrus.Warnln(ErrJWTInvalid)
		c.JSON(http.StatusUnauthorized, CommonResponse{
			Message: ErrJWTInvalid,
		})
		return
	}
	claim := token.Claims.(jwt.MapClaims)
	username := claim["user"].(string)
	if username == "" {
		c.JSON(http.StatusUnprocessableEntity, CommonResponse{
			Message: ErrJWTInvalid,
		})
		return
	}

	// upgrade to websocket
	ws, err := appService.WebSocketUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logrus.Warnln(ErrWebSocketUpgrade, err)
		c.JSON(http.StatusInternalServerError, CommonResponse{
			Message: ErrWebSocketUpgrade,
		})
		return
	}

	// save ws to session
	if appService.WebSocketSession.HasKey(username) {
		logrus.Warnln(ErrUserWebSocketExisted)
		c.JSON(http.StatusUnprocessableEntity, CommonResponse{
			Message: ErrUserWebSocketExisted,
		})
		return
	}
	appService.WebSocketSession.Set(username, ws)
	go HandleMessageFromWebSocket(appService, username)
	go HandleMessageFromChannel(appService)
}

// HandleMessageFromWebSocket get message from channel and handle it according to sender and receiver state
func HandleMessageFromWebSocket(appService *ServiceProvider, username string) {
	ws := appService.WebSocketSession.Get(username)
	for {
		var msg entity.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			logrus.Warnln(ErrReadMessageFromWebSocket, err)
			return
		}
		appService.MessageChan <- msg
	}
}

// HandleMessageFromChannel read message from websocket and send to message channel
func HandleMessageFromChannel(appService *ServiceProvider) {
	for {
		msg := <-appService.MessageChan
		// check receiver login status, use WebSocket for now
		// will change in new version
		if appService.WebSocketSession.HasKey(msg.Receiver) {
			msg.State = "msg_receiver"
			recevierWS := appService.WebSocketSession.Get(msg.Receiver)
			err := recevierWS.WriteJSON(msg)
			if err != nil {
				logrus.Warnln(ErrWriteMessageToReveiver, err)
				msg.State = "msg_cache"
			}
			// save message to db
			appService.MysqlClient.DB.Create(&msg)
		} else {
			msg.State = "msg_cache"
			appService.MysqlClient.DB.Create(&msg)
		}
	}
}
