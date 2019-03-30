package day05

import (
	"fmt"
	"regexp"
)

// 正则表达式
func RegexUsing() {

	buf := "1.2 2.33 2.3 5. 999"

	// 解析正则
	regex := regexp.MustCompile(`\d+\.\d+`)

	// 根据正确查找子串
	res := regex.FindAllStringSubmatch(buf,  -1) // [][]string
	fmt.Println(res)
	//fmt.Printf("%T\n", res)
}


