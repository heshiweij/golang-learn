package main

import "fmt"

func init() {

}

func main() {

	if false {
		// 创建切片
		// 01
		var slice []int = []int{1, 2, 3}
		fmt.Println(slice, len(slice), cap(slice)) // 3 3

		// 02
		var slice1 = make([]int, 3)
		fmt.Println(slice1, len(slice1), cap(slice1)) // 3 3
	}

	if false {
		// reslice
		var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		// arr[start:end:max] 且 [start, end)
		// len = end - start
		// cap = max - start
		s1 := arr[2:]                     // 3 ~ 11，end 不写，则 end = len(arr), max =  end
		fmt.Println(s1, len(s1), cap(s1)) // len = 8 cap = 8

		s2 := arr[:5]                     // 1 ~ 5, start 不写，则 start = 0, max = 11
		fmt.Println(s2, len(s2), cap(s2)) // len =5 cap = 11

		//s3 := arr[0:11:11]
		//fmt.Println(s3)
	}

	if false {
		// append
		var slice = make([]int, 2, 5)
		slice = append(slice, 10)
		fmt.Printf("%p\n", slice)
		slice = append(slice, 10)
		fmt.Printf("%p\n", slice)
		slice = append(slice, 10)
		fmt.Printf("%p\n", slice)

		slice = append(slice, 10) // 实际元素个数超过 cap 容量，则重新分配
		fmt.Printf("%p\n", slice)

		// 输出
		/**
			0xc000018120
			0xc000018120
			0xc000018120
			0xc00001c0a0
		 */

	}

	if true {
		var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		var s1 = arr[2:3]
		var s2 = arr[3:4]
		var s3 = arr[2:8]

		// 打印的不是 s1 变量的地址，而是 s1 切片结构体中保存的实际底层数组的指针变量的值
		// 	- 该值指针就是指向的底层数组的首地址, 也就是切片 start 的地址
		fmt.Printf("%p\n", s1) // 0xc00008e010
		fmt.Printf("%p\n", s2) // 0xc00008e018  相差 8 byte ，就是一个 int64 == int(当前 64 位系统下)
		fmt.Printf("%p\n", s3) // 0xc000092010 和第一个一样
	}

}
