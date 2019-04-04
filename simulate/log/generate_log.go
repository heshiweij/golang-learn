package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {

}

type Url struct {
	url    string
	target string
	start  int
	end    int
}

// 构建 URLs 规则集
func (url *Url) ruleUrl() []Url {
	urls := make([]Url, 3)
	r1 := Url{
		url:    "http://127.0.0.1:8800/",
		target: "",
		start:  0,
		end:    0,
	}

	r2 := Url{
		url:    "http://127.0.0.1:8800/list/{$id}.html",
		target: "{$id}",
		start:  1,
		end:    20,
	}

	r3 := Url{
		url:    "http://127.0.0.1:8800/movie/{$id}.html",
		target: "{$id}",
		start:  1,
		end:    65535,
	}

	urls[0] = r1
	urls[1] = r2
	urls[2] = r3

	return urls
}

// 生成 URL 列表
func (url *Url) generateUrlList(rules []Url) (urls []string) {
	if rules == nil || len(rules) <= 0 {
		panic("Urls is empty!")
	}

	urls = make([]string, 0)

	for _, urlRule := range rules {
		if len(urlRule.target) <= 0 {
			urls = append(urls, urlRule.url)
		} else {
			for i := urlRule.start; i <= urlRule.end; i++ {
				url := strings.Replace(urlRule.url, urlRule.target, strconv.Itoa(i), -1)
				urls = append(urls, url)
			}
		}
	}

	return
}

// 生成日志
func makeLog(currentURL, referURL, ua string) (log string) {
	values := url.Values{}
	values.Set("cur", currentURL)
	values.Set("ref", referURL)
	values.Set("ua", referURL)
	queryString := values.Encode()

	template := `
%s - - [02/Apr/2019:11:31:42 +0800] "GET %s HTTP/1.1" 200 612 "-" "%s" "-"`

	ip := makeRandIP()
	return fmt.Sprintf(template, ip, queryString, ua)
}

// 生成随机 UA
func makeRandUserAgent() string {
	templates := [2]string{`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36`,
		`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/11.1.2 Safari/605.1.15`}
	return templates[makeRandNumber(0, len(templates))]
}

// 生成随机 IP
func makeRandIP() string {
	return fmt.Sprint(makeRandNumber(0, 255), ".", makeRandNumber(0, 255), ".", makeRandNumber(0, 255), ".", makeRandNumber(0, 255), ".")
}

// 生成随机数
func makeRandNumber(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if min > max {
		return max
	}

	return r.Intn(max-min) + min
}

// 写入文件
func writeFile(filePath, content string) bool {
	if len(filePath) <= 0 || len(content) <= 0 {
		panic("args is empty")
	}
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	_, err = file.Write([]byte(content))
	if err != nil {
		panic(err)
	}

	return true
}

func main() {

	//if false {
	total := flag.Int("total", 0, "Total count by created")
	filePath := flag.String("filePath", "", "File path by created")
	// 必须要解析才能生效
	flag.Parse()

	//fmt.Println(total, filePath)
	//}

	// check args
	if *total <= 0 {
		panic("Total is empty")
	}

	if len(*filePath) <= 0 {
		panic("File path is empty")
	}

	if true {
		var url = new(Url)
		urls := url.ruleUrl()
		list := url.generateUrlList(urls)

		//
		if false {
			for _, data := range list {
				fmt.Println(data)
			}
		}

		var logStr string

		for i := 0; i < *total; i++ {
			currentURL := list[makeRandNumber(0, len(list))]
			referURL := list[makeRandNumber(0, len(list))]
			ua := makeRandUserAgent()

			logStr = fmt.Sprint(logStr, makeLog(currentURL, referURL, ua))
		}

		// 写入文件
		ok := writeFile(*filePath, logStr)

		if ok {
			fmt.Println("done.")
		}

		//fmt.Println(len(list))
	}
}
