package day03

import "fmt"

// 切片和底层数组的关系
func SliceArrayRelation() {
	// 定义数组
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 截取
	slice := arr[2:4:5] // 3 4

	// 再次截取
	slice2 := slice[0:3:3]
	fmt.Println(slice2)
	fmt.Println(len(slice2), cap(slice2))
}

// 结论：
//	- 对切片进行截取时，high、max 不能超过其容量
