package service

import (
	"time"
)

// User entity
type User struct {
	ID          uint64    `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	CreatedTime time.Time `json:"created_time"`
}

type AddFriend struct {
	ID        uint64    `json:"id"`
	FriendMin string    `json:"fmin"`
	FriendMax string    `json:"fmax"`
	AddedTime time.Time `json:"added_time"`
}
