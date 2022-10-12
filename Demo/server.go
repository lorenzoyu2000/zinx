package main

import (
	"fmt"
	"github.com/lorenzoyu2000/zinx/ziface"
	"github.com/lorenzoyu2000/zinx/znet"
)

// 模拟服务端
func main() {
	s := znet.NewServer()
	s.AddRouter(1, &PingRouter{})
	s.AddRouter(2, &PongRouter{})
	s.Serve()
}

type PingRouter struct {
	znet.BaseRouter
}

func (t *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("ping ping ping.....")
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

type PongRouter struct {
	znet.BaseRouter
}

func (t *PongRouter) Handle(request ziface.IRequest) {
	fmt.Println("MsgID: ", request.GetMsgId(), ", MsgData: ", string(request.GetData()))
	msg := znet.NewMessage(202, []byte("pong pong pong..."))
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

func ConnCreate(conn ziface.IConnection) {
	err := conn.Send(2, []byte("Connection Created...."))
	if err != nil {
		fmt.Println("Create err: ", err)
		return
	}
}

func ConnDestroy(conn ziface.IConnection) {
	fmt.Println("connID", conn.GetConnID(), " is lost")
}
