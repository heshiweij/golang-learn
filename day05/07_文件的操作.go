package day05

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 文件的操作
func FileUsing() {
	path := "./hello.txt"

	// 写文件
	if false {

		err := WriteFile(path, "heshiwei")

		if err != nil {
			fmt.Println("写入文件发生错误：", err)
		}
	}

	// 读文件
	if false {
		content, err := ReadFile(path)
		if err != nil {
			fmt.Println("读取文件发生错误：", err)
		}

		fmt.Println("内容是：", content)
	}

	// 使用 bufio 按行读取
	if true {
		content, err := BufIOReadFile(path)
		if err != nil {
			fmt.Println("出现错误：", err)
			return
		}

		fmt.Println(content)
	}
}

func ReadFile(path string) (content string, err error) {
	file, err := os.Open(path)

	defer func() {
		err = file.Close()
	}()

	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	bytes := make([]byte, 1024)
	n, err := file.Read(bytes)

	if err != nil && err != io.EOF {
		fmt.Println("error: ", err)
	}

	content = string(bytes)

	fmt.Println("已经成功读取: ", n, "个字节")
	return
}

// 写文件
func WriteFile(path, content string) (err error) {
	file, err := os.Create(path)

	// 发生错误则关闭
	defer func() {
		err = file.Close()
	}()

	if err != nil {
		fmt.Println("error: ", err)
	}

	var n int
	n, err = file.WriteString(content)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println("已经成功写入：", n, "个字节")
	return
}

// 使用 BufIO 包去读取
func BufIOReadFile(path string) (content string, err error) {

	file, err := os.Open(path)

	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	reader := bufio.NewReader(file)

	for {
		var buf []byte
		buf, err = reader.ReadBytes('\n')

		content = content + string(buf)

		if err != nil {
			if err == io.EOF {
				fmt.Println("读取完成")
				err = nil
				break
			}
		}
	}

	return
}
