package service

import (
	"database/sql"
	"fmt"

	"github.com/Sirupsen/logrus"
)

// AddFriend handle add friend
func (im *IMService) AddFriendRelationship(friend FriendRelationship) error {
	// check each username validate
	if isExist := im.CheckUser(friend.FriendMin); !isExist {
		return fmt.Errorf("user %s does not exist", friend.FriendMin)
	}
	if isExist := im.CheckUser(friend.FriendMax); !isExist {
		return fmt.Errorf("user %s does not exist", friend.FriendMax)
	}

	// check if friend pair exist
	if isExist := im.CheckFriendRelationshipExist(friend); isExist {
		// check friend relationship state
		if isActive := im.CheckFriendRelationshipState(friend); isActive {
			return fmt.Errorf("friend relationship exists")
		}

		// activate relationship
		db := im.dbs.CreateDB()
		defer db.Close()
		activeSQL := `UPDATE friend_relationship SET state="active" 
		WHERE fmin=? and fmax=?;`
		stmt := im.dbs.STMTFactory(activeSQL, db)
		defer stmt.Close()

		_, err := stmt.Exec(friend.FriendMin, friend.FriendMax)
		if err != nil {
			logrus.Warnf("active friend relationship between %s and %s error: %v",
				friend.FriendMin, friend.FriendMax, err)
			return fmt.Errorf("add friend error, please try later")
		}
		return nil
	} else {
		db := im.dbs.CreateDB()
		defer db.Close()
		preSQL := `
			INSERT INTO friend_relationship SET fmin=?,
			fmax=?,added_time=?,state=?;
		`
		stmt := im.dbs.STMTFactory(preSQL, db)
		defer stmt.Close()

		res, err := stmt.Exec(friend.FriendMin, friend.FriendMax, GetTimestamp(), "active")
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
}

// CheckFriendRelationshipExist check if friend relationship exists in db
// true exist false or not
func (im *IMService) CheckFriendRelationshipExist(friend FriendRelationship) bool {
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
			logrus.Fatalf("query friend relationship between %s and %s error: %s ",
				friend.FriendMin, friend.FriendMax, err.Error())
		}
	}

	return true
}

// CheckFriendRelationshipState assume friend relationship has existed, this method check
// relationship state, true "active" false "inactive"
func (im *IMService) CheckFriendRelationshipState(friend FriendRelationship) bool {
	db := im.dbs.CreateDB()
	defer db.Close()

	checkStateSQL := `SELECT state FROM friend_relationship WHERE fmin=? AND fmax=?;`
	stmt := im.dbs.STMTFactory(checkStateSQL, db)
	defer stmt.Close()

	var state string
	err := stmt.QueryRow(friend.FriendMin, friend.FriendMax).Scan(&state)
	if err != nil {
		logrus.Fatalf("query friend relationship state between %s and %s error: %s ",
			friend.FriendMin, friend.FriendMax, err.Error())
	}

	if state == "active" {
		return true
	} else {
		return false
	}
}

// DeleteFriendRelationship delete friend relationship,change relationship state
// from "active" to "inactive"
func (im *IMService) DeleteFriendRelationship(friend FriendRelationship) error {
	db := im.dbs.CreateDB()
	defer db.Close()

	deleteSQL := `UPDATE friend_relationship SET state="inactive" 
		WHERE fmin=? and fmax=?;`
	stmt := im.dbs.STMTFactory(deleteSQL, db)
	defer stmt.Close()

	_, err := stmt.Exec(friend.FriendMin, friend.FriendMax)
	if err != nil {
		logrus.Warnf("update friend relationship between %s and %s error: %v",
			friend.FriendMin, friend.FriendMax, err)
		return err
	}

	return nil
}

// ListFriendRelationship return all friends of username
func (im *IMService) ListFriendRelationship(username string) (*[]string, error) {
	listSQL := `SELECT fmin, fmax, state FROM friend_relationship WHERE 
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
			fmin  string
			fmax  string
			state string
		)
		if err := rows.Scan(&fmin, &fmax, &state); err != nil {
			logrus.Warnf("query friends error: %s", err.Error())
			return &[]string{""}, fmt.Errorf("get all friends of user %s error", username)
		}
		if state == "active" {
			if fmin == username {
				friendList = append(friendList, fmax)
			} else {
				friendList = append(friendList, fmin)
			}
		}
	}

	return &friendList, nil
}
