package util

import (
	"time"

	"github.com/Sirupsen/logrus"
	"golang.org/x/crypto/scrypt"
)

const PASS_SALT = "WebIM@adolph*lwq!"

// GetCurrentTimestamp return current utc timestamp
func GetCurrentTimestamp() int64 {
	return time.Now().UTC().Unix()
}

// EncryptPassword encrypt password
func EncryptPassword(password string) string {
	dk, err := scrypt.Key([]byte(password), []byte(PASS_SALT),
		16384, 8, 1, 32)
	if err != nil {
		logrus.Fatal(err)
	}

	return string(dk)
}
