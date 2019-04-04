package balance

import "fmt"

type Instance struct {
	// fields
	host   string // 主机地址
	port   int    // 主机端口
	weight byte   // 轮询权重
}

func NewInstance(host string, port int, weight byte) *Instance {
	return &Instance{
		host:   host,
		port:   port,
		weight: weight,
	}
}

// 获取主机地址
func (i *Instance) GetHost() string {
	return i.host
}

// 获取端口
func (i *Instance) GetPort() int {
	return i.port
}

func (i *Instance) String() string {
	return fmt.Sprint(i.host, ":", i.port)
}
