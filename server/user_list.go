package server

import (
	"net/http"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/adolphlwq/webim/entity"
	"github.com/gin-gonic/gin"
)

// UserList list all user
func UserList(c *gin.Context, appService *ServiceProvider) {
	var (
		limit    string
		limitInt int64
	)

	limit = c.Query("limit")
	if limit == "" {
		limit = "10"
	}
	limitInt, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		logrus.Warnf("parse int error: %v\n", err)

		c.JSON(http.StatusInternalServerError, CommonResponse{
			Message: ErrInternel,
		})
		return
	}

	var users []entity.User
	appService.MysqlClient.DB.Select("id, username").Limit(limitInt).First(&users)
	c.JSON(http.StatusOK, gin.H{
		"message": UserListSuccess,
		"users":   users,
	})
}
