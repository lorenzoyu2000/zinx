package main

import (
	"fmt"
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
		_, err := conn.Write([]byte("Hello this is zinx!!!"))
		if err != nil {
			fmt.Println("client conn write err ", err)
			continue
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("client conn read err ", err)
			continue
		}
		fmt.Printf("client recv data is [%s], cnt is [%d]\n", buf, cnt)
		time.Sleep(1 * time.Second)
	}

}
