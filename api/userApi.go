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
		logrus.Infof("err when login is: %v", err)
		c.JSON(http.StatusOK, gin.H{"status": 400, "data": err.Error()})
	} else {
		// handle session
		session := sessions.Default(c)
		session.Set(user.Username, user.Username)
		// do not forget to save session
		session.Save()
		logrus.Infof("register %s to session", session.Get(user.Username))
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   user.Username})
	}
}

// LogOut handle logout
func LogOut(c *gin.Context) {
	var user service.User
	c.BindJSON(&user)

	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	logrus.Infof("user to logout is %v", user.Username)

	// TODO clean user websocket
	// isWS true user has websocket false or not
	// if isWS := im.UserWSMap.HasKey(user.Username); isWS {
	// 	ws := im.UserWSMap.Get(user.Username)
	// 	err := ws.Close()
	// 	if err != nil {
	// 		logrus.Fatalf("close websocket error: %v", err)
	// 	}
	// 	im.UserWSMap.Delete(user.Username)
	// }
	//  close websocket will be done in im.HandleMsgFromWS

	session := sessions.Default(c)
	username := session.Get(user.Username)
	logrus.Infof("username in session is %v", username)
	if username == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   "logout success"})
		logrus.Infof("username %s does not exists in session, return directly", user.Username)
	} else {
		session.Delete(user.Username)
		session.Save()
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   "logout success"})
		logrus.Infof("remove username %s in session, return directly", user.Username)
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
