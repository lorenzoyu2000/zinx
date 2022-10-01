package ziface

/**
服务器接口
*/
type IServer interface {
	// 服务启动
	Start()
	// 服务停止
	Stop()
	// 服务运行
	Serve()
}
