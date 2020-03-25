# simple



> play1 ws://127.0.0.1:8080/ws/111,hcy
```json
{
"uid":"111",
"touid":"222",
"msg":"hello jyj",
"msgtype":0
}
```

> play2 ws://127.0.0.1:8080/ws/222,jyj
```json
{
"uid":"222",
"touid":"111",
"msg":"hello hcy",
"msgtype":0
}
```