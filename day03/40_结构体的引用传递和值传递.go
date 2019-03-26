package day03

import "fmt"

// 结构体和引用传递和值传递
func StructValueAndReferPass() {
	p := Person{}

	// 值传递
	fmt.Println(p)
	invoke05(p)
	fmt.Println(p)

	// 引用传递
	fmt.Println(p)
	invoke06(&p)
	fmt.Println(p)

}

// 值传递
func invoke05(p Person) {
	p.name = "hsw"
}

func invoke06(p *Person) {
	p.name = "hsw"
}
