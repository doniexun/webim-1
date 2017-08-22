package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/adolphlwq/webim/entity"
	"github.com/gin-gonic/gin"
)

// ContactAdd add contact relationship handler
func ContactAdd(c *gin.Context, appService *ServiceProvider) {
	var contact entity.Contact
	err := c.BindJSON(&contact)
	if err != nil {
		logrus.Warnln("parse contact info from request error: ", err.Error())
		c.JSON(http.StatusBadRequest, CommonResponse{
			Message: ErrMoreContactInfo,
		})
		return
	}

	if contact.FriendMax == "" || contact.FriendMin == "" {
		logrus.Warnln("need friendmax or friendmin of contact.")
		c.JSON(http.StatusBadRequest, CommonResponse{
			Message: ErrMoreContactInfo,
		})
		return
	}

	// check friend name existed
	contact.Standerize()
	isExisted1 := appService.MysqlClient.ChekcUserExistedByUsername(contact.FriendMax)
	if !isExisted1 {
		c.JSON(http.StatusBadRequest, CommonResponse{
			Message: fmt.Sprintf(ErrUsernameNotExistedTmpl, contact.FriendMax),
		})
		return
	}
	isExisted2 := appService.MysqlClient.ChekcUserExistedByUsername(contact.FriendMin)
	if !isExisted2 {
		c.JSON(http.StatusBadRequest, CommonResponse{
			Message: fmt.Sprintf(ErrUsernameNotExistedTmpl, contact.FriendMin),
		})
		return
	}

	// check if contact relationship existed
	var tmpContact entity.Contact
	appService.MysqlClient.DB.Where("friend_min = ? and friend_max = ?", contact.FriendMin, contact.FriendMax).First(&tmpContact)
	if tmpContact.FriendMax != "" {
		c.JSON(http.StatusUnprocessableEntity, CommonResponse{
			Message: ErrContactRealitionshipExisted,
		})
		return
	}

	// contact relationship does not exist, save to db
	contact.AddedTime = time.Now().UTC()
	contact.State = "active"
	appService.MysqlClient.DB.Create(&contact)
	c.JSON(http.StatusOK, CommonResponse{
		Message: ContactAddSuccess,
	})
}
