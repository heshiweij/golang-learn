package day05

import (
	"fmt"
	"strings"
	"unicode"
)

// 字符串函数的使用
func StringFuncUsing() {
	if strings.Contains("heshiwei", "he") {
		fmt.Println("包含")
	}

	if strings.Index("heshiwei", "he") >= 0 {
		fmt.Println("包含")
	}

	result := strings.Join([]string{"1", "2", "3"}, "@")
	fmt.Println(result)

	result = strings.Repeat("bala", 3)
	fmt.Println(result)

	result = strings.Replace("heshiheshiwei", "he", "@", 1)
	fmt.Println(result)

	splits := strings.Split("he@shi@wei", "@")
	fmt.Println(splits)
	printSplits(splits)

	// 去除两倍指定字符串
	// 	- 去除 string 在 cutset 包含的最大的子串
	//	- cutset = " "，去除首位空格
	result = strings.Trim(" he shi wei    ", " ")
	result = strings.Trim("he shi wei", " ")
	fmt.Println(result)

	// 以空格为分隔符，分割字符为切片
	splits = strings.Fields("h e   shi   wei")
	fmt.Println(splits)
	printSplits(splits)

	// 判断是否为空白字符
	isBool := unicode.IsSpace('f')
	fmt.Println(isBool)

}

func printSplits(names []string) {
	if names == nil {
		panic("names 不能为空")
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
