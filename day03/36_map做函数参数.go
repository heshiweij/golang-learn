package day03

import "fmt"

// map 作函数参数
func MapToFuncParams()  {
	var mmap = map[int]string{1:"he", 2:"shi", 3:"wei"}

	invoke04(mmap)

	fmt.Println(mmap)
}

func invoke04(mmap map[int]string) {
	mmap[2] = "fuck"
}
