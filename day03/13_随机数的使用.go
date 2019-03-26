package day03

import (
	"fmt"
	"math/rand"
	"time"
)

// 随机数的使用
//  - 随机种子
func RandNumber() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		// 产生非常大的非负整数
		fmt.Println(rand.Int());
	}

	for i := 0; i < 5; i++ {
		// 自定义范围限定
		fmt.Println(rand.Intn(100));
	}
}
