package day03

import "fmt"

// 切片的创建
func SliceCreating()  {
	// 1. 自动类型推倒
	//  - len 和 cap 一样
	//	- 切片 append 后超出了容量，则会自动 2 倍扩容
	slice := []int{1,2,3,4}
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice)) // 4 4

	slice = append(slice, 100)
	fmt.Println(len(slice), cap(slice)) // 5, 8

	// 2. 从数组创建而来
	//	- 如果是从原数组创建，则对切片的修改直接反映到数组
	//	- 对切片进行 append 后，切片底层数组会重新创建，抛弃原数组
	var arr [5]int = [5]int{1,2,3,4,5}
	slice2 := arr[2:3:3]
	fmt.Println(slice2)
	slice2 = append(slice2, 1)
	slice2 = append(slice2, 2)
	slice2 = append(slice2, 3)
	fmt.Println(slice2)

	slice2[0] = 10
	fmt.Println(slice2)
	fmt.Println(arr)

	// 3. 使用 make
	//	- 参数：[]Type len cap
	//	- cap 省略，cap = len
	var slice3 = make([]int, 3)
	fmt.Println(slice3)
	fmt.Println(len(slice3), cap(slice3))

	var slice4 = make([]int, 3,5)
	fmt.Println(slice4)
	fmt.Println(len(slice4), cap(slice4))


}
