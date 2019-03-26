package day03

import "fmt"

// Map 基本操作
func MapBasicOperator() {
	// 创建
	// - map 没有 cap，只有 len， len 相当于 cap，表示容量，创建时指定，使用( map[n] = m )多少个 len(map) 就是多少
	// - map
	var mmap map[int]string = make(map[int]string, 10)

	mmap[0] = "he"
	mmap[1] = "he"
	mmap[2] = "he"

	fmt.Println(mmap, len(mmap))

	// 初始化
	var mmap2 = map[int]string{1: "he", 2: "shi", 3: "wei"}
	fmt.Println(mmap2)

	// 赋值
	// - 已经存在：覆盖
	// - 不存在：新增
	var mmap3 = map[string]int{"he": 1, "shi": 2}
	mmap3["he"] = 100
	fmt.Println(mmap3)

	// 遍历
	//	- 无序
	var mmap4 = map[int]string{1: "he", 2: "shi", 3: "wei"}
	for key, value := range mmap4 {
		fmt.Printf("key: %d, value: %s\n", key, value)
	}

	// 判断元素是否存在
	var mmap5 = map[int]string{1: "he", 2: "shi", 3: "wei"}
	mmap5[1] = "" // 只能复制 map 指定的值类型
	val, ok := mmap5[1]
	if ok {
		fmt.Println("存在", val)
	} else {
		fmt.Println("不存在", val)
		fmt.Printf("%T\n", val)

		// 不存在则返回的是值类型的默认类型
		if val == "" {
			fmt.Println("就是个空字符串")
		}
	}


	// 删除元素
	var mmap6 = map[int]string{1: "he", 2: "shi", 3: "wei"}
	delete(mmap6, 1)
	fmt.Println(mmap6)
	_, ok = mmap6[1]
	if ok {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}
}
