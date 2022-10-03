package ziface

/*
	将请求的消息封装到Message中
*/
type IMessage interface {
	/*
		get、set方法
	*/
	GetMsgID() uint32
	GetMsgLen() uint32
	GetMsgData() []byte

	SetMsgID(uint32)
	SetMsgLen(uint32)
	SetMsgData([]byte)
}
