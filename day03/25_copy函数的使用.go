package day03

import "fmt"

// copy 函数的使用
func CopyUsing() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{6, 6, 6, 6, 6, 6}

	// 输出原数组
	fmt.Println(len(slice1), cap(slice1))
	fmt.Println(len(slice2), cap(slice2))
	fmt.Println(slice1, slice2)

	// 拷贝
	copy(slice1, slice2)

	// 输出从 2 拷贝到 1 的结果
	fmt.Println(len(slice1), cap(slice1))
	fmt.Println(len(slice2), cap(slice2))

	fmt.Println(slice1, slice2)

	// 拷贝
	copy(slice2, slice1)

	// 输出从 1 拷贝到 2 的结果
	fmt.Println(len(slice1), cap(slice1))
	fmt.Println(len(slice2), cap(slice2))
	fmt.Println(slice1, slice2)

	// 即使 dst 切片的容量足够，拷贝也不会被使用 cap 的元素
	slice3 := make([]int, 3, 5)
	slice3[0] = 1
	slice3[1] = 2
	slice3[2] = 3
	copy(slice3, slice2)
	fmt.Println(len(slice3), cap(slice3))


}

// 结论：
//	- 拷贝不会改变原切片的 len 和 cap
//	- 有多少拷贝多少，dst 容量不足，就不拷贝