package api

import (
	"webim/db"

	"net/http"

	"encoding/json"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// AddFriend add friend pair to db
func AddFriend(c *gin.Context) {
	var fpair db.AddFriend
	c.BindJSON(&fpair)

	err := us.AddFriend(fpair)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": 400, "data": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 200})
	}
}

// ListFriend list all friends of one user
func ListFriend(c *gin.Context) {
	username := c.Query("username")
	friendList, err := us.ListFriend(username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"data":   err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   friendList})

	}
}

func jsonify(input interface{}) (string, error) {
	retJSON, err := json.Marshal(input)
	if err != nil {
		logrus.Warnf("jsonify info error")
		return "", err
	}

	return string(retJSON), nil
}
