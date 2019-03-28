package day04

import "fmt"

// 鸭子接口
type Duck interface {

	// 呱呱叫
	sayGua()
}

// 玩具鸭
type ToyDuck struct {
	name string
}

// 唐老鸭
type DonaldDuck struct {
	name string
}

func (this ToyDuck) sayGua(){
	fmt.Println("toy gua")
}

func (this DonaldDuck) sayGua() {
	fmt.Println("donald gua")
}

// 接口的定义和实现
func InterfaceDefine() {
	// 定义接口变量
	var i Duck

	// 接口接收 ToyDuck
	i = ToyDuck{"toy"}
	i.sayGua()

	// 接口接收 Donald
	i = DonaldDuck{"donald"}
	i.sayGua()

}
