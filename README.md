## GOWS


> a simple ws program in go
> 这只是个demo，不建议生产使用，有的地方还需优化，，，，，



* 可以聊天
* 用户身份验证
* 群组聊天
* 消息通知
* 添加简单演示页面



> 说明



App —> Group —> Member

全局 —>   群组  —>    成员



成员只允许在群内聊天，成员拥有 群发，和群内单聊权限

全局可以通过https/http触发通知 

当用户解析的token内不含有gid记录时，默认在ALL群，即公开群，可以当做是一个大群



[API文档](server/param.md)

![演示页面](./img/wstest.gif)

