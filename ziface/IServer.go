package ziface

/*
	服务器接口
*/
type IServer interface {
	// 服务启动
	Start()
	// 服务停止
	Stop()
	// 服务运行
	Serve()
	// 给当前服务添加一个路由功能，供客户端的连接使用
	AddRouter(uint32, IRouter)
	// 获取ConnManager
	GetConnMgr() IConnectionManager
	// 设置钩子函数
	SetOnConnCreate(func(IConnection))
	SetOnConnDestroy(func(IConnection))
	// 调用钩子函数
	CallOnConnCreate(IConnection)
	CallOnConnDestroy(IConnection)
}
