package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8800")

	defer func() {
		err = conn.Close()
	}()

	if err != nil {
		fmt.Println("dial error: ", err)
	}

	// 开启一个协程用户接收服务端的消息
	go receive(conn)

	// 主线程不断接收用户输入
	for {
		bytes := make([]byte, 1024)
		n, err := os.Stdin.Read(bytes)
		if err != nil {
			fmt.Println("stdin read error: ", err)
			return
		}

		// 读取的时候，也需要注意最后一个 \n 问题
		var readStr = string(bytes[:n-1])

		if "exit" == readStr {
			return
		}

		n, err = conn.Write([]byte(readStr))
		if err != nil {
			fmt.Println("conn write error: ", err)
			return
		}
	}
}

func receive(conn net.Conn) {
	for {

		bytes := make([]byte, 1024)
		n, err := conn.Read(bytes)

		if err != nil {
			fmt.Println("read error: ", err)
			return
		}

		fmt.Println("接收到服务器的内容：", string(bytes[:n]))
	}
}
