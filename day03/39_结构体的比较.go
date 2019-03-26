package day03

import "fmt"

// 结构体的比较
func StructCompare() {
	p1 := new(Person)
	p2 := new(Person)

	// 地址不相等
	if p1 == p2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	p1.name = "hsw" // 修改内容，就不相等了

	// 内容相等
	if *p1 == *p2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
