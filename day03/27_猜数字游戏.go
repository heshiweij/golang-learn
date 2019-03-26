package day03

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 猜数字游戏
func Game() {
	var number int
	var randSlice = make([]int, 4)

	// 生成一个随机四位数
	gameGetRandNumber(&number)

	// 将 4 位数的每一位拆分到长度位 4 的切片
	gameSplitNumberToSlice2(randSlice, &number)

	// 开始游戏
	gameStart(randSlice)
}

// 生成一个随机四位数
func gameGetRandNumber(p *int) {
	rand.Seed(time.Now().UnixNano())

	for {
		*p = rand.Intn(10000)

		if *p > 1000 {
			break
		}
	}
}

// 将 4 位数的每一位拆分到长度位 4 的切片
func gameSplitNumberToSlice(slice []int, number *int) {
	if slice == nil || len(slice) != 4 {
		panic("Slice 长度必须是 4")
	}

	if *number < 1000 || *number > 10000 {
		panic("Number 必须是 4 位")
	}

	slice[0] = *number / 1000
	slice[1] = *number / 100 % 10
	slice[2] = *number / 10 % 10
	slice[3] = *number % 10
}

// 将 4 位数的每一位拆分到长度位 4 的切片(2)
func gameSplitNumberToSlice2(slice []int, number *int) {
	if slice == nil || len(slice) != 4 {
		panic("Slice 长度必须是 4")
	}

	if *number < 1000 || *number > 10000 {
		panic("Number 必须是 4 位")
	}

	str := strconv.Itoa(*number)

	slice[0] = int(str[0] - 48)
	slice[1] = int(str[1] - 48)
	slice[2] = int(str[2] - 48)
	slice[3] = int(str[3] - 48)
}

// 开始游戏
func gameStart(randSlice []int) {

	if randSlice == nil || len(randSlice) != 4 {
		panic("Slice 长度必须是 4")
	}

	count := 0

	for {

		input := 0

		for {
			fmt.Println("请输入一个四位数")
			fmt.Scan(&input)

			if input > 999 && input < 10000 {
				break
			}

			fmt.Println("请重新输入")
		}

		var slice = make([]int, 4)

		// 将 4 位数的每一位拆分到长度位 4 的切片
		gameSplitNumberToSlice2(slice, &input)

		// 正确的位数
		n := 0

		rightIndex := make([]int, 0, 4)

		for i := 0; i < len(slice); i++ {
			position := i + 1
			if randSlice[i] > slice[i] {
				fmt.Printf("第 %d 位小了\n", position)
			} else if randSlice[i] < slice[i] {
				fmt.Printf("第 %d 位大了\n", position)
			} else {
				fmt.Printf("第 %d 位猜对了\n", position)
				rightIndex = append(rightIndex, i)
				n++
			}
		}

		// 如果有 1 位以上猜对，则输出包含已经猜对的数字
		for i, data := range randSlice {

			exists := false

			for _, item := range rightIndex {
				if i == item {
					exists = true
				}
			}

			if exists {
				fmt.Print(data)
			} else {
				fmt.Print("*")
			}
		}
		//fmt.Println()
		//fmt.Println("right index: ", rightIndex)
		//fmt.Println("rand index", randSlice)

		count++


		if n == len(slice) {
			break
		}

	}

	// 公布答案
	fmt.Println("==============")
	fmt.Printf("游戏结束，您一共猜了 %d 次，答案：\n", count)
	for _, data := range randSlice {
		fmt.Print(data)
	}
	fmt.Println()
	fmt.Println("==============")
}
