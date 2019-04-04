package balance

type Balance interface {
	// 获取主机
	DoBalance([]*Instance) (*Instance, error)
}

