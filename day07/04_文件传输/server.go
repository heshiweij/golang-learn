package main

import (
	"fmt"
	"net"
	"os"
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

	for {
		// 阻塞等待客户端连接
		conn, err := listener.Accept()

		defer func() {
			err = conn.Close()
		}()

		if err != nil {
			fmt.Println("Accept error", err)
			return
		}

		// 开启协程去处理客户端的请求
		go handleClient(conn)

		fmt.Println("listener", listener, err)
	}
}

// 处理客户端
func handleClient(conn net.Conn) {

	fmt.Println("欢迎：", conn.RemoteAddr().String())

	file, err := os.Create("dest.file")
	defer func() {
		err = file.Close()
	}()
	if err != nil {
		fmt.Println("create error: ", err)
		return
	}

	for {
		// 接收客户端的数据
		bytes := make([]byte, 1024)
		n, err := conn.Read(bytes)

		if err != nil {
			fmt.Println("Read error", err)
			return
		}

		fmt.Println("已经接收到", n, "字节")

		n, err = file.Write(bytes[:n])
		if err != nil {
			fmt.Println("write error: ", err)
			return
		}

	}
}
