package api

import (
	"net/http"
	"webim/db"

	"github.com/Sirupsen/logrus"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister handle user register
func UserRegister(c *gin.Context) {
	var user db.User
	c.BindJSON(&user)

	err := us.UserRegister(user.Username, user.Password)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"status": 200,
			"data": "register successfully, please login."})
	} else {
		logrus.Warn("register user info error: ", err)
		c.JSON(http.StatusOK, gin.H{"status": 400, "data": err.Error()})
	}
}

// UserLogin handle user login
func UserLogin(c *gin.Context) {
	var user db.User
	c.BindJSON(&user)

	err := us.UserLogin(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": 400, "data": err.Error()})
		return
	}

	// handle session
	session := sessions.Default(c)
	session.Set(user.Username, user.Username)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user.Username})
}
