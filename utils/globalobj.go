package utils

import (
	"encoding/json"
	"fmt"
	"github.com/lorenzoyu2000/zinx/ziface"
	"io/ioutil"
)

/*
	存储zinx的框架的信息，用户可通过conf/zinx.json配置
*/

type GlobalObj struct {
	// zinx的服务对象
	TCPServer ziface.IServer
	// 当前服务器主机IP
	Host string
	// 当前服务器主机监听端口号
	TcpPort int
	// 服务器名臣
	Name string
	// zinx版本
	ZinxVersion string
	// 最大连接数
	MaxConn int
	// 数据包最大值
	MaxPackage uint32
}

var GlobalObject *GlobalObj

func init() {
	GlobalObject = &GlobalObj{}
	GlobalObject.Reload()
}

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
	fmt.Println(GlobalObject)
}
