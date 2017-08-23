package jwt

import (
	"testing"

	"github.com/adolphlwq/webim/entity"
)

func TestGenerateToken(t *testing.T) {
	user := &entity.User{
		Username: "test",
		Password: "test",
	}
	token := GenerateToken(user)
	if token == "" {
		t.Error()
	}
}

func TestParseToken(t *testing.T) {
	user := &entity.User{
		Username: "test",
		Password: "test",
	}
	tokenString := GenerateToken(user)

	jwtToken := ParseToken(tokenString)

	if !jwtToken.Valid {
		t.Error()
	}
}
