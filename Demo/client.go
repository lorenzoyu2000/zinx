package main

import (
	"fmt"
	"github.com/lorenzoyu2000/zinx/znet"
	"io"
	"net"
	"time"
)

// 模拟客户端
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Println("client conn is err ", err)
		return
	}

	for {
		dataPack := znet.NewDataPack()
		m := &znet.Message{
			MsgID:  2,
			MsgLen: 5,
			Data:   []byte{'h', 'e', 'l', 'l', 'o'},
		}
		binaryM, err := dataPack.Pack(m)
		if err != nil {
			fmt.Println("pack msg err: ", err)
			return
		}
		_, err = conn.Write(binaryM)
		if err != nil {
			fmt.Println("client conn write err ", err)
			continue
		}

		headData := make([]byte, dataPack.GetHeadLen())
		_, err = io.ReadFull(conn, headData)
		if err != nil {
			fmt.Println("read headData err: ", err)
			return
		}

		msgHead, err := dataPack.UnPack(headData)
		if err != nil {
			fmt.Println("unpack headData err: ", err)
			return
		}

		if msgHead.GetMsgLen() > 0 {
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msgHead.GetMsgLen())
			_, err = io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("read data err: ", err)
				return
			}
			fmt.Println("MsgID: ", msg.MsgID, ", MsgData: ", string(msg.Data))
		}
		time.Sleep(1 * time.Second)
	}
}
