package balance

import "fmt"

type BalanceMgr struct {
	allBalance map[string]Balance
}

var (
	mgr = BalanceMgr{
		allBalance: map[string]Balance{},
	}
)

func (b *BalanceMgr) registerBalance(name string, balance Balance) {
	b.allBalance[name] = balance
}

// 注册
func RegisterBalance(name string, balance Balance) {
	mgr.registerBalance(name, balance)
}

// 取出 & 选举
func DoBalance(name string, instances []*Instance) (inst *Instance, err error) {
	balance, ok := mgr.allBalance[name]
	if !ok {
		err = fmt.Errorf("not found %s balance", name)
		return
	}

	inst, err = balance.DoBalance(instances)
	return
}
