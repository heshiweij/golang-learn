package day04

import "fmt"

type Person6 struct {
	name string
	age  int
}

// 获得信息
// 	- 值拷贝
func (this Person6) getInfo() (name string, age int) {
	name = this.name
	age = this.age

	// 查看接收到的结构体是否和调用的是同一个
	fmt.Printf("Addr in getinfo: %v\n ", &this)

	return
}

// 设置信息
//	- 地址拷贝
func (this *Person6)setInfo(name string, age int)  {
	this.name = name
	this.age = age
}

// 设置信息
//	- 值拷贝
func (this Person6)setInfoValueCopy(name string, age int)  {
	this.name = name
	this.age = age

	fmt.Println("In setInfoValueCopy: ", this)
}

// 结构体方法 给结构体添加绑定方法
func StructBindFunc() {
	var p = Person6{name: "hsw", age: 100}

	// 查看调用处的地址
	fmt.Printf("Addr in outter: %v\n", &p)

	// 调用值拷贝方法
	name, age := p.getInfo() // p 是值，如果该方法的接收者是指针，则会自动转成指针类型
	fmt.Println(name, age)

	// 调用地址拷贝方法
	p.setInfo("wnm", 10)
	fmt.Println(p)

	// 调用值拷贝方法
	p.setInfoValueCopy("hahaha", 20)
	fmt.Println(p)

}

