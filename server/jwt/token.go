package jwt

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/adolphlwq/webim/entity"
	"github.com/dgrijalva/jwt-go"
)

// TokenKey default token key
const TokenKey = "webim&^%Go#@Slang"

// GenerateToken generate jwt token for user
func GenerateToken(user *entity.User) string {
	jwtToken := jwt.New(jwt.SigningMethodHS256)
	claims := jwtToken.Claims.(jwt.MapClaims)
	claims["user"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	tokenString, err := jwtToken.SignedString([]byte(TokenKey))
	if err != nil {
		logrus.Fatalf("generate token string error: %v\n", err)
	}

	return tokenString
}

// ParseToken parse token string and return jwt.Token
func ParseToken(tokenString string) *jwt.Token {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(TokenKey), nil
	})

	return token
}
