package day05

import (
	"fmt"
	"os"
)

// 设备文件的使用
func DeviceFileUsing() {
	// 输入
	os.Stdin.Close() // 关闭标准输入后无法输入

	var a int
	fmt.Scanf("%d", &a)
	fmt.Println("a = ", a)

	// 另一种读入方法
	bytes := make([]byte, 3)
	os.Stdin.Read(bytes)

	fmt.Println(bytes)

	// 输出
	os.Stdout.Close()    // 关闭标准输出后无法再输出内容
	fmt.Println("hello") // 调用的就是 os.Stdout.Write() 方法
	os.Stdout.WriteString("hello")
}
