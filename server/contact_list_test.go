package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContactList(t *testing.T) {
	startTestServer(t)
	contactListAPI := TestAddr + "/api/v1/contact/list"

	m := getRequest(contactListAPI)
	assert.NotNil(t, m)
	assert.Equal(t, m["message"], ContactRelationshipListSuccess)
}
