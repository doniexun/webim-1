# webim

A study project of Golang,this is the webim Golang backend program.For front end just go to [webimfe](https://github.com/adolphlwq/webimfe).

## TODOs
- [ ] login
    - [X] basic login
    - [ ] cookie/session/token
    - [ ] username and password validate (frontend)
- [ ] register
    - [X] basic register
    - [ ] username and password validate (frontend)
- [ ] friend
    - [X] add friend
    - [X] list friends of specific user without msginfo
    - [ ] list friends of sepcific user with info
- [ ] message
    - [ ] message page
    - [ ] send msg 
    - [ ] collect msg
    - [ ] real time notify logined user when new msg comes
- [ ] Dockerrize
    - [ ] webim image
    - [ ] webimfe image
    - [ ] docker-compose
    
## Reference
- [使用Golang scrypt包加密后存储MySQL编码问题](http://stackoverflow.com/questions/8291184/mysql-general-error-1366-incorrect-string-value?rq=1)...(Error 1366: Incorrect string value: '\xC9c\x8B~\xB9\xA0...' for column 'password')
- [change utf8 to utf8mb4 in mysql 5.5+](https://mathiasbynens.be/notes/mysql-utf8mb4)
- [DB:如何存储好友关系](https://www.zhihu.com/question/20216864)