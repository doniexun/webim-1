package server

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/adolphlwq/webim/entity"
	"github.com/gin-gonic/gin"
)

// MessageUnread get unread message of user
func MessageUnread(c *gin.Context, appService *ServiceProvider) {
	username := c.Param("username")
	if username == "" {
		logrus.Warnln(fmt.Sprintf(ErrParamInPath, "username"))
		c.JSON(http.StatusBadRequest, CommonResponse{
			Message: fmt.Sprintf(ErrParamInPath, "username"),
		})
		return
	}

	// check if user register
	if !appService.MysqlClient.ChekcUserExistedByUsername(username) {
		// user does not register
		logrus.Warnln("user does not exist in db.")
		c.JSON(http.StatusUnprocessableEntity, CommonResponse{
			Message: ErrUserNotExisted,
		})
		return
	}

	// get unread message of username
	var messages []entity.Message
	appService.MysqlClient.DB.Where("receiver = ? AND state = ?", username, "msg_cache").Find(&messages)
	c.JSON(http.StatusOK, gin.H{
		"message":  GetUnreadMessageSuccess,
		"messages": messages,
	})
}
