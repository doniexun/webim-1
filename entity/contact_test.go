package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStandarize(t *testing.T) {
	c := &Contact{
		FriendMax: "a",
		FriendMin: "b",
	}
	c.Standerize()
	assert.Equal(t, c.FriendMax, "b")
	assert.Equal(t, c.FriendMin, "a")
}
