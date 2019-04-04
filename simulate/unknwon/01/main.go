package main

import "fmt"

/**
// 批量导入
import (
	"time"
)
 */

/**
// 别名
import (
	T "time"
	. "os" // 省略调用(不能和别名同时使用)
)
 */

func init() {

}

// 常量的定义
const PI = 3.14

// 全局变量的声明和赋值
var name = "gopher"

// 一般类型声明
type NewType int

// 结构的声明
type gopher struct {
}

// 接口的声明
type golang interface {
}

func main() {
	fmt.Println("hello world")
}
