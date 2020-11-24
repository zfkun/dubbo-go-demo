# dubbo-go-demo
demo of dubbo-go for bugfix

## Intro
- 示例包含 gate 和 game 两个服务
- 两个服务会互相 RPC 通讯, 即均注册 provider 和 consumer
- gate 额外启动了 http 服务 (8000端口), 用于手工触发 gate -> game 的RPC调用
- 每次 gate -> game RPC调用(Message) 后, game -> gate 同步RPC调用(Send) 推送相同消息

## Issure
当 gate 和 game 之间 长时间(几分钟或几小时) 无RPC调用后, 从日志可观测到 session 的清理释放, 直到出现 `left session number:0` 后, 若再触发RPC调用，从日志可观测到 `client new session` 重新分配会话, 但是调用会立即失败, 出现以下错误:

```
dubbo/dubbo_invoker.go:144	result.Err: client read timeout, result.Rest: <nil>
proxy/proxy.go:174	result err: client read timeout
```

## Log

- log/gate.log
- log/game.log


## Run

```shell
# start gate & game server
make run

# or

# start gate server
make run-gate

# start game server
make run-game
```

## Test

```shell
# 手工触发 gate -> game RPC调用
curl -X GET http://127.0.0.1:8000/test
```
