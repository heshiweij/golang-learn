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

	var path string
	fmt.Println("请输入一个文件路径：")
	_, err = fmt.Scanf("%s", &path)

	if err != nil {
		fmt.Println("scanf error: ", err)
		return
	}

	// 开启一个协程，不断读取文件内容
	go readFile(conn, path)

	for {
		;
	}
}

func readFile(conn net.Conn, path string) {
	file, err := os.Open(path)
	defer func() {
		err = file.Close()
	}()

	if err != nil {
		fmt.Println("file open error: ", err)
		return
	} else {
		fmt.Println("打开文件成功")
	}

	for {

		bytes := make([]byte, 1024)
		n, err := file.Read(bytes)
		if err != nil {
			fmt.Println("read error: ", err)
			return
		}

		n, err = conn.Write(bytes[:n])

		if err != nil {
			fmt.Println("write error: ", err)
			return
		}
	}
}
