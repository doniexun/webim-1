package entity

import "time"

type Contact struct {
	ID        int64     `gorm:"primary_key"`
	FriendMin string    `json:"fmin"`
	FriendMax string    `json:"fmax"`
	AddedTime time.Time `json:"added_time"`
	State     string    `json:"state"`
}
