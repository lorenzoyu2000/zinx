package znet

import (
	"fmt"
	"github.com/lorenzoyu2000/zinx/ziface"
	"net"
)

// IServer 的接口实现，定义一个Server的服务器模块
type Server struct {
	// 服务器名称
	Name string
	// 协议版本
	IPVersion string
	// ip地址
	IP string
	// 端口号
	Port int
}

func (s *Server) Start() {
	fmt.Printf("[Start] server Listener at IP %s, Port %d, is starting\n", s.IP, s.Port)
	// 开启协程防止在Start()方法中阻塞，将阻塞点推迟到Server()方法中，为了在Server()中做一些启动服务之外的服务
	go func() {
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

		fmt.Println("start [zinx] server succ ", s.Name)

		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err: ", err)
				continue
			}
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("conn read err ", err)
						continue
					}
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("conn write err ", err)
					}
				}
			}()
		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	// 启动Server服务
	s.Start()
	// TODO 在服务启动之后做一些额外处理。
	// 这里考虑到Start()方法只启动服务，其职责单一，而把阻塞点设置在Server()，是为了以后的扩展性需求
	// 阻塞点
	select {}
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "127.0.0.1",
		Port:      9090,
	}
	return s
}
