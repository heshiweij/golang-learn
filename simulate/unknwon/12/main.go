package main

import (
	"fmt"
	"reflect"
)

func init() {

}

type User struct {
	Id   byte
	Name string
	Age  int16
}

type IHello interface {
	Hello()
}

// String 方法（值拷贝） 可以当做 Java 的 toString() 使用
func (this User) String() string {
	return "hello struct"
}

// String 方法（值拷贝） 可以当做 Java 的 toString() 使用
func (this User) Print(a, b int, c string) string {
	fmt.Println(a, b, c)
	return "hello, I am print method"
}

func (this *User) Hello() {
	fmt.Println("hello world")
}

func main() {

	var user User = User{1, "2", 3}

	// 探究 reflect type 的相关方法
	//Info(user)

	// 一些例子
	example(user)

}

func example(o interface{}) {
	t := reflect.TypeOf(o)

	v := reflect.ValueOf(o)

	// o 不能是 pointer ，这里需要做个判断
	if t.Kind() != reflect.Struct {
		panic("Not struct")
	}

	if false {
		// 取出结构体字段名称和值
		for i := 0; i < t.NumField(); i++ {
			var structField = t.Field(i)

			v := v.Field(i).Interface()
			fmt.Printf("%s %s = %d\n", structField.Name, structField.Type, v)
		}

	}

	if false {
		// 取出方法名称和值
		for i := 0; i < t.NumMethod(); i++ {
			m := t.Method(i)
			fmt.Println(m.Name, m.Type)
		}
	}

	if false {
		// 匿名组合字段的键和值的获取
		type Humen struct {
			Name string
		}

		type Person struct {
			Humen
			Name string
			age  int
		}

		var p = Person{}
		var t = reflect.TypeOf(p)
		// 匿名字段底层是的当成： Humen Humen 来存储的，field name 也是 Humen
		fmt.Println(t.Field(0)) //  {Humen  main.Humen  0 [0] true}
		// 第一个 0：外层字段索引 第二个0：内层字段索引
		fmt.Println(t.FieldByIndex([]int{0, 0})) // {Name  string  0 [0] false}
	}

	if false {
		// 通过 reflect 修改变量的值
		var a int = 1

		// 01
		v := reflect.ValueOf(a)
		fmt.Printf("%T\n", v) // reflect.Value
		v.Elem().SetInt(102)  //  panic: reflect: call of reflect.Value.Elem on int Value
		fmt.Println(a)

		// 02
		v2 := reflect.ValueOf(&a)
		fmt.Printf("%T\n", v2)
		v2.Elem().SetInt(100)
		fmt.Println(a)

	}

	if false {
		// 通过 reflect 修改 结构体字段的值
		var user User = User{
			Id:   1,
			Name: "2",
			Age:  20,
		}

		Set(&user)
		fmt.Printf("%T\n", user)
		fmt.Println(user.Id, user.Name, user.Age)
		fmt.Println(user)
	}

	if true {
		// 方法调用
		var user User = User{
			Id:   1,
			Name: "2",
			Age:  20,
		}
		Invoke(user)
	}
}

func Invoke(o interface{}) {
	v := reflect.ValueOf(o)

	mv := v.MethodByName("Print")

	args := []reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(2),
		reflect.ValueOf("a"),
	}

	res := mv.Call(args)
	fmt.Println(res) // []reflect.Value 由返回值组成的 reflect.Value 切片
	fmt.Println(res[0].String())
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("不能修改")
		return
	}

	v = v.Elem()

	f := v.FieldByName("Name")

	if ! v.IsValid() {
		fmt.Println("BAD")
		return
	}

	if f.Kind() == reflect.String {
		f.SetString("hello")
	}
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)

	if false {
		fmt.Println(t.Name())   // 结构体名字: User
		fmt.Println(t.String()) // main.user
		fmt.Println(t.Align())  // 内存对齐（最长字段占用的内存大小 ）
	}

	if false {
		var a int
		fmt.Println(t.AssignableTo(reflect.TypeOf(a))) // 是否 o 的可以赋值给 a 的变量
	}

	if false {
		// TODO:?
		// Int, Uint, Float, or Complexs
		var a byte
		var t = reflect.TypeOf(a)
		fmt.Println(t.Bits()) // ?
	}

	if false {
		var ch chan string
		var ch1 chan<- string
		var ch2 <-chan string
		var t = reflect.TypeOf(ch)
		var t1 = reflect.TypeOf(ch1)
		var t2 = reflect.TypeOf(ch2)

		fmt.Println(t.ChanDir(), t1.ChanDir(), t2.ChanDir())
	}

	if false {

		// 是否该类型可以使用 ==
		var fun func() string
		var mmap map[int]string
		fmt.Println(reflect.TypeOf(fun).Comparable())
		fmt.Println(reflect.TypeOf(mmap).Comparable())
	}

	if false {
		// 两个值是否可以相互转换
		var a bool
		var b complex64
		fmt.Println(reflect.TypeOf(a).ConvertibleTo(reflect.TypeOf(b)))
	}

	if false {
		// 下面几种类型的元素的类型
		// Array, Chan, Map, Ptr, or Slice.
		//var a map[int]string
		var a chan string
		//var a *int
		//var a string
		var t = reflect.TypeOf(a)
		fmt.Println(t.Elem())
	}

	if false {
		type User struct {
			id   byte `db:"id=1,name=2"`
			Name byte
			Age  int16
		}
		var u User
		t := reflect.TypeOf(u)
		fmt.Println(t.Field(0))

		// StructFiled : ==>  Name 字段名称、pkgName, Type 类型、Tag 标记、offset 位置、Index 对于 FieldByIndex 的位置、Any 是否匿名
		fmt.Printf("%T\n", t.Field(0)) // reflect.StructField
		fieldStruct := t.Field(0)
		fmt.Println(fieldStruct.Name)
		fmt.Println(fieldStruct.PkgPath) // 对于 lower case 的字段（unexported field） 输出包名
		fmt.Println(fieldStruct.Type)    // 任何别名都会输出其最原始的类型
		fmt.Println(fieldStruct.Tag)     // 输出 tag 的内容，做特殊处理
		fmt.Println(fieldStruct.Offset)
		fmt.Println(fieldStruct.Index) // slice
		fmt.Println(fieldStruct.Anonymous)
	}

	if false {
		fmt.Println(t.PkgPath()) // 输出该类型所在的包名
	}

	if false {
		fmt.Println("===")
		type User struct {
			id   byte `db:"id=1,name=2"`
			Name byte
			Age  int64
		}
		var user User
		//var a int
		t := reflect.TypeOf(user)
		// TODO:?
		fmt.Println(t.FieldAlign()) // ? 可能和 alignment 一样
	}

	if false {
		structField := t.FieldByIndex([]int{1}) // 切片只能有  1 个元素，索引超出范围也不行
		fmt.Println(structField)
	}

	if false {
		type User struct {
			id   byte `db:"id=1,name=2"`
			Name byte
			Age  int64
		}
		var user User
		t := reflect.TypeOf(user)
		// 根据字段名查找字段的structField ，大小写敏感，bool 表示是否找到
		fmt.Println(t.FieldByName("Id"))
	}

	if false {
		// TODO:?
		fmt.Println("===")
		structField, bool := t.FieldByNameFunc(func(s string) bool {
			return false
		})

		fmt.Println(structField, bool)
	}

	if false {
		// TODO:?
		var hello IHello = &User{}

		fmt.Println(t.Implements(reflect.TypeOf(hello).Elem()))
	}

	if false {
		fmt.Println("===")
		var fun = func(a, b int, c string, d User, e ...int) {
			fmt.Println(a, b)
		}
		var t = reflect.TypeOf(fun)
		fmt.Println(t.In(0))
		fmt.Println(t.In(1))
		fmt.Println(t.In(2))
		fmt.Println(t.In(3))
		//fmt.Println(t.In(5))

		fmt.Println(t.NumIn()) //  如果 t 是 func，则返回 func 的参数个数

		if false {
			var a int
			var t2 = reflect.TypeOf(a)
			fmt.Println(t2.NumIn()) // wrong: NumIn of non-func type
		}
	}

	if false {
		// 是否函数类型的 Type 是可变参数
		if false {
			var fun = func(a, b int) {

			}
			var t = reflect.TypeOf(fun)
			fmt.Println(t.IsVariadic()) // false
		}

		if false {
			var fun = func(a, b int, c ...int) {

			}
			var t = reflect.TypeOf(fun)
			fmt.Println(t.IsVariadic()) // true
		}

		if false {
			var fun = func(a int, c []int) {

			}
			var t = reflect.TypeOf(fun)
			fmt.Println(t.IsVariadic()) // false

			fmt.Println(t.In(t.NumIn() - 1))
		}

	}

	if false {
		// 返回 map 的 Type 的 key 的类型
		var mmap = map[int]string{}
		t := reflect.TypeOf(mmap)
		fmt.Println(t.Key())
	}

	if false {
		// 数组的len
		fmt.Println("===")
		var arr [10]int
		t := reflect.TypeOf(arr)
		fmt.Println(t.Len())
	}

	if false {
		// 返回的类型的名称
		fmt.Println("===")
		var a int
		var t = reflect.TypeOf(a)
		fmt.Println(t.Kind())
		fmt.Println(t.Kind() == reflect.Int) // true , reflect.Int == 2
		fmt.Println(t.Kind() == 2)           // true
	}

	if false {
		fmt.Println("===")
		fmt.Println(getKind())
	}

	if false {
		fmt.Println(Invalid)
	}

	if false {
		// method 相关（methodName() 和 method() 只认识 非指针绑定的方法）
		fmt.Println(t.NumMethod())
		//fmt.Println(t.MethodByName("Hello"))
		fmt.Println(t.Method(0)) // Method Struct
		method := t.Method(1)    //
		fmt.Println("---")
		fmt.Println(method.Name) // 名称
		fmt.Println(method.PkgPath)
		fmt.Println(method.Index)
		fmt.Println(method.Type)
		fmt.Println("---")

	}

	if false {
		// 返回值列表相关参数：output parameters
		fmt.Println("===")
		var fun = func(a, b int, c string, d User, e ...int) (int, string, User) {
			fmt.Println(a, b)
			return 1, "", User{}
		}

		var t = reflect.TypeOf(fun)
		fmt.Println(t.NumOut())
		fmt.Println(t.Out(0)) // int
	}

	if false {
		// 输出结构体字段的个数
		fmt.Println(t.NumField())

	}
}

type kind uint8

const (
	Invalid kind = iota
	a
	b
)

func getKind() kind {
	return a
}
