package main

import (
	"fmt"
	"net"
	"strings"
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

	for {
		// 接收客户端的数据
		bytes := make([]byte, 1024)
		n, err := conn.Read(bytes)

		if err != nil {
			fmt.Println("Read error", err)
			return
		}

		clientRepy := string(bytes[:n])

		// 去除最后的 \n
		if string(bytes[:n-1]) == "exit" {
			err = conn.Close()
			if err != nil {
				fmt.Println("close conn error: ", err)
			}
			fmt.Println("客户端：", conn.RemoteAddr().String(), "已经退出")
			return
		}

		n, err = conn.Write([]byte(strings.ToUpper(clientRepy)))

		fmt.Println("客户端的数据：", clientRepy)
	}
}
