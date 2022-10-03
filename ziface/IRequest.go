package ziface

/*
	把客户端发送的请求封装到IRequest
*/

type IRequest interface {
	// 得到当前连接
	GetConnection() IConnection
	// 得到请求消息数据
	GetData() []byte
	// 获取消息ID
	GetMsgId() uint32
}
