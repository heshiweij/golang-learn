package day03

import "fmt"

// 结构体指针的使用
func StructPointer() {
	var p Person
	p.id = 1
	p.name = "he"
	p.sex = 0
	p.addr = "sy"

	var pp *Person = &p
	fmt.Println(pp.id, pp.name, pp.sex, pp.addr) // pp.name 和 (*pp).name 相同

	// 使用 new 初始化
	pp2 := new(Person)
	fmt.Println(pp2.id, pp2.name, pp2.addr, pp2.sex)

}
