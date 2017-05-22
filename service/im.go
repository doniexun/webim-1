package service

import (
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
)

// SafeUserWSMap safe read or write UserWSMap
type SafeUserWSMap struct {
	userWS map[string]*websocket.Conn
	mutex  *sync.Mutex
}

// NewSafeUserWSMap return *SafeUserWSMap
func NewSafeUserWSMap() *SafeUserWSMap {
	userWSMap := &SafeUserWSMap{
		userWS: make(map[string]*websocket.Conn),
		mutex:  &sync.Mutex{},
	}

	return userWSMap
}

// Set safe map user to ws
func (uws *SafeUserWSMap) Set(username string, ws *websocket.Conn) {
	uws.mutex.Lock()
	defer uws.mutex.Unlock()
	uws.userWS[username] = ws
}

// Get safe get *websocket.Conn by username
func (uws *SafeUserWSMap) Get(username string) *websocket.Conn {
	uws.mutex.Lock()
	defer uws.mutex.Unlock()
	// if uws.HasKey(username) {
	// 	return uws.userWS[username]
	// }

	// return nil
	return uws.userWS[username]
}

// Delete safe delete user in map, note it does not close ws
func (uws *SafeUserWSMap) Delete(user string) {
	uws.mutex.Lock()
	defer uws.mutex.Unlock()

	if uws.HasKey(user) {
		delete(uws.userWS, user)
	}
}

// HasKey true has false or not
func (uws *SafeUserWSMap) HasKey(key string) bool {
	uws.mutex.Lock()
	defer uws.mutex.Unlock()
	_, ok := uws.userWS[key]
	return ok
}

// String return a copy of current UserWSMap
func (uws *SafeUserWSMap) String() map[string]websocket.Conn {
	uws.mutex.Lock()
	defer uws.mutex.Unlock()

	var dist = make(map[string]websocket.Conn)
	for k, v := range uws.userWS {
		dist[k] = *v
	}

	return dist
}

// Length count of websockets connections
func (uws *SafeUserWSMap) Length() int {
	uws.mutex.Lock()
	defer uws.mutex.Unlock()

	return len(uws.userWS)
}

// SafeMsgID safe generate msg id
// refer https://github.com/nlopes/slack/blob/db68538b374a5f052bf6befea117ac74760f4e8e/messageID.go
type SafeMsgID struct {
	msgID uint64
	mutex *sync.Mutex
}

// NewSafeMsgID new *SafeMsgId
func NewSafeMsgID() *SafeMsgID {
	mid := &SafeMsgID{
		msgID: 0,
		mutex: &sync.Mutex{},
	}

	return mid
}

// NextID return int64 msg id
func (mid *SafeMsgID) NextID() uint64 {
	mid.mutex.Lock()
	defer mid.mutex.Unlock()

	id := mid.msgID
	mid.msgID++
	return id
}

// SyncLatestID make msgID greater than latestID
func (mid *SafeMsgID) SyncLatestID(latestID uint64) {
	mid.mutex.Lock()
	defer mid.mutex.Unlock()

	if mid.msgID <= latestID {
		mid.msgID = latestID + 1
	}
}

// IMService im service entrypoint
type IMService struct {
	dbs       *DBService
	msgID     *SafeMsgID
	UserWSMap *SafeUserWSMap
}

// NewIMService return *IMService
func NewIMService(dbs *DBService) *IMService {
	im := &IMService{
		dbs:       dbs,
		msgID:     NewSafeMsgID(),
		UserWSMap: NewSafeUserWSMap(),
	}

	// sync msgID to max id
	im.msgID.SyncLatestID(im.GetMaxMsgID())

	return im
}

// GetTimestamp return timestamp in ms
func GetTimestamp() int64 {
	return time.Now().Unix()
}

// CheckErr handle err and fatal it
func CheckErr(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}
