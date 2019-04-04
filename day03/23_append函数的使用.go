package day03

import "fmt"

// append 函数的使用
func AppendUsing()  {
	slice := make([]int, 5, 10)

	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)

	fmt.Println(len(slice), cap(slice)) // 8 11

}
