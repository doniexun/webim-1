package server

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/adolphlwq/webim/entity"
	"github.com/adolphlwq/webim/server/jwt"
	"github.com/adolphlwq/webim/util"
	"github.com/gin-gonic/gin"
)

// UserLogin user login handler
func UserLogin(c *gin.Context, appService *ServiceProvider) {
	var user entity.User
	err := c.BindJSON(&user)
	if err != nil {
		logrus.Warnln("parse user info from login request error: ", err.Error())
		c.JSON(http.StatusBadRequest, CommonResponse{
			Message: ErrMoreUserInfo,
		})
		return
	}

	if user.Username == "" || user.Password == "" {
		logrus.Warnln("need username or password of user.")
		c.JSON(http.StatusBadRequest, CommonResponse{
			Message: ErrMoreUserInfo,
		})
		return
	}

	// check if user existed
	user.Password = util.EncryptPasswordWithSalt(user.Username, user.Password)
	var tmpUser entity.User
	appService.MysqlClient.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&tmpUser)
	if tmpUser.Username == "" {
		logrus.Warnln("user does not exist in db.")
		c.JSON(http.StatusUnprocessableEntity, CommonResponse{
			Message: ErrUserNotExisted,
		})
		return
	}

	token := jwt.GenerateToken(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": UserLoginSuccess,
		"token":   token,
	})
}
