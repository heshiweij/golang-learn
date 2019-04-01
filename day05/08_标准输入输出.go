package day05

import (
	"bufio"
	"fmt"
	"os"
)

// 标准输入输出的使用
func StdinStdoutUsing() {

	// 01 Sprint sprintf sprintln
	var str = fmt.Sprint("1", "2", "3", "hello", "world")
	fmt.Println(str)

	var str2 = fmt.Sprintf("%d-%s-%v", 1, "xixi", 1.2)
	fmt.Println(str2)

	// 02 print printf println
	fmt.Print("1", "2", "3")

	// 03 Fprint
	// 当 write = file
	file, err := os.Create("hello")
	if err != nil {
		fmt.Println("create file error: ", err)
		return
	}
	n, err := fmt.Fprint(file, "1", "2", 3)
	if err != nil {
		fmt.Println("fprint error", err)
		return
	}
	fmt.Println("写入：", n, "个字节")
	// 当 write = stdout
	file = os.Stdout
	n, err = fmt.Fprint(file, "1", "2", "3", 4)
	fmt.Println("写入：", n, "个字节")

	// 04 scanf
	var a, b, c int
	//fmt.Scanf("%d %d %d", &a, &b, &c)
	//fmt.Println(a, b, c)

	// 必须按照模板格式输入
	//fmt.Scanf("输入:%d %d %d", &a, &b, &c)
	//fmt.Println(a, b, c)

	// 多个连续的 scan 变量，如果有一个格式不匹配，则后面的都无法读取
	//fmt.Scan(&a)
	//fmt.Println("a = ", a)
	//fmt.Scan(&b)
	//fmt.Println("b = ", b)
	//fmt.Scan(&c)
	//fmt.Println("c = ", c)

	file, err = os.Open("hello")

	if err != nil {
		fmt.Println("open file error: ", err)
		return
	}

	n, err = fmt.Fscanf(file, "%d %d %d", &a, &b, &c)
	if err != nil {
		fmt.Println("fscanf error: ", err)
		return
	}

	fmt.Println(a, b, c)
	fmt.Println("读取了", n, "个数")

	// Sscanf
	var d, e, f int
	fmt.Sscanf("1 2 3", "%d %d %d", &d, &e, &f)
	fmt.Println(d, e, f)

	for {
		// 05 带缓冲区的 reader
		r := bufio.NewReader(os.Stdin)
		bytes, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("read error: ", err)
			return
		}

		fmt.Println(string(bytes))
	}

}
