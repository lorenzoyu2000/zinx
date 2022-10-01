package znet

import (
	"fmt"
	"github.com/lorenzoyu2000/zinx/ziface"
	"net"
)

/*
	连接接口实现
*/
type Connection struct {
	// 当前连接的套接字
	Conn *net.TCPConn
	// 连接的ID
	ConnID uint32
	// 当前连接状态
	isClosed bool
	// 告知当前连接已经停止的Channel
	ExitChan chan bool
	// 该连接处理的方法
	Router ziface.IRouter
}

func (c *Connection) Start() {
	fmt.Println("Conn Start, ConnID is ", c.ConnID)
	go c.StartReader()
	// TODO 读写业务分离
}

func (c *Connection) StartReader() {
	defer fmt.Println("Reader exit, ConnID is ", c.ConnID)
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("Reader err ", err)
			continue
		}

		// 将客户端请求封装为Request
		req := Request{
			conn: c,
			data: buf,
		}
		// 从路由中找到注册绑定Conn对应的Router
		go func(req ziface.IRequest) {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.PostHandle(req)
		}(&req)
	}
}

func (c *Connection) Stop() {
	defer fmt.Println("Conn stopped, ConnID is ", c.ConnID)

	if c.isClosed == true {
		return
	}
	c.isClosed = true
	close(c.ExitChan)
	c.Conn.Close()
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send([]byte) error {
	return nil
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		Router:   router,
		ExitChan: make(chan bool, 1),
	}

	return c
}
