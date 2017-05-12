package service

import (
	"fmt"
	"webim/db"

	"github.com/Sirupsen/logrus"
	"golang.org/x/crypto/scrypt"
)

const (
	PASS_SALT = "AES$%x09^@s!d:<X"
)

type UserService struct {
	dbs *db.DBService
}

func NewUserService(dbs *db.DBService) *UserService {
	us := &UserService{
		dbs: dbs,
	}
	return us
}

// UserLogin
func (us *UserService) UserLogin() string {
	return "user login"
}

// UserRegister
func (us *UserService) UserRegister(username, password string) error {
	err := us.ValidateUser(username, password)
	if err != nil {
		return err
	}

	encryptpass := us.EncryptPass(password)
	err = us.dbs.InsertUser(username, encryptpass)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser check if username is existed
// check username and password validation
func (us *UserService) ValidateUser(username, password string) error {
	if userLen := len(username); userLen < 1 || userLen > 20 {
		return fmt.Errorf("username is too short or too long. it should between [1,20]")
	}

	if passLen := len(password); passLen < 6 || passLen > 20 {
		return fmt.Errorf("password is too short or too long. it should between [6,20]")
	}

	isExist := us.dbs.CheckUser(username)
	if isExist {
		return fmt.Errorf("username has existed, login directly")
	}

	return nil
}

// EncryptPass use salt encrypt pass
func (us *UserService) EncryptPass(password string) string {
	dk, err := scrypt.Key([]byte(password), []byte(PASS_SALT),
		16384, 8, 1, 32)
	if err != nil {
		logrus.Fatal(err)
	}

	return string(dk)
}
