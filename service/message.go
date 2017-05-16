package service

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
)

// GetUnreadMsg get unread msgs by receiver
func (im *IMService) GetUnreadMsg(receiver string) (*[]Message, error) {
	// find all active friends
	activeFriends, err := im.ListFriendRelationship(receiver)
	if err != nil {
		return &[]Message{}, err
	}
	if len(*activeFriends) == 0 {
		return &[]Message{}, nil
	}

	var msgList []Message
	// loop get unread msgs by sender
	for _, sender := range *activeFriends {
		msgList_, err := im.GetUnreadMsgBySender(sender)
		if err != nil {
			return &msgList, err
		}
		msgList = append(msgList, *msgList_...)
	}

	return &msgList, nil
}

// GetUnreadMsgBySender get all unread msgs by sender name
// note sender and receiver has active friend relationship
func (im *IMService) GetUnreadMsgBySender(sender string) (*[]Message, error) {
	db := im.dbs.CreateDB()
	defer db.Close()

	unreadSQL := `SELECT * FROM message WHERE sender=? and state='msg_cache';`
	stmt := im.dbs.STMTFactory(unreadSQL, db)
	defer stmt.Close()
	rows, err := stmt.Query(sender)
	if err != nil {
		if err == sql.ErrNoRows {
			return &[]Message{}, nil
		} else {
			logrus.Warnf("get unread msgs by sender error: %v", err)
			return &[]Message{}, fmt.Errorf("get unread msgs by sender error")
		}
	}

	var msgList []Message
	for rows.Next() {
		var msg Message
		err := rows.Scan(&msg.ID, &msg.Sender, &msg.Receiver,
			&msg.Msg, &msg.SendTime, &msg.State)
		if err != nil {
			return &msgList, fmt.Errorf("parse msg from rows error: %v", err)
		}
		msgList = append(msgList, msg)
	}

	return &msgList, nil
}

// SendMsgToReceiver
func (im *IMService) SendMsgToReceiver() {}

// SaveMsgToDB
func (im *IMService) SaveMsgToDB() {}

// SyncMsgToDB
func (im *IMService) SyncMsgToDB() {}

// CacheMsgToDB
func (im *IMService) CacheMsgToDB() {}

// MessageWS a websocket to handle multi user message
func (im *IMService) MessageWS(w http.ResponseWriter, r *http.Request) {
	return
}
