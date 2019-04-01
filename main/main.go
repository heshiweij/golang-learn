package main

import (
	"go_learner/day05"
	"go_learner/day07"
	"go_learner/libraries"
)

func main() {

	//funcDay03()
	//
	//funcDay04()
	//
	funcDay05()
	//
	//funcDay06()
	//
	//funcDay07()

	//funcLibraries()
}

func funcLibraries() {
	libraries.TestRedis()
}

func funcDay07() {
	// Socket的实现
	day07.Socket()
}

func funcDay06() {
	// 协程的使用
	//day06.RoutineUsing()

	// 协程资源竞争问题
	//day06.RoutineResCompete()

	// 协程资源竞争问题
	//day06.RoutineChannel()

	// Chan 实现数据的交互
	//day06.ChanDataShare()

	// 有无缓存Chan的区别
	//day06.RoutineChenCache()

	//关闭chan
	//day06.CloseChan()

	// 单向管道
	//day06.UniDirectChan()

	// Timer的使用
	//day06.UsingTimer()

	// Timer的停止和重置
	//day06.TimerStoppingReset()

	// Ticker 的使用
	//day06.UsingTicker()

	// select实现斐波那契
	//day06.UsingSelectFibonacci()

	// select实现超时机制
	//day06.SelectTimeout()
}

func funcDay05() {
	// error接口的使用
	//day05.ErrorUsing()

	// 字符串函数的使用
	//day05.StringFuncUsing()

	// 字符串转换
	//day05.StringConvert()

	// 正则表达式
	//day05.RegexUsing()

	// JSON 使用
	//day05.JSONUsing()

	// 设备文件的操作
	//day05.DeviceFileUsing()

	// 文件的操作
	//day05.FileUsing()

	// 标准输入输出
	day05.StdinStdoutUsing()
}

func funcDay04() {
	// 匿名字段初始化
	//day04.InitializeAnony()

	// 成员的操作
	//day04.OperateAnony()

	// 非结构体匿名字段
	//day04.NoneStructAnonyField()

	// 结构体指针类型匿名字段
	//day04.StructAnonyPointerField()

	// 结构体方法 给基本类型绑定方法
	//day04.StructBasicTypeBindFunc()

	// 结构体方法 给结构体添加绑定方法
	//day04.StructBindFunc();

	// 方法的继承
	//day04.StructFuncInhert()

	// 方法的重写
	//day04.StructFuncRewrite()

	// 方法值和方法表达式
	//day04.StructFuncValueAndExpress()

	// 接口的定义和实现
	//day04.InterfaceDefine()

	// 多态的表现
	//day04.PolyDefine()

	// 接口的继承和转换
	//day04.InterfaceInheritAndTransfer()

	// 空接口
	//day04.InterfaceEmpty()

	// 类型断言
	//day04.InterfaceTypeAssert()
}

func funcDay03() {
	// 随机数的使用
	//day03.RandNumber()

	// 冒泡排序
	//day03.Bubbling()

	// 数组做函数参数是值拷贝
	//day03.ArrayToFuncParams()

	// 数组指针做函数参数
	//day03.ArrayPtrToParams()

	// 切片的长度和容量
	//day03.SliceLengthAndCap()

	// 切片和数组的区别
	//day03.SliceAndArrayDiff()

	// 切片的创建
	//day03.SliceCreating()

	// 切片的截取
	//day03.SliceCutting()

	// 切片和底层数组的关系
	//day03.SliceArrayRelation()

	// append 函数的使用
	//day03.AppendUsing()

	// append扩容特点
	//day03.AppendExpandCap()

	// copy 函数的使用
	//day03.CopyUsing()

	// 切片做函数的参数
	//day03.SliceFuncParams()

	// 猜数字游戏
	//day03.Game()

	// map 基本操作
	//day03.MapBasicOperator()

	// map 作函数参数
	//day03.MapToFuncParams()

	// 结构体初始化
	//day03.StructInitialize()

	// 结构体指针的使用
	//day03.StructPointer()

	// 结构体的比较
	//day03.StructCompare()

	// 结构体的值传递和引用传递
	//day03.StructValueAndReferPass()

	// Go 的可见性规则
	//day03.VisuableRegular()
}
