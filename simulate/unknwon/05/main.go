package main

import (
	"fmt"
)

func init() {

}

func main() {
	// if
	// 01
	if true {
		fmt.Println("true")
	}

	// 02
	var a = 1
	if a = 2; true {
		fmt.Println(a) // 2
	}

	// for
	// 01
	/*
	for {
		;
	}
	*/

	// 02
	i := 0
	for i <= 3 {
		i++
		fmt.Println("i = ", i)
	}

	// 03
	for i := 0; i < 3; i++ {
		fmt.Println("i = ", i)
	}

	// 04
	for i = 0; i < 3; {
		i ++
		fmt.Println("i = ", i)
	}

	// switch
	// 01
	var a1 = 1
	var b = 1
	switch a1 { // 任意表达式
	case b: // 任意表达式
		fmt.Println("a = 1")
	case 2:
		fmt.Println("a = 2")
	case 3:
		fmt.Println("a = 3")
	}

	// 02
	var a2 = 1
	switch {
	case a2 > 1: // 必须是 bool
		fmt.Println("a > 1")
	case a2 > 10:
		fmt.Println("a > 11")
	}

	switch a2 = 1; { // 初始化（本质上和上面的一样）
	case a2 > 1: // 必须是 bool
		fmt.Println("a > 1")
	case a2 > 10:
		fmt.Println("a > 11")
	}

	var a3 = 1
	switch a3 {
	case 1:
		fmt.Println(a3)
		fallthrough // 继续下一个 case
	case 2:
		fmt.Println(a3)
	}

}
