package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserList(t *testing.T) {
	startTestServer(t)
	userListAPI := TestAddr + "/api/v1/user/list"

	m := getRequest(userListAPI)
	assert.NotNil(t, m)
	assert.Equal(t, m["message"], UserListSuccess)
}
