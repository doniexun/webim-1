package server

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/adolphlwq/webim/entity"
	"github.com/gin-gonic/gin"
)

// ContactDelete delete contact handler
func ContactDelete(c *gin.Context, appService *ServiceProvider) {
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

	// check if contact relationship exists in db
	var tmpContact entity.Contact
	appService.MysqlClient.DB.Where("friend_min = ? and friend_max = ?",
		contact.FriendMin, contact.FriendMax).First(&tmpContact)
	if tmpContact.FriendMax == "" {
		logrus.Warnf("relationship between %s and %s does not exist.",
			contact.FriendMin, contact.FriendMax)
		c.JSON(http.StatusUnprocessableEntity, CommonResponse{
			Message: ErrContactRealitionshipNotExisted,
		})
		return
	}

	// change relationship state
	appService.MysqlClient.DB.Model(&tmpContact).Update("state", "inactive")
	c.JSON(http.StatusOK, CommonResponse{
		Message: ContactDeleteSuccess,
	})
}
