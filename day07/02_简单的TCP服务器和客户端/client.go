package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8800")

	defer func() {
		err = conn.Close()
	}()

	if err != nil {
		fmt.Println("dial error: ", err)
	}

	n, err := conn.Write([]byte("are you ok?"))

	if err != nil {
		fmt.Println("conn error: ", err)
	}

	fmt.Println("写入", n , "个字节")

}
