package znet

type Message struct {
	// 消息ID
	MsgID uint32
	// 消息长度
	MsgLen uint32
	// 消息数据
	Data []byte
}

func (m *Message) GetMsgID() uint32 {
	return m.MsgID
}

func (m *Message) GetMsgLen() uint32 {
	return m.MsgLen
}
func (m *Message) GetMsgData() []byte {
	return m.Data
}

func (m *Message) SetMsgID(id uint32) {
	m.MsgID = id
}
func (m *Message) SetMsgLen(len uint32) {
	m.MsgLen = len
}
func (m *Message) SetMsgData(data []byte) {
	m.Data = data
}

func NewMessage(id uint32, data []byte) *Message {
	m := &Message{
		MsgID:  id,
		MsgLen: uint32(len(data)),
		Data:   data,
	}
	return m
}
