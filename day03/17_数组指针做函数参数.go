package day03

import "fmt"

// 数值指针做函数参数
//  - 引用传递
func ArrayPtrToParams() {

	var arr [2]int = [2]int{1, 2}

	invoke02(&arr)

	fmt.Print(arr)
}

func invoke02(arr *[2]int) {
	arr[0] = 100
	arr[1] = 200
}

// 结论：
// 	- 数组通过指针实现引用传递
