package db

import (
	"database/sql"
	"time"

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
	}

	return nil
}
