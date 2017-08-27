package server

import (
	"fmt"
	"testing"

	"github.com/adolphlwq/webim/entity"
	"github.com/stretchr/testify/assert"
)

func TestMessageUnread(t *testing.T) {
	startTestServer(t)
	truncateTable("contacts")
	truncateTable("users")
	unreadAPI := TestAddr + "/api/v1/message/" + TestUsername + "/unread"
	unreadBadAPI := TestAddr + "/api/v1/message//unread"

	// test with bad api
	m0 := getRequest(unreadBadAPI)
	assert.NotNil(t, m0)
	assert.Equal(t, m0["message"], fmt.Sprintf(ErrParamInPath, "username"))

	// test user not register
	m1 := getRequest(unreadAPI)
	assert.NotNil(t, m1)
	assert.Equal(t, m1["message"], ErrUserNotExisted)

	// test user has registered
	user := entity.User{
		Username: TestUsername,
		Password: encryptPassword,
	}
	insertTestUser(user)
	m2 := getRequest(unreadAPI)
	assert.NotNil(t, m2)
	assert.Equal(t, m2["message"], GetUnreadMessageSuccess)
	assert.NotNil(t, m2["messages"])
}
