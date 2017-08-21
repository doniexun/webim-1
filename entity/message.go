package entity

// Message entity
type Message struct {
	ID       uint64 `gorm:"primary_key"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Msg      string `json:"msg"`
	SendTime uint64 `json:"send_time"`
	State    string `json:"state"`
}
