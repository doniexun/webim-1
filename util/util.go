package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"

	"github.com/Sirupsen/logrus"
	"golang.org/x/crypto/scrypt"
)

const (
	PASS_SALT1 = "WebIM@adolph*lwq!"
	PASS_SALT2 = "Love^%@Go+_*lang"
)

// GetCurrentTimestamp return current utc timestamp
func GetCurrentTimestamp() int64 {
	return time.Now().UTC().Unix()
}

// EncryptPassword encrypt password
func EncryptPassword(password string) string {
	dk, err := scrypt.Key([]byte(password), []byte(PASS_SALT1),
		16384, 8, 1, 32)
	if err != nil {
		logrus.Fatal(err)
	}

	return string(dk)
}

// EncryptPasswordWithSalt encrypt password with salt
func EncryptPasswordWithSalt(username, password string) string {
	h := md5.New()
	io.WriteString(h, password)
	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))

	io.WriteString(h, PASS_SALT1)
	io.WriteString(h, username)
	io.WriteString(h, PASS_SALT2)
	io.WriteString(h, pwmd5)

	last := fmt.Sprintf("%x", h.Sum(nil))
	return last
}
