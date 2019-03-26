package day03

import (
	"fmt"
	"math/rand"
	"time"
)

// 切片做函数的参数
func SliceFuncParams() {
	// 01
	slice := []int{1, 2}

	fmt.Println(slice)

	invoke03(slice)

	fmt.Println(slice)

	// 02
	var slice2 = make([]int, 5)

	// 产生随机数
	generateSliceRand(slice2)

	fmt.Println("排序前: ", slice2)

	// 排序
	bubblingSortSlice(slice2)

	// 输出结果
	fmt.Println("排序后: ", slice2)
	fmt.Println(slice2)
}

func invoke03(slice []int) {
	slice[0], slice[1] = slice[1], slice[0]
}

// 产生切片随机数据
func generateSliceRand(slice []int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(100)
	}
}

// 切片冒泡排序
func bubblingSortSlice(slice []int) {
	length := len(slice)

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

// 结论：
//	- 切片作为函数的参数时，引用传递
