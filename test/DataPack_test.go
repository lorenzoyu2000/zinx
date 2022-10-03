package test

import (
	"fmt"
	"github.com/lorenzoyu2000/zinx/znet"
	"io"
	"net"
	"testing"
)

// 测试封包，拆包功能
func TestDataPack(t *testing.T) {
	// 模拟服务端
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("server listen err ", err)
		return
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("server accept err ", err)
				continue
			}
			// 每次进行syscall时用协程执行，来提高并发能力
			go func(conn net.Conn) {
				dataPack := znet.NewDataPack()
				for {
					headData := make([]byte, dataPack.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("head read err ", err)
						break
					}

					msgHead, err := dataPack.UnPack(headData)
					if err != nil {
						fmt.Println("unpack err ", err)
						return
					}

					// MSG是有数据的，需要进行二次读取
					if msgHead.GetMsgLen() > 0 {
						msg := msgHead.(*znet.Message)
						msg.Data = make([]byte, msg.GetMsgLen())
						_, err = io.ReadFull(conn, msg.Data)
						if err != nil {
							fmt.Println("data read err ", err)
							return
						}
						fmt.Println("MsgID: ", msg.MsgID, ", MsgLen: ", msg.MsgLen, ", MsgData: ", string(msg.Data))
					}
				}
			}(conn)
		}
	}()

	// 模拟客户端
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client connection err ", err)
		return
	}
	msg1 := &znet.Message{
		MsgID:  1,
		MsgLen: 5,
		Data:   []byte{'h', 'e', 'l', 'l', 'o'},
	}
	msg2 := &znet.Message{
		MsgID:  2,
		MsgLen: 5,
		Data:   []byte{'w', 'o', 'r', 'l', 'd'},
	}
	dataPack := znet.NewDataPack()
	byte1, err := dataPack.Pack(msg1)
	if err != nil {
		fmt.Println("pack message err ", err)
		return
	}
	byte2, err := dataPack.Pack(msg2)
	if err != nil {
		fmt.Println("pack message err ", err)
		return
	}
	byte1 = append(byte1, byte2...)
	conn.Write(byte1)
	select {}
}
