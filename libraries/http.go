package libraries

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"html/template"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"
)

func Hello(w http.ResponseWriter, r *http.Request) {

	incr()
	cnt := getCount()

	_, err := fmt.Fprint(w, "hello -> ", cnt)

	if err != nil {
		fmt.Println("http write error: ", err)
		return
	}

	fmt.Println("http write success")
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "user/login")

	if err != nil {
		fmt.Println("http write error: ", err)
		return
	}

	fmt.Println("http write success")
}

func incr() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}

	c.Do("incr", "count")
}

func getCount() int {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}

	count, err := redis.Int(c.Do("get", "count"))
	return count
}

// http 服务端
func testServer() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/user/login", UserLogin)
	http.HandleFunc("/user/login/1", UserLogin)

	err := http.ListenAndServe("0.0.0.0:8800", nil)

	if err != nil {
		fmt.Println("listen and serve error: ", err)
		return
	}

	fmt.Println("listen at *:8800")
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	//panic("恐怖了")
	_, err := io.WriteString(w, "hello-world")
	if err != nil {
		panic(err)
	}
}

// http 全局异常处理、日志
func testPanic() {
	http.HandleFunc("/", logPanic(HelloWorld))

	err := http.ListenAndServe("0.0.0.0:8800", nil)

	if err != nil {
		fmt.Println("listen and serve error: ", err)
		return
	}

	fmt.Println("listen at *:8800")
}

// 全局错误处理 01
/*
func logPanic(handle func(response http.ResponseWriter, r *http.Request)) func(response http.ResponseWriter, r *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			// 捕捉异常
			if r := recover(); r != nil {
				fmt.Println("catch exception: ", r)
			}
		}()

		handle(writer, request)
	}
}
*/

// 全局错误处理 02 精简版
func logPanic(handle http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			// 捕捉异常
			if r := recover(); r != nil {
				fmt.Println("catch exception: ", r)
			}
		}()

		handle(writer, request)
	}
}

// http 表单
func testForm() {
	html := `
<html>
<head></head>
<body>
	<form action="" method="post">
		<input type="text" name="in" />
		<input type="text" name="in" />
		<input type="text" name="out" />
		<input type="submit" name="提交" />
	</form>
</body>
</html>
`
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		var r string
		switch request.Method {
		case "GET":
			// 设置 header 头：content-type: text/html
			writer.Header().Set("Content-Type", "text/html")
			r = html
		case "POST":
			err := request.ParseForm()
			if err != nil {
				panic(err)
			}

			//fmt.Println(request.Form["in"])

			in0 := request.Form["in"][0]
			in1 := request.Form["in"][1]

			r = fmt.Sprint(r, in0, in1)

			r = fmt.Sprint(r, request.FormValue("in"))

			r = fmt.Sprint(r, request.FormValue("out"))
		}

		_, err := io.WriteString(writer, r)
		if err != nil {
			panic(err)
		}
	})

	err := http.ListenAndServe("0.0.0.0:8800", nil)

	if err != nil {
		fmt.Println("listen and serve error: ", err)
		return
	}

	fmt.Println("listen at *:8800")
}

type Person struct {
	run func()
}

type Good struct {
	Name    string
	Age     int
	Address string
}

func testTemplate() {
	t, err := template.ParseFiles("/Users/heshiwei/projects/GoProjects/src/go_learner/libraries/template/index.html")
	if err != nil {
		panic(err)
	}

	good := Good{
		Name:    "Hsw",
		Age:     11,
		Address: "浙江上虞",
	}

	// 特点：
	//	- 会把非 string 字段的转成 string
	//	- {{.}} 会打印整个结构体
	err = t.Execute(os.Stdout, good)
	if err != nil {
		panic(err)
	}
}

func testClient() {
	//testGet()
	//testHead()
	//testTimeout()
}

// 超时
func testTimeout() {
	var urls = []string{
		"http://www.baidu.com",
		"http://www.facebook.com",
		"http://www.163.com",
	}

	// 构造支持超时 http 客户端
	c := http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (conn net.Conn, e error) {
				timeout := time.Second * 1
				return net.DialTimeout(network, addr, timeout)
			},
		},
	}

	for _, url := range urls {
		//fmt.Println(url)

		resp, err := c.Head(url)
		if err != nil {
			panic(err)
		}

		fmt.Println(resp.StatusCode, resp.Status)
	}

}

func testGet() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}

	if false {
		// 打印 resp 的各种信息
		fmt.Println(resp)
		fmt.Println(resp.Status)
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Proto)
		fmt.Println(resp.ProtoMajor) // 协议主版本号
		fmt.Println(resp.ProtoMinor) // 协议最小版本号
		fmt.Println(resp.Header)     // map[string][]string
		fmt.Println(resp.Body)       // io.ReadCloser
		fmt.Println(resp.ContentLength)
		fmt.Println(resp.TransferEncoding)
		fmt.Println(resp.Close)
		fmt.Println(resp.Uncompressed)
		fmt.Println(resp.Request)
	}

	// 输出 header

	if false {
		// 01
		var arr = []string{"1", "@"}
		fmt.Println(arr)

		// 02
		header := resp.Header
		for key, val := range header {
			fmt.Println(key, val)

			for _, v := range val {
				fmt.Println("    VAL: ==> ", v)
			}
		}
	}

	// 输出 body
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil read all error: ", err)
		return
	}
	fmt.Println(string(bytes))
}

func TestHttp() {
	// http 服务端
	//testServer()

	// http 客户端
	//testClient()

	// http 表单
	//testForm()

	// http 全局异常处理、日志
	//testPanic()

	// text 模板
	//testTemplate()

	// http text 模板
	testHttpTemplate()
}

var myTemplate template.Template

func testHttpTemplate() {

	// 初始化模板
	myTemplate, err := template.ParseFiles("/Users/heshiwei/projects/GoProjects/src/go_learner/libraries/template/index.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", logPanic(func(writer http.ResponseWriter, request *http.Request) {

		// 01 使用 struct
		if false {
			good := Good{
				Name:    "wnm",
				Age:     10,
				Address: "广东深圳",
			}

			err = myTemplate.Execute(writer, good)
		}

		// 02 使用 map
		if true {
			slice := make([]string, 3)
			slice[0] = "slice ele 01"
			slice[1] = "slice ele 02"
			slice[2] = "slice ele 03"
			mmap := map[string]interface{}{
				"name":          "hsw",
				"age":           "199",
				"address":       "嘻嘻",
				"elements":      slice,
				"element_empty": nil,
			}

			err = myTemplate.Execute(writer, mmap)
		}

		if err != nil {
			panic(err)
		}

	}))

	err = http.ListenAndServe("0.0.0.0:8800", nil)

	if err != nil {
		fmt.Println("listen and serve error: ", err)
		return
	}

	fmt.Println("listen at *:8800")
}
