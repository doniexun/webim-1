package service

import (
	"fmt"
	"webim/db"
)

// AddFriend handle add friend
func (us *UserService) AddFriend(friend db.AddFriend) error {
	// check each username validate
	if isExist := us.dbs.CheckUser(friend.FriendMin); !isExist {
		return fmt.Errorf("user %s does not exist", friend.FriendMin)
	}
	if isExist := us.dbs.CheckUser(friend.FriendMax); !isExist {
		return fmt.Errorf("user %s does not exist", friend.FriendMax)
	}

	err := us.dbs.AddFriend(friend)
	if err != nil {
		return err
	}

	return nil
}

// ListFriend return all friends of username
func (us *UserService) ListFriend(username string) (*[]string, error) {
	return us.dbs.ListFriend(username)
}
