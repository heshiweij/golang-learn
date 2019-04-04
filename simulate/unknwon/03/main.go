package main

import (
	"fmt"
)

func init() {

}

/******* 常量的定义 *********/

const a int = 1

const b = 2

const (
	a1 = 1
	a2 = 2
)

const b1, b2, b3 = 1, 2, 'A'

const (
	c1, c2, c3 = 1, 2, 3
	d1, d2     = 4, 5
	e1         = 6
)

// 定义内部常量
const (
	cMAX = 1 // 无法在其他包使用
)

func main() {
	// 自动赋值
	if false {
		// 01 单个
		const (
			a = 1
			b
			c
		)

		fmt.Println(a, b, c)

		// 02 多个
		const (
			a1, b1 = 1, 2
			a2, b2  // 个数必须和上一行保持一致
		)

		fmt.Println(a2, b2) // a2 = 1,  b2 = 2
	}

	if false {
		// 常量可以是表达式，但表达式（和内置函数）中必须全常量组成
		// 01
		const f1 = a1 + 1

		// 02
		const NAME = "hsw"
		const f2 = len(NAME) // ok

		// 03
		//var str = "hsw"
		//const f3 = len(str) // wrong
	}

	// 常量计数器 iota
	//	- 遇 const ， ++
	//	- 每走一行，++
	if false {
		const (
			a = 1
			b = 2
			f = 22
			c = iota
		)

		fmt.Println(a, b, c)
	}
}
