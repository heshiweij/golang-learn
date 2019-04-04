package main

import "fmt"

/*********** 变量定义 *************/

// 批量定义常量
const (
	NAME = 1
	ADDR = 2
)

// 批量定义变量
var (
	name = "string"
	addr = "sy"
)

// 批量定义类型别名
type (
	newInt int
	newFloat float64
	newByte byte
)

func main() {

	if false {
		// 各类型 0 值
		var int_ int
		var int_8 int8
		var int_16 int16
		var int_32 int32
		var int_64 int64

		var uint_ uint
		var uint_8 uint8
		var uint_16 uint16
		var uint_32 uint32
		var uint_64 uint64

		fmt.Println(int_, int_8, int_16, int_32, int_64)
		fmt.Println(uint_, uint_8, uint_16, uint_32, uint_64)

		var float_32 float32
		var float_64 float64

		fmt.Println(float_32, float_64)

		var bool_ bool
		var string_ string
		fmt.Println(bool_, string_)
	}

	if true {
		convert()
	}

}

/*********** 变量赋值 *************/

// 全局变量
var a = 1

var a2 = 1

var (
	a3 = 1
	a4 = 2
	a5 = 3
)

func test() {
	// 局部变量
	var b = 1

	var b1 int = 1

	var (
		b2 = 1
		b3 = 2
		b4 = 3
	)

	// 垃圾桶
	_, _, _, _, _ = b, b1, b2, b3, b4 // 忽略所有

	_, a = b, b1 // 忽略 b1

	fmt.Println(b, b1, b2, b3, b4)

}

/*********** 变量转换 *************/

// 1、uint[n]、int[n] 以及相关的别名（byte rune） 之间可以相互转换
// 2、bool
// 3、string 和 []uint8 以及 uint8 的别名切片（[]byte 等）可以相互转换
// 4、string 和 []int32 以及 int32 的别名切片（[]rune 等）可以相互转换
// 5、string 和 uint[n]、int[n] 以及相关的别名可以单向转换（int -> string）
// 5、隐式转换特例 int32 <-> rune、byte <-> uint8；其他的即使是别名也无法转换
func convert() {
	if false {
		// uint uint8 uint16 uint32 uint64 -> int ✅
		var val int = 1
		fmt.Println(int(val))
	}

	if false {
		// uint uint8 uint16 uint32 uint64 -> int ✅
		var val int = 1
		fmt.Println(int8(val))
	}

	if false {
		// uint uint8 uint16 uint32 uint64 -> int16 ✅
		var val int = 1
		fmt.Println(int16(val))
	}

	if false {
		// uint uint8 uint16 uint32 uint64 -> int32 ✅
		var val int = 1
		fmt.Println(int32(val))
	}

	if false {
		// uint uint8 uint16 uint32 uint64 -> int64 ✅
		var val int = 1
		fmt.Println(int64(val))
	}

	if false {
		// uint uint8 uint16 uint32 uint64 -> uint ✅
		var val int = 1
		fmt.Println(uint(val))
	}

	if false {
		// uint uint8 uint16 uint32 uint64 -> uint8 ✅
		var val int = 1
		fmt.Println(uint8(val))
	}

	if false {
		// uint uint8 uint16 uint32 uint64 -> uint16 ✅
		var val int = 1
		fmt.Println(uint16(val))
	}

	if false {
		// uint uint8 uint16 uint32 uint64 -> uint32 ✅
		var val int = 1
		fmt.Println(uint32(val))
	}

	if false {
		// uint uint8 uint16 uint32 uint64 -> uint64 ✅
		var val int = 1
		fmt.Println(uint64(val))
	}

	if false {
		var val string = "string"
		fmt.Println([]uint8(val))
		fmt.Println([]byte(val))

		type myUint8 uint8
		fmt.Println([]myUint8(val))

		type myByte byte
		fmt.Println([]myByte(val))
	}

	if false {
		var val string = "string和"
		fmt.Println([]rune(val))  // [115 116 114 105 110 103 21644]
		fmt.Println([]byte(val))  // [115 116 114 105 110 103 229 146 140]
		fmt.Println([]int32(val)) // [115 116 114 105 110 103 21644]

		type myRune rune
		fmt.Println([]myRune(val)) // [115 116 114 105 110 103 21644]
	}

	// 隐式转换特例 int32 <-> rune、byte <-> uint8
	if false {
		var a rune = 1
		var b int32 = 2
		a = b

		var a1 byte = 1
		var b1 uint8 = 2
		a1 = b1

		_, _ = a, a1
	}

	if false {
		var a int64 = 65
		var str string = string(a)

		type myint int
		var b myint = 65
		var str2 string = string(b)

		fmt.Println(str, str2)
	}

}
