# simple

> play ws://127.0.0.1:8080/ws/{token}
> token:{UID},{NAME},{Group}

> play1 ws://127.0.0.1:8080/ws/111,hcy,aaa
```json
{
"touid":"222",
"msg":"hello sssssssssssss",
"sendtype":0
}
```

> play2 ws://127.0.0.1:8080/ws/222,jyj,aaa
```json
{
"uid":"222",
"touid":"111",
"msg":"hello hcy",
"sendtype":0
}
```


