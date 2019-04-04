package main

import (
	"fmt"
	"sort"
)

func init() {

}

func main() {
	// map 的初始化
	if false {
		// 01
		var mmap = make(map[int]string)
		mmap[0] = "s"
		fmt.Println(mmap)

		// 02
		var mmap2 = map[int]string{
			1: "hello",
		}
		fmt.Println(mmap2)
	}

	// map 的排序
	if false {
		var mmap = make(map[int]string)
		mmap[1] = "1"
		mmap[2] = "2"
		mmap[3] = "3"
		mmap[4] = "4"
		mmap[5] = "5"

		/*
	2 2
	3 3
	4 4
	5 5
	1 1
	 */
		for k, v := range mmap {
			fmt.Println(k, v)
		}

		var s = make([]int, len(mmap))
		i := 0
		for k, _ := range mmap {
			s[i] = k
			i++
		}

		sort.Ints(s)

		/*
	1
	2
	3
	4
	5
	 */
		for _, k := range s {
			fmt.Println(mmap[k])
		}

	}

	// 将 map[int]string -> map[string]int
	if true {
		mapInt := make(map[int]string)
		mapStr := make(map[string]int)

		mapInt[1] = "str1"
		mapInt[2] = "str2"
		mapInt[3] = "str3"
		mapInt[4] = "str4"
		mapInt[5] = "str5"

		for k, v := range mapInt {
			mapStr[v] = k
		}

		fmt.Println(mapStr)
	}

}

