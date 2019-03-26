package day03

import "fmt"

type Person struct {
	id   int
	name string
	sex  byte
	addr string
}

// 结构体初始化
func StructInitialize() {
	// 01
	// 为赋值，默认为空
	var p Person
	fmt.Println(p)

	// 02
	var p2 = Person{
		1,
		"2",
		10,
		"fs", // 如果是竖着写，最后一个必须加 ","
	}

	fmt.Println(p2)

	// 03
	var p3 = Person{1, "fs", 2, "fs"}
	fmt.Println(p3)

	// 04
	var p4 = Person{id: 242, name: "fs"} // 带 key 不带 key 必须统一
	fmt.Println(p4)
	fmt.Println(p4.id, p4.name, p4.sex, p4.addr)
}
