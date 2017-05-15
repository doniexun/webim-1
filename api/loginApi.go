package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister handle user register
func UserRegister(c *gin.Context) {
	var user db.User
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
	var user db.User
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
	var user db.User
	c.BindJSON(&user)

	session := sessions.Default(c)
	username := session.Get(user.Username)
	if username == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   "logout success"})
		return
	} else {
		session.Delete(user.Username)
		session.Save()
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   "logout success"})
		return
	}
}
