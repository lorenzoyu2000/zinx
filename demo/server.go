package main

import "github.com/lorenzoyu2000/zinx/znet"

// 模拟服务端
func main() {
	s := znet.NewServer("[zinx v1.0]")
	s.Serve()

}
