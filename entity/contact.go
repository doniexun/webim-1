package entity

import "time"
import "strings"

type Contact struct {
	ID        int64     `gorm:"primary_key"`
	FriendMin string    `json:"fmin"`
	FriendMax string    `json:"fmax"`
	AddedTime time.Time `json:"added_time"`
	State     string    `json:"state"`
}

// Standerize make friendMax and friendMin lexicographically
func (c *Contact) Standerize() {
	if strings.Compare(c.FriendMax, c.FriendMin) < 0 {
		tmp := c.FriendMax
		c.FriendMax = c.FriendMin
		c.FriendMin = tmp
	}
}
