package day05

import (
	"fmt"
	"strconv"
)

// 字符串转换
func StringConvert() {

	// 01 将各种类型的值转换为 string，再转成 byte 追加到 []byte 中
	bytes := make([]byte, 0);
	bytes = strconv.AppendBool(bytes, true)
	bytes = strconv.AppendBool(bytes, false)

	fmt.Println(bytes, string(bytes))

	// 02 将各种类型转换成字符串
	var result string
	result = strconv.FormatFloat(1.2, 'f', 2, 32)
	fmt.Println(result)

	result = strconv.FormatBool(true)
	fmt.Println(result)

	result = strconv.FormatInt(10, 10)
	fmt.Println(result)

	result = strconv.FormatUint(10, 10)
	fmt.Println(result)

	// 整数转字符串 (itoa  atoa 相对，是 int 和 string 想换转换一堆方法)
	result = strconv.Itoa(10)
	fmt.Println("i2a: ", result)

	// 03 字符串可以通过 append 追加到 bytes 切片里
	//	- string... 可以转换为 []bytes，仅限于方法参数（形参和实参）
	bytes = append(bytes, strconv.FormatBool(true)...)

	// 对于中文汉字，会进行拆分，用 utf8 方式表示
	names := "何世威"
	bytes2 := append([]byte{}, names...)
	fmt.Println("===")
	fmt.Println(bytes2)

	// 03 字符串转各类型

	// string -> bool
	var (
		r   interface{}
		err error
	)
	r, err = strconv.ParseBool("true2")
	if err == nil {
		fmt.Println(r)
	} else {
		fmt.Println("convert error")
	}

	// string -> int
	r, err = strconv.ParseInt("100", 10, 32) // == ParseInt("100", 10, 0)

	if err != nil {
		fmt.Println("convert error")
	} else {
		fmt.Println(r)
	}

	// string -> float
	r, err = strconv.ParseFloat("1.2", 32)
	if err != nil {
		fmt.Println("convert error")
	} else {
		fmt.Println(r)
	}

	r, err = strconv.Atoi("10")
	if err != nil {
		fmt.Println("convert error")
	} else {
		fmt.Println(r)
	}

}
