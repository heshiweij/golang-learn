package main

import "fmt"

func init() {

}

type ZT int

func (zt *ZT) add(b int) {
	*zt = *zt + ZT(b)
}

func main() {
	var a ZT = 0
	a.add(100)
	fmt.Println(a)
}
