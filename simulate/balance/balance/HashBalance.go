package balance

import "errors"

func init() {
	RegisterBalance("hash", &HashBalance{})
}

type HashBalance struct {
}

func (h *HashBalance) DoBalance(insts []*Instance) (inst *Instance, err error) {
	if len(insts) <= 0 {
		err = errors.New("can't find any instance")
	}

	inst = insts[0]
	return
}
