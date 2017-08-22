package server

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRegister(t *testing.T) {
	startTestServer(t)
	userRegisterAPI := TestAddr + "/api/v1/user/register"

	data := `{"username": "test","password": "123456"}`
	m1 := postJSON(userRegisterAPI, strings.NewReader(data))
	assert.NotNil(t, m1)
	assert.Equal(t, m1["message"], UserRegisterSuccess)

	// test user has existed
	m2 := postJSON(userRegisterAPI, strings.NewReader(data))
	assert.NotNil(t, m2)
	assert.Equal(t, m2["message"], ErrUserExisted)
}
