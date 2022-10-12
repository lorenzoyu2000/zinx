package ziface

/*
	连接管理模块
*/
type IConnectionManager interface {
	// 添加连接
	AddConn(IConnection)
	// 删除连接
	RemoveConn(IConnection)
	// 获取连接
	GetConn(uint32) (IConnection, error)
	// 获取连接个数
	Size() int
	// 断开所有连接
	ClearAllConn()
}
