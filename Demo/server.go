package main

import (
	"fmt"
	"github.com/lorenzoyu2000/zinx/ziface"
	"github.com/lorenzoyu2000/zinx/znet"
)

// 模拟服务端
func main() {
	s := znet.NewServer()
	s.AddRouter(&TestRouter{})
	s.Serve()
}

type TestRouter struct {
	znet.BaseRouter
}

func (t *TestRouter) Handle(request ziface.IRequest) {
	fmt.Println("Router preHandler.....")
	fmt.Println("MsgID: ", request.GetMsgId(), ", MsgData: ", string(request.GetData()))
	msg := &znet.Message{
		MsgID:  2,
		MsgLen: 7,
		Data:   []byte{'H', 'a', 'n', 'd', 'l', 'e', 'r'},
	}
	dataPack := znet.NewDataPack()
	data, err := dataPack.Pack(msg)
	if err != nil {
		fmt.Println("pack msg error: ", err)
		return
	}
	_, err = request.GetConnection().GetTCPConnection().Write(data)
	if err != nil {
		fmt.Println("preHandle err ", err)
	}
}
