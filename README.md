# webim

A study project of Golang,this is the webim Golang backend program.For front end just go to [webimfe](https://github.com/adolphlwq/webimfe).

## Requisite
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

Then browse [localhost:8080](http://localhost:8080)

## TODOs(zh)
- [X] 注册和登录
    - [X] 登录（username唯一id）
    - [X] 注册
- [ ] 联系人页面
    - [X] 显示**自己所有的联系人**
        - [X] 显示对方id
        - [ ] 显示未读私信数量提醒
            - [X] API
            - [ ] 前端渲染
    - [X] 通过id添加新联系人（不需要对方同意）
    - [X] 删除某个联系人
        - [X] 删除联系人
        - [X] 保留与删除用户的消息等数据
        - [X] 再次添加被删除联系人时，消息等数据都还在
- [ ] 聊天
    - [X] 点击一个联系人会进入聊天界面
    - [ ] 点击联系人进入聊天界面，未读消息置为 0
    - [ ] 查看某个用户的历史消息
    - [X] 收发私信（实时）
    - [ ] 删除自己发的消息
- [ ] 部署
    - [X] server
    - [X] docker-compose

## Architecture design
### 后端代码目录
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

- `api目录`：存放API映射关系
- `service目录`：内部主要逻辑实现

### 聊天部分
1. 聊天信息存储到数据库
2. 将每条消息设置状态，根据消息的状态确定已读未读等：
- `msg_send`：sender发送消息到chat-server
- `msg_cache`：chat-server接收到sender的消息，但是receiver不在线，缓存到服务器
- `msg_done`：消息顺利从sender发到receiver
- `msg_delete`：消息被删除
3. 聊天主要有3中情况：
  - 双方都不在线
  - 双方都在线（实时通信）
  - 一方在线，一方不在线（缓存信息）

#### 都在线
```
          发送消息                  1. transfer msg to receiver by websocket
sender --------------> ChatServer -------------------------------------------> receiver 
        `msg_send`                 2. save msg to db with state `msg_done`
```

#### 只有一方在线
```
          发送消息                  1. save msg to db with state `msg_cache`
sender --------------> ChatServer -------------------------------------------> receiver (`offline`)
        `msg_send`                 
```

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

## Features
- 初始化系统后消息ID自动和数据库中最新值同步
- 用户退出要清理websocket.

## Reference
- [使用Golang scrypt包加密后存储MySQL编码问题](http://stackoverflow.com/questions/8291184/mysql-general-error-1366-incorrect-string-value?rq=1)...(Error 1366: Incorrect string value: '\xC9c\x8B~\xB9\xA0...' for column 'password')
- [change utf8 to utf8mb4 in mysql 5.5+](https://mathiasbynens.be/notes/mysql-utf8mb4)
- [DB:如何存储好友关系](https://www.zhihu.com/question/20216864)
