package db

import (
	"database/sql"
	"time"

	"fmt"

	"github.com/Sirupsen/logrus"
)

// CheckUser check if user exists. true exist or false not exist
func (dbs *DBService) CheckUser(username string) bool {
	db := dbs.CreateDB()
	defer db.Close()

	stmt, err := db.Prepare("SELECT username FROM user WHERE username=?")
	defer stmt.Close()
	if err != nil {
		logrus.Fatal("prepare check username stmt error: ", err)
	}

	var username_ string
	err = stmt.QueryRow(username).Scan(&username_)
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
func (dbs *DBService) CheckUserPass(username, encryptPass string) error {
	db := dbs.CreateDB()
	defer db.Close()

	stmt, err := db.Prepare("SELECT username FROM user WHERE username=? " +
		" AND password=?")
	defer stmt.Close()
	if err != nil {
		logrus.Warn("prepared stmt error: ", err)
		return err
	}

	var username_ string
	err = stmt.QueryRow(username, encryptPass).Scan(&username_)
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

func (dbs *DBService) InsertUser(username, encryptpass string) error {
	db := dbs.CreateDB()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO user SET username=?," +
		"password=?, created_time=?")
	defer stmt.Close()
	if err != nil {
		logrus.Warn("prepared stmt error: ", err)
		return err
	}

	res, err := stmt.Exec(username, encryptpass, time.Now())
	if err != nil {
		logrus.Warn("insert username and password to db error: ", err)
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		logrus.Warn("get affected lines from restult error: ", err)
		return err
	}

	return nil
}

func handleStmtErr(err error) {
	if err != nil {
		logrus.Fatal("prepared stmt error: ", err)
	}
}
