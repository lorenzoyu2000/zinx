package ziface

/*
	拆包、封包模块，用于解决TCP Nagle算法的粘包、拆包问题
*/
type IDataPack interface {
	// 获取头部长度
	GetHeadLen() uint32
	// 封包
	Pack(IMessage) ([]byte, error)
	// 拆包
	UnPack([]byte) (IMessage, error)
}
