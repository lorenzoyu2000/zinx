package znet

import (
	"fmt"
	"github.com/lorenzoyu2000/zinx/utils"
	"github.com/lorenzoyu2000/zinx/ziface"
	"math/rand"
	"net"
	"time"
)

/*
	IServer 的接口实现，定义一个Server的服务器模块
*/
type Server struct {
	// 服务器名称
	Name string
	// 协议版本
	IPVersion string
	// ip地址
	IP string
	// 端口号
	Port int
	// Router处理连接对应的业务
	MsgHandler ziface.IMsgHandle
	// 连接管理器
	ConnMgr ziface.IConnectionManager
	// 连接前的钩子函数
	OnConnCreate func(ziface.IConnection)
	// 连接后的钩子函数
	OnConnDestroy func(ziface.IConnection)
}

func (s *Server) Start() {
	fmt.Printf("[Start] server Listener at IP %s, Port %d, is starting\n", s.IP, s.Port)
	// 开启协程防止在Start()方法中阻塞，将阻塞点推迟到Server()方法中，为了在Server()中做一些启动服务之外的服务
	go func() {
		// 开启工作池，工作池全局唯一
		s.MsgHandler.StartWorkPool()
		// 获取连接
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", addr, " err: ", err)
			return
		}

		fmt.Println("start [zinx] server successed ", s.Name)
		rand.Seed(time.Now().Unix())
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err: ", err)
				continue
			}

			// 设置最大连接数的判断，如果超出最大连接数，则拒绝连接
			if s.ConnMgr.Size() >= utils.GlobalObject.MaxConn {
				// TODO 返回给用户错误信息
				fmt.Println("too many connections")
				conn.Close()
				continue
			}

			// 将连接和conn进行绑定，得到连接模块
			dealConn := NewConnection(s, conn, rand.Uint32(), s.MsgHandler)
			go dealConn.Start()
		}
	}()
}

// 关闭服务器
func (s *Server) Stop() {
	fmt.Println("zinx server stop")
	s.ConnMgr.ClearAllConn()
}

func (s *Server) Serve() {
	// 启动Server服务
	s.Start()
	// TODO 在服务启动之后做一些额外处理。
	// 这里考虑到Start()方法只启动服务，其职责单一，而把阻塞点设置在Server()，是为了以后的扩展性需求
	// 阻塞点
	select {}
}

// 提供给用户自定义Router方法的接口
func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)
}

func (s *Server) GetConnMgr() ziface.IConnectionManager {
	return s.ConnMgr
}

func NewServer() ziface.IServer {
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandle(),
		ConnMgr:    NewConnManager(),
	}
	return s
}

func (s *Server) SetOnConnCreate(hookFunc func(connection ziface.IConnection)) {
	s.OnConnCreate = hookFunc
}

func (s *Server) SetOnConnDestroy(hookFunc func(connection ziface.IConnection)) {
	s.OnConnDestroy = hookFunc
}

func (s *Server) CallOnConnCreate(conn ziface.IConnection) {
	if s.OnConnCreate != nil {
		fmt.Println("call OnConnCreate func")
		s.OnConnCreate(conn)
	}
}

func (s *Server) CallOnConnDestroy(conn ziface.IConnection) {
	if s.OnConnDestroy != nil {
		fmt.Println("call CallOnConnDestroy func")
		s.OnConnDestroy(conn)
	}
}
