package main

import (
	"fmt"
	"net"
)

func main() {
	// 监听
	listener, err := net.Listen("tcp", "0.0.0.0:8800")
	defer func() {
		err = listener.Close()
	}()

	if err != nil {
		fmt.Println("Listen error: ", err)
		return
	}

	fmt.Println("listener", listener, err)

	// 阻塞等待客户端连接
	conn, err := listener.Accept()
	defer func() {
		err = conn.Close()
	}()

	if err != nil {
		fmt.Println("Accept error", err)
	}

	fmt.Println("conn", conn, err)

	// 接收客户端的数据
	bytes := make([]byte, 1024)
	n, err := conn.Read(bytes)

	if err != nil {
		fmt.Println("Read error", err)
	}

	fmt.Println("客户端的数据：", string(bytes[:n]))
}
