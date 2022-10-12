package znet

import (
	"errors"
	"fmt"
	"github.com/lorenzoyu2000/zinx/utils"
	"github.com/lorenzoyu2000/zinx/ziface"
	"io"
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
	// 使用chan进行读写分离
	MsgChan chan []byte
	// 该连接处理的方法
	MsgHandler ziface.IMsgHandle
	// 加入server，方便获取ConnMgr
	TcpServer ziface.IServer
}

func (c *Connection) Start() {
	fmt.Println("Conn Start, ConnID is ", c.ConnID)
	go c.startReader()
	// 读写业务分离
	go c.startWriter()
	// 调用开发者传递的钩子函数
	c.TcpServer.CallOnConnCreate(c)
}

func (c *Connection) startReader() {
	fmt.Println("[Reader is running...]")
	defer fmt.Println("Reader exit, ConnID is ", c.ConnID)
	defer c.Stop()

	for {
		/*buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("Reader err ", err)
			continue
		}*/

		// 优化使用拆包器
		dataPack := NewDataPack()
		headData := make([]byte, dataPack.GetHeadLen())
		_, err := io.ReadFull(c.GetTCPConnection(), headData)
		if err != nil {
			fmt.Println("[Connection] read headData err ", err)
			return
		}

		msg, err := dataPack.UnPack(headData)
		if err != nil {
			fmt.Println("[Connection] unpack headData err ", err)
			return
		}
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			_, err := io.ReadFull(c.GetTCPConnection(), data)
			if err != nil {
				fmt.Println("[Connection] read data err ", err)
				continue
			}
			msg.SetMsgData(data)
		}
		// 将客户端请求封装为Request
		req := &Request{
			conn:    c,
			message: msg,
		}

		// 如果工作池已开启，就放入到工作池中
		if utils.GlobalObject.WorkPoolSize > 0 {
			c.MsgHandler.SendMsgToTaskQueue(req)
		} else {
			// 从路由中找到注册绑定Conn对应的Router
			go c.MsgHandler.DoMsgRouter(req)
		}
	}
}

func (c *Connection) startWriter() {
	fmt.Println("[starWriter is running...]")
	defer fmt.Println("Writer exit, ConnID is ", c.ConnID)

	for {
		select {
		case data := <-c.MsgChan:
			{
				_, err := c.Conn.Write(data)
				if err != nil {
					fmt.Println("Writer write data err: ", err)
					return
				}
			}

		case <-c.ExitChan:
			return
		}
	}
}

func (c *Connection) Stop() {
	defer fmt.Println("Conn stopped, ConnID is ", c.ConnID)

	if c.isClosed == true {
		return
	}

	c.TcpServer.CallOnConnDestroy(c)

	c.isClosed = true
	c.ExitChan <- true
	close(c.ExitChan)
	close(c.MsgChan)
	c.TcpServer.Stop()
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

// 发送消息之前，先使用封包器处理消息，再发送
func (c *Connection) Send(msgID uint32, data []byte) error {
	if c.isClosed {
		return errors.New("send msg to a closed connection")
	}

	dataPack := NewDataPack()
	msg := NewMessage(msgID, data)
	binaryMsg, err := dataPack.Pack(msg)
	if err != nil {
		return errors.New("pack msg error")
	}

	c.MsgChan <- binaryMsg
	return nil
}

func NewConnection(server ziface.IServer, conn *net.TCPConn, connID uint32, handler ziface.IMsgHandle) *Connection {
	c := &Connection{
		Conn:       conn,
		ConnID:     connID,
		isClosed:   false,
		MsgHandler: handler,
		MsgChan:    make(chan []byte),
		ExitChan:   make(chan bool, 1),
		TcpServer:  server,
	}

	server.GetConnMgr().AddConn(c)
	return c
}
