# zinx

## introduce

zinx是一个使用TCP协议的服务端开发框架，提供以下几项功能：

- 连接管理：对每个和服务端连接的客户端分配一个 Connection 来对连接进行管理，并且为了提升拓展性，为每个 Conn 提供自定义的属性。
- 自定义处理方法：可在服务器提供自定义的多种消息处理方法，并且使用 Goroutine 进行读写分离，采用多工作消息队列提高消息处理并发性能。
- 消息格式：自定义消息格式来解决 nagle 导致的封包拆包问题，并实现封包、拆包器来进行消息发送。

![zinx消息处理流程](https://imgs-1306864474.cos.ap-beijing.myqcloud.com/img/zinx%E6%B6%88%E6%81%AF%E5%A4%84%E7%90%86%E6%B5%81%E7%A8%8B.jpg)

## quick start

获取包

```go
go get github.com/lorenzoyu2000/zinx
```

 注册服务端信息

```go
func main() {
	s := znet.NewServer()
	// 注册钩子函数
	s.SetOnConnCreate(onConnCreate)
	s.SetOnConnDestroy(onConnDestroy)
	// 注册处理函数
	s.AddRouter(2, &apis.WorldChat{})
	s.AddRouter(3, &apis.MoveApi{})
	s.Serve()
}
```

开启服务

```go
go run main.go
```

![image-20221018164408310](https://imgs-1306864474.cos.ap-beijing.myqcloud.com/img/image-20221018164408310.png)

## application

可用服务器执行的程序[mmo_game](github.com/lorenzoyu2000/mmo_game)
