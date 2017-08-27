package server

import (
	"testing"

	"strings"

	"github.com/adolphlwq/webim/entity"
	"github.com/stretchr/testify/assert"
)

func TestUserLogin(t *testing.T) {
	startTestServer(t)
	truncateTable("users")
	userLoginAPI := TestAddr + "/api/v1/user/login"

	// test not enough info
	data := `{"username": "test"}`
	m1 := postJSON(userLoginAPI, strings.NewReader(data))
	assert.NotNil(t, m1)
	assert.Equal(t, m1["message"], ErrMoreUserInfo)

	// test user not resiger
	data1 := `{"username": "test","password": "123456"}`
	m2 := postJSON(userLoginAPI, strings.NewReader(data1))
	assert.NotNil(t, m2)
	assert.Equal(t, m2["message"], ErrUserNotExisted)

	// insert test data
	user := entity.User{
		Username: TestUsername,
		Password: encryptPassword,
	}
	insertTestUser(user)
	m3 := postJSON(userLoginAPI, strings.NewReader(data1))
	assert.NotNil(t, m3)
	assert.Equal(t, m3["message"], UserLoginSuccess)
	token := m3["token"].(string)
	if len(token) <= 0 {
		t.Errorf("bad token in login request response.")
	}

	truncateTable("users")
}
