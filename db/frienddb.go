package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
)

// AddFriend add friend pair (minName, maxName) to friend_relationship
func (dbs *DBService) AddFriend(friend AddFriend) error {
	// check if friend pair exist
	if isExist := dbs.CheckFriendExist(friend); isExist {
		return fmt.Errorf("friend relationship exists")
	}

	db := dbs.CreateDB()
	defer db.Close()
	preSQL := `
		INSERT INTO friend_relationship SET fmin=?,
		fmax=?,added_time=?
	`
	stmt := dbs.STMTFactory(preSQL, db)
	defer stmt.Close()
	res, err := stmt.Exec(friend.FriendMin, friend.FriendMax, time.Now())
	if err != nil {
		logrus.Warnf("insert friend relationship between %s and %s error %s",
			friend.FriendMin, friend.FriendMax, err.Error())
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		logrus.Warn("get affected lines from restult error: ", err)
		return err
	}

	return nil
}

// CheckFriendExist check if friend relationship exists in db
// true exist or false not
func (dbs *DBService) CheckFriendExist(friend AddFriend) bool {
	checkSQL := `
		SELECT id FROM friend_relationship WHERE fmin=? 
		AND fmax=?;
	`
	db := dbs.CreateDB()
	defer db.Close()
	stmt := dbs.STMTFactory(checkSQL, db)
	defer stmt.Close()

	var id uint64
	err := stmt.QueryRow(friend.FriendMin, friend.FriendMax).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			logrus.Fatalf("query friend relationship between %s and %s error %s ",
				friend.FriendMin, friend.FriendMax, err.Error())
		}
	}

	return true
}

func (dbs *DBService) ListFriend(user string) (*[]string, error) {
	listSQL := `SELECT fmin, fmax FROM friend_relationship WHERE 
		fmin=? OR fmax=?;`
	db := dbs.CreateDB()
	defer db.Close()
	stmt := dbs.STMTFactory(listSQL, db)
	defer stmt.Close()

	rows, err := stmt.Query(user, user)
	defer rows.Close()
	if err != nil {
		logrus.Warnf("query friends error: %s", err.Error())
		return &[]string{""}, fmt.Errorf("get all friends of user %s error", user)
	}

	var friendList []string
	for rows.Next() {
		var (
			fmin string
			fmax string
		)
		if err := rows.Scan(&fmin, &fmax); err != nil {
			logrus.Warnf("query friends error: %s", err.Error())
			return &[]string{""}, fmt.Errorf("get all friends of user %s error", user)
		}
		if fmin == user {
			friendList = append(friendList, fmax)
		} else {
			friendList = append(friendList, fmin)
		}
	}

	return &friendList, nil
}
