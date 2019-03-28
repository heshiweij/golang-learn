package day04

import "fmt"

type IHumen interface {
	run()
}

type IPerson interface {
	IHumen
	say()
}

type Worker struct {
}

func (this Worker) run() {
	fmt.Println("run...")
}

func (this Worker) say() {
	fmt.Println("say...")
}

// 接口的继承和转换
func InterfaceInheritAndTransfer() {
	// 01 接口继承
	var work Worker = Worker{}
	var i IHumen =  work
	i.run()

	var ip IPerson = work
	ip.run()
	ip.say()

	// 02 接口的转换
	i = ip // 小 = 大（✔️）
	//ip = i // 大 = 小（✘）
}
