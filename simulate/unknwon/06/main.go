package main

import "fmt"

func init() {

}

func main() {
	// 数组定义的格式
	//	- 数组是值类型的一部分
	// 01
	var arr [3]int
	fmt.Println(arr) // 0 0 0

	// 02
	var arr_1 = [...]int{}
	fmt.Println("len = ", len(arr_1)) // 0
	var arr_2 = [...]int{1, 2, 3}
	fmt.Println("len = ", len(arr_2)) // 3

	// 03 使用 new 创建数组的指针
	var arr_3 = new([3]int)
	fmt.Println(arr_3) // &[0 0 0]

	// 数组的初始化
	// 01
	arr2 := [3]int{}
	fmt.Println(arr2) // 0 0 0

	// 02
	arr3 := [3]int{1, 2}
	fmt.Println(arr3) // 1 2 0

	// 数组的长度也是数组的一部分
	/*
	var arr4 [3]int
	var arr5 [4]int

	arr4 = arr5 // wrong
	*/

	// "数组的指针" 和 "指针的数组"
	var arr6 = &[...]int{1, 2, 3,}
	fmt.Println(arr6)                      // arr6 数数组的指针
	fmt.Println(arr6[0], arr6[1], arr6[2]) // 数组的指针可以直接使用 []

	x, y := 1, 2
	var arr7 = [...]*int{&x, &y}
	fmt.Println(arr7) // arr7 是一个数组，里面存放的是指针
	fmt.Println("x = arr7[0] = ", *arr7[0])
	fmt.Println("y = arr7[1] = ", *arr7[1])

	// 数组的比较
	//	- 类型必须一样
	//	- 数组只有 == （所有的值相等）和 !=
	var arr8 = [2]int{1, 2}
	var arr9 = [2]int{1, 2}
	var arr10 = [2]int{3, 4}
	//var arr11 = [3]int{3, 4, 5}

	fmt.Println(arr8 == arr9)
	fmt.Println(arr8 == arr10)
	fmt.Println(arr9 == arr10)
	//fmt.Println(arr11 == arr8) //  wrong: 类型不一致

}
