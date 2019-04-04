package main

import "fmt"

/******** 运算符 ********/

const (
	BYTE = iota + 1
	KB   = BYTE << 10 // 1 * 1024
	MB   = KB << 10   // 1 * 1024 * 1024
	GB   = MB << 10   // 1 * 1024 * 1024
	TB   = GB << 10   // 1 * 1024 * 1024 * 1024
	PB   = TB << 10   // 1 * 1024 * 1024 * 1024
)

// 或
/*
const (
	_ = iota
	BYTE = iota
	KB   = BYTE << 11 // 1 * 1024
	MB   = KB << 11   // 1 * 1024 * 1024
	GB   = MB << 11   // 1 * 1024 * 1024
	TB   = GB << 11   // 1 * 1024 * 1024 * 1024
	PB   = TB << 11   // 1 * 1024 * 1024 * 1024
)
*/

func main() {
	if false {
		fmt.Println(BYTE)
		fmt.Println(KB)
		fmt.Println(MB)
		fmt.Println(GB)
		fmt.Println(TB)
		fmt.Println(PB)
	}

	// 位运算
	if true {
		//a := 6
		//b := 11
		// -----------
		// a = 0110
		// b = 1011
		// ==============
		// | 只要有一个是 1，就是 1
		// & 两个是 1，才是 1
		// ^ 不一样的就是 1
		// &^ 第二个是 1，则强制把第 1 个改成 0

	}
}
