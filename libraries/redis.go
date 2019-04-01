package libraries

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

const (
	host = "127.0.0.1"
	port = 6379
)

func TestRedis() {

	conn, err := redis.Dial("tcp", host+":"+strconv.Itoa(port))

	if err != nil || conn == nil {
		fmt.Println("conn error: ", err)
		return
	}

	//_, err = conn.Do("get", "name")
	//
	//if err != nil {
	//	fmt.Println("do error: ", err)
	//	return
	//}

	// 执行到这里，说明 redis conn、do 正常，才有关闭的必要
	defer func() {
		err = conn.Close()
		if err != nil {
			fmt.Println("conn close error: ", err)
		}
	}()

	//testByteSlice(reply)
	//testUint8Slice(reply)
	//testReplyDetail(reply)
	//showReplayString(reply)

	//writeForeverValue(conn)
	//writeExpireValue(conn)

	//multiGet(conn)
	//multiSet(conn)
	//existsValue(conn)
	//expireValue(conn)

	//pipline(conn)

	trans(conn)
}

func trans(conn redis.Conn) {
	err := conn.Send("multi")
	if err != nil {
		fmt.Println("multi error: ", err)
		return
	}
	err = conn.Send("incr", "count")
	if err != nil {
		fmt.Println("incr count error: ", err)
		return
	}
	err = conn.Send("incr", "aaa")
	if err != nil {
		fmt.Println("incr aaa error: ", err)
		return
	}
	_, err = conn.Do("exec")
	if err != nil {
		fmt.Println("exec error: ", err)
		return
	}
	fmt.Println("ok")
}

func pipline(conn redis.Conn) {
	err := conn.Send("set", "name", "hsw2")
	if err != nil {
		fmt.Println("set error: ", err)
		return
	}
	err = conn.Send("get", "name")
	if err != nil {
		fmt.Println("get error: ", err)
		return
	}
	err = conn.Flush()
	if err != nil {
		fmt.Println("flush error: ", err)
		return
	}

	funcFor(2, func(i int) bool {
		fmt.Println("这是第", i, "次")
		reply, err := conn.Receive()
		if err != nil {
			fmt.Println("receive error: ", err)
			return false
		}

		fmt.Println(reply)

		return true
	})
}

func funcFor(n int, fun func(index int) bool) {
	res := true
	for i := 0; i < n; i ++ {
		res = fun(i)
		if ! res {
			break
		}
	}
}

func expireValue(conn redis.Conn) {
	reply, err := redis.Bool(conn.Do("expire", "name2", 10))
	if err != nil {
		fmt.Println("write error: ", err)
		return
	}

	if reply {
		fmt.Println("设置成功")
	} else {
		fmt.Println("设置失败")
	}

}

func existsValue(conn redis.Conn) {
	reply, err := redis.Bool(conn.Do("exists", "name1"))
	if err != nil {
		fmt.Println("write error: ", err)
		return
	}

	fmt.Println(reply)
	fmt.Printf("%T\n", reply)

	if reply {
		fmt.Println("exists")
	} else {
		fmt.Println("no exists")
	}
}

func multiSet(conn redis.Conn) {
	reply, err := redis.String(conn.Do("mset", "name1", "hsw11", "name2", "hsw22"))
	if err != nil {
		fmt.Println("write error: ", err)
		return
	}

	fmt.Println(reply)
}

func multiGet(conn redis.Conn) {
	reply, err := redis.Strings(conn.Do("mget", "name1", "name2", "name3"))
	if err != nil {
		fmt.Println("write error: ", err)
		return
	}

	//fmt.Println(reply)

	for i, data := range reply {
		fmt.Println("i = ", i, " data = ", data)
	}
}

func writeExpireValue(conn redis.Conn) {
	_, err := redis.String(conn.Do("set", "name", "wnm2", "ex", 10))
	if err != nil {
		fmt.Println("write error: ", err)
		return
	}
}

func writeForeverValue(conn redis.Conn) {
	// 写入值不过期
	_, err := redis.String(conn.Do("set", "name", "wnm2"))

	if err != nil {
		fmt.Println("write error: ", err)
		return
	}
}

func showReplayString(reply interface{}, err error) {
	//show reply string
	// 01
	if value, ok := reply.([]byte); ok {
		fmt.Println(string(value))
	}

	// 02 using redis.String()
	reply, err = redis.String(reply, err)
	fmt.Println(reply, err)
}

func testReplyDetail(reply interface{}) {
	fmt.Printf("%T %v\n", reply, reply)
	fmt.Println(reply)
}

func testByteSlice(reply interface{}) {
	if value, ok := reply.([]byte); ok {
		fmt.Println("[]byte", value)
	} else {
		fmt.Println("not []uint8")
	}
}

func testUint8Slice(reply interface{}) {
	if value, ok := reply.([]uint8); ok {
		fmt.Println("[]uint8", value)
	} else {
		fmt.Println("not []uint8")
	}
}
