package entity

// User entity
type User struct {
	ID          uint64 `gorm:"primary_key"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	CreatedTime int64  `json:"created_time"`
}
