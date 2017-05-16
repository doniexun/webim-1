package api

import "github.com/gin-gonic/gin"
import "net/http"

// GetUnreadMsg
func GetUnreadMsg(c *gin.Context) {
	receiver := c.Query("receiver")
	msgList, err := im.GetUnreadMsg(receiver)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"data":   err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   *msgList})
	}
}
