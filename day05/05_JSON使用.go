package day05

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"-"`
	Age   int    `json:",string"`
	hobby []string
}

type Person2 struct {
	Name   string `json:"name"`    // 对字段进行二次编码，用在 JSON 的编码和解码
	Age    int    `json:"-"`       // 编码解码时排除该字段
	Gender bool   `json:",string"` // 编码时转换成 string，而不是 bool
	Hobby  []string
}

// JSON 使用
func JSONUsing() {

	// 通过结构体生成
	//structToJSON()

	// 通过 MAP 生成
	//mapToJSON()

	// JSON 转 struct
	//jsonToStruct()

	// JSON 转 map
	JSONToMap()

	//var data interface{}
	//data = "11"
	//GetOriginValue(data)
}

func GetOriginValue(val interface{}) {
	switch res := val.(type) {
	default:
		fmt.Println(res)
	}
}

func jsonToStruct() {
	// 二次编码
	str := `
{
	"name": "hsw",
	"age": 11,
	"gender": true,
	"hobby": ["1", "2", "3"]
}
`
	var v Person
	err := json.Unmarshal([]byte(str), &v)

	if err != nil {
		fmt.Println("解析错误", err)
		return
	}

	fmt.Println("解析成功：", v)
}

func JSONToMap() {
	var str = `
{
	"name": "hsw",
	"age": 11,
	"hobby": ["1", "2", "3"]
}
`

	mmap := map[string]interface{}{}
	err := json.Unmarshal([]byte(str), &mmap)

	if err != nil {
		fmt.Println("出错了", err)
		return
	}

	//fmt.Println(mmap)
	//fmt.Println(mmap["name"])

	//var name string = mmap["name"]
	//var name = mmap["name"] // name is interface{}
	//fmt.Println(name)
	//fmt.Printf("%T\n", name)

	//fmt.Printf("%+v\n",  mmap)

	// 迭代取值

	for key, val := range mmap {
		switch data := val.(type) {
		case string:
			fmt.Printf("%v => %v (string))", key, data)
		case int:
			fmt.Printf("%v => %v (int))", key, data)
		case []string:
			fmt.Printf("%v => %v ([]string))", key, data)
		case []interface{}:
			fmt.Printf("%v => %v ([]interface{}))", key, data)

			fmt.Println(">>>>>>>>>>>>")
			// 继续对 []interface{} 的值进行拆解
			for _, val2 := range data {
				//fmt.Printf("%v %T\n", key2, val2)

				switch data2 := val2.(type) {
				case string:
					fmt.Println("string: => ", data2)
				}
			}

			fmt.Println("<<<<<<<<<<<<<")

		case float64:
			fmt.Printf("%v => %v (float64))", key, data)
		default:
			fmt.Printf("%v => %v (unknown))", key, data)
		}
	}

	//fmt.Println("==============")
	//
	//fmt.Printf("%T\n", mmap["age"])
	//fmt.Printf("%T\n", mmap["hobby"])
	//fmt.Printf("%T\n", mmap["name"])

}

func structToJSON() {
	p := Person{"hsw", 10, []string{"ball", "flower"}}
	r, err := json.Marshal(p)
	if err != nil {
		fmt.Println("marshal 出错了")
	} else {
		fmt.Println(r)
		fmt.Println(string(r)) // bytes 转 string
	}
}

func mapToJSON() {
	var mmap = map[string]interface{}{
		"name":   "hsw",
		"age":    10,
		"person": Person{"hsw", 10, []string{"1", "2", "3"}},
	}

	r, err := json.Marshal(mmap)

	if err != nil {
		fmt.Println("json error: ", err)
		return
	}

	fmt.Println(string(r))
}
