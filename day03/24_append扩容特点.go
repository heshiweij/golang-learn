package day03

import "fmt"

// append 扩容特点
func AppendExpandCap() {
	slice := make([]int, 1, 1)

	for i := 0; i < 100; i ++ {
		slice = append(slice, i)
		fmt.Println(len(slice), cap(slice))
	}
}

// 结论：
//	- 两倍扩容
