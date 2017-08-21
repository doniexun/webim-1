# webim

[![Build Status](https://travis-ci.org/adolphlwq/webim.svg?branch=dev)](https://travis-ci.org/adolphlwq/webim)  [![Go Report Card](https://goreportcard.com/badge/github.com/adolphlwq/webim)](https://goreportcard.com/report/github.com/adolphlwq/webim)  [![GoDoc](https://godoc.org/github.com/adolphlwq/webim?status.svg)](https://godoc.org/github.com/adolphlwq/webim)

A web instant message(im) service written by Golang,this is the backend program.

## Requisites
- Golang
- [govendor](https://github.com/kardianos/govendor)
- Mysql
- Node
- Docker
- docker-compose

## Quick start
```
git clone https://github.com/adolphlwq/webim $GOPATH/src/github.com/adolphlwq/webim
docker-compose up -d
```

Then browse [localhost:8080](localhost:8080)

## TODOs
- [X] auth(login and register)
    - [X] login
    - [X] register
- [X] deploy
    - [X] server
    - [X] [docker-compose0(/docker-compose.yaml)
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
- [ ] test case
    
## Architecture design
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