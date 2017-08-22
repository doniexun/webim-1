package server

import (
	"fmt"
	"strings"
	"testing"

	"github.com/adolphlwq/webim/entity"
	"github.com/stretchr/testify/assert"
)

func TestContactAdd(t *testing.T) {
	startTestServer(t)
	truncateTable("contacts")
	truncateTable("users")

	contactAddAPI := TestAddr + "/api/v1/contact/add"
	missedPostData := `{"fmin": "aa"}`
	postData := `{"fmin": "aa", "fmax": "bb"}`

	// test without all needed info
	m1 := postJSON(contactAddAPI, strings.NewReader(missedPostData))
	assert.NotNil(t, m1)
	assert.Equal(t, m1["message"], ErrMoreContactInfo)

	// test with contact does not exist
	m2 := postJSON(contactAddAPI, strings.NewReader(postData))
	assert.NotNil(t, m2)
	assert.Equal(t, m2["message"], fmt.Sprintf(ErrUsernameNotExistedTmpl, "bb"))

	// test with fmax and fmin all exist
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
	m3 := postJSON(contactAddAPI, strings.NewReader(postData))
	assert.NotNil(t, m3)
	assert.Equal(t, m3["message"], ContactAddSuccess)

	truncateTable("contacts")
	truncateTable("users")
}
