package service

import (
	"database/sql"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
)

// GetUnreadMsg get unread msgs of receiver
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
	// loop to get unread msgs by sender
	for _, sender := range *activeFriends {
		newMsgList, err := im.GetUnreadMsgBySender(sender)
		if err != nil {
			return &msgList, err
		}
		msgList = append(msgList, *newMsgList...)
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
		}

		logrus.Warnf("get unread msgs by sender error: %v", err)
		return &[]Message{}, fmt.Errorf("get unread msgs by sender error")
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

// HandleMsgFromMsgChan get msg from chan and handle according sender and receiver state
func (im *IMService) HandleMsgFromMsgChan(msgChan chan Message) {
	for msg := range msgChan {
		msg.ID = im.msgID.NextID()
		logrus.Infof("get msg: %v", msg)
		// check receiver state
		if isReceiverLogin := im.UserWSMap.HasKey(msg.Receiver); isReceiverLogin {
			// receiver is login, send msg to recevier
			logrus.Infof("receiver: %s is login, send to it", msg.Receiver)
			msg.State = "msg_receiver"
			im.SendMsgToReceiver(msg)
			logrus.Infof("after send msg to receiver, save msg to db with state msg_done")
			msg.State = "msg_done"
			err := im.SaveMsgToDB(msg)
			CheckErr(err)
		} else {
			// receiver is logout, cache msg to db
			logrus.Infof("receiver: %s is logout, cache msg to db", msg.Receiver)
			msg.State = "msg_cache"
			err := im.CacheMsgToDB(msg)
			CheckErr(err)
		}
	}
}

// HandleMsgFromWS read msg from websocket and send to msgChan
func (im *IMService) HandleMsgFromWS(ws *websocket.Conn, msgChan chan Message, username string) {
	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		// refer to https://godoc.org/github.com/gorilla/websocket#IsUnexpectedCloseError
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				logrus.Warnf("read data from websocket error: %v", err)
			}
			// delete ws of username
			logrus.Infof("message.go HandleMsgFromWS delete %s in UserWSMap", username)
			im.UserWSMap.Delete(username)
			logrus.Infof("ws of %s has closed, exit this goroutine", username)
			return
		}
		logrus.Infof("get msg from ws: %v", msg)
		msgChan <- msg
	}
}

// SendMsgToReceiver
func (im *IMService) SendMsgToReceiver(msg Message) {
	logrus.Infof("message.go SendMsgToReceiver get ws of receiver: %s", msg.Receiver)
	ws := im.UserWSMap.Get(msg.Receiver)
	// TODOs validate ws?
	ws.WriteJSON(msg)
}

// SaveMsgToDB
func (im *IMService) SaveMsgToDB(msg Message) error {
	db := im.dbs.CreateDB()
	defer db.Close()
	saveSQL := `INSERT INTO message SET id=?,sender=?,
		receiver=?,msg=?,send_time=?,state=?;`
	stmt := im.dbs.STMTFactory(saveSQL, db)
	defer stmt.Close()

	res, err := stmt.Exec(msg.ID, msg.Sender, msg.Receiver,
		msg.Msg, msg.SendTime, msg.State)
	if err != nil {
		logrus.Warnf("insert into message table error: %v", err)
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

// CacheMsgToDB cache msg to message table with state "msg_cache"
func (im *IMService) CacheMsgToDB(msg Message) error {
	return im.SaveMsgToDB(msg)
}

// GetMaxMsgID get max id of message table, 0 if message table is empty.
func (im *IMService) GetMaxMsgID() uint64 {
	db := im.dbs.CreateDB()
	defer db.Close()

	rSQL := `SELECT MAX(id) FROM message;`
	stmt := im.dbs.STMTFactory(rSQL, db)
	defer stmt.Close()

	var tmpID interface{}
	err := stmt.QueryRow().Scan(&tmpID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0
		}
		logrus.Fatalf("get max id from message error: %v", err)
	}

	if tmpID == nil {
		return 0
	}
	return uint64(tmpID.(int64))
}

// UpdateMsgListState update msg arrays state
func (im *IMService) UpdateMsgListState(msgList *[]Message, state string) {
	for _, msg := range *msgList {
		im.UpdateMsgStateByID(msg.ID, state)
	}
}

// UpdateMsgStateByID update msg state of id to aimed state
func (im *IMService) UpdateMsgStateByID(id uint64, state string) {
	db := im.dbs.CreateDB()
	defer db.Close()

	updateSQL := `UPDATE message SET state=? WHERE id=?;`
	stmt := im.dbs.STMTFactory(updateSQL, db)
	defer stmt.Close()

	res, err := stmt.Exec(state, id)
	if err != nil {
		logrus.Fatalf("update msg state of id: %d to %s error: %s",
			id, state, err.Error())
	}

	_, err = res.RowsAffected()
	if err != nil {
		logrus.Fatal(err)
	}
}
