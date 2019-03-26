package day03

import "fmt"

// 切片的长度和容量
func SliceLengthAndCap()  {
	var arr [5]int = [5]int{1,2,3,4,5}
	var slice []int =  arr[1:3:5] // arr[low:high:max]

	// 输出切片：low 和 high 之间的内容
	fmt.Println(slice)

	fmt.Println("len: ", len(slice)) // len = high - low = 2
	fmt.Println("cap: ", cap(slice)) // cap = max - low = 4
}

// 结论：
//  - 切片是一种数据结构，是对数组的引用，内部保存了：原数组指针、长度、容量
//  - 长度和容量的计算公式：
//  	len = high - low
//      cap = max - low
//  - [low:high:max] low 可以取， high 不取


