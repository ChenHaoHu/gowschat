# simple

> play ws://127.0.0.1:8080/ws/{token} 聊天接口
> token:{UID},{NAME},{Group}

> play1 ws://127.0.0.1:8080/ws/111,hcy,aaa

```
	P2P            //somone msg to notify somone 0
	P2G            // somone msg to notify group 1
	N2G            // msg to notify all 2
	N2A            // msg to notify all 3
	N2P            // msg to notify someone 4
```

```json
{
    "touid": "222",
    "msg": "hello sssssssssssss",
    "sendtype": 0    
}
```

> play2 ws://127.0.0.1:8080/ws/222,jyj,aaa
```json
{
    "touid": "222",
    "msg": "hello sssssssssssss",
    "sendtype": 0
}
```

> GET http://127.0.0.1:8080/chat/group  获取登陆信息
```json
{
    "message": {
        "ALL": {},
        "aaa": {
            "111": {
                "Uid": "111",
                "Name": "hcy111",
                "Conn": {},
                "LoginTime": "Mar 27, 2020 00:11:12 UTC"
            },
            "333": {
                "Uid": "333",
                "Name": "hcy333",
                "Conn": {},
                "LoginTime": "Mar 27, 2020 00:11:15 UTC"
            }
        },
        "bbb": {
            "222": {
                "Uid": "222",
                "Name": "hcy222",
                "Conn": {},
                "LoginTime": "Mar 27, 2020 00:11:14 UTC"
            }
        }
    },
    "status": "ok"
}

```


> POST http://127.0.0.1:8080/chat/notice NOTICE通知

> form

sendtype:4
msg:这是一个全局消息哦
gid:aaa
uid:111

```json
{
    "status": "ok"
}
```



