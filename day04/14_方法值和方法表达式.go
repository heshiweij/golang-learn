package day04

import "fmt"

type Person9 struct {
	name string
}

// 绑定方法
func (this Person9) getInfo() {
	fmt.Println("hello person: ", this.name)

	this.name = "wnm"
}

// 方法值和方法表达式
func StructFuncValueAndExpress() {
	// 方法值
	//structFuncValue()

	// 方法表达式
	structFuncExpress()
}

// 方法值
func structFuncValue() {
	// 直接调用
	var p Person9 = Person9{name:"hsw"}
	p.getInfo()

	// 赋值为方法值后再调用：保留接收者
	myFunc :=  p.getInfo
	myFunc()
}

// 方法表达式
func structFuncExpress() {
	var p Person9 = Person9{name:"hsw"}

	// 通过 (*结构体).方法() 去调用，需要手动传入接收者
	f1 := (*Person9).getInfo // 01
	f1(&p)

	f2 := (Person9).getInfo // 02. 因为 getInfo 是指针接收者，因此无法这样调用
	f2(p)

	//fmt.Println(p)

}


