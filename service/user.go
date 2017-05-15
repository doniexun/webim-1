package service

import (
	"database/sql"
	"fmt"

	"github.com/Sirupsen/logrus"
	"golang.org/x/crypto/scrypt"
)

const (
	PASS_SALT = "AES$%x09^@s!d:<X"
)

type IMService struct {
	dbs *DBService
}

func NewIMService(dbs *DBService) *IMService {
	im := &IMService{
		dbs: dbs,
	}
	return im
}

// UserLogin
func (im *IMService) UserLogin(username, password string) error {
	err := im.userLoginValidate(username, password)
	if err != nil {
		return err
	}

	err = im.CheckUserPass(username, im.encryptPass(password))
	if err != nil { // login err
		return err
	}

	// handle session in api layer
	return nil
}

// userLoginValidate validate user login info
func (im *IMService) userLoginValidate(username, password string) error {
	if userLen := len(username); userLen < 1 || userLen > 20 {
		return fmt.Errorf("username is too short or too long. it should between [1,20]")
	}

	if passLen := len(password); passLen < 6 || passLen > 20 {
		return fmt.Errorf("password is too short or too long. it should between [6,20]")
	}

	isExist := im.CheckUser(username)
	if !isExist {
		return fmt.Errorf("username does not exist, please register first.")
	}

	return nil
}

// CheckUser check if user exists. true exist or false not exist
func (im *IMService) CheckUser(username string) bool {
	db := im.dbs.CreateDB()
	defer db.Close()

	stmt := im.dbs.STMTFactory("SELECT username FROM user WHERE username=?", db)
	defer stmt.Close()

	var username_ string
	err := stmt.QueryRow(username).Scan(&username_)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			logrus.Fatal("query username ", username, " error: ", err)
		}
	}

	return true
}

// CheckUserPass check if user and pass exist in db
func (im *IMService) CheckUserPass(username, encryptPass string) error {
	db := im.dbs.CreateDB()
	defer db.Close()

	stmt := im.dbs.STMTFactory("SELECT username FROM user WHERE username=? "+
		" AND password=?", db)
	defer stmt.Close()

	var username_ string
	err := stmt.QueryRow(username, encryptPass).Scan(&username_)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("username or password error")
		} else {
			logrus.Warn("query username ", username, " error: ", err)
			return fmt.Errorf("please try again")
		}
	}

	return nil
}

// UserRegister
func (im *IMService) UserRegister(username, password string) error {
	err := im.userRegValidate(username, password)
	if err != nil {
		return err
	}

	encryptpass := im.encryptPass(password)
	err = im.InsertUser(username, encryptpass)
	if err != nil {
		return err
	}

	return nil
}

// UserRegValidate check if username is existed
// check username and password validation
func (im *IMService) userRegValidate(username, password string) error {
	if userLen := len(username); userLen < 1 || userLen > 20 {
		return fmt.Errorf("username is too short or too long. it should between [1,20]")
	}

	if passLen := len(password); passLen < 6 || passLen > 20 {
		return fmt.Errorf("password is too short or too long. it should between [6,20]")
	}

	isExist := im.CheckUser(username)
	if isExist {
		return fmt.Errorf("username has existed, login directly")
	}

	return nil
}

// InsertUser insert username and pass to db
func (im *IMService) InsertUser(username, encryptpass string) error {
	db := im.dbs.CreateDB()
	defer db.Close()

	stmt := im.dbs.STMTFactory("INSERT INTO user SET username=?,"+
		"password=?, created_time=?", db)
	defer stmt.Close()

	res, err := stmt.Exec(username, encryptpass, GetTimestamp())
	if err != nil {
		logrus.Warn("insert username and password to db error: ", err)
		return fmt.Errorf("register user error, please try again")
	}

	_, err = res.RowsAffected()
	if err != nil {
		logrus.Warn("get affected lines from restult error: ", err)
		return fmt.Errorf("register user error, please try again")
	}

	return nil
}

// encryptPass use salt encrypt pass
func (us *IMService) encryptPass(password string) string {
	dk, err := scrypt.Key([]byte(password), []byte(PASS_SALT),
		16384, 8, 1, 32)
	if err != nil {
		logrus.Fatal(err)
	}

	return string(dk)
}

// GetUserByName
func (im *IMService) GetUserByName(username string) (*User, error) {
	db := im.dbs.CreateDB()
	defer db.Close()

	userSQL := "SELECT id, username, created_time FROM user WHERE username=?;"
	stmt := im.dbs.STMTFactory(userSQL, db)
	defer stmt.Close()

	var user User
	err := stmt.QueryRow(username).Scan(&user.ID, &user.Username, &user.CreatedTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%s does not exist")
		} else {
			return nil, fmt.Errorf("get user error, please try later")
		}
	}

	return &user, nil
}
