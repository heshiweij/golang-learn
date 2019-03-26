package day03

import "fmt"

// 数组做函数参数
func ArrayToFuncParams() {
	var arr [2]int = [2]int{2, 3}
	// 输出传递给函数前的值
	fmt.Println(arr)
	// 调用函数，传入数组
	invoke01(arr)
	// 输出传递给函数后的值
	fmt.Println(arr)
}

func invoke01(arr [2]int) {
	arr[0] = 100
	arr[1] = 200
}

// 结论：
//  - 数组传递给函数是值传递
