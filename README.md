# webim

a study project of Golang

## TODOs
- login and register
    - [X] login
        - [ ] username and password validate (frontend)
    - [X] register
        - [ ] username and password validate (frontend)
- friend
    - [X] add friend
    - [X] list friend
        - [ ] with msgs
    
## Debug
### Error 1366: Incorrect string value: '\xC9c\x8B~\xB9\xA0...' for column 'password'

http://stackoverflow.com/questions/10957238/incorrect-string-value-when-trying-to-insert-utf-8-into-mysql-via-jdbc

[change utf8 to utf8mb4 in mysql 5.5+](https://mathiasbynens.be/notes/mysql-utf8mb4)

[this url helps solve problem](http://stackoverflow.com/questions/8291184/mysql-general-error-1366-incorrect-string-value?rq=1)

## Reference
- [Golang包依赖管理工具gb简介](https://segmentfault.com/a/1190000004346513)
- [Golang包依赖管理工具govendor](https://github.com/kardianos/govendor)
- [DB:如何存储好友关系](https://www.zhihu.com/question/20216864)