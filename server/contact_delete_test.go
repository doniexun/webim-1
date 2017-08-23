package server

import (
	"fmt"
	"testing"
	"time"

	"github.com/adolphlwq/webim/entity"
	"github.com/stretchr/testify/assert"
)

func TestContactDelete(t *testing.T) {
	startTestServer(t)
	truncateTable("contacts")
	truncateTable("users")

	contactDeleteAPI := TestAddr + "/api/v1/contact/delete"
	missedPostData := `{"fmin": "aa"}`
	postData := `{"fmin": "aa", "fmax": "bb"}`

	// test without all needed info
	m1 := deleteRequest(contactDeleteAPI, missedPostData)
	assert.NotNil(t, m1)
	assert.Equal(t, m1["message"], ErrMoreContactInfo)

	// test with contact does not exist
	m2 := deleteRequest(contactDeleteAPI, postData)
	assert.NotNil(t, m2)
	assert.Equal(t, m2["message"], fmt.Sprintf(ErrUsernameNotExistedTmpl, "bb"))

	// fmax and fmin all exist
	fMax := entity.User{
		Username: "bb",
		Password: TestPassword,
	}
	appService.MysqlClient.DB.Create(&fMax)
	fMin := entity.User{
		Username: "aa",
		Password: TestPassword,
	}
	appService.MysqlClient.DB.Create(&fMin)

	// test contact realitionship does not exist
	m3 := deleteRequest(contactDeleteAPI, postData)
	assert.NotNil(t, m3)
	assert.Equal(t, m3["message"], ErrContactRealitionshipNotExisted)

	// test contact realitionship exist and active
	truncateTable("contacts")
	relationship := entity.Contact{
		ID:        1,
		FriendMax: "bb",
		FriendMin: "aa",
		State:     "active",
		AddedTime: time.Now().UTC(),
	}
	appService.MysqlClient.DB.Create(&relationship)
	m4 := deleteRequest(contactDeleteAPI, postData)
	assert.NotNil(t, m4)
	assert.Equal(t, m4["message"], ContactDeleteSuccess)

	// test contact relationship is already inactive
	m5 := deleteRequest(contactDeleteAPI, postData)
	assert.NotNil(t, m5)
	assert.Equal(t, m5["message"], ContactRelationshipIsInactive)

	truncateTable("contacts")
	truncateTable("users")
}
