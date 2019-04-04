package day03

import "fmt"

// 切片的截取
func SliceCutting()  {
	// 1: arr[:]
	arr := [5]int{1,2,3,4,5}
	slice := arr[:]
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice)) // 5 5

	// cap 和原数组的关系
	// 	- cap 是切片对原数组最大能使用的容量，后续添加超过容量，则会重新创建，即使原数组的空间还没有用完
	//arr2 := [100]int{}
	//slice2 := arr2[2:5:11]
	//fmt.Println(slice2)
	//slice2 = append(slice2, 100)
	//fmt.Println("len: ", len(slice2), "cap: ", cap(slice2))
	//slice2[0] = 999
	//fmt.Println(slice2, arr2)

	// low high max, len = high - low, cap = max - low
	//arr := [11]int{1,2,3,4,5,6,7,8,9,11}
	//slice := arr[2:5:6]
	//fmt.Println(len(slice), cap(slice))
	//slice = append(slice, 100)
	//fmt.Println(len(slice), cap(slice))

	// 2. arr[1:]
	arr2 := [5]int{1,2,3,4,5}
	slice2 := arr2[1:] // <==>  1 5 5
	fmt.Println(len(slice2), cap(slice2)) //

	// 3. arr[:2]
	arr3 := [5]int{1,2,3,4,5}
	slice3 := arr3[:3]
	fmt.Println(len(slice3), cap(slice3)) // 0 3 5

	// 4. arr[2:4]
	arr4 := [5]int{1,2,3,4,5}
	slice4 := arr4[2:3] // 2 3 5
	fmt.Println(len(slice4), cap(slice4)) // 1 3

	// 5. arr[2:4:5]
	arr5 := [5]int{1,2,3,4,5}
	slice5 := arr5[2:4:5]
	fmt.Println(len(slice5), cap(slice5)) // 2 3



}

