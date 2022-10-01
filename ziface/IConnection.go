package ziface

import "net"

/*
	连接模块接口
*/
type IConnection interface {
	// 启动连接
	Start()
	//停止连接
	Stop()
	// 获取当前连接ID
	GetConnID() uint32
	// 获取连接的socket conn
	GetTCPConnection() *net.TCPConn
	// 获取远程连接客户端的TCP状态 IP PORT
	GetRemoteAddr() net.Addr
	// 发送数据
	Send([]byte) error
}

// 处理连接业务的方法
type HandleFun func(*net.TCPConn, []byte, int) error
