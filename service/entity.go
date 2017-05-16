package service

import (
	"time"
)

// User entity
type User struct {
	ID          uint64 `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	CreatedTime int64  `json:"created_time"`
}

type FriendRelationship struct {
	ID        uint64    `json:"id"`
	FriendMin string    `json:"fmin"`
	FriendMax string    `json:"fmax"`
	AddedTime time.Time `json:"added_time"`
	State     string    `json:"state"`
}

type Message struct {
	ID       uint64 `json:"id"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Msg      string `json:"msg"`
	State    string `json:"state"`
	SendTime uint64 `json:"send_time"`
}
