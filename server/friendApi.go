package api

import (
	"net/http"

	"github.com/adolphlwq/webim/service"
	"github.com/gin-gonic/gin"
)

// AddFriendRelationship add friend pair to db
func AddFriendRelationship(c *gin.Context) {
	var fpair service.FriendRelationship
	c.BindJSON(&fpair)

	err := im.AddFriendRelationship(fpair)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": 400, "data": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 200})
	}
}

// ListFriendRelationship list all friends of one user
func ListFriendRelationship(c *gin.Context) {
	username := c.Query("username")
	friendList, err := im.ListFriendRelationship(username)
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

// DeleteFriendRelationship
func DeleteFriendRelationship(c *gin.Context) {
	var fpair service.FriendRelationship
	c.BindJSON(&fpair)

	err := im.DeleteFriendRelationship(fpair)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"data":   "delete friend error, try later."})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   "delete success."})
	}
}
