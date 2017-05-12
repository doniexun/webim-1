package db

import(
	"time"
)

// User entity
type User struct {
	ID          uint64	`json:"id"`
	Username    string	`json:"username"`
	Password    string	`json:"password"`
	CreatedTime time.Time	`json:"created_time"`
}
