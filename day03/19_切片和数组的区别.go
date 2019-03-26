package day03

import "fmt"

// 切片和数组的区别
func SliceAndArrayDiff() {
	var arr [5]int
	fmt.Println(len(arr), cap(arr)) // len: 5; cap: 5

	var slice = []int{}
	fmt.Println(len(slice), cap(slice)) // len: 0; cap: 0

	slice = append(slice, 1)
	fmt.Println(len(slice), cap(slice)) // len: 1; cap: 1

}

// 结论：
//  - 数组的长度和容量是固定
//  - 切片的长度和容量会变的
