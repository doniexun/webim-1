package service

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
)

// AddFriend handle add friend
func (im *IMService) AddFriend(friend AddFriend) error {
	// check each username validate
	if isExist := im.CheckUser(friend.FriendMin); !isExist {
		return fmt.Errorf("user %s does not exist", friend.FriendMin)
	}
	if isExist := im.CheckUser(friend.FriendMax); !isExist {
		return fmt.Errorf("user %s does not exist", friend.FriendMax)
	}

	// check if friend pair exist
	if isExist := im.CheckFriendExist(friend); isExist {
		return fmt.Errorf("friend relationship exists")
	}

	db := im.dbs.CreateDB()
	defer db.Close()
	preSQL := `
		INSERT INTO friend_relationship SET fmin=?,
		fmax=?,added_time=?
	`
	stmt := im.dbs.STMTFactory(preSQL, db)
	defer stmt.Close()

	res, err := stmt.Exec(friend.FriendMin, friend.FriendMax, time.Now())
	if err != nil {
		logrus.Warnf("insert friend relationship between %s and %s error %s",
			friend.FriendMin, friend.FriendMax, err.Error())
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		logrus.Fatal("get affected lines from restult error: ", err)
	}

	return nil
}

// CheckFriendExist check if friend relationship exists in db
// true exist false or not
func (im *IMService) CheckFriendExist(friend AddFriend) bool {
	checkSQL := `
		SELECT id FROM friend_relationship WHERE fmin=? 
		AND fmax=?;
	`
	db := im.dbs.CreateDB()
	defer db.Close()
	stmt := im.dbs.STMTFactory(checkSQL, db)
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

// ListFriend return all friends of username
func (im *IMService) ListFriend(username string) (*[]string, error) {
	listSQL := `SELECT fmin, fmax FROM friend_relationship WHERE 
		fmin=? OR fmax=?;`
	db := im.dbs.CreateDB()
	defer db.Close()
	stmt := im.dbs.STMTFactory(listSQL, db)
	defer stmt.Close()

	rows, err := stmt.Query(username, username)
	defer rows.Close()
	if err != nil {
		logrus.Warnf("query friends error: %s", err.Error())
		return &[]string{""}, fmt.Errorf("get all friends of user %s error", username)
	}

	var friendList []string
	for rows.Next() {
		var (
			fmin string
			fmax string
		)
		if err := rows.Scan(&fmin, &fmax); err != nil {
			logrus.Warnf("query friends error: %s", err.Error())
			return &[]string{""}, fmt.Errorf("get all friends of user %s error", username)
		}
		if fmin == username {
			friendList = append(friendList, fmax)
		} else {
			friendList = append(friendList, fmin)
		}
	}

	return &friendList, nil
}
