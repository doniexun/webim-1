# webim

A web instant message(im) service written by Golang,this is the backend program.For front end just go to [webimfe](https://github.com/adolphlwq/webimfe).

## Requisites
- Golang
- Mysql
- Node
- Docker
- docker-compose

## Quick start
```
git clone https://github.com/adolphlwq/webim
docker-compose up -d
```

Then browse [localhost:8080](localhost:8080)

## TODOs
- [ ] test case
- [ ] improve auth and session logic
    - [ ] map websocket to user
- [X] auth(login and register)
    - [X] login
    - [X] register
- [ ] contacts 
    - [X] list **all contacts of you**
        - [X] id
        - [ ] list messafes unread
            - [X] API
            - [ ] front end render
    - [X] add new contact by id
    - [X] delete contact
        - [X] delete
        - [X] reserve messages for add contact again
- [ ] **chat**
    - [X] into chat page when click a contact of contact list
    - [ ] unread set to zero
    - [ ] see history messages
    - [X] send and receiver message(real time if both are online)
    - [ ] delete messages
- [X] deploy
    - [X] server
    - [X] [docker-compose0(/docker-compose.yaml)

## APIs
- login
```
application/json
POST   /api/v1/user/login
```
- register
```
application/json
POST   /api/v1/user/register
```
- logout
```
application/json
POST   /api/v1/user/logout
```
- get user info
```
GET    /api/v1/user/get?username=luwenquan
// http://localhost:8877/api/v1/user/get?username=luwenquan
{
  "data": {
    "id": 1,
    "username": "luwenquan",
    "password": "",
    "created_time": 1495034536
  },
  "status": 200
}
```
- add friend
```
POST   /api/v1/friend/add
```
- list friends
```
GET    /api/v1/friend/list
// http://localhost:8877/api/v1/friend/list?username=luwenquan
{
  "data": [
    "test1",
    "test2"
  ],
  "status": 200
}
```
- delete friend
```
PUT    /api/v1/friend/delete
{
  "data": "delete success.",
  "status": 200
}
```
- get all unread msgs of specific user
```
GET    /api/v1/message/unread?receiver=test2
{
  "data": [
    {
      "id": 33,
      "sender": "luwenquan",
      "receiver": "test2",
      "msg": "we",
      "send_time": 1495094984,
      "state": "msg_cache"
    }
  ],
  "status": 200
}
```
- websocket chat
```
GET    /api/v1/message/ws/:username
```

## Architecture design
### backend directory
```
.
├── api
│   ├── api.go
│   ├── friendApi.go
│   ├── messageApi.go
│   └── userApi.go
├── docker-compose.yaml
├── Dockerfile.dev
├── LICENSE
├── Makefile
├── README.md
├── service
│   ├── db.go
│   ├── entity.go
│   ├── friend.go
│   ├── im.go
│   ├── message.go
│   └── user.go
├── vendor
│   ├── appengine
│   ├── github.com
│   ├── golang.org
│   └── vendor.json
└── webim.go
```

### chat logic design
1. messages are saved to db(mysql)
2. set `state` to each message
3. three main situations when chat：
  - both are offline
  - both are online(use websocket)
  - one is online and the other offline(cache messages)

#### both online
```
        send message                  1. transfer msg to receiver by websocket
sender --------------> ChatServer -------------------------------------------> receiver 
        `msg_send`                 2. save msg to db with state `msg_done`
```

#### only one online
```
        send message                 1. save msg to db with state `msg_cache`
sender --------------> ChatServer -------------------------------------------> receiver (`offline`)
        `msg_send`                 
```

## Reference
- [使用Golang scrypt包加密后存储MySQL编码问题](http://stackoverflow.com/questions/8291184/mysql-general-error-1366-incorrect-string-value?rq=1)...(Error 1366: Incorrect string value: '\xC9c\x8B~\xB9\xA0...' for column 'password')
- [change utf8 to utf8mb4 in mysql 5.5+](https://mathiasbynens.be/notes/mysql-utf8mb4)
- [DB:如何存储好友关系](https://www.zhihu.com/question/20216864)