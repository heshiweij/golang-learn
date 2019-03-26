package day03

import "fmt"

// 冒泡排序
// 最简单的排序算法
func Bubbling() {
	// 产生随机数
	var arr [5]int = [5]int{55, 44, 33, 22, 11}

	// 输出
	fmt.Println(arr)

	length := len(arr)

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}

	// 输出结果
	fmt.Println(arr)

}
