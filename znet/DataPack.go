package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/lorenzoyu2000/zinx/utils"
	"github.com/lorenzoyu2000/zinx/ziface"
)

type DataPack struct{}

func (d *DataPack) GetHeadLen() uint32 {
	return 8
}

func (d *DataPack) Pack(message ziface.IMessage) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})

	err := binary.Write(dataBuff, binary.LittleEndian, message.GetMsgLen())
	if err != nil {
		return nil, err
	}

	err = binary.Write(dataBuff, binary.LittleEndian, message.GetMsgID())
	if err != nil {
		return nil, err
	}

	err = binary.Write(dataBuff, binary.LittleEndian, message.GetMsgData())
	if err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

func (d *DataPack) UnPack(data []byte) (ziface.IMessage, error) {
	dataBuff := bytes.NewReader(data)

	msg := &Message{}

	err := binary.Read(dataBuff, binary.LittleEndian, &msg.MsgLen)
	if err != nil {
		return nil, err
	}

	err = binary.Read(dataBuff, binary.LittleEndian, &msg.MsgID)
	if err != nil {
		return nil, err
	}

	if utils.GlobalObject.MaxPackageSize > 0 && utils.GlobalObject.MaxPackageSize < msg.MsgLen {
		return nil, errors.New("recv msg is too long")
	}
	return msg, nil
}

func NewDataPack() ziface.IDataPack {
	dataPack := &DataPack{}
	return dataPack
}
