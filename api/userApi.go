package api

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/adolphlwq/webim/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister handle user register
func UserRegister(c *gin.Context) {
	var user service.User
	c.BindJSON(&user)

	err := im.UserRegister(user.Username, user.Password)
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
	var user service.User
	c.BindJSON(&user)

	err := im.UserLogin(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": 400, "data": err.Error()})
	} else {
		// handle session
		session := sessions.Default(c)
		session.Set(user.Username, user.Username)
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   user.Username})
	}
}

// LoginOut handle logout
func LoginOut(c *gin.Context) {
	var user service.User
	c.BindJSON(&user)

	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	logrus.Infof("user to logout is %v", user)

	// TODO clean user websocket
	// isWS true user has websocket false or not
	go func(im *service.IMService, username string) {
		if isWS := im.UserWSMap.HasKey(username); isWS {
			ws := im.UserWSMap.Get(username)
			err := ws.Close()
			if err != nil {
				logrus.Fatalf("close websocket error: %v", err)
			}
			im.UserWSMap.Delete(username)
		}
	}(im, user.Username)

	session := sessions.Default(c)
	username := session.Get(user.Username)
	if username == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   "logout success"})
		logrus.Info("username is nil, return directly")
	} else {
		session.Delete(user.Username)
		session.Save()
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   "logout success"})
		logrus.Info("remove username in session, return directly")
	}
}

// GetUserByName
func GetUserByName(c *gin.Context) {
	username := c.Query("username")
	fmt.Println(username)
	if username == "" || len(username) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"data":   ""})
		return
	}

	user, err := im.GetUserByName(username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"data":   err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   *user})
}
