package main

import (
	"fmt"
	"github.com/lorenzoyu2000/zinx/ziface"
	"github.com/lorenzoyu2000/zinx/znet"
)

// 模拟服务端
func main() {
	s := znet.NewServer("[zinx v3.0]")
	s.AddRouter(&TestRouter{})
	s.Serve()
}

type TestRouter struct {
	b znet.BaseRouter
}

func (t *TestRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Router preHandler.....")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("preHandler do something...."))
	if err != nil {
		fmt.Println("preHandle err ", err)
	}
}

func (t *TestRouter) Handle(request ziface.IRequest) {}

func (t *TestRouter) PostHandle(request ziface.IRequest) {}
