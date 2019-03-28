package day04

import "fmt"

// 类型断言
func InterfaceTypeAssert() {
	// 01 if else
	slice := make([]interface{}, 3)
	slice[0] = 1
	slice[1] = "string"
	slice[2] = Person10{"hello"}

	//value , ok := slice[0].(int)
	//fmt.Println(value, ok)

	// 通过 value,ok := data.(type) 断言类型
	for i, data := range slice {
		if value, ok := data.(int); ok {
			fmt.Printf("slice[%d] is int, The value is %v \n", i, value)
		} else if value, ok := data.(string); ok {
			fmt.Printf("slice[%d] is string, The value is %v \n", i, value)
		} else if value, ok := data.(Person10); ok {
			fmt.Printf("slice[%d] is struct, The value is %v \n", i, value)
		}
	}

	// 02 switch
	for _, data := range slice {
		switch val := data.(type) {
		case int:
			fmt.Println(val)
		case string:
			fmt.Println(val)
		case Person10:
			fmt.Println(val)
		}
	}

}
